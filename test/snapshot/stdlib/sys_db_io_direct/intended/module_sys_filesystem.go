package main

import "snapshot/hxrt"

func sys__FileSystem_deleteFile(path *string) {
	hxrt.FileSystemDeleteFile(path)
}
