package main

import "snapshot/hxrt"

type I_CollectionFeatureKey interface {
}

type CollectionFeatureKey struct {
	__hx_this I_CollectionFeatureKey
}

func (self *CollectionFeatureKey) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_CollectionFeatureKey() *CollectionFeatureKey {
	return hx_init_CollectionFeatureKey(hx_alloc_CollectionFeatureKey(nil))
}

func hx_alloc_CollectionFeatureKey(owner any) *CollectionFeatureKey {
	self := &CollectionFeatureKey{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_CollectionFeatureKey)
	}
	return self
}

func hx_init_CollectionFeatureKey(self *CollectionFeatureKey) *CollectionFeatureKey {
	return self
}

func main() {
	key := New_CollectionFeatureKey()
	values := New_haxe__ds__ObjectMap()
	values.__hx_this.set(key, hxrt.StringFromLiteral("one"))
	func(hx_value_1 any) bool {
		if hx_value_1 == nil {
			var hx_zero_2 bool
			return hx_zero_2
		}
		return hx_value_1.(bool)
	}(values.__hx_this.exists(key))
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
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(hxrt.FunctionCall(source)())
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
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *haxe__ds__ObjectMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__ObjectMap:
		return source
	default:
		return value.(*haxe__ds__ObjectMap)
	}
}
