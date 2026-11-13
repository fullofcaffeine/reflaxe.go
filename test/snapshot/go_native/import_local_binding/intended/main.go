package main

import (
	"bytes"
	"snapshot/hxrt"
	"strings"
)

type I_Box interface {
}

type Box struct {
	__hx_this I_Box
	value     any
}

func (self *Box) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Box(value any) *Box {
	return hx_init_Box(hx_alloc_Box(nil), value)
}

func hx_alloc_Box(owner any) *Box {
	self := &Box{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Box)
	}
	return self
}

func hx_init_Box(self *Box, value any) *Box {
	self.value = value
	return self
}

func main() {
	bytes_local := New_Box(bytes.NewBufferString(*hxrt.StdString(hxrt.StringFromLiteral("abc"))))
	if func(hx_value_1 any) *bytes.Buffer {
		if hx_value_1 == nil {
			var hx_zero_2 *bytes.Buffer
			return hx_zero_2
		}
		return hx_value_1.(*bytes.Buffer)
	}(bytes_local.value).Len() != 3 {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("generic native recovery")))
	}
	if !strings.Contains(*hxrt.StdString(hxrt.StringFromLiteral("abc")), *hxrt.StdString(hxrt.StringFromLiteral("b"))) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("real second import")))
	}
	if (((Local_read() != 7) || (Local_argument(New_Local()) != 7)) || (Local_loops() != 6)) || !hxrt.StringEqualStringPtr(Local_catchValue(), hxrt.StringFromLiteral("caught")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("local values")))
	}
	if (Local_closure() != 73) || (Local_runtimeAlias() != 10) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("nested or runtime binding")))
	}
	if Local_suffix() != 15 {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("suffix collision")))
	}
	hxrt.Println(any(hxrt.StringFromLiteral("import binding: PASS")))
}

func hx_recover_class_0(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_1(value any) *haxe__ValueException {
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
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}
