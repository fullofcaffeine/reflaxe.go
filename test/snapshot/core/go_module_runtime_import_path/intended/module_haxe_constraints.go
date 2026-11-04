package main

import "example.com/acme/reflaxe-go-smoke/hxrt"

type haxe__IMap interface {
	getIMap(k any) any
	setIMap(k any, v any)
	existsIMap(k any) bool
	removeIMap(k any) bool
	keys() hxrt.Anonymous
	iterator() hxrt.Anonymous
	keyValueIterator() hxrt.Anonymous
	copyIMap() haxe__IMap
	toString() *string
	clear()
}
