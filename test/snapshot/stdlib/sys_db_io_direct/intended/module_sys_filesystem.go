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
	return hxrt.ArrayFromValues(func(hx_sort_src_30 []*string) []any {
		hx_sort_out_32 := make([]any, 0, len(hx_sort_src_30))
		for _, hx_sort_item_31 := range hx_sort_src_30 {
			hx_sort_out_32 = append(hx_sort_out_32, hx_sort_item_31)
		}
		return hx_sort_out_32
	}(hxrt.FileSystemReadDirectory(path)))
}

func sys__FileSystem_rename(path *string, newPath *string) {
	hxrt.FileSystemRename(path, newPath)
}

func sys__FileSystem_stat(path *string) map[string]any {
	value := hxrt.FileSystemStatPath(path)
	hx_obj_33 := map[string]any{}
	hx_obj_33["gid"] = value.Gid
	hx_obj_33["uid"] = value.Uid
	hx_obj_33["atime"] = Date_fromTime(value.AtimeMs)
	hx_obj_33["mtime"] = Date_fromTime(value.MtimeMs)
	hx_obj_33["ctime"] = Date_fromTime(value.CtimeMs)
	hx_obj_33["size"] = value.Size
	hx_obj_33["dev"] = value.Dev
	hx_obj_33["ino"] = value.Ino
	hx_obj_33["nlink"] = value.Nlink
	hx_obj_33["rdev"] = value.Rdev
	hx_obj_33["mode"] = value.Mode
	return hx_obj_33
}
