package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"go/token"
	"go/types"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDeriveHaxePackageSanitizesSegments(t *testing.T) {
	t.Parallel()

	got, err := deriveHaxePackage("goextern", "github.com/go-sql-driver/mysql")
	if err != nil {
		t.Fatalf("deriveHaxePackage returned error: %v", err)
	}

	want := "goextern.github_com.go_sql_driver.mysql"
	if got != want {
		t.Fatalf("deriveHaxePackage mismatch: got %q, want %q", got, want)
	}
}

func TestBuildEmissionFmtIsDeterministic(t *testing.T) {
	cfg := Config{
		GoImportPath:      "fmt",
		OutputRoot:        t.TempDir(),
		HaxePackagePrefix: "goextern",
	}

	first, err := BuildEmission(cfg)
	if err != nil {
		t.Fatalf("BuildEmission first run failed: %v", err)
	}

	second, err := BuildEmission(cfg)
	if err != nil {
		t.Fatalf("BuildEmission second run failed: %v", err)
	}

	if len(first.Files) != len(second.Files) {
		t.Fatalf("file count mismatch: first=%d second=%d", len(first.Files), len(second.Files))
	}
	for i := range first.Files {
		if first.Files[i].Name != second.Files[i].Name {
			t.Fatalf("file[%d] name mismatch: %q != %q", i, first.Files[i].Name, second.Files[i].Name)
		}
		if first.Files[i].Contents != second.Files[i].Contents {
			t.Fatalf("file[%d] contents mismatch for %s", i, first.Files[i].Name)
		}
	}

	pkgFile := fileContentsByName(t, first, "FmtPkg.hx")
	if !strings.Contains(pkgFile, "// Source package: fmt\npackage goextern.fmt;") {
		t.Fatalf("FmtPkg.hx header/package spacing drifted")
	}
	if !strings.Contains(pkgFile, "@:go.import(\"fmt\")") {
		t.Fatalf("FmtPkg.hx missing @:go.import metadata")
	}
	if !strings.Contains(pkgFile, "extern class FmtPkg") {
		t.Fatalf("FmtPkg.hx missing class declaration")
	}
	if !strings.Contains(pkgFile, "@:go.name(\"Println\")") {
		t.Fatalf("FmtPkg.hx missing Println symbol mapping")
	}
}

func TestBuildEmissionLoadsPackageFromCallerModuleWithoutMutation(t *testing.T) {
	moduleDir := t.TempDir()
	writeTestFile(t, filepath.Join(moduleDir, "go.mod"), "module example.com/goexternfixture\n\ngo 1.22\n")
	writeTestFile(t, filepath.Join(moduleDir, "model", "record.go"), `package model

type Record struct {
	ID string
}
`)
	writeTestFile(t, filepath.Join(moduleDir, "api", "api.go"), `package api

import "example.com/goexternfixture/model"

func Find(id string) *model.Record {
	return &model.Record{ID: id}
}
`)

	goModBefore, err := os.ReadFile(filepath.Join(moduleDir, "go.mod"))
	if err != nil {
		t.Fatalf("read fixture go.mod: %v", err)
	}

	emission, err := BuildEmission(Config{
		GoImportPath:      "example.com/goexternfixture/api",
		WorkingDirectory:  moduleDir,
		OutputRoot:        t.TempDir(),
		HaxePackagePrefix: "goextern",
	})
	if err != nil {
		t.Fatalf("BuildEmission failed for caller-owned module package: %v", err)
	}

	apiFile := fileContentsByName(t, emission, "ApiPkg.hx")
	if !strings.Contains(apiFile, "@:go.name(\"Find\")") {
		t.Fatalf("ApiPkg.hx missing Find symbol mapping:\n%s", apiFile)
	}
	if !strings.Contains(apiFile, "public static function find(id:String):goextern.example_com.goexternfixture.model.Record;") {
		t.Fatalf("ApiPkg.hx must use the precise external named type:\n%s", apiFile)
	}
	recordFile := fileContentsByName(t, emission, "Record.hx")
	for _, snippet := range []string{
		"package goextern.example_com.goexternfixture.model;",
		"@:go.import(\"example.com/goexternfixture/model\")",
		"@:go.package(\"model\")",
		"@:go.name(\"Record\")",
		"@:go.struct",
		"public var id:String;",
	} {
		if !strings.Contains(recordFile, snippet) {
			t.Fatalf("Record.hx missing cross-package contract %q:\n%s", snippet, recordFile)
		}
	}
	for _, fallback := range emission.DynamicFallbacks {
		if fallback.Symbol == "Find" && fallback.Reason == "external_named_type" {
			t.Fatalf("Find must not retain an external_named_type fallback: %+v", fallback)
		}
	}

	goModAfter, err := os.ReadFile(filepath.Join(moduleDir, "go.mod"))
	if err != nil {
		t.Fatalf("read fixture go.mod after inspection: %v", err)
	}
	if string(goModAfter) != string(goModBefore) {
		t.Fatalf("goextern changed the caller's go.mod")
	}
	if exists(filepath.Join(moduleDir, "go.sum")) {
		t.Fatalf("goextern created go.sum while inspecting a local-only module")
	}
}

func TestBuildEmissionRejectsPackagePatternIdentity(t *testing.T) {
	moduleDir := t.TempDir()
	outputDir := t.TempDir()
	writeTestFile(t, filepath.Join(moduleDir, "go.mod"), "module example.com/patternfixture\n\ngo 1.22\n")
	writeTestFile(t, filepath.Join(moduleDir, "api", "api.go"), "package api\n\ntype Record struct{}\n")

	_, err := BuildEmission(Config{
		GoImportPath:      "example.com/patternfixture/...",
		WorkingDirectory:  moduleDir,
		OutputRoot:        outputDir,
		HaxePackagePrefix: "goextern",
	})
	expectGeneratorErrorCode(t, err, "package_load_failed")
	if !strings.Contains(err.Error(), "did not match exact import path") {
		t.Fatalf("package pattern diagnostic must explain exact identity: %v", err)
	}
}

func TestBuildEmissionClosesSupportedCrossPackageDeclarationGraph(t *testing.T) {
	moduleDir := t.TempDir()
	writeTestFile(t, filepath.Join(moduleDir, "go.mod"), "module example.com/graphfixture\n\ngo 1.22\n")
	writeTestFile(t, filepath.Join(moduleDir, "contract", "context.go"), `package contract

type Context interface {
	Err() error
}
`)
	writeTestFile(t, filepath.Join(moduleDir, "detail", "detail.go"), `package detail

type Detail struct {
	Bonus int
}
`)
	writeTestFile(t, filepath.Join(moduleDir, "model", "item.go"), `package model

import "example.com/graphfixture/detail"

type Item struct {
	Value  int
	Detail *detail.Detail
	Next   *Item
	Pair   *Pair
}

type Pair struct {
	Item *Item
}
`)
	writeTestFile(t, filepath.Join(moduleDir, "api", "api.go"), `package api

import (
	"example.com/graphfixture/contract"
	"example.com/graphfixture/model"
)

type ItemAlias = model.Item
type SecondAlias = ItemAlias

type Page struct {
	Items []*model.Item
}

func Background() contract.Context { return nil }
func Alias() *ItemAlias { return nil }
func Second() *SecondAlias { return nil }
func Lookup(ctx contract.Context, seed int) (*model.Item, error) { return nil, nil }
func List(ctx contract.Context, seed int) (*Page, error) { return nil, nil }
`)

	cfg := Config{
		GoImportPath:      "example.com/graphfixture/api",
		WorkingDirectory:  moduleDir,
		OutputRoot:        t.TempDir(),
		HaxePackagePrefix: "goextern",
	}
	first, err := BuildEmission(cfg)
	if err != nil {
		t.Fatalf("BuildEmission failed: %v", err)
	}
	second, err := BuildEmission(cfg)
	if err != nil {
		t.Fatalf("second BuildEmission failed: %v", err)
	}
	if len(first.DynamicFallbacks) != 0 {
		t.Fatalf("supported graph must be exact, got fallbacks: %+v", first.DynamicFallbacks)
	}
	if len(first.Files) != len(second.Files) {
		t.Fatalf("graph file count changed: %d != %d", len(first.Files), len(second.Files))
	}
	for index := range first.Files {
		if first.Files[index] != second.Files[index] {
			t.Fatalf("graph output changed at index %d: %+v != %+v", index, first.Files[index], second.Files[index])
		}
	}

	api := fileContentsByPath(t, first, "example_com/graphfixture/api/ApiPkg.hx")
	for _, snippet := range []string{
		"background():goextern.example_com.graphfixture.contract.Context;",
		"lookup(ctx:goextern.example_com.graphfixture.contract.Context, seed:Int):LookupResult;",
		"list(ctx:goextern.example_com.graphfixture.contract.Context, seed:Int):ListResult;",
	} {
		if !strings.Contains(api, snippet) {
			t.Fatalf("ApiPkg.hx missing graph signature %q:\n%s", snippet, api)
		}
	}
	page := fileContentsByPath(t, first, "example_com/graphfixture/api/Page.hx")
	if !strings.Contains(page, "public var items:go.NativeSlice<goextern.example_com.graphfixture.model.Item>;") {
		t.Fatalf("Page.hx missing precise external slice field:\n%s", page)
	}
	alias := fileContentsByPath(t, first, "example_com/graphfixture/api/ItemAlias.hx")
	if !strings.Contains(alias, "typedef ItemAlias = goextern.example_com.graphfixture.model.Item;") {
		t.Fatalf("ItemAlias.hx must preserve the public Go alias:\n%s", alias)
	}
	secondAlias := fileContentsByPath(t, first, "example_com/graphfixture/api/SecondAlias.hx")
	if !strings.Contains(secondAlias, "typedef SecondAlias = goextern.example_com.graphfixture.model.Item;") {
		t.Fatalf("SecondAlias.hx must retain an exact canonical alias target:\n%s", secondAlias)
	}
	item := fileContentsByPath(t, first, "example_com/graphfixture/model/Item.hx")
	for _, snippet := range []string{
		"public var detail:goextern.example_com.graphfixture.detail.Detail;",
		"public var next:Item;",
		"public var pair:Pair;",
	} {
		if !strings.Contains(item, snippet) {
			t.Fatalf("Item.hx missing recursive graph field %q:\n%s", snippet, item)
		}
	}
	_ = fileContentsByPath(t, first, "example_com/graphfixture/model/Pair.hx")
	_ = fileContentsByPath(t, first, "example_com/graphfixture/detail/Detail.hx")
	contextType := fileContentsByPath(t, first, "example_com/graphfixture/contract/Context.hx")
	if !strings.Contains(contextType, "@:go.package(\"contract\")") || !strings.Contains(contextType, "extern interface Context") {
		t.Fatalf("Context.hx missing exact package metadata or interface declaration:\n%s", contextType)
	}
}

func TestDeclarationGraphRejectsPackageProjectionCollisions(t *testing.T) {
	graph := &declarationGraph{
		config:   Config{HaxePackagePrefix: "goextern"},
		packages: make(map[string]*packagePlan),
	}
	if _, err := graph.packagePlan(types.NewPackage("example.com/dep-one", "depone")); err != nil {
		t.Fatalf("first package plan failed: %v", err)
	}
	_, err := graph.packagePlan(types.NewPackage("example.com/dep_one", "depunder"))
	expectGeneratorErrorCode(t, err, "haxe_package_collision")
}

func TestDeclarationGraphRejectsDuplicateGoPackageQualifiers(t *testing.T) {
	graph := &declarationGraph{
		config:   Config{HaxePackagePrefix: "goextern"},
		packages: make(map[string]*packagePlan),
	}
	if _, err := graph.packagePlan(types.NewPackage("example.com/first", "shared")); err != nil {
		t.Fatalf("first package plan failed: %v", err)
	}
	_, err := graph.packagePlan(types.NewPackage("example.net/second", "shared"))
	expectGeneratorErrorCode(t, err, "go_import_alias_required")
}

func TestBuildEmissionRejectsCaseFoldedTypePathCollision(t *testing.T) {
	moduleDir := t.TempDir()
	writeTestFile(t, filepath.Join(moduleDir, "go.mod"), "module example.com/collisionfixture\n\ngo 1.22\n")
	writeTestFile(t, filepath.Join(moduleDir, "model", "model.go"), `package model

type URL struct{}
type Url struct{}
`)
	_, err := BuildEmission(Config{
		GoImportPath:      "example.com/collisionfixture/model",
		WorkingDirectory:  moduleDir,
		OutputRoot:        t.TempDir(),
		HaxePackagePrefix: "goextern",
	})
	expectGeneratorErrorCode(t, err, "output_path_collision")
}

func TestBuildEmissionReportsStablePackageLoadErrorBeforeWriting(t *testing.T) {
	moduleDir := t.TempDir()
	outputDir := t.TempDir()
	writeTestFile(t, filepath.Join(moduleDir, "go.mod"), "module example.com/loadfixture\n\ngo 1.22\n")
	_, err := BuildEmission(Config{
		GoImportPath:      "example.com/loadfixture/missing",
		WorkingDirectory:  moduleDir,
		OutputRoot:        outputDir,
		HaxePackagePrefix: "goextern",
	})
	expectGeneratorErrorCode(t, err, "package_load_failed")
	entries, readErr := os.ReadDir(outputDir)
	if readErr != nil {
		t.Fatalf("read output directory: %v", readErr)
	}
	if len(entries) != 0 {
		t.Fatalf("package load error wrote output: %+v", entries)
	}
}

func TestBuildEmissionReportsGenericNamedTypeWithoutFalseExtern(t *testing.T) {
	moduleDir := t.TempDir()
	writeTestFile(t, filepath.Join(moduleDir, "go.mod"), "module example.com/genericfixture\n\ngo 1.22\n")
	writeTestFile(t, filepath.Join(moduleDir, "model", "model.go"), `package model

type Box[T any] struct {
	Value T
}
`)
	emission, err := BuildEmission(Config{
		GoImportPath:      "example.com/genericfixture/model",
		WorkingDirectory:  moduleDir,
		OutputRoot:        t.TempDir(),
		HaxePackagePrefix: "goextern",
	})
	if err != nil {
		t.Fatalf("BuildEmission failed: %v", err)
	}
	for _, file := range emission.Files {
		if filepath.Base(file.Name) == "Box.hx" {
			t.Fatalf("generic Box must not become a non-generic extern: %+v", file)
		}
	}
	found := false
	for _, fallback := range emission.DynamicFallbacks {
		if fallback.Symbol == "Box" && fallback.Position == "type" && fallback.Reason == "generic_named_type" {
			found = true
		}
	}
	if !found {
		t.Fatalf("missing generic named type fallback: %+v", emission.DynamicFallbacks)
	}
}

func TestUnsupportedNamedValueSliceDoesNotScheduleDependency(t *testing.T) {
	pkg := types.NewPackage("example.com/model", "model")
	typeName := types.NewTypeName(token.NoPos, pkg, "Value", nil)
	named := types.NewNamed(typeName, types.NewStruct(nil, nil), nil)
	scheduled := make([]*types.TypeName, 0)
	ctx := mappingContext{
		currentPackagePath: "example.com/root",
		haxePackagePrefix:  "goextern",
		schedule: func(candidate *types.TypeName) {
			scheduled = append(scheduled, candidate)
		},
	}

	mapped, reason := mapNativeSliceFieldElementWithReason(named, ctx)
	if mapped != "Dynamic" || reason != "slice_element_abi" {
		t.Fatalf("unsupported named value slice mismatch: got %q (%q)", mapped, reason)
	}
	if len(scheduled) != 0 {
		t.Fatalf("unsupported field scheduled an unused dependency: %+v", scheduled)
	}
}

func TestBuildEmissionExportsSupportedStructFieldsAndReportsSkippedFields(t *testing.T) {
	moduleDir := t.TempDir()
	writeTestFile(t, filepath.Join(moduleDir, "go.mod"), "module example.com/goexternstructfixture\n\ngo 1.22\n")
	writeTestFile(t, filepath.Join(moduleDir, "model", "record.go"), `package model

import "time"

type Related struct {
	Label string
}

type Record struct {
	ID       string
	Count    int
	Active   bool
	Values   []int
	Lookup   map[string]int
	Link     *Related
	Optional *int
	Related
	Inline   struct{ X int }
	Created  time.Time
	private  string
}
`)

	emission, err := BuildEmission(Config{
		GoImportPath:      "example.com/goexternstructfixture/model",
		WorkingDirectory:  moduleDir,
		OutputRoot:        t.TempDir(),
		HaxePackagePrefix: "goextern",
	})
	if err != nil {
		t.Fatalf("BuildEmission failed for struct fixture: %v", err)
	}

	record := fileContentsByName(t, emission, "Record.hx")
	related := fileContentsByName(t, emission, "Related.hx")
	if !strings.Contains(related, "@:go.struct") || !strings.Contains(related, "@:go.name(\"Label\")\n\tpublic var label:String;") {
		t.Fatalf("Related.hx must expose its exact string field and zero-value contract:\n%s", related)
	}
	for _, snippet := range []string{
		"@:go.struct",
		"public function new();",
		"@:go.name(\"ID\")\n\tpublic var id:String;",
		"@:go.name(\"Count\")\n\tpublic var count:Int;",
		"@:go.name(\"Active\")\n\tpublic var active:Bool;",
		"@:go.name(\"Values\")\n\tpublic var values:go.NativeSlice<Int>;",
		"@:go.name(\"Link\")\n\tpublic var link:Related;",
	} {
		if !strings.Contains(record, snippet) {
			t.Fatalf("Record.hx missing generated struct contract %q\n%s", snippet, record)
		}
	}
	for _, forbidden := range []string{"var related", "var inline", "var created", "var lookup", "var optional", "private", "dynamic"} {
		if strings.Contains(strings.ToLower(record), forbidden) {
			t.Fatalf("Record.hx must not expose unsupported or unexported field %q\n%s", forbidden, record)
		}
	}

	wantFallbacks := map[string]string{
		"Record.Related":  "embedded_field",
		"Record.Inline":   "struct",
		"Record.Created":  "named_value_field_abi",
		"Record.Lookup":   "map_field_abi",
		"Record.Optional": "pointer_field_abi",
	}
	if len(emission.DynamicFallbacks) != len(wantFallbacks) {
		t.Fatalf("unexpected struct-field fallback count: got %+v", emission.DynamicFallbacks)
	}
	for _, fallback := range emission.DynamicFallbacks {
		if wantReason, ok := wantFallbacks[fallback.Symbol]; ok {
			if fallback.Position != "field:"+strings.TrimPrefix(fallback.Symbol, "Record.") || fallback.Reason != wantReason {
				t.Fatalf("unexpected struct-field fallback: %+v", fallback)
			}
			delete(wantFallbacks, fallback.Symbol)
		}
	}
	if len(wantFallbacks) != 0 {
		t.Fatalf("missing struct-field fallbacks: %+v (all: %+v)", wantFallbacks, emission.DynamicFallbacks)
	}
}

func TestBuildEmissionFmtMultiReturnBoundary(t *testing.T) {
	cfg := Config{
		GoImportPath:      "fmt",
		OutputRoot:        t.TempDir(),
		HaxePackagePrefix: "goextern",
	}

	emission, err := BuildEmission(cfg)
	if err != nil {
		t.Fatalf("BuildEmission failed: %v", err)
	}

	pkgFile := fileContentsByName(t, emission, "FmtPkg.hx")
	if !strings.Contains(pkgFile, "public static function sprintf(format:String, a:haxe.Rest<Dynamic>):String;") {
		t.Fatalf("single-return fmt.Sprintf should stay strongly typed")
	}
	for _, signature := range []string{
		"@:go.tupleReturn\n\t@:go.name(\"Fprint\")\n\tpublic static function fprint(w:goextern.io.Writer, a:haxe.Rest<Dynamic>):FprintResult;",
		"@:go.tupleReturn\n\t@:go.name(\"Fscan\")\n\tpublic static function fscan(r:goextern.io.Reader, a:haxe.Rest<Dynamic>):FscanResult;",
		"@:go.tupleReturn\n\t@:go.name(\"Sscanf\")\n\tpublic static function sscanf(str:String, format:String, a:haxe.Rest<Dynamic>):SscanfResult;",
	} {
		if !strings.Contains(pkgFile, signature) {
			t.Fatalf("supported multi-return signature should use a generated typed carrier: %s", signature)
		}
	}
	fprintResult := fileContentsByName(t, emission, "FprintResult.hx")
	for _, snippet := range []string{
		"class FprintResult",
		"public var n(default, null):Int;",
		"public var err(default, null):Null<go.Error>;",
		"public function new(n:Int, err:Null<go.Error>)",
	} {
		if !strings.Contains(fprintResult, snippet) {
			t.Fatalf("FprintResult.hx missing generated carrier snippet %q\n%s", snippet, fprintResult)
		}
	}
	if strings.Contains(pkgFile, "@:go.valueError\n\t@:go.name(\"Fprint\")") {
		t.Fatalf("fmt.Fprint returns (int, error), but goextern must not apply @:go.valueError without an explicit go.Result<T> facade")
	}
}

func TestBuildEmissionTimeTupleReturnCarrierUsesNamedFields(t *testing.T) {
	cfg := Config{
		GoImportPath:      "time",
		OutputRoot:        t.TempDir(),
		HaxePackagePrefix: "goextern",
	}

	emission, err := BuildEmission(cfg)
	if err != nil {
		t.Fatalf("BuildEmission failed: %v", err)
	}

	timeFile := fileContentsByName(t, emission, "Time.hx")
	if !strings.Contains(timeFile, "@:go.tupleReturn\n\t@:go.name(\"Zone\")\n\tpublic function zone():TimeZoneResult;") {
		t.Fatalf("Time.Zone should use a typed tuple carrier")
	}
	if !strings.Contains(timeFile, "@:go.tupleReturn\n\t@:go.name(\"Date\")\n\tpublic function date():TimeDateResult;") {
		t.Fatalf("Time.Date should use a typed tuple carrier instead of Dynamic")
	}
	if !strings.Contains(timeFile, "@:go.tupleReturn\n\t@:go.name(\"MarshalBinary\")\n\tpublic function marshalBinary():TimeMarshalBinaryResult;") {
		t.Fatalf("Time.MarshalBinary should use a typed tuple carrier for []byte,error")
	}

	zoneResult := fileContentsByName(t, emission, "TimeZoneResult.hx")
	for _, snippet := range []string{
		"class TimeZoneResult",
		"public var name(default, null):String;",
		"public var offset(default, null):Int;",
		"public function new(name:String, offset:Int)",
	} {
		if !strings.Contains(zoneResult, snippet) {
			t.Fatalf("TimeZoneResult.hx missing generated carrier snippet %q\n%s", snippet, zoneResult)
		}
	}

	dateResult := fileContentsByName(t, emission, "TimeDateResult.hx")
	for _, snippet := range []string{
		"class TimeDateResult",
		"public var year(default, null):Int;",
		"public var month(default, null):Month;",
		"public var day(default, null):Int;",
		"public function new(year:Int, month:Month, day:Int)",
	} {
		if !strings.Contains(dateResult, snippet) {
			t.Fatalf("TimeDateResult.hx missing generated carrier snippet %q\n%s", snippet, dateResult)
		}
	}

	marshalBinaryResult := fileContentsByName(t, emission, "TimeMarshalBinaryResult.hx")
	for _, snippet := range []string{
		"class TimeMarshalBinaryResult",
		"public var value1(default, null):go.NativeSlice<Int>;",
		"public var value2(default, null):Null<go.Error>;",
		"public function new(value1:go.NativeSlice<Int>, value2:Null<go.Error>)",
	} {
		if !strings.Contains(marshalBinaryResult, snippet) {
			t.Fatalf("TimeMarshalBinaryResult.hx missing generated carrier snippet %q\n%s", snippet, marshalBinaryResult)
		}
	}
}

func TestBuildEmissionContextIncludesInterfaceAndPackageClass(t *testing.T) {
	cfg := Config{
		GoImportPath:      "context",
		OutputRoot:        t.TempDir(),
		HaxePackagePrefix: "goextern",
	}

	emission, err := BuildEmission(cfg)
	if err != nil {
		t.Fatalf("BuildEmission failed: %v", err)
	}

	contextType := fileContentsByName(t, emission, "Context.hx")
	if !strings.Contains(contextType, "extern interface Context") {
		t.Fatalf("Context.hx must emit Context as extern interface")
	}

	contextPkg := fileContentsByName(t, emission, "ContextPkg.hx")
	if !strings.Contains(contextPkg, "@:go.name(\"Background\")") {
		t.Fatalf("ContextPkg.hx missing Background mapping")
	}

	cancelFunc := fileContentsByName(t, emission, "CancelFunc.hx")
	if !strings.Contains(cancelFunc, "extern class CancelFunc {}") {
		t.Fatalf("CancelFunc.hx must emit empty extern class on a single line for formatter-stable fixtures")
	}
}

func TestBuildEmissionErrorsPackageClass(t *testing.T) {
	cfg := Config{
		GoImportPath:      "errors",
		OutputRoot:        t.TempDir(),
		HaxePackagePrefix: "goextern",
	}

	emission, err := BuildEmission(cfg)
	if err != nil {
		t.Fatalf("BuildEmission failed: %v", err)
	}

	if len(emission.Files) != 1 {
		t.Fatalf("expected 1 emitted file for errors package, got %d", len(emission.Files))
	}
	if emission.Files[0].Name != filepath.Join("errors", "ErrorsPkg.hx") {
		t.Fatalf("unexpected file name: %s", emission.Files[0].Name)
	}

	pkgFile := emission.Files[0].Contents
	if !strings.Contains(pkgFile, "@:go.import(\"errors\")") {
		t.Fatalf("ErrorsPkg.hx missing @:go.import metadata")
	}
	for _, symbol := range []string{"As", "Is", "Join", "New", "Unwrap"} {
		if !strings.Contains(pkgFile, "@:go.name(\""+symbol+"\")") {
			t.Fatalf("ErrorsPkg.hx missing symbol mapping for %s", symbol)
		}
	}
}

func TestMapTypeWithReasonReportsStableDynamicReasonCodes(t *testing.T) {
	t.Parallel()

	ctx := mappingContext{
		currentPackagePath: "example/current",
	}

	externalPkg := types.NewPackage("net/http", "http")
	externalNamed := types.NewNamed(
		types.NewTypeName(token.NoPos, externalPkg, "Request", nil),
		types.NewStruct(nil, nil),
		nil,
	)
	genericNamed := types.NewNamed(
		types.NewTypeName(token.NoPos, externalPkg, "Box", nil),
		types.NewStruct(nil, nil),
		nil,
	)
	emptyInterface := types.NewInterfaceType(nil, nil)
	emptyInterface.Complete()
	readSig := types.NewSignatureType(nil, nil, nil, types.NewTuple(), types.NewTuple(), false)
	readMethod := types.NewFunc(token.NoPos, nil, "Read", readSig)
	nonEmptyInterface := types.NewInterfaceType([]*types.Func{readMethod}, nil)
	nonEmptyInterface.Complete()
	constraint := types.NewInterfaceType(nil, nil)
	constraint.Complete()
	typeParam := types.NewTypeParam(types.NewTypeName(token.NoPos, nil, "T", nil), constraint)
	genericNamed.SetTypeParams([]*types.TypeParam{typeParam})
	privateNamed := types.NewNamed(
		types.NewTypeName(token.NoPos, externalPkg, "privateRecord", nil),
		types.NewStruct(nil, nil),
		nil,
	)

	cases := []struct {
		name   string
		typ    types.Type
		reason string
	}{
		{name: "callback", typ: types.NewSignatureType(nil, nil, nil, types.NewTuple(), types.NewTuple(), false), reason: "callback_signature"},
		{name: "unsupported map key", typ: types.NewMap(types.Typ[types.Int], types.Typ[types.String]), reason: "unsupported_map_key"},
		{name: "fixed array", typ: types.NewArray(types.Typ[types.Int], 4), reason: "fixed_array"},
		{name: "struct", typ: types.NewStruct(nil, nil), reason: "struct"},
		{name: "empty interface", typ: emptyInterface, reason: "empty_interface"},
		{name: "non-empty interface", typ: nonEmptyInterface, reason: "non_empty_interface"},
		{name: "channel", typ: types.NewChan(types.SendRecv, types.Typ[types.Int]), reason: "channel"},
		{name: "generic type parameter", typ: typeParam, reason: "type_parameter"},
		{name: "generic named type", typ: genericNamed, reason: "generic_named_type"},
		{name: "unexported named type", typ: privateNamed, reason: "unexported_named_type"},
		{name: "unsafe pointer", typ: types.Typ[types.UnsafePointer], reason: "unsafe_pointer"},
	}

	mappedExternal, reason := mapTypeWithReason(externalNamed, ctx)
	if mappedExternal != "net.http.Request" || reason != "" {
		t.Fatalf("external named type must map precisely: got %q (%q)", mappedExternal, reason)
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mapped, reason := mapTypeWithReason(tc.typ, ctx)
			if !containsDynamicType(mapped) {
				t.Fatalf("expected %s to map to a Dynamic-containing type, got %q", tc.name, mapped)
			}
			if reason != tc.reason {
				t.Fatalf("reason mismatch for %s: got %q, want %q", tc.name, reason, tc.reason)
			}
		})
	}
}

func TestWriteDynamicFallbackReportIsDeterministic(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "reports", "dynamic.json")
	emission := &Emission{
		DynamicFallbacks: []DynamicFallback{
			{Package: "pkg", Symbol: "B", Position: "param:x", GoType: "func()", Reason: "callback_signature"},
			{Package: "pkg", Symbol: "A", Position: "result:1", GoType: "struct{}", Reason: "struct"},
			{Package: "pkg", Symbol: "A", Position: "result:1", GoType: "struct{}", Reason: "struct"},
		},
	}

	if err := writeDynamicFallbackReport(target, emission); err != nil {
		t.Fatalf("writeDynamicFallbackReport failed: %v", err)
	}

	payload, err := os.ReadFile(target)
	if err != nil {
		t.Fatalf("read report failed: %v", err)
	}

	var report dynamicFallbackReport
	if err := json.Unmarshal(payload, &report); err != nil {
		t.Fatalf("report is not valid JSON: %v\n%s", err, string(payload))
	}
	if report.SchemaVersion != 1 {
		t.Fatalf("schemaVersion mismatch: got %d", report.SchemaVersion)
	}
	if len(report.Fallbacks) != 2 {
		t.Fatalf("expected sorted/deduplicated fallback list, got %d entries: %+v", len(report.Fallbacks), report.Fallbacks)
	}
	if report.Fallbacks[0].Symbol != "A" || report.Fallbacks[1].Symbol != "B" {
		t.Fatalf("fallback report ordering is not stable: %+v", report.Fallbacks)
	}
}

func TestRunWritesDynamicFallbackReport(t *testing.T) {
	dir := t.TempDir()
	outDir := filepath.Join(dir, "out")
	reportPath := filepath.Join(dir, "dynamic.json")

	var stdout bytes.Buffer
	if err := run([]string{
		"-package", "fmt",
		"-out", outDir,
		"-dynamic-report", reportPath,
	}, &stdout, io.Discard); err != nil {
		t.Fatalf("run failed: %v", err)
	}
	if !strings.Contains(stdout.String(), "precision=partial;") || !strings.Contains(stdout.String(), "fallbacks=") {
		t.Fatalf("run must report partial precision and fallback count: %q", stdout.String())
	}

	payload, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("read report failed: %v", err)
	}

	var report dynamicFallbackReport
	if err := json.Unmarshal(payload, &report); err != nil {
		t.Fatalf("report is not valid JSON: %v\n%s", err, string(payload))
	}
	if report.SchemaVersion != 1 {
		t.Fatalf("schemaVersion mismatch: got %d", report.SchemaVersion)
	}
	if len(report.Fallbacks) == 0 {
		t.Fatalf("expected fmt to expose at least one Dynamic boundary")
	}

	found := false
	for _, fallback := range report.Fallbacks {
		if fallback.Package == "fmt" && fallback.Symbol == "Fprint" && fallback.Position == "param:a" && fallback.GoType == "[]any" && fallback.Reason == "empty_interface" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected fmt.Fprint writer boundary in report, got %+v", report.Fallbacks)
	}
}

func TestWriteEmissionRemovesStaleHxFiles(t *testing.T) {
	dir := t.TempDir()

	initial := &Emission{
		OutputDir: dir,
		RootKey:   "root-a",
		Files: []EmittedFile{
			{Name: "First.hx", Contents: "package;\n\nextern class First {}\n"},
		},
	}

	if err := writeEmission(initial); err != nil {
		t.Fatalf("writeEmission initial failed: %v", err)
	}

	next := &Emission{
		OutputDir: dir,
		RootKey:   "root-a",
		Files: []EmittedFile{
			{Name: "Second.hx", Contents: "package;\n\nextern class Second {}\n"},
		},
	}

	if err := writeEmission(next); err != nil {
		t.Fatalf("writeEmission next failed: %v", err)
	}

	if exists(filepath.Join(dir, "First.hx")) {
		t.Fatalf("stale file First.hx should have been removed")
	}
	if !exists(filepath.Join(dir, "Second.hx")) {
		t.Fatalf("Second.hx should exist")
	}
}

func TestWriteEmissionPreservesUnownedFilesAndRejectsConflicts(t *testing.T) {
	dir := t.TempDir()
	unowned := filepath.Join(dir, "user", "Custom.hx")
	writeTestFile(t, unowned, "package user;\n\nclass Custom {}\n")

	emission := &Emission{
		OutputDir: dir,
		RootKey:   "root-a",
		Files: []EmittedFile{
			{Name: "generated/Binding.hx", Contents: "package generated;\n\nextern class Binding {}\n"},
		},
	}
	if err := writeEmission(emission); err != nil {
		t.Fatalf("writeEmission failed: %v", err)
	}
	if !exists(unowned) {
		t.Fatalf("unowned Haxe file must be preserved")
	}

	conflict := &Emission{
		OutputDir: dir,
		RootKey:   "root-b",
		Files: []EmittedFile{
			{Name: "user/Custom.hx", Contents: "package user;\n\nextern class Different {}\n"},
		},
	}
	expectGeneratorErrorCode(t, writeEmission(conflict), "unowned_output_conflict")
	payload, err := os.ReadFile(unowned)
	if err != nil {
		t.Fatalf("read unowned file after conflict: %v", err)
	}
	if string(payload) != "package user;\n\nclass Custom {}\n" {
		t.Fatalf("conflicting generation changed an unowned file")
	}
}

func TestWriteEmissionRejectsModifiedOwnedFile(t *testing.T) {
	dir := t.TempDir()
	initial := &Emission{
		OutputDir: dir,
		RootKey:   "root-a",
		Files: []EmittedFile{
			{Name: "pkg/Binding.hx", Contents: "original\n"},
		},
	}
	if err := writeEmission(initial); err != nil {
		t.Fatalf("write initial emission: %v", err)
	}
	writeTestFile(t, filepath.Join(dir, "pkg", "Binding.hx"), "user edit\n")

	next := &Emission{
		OutputDir: dir,
		RootKey:   "root-a",
		Files: []EmittedFile{
			{Name: "pkg/Binding.hx", Contents: "next\n"},
		},
	}
	expectGeneratorErrorCode(t, writeEmission(next), "owned_output_modified")
	payload, err := os.ReadFile(filepath.Join(dir, "pkg", "Binding.hx"))
	if err != nil {
		t.Fatalf("read modified owned file: %v", err)
	}
	if string(payload) != "user edit\n" {
		t.Fatalf("modified owned file was overwritten")
	}
}

func TestWriteEmissionKeepsSharedOutputUntilLastOwnerDropsIt(t *testing.T) {
	dir := t.TempDir()
	shared := EmittedFile{Name: "dep/Shared.hx", Contents: "shared\n"}
	if err := writeEmission(&Emission{OutputDir: dir, RootKey: "root-a", Files: []EmittedFile{shared}}); err != nil {
		t.Fatalf("write first shared owner: %v", err)
	}
	conflict := EmittedFile{Name: shared.Name, Contents: "different\n"}
	expectGeneratorErrorCode(
		t,
		writeEmission(&Emission{OutputDir: dir, RootKey: "root-b", Files: []EmittedFile{conflict}}),
		"owned_output_conflict",
	)
	if err := writeEmission(&Emission{OutputDir: dir, RootKey: "root-b", Files: []EmittedFile{shared}}); err != nil {
		t.Fatalf("write second shared owner: %v", err)
	}
	if err := writeEmission(&Emission{OutputDir: dir, RootKey: "root-a"}); err != nil {
		t.Fatalf("drop first shared owner: %v", err)
	}
	sharedPath := filepath.Join(dir, "dep", "Shared.hx")
	if !exists(sharedPath) {
		t.Fatalf("shared output was removed while another root still owned it")
	}
	if err := writeEmission(&Emission{OutputDir: dir, RootKey: "root-b"}); err != nil {
		t.Fatalf("drop final shared owner: %v", err)
	}
	if exists(sharedPath) {
		t.Fatalf("shared output remained after its final owner dropped it")
	}
}

func TestWriteEmissionRejectsConflictingExistingRootClaims(t *testing.T) {
	dir := t.TempDir()
	manifestDir := filepath.Join(dir, ".goextern", "roots")
	writeTestFile(t, filepath.Join(manifestDir, "root-a.json"), `{
	"schemaVersion": 1,
	"rootKey": "root-a",
	"root": {},
	"files": [{"path": "dep/Shared.hx", "sha256": "0000000000000000000000000000000000000000000000000000000000000000"}]
}
`)
	writeTestFile(t, filepath.Join(manifestDir, "root-b.json"), `{
	"schemaVersion": 1,
	"rootKey": "root-b",
	"root": {},
	"files": [{"path": "dep/Shared.hx", "sha256": "1111111111111111111111111111111111111111111111111111111111111111"}]
}
`)

	emission := &Emission{
		OutputDir: dir,
		RootKey:   "root-c",
		Files: []EmittedFile{
			{Name: "root/Binding.hx", Contents: "binding\n"},
		},
	}
	expectGeneratorErrorCode(t, writeEmission(emission), "owned_output_conflict")
	if exists(filepath.Join(dir, "root", "Binding.hx")) {
		t.Fatalf("conflicting ownership records must fail before writing")
	}
}

func TestWriteEmissionRejectsUnsafeRelativePathBeforeWriting(t *testing.T) {
	dir := t.TempDir()
	emission := &Emission{
		OutputDir: dir,
		RootKey:   "root-a",
		Files: []EmittedFile{
			{Name: "safe/Binding.hx", Contents: "safe\n"},
			{Name: "../escape.hx", Contents: "unsafe\n"},
		},
	}
	expectGeneratorErrorCode(t, writeEmission(emission), "unsafe_output_path")
	if exists(filepath.Join(dir, ".goextern")) {
		t.Fatalf("unsafe plan must fail before creating ownership state")
	}
	if exists(filepath.Join(dir, "safe", "Binding.hx")) {
		t.Fatalf("unsafe plan wrote a valid sibling before the error")
	}
}

func TestWriteEmissionRejectsSymbolicLinkBoundary(t *testing.T) {
	dir := t.TempDir()
	outside := t.TempDir()
	link := filepath.Join(dir, "linked")
	if err := os.Symlink(outside, link); err != nil {
		t.Skipf("symbolic links unavailable: %v", err)
	}
	emission := &Emission{
		OutputDir: dir,
		RootKey:   "root-a",
		Files: []EmittedFile{
			{Name: "linked/Escape.hx", Contents: "unsafe\n"},
		},
	}
	expectGeneratorErrorCode(t, writeEmission(emission), "unsafe_output_path")
	if exists(filepath.Join(outside, "Escape.hx")) {
		t.Fatalf("generator wrote through a symbolic-link boundary")
	}
}

func TestWriteEmissionRejectsOwnershipDirectorySymbolicLink(t *testing.T) {
	dir := t.TempDir()
	outside := t.TempDir()
	if err := os.Symlink(outside, filepath.Join(dir, ".goextern")); err != nil {
		t.Skipf("symbolic links unavailable: %v", err)
	}
	emission := &Emission{
		OutputDir: dir,
		RootKey:   "root-a",
		Files: []EmittedFile{
			{Name: "safe/Binding.hx", Contents: "safe\n"},
		},
	}
	expectGeneratorErrorCode(t, writeEmission(emission), "unsafe_output_path")
	if exists(filepath.Join(outside, "roots", "root-a.json")) {
		t.Fatalf("generator wrote ownership state through a symbolic link")
	}
}

func fileContentsByName(t *testing.T, emission *Emission, fileName string) string {
	t.Helper()

	for _, file := range emission.Files {
		if file.Name == fileName || filepath.Base(file.Name) == fileName {
			return file.Contents
		}
	}
	t.Fatalf("expected file %q not found in emission", fileName)
	return ""
}

func fileContentsByPath(t *testing.T, emission *Emission, filePath string) string {
	t.Helper()
	filePath = filepath.ToSlash(filePath)
	for _, file := range emission.Files {
		if filepath.ToSlash(file.Name) == filePath {
			return file.Contents
		}
	}
	t.Fatalf("expected file %q not found in emission", filePath)
	return ""
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func expectGeneratorErrorCode(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected generator error %q", want)
	}
	var failure *generatorError
	if !errors.As(err, &failure) {
		t.Fatalf("expected generatorError %q, got %T: %v", want, err, err)
	}
	if failure.Code != want {
		t.Fatalf("generator error code mismatch: got %q, want %q (%v)", failure.Code, want, err)
	}
}

func writeTestFile(t *testing.T, path string, contents string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("create fixture directory: %v", err)
	}
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("write fixture file: %v", err)
	}
}
