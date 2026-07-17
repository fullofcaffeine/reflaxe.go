package hxrt

import "testing"

func TestHaxeEqualPreservesValuesAndReferenceIdentity(t *testing.T) {
	leftString := "same"
	rightString := "same"
	if !HaxeEqual(&leftString, &rightString) {
		t.Fatal("HaxeEqual rejected distinct string carriers with equal contents")
	}
	if !HaxeEqual(1, 1.0) {
		t.Fatal("HaxeEqual rejected equal Haxe numeric values with different Go carriers")
	}
	if HaxeEqual(1, 2.0) {
		t.Fatal("HaxeEqual accepted different numeric values")
	}

	type box struct{ value int }
	firstBox := &box{value: 1}
	if !HaxeEqual(firstBox, firstBox) {
		t.Fatal("HaxeEqual rejected the same object reference")
	}
	if HaxeEqual(firstBox, &box{value: 1}) {
		t.Fatal("HaxeEqual compared distinct objects by contents")
	}

	firstMap := map[string]int{"value": 1}
	if !HaxeEqual(firstMap, firstMap) {
		t.Fatal("HaxeEqual rejected the same map reference")
	}
	if HaxeEqual(firstMap, map[string]int{"value": 1}) {
		t.Fatal("HaxeEqual compared distinct maps by contents")
	}

	var nilBox *box
	if !HaxeEqual(nilBox, nil) {
		t.Fatal("HaxeEqual rejected equivalent typed and untyped nil values")
	}
}
