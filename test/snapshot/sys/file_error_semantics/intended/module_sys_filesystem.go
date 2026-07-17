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
	return hxrt.ArrayFromValues(func(hx_sort_src_33 []*string) []any {
		hx_sort_out_35 := make([]any, 0, len(hx_sort_src_33))
		for _, hx_sort_item_34 := range hx_sort_src_33 {
			hx_sort_out_35 = append(hx_sort_out_35, hx_sort_item_34)
		}
		return hx_sort_out_35
	}(hxrt.FileSystemReadDirectory(path)))
}

func sys__FileSystem_rename(path *string, newPath *string) {
	hxrt.FileSystemRename(path, newPath)
}

func sys__FileSystem_stat(path *string) map[string]any {
	value := hxrt.FileSystemStatPath(path)
	hx_obj_36 := map[string]any{}
	hx_obj_36["gid"] = value.Gid
	hx_obj_36["uid"] = value.Uid
	hx_obj_36["atime"] = Date_fromTime(value.AtimeMs)
	hx_obj_36["mtime"] = Date_fromTime(value.MtimeMs)
	hx_obj_36["ctime"] = Date_fromTime(value.CtimeMs)
	hx_obj_36["size"] = value.Size
	hx_obj_36["dev"] = value.Dev
	hx_obj_36["ino"] = value.Ino
	hx_obj_36["nlink"] = value.Nlink
	hx_obj_36["rdev"] = value.Rdev
	hx_obj_36["mode"] = value.Mode
	return hx_obj_36
}
