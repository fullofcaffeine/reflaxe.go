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
- none

## inferred features
- `array`
- `core`
- `exception`
- `print`
- `stack`
- `string`

## selected features
- `array`
- `core`
- `exception`
- `print`
- `stack`
- `string`

## runtime files
- `array.go`
- `core.go`
- `exception.go`
- `hxrt.go`
- `print.go`
- `stack.go`
- `string.go`

## selection reasons
- `core` <- `baseline` (`compiler_baseline`)
- `array` <- `generated_surface` (`hxrt.Array`)
- `string` <- `baseline` (`compiler_baseline`)
- `print` <- `baseline` (`compiler_baseline`)
- `exception` <- `baseline` (`compiler_baseline`)
- `stack` <- `class_usage` (`hxrt.stack.NativeStack`)
- `stack` <- `class_usage` (`hxrt.stack.NativeStackFrame`)
