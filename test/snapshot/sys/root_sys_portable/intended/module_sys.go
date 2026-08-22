package main

import "snapshot/hxrt"

type Sys struct {
}

func Sys_args() *hxrt.Array {
	return hxrt.ArrayFromValues(func(hx_sort_src_1 []*string) []any {
		hx_sort_out_3 := make([]any, 0, len(hx_sort_src_1))
		for _, hx_sort_item_2 := range hx_sort_src_1 {
			hx_sort_out_3 = append(hx_sort_out_3, hx_sort_item_2)
		}
		return hx_sort_out_3
	}(hxrt.SysArgs()))
}

func Sys_command(cmd *string, args *hxrt.Array) int {
	return hxrt.SysCommand(cmd, func() []*string {
		var hx_if_9 []*string
		if args == nil {
			hx_if_9 = nil
		} else {
			hx_if_9 = func(hx_lambda_raw_4 []any) []*string {
				hx_lambda_out_5 := make([]*string, 0, len(hx_lambda_raw_4))
				for _, hx_lambda_item_6 := range hx_lambda_raw_4 {
					hx_lambda_out_5 = append(hx_lambda_out_5, func(hx_value_7 any) *string {
						if hx_value_7 == nil {
							var hx_zero_8 *string
							return hx_zero_8
						}
						return hx_value_7.(*string)
					}(hx_lambda_item_6))
				}
				return hx_lambda_out_5
			}(args.Values())
		}
		return hx_if_9
	}())
}

func Sys_cpuTime() float64 {
	hxrt.Throw(hxrt.StringFromLiteral("Sys.cpuTime is unsupported on haxe.go: process CPU time is not implemented"))
	var hx_throw_zero_10 float64
	return hx_throw_zero_10
}

func Sys_environment() *haxe__ds__StringMap {
	environment := New_haxe__ds__StringMap()
	_g := 0
	_g1 := hxrt.SysEnvironmentEntries()
	for _g < len(_g1) {
		entry := _g1[_g]
		_g = int(int32((_g + 1)))
		environment.__hx_this.set(entry.Key, entry.Value)
	}
	return environment
}

func Sys_executablePath() *string {
	return hxrt.StdString(hxrt.SysCurrentProgramPath())
}

func Sys_exit(code int) {
	hxrt.SysExit(code)
}

func Sys_getChar(echo bool) int {
	value := hxrt.SysReadCharValue()
	if value < 0 {
		hxrt.Throw(New_haxe__io__Eof())
	}
	if echo {
		New_sys__io__FileOutput(hxrt.SysStdout()).__hx_this.writeByte(value)
	}
	return value
}

func Sys_getCwd() *string {
	return hxrt.StdString(hxrt.SysGetCwd())
}

func Sys_getEnv(s *string) *string {
	return hxrt.SysGetEnv(s)
}

func Sys_print(v any) {
	hxrt.Print(v)
}

func Sys_println(v any) {
	hxrt.Println(v)
}

func Sys_programPath() *string {
	return hxrt.StdString(hxrt.SysCurrentProgramPath())
}

func Sys_putEnv(s *string, v *string) {
	hxrt.SysSetEnvironment(s, v)
}

func Sys_setCwd(s *string) {
	hxrt.SysChangeCwd(s)
}

func Sys_setTimeLocale(loc *string) bool {
	ignored := loc
	_ = ignored
	return false
}

func Sys_sleep(seconds float64) {
	hxrt.SysSleep(seconds)
}

func Sys_stderr() *haxe__io__Output {
	return New_sys__io__FileOutput(hxrt.SysStderr()).haxe__io__Output
}

func Sys_stdin() *haxe__io__Input {
	return New_sys__io__FileInput(hxrt.SysStdin()).haxe__io__Input
}

func Sys_stdout() *haxe__io__Output {
	return New_sys__io__FileOutput(hxrt.SysStdout()).haxe__io__Output
}

func Sys_systemName() *string {
	return hxrt.StdString(hxrt.SysSystemName())
}

func Sys_time() float64 {
	return hxrt.SysTime()
}
