package main

import (
	"encoding/json"
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

func Find(id string) model.Record {
	return model.Record{ID: id}
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
	if !strings.Contains(apiFile, "public static function find(id:String):Dynamic;") {
		t.Fatalf("ApiPkg.hx must preserve the honest external named-type fallback:\n%s", apiFile)
	}
	if len(emission.DynamicFallbacks) != 1 {
		t.Fatalf("expected one cross-package fallback, got %+v", emission.DynamicFallbacks)
	}
	fallback := emission.DynamicFallbacks[0]
	if fallback.GoType != "example.com/goexternfixture/model.Record" || fallback.Reason != "external_named_type" {
		t.Fatalf("unexpected cross-package fallback: %+v", fallback)
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
		"@:go.tupleReturn\n\t@:go.name(\"Fprint\")\n\tpublic static function fprint(w:Dynamic, a:haxe.Rest<Dynamic>):FprintResult;",
		"@:go.tupleReturn\n\t@:go.name(\"Fscan\")\n\tpublic static function fscan(r:Dynamic, a:haxe.Rest<Dynamic>):FscanResult;",
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
	if emission.Files[0].Name != "ErrorsPkg.hx" {
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
		exportedTypeNames:  map[string]bool{"Local": true},
	}

	externalPkg := types.NewPackage("net/http", "http")
	externalNamed := types.NewNamed(
		types.NewTypeName(token.NoPos, externalPkg, "Request", nil),
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

	cases := []struct {
		name   string
		typ    types.Type
		reason string
	}{
		{name: "callback", typ: types.NewSignatureType(nil, nil, nil, types.NewTuple(), types.NewTuple(), false), reason: "callback_signature"},
		{name: "external named type", typ: externalNamed, reason: "external_named_type"},
		{name: "unsupported map key", typ: types.NewMap(types.Typ[types.Int], types.Typ[types.String]), reason: "unsupported_map_key"},
		{name: "fixed array", typ: types.NewArray(types.Typ[types.Int], 4), reason: "fixed_array"},
		{name: "struct", typ: types.NewStruct(nil, nil), reason: "struct"},
		{name: "empty interface", typ: emptyInterface, reason: "empty_interface"},
		{name: "non-empty interface", typ: nonEmptyInterface, reason: "non_empty_interface"},
		{name: "channel", typ: types.NewChan(types.SendRecv, types.Typ[types.Int]), reason: "channel"},
		{name: "generic type parameter", typ: typeParam, reason: "type_parameter"},
		{name: "unsafe pointer", typ: types.Typ[types.UnsafePointer], reason: "unsafe_pointer"},
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

	if err := run([]string{
		"-package", "fmt",
		"-out", outDir,
		"-dynamic-report", reportPath,
	}, io.Discard, io.Discard); err != nil {
		t.Fatalf("run failed: %v", err)
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
		if fallback.Package == "fmt" && fallback.Symbol == "Fprint" && fallback.Position == "param:w" && fallback.GoType == "io.Writer" && fallback.Reason == "external_named_type" {
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
		Files: []EmittedFile{
			{Name: "First.hx", Contents: "package;\n\nextern class First {}\n"},
		},
	}

	if err := writeEmission(initial); err != nil {
		t.Fatalf("writeEmission initial failed: %v", err)
	}

	next := &Emission{
		OutputDir: dir,
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

func fileContentsByName(t *testing.T, emission *Emission, fileName string) string {
	t.Helper()

	for _, file := range emission.Files {
		if file.Name == fileName {
			return file.Contents
		}
	}
	t.Fatalf("expected file %q not found in emission", fileName)
	return ""
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
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
