package main

import "snapshot/hxrt"

func main() {
	var parsed any = New_haxe__format__JsonParser(hxrt.StringFromLiteral("[1,true,\"x\"]")).doParse()
	_ = parsed
	hxrt.Println(hxrt.JsonStringify(parsed))
}

type haxe__Json struct {
}

type haxe__format__JsonParser struct {
	source *string
}

func New_haxe__format__JsonParser(source *string) *haxe__format__JsonParser {
	return &haxe__format__JsonParser{source: source}
}

func (self *haxe__format__JsonParser) doParse() any {
	return hxrt.JsonParse(self.source)
}

func haxe__format__JsonPrinter_print(value any, rest ...any) *string {
	return hxrt.JsonStringify(value)
}

func haxe__Json_parse(source *string) any {
	return hxrt.JsonParse(source)
}

func haxe__Json_stringify(value any, rest ...any) *string {
	return hxrt.JsonStringify(value)
}
