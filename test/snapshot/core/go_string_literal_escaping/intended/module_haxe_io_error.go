package main

type haxe__io__Error struct {
	tag    int
	params []any
}

var haxe__io__Error_Blocked *haxe__io__Error = &haxe__io__Error{tag: 0}

var haxe__io__Error_Overflow *haxe__io__Error = &haxe__io__Error{tag: 1}

var haxe__io__Error_OutsideBounds *haxe__io__Error = &haxe__io__Error{tag: 2}

func haxe__io__Error_Custom(e any) *haxe__io__Error {
	enumValue := &haxe__io__Error{tag: 3}
	enumValue.params = []any{e}
	return enumValue
}
