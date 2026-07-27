package reflaxe.go.compiler;

/**
	Closed names for compiler defines owned or consumed by haxe.go.

	Why
	String literals at macro call sites are easy to mistype and difficult to audit.
	A closed type makes the supported configuration vocabulary discoverable and
	keeps each spelling in one place.

	What
	Each value is the exact name Haxe exposes through `Context.defined(...)`,
	`Context.definedValue(...)`, or `Compiler.define(...)`.

	How
	The abstract converts outward to `String` for the Haxe macro API, but it does
	not convert arbitrary strings inward. Dynamic define names supplied by generic
	helpers remain ordinary `String` values at those explicit boundaries.
**/
enum abstract GoCompilerDefine(String) to String {
	var DefineGoOutput = "go_output";
	var DefineGoTarget = "go";
	var DefineTargetName = "target.name";
	var DefineTargetAtomics = "target.atomics";
	var DefineGoNoBuild = "go_no_build";
	var DefineGoCodegenOnly = "go_codegen_only";
	var DefineGoCommand = "go_cmd";
	var DefineGoBuildOutput = "go_build_output";
	var DefineProfile = "reflaxe_go_profile";
	var DefinePortableProfile = "reflaxe_go_portable";
	var DefineIdiomaticProfile = "reflaxe_go_idiomatic";
	var DefineGopherProfile = "reflaxe_go_gopher";
	var DefineMetalProfile = "reflaxe_go_metal";
	var DefineRawNativeMode = "reflaxe_go_raw_native_mode";
	var DefineStrictExamples = "reflaxe_go_strict_examples";
	var DefineStrict = "reflaxe_go_strict";
	var DefineStrictPolicy = "reflaxe_go_strict_policy";
	var DefineMetalAllowFallback = "reflaxe_go_metal_allow_fallback";
	var DefineNativeAuthority = "reflaxe_go_native_authority";
	var DefineNativeSpecialization = "reflaxe_go_native_specialization";
	var DefineNativeFallback = "reflaxe_go_native_fallback";
	var DefineAutoMode = "reflaxe_go_auto";
	var DefineGoModule = "go_module";
	var DefineLineDirectives = "reflaxe_go_line_directives";
	var DefineHxrtDefaultFeatures = "reflaxe_go_hxrt_default_features";
	var DefineHxrtFeatures = "reflaxe_go_hxrt_features";
	var DefineHxrtNoFeatureInfer = "reflaxe_go_hxrt_no_feature_infer";
	var DefineContractReport = "reflaxe_go_contract_report";
	var DefineRuntimePlanReport = "reflaxe_go_runtime_plan_report";
	var DefineOptimizerPlanReport = "reflaxe_go_optimizer_plan_report";
	var DefineTypeUsageReport = "reflaxe_go_type_usage_report";
	var DefineSurfaceContractReport = "reflaxe_go_surface_contract_report";
	var DefineOptimizationPreset = "reflaxe_go_opt";
	var DefinePortableConcurrencyFastpath = "reflaxe_go_opt_go_concurrency_fastpath";
	var DefineNativeStackTrace = "reflaxe_go_native_stack_trace";
	var DefinePortableNativePolicy = "reflaxe_go_portable_native_policy";
	var DefinePortableNativeAllow = "reflaxe_go_portable_native_allow";
	var DefinePortableNativeScanMode = "reflaxe_go_portable_native_scan_mode";
	var DefineAutoEmptyConstructorInterfaces = "reflaxe_go_auto_empty_ctor_interfaces";
	var DefineGranularPassRegistry = "go_granular_pass_registry";
	var DefineLegacyPassBundle = "reflaxe_go_legacy_pass_bundle";
	var DefineTestPassRegistryCase = "reflaxe_go_test_registry_case";
	var DefineTestAstStatementCase = "reflaxe_go_test_ast_stmt_case";
}
