package main

import (
	"os"
	"runtime"
	"snapshot/hxrt"
)

type I_CallerResult interface {
}

type CallerResult struct {
	__hx_this I_CallerResult
	pc        uintptr
	file      *string
	line      int
	ok        bool
}

func (self *CallerResult) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_CallerResult(pc uintptr, file *string, line int, ok bool) *CallerResult {
	return hx_init_CallerResult(hx_alloc_CallerResult(nil), pc, file, line, ok)
}

func hx_alloc_CallerResult(owner any) *CallerResult {
	self := &CallerResult{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_CallerResult)
	}
	return self
}

func hx_init_CallerResult(self *CallerResult, pc uintptr, file *string, line int, ok bool) *CallerResult {
	self.pc = pc
	self.file = file
	self.line = line
	self.ok = ok
	return self
}

func (self *CallerResult) HxClassName() string {
	return "CallerResult"
}

type I_DescriptorHolder interface {
}

type DescriptorHolder struct {
	__hx_this I_DescriptorHolder
	fd        uintptr
}

func (self *DescriptorHolder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_DescriptorHolder(fd uintptr) *DescriptorHolder {
	return hx_init_DescriptorHolder(hx_alloc_DescriptorHolder(nil), fd)
}

func hx_alloc_DescriptorHolder(owner any) *DescriptorHolder {
	self := &DescriptorHolder{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_DescriptorHolder)
	}
	return self
}

func hx_init_DescriptorHolder(self *DescriptorHolder, fd uintptr) *DescriptorHolder {
	self.fd = fd
	return self
}

func (self *DescriptorHolder) HxClassName() string {
	return "DescriptorHolder"
}

func check(ok bool, label *string) {
	if !ok {
		hxrt.Throw(haxe__Exception_thrown(label))
	}
}

func hasFunction(pc uintptr) bool {
	fn := runtime.FuncForPC(pc)
	return ((fn != nil) && !hxrt.StringEqualStringPtr(hxrt.StdString(fn.Name()), hxrt.StringFromLiteral("")))
}

func main() {
	stdoutFd := os.Stdout.Fd()
	fdNumber := int(stdoutFd)
	check(sameValue(stdoutFd, os.Stdout.Fd()), hxrt.StringFromLiteral("parameter"))
	holder := New_DescriptorHolder(stdoutFd)
	check(sameValue(holder.fd, stdoutFd), hxrt.StringFromLiteral("class field read"))
	holder.fd = uintptr(int((hxrt.Int32Wrap(fdNumber) + hxrt.Int32Wrap(41))))
	check((int(holder.fd) == int((hxrt.Int32Wrap(fdNumber) + hxrt.Int32Wrap(41)))), hxrt.StringFromLiteral("class field write"))
	caller := func() *CallerResult {
		hx_tuple_1, hx_tuple_2, hx_tuple_3, hx_tuple_4 := runtime.Caller(0)
		return New_CallerResult(hx_tuple_1, hxrt.StdString(hx_tuple_2), hx_tuple_3, hx_tuple_4)
	}()
	check(((caller.ok && (caller.line > 0)) && !hxrt.StringEqualStringPtr(caller.file, hxrt.StringFromLiteral(""))), hxrt.StringFromLiteral("runtime.Caller"))
	check(hasFunction(caller.pc), hxrt.StringFromLiteral("tuple field to uintptr parameter"))
	frame := &runtime.Frame{}
	frame.PC = caller.pc
	frame.Entry = runtime.FuncForPC(frame.PC).Entry()
	check((sameValue(frame.PC, caller.pc) && (int(frame.Entry) != 0)), hxrt.StringFromLiteral("struct field"))
	pcs := func() []uintptr {
		_g := hxrt.NewArray()
		_g.Push(uintptr(0))
		_g.Push(uintptr(0))
		_g.Push(uintptr(0))
		_g.Push(uintptr(0))
		return func(hx_lambda_raw_9 []any) []uintptr {
			hx_lambda_out_10 := make([]uintptr, 0, len(hx_lambda_raw_9))
			for _, hx_lambda_item_11 := range hx_lambda_raw_9 {
				hx_lambda_out_10 = append(hx_lambda_out_10, func(hx_value_12 any) uintptr {
					if hx_value_12 == nil {
						var hx_zero_13 uintptr
						return hx_zero_13
					}
					return hx_value_12.(uintptr)
				}(hx_lambda_item_11))
			}
			return hx_lambda_out_10
		}(_g.Values())
	}()
	count := runtime.Callers(1, pcs)
	pcs[3] = pcs[0]
	check(((count > 0) && hasFunction(pcs[3])), hxrt.StringFromLiteral("slice element"))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("stdout fd: "), fdNumber)))
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("round trip: "), int(uintptr(42))))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.StringFromLiteral("native-uintptr:ok")))
}

func sameValue(left uintptr, right uintptr) bool {
	return (left == right)
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
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

func hx_is_nominal_haxe__Exception(value any) bool {
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

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
