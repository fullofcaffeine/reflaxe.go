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
	switch value := hxrt.ObjectSource(obj).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
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
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
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

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_58 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_58.__hx_this.details()
			}, func() any {
				return hx_receiver_58.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_59 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_59.__hx_this.get_message()
			}, func() any {
				return hx_receiver_59.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_60 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_60.__hx_this.get_native()
			}, func() any {
				return hx_receiver_60.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_61 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_61.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_61.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_62 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_62.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_62.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_63 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_63.__hx_this.toString()
			}, func() any {
				return hx_receiver_63.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_64 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_64.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_64.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_65 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_65.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_65.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__haxe__ds__BalancedTree(value *haxe__ds__BalancedTree, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "balance":
		return func() *hxrt.Function[func(*haxe__ds__TreeNode, any, any, *haxe__ds__TreeNode) *haxe__ds__TreeNode] {
			hx_receiver_66 := value
			return hxrt.BoundFunction(func(l *haxe__ds__TreeNode, k any, v any, r *haxe__ds__TreeNode) *haxe__ds__TreeNode {
				return hx_receiver_66.__hx_this.balance(l, k, v, r)
			}, func() any {
				return hx_receiver_66.__hx_this
			}, "balance")
		}()
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_67 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_67.__hx_this.clear()
			}, func() any {
				return hx_receiver_67.__hx_this
			}, "clear")
		}()
	case "compare":
		return func() *hxrt.Function[func(any, any) int] {
			hx_receiver_68 := value
			return hxrt.BoundFunction(func(k1 any, k2 any) int {
				return hx_receiver_68.__hx_this.compare(k1, k2)
			}, func() any {
				return hx_receiver_68.__hx_this
			}, "compare")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__BalancedTree] {
			hx_receiver_69 := value
			return hxrt.BoundFunction(func() *haxe__ds__BalancedTree {
				return hx_receiver_69.__hx_this.copy()
			}, func() any {
				return hx_receiver_69.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_70 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_70.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_70.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_71 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_71.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_71.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_72 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_72.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_72.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_73 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_73.__hx_this.get(key)
			}, func() any {
				return hx_receiver_73.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_74 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_74.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_74.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_75 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_75.__hx_this.iterator()
			}, func() any {
				return hx_receiver_75.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_76 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_76.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_76.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_77 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_77.__hx_this.keys()
			}, func() any {
				return hx_receiver_77.__hx_this
			}, "keys")
		}()
	case "keysLoop":
		return func() *hxrt.Function[func(*haxe__ds__TreeNode, *hxrt.Array) *hxrt.Array] {
			hx_receiver_78 := value
			return hxrt.BoundFunction(func(node *haxe__ds__TreeNode, acc *hxrt.Array) *hxrt.Array {
				return hx_receiver_78.__hx_this.keysLoop(node, acc)
			}, func() any {
				return hx_receiver_78.__hx_this
			}, "keysLoop")
		}()
	case "merge":
		return func() *hxrt.Function[func(*haxe__ds__TreeNode, *haxe__ds__TreeNode) *haxe__ds__TreeNode] {
			hx_receiver_79 := value
			return hxrt.BoundFunction(func(t1 *haxe__ds__TreeNode, t2 *haxe__ds__TreeNode) *haxe__ds__TreeNode {
				return hx_receiver_79.__hx_this.merge(t1, t2)
			}, func() any {
				return hx_receiver_79.__hx_this
			}, "merge")
		}()
	case "minBinding":
		return func() *hxrt.Function[func(*haxe__ds__TreeNode) *haxe__ds__TreeNode] {
			hx_receiver_80 := value
			return hxrt.BoundFunction(func(t *haxe__ds__TreeNode) *haxe__ds__TreeNode {
				return hx_receiver_80.__hx_this.minBinding(t)
			}, func() any {
				return hx_receiver_80.__hx_this
			}, "minBinding")
		}()
	case "remove":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_81 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_81.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_81.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_82 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_82.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_82.__hx_this
			}, "removeIMap")
		}()
	case "removeLoop":
		return func() *hxrt.Function[func(any, *haxe__ds__TreeNode) *haxe__ds__TreeNode] {
			hx_receiver_83 := value
			return hxrt.BoundFunction(func(k any, node *haxe__ds__TreeNode) *haxe__ds__TreeNode {
				return hx_receiver_83.__hx_this.removeLoop(k, node)
			}, func() any {
				return hx_receiver_83.__hx_this
			}, "removeLoop")
		}()
	case "removeMinBinding":
		return func() *hxrt.Function[func(*haxe__ds__TreeNode) *haxe__ds__TreeNode] {
			hx_receiver_84 := value
			return hxrt.BoundFunction(func(t *haxe__ds__TreeNode) *haxe__ds__TreeNode {
				return hx_receiver_84.__hx_this.removeMinBinding(t)
			}, func() any {
				return hx_receiver_84.__hx_this
			}, "removeMinBinding")
		}()
	case "set":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_85 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_85.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_85.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_86 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_86.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_86.__hx_this
			}, "setIMap")
		}()
	case "setLoop":
		return func() *hxrt.Function[func(any, any, *haxe__ds__TreeNode) *haxe__ds__TreeNode] {
			hx_receiver_87 := value
			return hxrt.BoundFunction(func(k any, v any, node *haxe__ds__TreeNode) *haxe__ds__TreeNode {
				return hx_receiver_87.__hx_this.setLoop(k, v, node)
			}, func() any {
				return hx_receiver_87.__hx_this
			}, "setLoop")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_88 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_88.__hx_this.toString()
			}, func() any {
				return hx_receiver_88.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__GenericStack(value *haxe__ds__GenericStack, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return func() *hxrt.Function[func(any)] {
			hx_receiver_89 := value
			return hxrt.BoundFunction(func(item any) {
				hx_receiver_89.__hx_this.add(item)
			}, func() any {
				return hx_receiver_89.__hx_this
			}, "add")
		}()
	case "first":
		return func() *hxrt.Function[func() any] {
			hx_receiver_90 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_90.__hx_this.first()
			}, func() any {
				return hx_receiver_90.__hx_this
			}, "first")
		}()
	case "isEmpty":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_91 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_91.__hx_this.isEmpty()
			}, func() any {
				return hx_receiver_91.__hx_this
			}, "isEmpty")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_92 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_92.__hx_this.iterator()
			}, func() any {
				return hx_receiver_92.__hx_this
			}, "iterator")
		}()
	case "pop":
		return func() *hxrt.Function[func() any] {
			hx_receiver_93 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_93.__hx_this.pop()
			}, func() any {
				return hx_receiver_93.__hx_this
			}, "pop")
		}()
	case "remove":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_94 := value
			return hxrt.BoundFunction(func(v any) bool {
				return hx_receiver_94.__hx_this.remove(v)
			}, func() any {
				return hx_receiver_94.__hx_this
			}, "remove")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_95 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_95.__hx_this.toString()
			}, func() any {
				return hx_receiver_95.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__TreeNode(value *haxe__ds__TreeNode, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_96 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_96.__hx_this.toString()
			}, func() any {
				return hx_receiver_96.__hx_this
			}, "toString")
		}()
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
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
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
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
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
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
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
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
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
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
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
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
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
	switch value := hxrt.ObjectSource(object).(type) {
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
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
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__BalancedTree:
		return hxrt.NewArray(hxrt.StringFromLiteral("root"))
	case *haxe__ds__GenericCell:
		return hxrt.NewArray(hxrt.StringFromLiteral("elt"), hxrt.StringFromLiteral("next"))
	case *haxe__ds__GenericStack:
		return hxrt.NewArray(hxrt.StringFromLiteral("head"))
	case *haxe__ds__TreeNode:
		return hxrt.NewArray(hxrt.StringFromLiteral("left"), hxrt.StringFromLiteral("right"), hxrt.StringFromLiteral("key"), hxrt.StringFromLiteral("value"), hxrt.StringFromLiteral("_height"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "savedMessage":
		return hxrt.ReadStoredField(value.savedMessage)
	case "savedNative":
		return hxrt.ReadStoredField(value.savedNative)
	case "savedPrevious":
		return hxrt.ReadStoredField(value.savedPrevious)
	case "savedStack":
		return hxrt.ReadStoredField(value.savedStack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Exception(value *haxe__Exception, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		return true
	case "savedNative":
		return true
	case "savedPrevious":
		return true
	case "savedStack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Exception(value *haxe__Exception, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		if incoming == nil {
			var zero *string
			value.savedMessage = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.savedMessage = typed
			return true
		default:
			return false
		}
	case "savedNative":
		if incoming == nil {
			var zero any
			value.savedNative = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedNative = typed
			return true
		default:
			return false
		}
	case "savedPrevious":
		if incoming == nil {
			var zero *haxe__Exception
			value.savedPrevious = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__Exception:
			value.savedPrevious = typed
			return true
		default:
			return false
		}
	case "savedStack":
		if incoming == nil {
			var zero any
			value.savedStack = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedStack = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__ValueException(value *haxe__ValueException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__ValueException(value *haxe__ValueException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
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
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return hxrt.ReadStoredField(value.high)
	case "low":
		return hxrt.ReadStoredField(value.low)
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
		return hxrt.ReadStoredField(value.root)
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
		return hxrt.ReadStoredField(value.elt)
	case "next":
		return hxrt.ReadStoredField(value.next)
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
		return hxrt.ReadStoredField(value.head)
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
		return hxrt.ReadStoredField(value._height)
	case "key":
		return hxrt.ReadStoredField(value.key)
	case "left":
		return hxrt.ReadStoredField(value.left)
	case "right":
		return hxrt.ReadStoredField(value.right)
	case "value":
		return hxrt.ReadStoredField(value.value)
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
	switch enumValue := value.(type) {
	case *haxe__StackItem:
		return (enumValue != nil)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*haxe__ds__TreeNode, any, any, *haxe__ds__TreeNode) *haxe__ds__TreeNode]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		var argument_3 any = nil
		if 3 < len(arguments) {
			argument_3 = arguments[3]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_0(argument_0), argument_1, argument_2, hx_recover_class_0(argument_3))
	case *hxrt.Function[func()]:
		hxrt.FunctionCall(source)()
		return nil
	case *hxrt.Function[func(any, any) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(argument_0, argument_1)
	case *hxrt.Function[func() *haxe__ds__BalancedTree]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(any) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(any) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(*haxe__ds__TreeNode, *hxrt.Array) *hxrt.Array]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_0(argument_0), func(hx_value_97 any) *hxrt.Array {
			if hx_value_97 == nil {
				var hx_zero_98 *hxrt.Array
				return hx_zero_98
			}
			return hx_value_97.(*hxrt.Array)
		}(any(argument_1)))
	case *hxrt.Function[func(*haxe__ds__TreeNode, *haxe__ds__TreeNode) *haxe__ds__TreeNode]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_0(argument_0), hx_recover_class_0(argument_1))
	case *hxrt.Function[func(*haxe__ds__TreeNode) *haxe__ds__TreeNode]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_0(argument_0))
	case *hxrt.Function[func(any, *haxe__ds__TreeNode) *haxe__ds__TreeNode]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(argument_0, hx_recover_class_0(argument_1))
	case *hxrt.Function[func(any, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(argument_0, argument_1)
		return nil
	case *hxrt.Function[func(any, any, *haxe__ds__TreeNode) *haxe__ds__TreeNode]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(argument_0, argument_1, hx_recover_class_0(argument_2))
	case *hxrt.Function[func(any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(argument_0)
		return nil
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_99 any) bool {
				if hx_value_99 == nil {
					var hx_zero_100 bool
					return hx_zero_100
				}
				return hx_value_99.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_101 any) bool {
				if hx_value_101 == nil {
					var hx_zero_102 bool
					return hx_zero_102
				}
				return hx_value_101.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_103 any) bool {
				if hx_value_103 == nil {
					var hx_zero_104 bool
					return hx_zero_104
				}
				return hx_value_103.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_105 any) bool {
				if hx_value_105 == nil {
					var hx_zero_106 bool
					return hx_zero_106
				}
				return hx_value_105.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_107 any) bool {
				if hx_value_107 == nil {
					var hx_zero_108 bool
					return hx_zero_108
				}
				return hx_value_107.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_109 any) bool {
				if hx_value_109 == nil {
					var hx_zero_110 bool
					return hx_zero_110
				}
				return hx_value_109.(bool)
			}(any(nil))
		})
	case *hxrt.Function[func() *haxe__ds__BalancedTree]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_111 any) bool {
				if hx_value_111 == nil {
					var hx_zero_112 bool
					return hx_zero_112
				}
				return hx_value_111.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_113 any) bool {
				if hx_value_113 == nil {
					var hx_zero_114 bool
					return hx_zero_114
				}
				return hx_value_113.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() *haxe__ds__BalancedTree]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *haxe__ds__TreeNode {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__TreeNode:
		return source
	default:
		return value.(*haxe__ds__TreeNode)
	}
}

func hx_recover_class_1(value any) *haxe__ds__BalancedTree {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__BalancedTree:
		return source
	default:
		return value.(*haxe__ds__BalancedTree)
	}
}

func hx_recover_class_2(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_3(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
