package main

import "snapshot/hxrt"

type Sys struct {
}

func Sys_executablePath() *string {
	return hxrt.StdString(hxrt.SysCurrentProgramPath())
}

func Sys_getChar(echo bool) int {
	value := hxrt.SysReadCharValue()
	if value < 0 {
		hxrt.Throw(haxe__Exception_thrown(New_haxe__io__Eof()))
	}
	if echo {
		New_sys__io__FileOutput(hxrt.SysStdout()).__hx_this.writeByte(value)
	}
	return value
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

func Sys_setCwd(s *string) {
	hxrt.SysChangeCwd(s)
}

func Sys_setTimeLocale(loc *string) bool {
	ignored := loc
	_ = ignored
	return false
}

func Sys_stderr() *haxe__io__Output {
	return func(hx_value *sys__io__FileOutput) *haxe__io__Output {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__Output
	}(New_sys__io__FileOutput(hxrt.SysStderr()))
}

func Sys_stdin() *haxe__io__Input {
	return func(hx_value *sys__io__FileInput) *haxe__io__Input {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__Input
	}(New_sys__io__FileInput(hxrt.SysStdin()))
}

func Sys_stdout() *haxe__io__Output {
	return func(hx_value *sys__io__FileOutput) *haxe__io__Output {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__Output
	}(New_sys__io__FileOutput(hxrt.SysStdout()))
}

func Sys_time() float64 {
	return hxrt.SysTime()
}
