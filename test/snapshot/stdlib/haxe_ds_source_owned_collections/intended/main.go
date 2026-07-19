package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func main() {
	tree := New_haxe__ds__BalancedTree()
	tree.__hx_this.set(2, hxrt.StringFromLiteral("two"))
	tree.__hx_this.set(1, hxrt.StringFromLiteral("one"))
	tree.__hx_this.set(3, hxrt.StringFromLiteral("three"))
	stack := New_haxe__ds__GenericStack()
	stack.head = New_haxe__ds__GenericCell(hxrt.StringFromLiteral("alpha"), stack.head)
	stack.head = New_haxe__ds__GenericCell(hxrt.StringFromLiteral("beta"), stack.head)
	var v any = any(func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(tree.__hx_this.get(1)))
	hxrt.Println(v)
	var v_1 any = any(func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(stack.__hx_this.toString()))
	hxrt.Println(v_1)
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := obj.(type) {
	case *haxe__ds__BalancedTree:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__GenericStack:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__TreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe__ds__BalancedTree:
		return hxrt__generated_method_field__haxe__ds__BalancedTree(value, key)
	case *haxe__ds__GenericStack:
		return hxrt__generated_method_field__haxe__ds__GenericStack(value, key)
	case *haxe__ds__TreeNode:
		return hxrt__generated_method_field__haxe__ds__TreeNode(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__haxe__ds__BalancedTree(value *haxe__ds__BalancedTree, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "balance":
		return value.balance
	case "clear":
		return value.clear
	case "compare":
		return value.compare
	case "copy":
		return value.copy
	case "copyIMap":
		return value.copyIMap
	case "exists":
		return value.exists
	case "existsIMap":
		return value.existsIMap
	case "get":
		return value.get
	case "getIMap":
		return value.getIMap
	case "iterator":
		return value.iterator
	case "keyValueIterator":
		return value.keyValueIterator
	case "keys":
		return value.keys
	case "keysLoop":
		return value.keysLoop
	case "merge":
		return value.merge
	case "minBinding":
		return value.minBinding
	case "remove":
		return value.remove
	case "removeIMap":
		return value.removeIMap
	case "removeLoop":
		return value.removeLoop
	case "removeMinBinding":
		return value.removeMinBinding
	case "set":
		return value.set
	case "setIMap":
		return value.setIMap
	case "setLoop":
		return value.setLoop
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__GenericStack(value *haxe__ds__GenericStack, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return value.add
	case "first":
		return value.first
	case "isEmpty":
		return value.isEmpty
	case "iterator":
		return value.iterator
	case "pop":
		return value.pop
	case "remove":
		return value.remove
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__TreeNode(value *haxe__ds__TreeNode, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "toString":
		return value.toString
	}
	return nil
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	classValue, ok := value.(*hxrt__TypeClassValue)
	if !ok || classValue == nil {
		return nil, false
	}
	className := *hxrt.StdString(classValue.name)
	switch className {
	default:
		return nil, false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedField(object any, field *string) any {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__BalancedTree:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__GenericCell:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__GenericStack:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__TreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe__ds__BalancedTree:
		return hxrt__generated_field_lookup__haxe__ds__BalancedTree(value, key)
	case *haxe__ds__GenericCell:
		return hxrt__generated_field_lookup__haxe__ds__GenericCell(value, key)
	case *haxe__ds__GenericStack:
		return hxrt__generated_field_lookup__haxe__ds__GenericStack(value, key)
	case *haxe__ds__TreeNode:
		return hxrt__generated_field_lookup__haxe__ds__TreeNode(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__BalancedTree:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__GenericCell:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__GenericStack:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__TreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe__ds__BalancedTree:
		return hxrt__generated_field_has__haxe__ds__BalancedTree(value, key)
	case *haxe__ds__GenericCell:
		return hxrt__generated_field_has__haxe__ds__GenericCell(value, key)
	case *haxe__ds__GenericStack:
		return hxrt__generated_field_has__haxe__ds__GenericStack(value, key)
	case *haxe__ds__TreeNode:
		return hxrt__generated_field_has__haxe__ds__TreeNode(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := object.(type) {
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__BalancedTree:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__GenericCell:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__GenericStack:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__TreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe__ds__BalancedTree:
		return hxrt__generated_field_set__haxe__ds__BalancedTree(value, key, incoming)
	case *haxe__ds__GenericCell:
		return hxrt__generated_field_set__haxe__ds__GenericCell(value, key, incoming)
	case *haxe__ds__GenericStack:
		return hxrt__generated_field_set__haxe__ds__GenericStack(value, key, incoming)
	case *haxe__ds__TreeNode:
		return hxrt__generated_field_set__haxe__ds__TreeNode(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := object.(type) {
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__BalancedTree:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__GenericCell:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__GenericStack:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__TreeNode:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__BalancedTree:
		return hxrt.NewArray(hxrt.StringFromLiteral("root"))
	case *haxe__ds__GenericCell:
		return hxrt.NewArray(hxrt.StringFromLiteral("elt"), hxrt.StringFromLiteral("next"))
	case *haxe__ds__GenericStack:
		return hxrt.NewArray(hxrt.StringFromLiteral("head"))
	case *haxe__ds__TreeNode:
		return hxrt.NewArray(hxrt.StringFromLiteral("_height"), hxrt.StringFromLiteral("key"), hxrt.StringFromLiteral("left"), hxrt.StringFromLiteral("right"), hxrt.StringFromLiteral("value"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return value.high
	case "low":
		return value.low
	}
	return nil
}

func hxrt__generated_field_has__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		return true
	case "low":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		if incoming == nil {
			var zero int
			value.high = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.high = typed
			return true
		default:
			return false
		}
	case "low":
		if incoming == nil {
			var zero int
			value.low = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.low = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__BalancedTree(value *haxe__ds__BalancedTree, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "root":
		return value.root
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__BalancedTree(value *haxe__ds__BalancedTree, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "root":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__BalancedTree(value *haxe__ds__BalancedTree, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "root":
		if incoming == nil {
			var zero *haxe__ds__TreeNode
			value.root = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds__TreeNode:
			value.root = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__GenericCell(value *haxe__ds__GenericCell, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "elt":
		return value.elt
	case "next":
		return value.next
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__GenericCell(value *haxe__ds__GenericCell, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "elt":
		return true
	case "next":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__GenericCell(value *haxe__ds__GenericCell, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "elt":
		if incoming == nil {
			var zero any
			value.elt = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.elt = typed
			return true
		default:
			return false
		}
	case "next":
		if incoming == nil {
			var zero *haxe__ds__GenericCell
			value.next = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds__GenericCell:
			value.next = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__GenericStack(value *haxe__ds__GenericStack, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "head":
		return value.head
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__GenericStack(value *haxe__ds__GenericStack, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "head":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__GenericStack(value *haxe__ds__GenericStack, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "head":
		if incoming == nil {
			var zero *haxe__ds__GenericCell
			value.head = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds__GenericCell:
			value.head = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__TreeNode(value *haxe__ds__TreeNode, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "_height":
		return value._height
	case "key":
		return value.key
	case "left":
		return value.left
	case "right":
		return value.right
	case "value":
		return value.value
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__TreeNode(value *haxe__ds__TreeNode, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "_height":
		return true
	case "key":
		return true
	case "left":
		return true
	case "right":
		return true
	case "value":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__TreeNode(value *haxe__ds__TreeNode, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "_height":
		if incoming == nil {
			var zero int
			value._height = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value._height = typed
			return true
		default:
			return false
		}
	case "key":
		if incoming == nil {
			var zero any
			value.key = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.key = typed
			return true
		default:
			return false
		}
	case "left":
		if incoming == nil {
			var zero *haxe__ds__TreeNode
			value.left = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds__TreeNode:
			value.left = typed
			return true
		default:
			return false
		}
	case "right":
		if incoming == nil {
			var zero *haxe__ds__TreeNode
			value.right = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds__TreeNode:
			value.right = typed
			return true
		default:
			return false
		}
	case "value":
		if incoming == nil {
			var zero any
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.value = typed
			return true
		default:
			return false
		}
	}
	return false
}

func reflaxe__go___internal__CompilerReflect_typeField(object any, field *string) any {
	key := *hxrt.StdString(field)
	value, found := hxrt_typeClassMetadataField(object, key)
	if !found {
		return nil
	}
	return value
}

func reflaxe__go___internal__CompilerReflect_hasTypeField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	_, found := hxrt_typeClassMetadataField(object, key)
	return found
}

func reflaxe__go___internal__CompilerReflect_generatedMethod(object any, field *string) any {
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	return false
}
