package main

import "snapshot/hxrt"

func main() {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("startup.throws="), hxrt.StdString(throws(hxrt.NewFunction(func() {
		New_sys__io__Process(hxrt.StringFromLiteral("__haxe_go_missing_process__"), hxrt.NewArray(), false)
	})))))
	hxrt.Println(v)
	empty := New_sys__io__Process(hxrt.StringFromLiteral("sh"), hxrt.NewArray(hxrt.StringFromLiteral("-c"), hxrt.StringFromLiteral("printf '\\n'")), false)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("empty.line="), hxrt.StdString(hxrt.StringEqualStringPtr(empty.stdout.__hx_this.readLine(), hxrt.StringFromLiteral("")))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("empty.eof="), hxrt.StdString(reachesEof(empty))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("empty.code="), empty.__hx_this.exitCode(true)))
	hxrt.Println(v_3)
	empty.__hx_this.close()
	shell := New_sys__io__Process(hxrt.StringFromLiteral("printf 'shell-form\\n'"), nil, false)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("shell.line="), shell.stdout.__hx_this.readLine()))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("shell.code="), shell.__hx_this.exitCode(true)))
	hxrt.Println(v_5)
	shell.__hx_this.close()
	piped := New_sys__io__Process(hxrt.StringFromLiteral("sh"), hxrt.NewArray(hxrt.StringFromLiteral("-c"), hxrt.StringFromLiteral("IFS= read -r line; printf 'out:%s\\n' \"$line\"; printf 'err:%s\\n' \"$line\" >&2; exit 7")), false)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("pid.positive="), hxrt.StdString((piped.__hx_this.getPid() > 0))))
	hxrt.Println(v_6)
	piped.stdin.__hx_this.writeString(hxrt.StringFromLiteral("hello\n"), nil)
	piped.stdin.__hx_this.close()
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("stdin.stdout="), piped.stdout.__hx_this.readLine()))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("stderr.line="), piped.stderr.__hx_this.readLine()))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("exit.code="), piped.__hx_this.exitCode(true)))
	hxrt.Println(v_9)
	piped.__hx_this.close()
	longOutput := New_sys__io__Process(hxrt.StringFromLiteral("python3"), hxrt.NewArray(hxrt.StringFromLiteral("-c"), hxrt.StringFromLiteral("print('x' * 70000)")), false)
	var v_10 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("long.length="), hxrt.StringLengthStringPtr(longOutput.stdout.__hx_this.readLine())))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("long.code="), longOutput.__hx_this.exitCode(true)))
	hxrt.Println(v_11)
	longOutput.__hx_this.close()
	nonblocking := New_sys__io__Process(hxrt.StringFromLiteral("sh"), hxrt.NewArray(hxrt.StringFromLiteral("-c"), hxrt.StringFromLiteral("sleep 0.2; exit 9")), false)
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("nonblock.running="), hxrt.StdString((nonblocking.__hx_this.exitCode(false) == nil))))
	hxrt.Println(v_12)
	waitForChild(hxrt.StringFromLiteral("0.3"))
	var nonblockingCode any = nonblocking.__hx_this.exitCode(false)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("nonblock.code="), nonblockingCode)))
	nonblocking.__hx_this.close()
	killed := New_sys__io__Process(hxrt.StringFromLiteral("sh"), hxrt.NewArray(hxrt.StringFromLiteral("-c"), hxrt.StringFromLiteral("sleep 5")), false)
	killed.__hx_this.kill()
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("kill.nonzero="), hxrt.StdString((killed.__hx_this.exitCode(true) != 0))))
	hxrt.Println(v_13)
	killed.__hx_this.close()
	marker := hxrt.StringFromLiteral("tmp_process_close_marker.txt")
	if sys__FileSystem_exists(marker) {
		sys__FileSystem_deleteFile(marker)
	}
	closing := New_sys__io__Process(hxrt.StringFromLiteral("sh"), hxrt.NewArray(hxrt.StringFromLiteral("-c"), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sleep 0.2; printf done > "), marker)), false)
	closing.__hx_this.close()
	waitForChild(hxrt.StringFromLiteral("0.5"))
	var v_14 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("close.keeps.running="), hxrt.StdString(sys__FileSystem_exists(marker))))
	hxrt.Println(v_14)
	if sys__FileSystem_exists(marker) {
		sys__FileSystem_deleteFile(marker)
	}
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("detached.throws="), hxrt.StdString(throws(hxrt.NewFunction(func() {
		New_sys__io__Process(hxrt.StringFromLiteral("sh"), hxrt.NewArray(hxrt.StringFromLiteral("-c"), hxrt.StringFromLiteral("exit 0")), true)
	})))))
	hxrt.Println(v_15)
}

func reachesEof(process *sys__io__Process) bool {
	hx_try_return_1 := false
	var hx_try_value_2 bool
	hxrt.TryCatch(func() {
		process.stdout.__hx_this.readLine()
		hx_try_value_2 = false
		hx_try_return_1 = true
		return
	}, func(hx_caught_3 any) {
		hx_exception_4 := haxe__Exception_caught(hx_caught_3)
		hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_exception_4
		if func(hx_value any) bool {
			switch hx_carrier := hx_value.(type) {
			case *haxe__io__Eof:
				if hx_carrier == nil {
					return false
				}
				return true
			default:
				return false
			}
		}(hxrt.ObjectSource(hxrt.ObjectSource(hx_payload_5))) {
			var hx_tmp_1 *haxe__io__Eof = hx_recover_class_0(hx_payload_5)
			_ = hx_tmp_1
			hx_try_value_2 = true
			hx_try_return_1 = true
			return
		} else {
			var hx_tmp any = hx_payload_5
			_ = hx_tmp
			hx_try_value_2 = false
			hx_try_return_1 = true
			return
		}
	})
	if hx_try_return_1 {
		return hx_try_value_2
	}
	return false
}

func throws(action *hxrt.Function[func()]) bool {
	hx_try_return_6 := false
	var hx_try_value_7 bool
	hxrt.TryCatch(func() {
		func() {
			var hx_callable_11 *hxrt.Function[func()] = action
			hxrt.FunctionCall(hx_callable_11)()
		}()
		hx_try_value_7 = false
		hx_try_return_6 = true
		return
	}, func(hx_caught_8 any) {
		hx_exception_9 := haxe__Exception_caught(hx_caught_8)
		hx_payload_10 := haxe__Exception_catchValue(hx_caught_8)
		_ = hx_exception_9
		var hx_tmp any = hx_payload_10
		_ = hx_tmp
		hx_try_value_7 = true
		hx_try_return_6 = true
		return
	})
	if hx_try_return_6 {
		return hx_try_value_7
	}
	return false
}

func waitForChild(seconds *string) {
	waiter := New_sys__io__Process(hxrt.StringFromLiteral("sh"), hxrt.NewArray(hxrt.StringFromLiteral("-c"), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sleep "), seconds)), false)
	waiter.__hx_this.exitCode(true)
	waiter.__hx_this.close()
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
	case *haxe__Exception:
		return source
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}
