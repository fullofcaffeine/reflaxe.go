package main

import "snapshot/hxrt"

func main() {
	path := hxrt.StringFromLiteral("runtime_hxrt_infer_file.bin")
	sys__io__File_saveBytes(path, haxe__io__Bytes_ofHex(hxrt.StringFromLiteral("00ff80")))
	output := sys__io__File_update(path, true)
	output.__hx_this.seek(1, sys__io__FileSeek_SeekBegin)
	output.__hx_this.writeByte(7)
	output.__hx_this.close()
	input := sys__io__File_read(path, true)
	input.__hx_this.readByte()
	input.__hx_this.close()
}
