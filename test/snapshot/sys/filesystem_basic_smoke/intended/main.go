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
		hx_tmp := hx_caught_4
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
	}, func(hx_caught_6 any) {
		hx_tmp_1 := hx_caught_6
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
	}, func(hx_caught_8 any) {
		hx_tmp_2 := hx_caught_8
		_ = hx_tmp_2
		deleteDirectoryFileThrows = true
	})
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("delete.directory.file.throws="), hxrt.StdString(deleteDirectoryFileThrows)))
	hxrt.Println(v_6)
	names := sys__FileSystem_readDirectory(root)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("entry="), firstEntry(names)))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("size="), func(hx_obj_10 map[string]any) int {
		hx_field_11 := hx_obj_10["size"]
		if hx_field_11 == nil {
			var hx_zero_12 int
			return hx_zero_12
		}
		return hx_field_11.(int)
	}(sys__FileSystem_stat(fileB))))
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
		entry := func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
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
