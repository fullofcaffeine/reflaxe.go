package main

import "snapshot/hxrt"

func main() {
	printFn := Sys_print
	printlnFn := Sys_println
	localeFn := Sys_setTimeLocale
	setCwdFn := Sys_setCwd
	timeFn := Sys_time
	executablePathFn := Sys_executablePath
	programPathFn := Sys_programPath
	getCharFn := Sys_getChar
	stdinFn := Sys_stdin
	stdoutFn := Sys_stdout
	stderrFn := Sys_stderr
	printFn(hxrt.StringFromLiteral("print="))
	printlnFn(hxrt.StringFromLiteral("ok"))
	printlnFn(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("locale="), hxrt.StdString(!localeFn(hxrt.StringFromLiteral("__haxe_go_missing_locale__")))))
	cwd := hxrt.StdString(hxrt.SysGetCwd())
	setCwdFn(cwd)
	printlnFn(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("cwd="), hxrt.StdString(hxrt.StringEqualStringPtr(hxrt.StdString(hxrt.SysGetCwd()), cwd))))
	started := timeFn()
	hxrt.SysSleep(0.01)
	finished := timeFn()
	printlnFn(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("time="), hxrt.StdString(((started > 0) && (finished >= started)))))
	programPath := programPathFn()
	printlnFn(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("programPath="), hxrt.StdString((hxrt.StringLengthStringPtr(programPath) > 0))))
	printlnFn(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("executableAlias="), hxrt.StdString(hxrt.StringEqualStringPtr(executablePathFn(), programPath))))
	stdin := stdinFn()
	stderr := stderrFn()
	printlnFn(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("inputFunctions="), hxrt.StdString((((stdin != nil) && (stderr != nil)) && (getCharFn != nil)))))
	stdout := stdoutFn()
	stdout.__hx_this.writeString(hxrt.StringFromLiteral("stdout=ok\n"), nil)
	stdout.__hx_this.flush()
	stdout.__hx_this.close()
	stdoutFn().__hx_this.writeString(hxrt.StringFromLiteral("stdoutAfterClose=ok\n"), nil)
}

var sysTypeProbe *Sys = nil
