package main

type haxe__ds__Either struct {
	tag    int
	params []any
}

func (self *haxe__ds__Either) HxEnumValue() (string, []any) {
	return [2]string{"Left", "Right"}[self.tag], self.params
}

func haxe__ds__Either_Left(v any) *haxe__ds__Either {
	enumValue := &haxe__ds__Either{tag: 0}
	enumValue.params = []any{v}
	return enumValue
}

func haxe__ds__Either_Right(v any) *haxe__ds__Either {
	enumValue := &haxe__ds__Either{tag: 1}
	enumValue.params = []any{v}
	return enumValue
}
