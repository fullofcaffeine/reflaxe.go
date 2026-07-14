# Runtime Plan Report

- schema version: `2`
- contract: `portable`
- policy preset: `portable_default`
- semantic boundary source: `typed_api_or_module`
- mode: `selective`
- selective enabled: `yes`
- full copy: `no`
- inference disabled: `no`

## manual features
- `process`

## inferred features
- `core`
- `exception`
- `json`
- `print`
- `string`

## selected features
- `core`
- `exception`
- `json`
- `print`
- `process`
- `string`
- `sys`

## runtime files
- `core.go`
- `exception.go`
- `hxrt.go`
- `json.go`
- `print.go`
- `process.go`
- `string.go`
- `sys.go`

## selection reasons
- `core` <- `baseline` (`compiler_baseline`)
- `string` <- `baseline` (`compiler_baseline`)
- `print` <- `baseline` (`compiler_baseline`)
- `exception` <- `baseline` (`compiler_baseline`)
- `json` <- `class_usage` (`haxe.Json`)
- `sys` <- `dependency_edge` (`process->sys`)
- `process` <- `manual_define` (`reflaxe_go_hxrt_features`)
