package main

import (
	"archive/zip"
	"image/color"
	"io/fs"
	"math"
	"math/bits"
	"slices"
	"snapshot/hxrt"
)

type I_Add32Result interface {
}

type Add32Result struct {
	__hx_this I_Add32Result
	sum       uint32
	carryOut  uint32
}

func (self *Add32Result) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Add32Result(sum uint32, carryOut uint32) *Add32Result {
	return hx_init_Add32Result(hx_alloc_Add32Result(nil), sum, carryOut)
}

func hx_alloc_Add32Result(owner any) *Add32Result {
	self := &Add32Result{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Add32Result)
	}
	return self
}

func hx_init_Add32Result(self *Add32Result, sum uint32, carryOut uint32) *Add32Result {
	self.sum = sum
	self.carryOut = carryOut
	return self
}

func (self *Add32Result) HxClassName() string {
	return "Add32Result"
}

type I_Counters interface {
}

type Counters struct {
	__hx_this I_Counters
	wide      uint32
	narrow    uint16
	word      uint
}

func (self *Counters) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Counters(wide uint32, narrow uint16, word uint) *Counters {
	return hx_init_Counters(hx_alloc_Counters(nil), wide, narrow, word)
}

func hx_alloc_Counters(owner any) *Counters {
	self := &Counters{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Counters)
	}
	return self
}

func hx_init_Counters(self *Counters, wide uint32, narrow uint16, word uint) *Counters {
	self.wide = wide
	self.narrow = narrow
	self.word = word
	return self
}

func (self *Counters) HxClassName() string {
	return "Counters"
}

func check(ok bool, label *string) {
	if !ok {
		hxrt.Throw(haxe__Exception_thrown(label))
	}
}

func main() {
	perm := uint32(420)
	check((bits.OnesCount32(perm) == 4), hxrt.StringFromLiteral("uint32 parameter"))
	reversed := bits.ReverseBytes32(perm)
	check(sameUint32(reversed, ^uint32(1543438335)), hxrt.StringFromLiteral("uint32 result and negative constant"))
	check(SyscallOpen_openReadOnly(hxrt.StringFromLiteral("go.mod"), perm), hxrt.StringFromLiteral("syscall.Open uint32 perm"))
	minusOne := int((hxrt.Int32Wrap(bits.OnesCount32(perm)) - hxrt.Int32Wrap(5)))
	allBits := ^uint32(0)
	check(sameUint32(allBits, uint32(minusOne)), hxrt.StringFromLiteral("uint32 wrap"))
	sum := func() *Add32Result {
		hx_tuple_1, hx_tuple_2 := bits.Add32(allBits, uint32(1), uint32(0))
		return New_Add32Result(hx_tuple_1, hx_tuple_2)
	}()
	check(((int(sum.sum) == 0) && (int(sum.carryOut) == 1)), hxrt.StringFromLiteral("uint32 tuple fields"))
	header := &zip.FileHeader{}
	header.SetMode(*func() *fs.FileMode {
		hx_extern_value_3 := fs.FileMode(perm)
		return &hx_extern_value_3
	}())
	permText := hxrt.StdString(func() *fs.FileMode {
		hx_extern_value_5 := func() *fs.FileMode {
			hx_extern_value_4 := header.Mode()
			return &hx_extern_value_4
		}().Perm()
		return &hx_extern_value_5
	}().String())
	check((int(header.ExternalAttrs) != 0), hxrt.StringFromLiteral("uint32 struct field read"))
	header.ExternalAttrs = perm
	check(sameUint32(header.ExternalAttrs, perm), hxrt.StringFromLiteral("uint32 struct field write"))
	header.Method = uint16(8)
	check((int(header.Method) == 8), hxrt.StringFromLiteral("uint16 struct field"))
	check(sameUint16(bits.ReverseBytes16(uint16(4660)), uint16(13330)), hxrt.StringFromLiteral("uint16 parameter"))
	check(sameUint16(uint16(9029), uint16(9029)), hxrt.StringFromLiteral("uint16 constant masking"))
	check(sameUint16(uint16(65535), uint16(minusOne)), hxrt.StringFromLiteral("uint16 wrap"))
	color_local := &color.RGBA64{}
	color_local.R = uint16(65535)
	color_local.A = uint16(65535)
	rgba := func() *RgbaResult {
		hx_tuple_6, hx_tuple_7, hx_tuple_8, hx_tuple_9 := color_local.RGBA()
		return New_RgbaResult(hx_tuple_6, hx_tuple_7, hx_tuple_8, hx_tuple_9)
	}()
	check((((int(rgba.r) == 65535) && (int(rgba.g) == 0)) && sameUint32(rgba.a, rgba.r)), hxrt.StringFromLiteral("uint32 method tuple"))
	check(((bits.OnesCount(uint(255)) == 8) && (bits.Len(uint(256)) == 9)), hxrt.StringFromLiteral("uint parameter"))
	allWord := ^uint(0)
	check((sameUint(allWord, uint(minusOne)) && (bits.OnesCount(allWord) == bits.Len(allWord))), hxrt.StringFromLiteral("uint wrap"))
	check((int(allWord) == -1), hxrt.StringFromLiteral("uint to int"))
	maxWide := uint32(math.MaxUint32)
	check((sameUint32(maxWide, allBits) && (bits.OnesCount32(maxWide) == 32)), hxrt.StringFromLiteral("uint32 constant local"))
	maxNarrow := uint16(math.MaxUint16)
	check(sameUint16(maxNarrow, uint16(65535)), hxrt.StringFromLiteral("uint16 constant local"))
	maxWords := func() []uint {
		_g := hxrt.NewArray()
		_g.Push(uint(math.MaxUint))
		_g.Push(uint(math.MaxUint))
		return func(hx_lambda_raw_12 []any) []uint {
			hx_lambda_out_13 := make([]uint, 0, len(hx_lambda_raw_12))
			for _, hx_lambda_item_14 := range hx_lambda_raw_12 {
				hx_lambda_out_13 = append(hx_lambda_out_13, func(hx_value_15 any) uint {
					if hx_value_15 == nil {
						var hx_zero_16 uint
						return hx_zero_16
					}
					return hx_value_15.(uint)
				}(hx_lambda_item_14))
			}
			return hx_lambda_out_13
		}(_g.Values())
	}()
	check(sameUint(maxWords[1], allWord), hxrt.StringFromLiteral("uint constant slice element"))
	counters := New_Counters(perm, uint16(7), uint(9))
	counters.wide = reversed
	counters.narrow = bits.ReverseBytes16(counters.narrow)
	check((sameUint32(counters.wide, reversed) && (int(counters.narrow) == 1792)), hxrt.StringFromLiteral("class fields"))
	check((int(counters.word) == 9), hxrt.StringFromLiteral("uint class field"))
	wide := []uint32{perm, uint32(7), reversed}
	check(sameUint32(slices.Max(wide), reversed), hxrt.StringFromLiteral("uint32 slice"))
	wide[1] = allBits
	check(sameUint32(slices.Max(wide), allBits), hxrt.StringFromLiteral("uint32 slice element write"))
	narrow := []uint16{uint16(1), uint16(2)}
	check((slices.Index(narrow, uint16(2)) == 1), hxrt.StringFromLiteral("uint16 slice"))
	words := []uint{uint(3), allWord}
	check(sameUint(slices.Max(words), allWord), hxrt.StringFromLiteral("uint slice"))
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("perm: "), permText)))
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("ones: "), bits.OnesCount32(perm)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("round trip: "), int(uint32(42))), hxrt.StringFromLiteral(" ")), int(uint16(42))), hxrt.StringFromLiteral(" ")), int(uint(42))))
	hxrt.Println(v_1)
	hxrt.Println(any(hxrt.StringFromLiteral("native-unsigned:ok")))
}

func sameUint(left uint, right uint) bool {
	return (left == right)
}

func sameUint16(left uint16, right uint16) bool {
	return (left == right)
}

func sameUint32(left uint32, right uint32) bool {
	return (left == right)
}

type I_RgbaResult interface {
}

type RgbaResult struct {
	__hx_this I_RgbaResult
	r         uint32
	g         uint32
	b         uint32
	a         uint32
}

func (self *RgbaResult) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_RgbaResult(r uint32, g uint32, b uint32, a uint32) *RgbaResult {
	return hx_init_RgbaResult(hx_alloc_RgbaResult(nil), r, g, b, a)
}

func hx_alloc_RgbaResult(owner any) *RgbaResult {
	self := &RgbaResult{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_RgbaResult)
	}
	return self
}

func hx_init_RgbaResult(self *RgbaResult, r uint32, g uint32, b uint32, a uint32) *RgbaResult {
	self.r = r
	self.g = g
	self.b = b
	self.a = a
	return self
}

func (self *RgbaResult) HxClassName() string {
	return "RgbaResult"
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
