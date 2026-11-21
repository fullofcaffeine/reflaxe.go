package main

type haxe__io__Encoding struct {
	tag    int
	params []any
}

func (self *haxe__io__Encoding) HxEnumValue() (string, []any) {
	return [2]string{"UTF8", "RawNative"}[self.tag], self.params
}

var haxe__io__Encoding_UTF8 *haxe__io__Encoding = &haxe__io__Encoding{tag: 0}

var haxe__io__Encoding_RawNative *haxe__io__Encoding = &haxe__io__Encoding{tag: 1}
