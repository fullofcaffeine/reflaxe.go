package main

import "snapshot/hxrt"

func main() {
	root := hxrt.StringFromLiteral("tmp_file_error_semantics")
	if sys__FileSystem_exists(root) {
		sys__FileSystem_deleteDirectory(root)
	}
	sys__FileSystem_createDirectory(root)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("missing.read.throws="), hxrt.StdString(throws(func() {
		sys__io__File_getContent(hxrt.StringConcatStringPtr(root, hxrt.StringFromLiteral("/missing.txt")))
	}))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("directory.read.throws="), hxrt.StdString(throws(func() {
		sys__io__File_getContent(root)
	}))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("directory.write.throws="), hxrt.StdString(throws(func() {
		sys__io__File_saveContent(root, hxrt.StringFromLiteral("not-a-file"))
	}))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("environment.invalid.throws="), hxrt.StdString(throws(func() {
		hxrt.SysSetEnvironment(hxrt.StringFromLiteral("HAXE_GO=INVALID"), hxrt.StringFromLiteral("value"))
	}))))
	hxrt.Println(v_3)
	locked := hxrt.StringConcatStringPtr(root, hxrt.StringFromLiteral("/locked.txt"))
	sys__io__File_saveContent(locked, hxrt.StringFromLiteral("secret"))
	args := hxrt.NewArray(hxrt.StringFromLiteral("000"), locked)
	hxrt.SysCommand(hxrt.StringFromLiteral("chmod"), func() []*string {
		var hx_if_6 []*string
		if args == nil {
			hx_if_6 = nil
		} else {
			hx_if_6 = func(hx_lambda_raw_1 []any) []*string {
				hx_lambda_out_2 := make([]*string, 0, len(hx_lambda_raw_1))
				for _, hx_lambda_item_3 := range hx_lambda_raw_1 {
					hx_lambda_out_2 = append(hx_lambda_out_2, func(hx_value_4 any) *string {
						if hx_value_4 == nil {
							var hx_zero_5 *string
							return hx_zero_5
						}
						return hx_value_4.(*string)
					}(hx_lambda_item_3))
				}
				return hx_lambda_out_2
			}(args.Values())
		}
		return hx_if_6
	}())
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("permission.read.throws="), hxrt.StdString(throws(func() {
		sys__io__File_getContent(locked)
	}))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("permission.write.throws="), hxrt.StdString(throws(func() {
		sys__io__File_saveContent(locked, hxrt.StringFromLiteral("replacement"))
	}))))
	hxrt.Println(v_5)
	args_1 := hxrt.NewArray(hxrt.StringFromLiteral("600"), locked)
	hxrt.SysCommand(hxrt.StringFromLiteral("chmod"), func() []*string {
		var hx_if_12 []*string
		if args_1 == nil {
			hx_if_12 = nil
		} else {
			hx_if_12 = func(hx_lambda_raw_7 []any) []*string {
				hx_lambda_out_8 := make([]*string, 0, len(hx_lambda_raw_7))
				for _, hx_lambda_item_9 := range hx_lambda_raw_7 {
					hx_lambda_out_8 = append(hx_lambda_out_8, func(hx_value_10 any) *string {
						if hx_value_10 == nil {
							var hx_zero_11 *string
							return hx_zero_11
						}
						return hx_value_10.(*string)
					}(hx_lambda_item_9))
				}
				return hx_lambda_out_8
			}(args_1.Values())
		}
		return hx_if_12
	}())
	sys__FileSystem_deleteFile(locked)
	sys__FileSystem_deleteDirectory(root)
}

func throws(action func()) bool {
	hx_try_return_13 := false
	var hx_try_value_14 bool
	hxrt.TryCatch(func() {
		action()
		hx_try_value_14 = false
		hx_try_return_13 = true
		return
	}, func(hx_caught_15 any) {
		hx_tmp := hx_caught_15
		_ = hx_tmp
		hx_try_value_14 = true
		hx_try_return_13 = true
		return
	})
	if hx_try_return_13 {
		return hx_try_value_14
	}
	return false
}
