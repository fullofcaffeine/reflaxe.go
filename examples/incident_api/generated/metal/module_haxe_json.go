package main

import "examples_incident_api_metal/hxrt"

func haxe__Json_parse(text *string) any {
	return hxrt.JsonParse(text)
}

func haxe__Json_stringify(value any, replacer func(any, any) any, space *string) *string {
	return hxrt.StdString(hxrt.JsonStringify(value))
}
