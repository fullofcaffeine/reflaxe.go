package main

import "snapshot/hxrt"

func Registry_install(value *string) bool {
	Registry_selected = value
	return true
}

var Registry_selected *string = hxrt.StringFromLiteral("common")
