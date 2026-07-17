package main

import "snapshot/hxrt"

func sys__FileSystem_absolutePath(relPath *string) *string {
	return hxrt.StdString(hxrt.FileSystemAbsolutePath(relPath))
}

func sys__FileSystem_createDirectory(path *string) {
	hxrt.FileSystemCreateDirectory(path)
}

func sys__FileSystem_deleteDirectory(path *string) {
	hxrt.FileSystemDeleteDirectory(path)
}

func sys__FileSystem_deleteFile(path *string) {
	hxrt.FileSystemDeleteFile(path)
}

func sys__FileSystem_exists(path *string) bool {
	return hxrt.FileSystemExists(path)
}

func sys__FileSystem_fullPath(relPath *string) *string {
	return hxrt.StdString(hxrt.FileSystemFullPath(relPath))
}

func sys__FileSystem_isDirectory(path *string) bool {
	return hxrt.FileSystemIsDirectory(path)
}

func sys__FileSystem_readDirectory(path *string) *hxrt.Array {
	return hxrt.ArrayFromValues(func(hx_sort_src_31 []*string) []any {
		hx_sort_out_33 := make([]any, 0, len(hx_sort_src_31))
		for _, hx_sort_item_32 := range hx_sort_src_31 {
			hx_sort_out_33 = append(hx_sort_out_33, hx_sort_item_32)
		}
		return hx_sort_out_33
	}(hxrt.FileSystemReadDirectory(path)))
}

func sys__FileSystem_rename(path *string, newPath *string) {
	hxrt.FileSystemRename(path, newPath)
}

func sys__FileSystem_stat(path *string) map[string]any {
	value := hxrt.FileSystemStatPath(path)
	hx_obj_34 := map[string]any{}
	hx_obj_34["gid"] = value.Gid
	hx_obj_34["uid"] = value.Uid
	hx_obj_34["atime"] = Date_fromTime(value.AtimeMs)
	hx_obj_34["mtime"] = Date_fromTime(value.MtimeMs)
	hx_obj_34["ctime"] = Date_fromTime(value.CtimeMs)
	hx_obj_34["size"] = value.Size
	hx_obj_34["dev"] = value.Dev
	hx_obj_34["ino"] = value.Ino
	hx_obj_34["nlink"] = value.Nlink
	hx_obj_34["rdev"] = value.Rdev
	hx_obj_34["mode"] = value.Mode
	return hx_obj_34
}
