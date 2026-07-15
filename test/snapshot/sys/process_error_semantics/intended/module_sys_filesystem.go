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

func sys__FileSystem_readDirectory(path *string) []*string {
	return hxrt.FileSystemReadDirectory(path)
}

func sys__FileSystem_rename(path *string, newPath *string) {
	hxrt.FileSystemRename(path, newPath)
}

func sys__FileSystem_stat(path *string) map[string]any {
	value := hxrt.FileSystemStatPath(path)
	hx_obj_21 := map[string]any{}
	hx_obj_21["gid"] = value.Gid
	hx_obj_21["uid"] = value.Uid
	hx_obj_21["atime"] = Date_fromTime(value.AtimeMs)
	hx_obj_21["mtime"] = Date_fromTime(value.MtimeMs)
	hx_obj_21["ctime"] = Date_fromTime(value.CtimeMs)
	hx_obj_21["size"] = value.Size
	hx_obj_21["dev"] = value.Dev
	hx_obj_21["ino"] = value.Ino
	hx_obj_21["nlink"] = value.Nlink
	hx_obj_21["rdev"] = value.Rdev
	hx_obj_21["mode"] = value.Mode
	return hx_obj_21
}
