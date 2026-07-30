package main

type haxe__IMap interface {
	getIMap(k any) any
	setIMap(k any, v any)
	existsIMap(k any) bool
	removeIMap(k any) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() map[string]any
	copyIMap() haxe__IMap
	toString() *string
	clear()
}
