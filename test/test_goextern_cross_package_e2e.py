#!/usr/bin/env python3

from __future__ import annotations

import hashlib
import json
import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest


ROOT = Path(__file__).resolve().parent.parent
GOEXTERN = ROOT / "tools" / "goextern"


def run(command: list[str], cwd: Path, timeout: int = 240) -> subprocess.CompletedProcess[str]:
    env = os.environ.copy()
    env["HAXE_NO_SERVER"] = "1"
    return subprocess.run(
        command,
        cwd=cwd,
        env=env,
        capture_output=True,
        text=True,
        timeout=timeout,
    )


def write(path: Path, contents: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(contents, encoding="utf-8")


def tree_digest(root: Path) -> str:
    digest = hashlib.sha256()
    for path in sorted(candidate for candidate in root.rglob("*") if candidate.is_file()):
        digest.update(path.relative_to(root).as_posix().encode("utf-8"))
        digest.update(b"\0")
        digest.update(path.read_bytes())
        digest.update(b"\0")
    return digest.hexdigest()


def compiler_command(module_root: Path, extern_root: Path) -> list[str]:
    return [
        shutil.which("haxe") or "haxe",
        "-cp",
        str(module_root),
        "-cp",
        str(extern_root),
        "-cp",
        str(ROOT / "src"),
        "-cp",
        str(ROOT / "std"),
        "-cp",
        str(ROOT / "std" / "go" / "_std"),
        "-cp",
        str(ROOT / "vendor" / "reflaxe" / "src"),
        "-D",
        "reflaxe=4.0.0-beta",
        "-D",
        "reflaxe.go=0.0.0-development",
        "--macro",
        'nullSafety("reflaxe.go")',
        "--macro",
        "reflaxe.go.CompilerBootstrap.Start()",
        "--macro",
        "reflaxe.go.CompilerInit.Start()",
        "-D",
        f"go_output={module_root}",
        "-D",
        f"reflaxe_go_project={module_root / 'reflaxe-go-project.json'}",
        "-D",
        "reflaxe_go_strict_examples",
        "-D",
        "no-traces",
        "-D",
        "reflaxe.dont_output_metadata_id",
        "-main",
        "Main",
    ]


@unittest.skipUnless(shutil.which("haxe") and shutil.which("go"), "requires Haxe and Go")
class GoexternCrossPackageEndToEndTest(unittest.TestCase):
    def test_generated_graph_compiles_and_runs_through_haxe_go(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-goextern-graph-") as raw:
            module_root = Path(raw)
            extern_root = module_root / "externs"
            generated_root = extern_root / "goextern"
            report_path = module_root / "fallbacks.json"

            write(module_root / "go.mod", "module example.com/goexternfixture\n\ngo 1.22\n")
            write(
                module_root / "contract" / "context.go",
                """package contract

type Context interface {
	Err() error
}
""",
            )
            write(
                module_root / "detail" / "detail.go",
                """package detail

type Detail struct {
	Bonus int
}
""",
            )
            write(
                module_root / "model" / "item.go",
                """package model

import "example.com/goexternfixture/detail"

type Item struct {
	Value  int
	Detail *detail.Detail
}
""",
            )
            write(
                module_root / "api" / "api.go",
                """package api

import (
	"example.com/goexternfixture/contract"
	"example.com/goexternfixture/detail"
	"example.com/goexternfixture/model"
)

type localContext struct{}
func (localContext) Err() error { return nil }

type ItemAlias = model.Item

type Page struct {
	Items []*model.Item
}

func Background() contract.Context { return localContext{} }
func Alias(seed int) *ItemAlias {
	return &model.Item{Value: seed, Detail: &detail.Detail{}}
}
func Value(seed int) detail.Detail { return detail.Detail{Bonus: seed} }
func Lookup(ctx contract.Context, seed int) (item *model.Item, err error) {
	return &model.Item{Value: seed, Detail: &detail.Detail{Bonus: 1}}, ctx.Err()
}
func List(ctx contract.Context, seed int) (page *Page, err error) {
	return &Page{Items: []*model.Item{{Value: seed, Detail: &detail.Detail{Bonus: 1}}}}, ctx.Err()
}
""",
            )

            generator = [
                shutil.which("go") or "go",
                "run",
                ".",
                "-package",
                "example.com/goexternfixture/api",
                "-dir",
                str(module_root),
                "-out",
                str(generated_root),
                "-haxe-package",
                "goextern",
                "-dynamic-report",
                str(report_path),
            ]
            first = run(generator, GOEXTERN)
            self.assertEqual(0, first.returncode, first.stdout + first.stderr)
            self.assertIn("precision=exact; fallbacks=0", first.stdout)
            first_digest = tree_digest(extern_root)
            second = run(generator, GOEXTERN)
            self.assertEqual(0, second.returncode, second.stdout + second.stderr)
            self.assertEqual(first_digest, tree_digest(extern_root))
            report = json.loads(report_path.read_text(encoding="utf-8"))
            self.assertEqual([], report["fallbacks"])

            write(
                module_root / "Main.hx",
                """import goextern.example_com.goexternfixture.api.ApiPkg;
import goextern.example_com.goexternfixture.api.ItemAlias;

@:go.import("fmt")
@:go.package("fmt")
extern class GoFmt {
	@:go.name("Println")
	public static function println(value:Int):Void;
}

class Main {
	static function main():Void {
		final context = ApiPkg.background();
		final lookup = ApiPkg.lookup(context, 20);
		final listed = ApiPkg.list(context, 20);
		final alias:ItemAlias = ApiPkg.alias(0);
		final value = ApiPkg.value(0);
		final first = listed.page.items[0];
		if (lookup.err == null && listed.err == null && alias.value == 0) {
			GoFmt.println(lookup.item.value + lookup.item.detail.bonus + first.value + first.detail.bonus + value.bonus);
		} else {
			GoFmt.println(-1);
		}
	}
}
""",
            )
            project = {
                "schemaVersion": 1,
                "mode": "existing-module",
                "moduleRoot": ".",
                "packageDir": ".",
                "packageName": "main",
                "runtimeDir": "hxrt",
                "entrypoint": {"kind": "compiler-main"},
                "build": {"kind": "none"},
            }
            write(
                module_root / "reflaxe-go-project.json",
                json.dumps(project, indent=2) + "\n",
            )

            compiled = run(compiler_command(module_root, extern_root), module_root)
            self.assertEqual(0, compiled.returncode, compiled.stdout + compiled.stderr)
            generated = "\n".join(
                path.read_text(encoding="utf-8")
                for path in sorted(module_root.glob("haxego_generated_*.go"))
            )
            for snippet in ["api.Background()", "api.Lookup(", "api.List(", "api.Value(", ".Items[", ".Detail.Bonus"]:
                self.assertIn(snippet, generated)

            for command in (
                ["go", "vet", ".", "./api", "./contract", "./detail", "./model"],
                ["go", "test", "./..."],
            ):
                checked = run(command, module_root)
                self.assertEqual(
                    0,
                    checked.returncode,
                    checked.stdout + checked.stderr + "\nGenerated application:\n" + generated,
                )
            executed = run(["go", "run", "."], module_root)
            self.assertEqual(0, executed.returncode, executed.stdout + executed.stderr)
            self.assertEqual("42\n", executed.stdout)


if __name__ == "__main__":
    unittest.main()
