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

func hx_recover_callable_0(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__io__Output]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__io__Input]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_9 any) bool {
				if hx_value_9 == nil {
					var hx_zero_10 bool
					return hx_zero_10
				}
				return hx_value_9.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_11 any) bool {
				if hx_value_11 == nil {
					var hx_zero_12 bool
					return hx_zero_12
				}
				return hx_value_11.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__io__Output]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__io__Input]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() *haxe__io__Output]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__io__Input]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_19 any) *string {
				if hx_value_19 == nil {
					var hx_zero_20 *string
					return hx_zero_20
				}
				return hx_value_19.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_21 any) *string {
				if hx_value_21 == nil {
					var hx_zero_22 *string
					return hx_zero_22
				}
				return hx_value_21.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_23 any) *string {
				if hx_value_23 == nil {
					var hx_zero_24 *string
					return hx_zero_24
				}
				return hx_value_23.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *haxe__io__Eof {
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

func hx_recover_class_1(value any) *haxe__ValueException {
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

func hx_recover_class_2(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_3(value any) *haxe__ds__StringMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return source
	default:
		return value.(*haxe__ds__StringMap)
	}
}
