package main

import "snapshot/hxrt"

func main() {
	printFn := hxrt.NamedFunction(Sys_print, "Sys.print")
	printlnFn := hxrt.NamedFunction(Sys_println, "Sys.println")
	localeFn := hxrt.NamedFunction(Sys_setTimeLocale, "Sys.setTimeLocale")
	setCwdFn := hxrt.NamedFunction(Sys_setCwd, "Sys.setCwd")
	timeFn := hxrt.NamedFunction(Sys_time, "Sys.time")
	executablePathFn := hxrt.NamedFunction(Sys_executablePath, "Sys.executablePath")
	programPathFn := hxrt.NamedFunction(Sys_programPath, "Sys.programPath")
	getCharFn := hxrt.NamedFunction(Sys_getChar, "Sys.getChar")
	stdinFn := hxrt.NamedFunction(Sys_stdin, "Sys.stdin")
	stdoutFn := hxrt.NamedFunction(Sys_stdout, "Sys.stdout")
	stderrFn := hxrt.NamedFunction(Sys_stderr, "Sys.stderr")
	func() {
		var hx_callable_1 *hxrt.Function[func(any)] = printFn
		var hx_argument_2 any = hxrt.StringFromLiteral("print=")
		hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
	func() {
		var hx_callable_3 *hxrt.Function[func(any)] = printlnFn
		var hx_argument_4 any = hxrt.StringFromLiteral("ok")
		hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}()
	func() {
		var hx_callable_7 *hxrt.Function[func(any)] = printlnFn
		var hx_argument_8 any = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("locale="), hxrt.StdString(!func() bool {
			var hx_callable_5 *hxrt.Function[func(*string) bool] = localeFn
			var hx_argument_6 *string = hxrt.StringFromLiteral("__haxe_go_missing_locale__")
			return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
		}()))
		hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}()
	cwd := hxrt.StdString(hxrt.SysGetCwd())
	func() {
		var hx_callable_9 *hxrt.Function[func(*string)] = setCwdFn
		var hx_argument_10 *string = cwd
		hxrt.FunctionCall(hx_callable_9)(hx_argument_10)
	}()
	func() {
		var hx_callable_11 *hxrt.Function[func(any)] = printlnFn
		var hx_argument_12 any = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("cwd="), hxrt.StdString(hxrt.StringEqualStringPtr(hxrt.StdString(hxrt.SysGetCwd()), cwd)))
		hxrt.FunctionCall(hx_callable_11)(hx_argument_12)
	}()
	started := func() float64 {
		var hx_callable_13 *hxrt.Function[func() float64] = timeFn
		return hxrt.FunctionCall(hx_callable_13)()
	}()
	hxrt.SysSleep(0.01)
	finished := func() float64 {
		var hx_callable_14 *hxrt.Function[func() float64] = timeFn
		return hxrt.FunctionCall(hx_callable_14)()
	}()
	func() {
		var hx_callable_15 *hxrt.Function[func(any)] = printlnFn
		var hx_argument_16 any = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("time="), hxrt.StdString(((started > 0) && (finished >= started))))
		hxrt.FunctionCall(hx_callable_15)(hx_argument_16)
	}()
	programPath := func() *string {
		var hx_callable_17 *hxrt.Function[func() *string] = programPathFn
		return hxrt.FunctionCall(hx_callable_17)()
	}()
	func() {
		var hx_callable_18 *hxrt.Function[func(any)] = printlnFn
		var hx_argument_19 any = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("programPath="), hxrt.StdString((hxrt.StringLengthStringPtr(programPath) > 0)))
		hxrt.FunctionCall(hx_callable_18)(hx_argument_19)
	}()
	func() {
		var hx_callable_21 *hxrt.Function[func(any)] = printlnFn
		var hx_argument_22 any = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("executableAlias="), hxrt.StdString(hxrt.StringEqualStringPtr(func() *string {
			var hx_callable_20 *hxrt.Function[func() *string] = executablePathFn
			return hxrt.FunctionCall(hx_callable_20)()
		}(), programPath)))
		hxrt.FunctionCall(hx_callable_21)(hx_argument_22)
	}()
	stdin := func() *haxe__io__Input {
		var hx_callable_23 *hxrt.Function[func() *haxe__io__Input] = stdinFn
		return hxrt.FunctionCall(hx_callable_23)()
	}()
	stderr := func() *haxe__io__Output {
		var hx_callable_24 *hxrt.Function[func() *haxe__io__Output] = stderrFn
		return hxrt.FunctionCall(hx_callable_24)()
	}()
	func() {
		var hx_callable_25 *hxrt.Function[func(any)] = printlnFn
		var hx_argument_26 any = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("inputFunctions="), hxrt.StdString((((stdin != nil) && (stderr != nil)) && (getCharFn != nil))))
		hxrt.FunctionCall(hx_callable_25)(hx_argument_26)
	}()
	stdout := func() *haxe__io__Output {
		var hx_callable_27 *hxrt.Function[func() *haxe__io__Output] = stdoutFn
		return hxrt.FunctionCall(hx_callable_27)()
	}()
	stdout.__hx_this.writeString(hxrt.StringFromLiteral("stdout=ok\n"), nil)
	stdout.__hx_this.flush()
	stdout.__hx_this.close()
	func() *haxe__io__Output {
		var hx_callable_28 *hxrt.Function[func() *haxe__io__Output] = stdoutFn
		return hxrt.FunctionCall(hx_callable_28)()
	}().__hx_this.writeString(hxrt.StringFromLiteral("stdoutAfterClose=ok\n"), nil)
}

var sysTypeProbe *Sys = nil

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_haxe__io__Eof(value any) *haxe__io__Eof {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Eof:
		return source
	default:
		return value.(*haxe__io__Eof)
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Eof(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}
