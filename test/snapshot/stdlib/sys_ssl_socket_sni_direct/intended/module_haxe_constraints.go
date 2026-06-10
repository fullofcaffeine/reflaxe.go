package main

type haxe__IMap interface {
	get(k any) any
	set(k any, v any)
	exists(k any) bool
	remove(k any) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() map[string]any
	copyIMap() haxe__IMap
	toString() *string
	clear()
}
