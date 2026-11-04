package main

import "snapshot/hxrt"

func main() {
	var value any = New_SameShape()
	if hx_is_nominal_0(value) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("structural interface match violated nominal membership")))
	}
	var v any = any(Registry_count)
	hxrt.Println(v)
}

type I_SameShape interface {
	code() int
}

type SameShape struct {
	__hx_this I_SameShape
}

func (self *SameShape) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_SameShape() *SameShape {
	return hx_init_SameShape(hx_alloc_SameShape(nil))
}

func hx_alloc_SameShape(owner any) *SameShape {
	self := &SameShape{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_SameShape)
	}
	return self
}

func hx_init_SameShape(self *SameShape) *SameShape {
	return self
}

func (self *SameShape) code() int {
	return 7
}

func hx_recover_class_1(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	}
	hx_recovered, hx_matched := hx_recover_class_1__probe_0(value)
	if hx_matched {
		return hx_recovered
	}
	return value.(*haxe__Exception)
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	hx_recovered, hx_matched := hx_is_nominal_0__probe_0(value)
	if hx_matched {
		return hx_recovered
	}
	return false
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	}
	hx_recovered, hx_matched := hx_is_nominal_1__probe_0(value)
	if hx_matched {
		return hx_recovered
	}
	return false
}

func hx_is_nominal_2(value any) bool {
	value = hxrt.ObjectSource(value)
	hx_recovered, hx_matched := hx_is_nominal_2__probe_0(value)
	if hx_matched {
		return hx_recovered
	}
	return false
}

func hx_is_nominal_3(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
