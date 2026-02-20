## [0.26.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.25.0...v0.26.0) (2026-02-20)

### Features

* add FileSystem parity coverage and throw-expression lowering ([e8bae77](https://github.com/fullofcaffeine/reflaxe.go/commit/e8bae7719d8abab4cfce9395130f24e68018065c))

## [0.25.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.24.2...v0.25.0) (2026-02-20)

### Features

* **core:** lower TTypeExpr class/enum value expressions ([2ebaf6f](https://github.com/fullofcaffeine/reflaxe.go/commit/2ebaf6f0367f019aab36f51381fbc22373e6af05))

## [0.24.2](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.24.1...v0.24.2) (2026-02-20)

### Performance Improvements

* **harness:** add atomic microcase to go profile baseline ([ded32bc](https://github.com/fullofcaffeine/reflaxe.go/commit/ded32bced6b6a7a2a2e1e0bc80f2188604d8e218))

## [0.24.1](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.24.0...v0.24.1) (2026-02-20)

### Performance Improvements

* **atomic:** use sync/atomic for AtomicInt runtime ops ([bfaa9e5](https://github.com/fullofcaffeine/reflaxe.go/commit/bfaa9e567ccabb9b5c5789e3166babde46603102))

## [0.24.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.23.0...v0.24.0) (2026-02-20)

### Features

* **atomic:** promote AtomicInt/AtomicBool to snapshot tier ([8b076de](https://github.com/fullofcaffeine/reflaxe.go/commit/8b076ded74e5598cb847fe732bef3248d2537336))

## [0.23.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.22.1...v0.23.0) (2026-02-20)

### Features

* add AtomicObject runtime shims and snapshot parity ([163324c](https://github.com/fullofcaffeine/reflaxe.go/commit/163324c7fbef286c01db570557b4ac8408a181b3))

## [0.22.1](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.22.0...v0.22.1) (2026-02-20)

### Bug Fixes

* add Int32 parity fixture and numeric lowering fixes ([d217980](https://github.com/fullofcaffeine/reflaxe.go/commit/d21798052d9cc50151e4b26c008001b234ed5b54))

## [0.22.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.21.1...v0.22.0) (2026-02-20)

### Features

* add Int64 parity fixtures and lowering/runtime support ([fed4f74](https://github.com/fullofcaffeine/reflaxe.go/commit/fed4f74e9cfdff7149c820a9fcc3500873ea40b2))

## [0.21.1](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.21.0...v0.21.1) (2026-02-20)

### Bug Fixes

* **core:** avoid panic on missing anonymous object fields ([f88c4f6](https://github.com/fullofcaffeine/reflaxe.go/commit/f88c4f65087faa5900d74650cbca13613b3436ae))

## [0.21.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.20.0...v0.21.0) (2026-02-20)

### Features

* **stdlib:** add haxe.PosInfos runtime parity coverage ([5dbb994](https://github.com/fullofcaffeine/reflaxe.go/commit/5dbb99472e7569511bfdc0577aaee02171233926))

## [0.20.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.19.0...v0.20.0) (2026-02-19)

### Features

* **sys:** add host parity fixtures and resolve/reverse semantics ([4654c0f](https://github.com/fullofcaffeine/reflaxe.go/commit/4654c0ff3abe7a22481db1be66194d933412488f))

## [0.19.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.18.0...v0.19.0) (2026-02-19)

### Features

* **stdlib:** add vector runtime parity coverage ([fa2a6ae](https://github.com/fullofcaffeine/reflaxe.go/commit/fa2a6aee11a3a4451a797aa7fb13a141815a1b3d))

### Bug Fixes

* **core:** lower new Array() constructors to native slices ([1df97e1](https://github.com/fullofcaffeine/reflaxe.go/commit/1df97e1baf54c4a080dfa03f1fd7fe4ed7563657))

## [0.18.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.17.1...v0.18.0) (2026-02-19)

### Features

* **stdlib:** move Sys/File/Process behavior into hxrt runtime ([638a7ea](https://github.com/fullofcaffeine/reflaxe.go/commit/638a7eac4432b58b44118c86681e7818953fcd7d))

## [0.17.1](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.17.0...v0.17.1) (2026-02-19)

### Performance Improvements

* **stdlib:** cache bytes raw conversion path ([94d4adb](https://github.com/fullofcaffeine/reflaxe.go/commit/94d4adbe6f97656c9c71e1e29f6af2fcaca7ecad))

## [0.17.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.16.0...v0.17.0) (2026-02-19)

### Features

* **stdlib:** migrate json shim declarations out of compiler core ([8b18b3f](https://github.com/fullofcaffeine/reflaxe.go/commit/8b18b3f07d1f6ad116b6101035883dd53d2aa90f))

## [0.16.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.15.0...v0.16.0) (2026-02-19)

### Features

* **serializer:** support resolver polymorphism and ref stress ([8fb3100](https://github.com/fullofcaffeine/reflaxe.go/commit/8fb3100b328fb738c68b5fcf91d08b271cf29029))

## [0.15.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.14.0...v0.15.0) (2026-02-19)

### Features

* **socket:** add advanced sys.net.Socket shim semantics ([952a175](https://github.com/fullofcaffeine/reflaxe.go/commit/952a1750a4e6bd4267b9304b941d94502b7ca5cb))

## [0.14.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.13.0...v0.14.0) (2026-02-19)

### Features

* **serializer:** align enum cache refs with Haxe semantics ([ec0cf87](https://github.com/fullofcaffeine/reflaxe.go/commit/ec0cf8740f63ca6d982b6ffeddef68267aa65e26))

## [0.13.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.12.0...v0.13.0) (2026-02-19)

### Features

* **serializer:** add custom token and resolver materialization ([c312c2f](https://github.com/fullofcaffeine/reflaxe.go/commit/c312c2fe418b5252c22ecb87c4d2b0e8bfc3d1ea))

## [0.12.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.11.0...v0.12.0) (2026-02-19)

### Features

* **serializer:** extend wire-token parity for maps refs and enum index ([7dc81ff](https://github.com/fullofcaffeine/reflaxe.go/commit/7dc81ffeb312376cc3c34ca8ada6d5153668eb18))

## [0.11.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.10.0...v0.11.0) (2026-02-19)

### Features

* **serializer:** add class/enum wire-token parity ([b563f4a](https://github.com/fullofcaffeine/reflaxe.go/commit/b563f4a1002000bcf333ef49a2db793fed5f777a))

## [0.10.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.9.0...v0.10.0) (2026-02-19)

### Features

* **serializer:** add date/bytes wire-token coverage ([2f36c5e](https://github.com/fullofcaffeine/reflaxe.go/commit/2f36c5e2d2439480c11321a389f58393ee994a85))

## [0.9.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.8.0...v0.9.0) (2026-02-19)

### Features

* **serializer:** add wire-format token baseline ([94b2b77](https://github.com/fullofcaffeine/reflaxe.go/commit/94b2b776c805dd27acbb99d2b2f2a4b1134b23a7))

## [0.8.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.7.0...v0.8.0) (2026-02-19)

### Features

* **ereg:** harden option and match-state parity ([193d30f](https://github.com/fullofcaffeine/reflaxe.go/commit/193d30f8acd7a5b18fab0722721236da313719d9))

## [0.7.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.6.0...v0.7.0) (2026-02-19)

### Features

* **perf:** gate metal profile budgets in CI ([8e51e69](https://github.com/fullofcaffeine/reflaxe.go/commit/8e51e690d85558a2014a41d674c71119491a6d0e))

## [0.6.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.5.0...v0.6.0) (2026-02-19)

### Features

* **socket:** add deterministic sys.net loopback parity fixture ([c2f4f2e](https://github.com/fullofcaffeine/reflaxe.go/commit/c2f4f2eed053d7f69c38754c11640770188d9eb4))

## [0.5.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.4.0...v0.5.0) (2026-02-19)

### Features

* **semantic-diff:** add serializer/ereg parity fixtures and shim group ([1b1e0e4](https://github.com/fullofcaffeine/reflaxe.go/commit/1b1e0e47eecd492f1a45c8bfcf77e369aa159e90))

## [0.4.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.3.0...v0.4.0) (2026-02-19)

### Features

* **http:** expand sys.Http parity with customRequest and multipart ([de7bc21](https://github.com/fullofcaffeine/reflaxe.go/commit/de7bc218b4b85de9d326e6dd2a376bb1938ad7bc))

## [0.3.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.2.1...v0.3.0) (2026-02-19)

### Features

* **stdlib:** add functional sys.Http shim semantics ([71dfebd](https://github.com/fullofcaffeine/reflaxe.go/commit/71dfebd53f9c0de75133ad1cf3e92fa170559a4f))

## [0.2.1](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.2.0...v0.2.1) (2026-02-19)

### Bug Fixes

* **stdlib:** close haxe.Http go-test parity gap ([503a364](https://github.com/fullofcaffeine/reflaxe.go/commit/503a3645446c4cbba5e9f1135cb20f539b129a50))

## [0.2.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.1.0...v0.2.0) (2026-02-19)

### Features

* **phase0:** harden stdlib sweep and add perf/release visibility gates ([56b2706](https://github.com/fullofcaffeine/reflaxe.go/commit/56b27063c084643e7efc051a615ce1806d86e14c))
