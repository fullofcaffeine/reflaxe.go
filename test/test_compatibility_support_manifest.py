#!/usr/bin/env python3

from __future__ import annotations

import copy
import json
import subprocess
import tempfile
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
GENERATOR = REPO_ROOT / "scripts" / "compatibility" / "generate_support_manifest.py"
SOURCE = REPO_ROOT / "docs" / "compatibility-support-source.json"
MANIFEST = REPO_ROOT / "docs" / "compatibility-support-manifest.json"
MATRIX_DOC = REPO_ROOT / "docs" / "compatibility-support-matrix.md"
RELEASE_DOC = REPO_ROOT / "docs" / "compatibility-release-status.md"
TOOLCHAIN_POLICY = REPO_ROOT / "docs" / "toolchain-policy.json"
PORTABLE_INVENTORY = REPO_ROOT / "test" / "portable_stdlib_inventory.json"

ALLOWED_STATES = {
    "semantic-diff-supported",
    "compile-go-test-run-supported",
    "compile-only",
    "experimental",
    "compatibility-only",
    "excluded",
}
RELEASE_CLAIM = (
    "Haxe.Go is a pre-1.0 beta for pinned, application-qualified portable workloads "
    "on the admitted toolchain, platform, and operation/member surface."
)


def run_generator(*args: str) -> subprocess.CompletedProcess[str]:
    return subprocess.run(
        ["python3", str(GENERATOR), *args],
        cwd=REPO_ROOT,
        capture_output=True,
        text=True,
    )


class CompatibilitySupportManifestTest(unittest.TestCase):
    def load_manifest(self) -> dict[str, object]:
        return json.loads(MANIFEST.read_text(encoding="utf-8"))

    def test_generated_artifacts_are_current_and_versioned(self) -> None:
        completed = run_generator("--check")
        self.assertEqual(completed.returncode, 0, completed.stdout + completed.stderr)

        manifest = self.load_manifest()
        self.assertEqual(manifest["schema_version"], 1)
        self.assertEqual(manifest["kind"], "haxe.go-compatibility-support-manifest")
        self.assertEqual(manifest["generated_by"], "scripts/compatibility/generate_support_manifest.py")
        self.assertEqual(manifest["release_claim"]["statement"], RELEASE_CLAIM)
        self.assertEqual(set(manifest["evidence_states"]), ALLOWED_STATES)

    def test_toolchains_and_portable_modules_are_derived_from_governed_inputs(self) -> None:
        manifest = self.load_manifest()
        toolchain = json.loads(TOOLCHAIN_POLICY.read_text(encoding="utf-8"))
        portable = json.loads(PORTABLE_INVENTORY.read_text(encoding="utf-8"))

        self.assertEqual(manifest["toolchains"], toolchain)
        generated_modules = manifest["portable_stdlib"]["modules"]
        self.assertEqual(len(generated_modules), len(portable["modules"]))
        self.assertEqual(
            [entry["module"] for entry in generated_modules],
            [entry["module"] for entry in portable["modules"]],
        )
        self.assertEqual(sum(manifest["portable_stdlib"]["status_counts"].values()), len(portable["modules"]))
        self.assertTrue(all(entry["admission"] == "module-evidence-only" for entry in generated_modules))
        generated_by_module = {entry["module"]: entry for entry in generated_modules}
        for source_entry in portable["modules"]:
            generated = generated_by_module[source_entry["module"]]
            self.assertFalse(generated["release_admitted"])
            if source_entry["status"] == "snapshot":
                self.assertEqual(generated["state"], "compile-go-test-run-supported")

    def test_every_platform_preset_and_operation_has_an_explicit_closed_state(self) -> None:
        manifest = self.load_manifest()

        for entry in [*manifest["platforms"], *manifest["presets"]]:
            self.assertIn(entry["state"], ALLOWED_STATES, entry["id"])
            self.assertIsInstance(entry["release_admitted"], bool, entry["id"])

        seen_symbols: set[str] = set()
        for surface in manifest["surfaces"]:
            self.assertTrue(surface["operations"], surface["id"])
            for operation in surface["operations"]:
                label = f"{surface['id']}:{operation['id']}"
                self.assertIn(operation["state"], ALLOWED_STATES, label)
                self.assertIn(operation["granularity"], {"operation", "member", "surface"}, label)
                self.assertTrue(operation["symbols"], label)
                for symbol in operation["symbols"]:
                    self.assertTrue(symbol, label)
                    self.assertNotIn(symbol, seen_symbols, label)
                    seen_symbols.add(symbol)
                self.assertTrue(operation["evidence_ids"], label)
                self.assertIsInstance(operation["release_admitted"], bool, label)
                self.assertTrue(operation["qualification"], label)

    def test_high_risk_surfaces_are_bounded_at_operation_level(self) -> None:
        manifest = self.load_manifest()
        surfaces = {entry["id"]: entry for entry in manifest["surfaces"]}

        file_ops = {
            symbol: entry
            for entry in surfaces["portable-file-io"]["operations"]
            for symbol in entry["symbols"]
        }
        self.assertTrue(file_ops["sys.io.File.getContent"]["release_admitted"])
        self.assertTrue(file_ops["sys.io.File.saveContent"]["release_admitted"])

        process_ops = {
            symbol: entry
            for entry in surfaces["portable-process"]["operations"]
            for symbol in entry["symbols"]
        }
        self.assertTrue(process_ops["sys.io.Process.stdout.readLine"]["release_admitted"])
        self.assertTrue(process_ops["sys.io.Process.exitCode"]["release_admitted"])
        self.assertFalse(process_ops["new sys.io.Process(..., detached=true)"]["release_admitted"])

        concurrency_ops = {
            entry["id"]: entry for entry in surfaces["portable-concurrency"]["operations"]
        }
        self.assertTrue(concurrency_ops["primitives"]["release_admitted"])
        self.assertTrue(concurrency_ops["threads-event-loops-pools"]["release_admitted"])
        tls_lifecycle = concurrency_ops["tls-lifecycle"]
        self.assertTrue(tls_lifecycle["release_admitted"])
        self.assertEqual(tls_lifecycle["state"], "semantic-diff-supported")
        self.assertFalse(tls_lifecycle["blockers"])
        self.assertIn("runtime:thread-regressions", tls_lifecycle["evidence_ids"])
        self.assertIn("snapshot:detached-thread-lifecycle", tls_lifecycle["evidence_ids"])
        self.assertTrue(
            any("outside portable Thread" in exclusion for exclusion in tls_lifecycle["exclusions"])
        )

        network_ops = {
            entry["id"]: entry for entry in surfaces["portable-networking"]["operations"]
        }
        self.assertEqual(
            {
                "tcp-ipv4-blocking-client-core",
                "http-ipv4-blocking-client-core",
                "http-ipv4-multipart-upload",
                "http-data-url-client",
                "http-proxy-and-custom-transport",
                "https-client",
                "tcp-server-and-listener-controls",
                "socket-timeout-nonblocking-readiness-controls",
                "host-dns-and-reverse-lookup",
                "udp-ipv4",
                "tls-socket",
            },
            set(network_ops),
        )
        tcp_client = network_ops["tcp-ipv4-blocking-client-core"]
        self.assertTrue(tcp_client["release_admitted"])
        self.assertEqual("semantic-diff-supported", tcp_client["state"])
        self.assertEqual([], tcp_client["blockers"])
        self.assertIn("Linux/amd64", tcp_client["qualification"])
        self.assertIn("pre-resolved numeric endpoint", tcp_client["qualification"])

        http_core = network_ops["http-ipv4-blocking-client-core"]
        self.assertFalse(http_core["release_admitted"])
        self.assertEqual("semantic-diff-supported", http_core["state"])
        self.assertEqual(["haxe_go-vfp.10.8"], http_core["blockers"])
        self.assertIn("Linux/amd64", http_core["qualification"])
        self.assertIn("numeric IPv4", http_core["qualification"])
        self.assertIn("application-controlled", http_core["qualification"])
        self.assertIn("resource convergence", http_core["qualification"].lower())

        multipart = network_ops["http-ipv4-multipart-upload"]
        self.assertFalse(multipart["release_admitted"])
        self.assertEqual("semantic-diff-supported", multipart["state"])
        self.assertEqual(["haxe_go-vfp.10.8"], multipart["blockers"])
        self.assertTrue(
            any("blocks forever" in exclusion for exclusion in multipart["exclusions"])
        )

        data_url = network_ops["http-data-url-client"]
        self.assertFalse(data_url["release_admitted"])
        self.assertEqual("compile-go-test-run-supported", data_url["state"])
        self.assertEqual(["haxe_go-vfp.10.8"], data_url["blockers"])

        for operation_id in ("http-proxy-and-custom-transport", "https-client"):
            operation = network_ops[operation_id]
            self.assertFalse(operation["release_admitted"], operation_id)
            self.assertEqual("experimental", operation["state"])
            self.assertEqual([], operation["blockers"], operation_id)

        for operation_id in (
            "tcp-server-and-listener-controls",
            "socket-timeout-nonblocking-readiness-controls",
            "host-dns-and-reverse-lookup",
            "udp-ipv4",
            "tls-socket",
        ):
            operation = network_ops[operation_id]
            self.assertFalse(operation["release_admitted"], operation_id)
            self.assertEqual(["haxe_go-vfp.10.9"], operation["blockers"], operation_id)

        for surface_id, blocker in {"go-native": "haxe_go-vfp.9.1"}.items():
            operations = surfaces[surface_id]["operations"]
            self.assertFalse(any(entry["release_admitted"] for entry in operations), surface_id)
            self.assertIn(blocker, {item for entry in operations for item in entry["blockers"]}, surface_id)

        known_blockers = {entry["id"] for entry in manifest["known_blockers"]}
        self.assertNotIn("haxe_go-vfp.10.4", known_blockers)
        self.assertTrue({"haxe_go-vfp.10.8", "haxe_go-vfp.10.9"} <= known_blockers)

        trust_ops = {entry["id"]: entry for entry in surfaces["compiler-input-trust"]["operations"]}
        self.assertTrue(trust_ops["trusted-source"]["release_admitted"])
        self.assertFalse(trust_ops["untrusted-source"]["release_admitted"])
        self.assertEqual([], trust_ops["untrusted-source"]["blockers"])
        self.assertIn("contract:output-confinement", trust_ops["untrusted-source"]["evidence_ids"])
        self.assertIn("not a sandbox", trust_ops["untrusted-source"]["qualification"])

    def test_admitted_symbols_match_the_named_fixture_contracts(self) -> None:
        manifest = self.load_manifest()
        surfaces = {entry["id"]: entry for entry in manifest["surfaces"]}

        def admitted_symbols(surface_id: str) -> set[str]:
            return {
                symbol
                for operation in surfaces[surface_id]["operations"]
                if operation["release_admitted"]
                for symbol in operation["symbols"]
            }

        self.assertEqual(
            admitted_symbols("portable-collections-and-text") & {
                "StringTools.trim",
                "StringTools.startsWith",
                "StringTools.replace",
                "StringTools.contains",
                "StringTools.endsWith",
            },
            {
                "StringTools.trim",
                "StringTools.startsWith",
                "StringTools.replace",
                "StringTools.contains",
                "StringTools.endsWith",
            },
        )
        self.assertNotIn("StringTools.urlEncode", admitted_symbols("portable-collections-and-text"))
        self.assertNotIn("Array.iterator", admitted_symbols("portable-collections-and-text"))

        reflection = admitted_symbols("portable-reflection")
        self.assertTrue(
            {
                "Reflect.callMethod",
                "Reflect.compare",
                "Reflect.compareMethods",
                "Reflect.copy",
                "Reflect.deleteField",
                "Reflect.field",
                "Reflect.fields",
                "Reflect.getProperty",
                "Reflect.hasField",
                "Reflect.isEnumValue",
                "Reflect.isFunction",
                "Reflect.isObject",
                "Reflect.makeVarArgs",
                "Reflect.setField",
                "Reflect.setProperty",
            }
            <= reflection
        )

        filesystem = admitted_symbols("portable-filesystem")
        self.assertTrue(
            {
                "sys.FileSystem.exists",
                "sys.FileSystem.rename",
                "sys.FileSystem.stat",
                "sys.FileSystem.fullPath",
                "sys.FileSystem.absolutePath",
                "sys.FileSystem.createDirectory",
                "sys.FileSystem.isDirectory",
                "sys.FileSystem.readDirectory",
                "sys.FileSystem.deleteFile",
                "sys.FileSystem.deleteDirectory",
            }
            <= filesystem
        )

        self.assertEqual(
            {
                'new sys.net.Host("IPv4 literal")',
                "sys.net.Host.toString (IPv4 literal)",
                "sys.net.Socket.connect (blocking IPv4 literal)",
                "sys.net.Socket.input.readByte/readBytes (blocking)",
                "sys.net.Socket.output.writeByte/writeBytes (blocking)",
                "sys.net.Socket.output.flush",
                "sys.net.Socket.host",
                "sys.net.Socket.peer",
                "sys.net.Socket.close",
            },
            admitted_symbols("portable-networking"),
        )

    def test_portable_is_the_only_release_admitted_preset_semantics(self) -> None:
        presets = {entry["id"]: entry for entry in self.load_manifest()["presets"]}
        self.assertEqual(presets["portable"]["state"], "semantic-diff-supported")
        self.assertTrue(presets["portable"]["release_admitted"])
        self.assertEqual(presets["metal"]["state"], "compatibility-only")
        self.assertFalse(presets["metal"]["release_admitted"])
        self.assertIn("not a second semantic product", presets["metal"]["qualification"])

    def test_generated_docs_and_release_status_use_the_bounded_claim(self) -> None:
        matrix = MATRIX_DOC.read_text(encoding="utf-8")
        release = RELEASE_DOC.read_text(encoding="utf-8")

        for document in (matrix, release):
            self.assertIn("generated; edit compatibility-support-source.json", document)
            self.assertIn(RELEASE_CLAIM, document)
            self.assertNotIn("beta-stable", document.lower())
        self.assertIn("Operation/member admission", matrix)
        self.assertIn("Not admitted by this release scope", release)
        self.assertIn("application-controlled, pre-resolved numeric TCP endpoints", release)
        self.assertNotIn("haxe_go-vfp.6.3", release)
        self.assertIn("docs/public-contract.md", release)
        self.assertNotIn("haxe_go-vfp.6.4", release)
        self.assertIn("docs/semver-lifecycle-policy.md", release)

    def test_release_wiring_consumes_the_generated_manifest(self) -> None:
        package_json = (REPO_ROOT / "package.json").read_text(encoding="utf-8")
        release_status = (REPO_ROOT / "scripts" / "release" / "check-release-state.sh").read_text(encoding="utf-8")
        release_visibility = (REPO_ROOT / "docs" / "release-visibility.md").read_text(encoding="utf-8")
        checklist = (REPO_ROOT / "docs" / "release-readiness-checklist.md").read_text(encoding="utf-8")
        runner = (REPO_ROOT / "test" / "run-release-contracts.py").read_text(encoding="utf-8")

        self.assertIn('"compatibility:generate"', package_json)
        self.assertIn('"compatibility:verify"', package_json)
        self.assertIn("generate_support_manifest.py --check", release_status)
        self.assertIn('require_file "docs/compatibility-support-manifest.json"', release_status)
        self.assertIn("compatibility-support-manifest.json", release_visibility)
        self.assertIn("npm run compatibility:verify", checklist)
        self.assertIn("test/test_compatibility_support_manifest.py", runner)

    def test_invalid_or_implicit_source_entries_fail_closed(self) -> None:
        source = json.loads(SOURCE.read_text(encoding="utf-8"))
        mutations: list[tuple[str, dict[str, object], str]] = []

        invalid_state = copy.deepcopy(source)
        invalid_state["platforms"][0]["state"] = "probably-supported"
        mutations.append(("invalid-state", invalid_state, "unknown evidence state"))

        implicit_state = copy.deepcopy(source)
        del implicit_state["surfaces"][0]["operations"][0]["state"]
        mutations.append(("implicit-state", implicit_state, "must declare state"))

        implicit_members = copy.deepcopy(source)
        implicit_members["surfaces"][0]["operations"] = []
        mutations.append(("implicit-members", implicit_members, "must declare at least one operation/member"))

        implicit_symbols = copy.deepcopy(source)
        del implicit_symbols["surfaces"][0]["operations"][0]["symbols"]
        mutations.append(("implicit-symbols", implicit_symbols, "must declare symbols"))

        empty_symbols = copy.deepcopy(source)
        empty_symbols["surfaces"][0]["operations"][0]["symbols"] = []
        mutations.append(("empty-symbols", empty_symbols, "must declare symbols"))

        missing_evidence = copy.deepcopy(source)
        missing_evidence["evidence"][0]["path"] = "test/does-not-exist.compatibility-evidence"
        mutations.append(("missing-evidence", missing_evidence, "evidence path does not exist"))

        unknown_reference = copy.deepcopy(source)
        unknown_reference["surfaces"][0]["operations"][0]["evidence_ids"] = ["missing:evidence"]
        mutations.append(("unknown-reference", unknown_reference, "references unknown evidence"))

        unknown_field = copy.deepcopy(source)
        unknown_field["surfaces"][0]["operations"][0]["probably_supported"] = True
        mutations.append(("unknown-field", unknown_field, "has unknown fields"))

        weak_platform_evidence = copy.deepcopy(source)
        weak_platform_evidence["platforms"][0]["evidence_ids"] = ["bead:release-assets"]
        mutations.append(
            ("weak-platform-evidence", weak_platform_evidence, "lacks compile-go-test-run evidence")
        )

        weak_preset_evidence = copy.deepcopy(source)
        weak_preset_evidence["presets"][0]["evidence_ids"] = ["policy:native-presets"]
        mutations.append(("weak-preset-evidence", weak_preset_evidence, "lacks semantic-diff evidence"))

        with tempfile.TemporaryDirectory() as raw_temp:
            temp = Path(raw_temp)
            for name, mutated, expected in mutations:
                with self.subTest(name=name):
                    source_path = temp / f"{name}.json"
                    source_path.write_text(json.dumps(mutated, indent=2) + "\n", encoding="utf-8")
                    completed = run_generator(
                        "--source",
                        str(source_path),
                        "--manifest",
                        str(temp / f"{name}-manifest.json"),
                        "--matrix-doc",
                        str(temp / f"{name}-matrix.md"),
                        "--release-doc",
                        str(temp / f"{name}-release.md"),
                    )
                    self.assertNotEqual(completed.returncode, 0, completed.stdout)
                    self.assertIn(expected, completed.stdout + completed.stderr)


if __name__ == "__main__":
    unittest.main()
