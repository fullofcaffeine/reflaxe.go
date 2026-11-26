package main

import "snapshot/hxrt"

func main() {
	root := hxrt.StringFromLiteral("tmp_file_error_semantics")
	if sys__FileSystem_exists(root) {
		sys__FileSystem_deleteDirectory(root)
	}
	sys__FileSystem_createDirectory(root)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("missing.read.throws="), hxrt.StdString(throws(hxrt.NewFunction(func() {
		sys__io__File_getContent(hxrt.StringConcatStringPtr(root, hxrt.StringFromLiteral("/missing.txt")))
	})))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("directory.read.throws="), hxrt.StdString(throws(hxrt.NewFunction(func() {
		sys__io__File_getContent(root)
	})))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("directory.write.throws="), hxrt.StdString(throws(hxrt.NewFunction(func() {
		sys__io__File_saveContent(root, hxrt.StringFromLiteral("not-a-file"))
	})))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("environment.invalid.throws="), hxrt.StdString(throws(hxrt.NewFunction(func() {
		hxrt.SysSetEnvironment(hxrt.StringFromLiteral("HAXE_GO=INVALID"), hxrt.StringFromLiteral("value"))
	})))))
	hxrt.Println(v_3)
	locked := hxrt.StringConcatStringPtr(root, hxrt.StringFromLiteral("/locked.txt"))
	sys__io__File_saveContent(locked, hxrt.StringFromLiteral("secret"))
	Sys_command(hxrt.StringFromLiteral("chmod"), hxrt.NewArray(hxrt.StringFromLiteral("000"), locked))
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("permission.read.throws="), hxrt.StdString(throws(hxrt.NewFunction(func() {
		sys__io__File_getContent(locked)
	})))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("permission.write.throws="), hxrt.StdString(throws(hxrt.NewFunction(func() {
		sys__io__File_saveContent(locked, hxrt.StringFromLiteral("replacement"))
	})))))
	hxrt.Println(v_5)
	Sys_command(hxrt.StringFromLiteral("chmod"), hxrt.NewArray(hxrt.StringFromLiteral("600"), locked))
	sys__FileSystem_deleteFile(locked)
	sys__FileSystem_deleteDirectory(root)
}

func throws(action *hxrt.Function[func()]) bool {
	hx_try_return_1 := false
	var hx_try_value_2 bool
	hxrt.TryCatch(func() {
		func() {
			var hx_callable_6 *hxrt.Function[func()] = action
			hxrt.FunctionCall(hx_callable_6)()
		}()
		hx_try_value_2 = false
		hx_try_return_1 = true
		return
	}, func(hx_caught_3 any) {
		hx_exception_4 := haxe__Exception_caught(hx_caught_3)
		hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_exception_4
		var hx_tmp any = hx_payload_5
		_ = hx_tmp
		hx_try_value_2 = true
		hx_try_return_1 = true
		return
	})
	if hx_try_return_1 {
		return hx_try_value_2
	}
	return false
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
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
	case *haxe__Exception:
		return (source != nil)
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	case *haxe__ValueException:
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
