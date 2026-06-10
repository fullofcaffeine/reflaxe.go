## [0.49.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.48.0...v0.49.0) (2026-06-10)

### Features

* close legacy text parity tranche ([1115976](https://github.com/fullofcaffeine/reflaxe.go/commit/1115976a88c2ae4fa4f3b8009e39bac512d932c4))
* move bytes helper algorithms into hxrt ([bd2ff22](https://github.com/fullofcaffeine/reflaxe.go/commit/bd2ff2292829d509fadc898f896e5ec35749afed))
* move io helper loops into staged std ([1d0bc89](https://github.com/fullofcaffeine/reflaxe.go/commit/1d0bc89fd0db3eba12a844e75dd0919737337d61))
* move sys http leaf helpers into staged std ([06984b0](https://github.com/fullofcaffeine/reflaxe.go/commit/06984b0d0cbfe8f19643adb690eed88b65e12bbc))
* promote haxe ds hashmap parity surface ([9ec0071](https://github.com/fullofcaffeine/reflaxe.go/commit/9ec00719659c035ef94d0ee67eab941d9c687411))
* promote haxe ds sort helper parity ([94f66fb](https://github.com/fullofcaffeine/reflaxe.go/commit/94f66fba233c3d20fa866dacaa5622e1d879ca1c))
* restore scoped raw go injection support ([16f794a](https://github.com/fullofcaffeine/reflaxe.go/commit/16f794a62c7880e80b4959c69031ad9e7d755c32))
* **rtti:** promote direct haxe.rtti parity ([5f24da3](https://github.com/fullofcaffeine/reflaxe.go/commit/5f24da388457b123a46766fd56d4202e08f5ca35))
* **rtti:** stabilize anonymous carrier mutation and class metadata lookup ([2c13cf2](https://github.com/fullofcaffeine/reflaxe.go/commit/2c13cf296d8f33f0248485e7a9b65c30a599856f))
* **std:** close sys.thread runtime parity tranche ([867624c](https://github.com/fullofcaffeine/reflaxe.go/commit/867624c4eb44b75e7511084ebef49beeb9dda099))
* **stdlib:** add opt-in native stack diagnostics ([8d3199a](https://github.com/fullofcaffeine/reflaxe.go/commit/8d3199a31f21127282b778f27da2eb130aa7317f))
* **stdlib:** close haxe.io misc direct parity tranche ([42240cd](https://github.com/fullofcaffeine/reflaxe.go/commit/42240cd84a8dfd6a1c95b5ecc654adef04d4ed08))
* **stdlib:** close sys db and io parity tranche ([0d9f947](https://github.com/fullofcaffeine/reflaxe.go/commit/0d9f947e67e1a4c0f5ef0dfccbcdb04ec31b59f1))
* **stdlib:** promote sys ssl leafs and net address ([20bf41c](https://github.com/fullofcaffeine/reflaxe.go/commit/20bf41c73432c20687467ec88ba3f787ea1050c3))
* **stdlib:** promote sys.net.UdpSocket direct baseline ([88f4f1c](https://github.com/fullofcaffeine/reflaxe.go/commit/88f4f1c2a2d2ce1f00b01598f9e5093cd7b2161a))
* **stdlib:** promote WeakMap platform contract on Go ([8cc8bff](https://github.com/fullofcaffeine/reflaxe.go/commit/8cc8bff90fa2f4eec65b4a9dcf89e533ed906217))
* **stdlib:** support direct haxe event loop surfaces ([582463b](https://github.com/fullofcaffeine/reflaxe.go/commit/582463b97e584a07c35eb842f29cfdf08f609561))
* **stdlib:** support direct haxe.http.HttpBase baseline ([be458c9](https://github.com/fullofcaffeine/reflaxe.go/commit/be458c977fb51ef94776688adb6ac4eeafbbe7ae))
* **std:** promote first-wave sys.thread primitives ([796b3d7](https://github.com/fullofcaffeine/reflaxe.go/commit/796b3d75529845840114670777636798a99a3961))
* **std:** promote haxe io typed arrays parity ([c724c74](https://github.com/fullofcaffeine/reflaxe.go/commit/c724c740b7c90d5057ea7a86bc30853e52605981))
* **std:** promote sys.ssl.Socket baseline support ([05dad56](https://github.com/fullofcaffeine/reflaxe.go/commit/05dad56be481c9eda5e3ec6efb4b40376f6c8039))
* wire embedded haxe.Resource payloads ([8404bb6](https://github.com/fullofcaffeine/reflaxe.go/commit/8404bb694ac93bae5119a01163c21ef9ffab0591))

### Bug Fixes

* **ci:** benchmark go profiles with selective hxrt ([7307c7c](https://github.com/fullofcaffeine/reflaxe.go/commit/7307c7cb9b28faf4a1f7710c7b768584a9595f1f))
* **compiler:** elide guarded nullable primitive casts ([c110fa0](https://github.com/fullofcaffeine/reflaxe.go/commit/c110fa01e23363075de201f21df4dc328850eed2))
* **compiler:** narrow nullable primitive locals ([ec2ba2d](https://github.com/fullofcaffeine/reflaxe.go/commit/ec2ba2dcfc601d8ec11b434982a44938805563a9))
* **compiler:** preserve nullable enum primitive payloads ([b514e9f](https://github.com/fullofcaffeine/reflaxe.go/commit/b514e9fac3dad5b0dafc4bea84c5869e7bb27fc7))
* **compiler:** support nullable primitive params ([689979f](https://github.com/fullofcaffeine/reflaxe.go/commit/689979f62cb6c2981ffa41e74b9121ac1071f7b0))
* **event-loop:** classify direct haxe loop surfaces as unsupported ([35d93fa](https://github.com/fullofcaffeine/reflaxe.go/commit/35d93fa622be05425fd9f746ff4b7625ff68bbb0))
* **go:** restore direct haxe.ValueException parity ([7648462](https://github.com/fullofcaffeine/reflaxe.go/commit/7648462dbf22a92e1e82285b30d8359f399fe9b1))
* **go:** restore safe instance default-arg padding ([34178a7](https://github.com/fullofcaffeine/reflaxe.go/commit/34178a7e843a4d8abe2ac293a2d46d95a7e6c367))
* **parity:** restore direct std exceptions and collections ([ed495f1](https://github.com/fullofcaffeine/reflaxe.go/commit/ed495f1e8675a87e9708287ecb32b591b938d167))
* **runtime:** keep native stack capture footprint explicit ([554e8a3](https://github.com/fullofcaffeine/reflaxe.go/commit/554e8a34e143961c988c3e704fa5e6c5c3aad5af))
* **runtime:** remove sys thread timeout dynamic bridge ([b275f7f](https://github.com/fullofcaffeine/reflaxe.go/commit/b275f7f297e0bea3c595db42a522c833a13b673b))
* **stdlib:** restore haxe Utf8 size constructor ([059f054](https://github.com/fullofcaffeine/reflaxe.go/commit/059f0540d077b23ad736d41b8fe293d81505f8dc))
* **stdlib:** use upstream haxe io path ([9f97fd6](https://github.com/fullofcaffeine/reflaxe.go/commit/9f97fd6cb970442cb72f7c08646c56113ef8d8c9))

## [0.48.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.47.1...v0.48.0) (2026-03-07)

### Features

* add hxrt runtime plan reason provenance ([bb2b0c7](https://github.com/fullofcaffeine/reflaxe.go/commit/bb2b0c78c2f49fda96b74e058a94f9455647650a))
* define stack and main-loop std snapshot scope ([8bf3d22](https://github.com/fullofcaffeine/reflaxe.go/commit/8bf3d228b72fa911af76294d404799509a3ec94c))
* move DateTools into staged std ([ff597ef](https://github.com/fullofcaffeine/reflaxe.go/commit/ff597ef932e91e9804cf16f957c3ec26ca8f56b5))
* move haxe.io.Path into staged std ([9e2f737](https://github.com/fullofcaffeine/reflaxe.go/commit/9e2f737fedfcdf59e859b8834e51949fcebd5bc7))
* move StringTools into staged std ([6acb400](https://github.com/fullofcaffeine/reflaxe.go/commit/6acb400b816da7f7dc02a6cfa4bb1d727969349d))
* promote direct haxe helper surfaces ([73a7e4c](https://github.com/fullofcaffeine/reflaxe.go/commit/73a7e4cdfc34e1c7eee0db529a3cd648202934d2))
* promote enum helper parity ([2c8056d](https://github.com/fullofcaffeine/reflaxe.go/commit/2c8056d6c1976a980a3fbf8bfdf8a2c029018e9e))
* promote haxe constraints and rest parity ([34350bf](https://github.com/fullofcaffeine/reflaxe.go/commit/34350bf52bb8e332c98d7a34bd6b4c097ecc465f))
* promote iterator family semantic-diff parity ([bc3c9ae](https://github.com/fullofcaffeine/reflaxe.go/commit/bc3c9ae4dd6ccd0ee73a94755b06b9f205308e17))
* promote root sys parity surface ([f0d5783](https://github.com/fullofcaffeine/reflaxe.go/commit/f0d578362dacd17018622d366c2ae284feb484c1))
* promote root xml parity surface ([65b35b4](https://github.com/fullofcaffeine/reflaxe.go/commit/65b35b4e892d738769b16dacf9005ad25b1dd1de))
* promote UnicodeString parity surface ([5e8430f](https://github.com/fullofcaffeine/reflaxe.go/commit/5e8430fb8ae97e9016b82924bb56edae066c3c90))

### Bug Fixes

* preserve parsed xml cdata nodes ([7c42cd6](https://github.com/fullofcaffeine/reflaxe.go/commit/7c42cd64a175c69c33802dba134aa748b3b38bed))
* realign ci stdlib governance and examples ([5cdb061](https://github.com/fullofcaffeine/reflaxe.go/commit/5cdb06125ea9ad7105fb90684da2ed950de266ee))
* restore ci snapshot parity gates ([e543acb](https://github.com/fullofcaffeine/reflaxe.go/commit/e543acb5782861bfb0f03db8cf2fbae7cb9b493d))

## [0.47.1](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.47.0...v0.47.1) (2026-03-04)

### Performance Improvements

* make go profile harness tolerate portable-only tui lane ([43364d5](https://github.com/fullofcaffeine/reflaxe.go/commit/43364d50c40392382f9ec508fcaafcfea4d20c1a))

## [0.47.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.46.0...v0.47.0) (2026-02-26)

### Features

* add lambda exists/has iterable fallback ([e1c3cf7](https://github.com/fullofcaffeine/reflaxe.go/commit/e1c3cf75fc2aeec504821756f7a13da2c5faeb15))
* add opt-in Go line directives for Haxe sources ([451f3e6](https://github.com/fullofcaffeine/reflaxe.go/commit/451f3e66814584c4a76f1d0622a11a037fa134de))
* add selective hxrt runtime slicing and contracts ([0f301a6](https://github.com/fullofcaffeine/reflaxe.go/commit/0f301a6253a658564c238506adb49f936d04c89b))
* **codegen:** elide redundant blank-identifier guards and refresh snapshots ([d496e4e](https://github.com/fullofcaffeine/reflaxe.go/commit/d496e4e942ecbeb799223bdb37f45602a33629c2))
* **devex:** add flag-driven auto empty constructor injection ([8220ede](https://github.com/fullofcaffeine/reflaxe.go/commit/8220ede1b1ca0f4108a0db06c9bbcdfb0b89d1bb))
* emit single-package file-per-module Go output ([16a4176](https://github.com/fullofcaffeine/reflaxe.go/commit/16a4176614710c6750984f49f950146cbcc4cf63))
* **examples:** add fluxproxy matrix and pure-go parity baselines ([cd5d10c](https://github.com/fullofcaffeine/reflaxe.go/commit/cd5d10c98b87752a35fd30d2293e3168e7a2c653))
* **examples:** add pulseforge scripted and interactive cli modes ([8629b30](https://github.com/fullofcaffeine/reflaxe.go/commit/8629b3041d2abd3fd6ad42e12b3979acb4b3cbbb))
* **examples:** implement pulseforge pipeline contract across profiles ([8f48495](https://github.com/fullofcaffeine/reflaxe.go/commit/8f484957dbab32da0650c1e84f68934f2dee3d42))
* **family:** add std sync verify and dual-run pin checks ([51ce2cd](https://github.com/fullofcaffeine/reflaxe.go/commit/51ce2cd34677658b1c1ed4b6fbe011f0df9014c9))
* **go-ast:** emit no-binding type switch when binding unused ([b7c0bd8](https://github.com/fullofcaffeine/reflaxe.go/commit/b7c0bd8cd3a7cfb0b31cd3b39d5e3a44ccda517d))
* **go:** add non-blocking Chan.tryRecv API ([95eff86](https://github.com/fullofcaffeine/reflaxe.go/commit/95eff86d0ce55957983d5c501eeb776fc9f79e8b))
* **go:** add typed select helper API with deterministic branching ([77c252b](https://github.com/fullofcaffeine/reflaxe.go/commit/77c252b1b45e6a485bcd42b3a3f1a3ebb7682fda))
* **goextern:** add deterministic generator and fixture CI checks ([9e6e1a2](https://github.com/fullofcaffeine/reflaxe.go/commit/9e6e1a2170c5f8d85dea120b3c1de1f35e9b402e))
* **interp+snapshots:** add std/go interop wrappers and realign fixtures ([c3bb620](https://github.com/fullofcaffeine/reflaxe.go/commit/c3bb62094a5c676971457f6b04aa91265e508cc7))
* **lambda:** add haxe.ds.List lowering and iterable diagnostics ([1b53122](https://github.com/fullofcaffeine/reflaxe.go/commit/1b53122eb3c264a4b21f77347e7d282bc30ab982))
* lower untyped identifier expressions ([b6b6f6d](https://github.com/fullofcaffeine/reflaxe.go/commit/b6b6f6d16853c88186397b0774f8aa017fbbb410))
* **metal:** add explicit __go__ fallback define ([acb86b8](https://github.com/fullofcaffeine/reflaxe.go/commit/acb86b8932f0cc0af9324ba49d321dfc418f8187))
* **metal:** add go.Chan typed monomorphization prototype ([737ceb2](https://github.com/fullofcaffeine/reflaxe.go/commit/737ceb245ecbd02abe2c2404b7e10fdceebcf3e0))
* **metal:** add typed go.Result lowering shims ([ef0def6](https://github.com/fullofcaffeine/reflaxe.go/commit/ef0def643cdad631ed22c9bfe05d58f1c5dc06c1))
* **metal:** extend monomorphization to go slice/map ([81bcf68](https://github.com/fullofcaffeine/reflaxe.go/commit/81bcf68f3315b25c21eaab5d6ebdd91db5eb4179))
* **perf:** add channel map generic go profile benches ([de21aa0](https://github.com/fullofcaffeine/reflaxe.go/commit/de21aa07148367ed5d714f7dd75c743021d990c5))
* **perf:** add flagship app benchmark baselines and CI stage ([904ab3c](https://github.com/fullofcaffeine/reflaxe.go/commit/904ab3cfa038f0922217861646fda2bf39cf708f))
* remove gopher profile and migrate portable baseline ([a7208e6](https://github.com/fullofcaffeine/reflaxe.go/commit/a7208e6e4c5a997540fce528fd1bf71e95f550c5))
* support lambda count/empty generic iterables ([e5c8fff](https://github.com/fullofcaffeine/reflaxe.go/commit/e5c8fff153b8f163b14aa3bdb11b810b93b7dbdd))

### Bug Fixes

* coerce nullable float/bool any-branches ([ee5f3c5](https://github.com/fullofcaffeine/reflaxe.go/commit/ee5f3c5c894ec41d29a164216ad1e8d8e8423a45))
* **go:** correct null equality lowering for Result semantics ([82e63d3](https://github.com/fullofcaffeine/reflaxe.go/commit/82e63d3bc74a3cd6836d102f0c2213c8a3a9fff7))
* **goextern:** stabilize empty extern formatting ([1ef9131](https://github.com/fullofcaffeine/reflaxe.go/commit/1ef913148cb2fed62932840cd256752602a2b200))
* normalize extern string returns to hxrt pointers ([b158c52](https://github.com/fullofcaffeine/reflaxe.go/commit/b158c520534a0d966bb422f5079b82a06420f57d))
* type-assert chan recv results in non-metal profiles ([7e8fe08](https://github.com/fullofcaffeine/reflaxe.go/commit/7e8fe088e0cc64357042f7d13d7d66bc6ea5c191))

### Performance Improvements

* add select microcase to go profile harness ([0b53260](https://github.com/fullofcaffeine/reflaxe.go/commit/0b53260d23dcb83ec4bd232a36dc38ccbac4c13b))
* add selective hxrt perf-size harness gating ([42b90ba](https://github.com/fullofcaffeine/reflaxe.go/commit/42b90ba23a9c8a6b77599b7e4d0817369b00cb38))
* add virtual and string go profile microcases ([ecf2597](https://github.com/fullofcaffeine/reflaxe.go/commit/ecf259776af3449b9327086c2c9a99325aa80a51))
* **apps:** add portable-vs-metal delta reporting and gating ([72dfc28](https://github.com/fullofcaffeine/reflaxe.go/commit/72dfc287b7b5c7af279a83120a2212028ec93bae))
* **go:** add portable-vs-metal delta budgets and docs ([88df1ca](https://github.com/fullofcaffeine/reflaxe.go/commit/88df1ca509142ff5b06e0d9153731d7d7ca033f2))
* **go:** add string_instance microcase to profile harness ([252b4ac](https://github.com/fullofcaffeine/reflaxe.go/commit/252b4ac8226e476c376e279a5bac3ce6f1683a1b))
* **go:** specialize go.Select helpers with @:generic ([6a52bbf](https://github.com/fullofcaffeine/reflaxe.go/commit/6a52bbfdfa1bd014b44c63e2d9f22d82eda96b65))

## [0.46.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.45.0...v0.46.0) (2026-02-23)

### Features

* **examples:** add typed interop smoke app across profiles ([080e06d](https://github.com/fullofcaffeine/reflaxe.go/commit/080e06daa55b2c65ed944633b4135cad5704552a))

## [0.45.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.44.0...v0.45.0) (2026-02-23)

### Features

* **interop:** add extern metadata lowering and import mapping ([2951fee](https://github.com/fullofcaffeine/reflaxe.go/commit/2951feea515e68a28e71cdffdbcf351db62970ae))

## [0.44.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.43.0...v0.44.0) (2026-02-23)

### Features

* **go-native:** add deterministic channel/select handshake contracts ([da6ac11](https://github.com/fullofcaffeine/reflaxe.go/commit/da6ac114053521b755330f5dbea8e7e7bba25fdc))

## [0.43.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.42.0...v0.43.0) (2026-02-23)

### Features

* **go-native:** enable real channel and goroutine primitives ([397e40a](https://github.com/fullofcaffeine/reflaxe.go/commit/397e40ad2f3aebafb557176c9858a5ccfaf626ae))

## [0.42.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.41.0...v0.42.0) (2026-02-23)

### Features

* **ast:** add range loop form ([90b9867](https://github.com/fullofcaffeine/reflaxe.go/commit/90b986751d0565b3cd07283a80b3313d4ff77b20))

## [0.41.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.40.0...v0.41.0) (2026-02-23)

### Features

* **ast:** add select statement form ([09ca5c4](https://github.com/fullofcaffeine/reflaxe.go/commit/09ca5c484e2ccc4ff2518d228a733f95b7405f64))

## [0.40.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.39.0...v0.40.0) (2026-02-23)

### Features

* **ast:** add channel send/recv forms ([22eceff](https://github.com/fullofcaffeine/reflaxe.go/commit/22eceffb1bec89d3bbd8485aa79c0f6f4d4b8405))

## [0.39.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.38.1...v0.39.0) (2026-02-23)

### Features

* **ast:** add go/defer statement nodes ([77d3083](https://github.com/fullofcaffeine/reflaxe.go/commit/77d30831c4c108d9fefdfd42d911e58e32b441c2))

## [0.38.1](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.38.0...v0.38.1) (2026-02-23)

### Bug Fixes

* **trycatch:** forward statement try/catch returns and seed phase2 roadmap ([59eecfd](https://github.com/fullofcaffeine/reflaxe.go/commit/59eecfda808f46e306fbd32289f0dd4d915ef74a))

## [0.38.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.37.0...v0.38.0) (2026-02-20)

### Features

* **stdlib:** add ReadOnlyArray semantic-diff parity ([2ab312c](https://github.com/fullofcaffeine/reflaxe.go/commit/2ab312c7b9043fb8a2ff15964035ba9e08b47806))

## [0.37.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.36.0...v0.37.0) (2026-02-20)

### Features

* **io:** add bytes hex shim parity ([42bd1c0](https://github.com/fullofcaffeine/reflaxe.go/commit/42bd1c0c4368f5301c1437a1becd9c712d8a2f6d))

## [0.36.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.35.0...v0.36.0) (2026-02-20)

### Features

* **io:** add Bytes.ofData shim and alias parity contract ([d368007](https://github.com/fullofcaffeine/reflaxe.go/commit/d3680071995a56e4abe7aa48dfce61ca6819f2a6))

## [0.35.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.34.0...v0.35.0) (2026-02-20)

### Features

* **io:** add configurable RawNative mode with snapshot guardrails ([8c763c3](https://github.com/fullofcaffeine/reflaxe.go/commit/8c763c3ac9c99dd52f80b761261eb0a57103c0cf))

## [0.34.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.33.0...v0.34.0) (2026-02-20)

### Features

* **io:** add Encoding enum constructors and Bytes.getString parity ([d481791](https://github.com/fullofcaffeine/reflaxe.go/commit/d481791fb140884f738badf505d2dcb5bbd515dc))

## [0.33.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.32.0...v0.33.0) (2026-02-20)

### Features

* **io:** throw haxe.io.Error constructors for blocked/outside paths ([e13b13e](https://github.com/fullofcaffeine/reflaxe.go/commit/e13b13e525f656b18e07a77db64725ee26267e00))

## [0.32.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.31.2...v0.32.0) (2026-02-20)

### Features

* **io:** model haxe.io.Error constructors in shim parity ([a7c71bc](https://github.com/fullofcaffeine/reflaxe.go/commit/a7c71bc9c7414983b1c6b7b409ac886ad0b31b3c))

## [0.31.2](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.31.1...v0.31.2) (2026-02-20)

### Bug Fixes

* **throw:** emit non-void fallback return after statement throws ([ceab5ec](https://github.com/fullofcaffeine/reflaxe.go/commit/ceab5ec51c2719b5ff0b571b74eb8d84b78df6b2))

## [0.31.1](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.31.0...v0.31.1) (2026-02-20)

### Bug Fixes

* **io:** support custom Input/Output subclasses ([b814d18](https://github.com/fullofcaffeine/reflaxe.go/commit/b814d186e220ec960696694a8ceac91e09556b35))

## [0.31.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.30.0...v0.31.0) (2026-02-20)

### Features

* **io:** gate helper shim emission by usage ([d22adf7](https://github.com/fullofcaffeine/reflaxe.go/commit/d22adf764fae72ca650ce75c9f249eac82a448c7))

## [0.30.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.29.0...v0.30.0) (2026-02-20)

### Features

* **io:** expand Input/Output helper subset parity ([5b5a49f](https://github.com/fullofcaffeine/reflaxe.go/commit/5b5a49f101a3629dbcae992c74df828371f1f3c1))

## [0.29.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.28.0...v0.29.0) (2026-02-20)

### Features

* add bytes input/output semantic parity subset ([096d0c9](https://github.com/fullofcaffeine/reflaxe.go/commit/096d0c9430304265ecc8aeccb9fc9389aa5b9c3f))

## [0.28.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.27.4...v0.28.0) (2026-02-20)

### Features

* add haxe.io.Bytes core ops shim parity ([d009513](https://github.com/fullofcaffeine/reflaxe.go/commit/d0095131587591dcca7e1efcca5148685f49d387))

## [0.27.4](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.27.3...v0.27.4) (2026-02-20)

### Bug Fixes

* normalize haxe.io.BytesBuffer byte pushes ([cdbf7ad](https://github.com/fullofcaffeine/reflaxe.go/commit/cdbf7ad064a3d148045751bbb7eb6d27f362776d))

## [0.27.3](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.27.2...v0.27.3) (2026-02-20)

### Bug Fixes

* wire customRequest socket transport injection ([6b7f32f](https://github.com/fullofcaffeine/reflaxe.go/commit/6b7f32f0dc23f080cfb338755e0ce79b239da033))

## [0.27.2](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.27.1...v0.27.2) (2026-02-20)

### Bug Fixes

* support resolver type-value class and enum returns ([446e055](https://github.com/fullofcaffeine/reflaxe.go/commit/446e055fc4ccc4116c5a4deebdd29eab94f13f25))

## [0.27.1](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.27.0...v0.27.1) (2026-02-20)

### Bug Fixes

* make haxe.ds typed reads nil-safe ([0254ac6](https://github.com/fullofcaffeine/reflaxe.go/commit/0254ac6628e3f35ed462066028c5545c486cced4))

## [0.27.0](https://github.com/fullofcaffeine/reflaxe.go/compare/v0.26.0...v0.27.0) (2026-02-20)

### Features

* align haxe.ds List.push parity and add ds semantic contract ([b0c392b](https://github.com/fullofcaffeine/reflaxe.go/commit/b0c392be1457bfa8799a53c42a9d032b0dae2279))

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
