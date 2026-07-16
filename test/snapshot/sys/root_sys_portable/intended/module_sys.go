package main

import "snapshot/hxrt"

type Sys struct {
}

func Sys_args() []*string {
	return hxrt.SysArgs()
}

func Sys_command(cmd *string, args []*string) int {
	return hxrt.SysCommand(cmd, args)
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
		environment.set(entry.Key, entry.Value)
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
		var hx_throw_zero_11 int
		return hx_throw_zero_11
	}
	if echo {
		New_sys__io__FileOutput(hxrt.SysStdout()).writeByte(value)
	}
	return value
}

func Sys_getCwd() *string {
	return hxrt.StdString(hxrt.SysGetCwd())
}

func Sys_getEnv(s *string) *string {
	return hxrt.StdString(hxrt.SysGetEnv(s))
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

func Sys_stderr() haxe__io__Output {
	return New_sys__io__FileOutput(hxrt.SysStderr())
}

func Sys_stdin() haxe__io__Input {
	return New_sys__io__FileInput(hxrt.SysStdin())
}

func Sys_stdout() haxe__io__Output {
	return New_sys__io__FileOutput(hxrt.SysStdout())
}

func Sys_systemName() *string {
	return hxrt.StdString(hxrt.SysSystemName())
}

func Sys_time() float64 {
	return hxrt.SysTime()
}
