package main

import (
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
