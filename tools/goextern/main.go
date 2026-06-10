package main

import (
	"errors"
	"flag"
	"fmt"
	"go/types"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/packages"
)

type Config struct {
	GoImportPath      string
	OutputRoot        string
	HaxePackagePrefix string
	PackageClassName  string
}

type Emission struct {
	OutputDir string
	Files     []EmittedFile
}

type EmittedFile struct {
	Name     string
	Contents string
}

type declaration struct {
	ClassName       string
	GoTypeName      string
	Interface       bool
	PackageClass    bool
	StaticMethods   []methodDecl
	InstanceMethods []methodDecl
}

type methodDecl struct {
	GoName      string
	HaxeName    string
	Params      []paramDecl
	ReturnType  string
	Static      bool
	TupleReturn bool
}

type paramDecl struct {
	Name string
	Type string
}

type tupleCarrierDecl struct {
	ClassName string
	Fields    []paramDecl
}

type mappingContext struct {
	currentPackagePath string
	exportedTypeNames  map[string]bool
}

var identifierPartPattern = regexp.MustCompile(`[A-Za-z0-9]+`)

var haxeKeywords = map[string]bool{
	"abstract":   true,
	"break":      true,
	"case":       true,
	"cast":       true,
	"catch":      true,
	"class":      true,
	"continue":   true,
	"default":    true,
	"do":         true,
	"dynamic":    true,
	"else":       true,
	"enum":       true,
	"extends":    true,
	"false":      true,
	"final":      true,
	"for":        true,
	"function":   true,
	"if":         true,
	"implements": true,
	"import":     true,
	"in":         true,
	"inline":     true,
	"interface":  true,
	"macro":      true,
	"new":        true,
	"null":       true,
	"operator":   true,
	"overload":   true,
	"override":   true,
	"package":    true,
	"private":    true,
	"public":     true,
	"return":     true,
	"static":     true,
	"super":      true,
	"switch":     true,
	"this":       true,
	"throw":      true,
	"true":       true,
	"try":        true,
	"typedef":    true,
	"untyped":    true,
	"using":      true,
	"var":        true,
	"while":      true,
}

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "goextern: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer, stderr io.Writer) error {
	fs := flag.NewFlagSet("goextern", flag.ContinueOnError)
	fs.SetOutput(stderr)

	goImportPath := fs.String("package", "", "Go import path to introspect (required)")
	outRoot := fs.String("out", "gen/goextern", "Output root directory")
	haxePackagePrefix := fs.String("haxe-package", "goextern", "Root Haxe package prefix")
	packageClassName := fs.String("package-class", "", "Override package static extern class name")
	stdoutOnly := fs.Bool("stdout", false, "Print generated files to stdout instead of writing to disk")

	if err := fs.Parse(args); err != nil {
		return err
	}

	cfg := Config{
		GoImportPath:      strings.TrimSpace(*goImportPath),
		OutputRoot:        strings.TrimSpace(*outRoot),
		HaxePackagePrefix: strings.TrimSpace(*haxePackagePrefix),
		PackageClassName:  strings.TrimSpace(*packageClassName),
	}

	if cfg.GoImportPath == "" {
		return errors.New("missing required -package flag")
	}

	emission, err := BuildEmission(cfg)
	if err != nil {
		return err
	}

	if *stdoutOnly {
		printEmission(stdout, emission)
		return nil
	}

	if err := writeEmission(emission); err != nil {
		return err
	}

	_, _ = fmt.Fprintf(stdout, "generated %d files in %s\n", len(emission.Files), emission.OutputDir)
	return nil
}

func BuildEmission(cfg Config) (*Emission, error) {
	if strings.TrimSpace(cfg.GoImportPath) == "" {
		return nil, errors.New("Go import path is required")
	}
	if strings.TrimSpace(cfg.OutputRoot) == "" {
		cfg.OutputRoot = "gen/goextern"
	}

	pkg, err := loadPackage(cfg.GoImportPath)
	if err != nil {
		return nil, err
	}

	haxePackage, err := deriveHaxePackage(cfg.HaxePackagePrefix, cfg.GoImportPath)
	if err != nil {
		return nil, err
	}

	outputDir, err := deriveOutputDir(cfg.OutputRoot, cfg.GoImportPath)
	if err != nil {
		return nil, err
	}

	decls, carriers, err := collectDeclarations(pkg, cfg.PackageClassName)
	if err != nil {
		return nil, err
	}

	sort.Slice(decls, func(i, j int) bool {
		return decls[i].ClassName < decls[j].ClassName
	})

	files := make([]EmittedFile, 0, len(decls))
	for _, decl := range decls {
		files = append(files, EmittedFile{
			Name:     decl.ClassName + ".hx",
			Contents: renderDeclaration(haxePackage, pkg.Types.Path(), decl),
		})
	}
	for _, carrier := range carriers {
		files = append(files, EmittedFile{
			Name:     carrier.ClassName + ".hx",
			Contents: renderTupleCarrier(haxePackage, carrier),
		})
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	return &Emission{
		OutputDir: outputDir,
		Files:     files,
	}, nil
}

func loadPackage(goImportPath string) (*packages.Package, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedTypes |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedModule,
	}

	pkgs, err := packages.Load(cfg, goImportPath)
	if err != nil {
		return nil, fmt.Errorf("load package %q: %w", goImportPath, err)
	}
	if len(pkgs) == 0 {
		return nil, fmt.Errorf("package %q not found", goImportPath)
	}
	if packages.PrintErrors(pkgs) > 0 {
		return nil, fmt.Errorf("failed loading package %q", goImportPath)
	}

	for _, pkg := range pkgs {
		if pkg != nil && pkg.Types != nil && pkg.Types.Path() == goImportPath {
			return pkg, nil
		}
	}

	for _, pkg := range pkgs {
		if pkg != nil && pkg.Types != nil {
			return pkg, nil
		}
	}

	return nil, fmt.Errorf("loaded package %q without type information", goImportPath)
}

func collectDeclarations(pkg *packages.Package, packageClassOverride string) ([]declaration, []tupleCarrierDecl, error) {
	scope := pkg.Types.Scope()
	if scope == nil {
		return nil, nil, fmt.Errorf("package %q has no scope", pkg.Types.Path())
	}

	names := scope.Names()
	sort.Strings(names)

	exportedTypeNames := make(map[string]bool)
	namedTypes := make([]*types.Named, 0)
	packageFuncs := make([]*types.Func, 0)

	for _, name := range names {
		obj := scope.Lookup(name)
		if obj == nil || !obj.Exported() {
			continue
		}

		switch typed := obj.(type) {
		case *types.TypeName:
			named := asNamedType(typed)
			if named == nil {
				continue
			}
			exportedTypeNames[typed.Name()] = true
			namedTypes = append(namedTypes, named)
		case *types.Func:
			packageFuncs = append(packageFuncs, typed)
		}
	}

	sort.Slice(namedTypes, func(i, j int) bool {
		return namedTypes[i].Obj().Name() < namedTypes[j].Obj().Name()
	})
	sort.Slice(packageFuncs, func(i, j int) bool {
		if packageFuncs[i].Name() == packageFuncs[j].Name() {
			return signatureSortKey(packageFuncs[i]) < signatureSortKey(packageFuncs[j])
		}
		return packageFuncs[i].Name() < packageFuncs[j].Name()
	})

	ctx := mappingContext{
		currentPackagePath: pkg.Types.Path(),
		exportedTypeNames:  exportedTypeNames,
	}
	usedCarrierNames := make(map[string]bool)
	for name := range exportedTypeNames {
		usedCarrierNames[name] = true
	}

	decls := make([]declaration, 0, len(namedTypes)+1)
	carriers := make([]tupleCarrierDecl, 0)
	for _, named := range namedTypes {
		decl, typeCarriers := buildTypeDeclaration(named, ctx, usedCarrierNames)
		decls = append(decls, decl)
		carriers = append(carriers, typeCarriers...)
	}

	if len(packageFuncs) > 0 {
		decl, packageCarriers := buildPackageDeclaration(pkg.Types, packageClassOverride, packageFuncs, exportedTypeNames, ctx, usedCarrierNames)
		decls = append(decls, decl)
		carriers = append(carriers, packageCarriers...)
	}

	sort.Slice(carriers, func(i, j int) bool {
		return carriers[i].ClassName < carriers[j].ClassName
	})

	return decls, carriers, nil
}

func asNamedType(typeName *types.TypeName) *types.Named {
	if typeName == nil || !typeName.Exported() {
		return nil
	}
	unalias := types.Unalias(typeName.Type())
	named, ok := unalias.(*types.Named)
	if !ok || named.Obj() == nil {
		return nil
	}
	if !named.Obj().Exported() {
		return nil
	}
	return named
}

func buildTypeDeclaration(named *types.Named, ctx mappingContext, usedCarrierNames map[string]bool) (declaration, []tupleCarrierDecl) {
	goName := named.Obj().Name()
	_, isInterface := named.Underlying().(*types.Interface)
	methods, carriers := collectMethods(named, false, ctx, goName, usedCarrierNames)

	return declaration{
		ClassName:       goName,
		GoTypeName:      goName,
		Interface:       isInterface,
		PackageClass:    false,
		StaticMethods:   nil,
		InstanceMethods: methods,
	}, carriers
}

func buildPackageDeclaration(pkg *types.Package, classOverride string, funcs []*types.Func, exportedTypes map[string]bool, ctx mappingContext, usedCarrierNames map[string]bool) (declaration, []tupleCarrierDecl) {
	className := ""
	if strings.TrimSpace(classOverride) != "" {
		className = sanitizeClassName(classOverride)
	}
	if className == "" {
		className = toExportedClassName(pkg.Name()) + "Pkg"
	}
	for exportedTypes[className] {
		className += "Api"
	}
	usedCarrierNames[className] = true

	methods := make([]methodDecl, 0, len(funcs))
	carriers := make([]tupleCarrierDecl, 0)
	for _, fn := range funcs {
		if fn == nil || !fn.Exported() {
			continue
		}
		sig, ok := types.Unalias(fn.Type()).(*types.Signature)
		if !ok {
			continue
		}
		method, carrier := signatureToMethod(fn.Name(), sig, true, ctx, "", usedCarrierNames)
		methods = append(methods, method)
		if carrier != nil {
			carriers = append(carriers, *carrier)
		}
	}

	sortMethodDecls(methods)

	return declaration{
		ClassName:       className,
		GoTypeName:      "",
		Interface:       false,
		PackageClass:    true,
		StaticMethods:   methods,
		InstanceMethods: nil,
	}, carriers
}

func collectMethods(named *types.Named, static bool, ctx mappingContext, ownerName string, usedCarrierNames map[string]bool) ([]methodDecl, []tupleCarrierDecl) {
	seen := make(map[string]methodDecl)
	carrierByName := make(map[string]tupleCarrierDecl)

	methodSets := []types.Type{
		named,
		types.NewPointer(named),
	}

	for _, methodTarget := range methodSets {
		set := types.NewMethodSet(methodTarget)
		for i := 0; i < set.Len(); i++ {
			selection := set.At(i)
			fn, ok := selection.Obj().(*types.Func)
			if !ok || !fn.Exported() {
				continue
			}
			sig, ok := types.Unalias(fn.Type()).(*types.Signature)
			if !ok {
				continue
			}

			method, carrier := signatureToMethod(fn.Name(), sig, static, ctx, ownerName, usedCarrierNames)
			key := fn.Name() + "|" + types.TypeString(sig, qualifierByPath)
			if _, exists := seen[key]; !exists {
				seen[key] = method
				if carrier != nil {
					carrierByName[carrier.ClassName] = *carrier
				}
			}
		}
	}

	out := make([]methodDecl, 0, len(seen))
	for _, method := range seen {
		out = append(out, method)
	}

	sortMethodDecls(out)
	carriers := make([]tupleCarrierDecl, 0, len(carrierByName))
	for _, carrier := range carrierByName {
		carriers = append(carriers, carrier)
	}
	sort.Slice(carriers, func(i, j int) bool {
		return carriers[i].ClassName < carriers[j].ClassName
	})
	return out, carriers
}

func signatureToMethod(goName string, sig *types.Signature, static bool, ctx mappingContext, ownerName string, usedCarrierNames map[string]bool) (methodDecl, *tupleCarrierDecl) {
	params := make([]paramDecl, 0, sig.Params().Len())
	usedNames := make(map[string]int)

	for i := 0; i < sig.Params().Len(); i++ {
		param := sig.Params().At(i)
		rawName := param.Name()
		if rawName == "" {
			rawName = "arg" + strconv.Itoa(i+1)
		}

		paramType := mapType(param.Type(), ctx)
		if sig.Variadic() && i == sig.Params().Len()-1 {
			sliceType, ok := types.Unalias(param.Type()).(*types.Slice)
			if ok {
				paramType = "haxe.Rest<" + mapType(sliceType.Elem(), ctx) + ">"
			} else {
				paramType = "haxe.Rest<Dynamic>"
			}
		}

		params = append(params, paramDecl{
			Name: sanitizeParamName(rawName, i, usedNames),
			Type: paramType,
		})
	}

	returnType := "Void"
	var carrier *tupleCarrierDecl
	tupleReturn := false
	switch sig.Results().Len() {
	case 0:
		returnType = "Void"
	case 1:
		returnType = mapType(sig.Results().At(0).Type(), ctx)
	default:
		if generatedCarrier, ok := buildTupleCarrier(goName, sig.Results(), ctx, ownerName, usedCarrierNames); ok {
			returnType = generatedCarrier.ClassName
			carrier = &generatedCarrier
			tupleReturn = true
		} else {
			returnType = "Dynamic"
		}
	}

	haxeName := sanitizeMethodName(lowerCamel(goName))
	if haxeName == "" {
		haxeName = "call"
	}

	return methodDecl{
		GoName:      goName,
		HaxeName:    haxeName,
		Params:      params,
		ReturnType:  returnType,
		Static:      static,
		TupleReturn: tupleReturn,
	}, carrier
}

func buildTupleCarrier(goName string, results *types.Tuple, ctx mappingContext, ownerName string, usedCarrierNames map[string]bool) (tupleCarrierDecl, bool) {
	if results == nil || results.Len() < 2 {
		return tupleCarrierDecl{}, false
	}

	fields := make([]paramDecl, 0, results.Len())
	usedNames := make(map[string]int)
	for i := 0; i < results.Len(); i++ {
		result := results.At(i)
		fieldType, ok := tupleResultType(result.Type())
		if !ok {
			return tupleCarrierDecl{}, false
		}
		rawName := result.Name()
		if rawName == "" {
			rawName = "value" + strconv.Itoa(i+1)
		}
		fields = append(fields, paramDecl{
			Name: sanitizeParamName(rawName, i, usedNames),
			Type: fieldType,
		})
	}

	return tupleCarrierDecl{
		ClassName: uniqueTupleCarrierName(ownerName, goName, usedCarrierNames),
		Fields:    fields,
	}, true
}

func tupleResultType(t types.Type) (string, bool) {
	if isBuiltinErrorType(t) {
		return "Null<go.Error>", true
	}
	basic, ok := types.Unalias(t).(*types.Basic)
	if !ok {
		return "", false
	}
	mapped := mapBasicType(basic)
	if mapped == "Dynamic" {
		return "", false
	}
	return mapped, true
}

func isBuiltinErrorType(t types.Type) bool {
	named, ok := types.Unalias(t).(*types.Named)
	if !ok || named.Obj() == nil {
		return false
	}
	obj := named.Obj()
	return obj.Pkg() == nil && obj.Name() == "error"
}

func uniqueTupleCarrierName(ownerName string, goName string, used map[string]bool) string {
	baseName := goName + "Result"
	if strings.TrimSpace(ownerName) != "" {
		baseName = ownerName + goName + "Result"
	}
	base := sanitizeClassName(baseName)
	if used == nil {
		return base
	}
	if !used[base] {
		used[base] = true
		return base
	}
	for index := 2; ; index++ {
		candidate := base + strconv.Itoa(index)
		if !used[candidate] {
			used[candidate] = true
			return candidate
		}
	}
}

func mapType(t types.Type, ctx mappingContext) string {
	if t == nil {
		return "Dynamic"
	}

	switch tt := types.Unalias(t).(type) {
	case *types.Basic:
		return mapBasicType(tt)
	case *types.Pointer:
		return mapType(tt.Elem(), ctx)
	case *types.Slice:
		return "Array<" + mapType(tt.Elem(), ctx) + ">"
	case *types.Array:
		return "Array<" + mapType(tt.Elem(), ctx) + ">"
	case *types.Map:
		keyType := mapType(tt.Key(), ctx)
		valueType := mapType(tt.Elem(), ctx)
		if keyType == "String" {
			return "haxe.DynamicAccess<" + valueType + ">"
		}
		return "Dynamic"
	case *types.Named:
		obj := tt.Obj()
		if obj == nil {
			return "Dynamic"
		}
		pkg := obj.Pkg()
		if pkg == nil {
			if obj.Name() == "error" {
				return "go.Error"
			}
			return "Dynamic"
		}
		if pkg.Path() == ctx.currentPackagePath {
			if ctx.exportedTypeNames[obj.Name()] {
				return obj.Name()
			}
			return "Dynamic"
		}
		return "Dynamic"
	case *types.Interface:
		if tt.NumMethods() == 0 {
			return "Dynamic"
		}
		return "Dynamic"
	case *types.Signature:
		return "Dynamic"
	case *types.Struct:
		return "Dynamic"
	case *types.Tuple:
		return "Dynamic"
	case *types.Chan:
		return "Dynamic"
	case *types.TypeParam:
		return "Dynamic"
	case *types.Union:
		return "Dynamic"
	default:
		return "Dynamic"
	}
}

func mapBasicType(basic *types.Basic) string {
	if basic == nil {
		return "Dynamic"
	}

	if basic.Kind() == types.String {
		return "String"
	}
	if basic.Info()&types.IsBoolean != 0 {
		return "Bool"
	}
	if basic.Info()&types.IsInteger != 0 {
		return "Int"
	}
	if basic.Info()&types.IsFloat != 0 {
		return "Float"
	}
	if basic.Kind() == types.UnsafePointer {
		return "Dynamic"
	}
	return "Dynamic"
}

func renderDeclaration(haxePackage string, goImportPath string, decl declaration) string {
	var b strings.Builder
	b.WriteString("// Code generated by tools/goextern; DO NOT EDIT.\n")
	b.WriteString("// Source package: ")
	b.WriteString(goImportPath)
	b.WriteString("\n")

	if haxePackage == "" {
		b.WriteString("package;\n\n")
	} else {
		b.WriteString("package ")
		b.WriteString(haxePackage)
		b.WriteString(";\n\n")
	}

	b.WriteString("@:go.import(\"")
	b.WriteString(goImportPath)
	b.WriteString("\")\n")

	if !decl.PackageClass {
		b.WriteString("@:go.name(\"")
		b.WriteString(decl.GoTypeName)
		b.WriteString("\")\n")
	}

	allMethods := make([]methodDecl, 0, len(decl.StaticMethods)+len(decl.InstanceMethods))
	allMethods = append(allMethods, decl.StaticMethods...)
	allMethods = append(allMethods, decl.InstanceMethods...)
	sortMethodDecls(allMethods)

	if decl.Interface {
		b.WriteString("extern interface ")
		b.WriteString(decl.ClassName)
		if len(allMethods) == 0 {
			b.WriteString(" {}\n")
			return b.String()
		}
		b.WriteString(" {\n")
	} else {
		b.WriteString("extern class ")
		b.WriteString(decl.ClassName)
		if len(allMethods) == 0 {
			b.WriteString(" {}\n")
			return b.String()
		}
		b.WriteString(" {\n")
	}

	for idx, method := range allMethods {
		if idx > 0 {
			b.WriteString("\n")
		}
		if method.TupleReturn {
			b.WriteString("\t@:go.tupleReturn\n")
		}
		b.WriteString("\t@:go.name(\"")
		b.WriteString(method.GoName)
		b.WriteString("\")\n")
		b.WriteString("\tpublic ")
		if method.Static {
			b.WriteString("static ")
		}
		b.WriteString("function ")
		b.WriteString(method.HaxeName)
		b.WriteString("(")
		for i, param := range method.Params {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(param.Name)
			b.WriteString(":")
			b.WriteString(param.Type)
		}
		b.WriteString("):")
		b.WriteString(method.ReturnType)
		b.WriteString(";\n")
	}

	b.WriteString("}\n")
	return b.String()
}

func renderTupleCarrier(haxePackage string, carrier tupleCarrierDecl) string {
	var b strings.Builder
	b.WriteString("// Code generated by tools/goextern; DO NOT EDIT.\n")
	b.WriteString("// Tuple carrier for Go multi-return extern calls.\n")

	if haxePackage == "" {
		b.WriteString("package;\n\n")
	} else {
		b.WriteString("package ")
		b.WriteString(haxePackage)
		b.WriteString(";\n\n")
	}

	b.WriteString("/**\n")
	b.WriteString("\tWhat: typed Haxe carrier for a Go function that returns more than one value.\n")
	b.WriteString("\tWhy: Haxe functions return one value, while Go can return multiple values directly.\n")
	b.WriteString("\tHow: haxe.go lowers the matching `@:go.tupleReturn` extern call into this carrier.\n")
	b.WriteString("**/\n")
	b.WriteString("class ")
	b.WriteString(carrier.ClassName)
	b.WriteString(" {\n")
	for _, field := range carrier.Fields {
		b.WriteString("\tpublic var ")
		b.WriteString(field.Name)
		b.WriteString("(default, null):")
		b.WriteString(field.Type)
		b.WriteString(";\n")
	}
	if len(carrier.Fields) > 0 {
		b.WriteString("\n")
	}
	b.WriteString("\tpublic function new(")
	for i, field := range carrier.Fields {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(field.Name)
		b.WriteString(":")
		b.WriteString(field.Type)
	}
	b.WriteString(") {\n")
	for _, field := range carrier.Fields {
		b.WriteString("\t\tthis.")
		b.WriteString(field.Name)
		b.WriteString(" = ")
		b.WriteString(field.Name)
		b.WriteString(";\n")
	}
	b.WriteString("\t}\n")
	b.WriteString("}\n")
	return b.String()
}

func writeEmission(emission *Emission) error {
	if emission == nil {
		return errors.New("nil emission")
	}
	if emission.OutputDir == "" {
		return errors.New("empty output directory")
	}

	if err := os.MkdirAll(emission.OutputDir, 0o755); err != nil {
		return fmt.Errorf("create output directory %q: %w", emission.OutputDir, err)
	}

	expected := make(map[string]bool, len(emission.Files))
	for _, file := range emission.Files {
		target := filepath.Join(emission.OutputDir, file.Name)
		expected[file.Name] = true
		if err := os.WriteFile(target, []byte(file.Contents), 0o644); err != nil {
			return fmt.Errorf("write %q: %w", target, err)
		}
	}

	entries, err := os.ReadDir(emission.OutputDir)
	if err != nil {
		return fmt.Errorf("read output directory %q: %w", emission.OutputDir, err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) != ".hx" {
			continue
		}
		if expected[entry.Name()] {
			continue
		}
		target := filepath.Join(emission.OutputDir, entry.Name())
		if err := os.Remove(target); err != nil {
			return fmt.Errorf("remove stale file %q: %w", target, err)
		}
	}

	return nil
}

func printEmission(stdout io.Writer, emission *Emission) {
	if emission == nil {
		return
	}
	for idx, file := range emission.Files {
		if idx > 0 {
			_, _ = io.WriteString(stdout, "\n")
		}
		_, _ = fmt.Fprintf(stdout, "=== %s ===\n", file.Name)
		_, _ = io.WriteString(stdout, file.Contents)
	}
}

func deriveHaxePackage(prefix string, goImportPath string) (string, error) {
	goSegments, err := sanitizePathSegments(goImportPath)
	if err != nil {
		return "", err
	}

	parts := make([]string, 0, 8)
	if strings.TrimSpace(prefix) != "" {
		for _, segment := range strings.Split(prefix, ".") {
			segment = strings.TrimSpace(segment)
			if segment == "" {
				continue
			}
			parts = append(parts, sanitizePackageSegment(segment))
		}
	}
	for _, segment := range goSegments {
		parts = append(parts, sanitizePackageSegment(segment))
	}

	return strings.Join(parts, "."), nil
}

func deriveOutputDir(outRoot string, goImportPath string) (string, error) {
	if strings.TrimSpace(outRoot) == "" {
		return "", errors.New("output root directory must not be empty")
	}
	segments, err := sanitizePathSegments(goImportPath)
	if err != nil {
		return "", err
	}
	all := append([]string{outRoot}, segments...)
	return filepath.Join(all...), nil
}

func sanitizePathSegments(importPath string) ([]string, error) {
	trimmed := strings.TrimSpace(importPath)
	if trimmed == "" {
		return nil, errors.New("Go import path must not be empty")
	}

	rawSegments := strings.Split(trimmed, "/")
	out := make([]string, 0, len(rawSegments))
	for _, segment := range rawSegments {
		segment = strings.TrimSpace(segment)
		if segment == "" || segment == "." || segment == ".." {
			return nil, fmt.Errorf("invalid import path segment %q", segment)
		}
		var b strings.Builder
		for _, r := range segment {
			switch {
			case unicode.IsLetter(r), unicode.IsDigit(r), r == '.', r == '-', r == '_':
				b.WriteRune(r)
			default:
				b.WriteRune('_')
			}
		}
		sanitized := b.String()
		if sanitized == "" || sanitized == "." || sanitized == ".." {
			return nil, fmt.Errorf("invalid sanitized import path segment %q", segment)
		}
		out = append(out, sanitized)
	}
	return out, nil
}

func sanitizePackageSegment(segment string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(strings.TrimSpace(segment)) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			b.WriteRune(r)
		} else {
			b.WriteRune('_')
		}
	}
	out := b.String()
	if out == "" {
		out = "_"
	}
	first, _ := utf8.DecodeRuneInString(out)
	if unicode.IsDigit(first) {
		out = "_" + out
	}
	if haxeKeywords[out] {
		out += "_pkg"
	}
	return out
}

func toExportedClassName(raw string) string {
	parts := identifierPartPattern.FindAllString(raw, -1)
	if len(parts) == 0 {
		return "Package"
	}
	var b strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		lower := strings.ToLower(part)
		r, size := utf8.DecodeRuneInString(lower)
		if r == utf8.RuneError && size == 0 {
			continue
		}
		b.WriteRune(unicode.ToUpper(r))
		b.WriteString(lower[size:])
	}
	name := b.String()
	if name == "" {
		name = "Package"
	}
	first, _ := utf8.DecodeRuneInString(name)
	if unicode.IsDigit(first) {
		name = "_" + name
	}
	return sanitizeClassName(name)
}

func lowerCamel(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return ""
	}
	if name == strings.ToUpper(name) {
		return strings.ToLower(name)
	}
	r, size := utf8.DecodeRuneInString(name)
	if r == utf8.RuneError && size == 0 {
		return ""
	}
	return string(unicode.ToLower(r)) + name[size:]
}

func sanitizeClassName(name string) string {
	clean := sanitizeIdentifier(name, "Type")
	r, size := utf8.DecodeRuneInString(clean)
	if r == utf8.RuneError && size == 0 {
		return "Type"
	}
	return string(unicode.ToUpper(r)) + clean[size:]
}

func sanitizeMethodName(name string) string {
	return sanitizeIdentifier(name, "call")
}

func sanitizeParamName(name string, index int, used map[string]int) string {
	base := sanitizeIdentifier(name, "arg"+strconv.Itoa(index+1))
	if used == nil {
		return base
	}
	if _, exists := used[base]; !exists {
		used[base] = 1
		return base
	}
	next := used[base]
	used[base] = next + 1
	return base + strconv.Itoa(next)
}

func sanitizeIdentifier(name string, fallback string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = fallback
	}
	var b strings.Builder
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			b.WriteRune(r)
		} else {
			b.WriteRune('_')
		}
	}
	out := b.String()
	if out == "" {
		out = fallback
	}
	first, _ := utf8.DecodeRuneInString(out)
	if unicode.IsDigit(first) {
		out = "_" + out
	}
	if haxeKeywords[out] {
		out += "_"
	}
	return out
}

func sortMethodDecls(methods []methodDecl) {
	sort.Slice(methods, func(i, j int) bool {
		if methods[i].GoName == methods[j].GoName {
			if methods[i].HaxeName == methods[j].HaxeName {
				return methodSignatureKey(methods[i]) < methodSignatureKey(methods[j])
			}
			return methods[i].HaxeName < methods[j].HaxeName
		}
		return methods[i].GoName < methods[j].GoName
	})
}

func methodSignatureKey(method methodDecl) string {
	var b strings.Builder
	b.WriteString(method.HaxeName)
	b.WriteString("|")
	for _, param := range method.Params {
		b.WriteString(param.Name)
		b.WriteString(":")
		b.WriteString(param.Type)
		b.WriteString(",")
	}
	b.WriteString("->")
	b.WriteString(method.ReturnType)
	return b.String()
}

func signatureSortKey(fn *types.Func) string {
	if fn == nil {
		return ""
	}
	sig, ok := types.Unalias(fn.Type()).(*types.Signature)
	if !ok {
		return fn.Name()
	}
	return fn.Name() + "|" + types.TypeString(sig, qualifierByPath)
}

func qualifierByPath(pkg *types.Package) string {
	if pkg == nil {
		return ""
	}
	return pkg.Path()
}
