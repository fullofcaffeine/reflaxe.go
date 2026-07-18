package main

import "examples_incident_api_portable/hxrt"

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
	return hxrt.ArrayFromValues(func(hx_sort_src_147 []*string) []any {
		hx_sort_out_149 := make([]any, 0, len(hx_sort_src_147))
		for _, hx_sort_item_148 := range hx_sort_src_147 {
			hx_sort_out_149 = append(hx_sort_out_149, hx_sort_item_148)
		}
		return hx_sort_out_149
	}(hxrt.FileSystemReadDirectory(path)))
}

func sys__FileSystem_rename(path *string, newPath *string) {
	hxrt.FileSystemRename(path, newPath)
}

func sys__FileSystem_stat(path *string) map[string]any {
	value := hxrt.FileSystemStatPath(path)
	hx_obj_150 := map[string]any{}
	hx_obj_150["gid"] = value.Gid
	hx_obj_150["uid"] = value.Uid
	hx_obj_150["atime"] = Date_fromTime(value.AtimeMs)
	hx_obj_150["mtime"] = Date_fromTime(value.MtimeMs)
	hx_obj_150["ctime"] = Date_fromTime(value.CtimeMs)
	hx_obj_150["size"] = value.Size
	hx_obj_150["dev"] = value.Dev
	hx_obj_150["ino"] = value.Ino
	hx_obj_150["nlink"] = value.Nlink
	hx_obj_150["rdev"] = value.Rdev
	hx_obj_150["mode"] = value.Mode
	return hx_obj_150
}
