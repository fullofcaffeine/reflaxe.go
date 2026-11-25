package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"snapshot/hxrt"
)

type I_DecodedHex interface {
}

type DecodedHex struct {
	__hx_this I_DecodedHex
	value1    []byte
	value2    *hxrt.NativeErrorValue
}

func (self *DecodedHex) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_DecodedHex(value1 []byte, value2 *hxrt.NativeErrorValue) *DecodedHex {
	return hx_init_DecodedHex(hx_alloc_DecodedHex(nil), value1, value2)
}

func hx_alloc_DecodedHex(owner any) *DecodedHex {
	self := &DecodedHex{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_DecodedHex)
	}
	return self
}

func hx_init_DecodedHex(self *DecodedHex, value1 []byte, value2 *hxrt.NativeErrorValue) *DecodedHex {
	self.value1 = value1
	self.value2 = value2
	return self
}

func (self *DecodedHex) HxClassName() string {
	return "DecodedHex"
}

func check(ok bool, label *string) {
	if !ok {
		hxrt.Throw(haxe__Exception_thrown(label))
	}
}

func compactData(source []byte) *string {
	output := bytes.NewBufferString(*hxrt.StdString(hxrt.StringFromLiteral("")))
	error_local := hxrt.ImportNativeError(json.Compact(output, source))
	if error_local != nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StdString(error_local.Text())))
	}
	return hxrt.StdString(output.String())
}

func keepInts(value []int) int {
	return len(value)
}

func main() {
	check(hxrt.StringEqualStringPtr(compactData(bytes.NewBufferString(*hxrt.StdString(hxrt.StringFromLiteral("{ \"ok\": true }"))).Bytes()), hxrt.StringFromLiteral("{\"ok\":true}")), hxrt.StringFromLiteral("typed []byte parameter"))
	var absent []byte = nil
	check(((absent == nil) && (bytes.Clone(absent) == nil)), hxrt.StringFromLiteral("nil byte slice"))
	decoded := func() *DecodedHex {
		hx_tuple_1, hx_tuple_2 := hex.DecodeString(*hxrt.StdString(hxrt.StringFromLiteral("00ff")))
		return New_DecodedHex(hx_tuple_1, hxrt.ImportNativeError(hx_tuple_2))
	}()
	check((decoded.value2 == nil), hxrt.StringFromLiteral("hex decode"))
	octets := decoded.value1
	check(((len(octets) == 2) && hxrt.StringEqualStringPtr(hxrt.StdString(hex.EncodeToString(octets)), hxrt.StringFromLiteral("00ff"))), hxrt.StringFromLiteral("byte 0 and 255"))
	high := octets[1]
	check((bytes.IndexByte(octets, high) == 1), hxrt.StringFromLiteral("byte scalar argument"))
	buffer := bytes.NewBufferString(*hxrt.StdString(hxrt.StringFromLiteral("ab")))
	view := buffer.Bytes()
	view[0] = high
	check(hxrt.StringEqualStringPtr(hxrt.StdString(hex.EncodeToString(buffer.Bytes())), hxrt.StringFromLiteral("ff62")), hxrt.StringFromLiteral("Buffer.Bytes aliases backing storage"))
	copy_local := bytes.Clone(view)
	copy_local[1] = octets[0]
	check((hxrt.StringEqualStringPtr(hxrt.StdString(hex.EncodeToString(copy_local)), hxrt.StringFromLiteral("ff00")) && hxrt.StringEqualStringPtr(hxrt.StdString(hex.EncodeToString(buffer.Bytes())), hxrt.StringFromLiteral("ff62"))), hxrt.StringFromLiteral("Clone owns a copy"))
	portable := hxrt.ArrayFromValues(func(hx_sort_src_3 []byte) []any {
		hx_sort_out_5 := make([]any, 0, len(hx_sort_src_3))
		for _, hx_sort_item_4 := range hx_sort_src_3 {
			hx_sort_out_5 = append(hx_sort_out_5, hx_sort_item_4)
		}
		return hx_sort_out_5
	}(view))
	copiedBack := func(hx_lambda_raw_6 []any) []byte {
		hx_lambda_out_7 := make([]byte, 0, len(hx_lambda_raw_6))
		for _, hx_lambda_item_8 := range hx_lambda_raw_6 {
			hx_lambda_out_7 = append(hx_lambda_out_7, func(hx_value_9 any) byte {
				if hx_value_9 == nil {
					var hx_zero_10 byte
					return hx_zero_10
				}
				return hx_value_9.(byte)
			}(hx_lambda_item_8))
		}
		return hx_lambda_out_7
	}(portable.Values())
	copiedBack[0] = octets[0]
	check((hxrt.StringEqualStringPtr(hxrt.StdString(hex.EncodeToString(copiedBack)), hxrt.StringFromLiteral("0062")) && hxrt.StringEqualStringPtr(hxrt.StdString(hex.EncodeToString(buffer.Bytes())), hxrt.StringFromLiteral("ff62"))), hxrt.StringFromLiteral("Array boundary copies"))
	check((keepInts([]int{1, 2}) == 2), hxrt.StringFromLiteral("ordinary NativeSlice<Int>"))
	hxrt.Println(any(hxrt.StringFromLiteral("native-byte:ok")))
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
