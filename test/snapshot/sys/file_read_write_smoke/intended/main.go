package main

import "snapshot/hxrt"

func main() {
	path := hxrt.StringFromLiteral("./tmp_sys_file_smoke.txt")
	sys__io__File_saveContent(path, hxrt.StringFromLiteral("hello"))
	content := sys__io__File_getContent(path)
	hxrt.Println(any(content))
}
