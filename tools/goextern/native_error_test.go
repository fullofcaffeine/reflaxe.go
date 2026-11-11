package main

import (
	"bytes"
	"go/types"
	"path/filepath"
	"strings"
	"testing"
)

func TestNativeErrorUnsupportedStorageReportsExactBoundary(t *testing.T) {
	errorType := types.Universe.Lookup("error").Type()
	ctx := mappingContext{errorRepresentation: nativeErrorRepresentation}
	for _, tc := range []struct {
		name   string
		typ    types.Type
		reason string
	}{
		{"slice", types.NewSlice(errorType), "native_error_collection_abi"},
		{"nested slice", types.NewSlice(types.NewSlice(errorType)), "native_error_collection_abi"},
		{"map value", types.NewMap(types.Typ[types.String], errorType), "native_error_collection_abi"},
		{"map key", types.NewMap(errorType, types.Typ[types.Int]), "native_error_collection_abi"},
		{"pointer", types.NewPointer(errorType), "native_error_pointer_abi"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := mapTypeResult(tc.typ, ctx)
			if got.Haxe != "Dynamic" || got.Reason != tc.reason {
				t.Fatalf("mapping = %+v, want explicit Dynamic fallback with %s", got, tc.reason)
			}
		})
	}
}

func TestNativeErrorRepresentationIsExplicitAndConsistent(t *testing.T) {
	dir := t.TempDir()
	writeTestFile(t, filepath.Join(dir, "go.mod"), "module example.com/errorabi\n\ngo 1.22\n")
	writeTestFile(t, filepath.Join(dir, "api", "api.go"), `package api
type Alias = error
type Box struct { Err error }
func Echo(value error) error { return value }
func Pair() (int, error) { return 1, nil }
func Aliased(value Alias) Alias { return value }
func Collection(value []error) []error { return value }
func Pointer(value *error) *error { return value }
`)
	for _, mode := range []string{"legacy", "native"} {
		var output, diagnostics bytes.Buffer
		err := run([]string{"-package", "example.com/errorabi/api", "-dir", dir,
			"-error-representation", mode, "-stdout"}, &output, &diagnostics)
		if err != nil {
			t.Fatalf("%s generation: %v (%s)", mode, err, &diagnostics)
		}
		text := output.String()
		name := "go.Error"
		if mode == "native" {
			name = "go.NativeError"
		}
		for _, want := range []string{
			"function echo(value:" + name + "):" + name,
			"Null<" + name + ">",
		} {
			if !strings.Contains(text, want) {
				t.Fatalf("%s missing %q:\n%s", mode, want, text)
			}
		}
		if mode == "native" {
			for _, want := range []string{"function aliased(value:go.NativeError):go.NativeError", "var err:go.NativeError"} {
				if !strings.Contains(text, want) {
					t.Fatalf("native missing %q:\n%s", want, text)
				}
			}
			if strings.Contains(text, "NativeSlice<go.NativeError>") {
				t.Fatal("native mode falsely admitted an error slice without element adapters")
			}
		}
	}
}

func TestInvalidErrorRepresentationFailsBeforePackageLoad(t *testing.T) {
	var output, diagnostics bytes.Buffer
	err := run([]string{"-package", "not-a-real-package", "-error-representation", "typo", "-stdout"}, &output, &diagnostics)
	if err == nil || !strings.Contains(err.Error(), "error representation") {
		t.Fatalf("expected explicit representation diagnostic, got %v", err)
	}
}

func TestNativeErrorOwnershipPreventsImplicitRepresentationChange(t *testing.T) {
	dir := t.TempDir()
	first := &Emission{OutputDir: dir, RootKey: "same-api", Root: ownershipRoot{ErrorRepresentation: nativeErrorRepresentation}, Files: []EmittedFile{{Name: "Example.hx", Contents: "native"}}}
	if err := writeEmission(first); err != nil {
		t.Fatal(err)
	}
	second := &Emission{OutputDir: dir, RootKey: "same-api", Files: []EmittedFile{{Name: "Example.hx", Contents: "legacy"}}}
	if err := writeEmission(second); err == nil || !strings.Contains(err.Error(), "error_representation_conflict") {
		t.Fatalf("expected representation conflict, got %v", err)
	}
	if err := writeEmission(first); err != nil {
		t.Fatalf("identical regeneration failed: %v", err)
	}
}
