package main

import "snapshot/hxrt"

func sys__io__File_read(path *string, binary bool) *sys__io__FileInput {
	return New_sys__io__FileInput(hxrt.FileOpenInput(path))
}

func sys__io__File_write(path *string, binary bool) *sys__io__FileOutput {
	return New_sys__io__FileOutput(hxrt.FileOpenWrite(path))
}
