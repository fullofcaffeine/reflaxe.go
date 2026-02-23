package packages

import (
	"fmt"
	"go/importer"
	"go/types"
	"os"
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
	_ = cfg

	loaded := make([]*Package, 0, len(patterns))
	imp := importer.Default()

	for _, pattern := range patterns {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}

		pkg, err := imp.Import(pattern)
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
