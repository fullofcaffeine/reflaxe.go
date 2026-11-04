package hxrt

import "fmt"

// Print writes a Haxe display value without adding a newline.
//
// What: Implements the portable Sys.print(Dynamic) output contract.
// Why: Dynamic is inherent at this public Haxe boundary; converting through
// StdString keeps that untyped island out of the compiler and preserves Haxe
// null/value formatting.
// How: Format exactly once through StdString, then write with fmt.Print.
func Print(value any) {
	fmt.Print(*StdString(value))
}

func Println(value any) {
	fmt.Println(*StdString(value))
}
