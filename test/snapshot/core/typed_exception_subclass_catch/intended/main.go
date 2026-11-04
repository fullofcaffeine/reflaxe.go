package main

import "snapshot/hxrt"

type I_DomainError interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type DomainError struct {
	*haxe__Exception
	__hx_this I_DomainError
	code      *string
}

func (self *DomainError) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_DomainError(code *string, message *string) *DomainError {
	return hx_init_DomainError(hx_alloc_DomainError(nil), code, message)
}

func hx_alloc_DomainError(owner any) *DomainError {
	self := &DomainError{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_DomainError)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_DomainError(self *DomainError, code *string, message *string) *DomainError {
	hx_init_haxe__Exception(self.haxe__Exception, message, nil, nil)
	self.code = code
	return self
}

func (self *DomainError) String() string {
	return *self.toString()
}

func main() {
	hxrt.TryCatch(func() {
		raiseDomain()
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_payload_3
		if hx_is_nominal_0(hx_exception_2) {
			var error *DomainError = hx_recover_class_1(hx_exception_2)
			var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed:"), error.code))
			hxrt.Println(v)
		} else {
			var error_1 *haxe__Exception = hx_recover_class_0(hx_exception_2)
			var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("base:"), error_1.__hx_this.get_message()))
			hxrt.Println(v_1)
		}
	})
	hxrt.TryCatch(func() {
		raisePlain()
	}, func(hx_caught_4 any) {
		hx_exception_5 := haxe__Exception_caught(hx_caught_4)
		hx_payload_6 := haxe__Exception_catchValue(hx_caught_4)
		_ = hx_payload_6
		if hx_is_nominal_0(hx_exception_5) {
			var error_2 *DomainError = hx_recover_class_1(hx_exception_5)
			var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed:"), error_2.code))
			hxrt.Println(v_2)
		} else {
			var error_3 *haxe__Exception = hx_recover_class_0(hx_exception_5)
			var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("base:"), error_3.__hx_this.get_message()))
			hxrt.Println(v_3)
		}
	})
}

func raiseDomain() {
	hxrt.Throw(haxe__Exception_thrown(New_DomainError(hxrt.StringFromLiteral("E42"), hxrt.StringFromLiteral("typed domain failure"))))
}

func raisePlain() {
	hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("plain")))
}

func hx_recover_class_0(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *DomainError:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_1(value any) *DomainError {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *DomainError:
		return source
	default:
		return value.(*DomainError)
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *DomainError:
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
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *DomainError:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}
