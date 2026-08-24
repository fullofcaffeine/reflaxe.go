package packages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/importer"
	"go/token"
	"go/types"
	"io"
	"os"
	"os/exec"
	"strings"
)

type LoadMode int

const (
	NeedName LoadMode = 1 << iota
	NeedFiles
	NeedCompiledGoFiles
	NeedImports
	NeedDeps
	NeedExportsFile
	NeedTypes
	NeedSyntax
	NeedTypesInfo
	NeedTypesSizes
	NeedModule
)

type Config struct {
	Mode LoadMode
	Dir  string
	Env  []string
}

type Error struct {
	Msg string
}

func (e Error) Error() string {
	return e.Msg
}

type Package struct {
	Name    string
	PkgPath string
	Types   *types.Package
	Errors  []Error
}

func Load(cfg *Config, patterns ...string) ([]*Package, error) {
	loaded := make([]*Package, 0, len(patterns))

	for _, pattern := range patterns {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}

		pkg, err := loadPackage(cfg, pattern)
		if err != nil {
			loaded = append(loaded, &Package{
				Name:    "",
				PkgPath: pattern,
				Types:   nil,
				Errors:  []Error{{Msg: err.Error()}},
			})
			continue
		}

		loaded = append(loaded, &Package{
			Name:    pkg.Name(),
			PkgPath: pkg.Path(),
			Types:   pkg,
			Errors:  nil,
		})
	}

	return loaded, nil
}

type listedPackage struct {
	ImportPath string
	Export     string
}

func loadPackage(cfg *Config, pattern string) (*types.Package, error) {
	args := []string{"list", "-mod=readonly", "-deps", "-export", "-json", pattern}
	cmd := exec.Command("go", args...)
	if cfg != nil {
		cmd.Dir = strings.TrimSpace(cfg.Dir)
		if len(cfg.Env) > 0 {
			cmd.Env = cfg.Env
		}
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		detail := strings.TrimSpace(stderr.String())
		if detail == "" {
			detail = err.Error()
		}
		return nil, fmt.Errorf("go list %q: %s", pattern, detail)
	}

	exportFiles := make(map[string]string)
	decoder := json.NewDecoder(&stdout)
	for {
		var listed listedPackage
		if err := decoder.Decode(&listed); err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("decode go list result for %q: %w", pattern, err)
		}
		if listed.ImportPath != "" && listed.Export != "" {
			exportFiles[listed.ImportPath] = listed.Export
		}
	}

	lookup := func(importPath string) (io.ReadCloser, error) {
		exportFile, ok := exportFiles[importPath]
		if !ok {
			return nil, fmt.Errorf("go list did not report export data for %q", importPath)
		}
		return os.Open(exportFile)
	}
	imp := importer.ForCompiler(token.NewFileSet(), "gc", lookup)
	return imp.Import(pattern)
}

func PrintErrors(pkgs []*Package) int {
	count := 0
	for _, pkg := range pkgs {
		if pkg == nil {
			continue
		}
		for _, pkgErr := range pkg.Errors {
			count++
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", pkgErr.Error())
		}
	}
	return count
}
