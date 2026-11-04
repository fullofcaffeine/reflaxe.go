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

func hx_recover_class_0(value any) *haxe__io__Eof {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Eof:
		return source
	default:
		return value.(*haxe__io__Eof)
	}
}

func hx_recover_class_1(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_2(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}
