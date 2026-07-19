package main

type haxe__ds__Option struct {
	tag    int
	params []any
}

func haxe__ds__Option_Some(v any) *haxe__ds__Option {
	enumValue := &haxe__ds__Option{tag: 0}
	enumValue.params = []any{v}
	return enumValue
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1}
