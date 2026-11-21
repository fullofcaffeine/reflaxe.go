package main

import "examples_portable_beta/hxrt"

func sys__FileSystem_deleteFile(path *string) {
	hxrt.FileSystemDeleteFile(path)
}

func sys__FileSystem_exists(path *string) bool {
	return hxrt.FileSystemExists(path)
}
