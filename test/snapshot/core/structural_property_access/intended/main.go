package main

import "snapshot/hxrt"

type I_Counter interface {
	get_value() int
	set_value(value int) int
}

type Counter struct {
	__hx_this I_Counter
	value     int
	reads     int
	writes    int
	stored    int
}

func (self *Counter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Counter() *Counter {
	return hx_init_Counter(hx_alloc_Counter(nil))
}

func hx_alloc_Counter(owner any) *Counter {
	self := &Counter{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Counter)
	}
	return self
}

func hx_init_Counter(self *Counter) *Counter {
	self.stored = 4
	self.writes = 0
	self.reads = 0
	return self
}

func (self *Counter) get_value() int {
	self.reads = int(int32((self.reads + 1)))
	return self.stored
}

func (self *Counter) set_value(value int) int {
	self.writes = int(int32((self.writes + 1)))
	self.stored = value
	return value
}

func main() {
	source := New_Counter()
	var erased any = source
	view := hx_structural_recover(hx_structural_recover(erased))
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("before="), func() int {
		var hx_callable_1 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(view, "get_value"))
		return hxrt.FunctionCall(hx_callable_1)()
	}()))
	hxrt.Println(v)
	assigned := func() int {
		var hx_callable_2 *hxrt.Function[func(int) int] = hx_recover_callable_1(hxrt.AnonymousGet(view, "set_value"))
		var hx_argument_3 int = 9
		return hxrt.FunctionCall(hx_callable_2)(hx_argument_3)
	}()
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("assigned="), assigned)))
	func() int {
		var hx_callable_5 *hxrt.Function[func(int) int] = hx_recover_callable_1(hxrt.AnonymousGet(view, "set_value"))
		var hx_argument_6 int = int((hxrt.Int32Wrap(func() int {
			var hx_callable_4 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(view, "get_value"))
			return hxrt.FunctionCall(hx_callable_4)()
		}()) + hxrt.Int32Wrap(3)))
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}()
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("after="), func() int {
		var hx_callable_7 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(view, "get_value"))
		return hxrt.FunctionCall(hx_callable_7)()
	}()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("reads="), source.reads))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("writes="), source.writes))
	hxrt.Println(v_3)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *Counter:
		return func(hx_structural_source_15 *Counter) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_15, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_15.value, Found: true}
				case "reads":
					return hxrt.AnonymousField{Value: hx_structural_source_15.reads, Found: true}
				case "writes":
					return hxrt.AnonymousField{Value: hx_structural_source_15.writes, Found: true}
				case "stored":
					return hxrt.AnonymousField{Value: hx_structural_source_15.stored, Found: true}
				case "get_value":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_22 := hx_structural_source_15
						return hxrt.BoundFunction(func() int {
							return hx_receiver_22.__hx_this.get_value()
						}, func() any {
							return hx_receiver_22.__hx_this
						}, "get_value")
					}(), Found: true}
				case "set_value":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_23 := hx_structural_source_15
						return hxrt.BoundFunction(func(value int) int {
							return hx_receiver_23.__hx_this.set_value(value)
						}, func() any {
							return hx_receiver_23.__hx_this
						}, "set_value")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "reads":
					hx_structural_source_15.reads = func(hx_value_16 any) int {
						if hx_value_16 == nil {
							var hx_zero_17 int
							return hx_zero_17
						}
						return hx_value_16.(int)
					}(incoming)
					return true
				case "writes":
					hx_structural_source_15.writes = func(hx_value_18 any) int {
						if hx_value_18 == nil {
							var hx_zero_19 int
							return hx_zero_19
						}
						return hx_value_18.(int)
					}(incoming)
					return true
				case "stored":
					hx_structural_source_15.stored = func(hx_value_20 any) int {
						if hx_value_20 == nil {
							var hx_zero_21 int
							return hx_zero_21
						}
						return hx_value_20.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_24 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_24, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_24.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_24.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_24.high = func(hx_value_25 any) int {
						if hx_value_25 == nil {
							var hx_zero_26 int
							return hx_zero_26
						}
						return hx_value_25.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_24.low = func(hx_value_27 any) int {
						if hx_value_27 == nil {
							var hx_zero_28 int
							return hx_zero_28
						}
						return hx_value_27.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func(int) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
