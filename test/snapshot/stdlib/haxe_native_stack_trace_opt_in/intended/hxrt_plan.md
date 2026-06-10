# Runtime Plan Report

- schema version: `1`
- contract: `portable`
- mode: `selective`
- selective enabled: `yes`
- full copy: `no`
- inference disabled: `no`

## manual features
- none

## inferred features
- `core`
- `exception`
- `print`
- `stack`
- `string`

## selected features
- `core`
- `exception`
- `print`
- `stack`
- `string`

## runtime files
- `core.go`
- `exception.go`
- `hxrt.go`
- `print.go`
- `stack.go`
- `string.go`

## selection reasons
- `core` <- `baseline` (`compiler_baseline`)
- `string` <- `baseline` (`compiler_baseline`)
- `print` <- `baseline` (`compiler_baseline`)
- `exception` <- `baseline` (`compiler_baseline`)
- `stack` <- `define` (`reflaxe_go_native_stack_trace`)
