package main

import "examples_incident_api_portable/hxrt"

func sys__io__File_getContent(path *string) *string {
	return hxrt.StdString(hxrt.FileReadContent(path))
}

func sys__io__File_saveContent(path *string, content *string) {
	hxrt.FileWriteContent(path, content)
}
