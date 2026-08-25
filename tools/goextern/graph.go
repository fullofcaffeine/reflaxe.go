package main

import (
	"crypto/sha256"
	"fmt"
	"go/types"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

type declarationID struct {
	PackagePath string
	GoName      string
}

type packagePlan struct {
	GoPackageName  string
	HaxePackage    string
	RelativeDir    string
	Declarations   []declaration
	Carriers       []tupleCarrierDecl
	UsedClassNames map[string]bool
}

type declarationGraph struct {
	config       Config
	pending      map[declarationID]*types.TypeName
	done         map[declarationID]bool
	packages     map[string]*packagePlan
	fallbacks    []DynamicFallback
	rootTypeIDs  []declarationID
	packageFuncs []*types.Func
}

type generatorError struct {
	Code    string
	Message string
}

func (failure *generatorError) Error() string {
	return failure.Code + ": " + failure.Message
}

func graphError(code string, format string, args ...any) error {
	return &generatorError{Code: code, Message: fmt.Sprintf(format, args...)}
}

func buildGraphEmission(cfg Config, root *packages.Package) (*Emission, error) {
	if root == nil || root.Types == nil || root.Types.Scope() == nil {
		return nil, fmt.Errorf("package %q has no type scope", cfg.GoImportPath)
	}

	graph := &declarationGraph{
		config:   cfg,
		pending:  make(map[declarationID]*types.TypeName),
		done:     make(map[declarationID]bool),
		packages: make(map[string]*packagePlan),
	}
	if err := graph.seedRoot(root.Types); err != nil {
		return nil, err
	}
	if err := graph.buildRootTypes(); err != nil {
		return nil, err
	}
	if err := graph.buildRootPackageClass(root.Types); err != nil {
		return nil, err
	}
	if err := graph.buildPending(); err != nil {
		return nil, err
	}

	files := make([]EmittedFile, 0)
	packagePaths := make([]string, 0, len(graph.packages))
	for packagePath := range graph.packages {
		packagePaths = append(packagePaths, packagePath)
	}
	sort.Strings(packagePaths)
	for _, packagePath := range packagePaths {
		plan := graph.packages[packagePath]
		sort.Slice(plan.Declarations, func(i, j int) bool {
			return plan.Declarations[i].ClassName < plan.Declarations[j].ClassName
		})
		sort.Slice(plan.Carriers, func(i, j int) bool {
			return plan.Carriers[i].ClassName < plan.Carriers[j].ClassName
		})
		for _, decl := range plan.Declarations {
			files = append(files, EmittedFile{
				Name: filepath.Join(plan.RelativeDir, decl.ClassName+".hx"),
				Contents: renderDeclaration(
					plan.HaxePackage,
					packagePath,
					plan.GoPackageName,
					decl,
				),
			})
		}
		for _, carrier := range plan.Carriers {
			files = append(files, EmittedFile{
				Name:     filepath.Join(plan.RelativeDir, carrier.ClassName+".hx"),
				Contents: renderTupleCarrier(plan.HaxePackage, carrier),
			})
		}
	}
	sort.Slice(files, func(i, j int) bool { return files[i].Name < files[j].Name })
	if _, _, err := plannedOutputs(files); err != nil {
		return nil, err
	}

	return &Emission{
		OutputDir: cfg.OutputRoot,
		RootKey:   rootKey(cfg),
		Root: ownershipRoot{
			GoImportPath:      cfg.GoImportPath,
			HaxePackagePrefix: cfg.HaxePackagePrefix,
			PackageClassName:  cfg.PackageClassName,
		},
		PackageCount:     len(graph.packages),
		Files:            files,
		DynamicFallbacks: sortedDynamicFallbacks(graph.fallbacks),
	}, nil
}

func (graph *declarationGraph) seedRoot(pkg *types.Package) error {
	if _, err := graph.packagePlan(pkg); err != nil {
		return err
	}
	scope := pkg.Scope()
	names := scope.Names()
	sort.Strings(names)
	for _, name := range names {
		obj := scope.Lookup(name)
		if obj == nil || !obj.Exported() {
			continue
		}
		switch typed := obj.(type) {
		case *types.TypeName:
			id := typeNameID(typed)
			graph.pending[id] = typed
			graph.rootTypeIDs = append(graph.rootTypeIDs, id)
		case *types.Func:
			graph.packageFuncs = append(graph.packageFuncs, typed)
		}
	}
	sort.Slice(graph.rootTypeIDs, func(i, j int) bool {
		return declarationIDLess(graph.rootTypeIDs[i], graph.rootTypeIDs[j])
	})
	sort.Slice(graph.packageFuncs, func(i, j int) bool {
		if graph.packageFuncs[i].Name() == graph.packageFuncs[j].Name() {
			return signatureSortKey(graph.packageFuncs[i]) < signatureSortKey(graph.packageFuncs[j])
		}
		return graph.packageFuncs[i].Name() < graph.packageFuncs[j].Name()
	})
	return nil
}

func (graph *declarationGraph) buildRootTypes() error {
	for _, id := range graph.rootTypeIDs {
		if err := graph.buildType(id, graph.pending[id]); err != nil {
			return err
		}
	}
	return nil
}

func (graph *declarationGraph) buildRootPackageClass(pkg *types.Package) error {
	if len(graph.packageFuncs) == 0 {
		return nil
	}
	plan, err := graph.packagePlan(pkg)
	if err != nil {
		return err
	}
	ctx := graph.mappingContext(pkg)
	exportedTypes := exportedTypeNames(pkg)
	decl, carriers, fallbacks := buildPackageDeclaration(
		pkg,
		graph.config.PackageClassName,
		graph.packageFuncs,
		exportedTypes,
		ctx,
		plan.UsedClassNames,
	)
	plan.Declarations = append(plan.Declarations, decl)
	plan.Carriers = append(plan.Carriers, carriers...)
	graph.fallbacks = append(graph.fallbacks, fallbacks...)
	return nil
}

func (graph *declarationGraph) buildPending() error {
	for {
		ids := make([]declarationID, 0)
		for id := range graph.pending {
			if !graph.done[id] {
				ids = append(ids, id)
			}
		}
		if len(ids) == 0 {
			return nil
		}
		sort.Slice(ids, func(i, j int) bool { return declarationIDLess(ids[i], ids[j]) })
		for _, id := range ids {
			if err := graph.buildType(id, graph.pending[id]); err != nil {
				return err
			}
		}
	}
}

func (graph *declarationGraph) buildType(id declarationID, typeName *types.TypeName) error {
	if graph.done[id] {
		return nil
	}
	graph.done[id] = true
	if typeName == nil || typeName.Pkg() == nil || !typeName.Exported() {
		return nil
	}
	plan, err := graph.packagePlan(typeName.Pkg())
	if err != nil {
		return err
	}
	ctx := graph.mappingContext(typeName.Pkg())
	if typeName.IsAlias() {
		decl, fallbacks := buildAliasDeclaration(typeName, ctx)
		plan.Declarations = append(plan.Declarations, decl)
		graph.fallbacks = append(graph.fallbacks, fallbacks...)
		return nil
	}
	named := asNamedType(typeName)
	if named == nil {
		return nil
	}
	if named.TypeParams() != nil && named.TypeParams().Len() > 0 {
		graph.fallbacks = append(graph.fallbacks, newDynamicFallback(ctx, typeName.Name(), "type", typeName.Type(), "generic_named_type"))
		return nil
	}
	decl, carriers, fallbacks := buildTypeDeclaration(named, ctx, plan.UsedClassNames)
	plan.Declarations = append(plan.Declarations, decl)
	plan.Carriers = append(plan.Carriers, carriers...)
	graph.fallbacks = append(graph.fallbacks, fallbacks...)
	return nil
}

func (graph *declarationGraph) mappingContext(pkg *types.Package) mappingContext {
	return mappingContext{
		currentPackagePath: pkg.Path(),
		haxePackagePrefix:  graph.config.HaxePackagePrefix,
		schedule: func(typeName *types.TypeName) {
			if typeName == nil || typeName.Pkg() == nil || !typeName.Exported() {
				return
			}
			id := typeNameID(typeName)
			if !graph.done[id] {
				graph.pending[id] = typeName
			}
		},
	}
}

func (graph *declarationGraph) packagePlan(pkg *types.Package) (*packagePlan, error) {
	if pkg == nil || strings.TrimSpace(pkg.Path()) == "" {
		return nil, fmt.Errorf("package has no import path")
	}
	if existing := graph.packages[pkg.Path()]; existing != nil {
		return existing, nil
	}
	haxePackage, err := deriveHaxePackage(graph.config.HaxePackagePrefix, pkg.Path())
	if err != nil {
		return nil, err
	}
	segments, err := sanitizePathSegments(pkg.Path())
	if err != nil {
		return nil, err
	}
	haxeSegments := make([]string, 0, len(segments))
	for _, segment := range segments {
		haxeSegments = append(haxeSegments, sanitizePackageSegment(segment))
	}
	relativeDir := filepath.Join(haxeSegments...)
	for otherPath, other := range graph.packages {
		if strings.EqualFold(other.HaxePackage, haxePackage) {
			return nil, graphError("haxe_package_collision", "%q and %q project to %q", otherPath, pkg.Path(), haxePackage)
		}
		if strings.EqualFold(other.RelativeDir, relativeDir) {
			return nil, graphError("output_path_collision", "%q and %q project to %q", otherPath, pkg.Path(), relativeDir)
		}
		if other.GoPackageName == pkg.Name() {
			return nil, graphError("go_import_alias_required", "%q and %q both use Go package name %q", otherPath, pkg.Path(), pkg.Name())
		}
	}
	used := exportedTypeNames(pkg)
	plan := &packagePlan{
		GoPackageName:  pkg.Name(),
		HaxePackage:    haxePackage,
		RelativeDir:    relativeDir,
		UsedClassNames: used,
	}
	graph.packages[pkg.Path()] = plan
	return plan, nil
}

func buildAliasDeclaration(typeName *types.TypeName, ctx mappingContext) (declaration, []DynamicFallback) {
	target := types.Unalias(typeName.Type())
	mapped := mapTypeResult(target, ctx)
	fallbacks := make([]DynamicFallback, 0, 1)
	if mapped.Reason != "" {
		fallbacks = append(fallbacks, newDynamicFallback(ctx, typeName.Name(), "alias", typeName.Type(), "alias_target_unsupported"))
		mapped = mappedType{Haxe: "Dynamic", Reason: "alias_target_unsupported"}
	} else {
		commitMappedType(ctx, mapped)
	}
	return declaration{
		ClassName:   typeName.Name(),
		GoTypeName:  typeName.Name(),
		Alias:       true,
		AliasTarget: mapped.Haxe,
	}, fallbacks
}

func exportedTypeNames(pkg *types.Package) map[string]bool {
	out := make(map[string]bool)
	if pkg == nil || pkg.Scope() == nil {
		return out
	}
	for _, name := range pkg.Scope().Names() {
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok && obj.Exported() {
			out[name] = true
		}
	}
	return out
}

func typeNameID(typeName *types.TypeName) declarationID {
	id := declarationID{GoName: typeName.Name()}
	if typeName.Pkg() != nil {
		id.PackagePath = typeName.Pkg().Path()
	}
	return id
}

func declarationIDLess(left declarationID, right declarationID) bool {
	if left.PackagePath != right.PackagePath {
		return left.PackagePath < right.PackagePath
	}
	return left.GoName < right.GoName
}

func rootKey(cfg Config) string {
	payload := cfg.GoImportPath + "\x00" + cfg.HaxePackagePrefix + "\x00" + cfg.PackageClassName
	return fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
}
