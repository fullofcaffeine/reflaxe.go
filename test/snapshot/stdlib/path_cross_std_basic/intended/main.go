package main

import "snapshot/hxrt"

func main() {
	unix := New_haxe__io__Path(hxrt.StringFromLiteral("/tmp/demo.txt"))
	var v any = any(unix.dir)
	hxrt.Println(v)
	var v_1 any = any(unix.file)
	hxrt.Println(v_1)
	var v_2 any = any(unix.ext)
	hxrt.Println(v_2)
	var v_3 any = any(unix.__hx_this.toString())
	hxrt.Println(v_3)
	dot := New_haxe__io__Path(hxrt.StringFromLiteral("."))
	var v_4 any = any(dot.dir)
	hxrt.Println(v_4)
	var v_5 any = any(dot.file)
	hxrt.Println(v_5)
	var v_6 any = any(dot.ext)
	hxrt.Println(v_6)
	var v_7 any = any(dot.__hx_this.toString())
	hxrt.Println(v_7)
	var v_8 any = any(haxe__io__Path_withoutExtension(hxrt.StringFromLiteral("/tmp/demo.txt")))
	hxrt.Println(v_8)
	var v_9 any = any(haxe__io__Path_withoutDirectory(hxrt.StringFromLiteral("/tmp/demo.txt")))
	hxrt.Println(v_9)
	var v_10 any = any(haxe__io__Path_directory(hxrt.StringFromLiteral("demo.txt")))
	hxrt.Println(v_10)
	var v_11 any = any(haxe__io__Path_extension(hxrt.StringFromLiteral("/tmp/demo.txt")))
	hxrt.Println(v_11)
	var v_12 any = any(haxe__io__Path_withExtension(hxrt.StringFromLiteral("/tmp/demo.txt"), hxrt.StringFromLiteral("log")))
	hxrt.Println(v_12)
	var v_13 any = any(haxe__io__Path_join(hxrt.NewArray(hxrt.StringFromLiteral("/tmp"), hxrt.StringFromLiteral("demo"), hxrt.StringFromLiteral(".."), hxrt.StringFromLiteral("out"), hxrt.StringFromLiteral("file.txt"))))
	hxrt.Println(v_13)
	var v_14 any = any(haxe__io__Path_normalize(hxrt.StringFromLiteral("/usr/local/../lib//./a\\b")))
	hxrt.Println(v_14)
	var v_15 any = any(haxe__io__Path_addTrailingSlash(hxrt.StringFromLiteral("a\\b")))
	hxrt.Println(v_15)
	var v_16 any = any(haxe__io__Path_addTrailingSlash(hxrt.StringFromLiteral("a/b")))
	hxrt.Println(v_16)
	var v_17 any = any(haxe__io__Path_removeTrailingSlashes(hxrt.StringFromLiteral("a///")))
	hxrt.Println(v_17)
	var v_18 any = any(haxe__io__Path_isAbsolute(hxrt.StringFromLiteral("/tmp/demo.txt")))
	hxrt.Println(v_18)
	var v_19 any = any(haxe__io__Path_isAbsolute(hxrt.StringFromLiteral("C:\\tmp\\demo.txt")))
	hxrt.Println(v_19)
	var v_20 any = any(haxe__io__Path_isAbsolute(hxrt.StringFromLiteral("\\\\server\\share")))
	hxrt.Println(v_20)
	var v_21 any = any(haxe__io__Path_isAbsolute(hxrt.StringFromLiteral("relative/path")))
	hxrt.Println(v_21)
}

type Std struct {
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}
