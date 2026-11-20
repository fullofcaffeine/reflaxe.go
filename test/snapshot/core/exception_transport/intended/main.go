package main

import "snapshot/hxrt"

type I_Child interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type Child struct {
	*Parent
	__hx_this I_Child
}

func (self *Child) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Child() *Child {
	return hx_init_Child(hx_alloc_Child(nil))
}

func hx_alloc_Child(owner any) *Child {
	self := &Child{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Child)
	}
	self.Parent = hx_alloc_Parent(self.__hx_this)
	return self
}

func hx_init_Child(self *Child) *Child {
	hx_init_Parent(self.Parent)
	return self
}

func (self *Child) toString() *string {
	return hxrt.StringFromLiteral("child-string")
}

func (self *Child) String() string {
	return *self.toString()
}

func main() {
	child := New_Child()
	base := func(hx_value *Child) *haxe__Exception {
		if hx_value == nil {
			return nil
		}
		return hx_value.Parent.haxe__Exception
	}(child)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("base-string="), base.__hx_this.toString()))
	hxrt.Println(v)
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(base))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		if hx_is_nominal_Parent(hx_exception_2) {
			var parent *Parent = hx_recover_class_Parent(hx_exception_2)
			var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("parent-catch="), hxrt.StdString(hxrt.HaxeEqual(parent, child))))
			hxrt.Println(v_1)
		} else {
			var other any = hx_payload_3
			_ = other
			hxrt.Println(any(hxrt.StringFromLiteral("parent-catch=false")))
		}
	})
	var hx_try_4 *string
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(child))
		hx_try_4 = hxrt.StringFromLiteral("missed")
	}, func(hx_caught_5 any) {
		hx_exception_6 := haxe__Exception_caught(hx_caught_5)
		hx_payload_7 := haxe__Exception_catchValue(hx_caught_5)
		_ = hx_payload_7
		var caught *haxe__Exception = hx_recover_class_haxe__Exception(hx_exception_6)
		hx_try_4 = caught.__hx_this.toString()
	})
	result := hx_try_4
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("expression="), result)))
	hxrt.TryCatch(func() {
		hxrt.TryCatch(func() {
			hxrt.Throw(haxe__Exception_thrown(child))
		}, func(hx_caught_11 any) {
			hx_exception_12 := haxe__Exception_caught(hx_caught_11)
			hx_payload_13 := haxe__Exception_catchValue(hx_caught_11)
			_ = hx_payload_13
			var caught_2 *haxe__Exception = hx_recover_class_haxe__Exception(hx_exception_12)
			hxrt.Throw(haxe__Exception_thrown(caught_2))
		})
	}, func(hx_caught_8 any) {
		hx_exception_9 := haxe__Exception_caught(hx_caught_8)
		hx_payload_10 := haxe__Exception_catchValue(hx_caught_8)
		if hx_is_nominal_Child(hx_exception_9) {
			var caught_1 *Child = hx_recover_class_Child(hx_exception_9)
			var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("rethrow="), hxrt.StdString((caught_1 == child))))
			hxrt.Println(v_2)
		} else {
			var other_1 any = hx_payload_10
			_ = other_1
			hxrt.Println(any(hxrt.StringFromLiteral("rethrow=false")))
		}
	})
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("base-not-value="), hxrt.StdString(!hx_is_nominal_haxe__ValueException(base))))
	hxrt.Println(v_3)
	hx_obj_14 := map[string]any{}
	hx_obj_14["marker"] = 42
	native := hxrt.AnonymousFromMap(hx_obj_14)
	attached := New_haxe__Exception(hxrt.StringFromLiteral("attached"), base, native)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native="), hxrt.StdString(hxrt.HaxeEqual(func() hxrt.Anonymous {
		var this1 any = attached.__hx_this.get_native()
		return hx_structural_recover(this1)
	}(), native))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("previous="), hxrt.StdString((attached.__hx_this.get_previous() == base))))
	hxrt.Println(v_5)
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(17))
	}, func(hx_caught_15 any) {
		hx_exception_16 := haxe__Exception_caught(hx_caught_15)
		hx_payload_17 := haxe__Exception_catchValue(hx_caught_15)
		_ = hx_payload_17
		var caught_3 *haxe__Exception = hx_recover_class_haxe__Exception(hx_exception_16)
		var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("primitive="), caught_3.__hx_this.get_message()))
		hxrt.Println(v_6)
		var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("primitive-value="), hxrt.StdString(hx_is_nominal_haxe__ValueException(caught_3))))
		hxrt.Println(v_7)
	})
}

type I_Parent interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type Parent struct {
	*haxe__Exception
	__hx_this I_Parent
}

func (self *Parent) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Parent() *Parent {
	return hx_init_Parent(hx_alloc_Parent(nil))
}

func hx_alloc_Parent(owner any) *Parent {
	self := &Parent{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Parent)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_Parent(self *Parent) *Parent {
	hx_init_haxe__Exception(self.haxe__Exception, hxrt.StringFromLiteral("parent"), nil, nil)
	return self
}

func (self *Parent) toString() *string {
	return hxrt.StringFromLiteral("parent-string")
}

func (self *Parent) String() string {
	return *self.__hx_this.toString()
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *Child:
		return func(hx_structural_source_1 *Child) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_2 := hx_structural_source_1
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_2.__hx_this.toString()
						}, func() any {
							return hx_receiver_2.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_1.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_1.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_1.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_1.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_5 := hx_structural_source_1
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_5.__hx_this.get_message()
						}, func() any {
							return hx_receiver_5.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_6 := hx_structural_source_1
						return hxrt.BoundFunction(func() any {
							return hx_receiver_6.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_6.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_7 := hx_structural_source_1
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_7.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_7.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_8 := hx_structural_source_1
						return hxrt.BoundFunction(func() any {
							return hx_receiver_8.__hx_this.get_native()
						}, func() any {
							return hx_receiver_8.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_9 := hx_structural_source_1
						return hxrt.BoundFunction(func() any {
							return hx_receiver_9.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_9.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_10 := hx_structural_source_1
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_10.__hx_this.details()
						}, func() any {
							return hx_receiver_10.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_1.savedMessage = func(hx_value_3 any) *string {
						if hx_value_3 == nil {
							var hx_zero_4 *string
							return hx_zero_4
						}
						return hx_value_3.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_1.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_1.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_1.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *Parent:
		return func(hx_structural_source_11 *Parent) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_11, func(name string) hxrt.AnonymousField {
				switch name {
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_12 := hx_structural_source_11
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_12.__hx_this.toString()
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_11.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_11.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_11.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_11.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_15 := hx_structural_source_11
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_15.__hx_this.get_message()
						}, func() any {
							return hx_receiver_15.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_16 := hx_structural_source_11
						return hxrt.BoundFunction(func() any {
							return hx_receiver_16.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_16.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_17 := hx_structural_source_11
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_17.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_18 := hx_structural_source_11
						return hxrt.BoundFunction(func() any {
							return hx_receiver_18.__hx_this.get_native()
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_19 := hx_structural_source_11
						return hxrt.BoundFunction(func() any {
							return hx_receiver_19.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_20 := hx_structural_source_11
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_20.__hx_this.details()
						}, func() any {
							return hx_receiver_20.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_11.savedMessage = func(hx_value_13 any) *string {
						if hx_value_13 == nil {
							var hx_zero_14 *string
							return hx_zero_14
						}
						return hx_value_13.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_11.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_11.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_11.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_21 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_21, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_21.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_21.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_21.high = func(hx_value_22 any) int {
						if hx_value_22 == nil {
							var hx_zero_23 int
							return hx_zero_23
						}
						return hx_value_22.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_21.low = func(hx_value_24 any) int {
						if hx_value_24 == nil {
							var hx_zero_25 int
							return hx_zero_25
						}
						return hx_value_24.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_26 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_26, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_26.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_27 := hx_structural_source_26
						return hxrt.BoundFunction(func() any {
							return hx_receiver_27.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_27.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_26.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_26.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_26.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_26.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_30 := hx_structural_source_26
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_30.__hx_this.get_message()
						}, func() any {
							return hx_receiver_30.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_31 := hx_structural_source_26
						return hxrt.BoundFunction(func() any {
							return hx_receiver_31.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_32 := hx_structural_source_26
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_32.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_32.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_33 := hx_structural_source_26
						return hxrt.BoundFunction(func() any {
							return hx_receiver_33.__hx_this.get_native()
						}, func() any {
							return hx_receiver_33.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_34 := hx_structural_source_26
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_34.__hx_this.toString()
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_35 := hx_structural_source_26
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_35.__hx_this.details()
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_26.savedMessage = func(hx_value_28 any) *string {
						if hx_value_28 == nil {
							var hx_zero_29 *string
							return hx_zero_29
						}
						return hx_value_28.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_26.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_26.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_26.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_36 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_36, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_36.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_36.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_36.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_36.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_39 := hx_structural_source_36
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_39.__hx_this.get_message()
						}, func() any {
							return hx_receiver_39.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_40 := hx_structural_source_36
						return hxrt.BoundFunction(func() any {
							return hx_receiver_40.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_40.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_41 := hx_structural_source_36
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_41.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_41.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_42 := hx_structural_source_36
						return hxrt.BoundFunction(func() any {
							return hx_receiver_42.__hx_this.get_native()
						}, func() any {
							return hx_receiver_42.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_43 := hx_structural_source_36
						return hxrt.BoundFunction(func() any {
							return hx_receiver_43.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_43.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_44 := hx_structural_source_36
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_44.__hx_this.toString()
						}, func() any {
							return hx_receiver_44.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_45 := hx_structural_source_36
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_45.__hx_this.details()
						}, func() any {
							return hx_receiver_45.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_36.savedMessage = func(hx_value_37 any) *string {
						if hx_value_37 == nil {
							var hx_zero_38 *string
							return hx_zero_38
						}
						return hx_value_37.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_36.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_36.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_36.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_class_Child(value any) *Child {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Child:
		return source
	default:
		return value.(*Child)
	}
}

func hx_recover_class_Parent(value any) *Parent {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Child:
		return source.Parent
	case *Parent:
		return source
	default:
		return value.(*Parent)
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Child:
		return source.Parent.haxe__Exception
	case *Parent:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
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

func hx_is_nominal_Child(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *Child:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_Parent(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *Child:
		return (source != nil)
	case *Parent:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *Child:
		return (source != nil)
	case *Parent:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
