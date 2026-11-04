package main

import "snapshot/hxrt"

func firstEntry(items *hxrt.Array) *string {
	var hx_if_3 *string
	if items.Len() > 0 {
		hx_if_3 = hxrt.StdString(func(hx_value_1 any) *string {
			if hx_value_1 == nil {
				var hx_zero_2 *string
				return hx_zero_2
			}
			return hx_value_1.(*string)
		}(items.Get(0)))
	} else {
		hx_if_3 = hxrt.StringFromLiteral("")
	}
	return hx_if_3
}

func main() {
	root := hxrt.StringFromLiteral("tmp_fs_smoke")
	fileA := hxrt.StringConcatStringPtr(root, hxrt.StringFromLiteral("/a.txt"))
	fileB := hxrt.StringConcatStringPtr(root, hxrt.StringFromLiteral("/b.txt"))
	rmDirRecursive(root)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("exists0="), hxrt.StdString(sys__FileSystem_exists(root))))
	hxrt.Println(v)
	missingThrows := false
	hxrt.TryCatch(func() {
		sys__FileSystem_isDirectory(root)
	}, func(hx_caught_4 any) {
		hx_exception_5 := haxe__Exception_caught(hx_caught_4)
		hx_payload_6 := haxe__Exception_catchValue(hx_caught_4)
		_ = hx_exception_5
		var hx_tmp any = hx_payload_6
		_ = hx_tmp
		missingThrows = true
	})
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("missing.throws="), hxrt.StdString(missingThrows)))
	hxrt.Println(v_1)
	sys__FileSystem_createDirectory(root)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dir1="), hxrt.StdString(sys__FileSystem_isDirectory(root))))
	hxrt.Println(v_2)
	absolute := sys__FileSystem_absolutePath(root)
	canonical := sys__FileSystem_fullPath(root)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("paths="), hxrt.StdString((sys__FileSystem_isDirectory(absolute) && sys__FileSystem_isDirectory(canonical)))))
	hxrt.Println(v_3)
	missingAbsolute := sys__FileSystem_absolutePath(hxrt.StringConcatStringPtr(root, hxrt.StringFromLiteral("/missing/child.txt")))
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("absolute.missing="), hxrt.StdString(!sys__FileSystem_exists(missingAbsolute))))
	hxrt.Println(v_4)
	directoryOnly := hxrt.StringConcatStringPtr(root, hxrt.StringFromLiteral("/directory-only"))
	sys__FileSystem_createDirectory(directoryOnly)
	deleteFileDirectoryThrows := false
	hxrt.TryCatch(func() {
		sys__FileSystem_deleteFile(directoryOnly)
	}, func(hx_caught_7 any) {
		hx_exception_8 := haxe__Exception_caught(hx_caught_7)
		hx_payload_9 := haxe__Exception_catchValue(hx_caught_7)
		_ = hx_exception_8
		var hx_tmp_1 any = hx_payload_9
		_ = hx_tmp_1
		deleteFileDirectoryThrows = true
	})
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("delete.file.directory.throws="), hxrt.StdString(deleteFileDirectoryThrows)))
	hxrt.Println(v_5)
	sys__FileSystem_deleteDirectory(directoryOnly)
	sys__io__File_saveContent(fileA, hxrt.StringFromLiteral("hello"))
	sys__FileSystem_rename(fileA, fileB)
	deleteDirectoryFileThrows := false
	hxrt.TryCatch(func() {
		sys__FileSystem_deleteDirectory(fileB)
	}, func(hx_caught_10 any) {
		hx_exception_11 := haxe__Exception_caught(hx_caught_10)
		hx_payload_12 := haxe__Exception_catchValue(hx_caught_10)
		_ = hx_exception_11
		var hx_tmp_2 any = hx_payload_12
		_ = hx_tmp_2
		deleteDirectoryFileThrows = true
	})
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("delete.directory.file.throws="), hxrt.StdString(deleteDirectoryFileThrows)))
	hxrt.Println(v_6)
	names := sys__FileSystem_readDirectory(root)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("entry="), firstEntry(names)))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("size="), func(hx_value_13 any) int {
		if hx_value_13 == nil {
			var hx_zero_14 int
			return hx_zero_14
		}
		return hx_value_13.(int)
	}(hxrt.AnonymousGet(sys__FileSystem_stat(fileB), "size"))))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("content="), sys__io__File_getContent(fileB)))
	hxrt.Println(v_9)
	sys__FileSystem_deleteFile(fileB)
	sys__FileSystem_deleteDirectory(root)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("exists1="), hxrt.StdString(sys__FileSystem_exists(root))))
	hxrt.Println(v_10)
}

func rmDirRecursive(path *string) {
	if !sys__FileSystem_exists(path) {
		return
	}
	_g := 0
	_g1 := sys__FileSystem_readDirectory(path)
	for _g < _g1.Len() {
		entry := func(hx_value_15 any) *string {
			if hx_value_15 == nil {
				var hx_zero_16 *string
				return hx_zero_16
			}
			return hx_value_15.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		child := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(path, hxrt.StringFromLiteral("/")), entry)
		if sys__FileSystem_isDirectory(child) {
			rmDirRecursive(child)
		} else {
			sys__FileSystem_deleteFile(child)
		}
	}
	sys__FileSystem_deleteDirectory(path)
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
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
