package hxrt

import "sort"

// ArraySort orders one Array in place with a Haxe comparator.
//
// What: Apply a source comparator to the carrier's erased element storage.
// Why: Portable Array identity lives in hxrt, so sorting a detached Go slice
// would not update aliases that observe the same Haxe Array.
// How: Keep the existing carrier and let Go's stable sorter call the typed
// comparator adapter emitted at the Haxe-to-runtime boundary.
func ArraySort(array *Array, compare func(any, any) int) {
	sort.SliceStable(array.values, func(left, right int) bool {
		return compare(array.values[left], array.values[right]) < 0
	})
}
