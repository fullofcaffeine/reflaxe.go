package main

import (
	"encoding/base64"
	"math"
	"reflect"
	"regexp"
	"snapshot/hxrt"
	"sort"
	"strconv"
	"strings"
	"unsafe"
)

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func main() {
	basic := New_haxe__Template(hxrt.StringFromLiteral("::name::"))
	var v any = any(basic.execute(func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["name"] = hxrt.StringFromLiteral("ok")
		return hx_obj_1
	}(), nil))
	hxrt.Println(v)
	cond := New_haxe__Template(hxrt.StringFromLiteral("::if enabled::yes::else::no::end::"))
	var v_1 any = any(cond.execute(func() map[string]any {
		hx_obj_2 := map[string]any{}
		hx_obj_2["enabled"] = true
		return hx_obj_2
	}(), nil))
	hxrt.Println(v_1)
	var v_2 any = any(cond.execute(func() map[string]any {
		hx_obj_3 := map[string]any{}
		hx_obj_3["enabled"] = false
		return hx_obj_3
	}(), nil))
	hxrt.Println(v_2)
	loop := New_haxe__Template(hxrt.StringFromLiteral("::foreach items::::__current__::::end::"))
	var v_3 any = any(loop.execute(func() map[string]any {
		hx_obj_4 := map[string]any{}
		hx_obj_4["items"] = hxrt.NewArray(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("c"))
		return hx_obj_4
	}(), nil))
	hxrt.Println(v_3)
}

type haxe__io__Encoding struct {
	tag int
}

type haxe__io__Input interface {
	get_bigEndian() bool
	set_bigEndian(e bool) bool
	readByte() int
	readBytes(buf *haxe__io__Bytes, pos int, len int) int
	close()
}

type haxe__io__Output interface {
	get_bigEndian() bool
	set_bigEndian(e bool) bool
	writeByte(c int)
	writeBytes(s *haxe__io__Bytes, pos int, len int) int
	flush()
	close()
}

type haxe__io__Eof struct {
}

type haxe__io__Error struct {
	tag    int
	params []any
}

type haxe__io__Bytes struct {
	b             []int
	length        int
	__hx_raw      []byte
	__hx_rawValid bool
}

type haxe__io__BytesBuffer struct {
	b []int
}

type haxe__io__BytesInput struct {
	bigEndian bool
	b         []int
	pos       int
	len       int
	totlen    int
}

type haxe__io__BytesOutput struct {
	bigEndian bool
	b         *haxe__io__BytesBuffer
}

func New_haxe__io__Input() haxe__io__Input {
	return New_haxe__io__BytesInput(&haxe__io__Bytes{b: []int{}, length: 0})
}

func New_haxe__io__Output() haxe__io__Output {
	return New_haxe__io__BytesOutput()
}

func New_haxe__io__Eof() *haxe__io__Eof {
	return &haxe__io__Eof{}
}

var haxe__io__Encoding_UTF8 *haxe__io__Encoding = &haxe__io__Encoding{tag: 0}

var haxe__io__Encoding_RawNative *haxe__io__Encoding = &haxe__io__Encoding{tag: 1}

func (self *haxe__io__Encoding) String() string {
	if self == nil {
		return "null"
	}
	switch self.tag {
	case 0:
		return "UTF8"
	case 1:
		return "RawNative"
	default:
		return "Encoding"
	}
}

func (self *haxe__io__Encoding) toString() *string {
	return hxrt.StringFromLiteral(self.String())
}

func haxe__io__resolveEncodingTag(encoding ...*haxe__io__Encoding) int {
	if len(encoding) == 0 || encoding[0] == nil {
		return 0
	}
	return encoding[0].tag
}

func haxe__io__bytes_fromStringRawNativeUTF16LE(value *string) *haxe__io__Bytes {
	raw := []byte(*hxrt.StdString(value))
	converted := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		converted[i] = int(raw[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}
}

func haxe__io__bytes_toStringRawNativeUTF16LE(value []int) *string {
	return hxrt.BytesToString(value)
}

func (self *haxe__io__Eof) toString() *string {
	return hxrt.StringFromLiteral("Eof")
}

var haxe__io__Error_Blocked *haxe__io__Error = &haxe__io__Error{tag: 0}

var haxe__io__Error_Overflow *haxe__io__Error = &haxe__io__Error{tag: 1}

var haxe__io__Error_OutsideBounds *haxe__io__Error = &haxe__io__Error{tag: 2}

func haxe__io__Error_Custom(e any) *haxe__io__Error {
	return &haxe__io__Error{tag: 3, params: []any{e}}
}

func (self *haxe__io__Error) String() string {
	if self == nil {
		return "null"
	}
	switch self.tag {
	case 0:
		return "Blocked"
	case 1:
		return "Overflow"
	case 2:
		return "OutsideBounds"
	case 3:
		if len(self.params) == 0 {
			return "Custom(null)"
		}
		return "Custom(" + *hxrt.StdString(self.params[0]) + ")"
	default:
		return "Error"
	}
}

func (self *haxe__io__Error) toString() *string {
	return hxrt.StringFromLiteral(self.String())
}

func New_haxe__io__Bytes(length int, b []int) *haxe__io__Bytes {
	if b == nil {
		b = make([]int, length)
	}
	return &haxe__io__Bytes{b: b, length: len(b)}
}

func haxe__io__Bytes_alloc(length int) *haxe__io__Bytes {
	return &haxe__io__Bytes{b: make([]int, length), length: length}
}

func haxe__io__Bytes_ofString(value *string, encoding ...*haxe__io__Encoding) *haxe__io__Bytes {
	if haxe__io__resolveEncodingTag(encoding...) == 1 {
		return haxe__io__bytes_fromStringRawNativeUTF16LE(value)
	}
	raw := []byte(*hxrt.StdString(value))
	converted := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		converted[i] = int(raw[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}
}

func haxe__io__Bytes_ofData(b []int) *haxe__io__Bytes {
	if b == nil {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return &haxe__io__Bytes{b: b, length: len(b)}
}

func haxe__io__Bytes_ofHex(s *string) *haxe__io__Bytes {
	decoded := hxrt.BytesOfHex(s)
	return &haxe__io__Bytes{b: decoded, length: len(decoded)}
}

func (self *haxe__io__Bytes) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.BytesToString(self.b)
}

func (self *haxe__io__Bytes) toHex() *string {
	if self == nil || self.length == 0 {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.BytesToHex(self.b, self.length)
}

func (self *haxe__io__Bytes) getData() []int {
	if self == nil {
		return []int{}
	}
	return self.b
}

func (self *haxe__io__Bytes) getString(pos int, len int, encoding ...*haxe__io__Encoding) *string {
	if self == nil || pos < 0 || len < 0 || pos+len > self.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return hxrt.StringFromLiteral("")
	}
	slice := self.b[pos : pos+len]
	if haxe__io__resolveEncodingTag(encoding...) == 1 {
		return haxe__io__bytes_toStringRawNativeUTF16LE(slice)
	}
	return hxrt.BytesToString(slice)
}

func (self *haxe__io__Bytes) readString(pos int, len int) *string {
	return self.getString(pos, len)
}

func (self *haxe__io__Bytes) get(pos int) int {
	return self.b[pos]
}

func (self *haxe__io__Bytes) set(pos int, value int) {
	self.b[pos] = value
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) blit(pos int, src *haxe__io__Bytes, srcpos int, len int) {
	if self == nil || src == nil || pos < 0 || srcpos < 0 || len < 0 || pos+len > self.length || srcpos+len > src.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	if len == 0 {
		return
	}
	if self == src && pos > srcpos {
		for i := len - 1; i >= 0; i-- {
			self.b[pos+i] = src.b[srcpos+i]
		}
	} else {
		for i := 0; i < len; i++ {
			self.b[pos+i] = src.b[srcpos+i]
		}
	}
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) fill(pos int, len int, value int) {
	if self == nil || pos < 0 || len < 0 || pos+len > self.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	masked := value & 255
	for i := 0; i < len; i++ {
		self.b[pos+i] = masked
	}
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) sub(pos int, len int) *haxe__io__Bytes {
	if self == nil || pos < 0 || len < 0 || pos+len > self.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	if len == 0 {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	copied := make([]int, len)
	copy(copied, self.b[pos:pos+len])
	return &haxe__io__Bytes{b: copied, length: len}
}

func (self *haxe__io__Bytes) compare(other *haxe__io__Bytes) int {
	if self == nil && other == nil {
		return 0
	}
	if self == nil {
		return -1
	}
	if other == nil {
		return 1
	}
	limit := self.length
	if other.length < limit {
		limit = other.length
	}
	for i := 0; i < limit; i++ {
		if self.b[i] < other.b[i] {
			return -1
		}
		if self.b[i] > other.b[i] {
			return 1
		}
	}
	if self.length < other.length {
		return -1
	}
	if self.length > other.length {
		return 1
	}
	return 0
}

func New_haxe__io__BytesBuffer() *haxe__io__BytesBuffer {
	return &haxe__io__BytesBuffer{b: []int{}}
}

func (self *haxe__io__BytesBuffer) addByte(value int) {
	self.b = hxrt.BytesBufferAddByte(self.b, value)
}

func (self *haxe__io__BytesBuffer) add(src *haxe__io__Bytes) {
	if src == nil {
		return
	}
	self.b = hxrt.BytesBufferAdd(self.b, src.b)
}

func (self *haxe__io__BytesBuffer) addBytes(src *haxe__io__Bytes, pos int, len int) {
	if src == nil || pos < 0 || len < 0 || pos+len > src.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	if len == 0 {
		return
	}
	self.b = hxrt.BytesBufferAddSlice(self.b, src.b, pos, len)
}

func (self *haxe__io__BytesBuffer) addString(value *string, encoding ...*haxe__io__Encoding) {
	self.add(haxe__io__Bytes_ofString(value))
}

func (self *haxe__io__BytesBuffer) getBytes() *haxe__io__Bytes {
	copied := hxrt.BytesClone(self.b)
	return &haxe__io__Bytes{b: copied, length: len(copied)}
}

func (self *haxe__io__BytesBuffer) get_length() int {
	return hxrt.BytesBufferLength(self.b)
}

func New_haxe__io__BytesInput(b *haxe__io__Bytes, opts ...int) *haxe__io__BytesInput {
	if b == nil {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__BytesInput{}
	}
	start := 0
	if len(opts) > 0 {
		start = opts[0]
	}
	sliceLen := (b.length - start)
	if len(opts) > 1 {
		sliceLen = opts[1]
	}
	if start < 0 || sliceLen < 0 || start+sliceLen > b.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__BytesInput{}
	}
	return &haxe__io__BytesInput{b: b.b, pos: start, len: sliceLen, totlen: sliceLen}
}

func (self *haxe__io__BytesInput) get_position() int {
	return self.pos
}

func (self *haxe__io__BytesInput) set_position(p int) int {
	if p < 0 {
		p = 0
	} else {
		if p > self.totlen {
			p = self.totlen
		}
	}
	self.len = (self.totlen - p)
	self.pos = p
	return p
}

func (self *haxe__io__BytesInput) get_length() int {
	return self.totlen
}

func (self *haxe__io__BytesInput) readByte() int {
	if self == nil || self.len == 0 {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	self.len = (self.len - 1)
	value := self.b[self.pos]
	self.pos = (self.pos + 1)
	return value
}

func (self *haxe__io__BytesInput) readBytes(buf *haxe__io__Bytes, pos int, len int) int {
	if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	if len > 0 && (self == nil || self.len == 0) {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	if self == nil {
		return 0
	}
	if self.len < len {
		len = self.len
	}
	for i := 0; i < len; i++ {
		buf.b[pos+i] = self.b[self.pos+i]
	}
	self.pos += len
	self.len -= len
	return len
}

func (self *haxe__io__BytesInput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.bigEndian
}

func (self *haxe__io__BytesInput) set_bigEndian(e bool) bool {
	if self != nil {
		self.bigEndian = e
	}
	return e
}

func (self *haxe__io__BytesInput) close() {
	_ = self
}

func New_haxe__io__BytesOutput() *haxe__io__BytesOutput {
	return &haxe__io__BytesOutput{b: &haxe__io__BytesBuffer{b: []int{}}}
}

func (self *haxe__io__BytesOutput) get_length() int {
	if self == nil || self.b == nil {
		return 0
	}
	return self.b.get_length()
}

func (self *haxe__io__BytesOutput) writeByte(c int) {
	if self == nil || self.b == nil {
		return
	}
	self.b.addByte(c)
}

func (self *haxe__io__BytesOutput) writeBytes(buf *haxe__io__Bytes, pos int, len int) int {
	if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	if self == nil || self.b == nil {
		return 0
	}
	self.b.addBytes(buf, pos, len)
	return len
}

func (self *haxe__io__BytesOutput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.bigEndian
}

func (self *haxe__io__BytesOutput) set_bigEndian(e bool) bool {
	if self != nil {
		self.bigEndian = e
	}
	return e
}

func (self *haxe__io__BytesOutput) flush() {
	_ = self
}

func (self *haxe__io__BytesOutput) close() {
	_ = self
}

func (self *haxe__io__BytesOutput) getBytes() *haxe__io__Bytes {
	if self == nil || self.b == nil {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return self.b.getBytes()
}

type Std struct {
}

type Type struct {
}

type Reflect struct {
}

func Reflect_compare(a any, b any) int {
	toFloat := func(value any) (float64, bool) {
		switch v := value.(type) {
		case int:
			return float64(v), true
		case int8:
			return float64(v), true
		case int16:
			return float64(v), true
		case int32:
			return float64(v), true
		case int64:
			return float64(v), true
		case uint:
			return float64(v), true
		case uint8:
			return float64(v), true
		case uint16:
			return float64(v), true
		case uint32:
			return float64(v), true
		case uint64:
			return float64(v), true
		case float32:
			return float64(v), true
		case float64:
			return v, true
		default:
			return 0, false
		}
	}
	if af, ok := toFloat(a); ok {
		if bf, okB := toFloat(b); okB {
			if af < bf {
				return -1
			}
			if af > bf {
				return 1
			}
			return 0
		}
	}
	aStr := *hxrt.StdString(a)
	bStr := *hxrt.StdString(b)
	if aStr < bStr {
		return -1
	}
	if aStr > bStr {
		return 1
	}
	return 0
}

func Reflect_compareMethods(a any, b any) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	if !av.IsValid() || !bv.IsValid() {
		return !av.IsValid() && !bv.IsValid()
	}
	if av.Kind() == reflect.Func && bv.Kind() == reflect.Func {
		if av.IsNil() || bv.IsNil() {
			return av.IsNil() && bv.IsNil()
		}
		return av.Pointer() == bv.Pointer()
	}
	return reflect.DeepEqual(a, b)
}

func Reflect_field(obj any, field *string) any {
	if obj == nil {
		return nil
	}
	key := *hxrt.StdString(field)
	if metadataValue, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return metadataValue
	}
	switch value := obj.(type) {
	case map[string]any:
		return value[key]
	case map[any]any:
		return value[key]
	case *map[string]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	case *map[any]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return nil
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if fieldValue := rv.FieldByName(key); fieldValue.IsValid() && fieldValue.CanInterface() {
			return fieldValue.Interface()
		}
	}
	method := reflect.ValueOf(obj).MethodByName(key)
	if method.IsValid() {
		return method.Interface()
	}
	return nil
}

func Reflect_hasField(obj any, field *string) bool {
	if obj == nil {
		return false
	}
	key := *hxrt.StdString(field)
	if _, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return true
	}
	switch value := obj.(type) {
	case map[string]any:
		_, ok := value[key]
		return ok
	case map[any]any:
		_, ok := value[key]
		return ok
	case *map[string]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	case *map[any]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return false
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if rv.FieldByName(key).IsValid() {
			return true
		}
	}
	return reflect.ValueOf(obj).MethodByName(key).IsValid()
}

func Reflect_setField(obj any, field *string, value any) {
	if obj == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	key := *hxrt.StdString(field)
	switch target := obj.(type) {
	case map[string]any:
		target[key] = value
		return
	case map[any]any:
		target[key] = value
		return
	case *map[string]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	case *map[any]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() || rv.Kind() != reflect.Pointer {
		return
	}
	if rv.IsNil() {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return
	}
	fieldValue := rv.FieldByName(key)
	if !fieldValue.IsValid() || !fieldValue.CanSet() {
		return
	}
	if value == nil {
		fieldValue.Set(reflect.Zero(fieldValue.Type()))
		return
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(fieldValue.Type()) {
		fieldValue.Set(incoming)
		return
	}
	if incoming.Type().ConvertibleTo(fieldValue.Type()) {
		fieldValue.Set(incoming.Convert(fieldValue.Type()))
		return
	}
	if fieldValue.Kind() == reflect.Interface {
		fieldValue.Set(incoming)
	}
}

func hxrt_haxeBytesToRaw(value *haxe__io__Bytes) []byte {
	if value == nil {
		return []byte{}
	}
	if value.__hx_rawValid && len(value.__hx_raw) == len(value.b) {
		return value.__hx_raw
	}
	raw := make([]byte, len(value.b))
	for i := 0; i < len(value.b); i++ {
		raw[i] = byte(value.b[i])
	}
	value.__hx_raw = raw
	value.__hx_rawValid = true
	return raw
}

func hxrt_rawToHaxeBytes(value []byte) *haxe__io__Bytes {
	converted := make([]int, len(value))
	for i := 0; i < len(value); i++ {
		converted[i] = int(value[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: value, __hx_rawValid: true}
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}

type ValueType struct {
	tag    int
	params []any
}

var ValueType_TNull *ValueType = &ValueType{tag: 0, params: []any{}}

var ValueType_TInt *ValueType = &ValueType{tag: 1, params: []any{}}

var ValueType_TFloat *ValueType = &ValueType{tag: 2, params: []any{}}

var ValueType_TBool *ValueType = &ValueType{tag: 3, params: []any{}}

var ValueType_TObject *ValueType = &ValueType{tag: 4, params: []any{}}

var ValueType_TFunction *ValueType = &ValueType{tag: 5, params: []any{}}

var ValueType_TUnknown *ValueType = &ValueType{tag: 8, params: []any{}}

func ValueType_TClass(c any) *ValueType {
	return &ValueType{tag: 6, params: []any{c}}
}

func ValueType_TEnum(e any) *ValueType {
	return &ValueType{tag: 7, params: []any{e}}
}

func hxrt_typeCallAny(callable any, args []any) (any, bool) {
	result := any(nil)
	ok := false
	defer func() {
		if recover() != nil {
			result = nil
			ok = false
		}
	}()
	if callable == nil {
		return nil, false
	}
	fn := reflect.ValueOf(callable)
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil, false
	}
	fnType := fn.Type()
	if fnType.NumIn() != len(args) {
		return nil, false
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		paramType := fnType.In(i)
		arg := args[i]
		if arg == nil {
			in[i] = reflect.Zero(paramType)
			continue
		}
		v := reflect.ValueOf(arg)
		if v.IsValid() && v.Type().AssignableTo(paramType) {
			in[i] = v
			continue
		}
		if v.IsValid() && v.Type().ConvertibleTo(paramType) {
			in[i] = v.Convert(paramType)
			continue
		}
		if paramType.Kind() == reflect.Interface && v.IsValid() {
			in[i] = v
			continue
		}
		return nil, false
	}
	out := fn.Call(in)
	if len(out) == 0 {
		return nil, true
	}
	first := out[0]
	if !first.IsValid() {
		return nil, true
	}
	result = first.Interface()
	ok = true
	return result, ok
}

func hxrt_typeArrayValues(value *hxrt.Array) []any {
	if value == nil {
		return []any{}
	}
	return value.Values()
}

func hxrt_typeResolvedClassName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeClassValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeClassValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeResolvedEnumName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeEnumValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeEnumValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeCreateClassInstance(className string, args []any) (any, bool) {
	switch className {
	case "Main":
		return nil, false
	case "StringTools":
		return nil, false
	case "haxe.Int64Helper":
		return nil, false
	case "haxe.Template":
		return hxrt_typeCallAny(New_haxe__Template, args)
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe._Template.ExprCursor":
		return hxrt_typeCallAny(New_haxe___Template__ExprCursor, args)
	case "haxe._Template.TokenCursor":
		return hxrt_typeCallAny(New_haxe___Template__TokenCursor, args)
	case "haxe.ds.IntMap":
		return hxrt_typeCallAny(New_haxe__ds__IntMap, args)
	case "haxe.ds.List":
		return hxrt_typeCallAny(New_haxe__ds__List, args)
	case "haxe.ds.ObjectMap":
		return hxrt_typeCallAny(New_haxe__ds__ObjectMap, args)
	case "haxe.ds.StringMap":
		return hxrt_typeCallAny(New_haxe__ds__StringMap, args)
	case "haxe.ds._List.GoListIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListIterator, args)
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListKeyValueIterator, args)
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__MapKeyValueIterator, args)
	case "haxe.iterators.StringIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringIterator, args)
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringKeyValueIterator, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "haxe.Template":
		return &haxe__Template{}, true
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
	case "haxe._Template.ExprCursor":
		return &haxe___Template__ExprCursor{}, true
	case "haxe._Template.TokenCursor":
		return &haxe___Template__TokenCursor{}, true
	case "haxe.ds.IntMap":
		return &haxe__ds__IntMap{}, true
	case "haxe.ds.List":
		return &haxe__ds__List{}, true
	case "haxe.ds.ObjectMap":
		return &haxe__ds__ObjectMap{}, true
	case "haxe.ds.StringMap":
		return &haxe__ds__StringMap{}, true
	case "haxe.ds._List.GoListIterator":
		return &haxe__ds___List__GoListIterator{}, true
	case "haxe.ds._List.GoListKeyValueIterator":
		return &haxe__ds___List__GoListKeyValueIterator{}, true
	case "haxe.iterators.MapKeyValueIterator":
		return &haxe__iterators__MapKeyValueIterator{}, true
	case "haxe.iterators.StringIterator":
		return &haxe__iterators__StringIterator{}, true
	case "haxe.iterators.StringKeyValueIterator":
		return &haxe__iterators__StringKeyValueIterator{}, true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
	case "ValueType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TNull, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TInt, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFloat, true
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TBool, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TObject, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFunction, true
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TClass, args)
			case 7:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TEnum, args)
			case 8:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TUnknown, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TNull":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TNull, true
		case "TInt":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TInt, true
		case "TFloat":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFloat, true
		case "TBool":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TBool, true
		case "TObject":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TObject, true
		case "TFunction":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFunction, true
		case "TClass":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TClass, args)
		case "TEnum":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TEnum, args)
		case "TUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TUnknown, true
		default:
			return nil, false
		}
	case "haxe._Template.TemplateExpr":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpVar, args)
			case 1:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpExpr, args)
			case 2:
				if len(args) != 3 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpIf, args)
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpStr, args)
			case 4:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpBlock, args)
			case 5:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpForeach, args)
			case 6:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpMacro, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "OpVar":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpVar, args)
		case "OpExpr":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpExpr, args)
		case "OpIf":
			if len(args) != 3 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpIf, args)
		case "OpStr":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpStr, args)
		case "OpBlock":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpBlock, args)
		case "OpForeach":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpForeach, args)
		case "OpMacro":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe___Template__TemplateExpr_OpMacro, args)
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func Type_getClass(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeClassValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeClassValue:
		copyValue := value
		return &copyValue
	case *hxrt.Array:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")}
	case *haxe__Template:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Template")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	case *haxe___Template__ExprCursor:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Template.ExprCursor")}
	case *haxe___Template__TokenCursor:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Template.TokenCursor")}
	case *haxe__ds__IntMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.IntMap")}
	case *haxe__ds__List:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.List")}
	case *haxe__ds__ObjectMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.ObjectMap")}
	case *haxe__ds__StringMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.StringMap")}
	case *haxe__ds___List__GoListIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._List.GoListIterator")}
	case *haxe__ds___List__GoListKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._List.GoListKeyValueIterator")}
	case *haxe__iterators__MapKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.MapKeyValueIterator")}
	case *haxe__iterators__StringIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.StringIterator")}
	case *haxe__iterators__StringKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.StringKeyValueIterator")}
	default:
		return nil
	}
}

func Type_getEnum(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeEnumValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeEnumValue:
		copyValue := value
		return &copyValue
	case *ValueType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("ValueType")}
	case *haxe___Template__TemplateExpr:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe._Template.TemplateExpr")}
	default:
		return nil
	}
}

func Type_getSuperClass(c any) any {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	switch className {
	case "Main":
		return nil
	case "StringTools":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe.Template":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe._Template.ExprCursor":
		return nil
	case "haxe._Template.TokenCursor":
		return nil
	case "haxe.ds.IntMap":
		return nil
	case "haxe.ds.List":
		return nil
	case "haxe.ds.ObjectMap":
		return nil
	case "haxe.ds.StringMap":
		return nil
	case "haxe.ds._List.GoListIterator":
		return nil
	case "haxe.ds._List.GoListKeyValueIterator":
		return nil
	case "haxe.iterators.MapKeyValueIterator":
		return nil
	case "haxe.iterators.StringIterator":
		return nil
	case "haxe.iterators.StringKeyValueIterator":
		return nil
	default:
		return nil
	}
}

func Type_getClassName(c any) *string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(className)
}

func Type_getClassFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("main"))
	case "StringTools":
		return hxrt.NewArray(hxrt.StringFromLiteral("MAX_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("contains"), hxrt.StringFromLiteral("containsImpl"), hxrt.StringFromLiteral("endsWith"), hxrt.StringFromLiteral("endsWithImpl"), hxrt.StringFromLiteral("fastCodeAt"), hxrt.StringFromLiteral("hex"), hxrt.StringFromLiteral("hexDigitValue"), hxrt.StringFromLiteral("htmlEscape"), hxrt.StringFromLiteral("htmlUnescape"), hxrt.StringFromLiteral("isEof"), hxrt.StringFromLiteral("isSpace"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("lpad"), hxrt.StringFromLiteral("ltrim"), hxrt.StringFromLiteral("replace"), hxrt.StringFromLiteral("rpad"), hxrt.StringFromLiteral("rtrim"), hxrt.StringFromLiteral("startsWith"), hxrt.StringFromLiteral("startsWithImpl"), hxrt.StringFromLiteral("trim"), hxrt.StringFromLiteral("unsafeCodeAt"), hxrt.StringFromLiteral("urlDecode"), hxrt.StringFromLiteral("urlEncode"), hxrt.StringFromLiteral("utf16CodePointAt"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Template":
		return hxrt.NewArray(hxrt.StringFromLiteral("addValues"), hxrt.StringFromLiteral("compareValues"), hxrt.StringFromLiteral("divideValues"), hxrt.StringFromLiteral("expr_float"), hxrt.StringFromLiteral("expr_int"), hxrt.StringFromLiteral("expr_splitter"), hxrt.StringFromLiteral("expr_trim"), hxrt.StringFromLiteral("globals"), hxrt.StringFromLiteral("isSpaceOnly"), hxrt.StringFromLiteral("joinDynamicArgs"), hxrt.StringFromLiteral("kwdEnd"), hxrt.StringFromLiteral("multiplyValues"), hxrt.StringFromLiteral("parseFloatLiteral"), hxrt.StringFromLiteral("parseIntLiteral"), hxrt.StringFromLiteral("peekExprToken"), hxrt.StringFromLiteral("peekToken"), hxrt.StringFromLiteral("popExprToken"), hxrt.StringFromLiteral("popToken"), hxrt.StringFromLiteral("splitter"), hxrt.StringFromLiteral("subtractValues"), hxrt.StringFromLiteral("trimExprToken"), hxrt.StringFromLiteral("valueAsBool"), hxrt.StringFromLiteral("valueAsFloat"))
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe._Template.ExprCursor":
		return hxrt.NewArray()
	case "haxe._Template.TokenCursor":
		return hxrt.NewArray()
	case "haxe.ds.IntMap":
		return hxrt.NewArray()
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("sameValue"))
	case "haxe.ds.ObjectMap":
		return hxrt.NewArray()
	case "haxe.ds.StringMap":
		return hxrt.NewArray()
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray()
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray()
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt.NewArray()
	default:
		return hxrt.NewArray()
	}
}

func Type_getInstanceFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Main":
		return hxrt.NewArray()
	case "StringTools":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Template":
		return hxrt.NewArray(hxrt.StringFromLiteral("context"), hxrt.StringFromLiteral("execute"), hxrt.StringFromLiteral("expr"), hxrt.StringFromLiteral("macros"), hxrt.StringFromLiteral("makeConst"), hxrt.StringFromLiteral("makeExpr"), hxrt.StringFromLiteral("makeExpr2"), hxrt.StringFromLiteral("makePath"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("parse"), hxrt.StringFromLiteral("parseBlock"), hxrt.StringFromLiteral("parseExpr"), hxrt.StringFromLiteral("parseTokens"), hxrt.StringFromLiteral("popStackValue"), hxrt.StringFromLiteral("resolve"), hxrt.StringFromLiteral("run"), hxrt.StringFromLiteral("skipSpaces"), hxrt.StringFromLiteral("stack"))
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe._Template.ExprCursor":
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("tokens"))
	case "haxe._Template.TokenCursor":
		return hxrt.NewArray(hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("tokens"))
	case "haxe.ds.IntMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("filter"), hxrt.StringFromLiteral("first"), hxrt.StringFromLiteral("isEmpty"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("join"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("last"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("pop"), hxrt.StringFromLiteral("push"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds.ObjectMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds.StringMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("next"))
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	default:
		return hxrt.NewArray()
	}
}

func Type_getEnumName(e any) *string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(enumName)
}

func Type_resolveClass(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringTools":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Template":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Template.ExprCursor":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Template.TokenCursor":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.IntMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.List":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.ObjectMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.StringMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.MapKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_resolveEnum(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "ValueType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Template.TemplateExpr":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_createInstance(cl any, args *hxrt.Array) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, hxrt_typeArrayValues(args))
	if !ok {
		return nil
	}
	return instance
}

func Type_createEmptyInstance(cl any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassEmptyInstance(className)
	if !ok {
		return nil
	}
	return instance
}

func Type_createEnum(e any, constr *string, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_enumConstructor(e any) *string {
	if hxrt.AnyEqualsNull(e) {
		return nil
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TNull")
		case 1:
			return hxrt.StringFromLiteral("TInt")
		case 2:
			return hxrt.StringFromLiteral("TFloat")
		case 3:
			return hxrt.StringFromLiteral("TBool")
		case 4:
			return hxrt.StringFromLiteral("TObject")
		case 5:
			return hxrt.StringFromLiteral("TFunction")
		case 6:
			return hxrt.StringFromLiteral("TClass")
		case 7:
			return hxrt.StringFromLiteral("TEnum")
		case 8:
			return hxrt.StringFromLiteral("TUnknown")
		default:
			return nil
		}
	case *haxe___Template__TemplateExpr:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("OpVar")
		case 1:
			return hxrt.StringFromLiteral("OpExpr")
		case 2:
			return hxrt.StringFromLiteral("OpIf")
		case 3:
			return hxrt.StringFromLiteral("OpStr")
		case 4:
			return hxrt.StringFromLiteral("OpBlock")
		case 5:
			return hxrt.StringFromLiteral("OpForeach")
		case 6:
			return hxrt.StringFromLiteral("OpMacro")
		default:
			return nil
		}
	default:
		return nil
	}
}

func Type_enumIndex(e any) int {
	if hxrt.AnyEqualsNull(e) {
		return -1
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe___Template__TemplateExpr:
		if value == nil {
			return -1
		}
		return value.tag
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown"))
	case "haxe._Template.TemplateExpr":
		return hxrt.NewArray(hxrt.StringFromLiteral("OpVar"), hxrt.StringFromLiteral("OpExpr"), hxrt.StringFromLiteral("OpIf"), hxrt.StringFromLiteral("OpStr"), hxrt.StringFromLiteral("OpBlock"), hxrt.StringFromLiteral("OpForeach"), hxrt.StringFromLiteral("OpMacro"))
	default:
		return hxrt.NewArray()
	}
}

func Type_enumParameters(e any) *hxrt.Array {
	if hxrt.AnyEqualsNull(e) {
		return hxrt.NewArray()
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe___Template__TemplateExpr:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	default:
		return hxrt.NewArray()
	}
}

func Type_allEnums(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown)
	case "haxe._Template.TemplateExpr":
		return hxrt.NewArray()
	default:
		return hxrt.NewArray()
	}
}

func Type_typeof(v any) any {
	if hxrt.AnyEqualsNull(v) {
		return ValueType_TNull
	}
	if enumValue := Type_getEnum(v); enumValue != nil {
		return ValueType_TEnum(enumValue)
	}
	if classValue := Type_getClass(v); classValue != nil {
		return ValueType_TClass(classValue)
	}
	switch v.(type) {
	case bool:
		return ValueType_TBool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return ValueType_TInt
	case float32, float64:
		return ValueType_TFloat
	case string, *string:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("String")})
	case *hxrt.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	}
	ref := reflect.ValueOf(v)
	if !ref.IsValid() {
		return ValueType_TNull
	}
	switch ref.Kind() {
	case reflect.Func:
		return ValueType_TFunction
	case reflect.Slice, reflect.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:
		return ValueType_TObject
	default:
		return ValueType_TUnknown
	}
}

func Type_enumEq(a any, b any) bool {
	return reflect.DeepEqual(a, b)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	className, ok := hxrt_typeResolvedClassName(value)
	if !ok {
		return nil, false
	}
	switch className {
	default:
		return nil, false
	}
}

type EReg struct {
	regex       *regexp.Regexp
	global      bool
	lastSource  *string
	lastIndices []int
}

func New_EReg(pattern *string, options *string) *EReg {
	rawPattern := *hxrt.StdString(pattern)
	rawOptions := *hxrt.StdString(options)
	global := false
	flagI := false
	flagM := false
	flagS := false
	for _, option := range rawOptions {
		switch option {
		case 'g':
			global = true
		case 'i':
			flagI = true
		case 'm':
			flagM = true
		case 's':
			flagS = true
		case 'u':
			// RE2 is UTF-8 aware by default; keep parity by accepting and ignoring this option.
		default:
			hxrt.Throw(hxrt.StringFromLiteral("Unsupported regexp option '" + string(option) + "'"))
			return &EReg{regex: regexp.MustCompile("a^"), global: false, lastSource: nil, lastIndices: nil}
		}
	}
	inlineFlags := ""
	if flagI {
		inlineFlags += "i"
	}
	if flagM {
		inlineFlags += "m"
	}
	if flagS {
		inlineFlags += "s"
	}
	if inlineFlags != "" {
		rawPattern = "(?" + inlineFlags + ")" + rawPattern
	}
	compiled, err := regexp.Compile(rawPattern)
	if err != nil {
		hxrt.Throw(err)
		compiled = regexp.MustCompile("a^")
	}
	return &EReg{regex: compiled, global: global, lastSource: nil, lastIndices: nil}
}

func hxrt_eregHasMatch(self *EReg) bool {
	return self != nil && self.lastSource != nil && len(self.lastIndices) >= 2 && self.lastIndices[0] >= 0 && self.lastIndices[1] >= self.lastIndices[0]
}

func hxrt_eregThrowNoMatch() {
	hxrt.Throw(hxrt.StringFromLiteral("Invalid regex operation because no match was made"))
}

func hxrt_eregThrowInvalidGroup() {
	hxrt.Throw(hxrt.StringFromLiteral("Invalid group"))
}

func (self *EReg) match(source *string) bool {
	if self == nil || self.regex == nil {
		return false
	}
	raw := *hxrt.StdString(source)
	found := self.regex.FindStringSubmatchIndex(raw)
	if found == nil {
		self.lastSource = nil
		self.lastIndices = nil
		return false
	}
	indices := make([]int, len(found))
	copy(indices, found)
	self.lastSource = hxrt.StringFromLiteral(raw)
	self.lastIndices = indices
	return true
}

func (self *EReg) matchSub(source *string, pos int) bool {
	if self == nil || self.regex == nil {
		return false
	}
	raw := *hxrt.StdString(source)
	if pos < 0 {
		pos = 0
	}
	if pos > len(raw) {
		return false
	}
	found := self.regex.FindStringSubmatchIndex(raw[pos:])
	if found == nil {
		return false
	}
	shifted := make([]int, len(found))
	for i := 0; i < len(found); i++ {
		if found[i] >= 0 {
			shifted[i] = found[i] + pos
		} else {
			shifted[i] = -1
		}
	}
	self.lastSource = hxrt.StringFromLiteral(raw)
	self.lastIndices = shifted
	return true
}

func (self *EReg) matched(index int) *string {
	if !hxrt_eregHasMatch(self) {
		hxrt_eregThrowNoMatch()
		return nil
	}
	if index < 0 {
		hxrt_eregThrowInvalidGroup()
		return nil
	}
	offset := index * 2
	if offset+1 >= len(self.lastIndices) {
		hxrt_eregThrowInvalidGroup()
		return nil
	}
	start := self.lastIndices[offset]
	end := self.lastIndices[offset+1]
	if start < 0 || end < start {
		return nil
	}
	raw := *hxrt.StdString(self.lastSource)
	if end > len(raw) {
		end = len(raw)
	}
	return hxrt.StringFromLiteral(raw[start:end])
}

func (self *EReg) matchedPos() map[string]any {
	if !hxrt_eregHasMatch(self) {
		hxrt_eregThrowNoMatch()
		return map[string]any{"pos": 0, "len": 0}
	}
	start := self.lastIndices[0]
	end := self.lastIndices[1]
	return map[string]any{"pos": start, "len": end - start}
}

func (self *EReg) matchedLeft() *string {
	if !hxrt_eregHasMatch(self) {
		hxrt_eregThrowNoMatch()
		return nil
	}
	raw := *hxrt.StdString(self.lastSource)
	start := self.lastIndices[0]
	if start > len(raw) {
		start = len(raw)
	}
	return hxrt.StringFromLiteral(raw[:start])
}

func (self *EReg) matchedRight() *string {
	if !hxrt_eregHasMatch(self) {
		hxrt_eregThrowNoMatch()
		return nil
	}
	raw := *hxrt.StdString(self.lastSource)
	end := self.lastIndices[1]
	if end > len(raw) {
		end = len(raw)
	}
	return hxrt.StringFromLiteral(raw[end:])
}

func (self *EReg) split(source *string) *hxrt.Array {
	raw := *hxrt.StdString(source)
	if (self == nil) || (self.regex == nil) {
		return hxrt.NewArray(hxrt.StringFromLiteral(raw))
	}
	parts := self.regex.Split(raw, -1)
	out := hxrt.NewArray()
	for _, part := range parts {
		out.Push(hxrt.StringFromLiteral(part))
	}
	return out
}

func (self *EReg) replace(source *string, by *string) *string {
	if self == nil || self.regex == nil {
		return source
	}
	rawSource := *hxrt.StdString(source)
	rawBy := *hxrt.StdString(by)
	if self.global {
		return hxrt.StringFromLiteral(self.regex.ReplaceAllString(rawSource, rawBy))
	}
	first := self.regex.FindStringSubmatchIndex(rawSource)
	if first == nil {
		return hxrt.StringFromLiteral(rawSource)
	}
	replacement := self.regex.ExpandString(nil, rawBy, rawSource, first)
	out := rawSource[:first[0]] + string(replacement) + rawSource[first[1]:]
	return hxrt.StringFromLiteral(out)
}

func (self *EReg) map_(source *string, callback func(*EReg) *string) *string {
	if self == nil || self.regex == nil {
		return source
	}
	raw := *hxrt.StdString(source)
	if callback == nil {
		return hxrt.StringFromLiteral(raw)
	}
	var matches [][]int
	if self.global {
		matches = self.regex.FindAllStringSubmatchIndex(raw, -1)
	} else {
		if first := self.regex.FindStringSubmatchIndex(raw); first != nil {
			matches = [][]int{first}
		}
	}
	if len(matches) == 0 {
		return hxrt.StringFromLiteral(raw)
	}
	var builder strings.Builder
	cursor := 0
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		start := match[0]
		end := match[1]
		if start < cursor {
			start = cursor
		}
		if end < start {
			end = start
		}
		if start > len(raw) {
			start = len(raw)
		}
		if end > len(raw) {
			end = len(raw)
		}
		builder.WriteString(raw[cursor:start])
		indices := make([]int, len(match))
		copy(indices, match)
		self.lastSource = hxrt.StringFromLiteral(raw)
		self.lastIndices = indices
		replacement := callback(self)
		builder.WriteString(*hxrt.StdString(replacement))
		cursor = end
	}
	builder.WriteString(raw[cursor:])
	return hxrt.StringFromLiteral(builder.String())
}

type haxe__SerializedDate struct {
	ms float64
}

type haxe__SerializedBytes struct {
	data []byte
}

type haxe__SerializedClassRef struct {
	name string
}

type haxe__SerializedEnumRef struct {
	name string
}

type haxe__SerializedClass struct {
	name        string
	fieldNames  []string
	fieldValues []any
}

type haxe__SerializedEnum struct {
	name                string
	constructor         string
	constructorIndex    int
	hasConstructorIndex bool
	args                []any
}

type haxe__Unserializer__DefaultResolver struct {
}

type haxe__Unserializer__NullResolver struct {
}

var haxe__Serializer_USE_CACHE bool = false

var haxe__Serializer_USE_ENUM_INDEX bool = false

var haxe__Unserializer_DEFAULT_RESOLVER any = &haxe__Unserializer__DefaultResolver{}

var haxe__Unserializer_NULL_RESOLVER any = &haxe__Unserializer__NullResolver{}

func hxrt_serializerLookupClassName(typeName string) (string, bool) {
	switch typeName {
	case "haxe__Template":
		return "haxe.Template", true
	case "haxe___Int64_____Int64":
		return "haxe._Int64.___Int64", true
	case "haxe___Template__ExprCursor":
		return "haxe._Template.ExprCursor", true
	case "haxe___Template__TokenCursor":
		return "haxe._Template.TokenCursor", true
	case "haxe__ds__IntMap":
		return "haxe.ds.IntMap", true
	case "haxe__ds__List":
		return "haxe.ds.List", true
	case "haxe__ds__ObjectMap":
		return "haxe.ds.ObjectMap", true
	case "haxe__ds__StringMap":
		return "haxe.ds.StringMap", true
	case "haxe__ds___List__GoListIterator":
		return "haxe.ds._List.GoListIterator", true
	case "haxe__ds___List__GoListKeyValueIterator":
		return "haxe.ds._List.GoListKeyValueIterator", true
	case "haxe__iterators__MapKeyValueIterator":
		return "haxe.iterators.MapKeyValueIterator", true
	case "haxe__iterators__StringIterator":
		return "haxe.iterators.StringIterator", true
	case "haxe__iterators__StringKeyValueIterator":
		return "haxe.iterators.StringKeyValueIterator", true
	default:
		return "", false
	}
}

func hxrt_serializerLookupEnumConstructor(typeName string, tag int) (string, string, bool) {
	switch typeName {
	case "haxe___Template__TemplateExpr":
		constructors := []string{"OpVar", "OpExpr", "OpIf", "OpStr", "OpBlock", "OpForeach", "OpMacro"}
		if tag < 0 || tag >= len(constructors) {
			return "", "", false
		}
		return "haxe._Template.TemplateExpr", constructors[tag], true
	default:
		return "", "", false
	}
}

func hxrt_serializerLookupEnumConstructorByName(enumName string, index int) (string, bool) {
	switch enumName {
	case "haxe._Template.TemplateExpr":
		constructors := []string{"OpVar", "OpExpr", "OpIf", "OpStr", "OpBlock", "OpForeach", "OpMacro"}
		if index < 0 || index >= len(constructors) {
			return "", false
		}
		return constructors[index], true
	default:
		return "", false
	}
}

func hxrt_serializerLookupEnumIndexByName(enumName string, constructorName string) (int, bool) {
	switch enumName {
	case "haxe._Template.TemplateExpr":
		if constructorName == "OpVar" {
			return 0, true
		}
		if constructorName == "OpExpr" {
			return 1, true
		}
		if constructorName == "OpIf" {
			return 2, true
		}
		if constructorName == "OpStr" {
			return 3, true
		}
		if constructorName == "OpBlock" {
			return 4, true
		}
		if constructorName == "OpForeach" {
			return 5, true
		}
		if constructorName == "OpMacro" {
			return 6, true
		}
		return 0, false
	default:
		return 0, false
	}
}

func (self *haxe__Unserializer__DefaultResolver) resolveClass(name *string) any {
	className := *hxrt.StdString(name)
	if !hxrt_unserializerHasClass(className) {
		return nil
	}
	return &haxe__SerializedClassRef{name: className}
}

func (self *haxe__Unserializer__DefaultResolver) resolveEnum(name *string) any {
	enumName := *hxrt.StdString(name)
	if !hxrt_unserializerHasEnum(enumName) {
		return nil
	}
	return &haxe__SerializedEnumRef{name: enumName}
}

func (self *haxe__Unserializer__NullResolver) resolveClass(name *string) any {
	return nil
}

func (self *haxe__Unserializer__NullResolver) resolveEnum(name *string) any {
	return nil
}

func hxrt_unserializerHasClass(className string) bool {
	switch className {
	case "haxe.Template":
		return true
	case "haxe._Int64.___Int64":
		return true
	case "haxe._Template.ExprCursor":
		return true
	case "haxe._Template.TokenCursor":
		return true
	case "haxe.ds.IntMap":
		return true
	case "haxe.ds.List":
		return true
	case "haxe.ds.ObjectMap":
		return true
	case "haxe.ds.StringMap":
		return true
	case "haxe.ds._List.GoListIterator":
		return true
	case "haxe.ds._List.GoListKeyValueIterator":
		return true
	case "haxe.iterators.MapKeyValueIterator":
		return true
	case "haxe.iterators.StringIterator":
		return true
	case "haxe.iterators.StringKeyValueIterator":
		return true
	default:
		return false
	}
}

func hxrt_unserializerHasEnum(enumName string) bool {
	switch enumName {
	case "haxe._Template.TemplateExpr":
		return true
	default:
		return false
	}
}

func hxrt_unserializerBindSelf(instance any) {
	if instance == nil {
		return
	}
	rv := reflect.ValueOf(instance)
	if !rv.IsValid() || rv.Kind() != reflect.Pointer || rv.IsNil() {
		return
	}
	elem := rv.Elem()
	if !elem.IsValid() || elem.Kind() != reflect.Struct {
		return
	}
	field := elem.FieldByName("__hx_this")
	if !field.IsValid() {
		return
	}
	if !rv.Type().AssignableTo(field.Type()) {
		return
	}
	if field.CanSet() {
		field.Set(rv)
		return
	}
	if !field.CanAddr() {
		return
	}
	lifted := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
	lifted.Set(rv)
}

func hxrt_unserializerInvokeResolver(resolver any, methodName string, name string) (any, bool) {
	if resolver == nil {
		return nil, false
	}
	result := any(nil)
	ok := false
	defer func() {
		if recover() != nil {
			result = nil
			ok = false
		}
	}()
	rv := reflect.ValueOf(resolver)
	if !rv.IsValid() {
		return nil, false
	}
	for rv.IsValid() && rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return nil, false
		}
		rv = rv.Elem()
	}
	method := reflect.Value{}
	if rv.IsValid() {
		method = rv.MethodByName(methodName)
	}
	if !method.IsValid() && rv.IsValid() && rv.Kind() != reflect.Pointer && rv.CanAddr() {
		method = rv.Addr().MethodByName(methodName)
	}
	if !method.IsValid() && rv.IsValid() && rv.Kind() == reflect.Pointer && !rv.IsNil() {
		method = rv.Elem().MethodByName(methodName)
	}
	if !method.IsValid() && rv.IsValid() {
		switch rv.Kind() {
		case reflect.Struct:
			field := rv.FieldByName(methodName)
			if field.IsValid() {
				for field.IsValid() && field.Kind() == reflect.Interface {
					if field.IsNil() {
						break
					}
					field = field.Elem()
				}
				if field.IsValid() && field.Kind() == reflect.Func {
					method = field
				}
			}
		case reflect.Pointer:
			if !rv.IsNil() && rv.Elem().Kind() == reflect.Struct {
				field := rv.Elem().FieldByName(methodName)
				if field.IsValid() {
					for field.IsValid() && field.Kind() == reflect.Interface {
						if field.IsNil() {
							break
						}
						field = field.Elem()
					}
					if field.IsValid() && field.Kind() == reflect.Func {
						method = field
					}
				}
			}
		case reflect.Map:
			if rv.Type().Key().Kind() == reflect.String {
				field := rv.MapIndex(reflect.ValueOf(methodName))
				if field.IsValid() {
					for field.IsValid() && field.Kind() == reflect.Interface {
						if field.IsNil() {
							break
						}
						field = field.Elem()
					}
					if field.IsValid() && field.Kind() == reflect.Func {
						method = field
					}
				}
			}
		}
	}
	if !method.IsValid() || method.Kind() != reflect.Func {
		return nil, false
	}
	methodType := method.Type()
	if methodType.NumIn() != 1 || methodType.NumOut() < 1 {
		return nil, false
	}
	argType := methodType.In(0)
	nameValue := reflect.ValueOf(name)
	var arg reflect.Value
	if nameValue.Type().AssignableTo(argType) {
		arg = nameValue
	} else if nameValue.Type().ConvertibleTo(argType) {
		arg = nameValue.Convert(argType)
	} else if argType.Kind() == reflect.Pointer && argType.Elem().Kind() == reflect.String {
		nameCopy := name
		arg = reflect.ValueOf(&nameCopy)
		if !arg.Type().AssignableTo(argType) {
			if arg.Type().ConvertibleTo(argType) {
				arg = arg.Convert(argType)
			} else {
				return nil, false
			}
		}
	} else {
		return nil, false
	}
	out := method.Call([]reflect.Value{arg})
	if len(out) == 0 {
		return nil, true
	}
	value := out[0]
	if !value.IsValid() {
		return nil, true
	}
	for value.IsValid() && value.Kind() == reflect.Interface {
		if value.IsNil() {
			return nil, true
		}
		value = value.Elem()
	}
	if !value.IsValid() {
		return nil, true
	}
	switch value.Kind() {
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:
		if value.IsNil() {
			return nil, true
		}
	}
	result = value.Interface()
	ok = true
	return result, ok
}

func hxrt_unserializerResolveClass(self *haxe__Unserializer, name string) any {
	var resolver any
	if self != nil {
		resolver = self.resolver
	}
	if resolver == nil {
		resolver = haxe__Unserializer_DEFAULT_RESOLVER
	}
	if resolver == nil {
		return nil
	}
	switch current := resolver.(type) {
	case interface{ resolveClass(*string) any }:
		return current.resolveClass(hxrt.StringFromLiteral(name))
	case interface{ resolveClass(string) any }:
		return current.resolveClass(name)
	case interface{ resolveClass(any) any }:
		return current.resolveClass(name)
	case interface{ ResolveClass(*string) any }:
		return current.ResolveClass(hxrt.StringFromLiteral(name))
	case interface{ ResolveClass(string) any }:
		return current.ResolveClass(name)
	case interface{ ResolveClass(any) any }:
		return current.ResolveClass(name)
	}
	resolved, ok := hxrt_unserializerInvokeResolver(resolver, "resolveClass", name)
	if ok {
		return resolved
	}
	resolved, ok = hxrt_unserializerInvokeResolver(resolver, "ResolveClass", name)
	if ok {
		return resolved
	}
	return nil
}

func hxrt_unserializerResolveEnum(self *haxe__Unserializer, name string) any {
	var resolver any
	if self != nil {
		resolver = self.resolver
	}
	if resolver == nil {
		resolver = haxe__Unserializer_DEFAULT_RESOLVER
	}
	if resolver == nil {
		return nil
	}
	switch current := resolver.(type) {
	case interface{ resolveEnum(*string) any }:
		return current.resolveEnum(hxrt.StringFromLiteral(name))
	case interface{ resolveEnum(string) any }:
		return current.resolveEnum(name)
	case interface{ resolveEnum(any) any }:
		return current.resolveEnum(name)
	case interface{ ResolveEnum(*string) any }:
		return current.ResolveEnum(hxrt.StringFromLiteral(name))
	case interface{ ResolveEnum(string) any }:
		return current.ResolveEnum(name)
	case interface{ ResolveEnum(any) any }:
		return current.ResolveEnum(name)
	}
	resolved, ok := hxrt_unserializerInvokeResolver(resolver, "resolveEnum", name)
	if ok {
		return resolved
	}
	resolved, ok = hxrt_unserializerInvokeResolver(resolver, "ResolveEnum", name)
	if ok {
		return resolved
	}
	return nil
}

func hxrt_unserializerExtractNameField(resolved any) (string, bool) {
	rv := reflect.ValueOf(resolved)
	if !rv.IsValid() {
		return "", false
	}
	for rv.IsValid() && (rv.Kind() == reflect.Interface || rv.Kind() == reflect.Pointer) {
		if rv.IsNil() {
			return "", false
		}
		rv = rv.Elem()
	}
	if !rv.IsValid() || rv.Kind() != reflect.Struct {
		return "", false
	}
	field := rv.FieldByName("name")
	if !field.IsValid() {
		return "", false
	}
	for field.IsValid() && field.Kind() == reflect.Interface {
		if field.IsNil() {
			return "", false
		}
		field = field.Elem()
	}
	if !field.IsValid() {
		return "", false
	}
	if field.Kind() == reflect.String {
		return field.String(), true
	}
	if field.Kind() == reflect.Pointer {
		if field.IsNil() {
			return "", false
		}
		if field.Elem().Kind() == reflect.String {
			return field.Elem().String(), true
		}
	}
	return "", false
}

func hxrt_unserializerResolvedClassName(resolved any) (string, bool) {
	switch current := resolved.(type) {
	case haxe__SerializedClassRef:
		return current.name, true
	case *haxe__SerializedClassRef:
		if current == nil {
			return "", false
		}
		return current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	}
	return hxrt_unserializerExtractNameField(resolved)
}

func hxrt_unserializerResolvedEnumName(resolved any) (string, bool) {
	switch current := resolved.(type) {
	case haxe__SerializedEnumRef:
		return current.name, true
	case *haxe__SerializedEnumRef:
		if current == nil {
			return "", false
		}
		return current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	}
	return hxrt_unserializerExtractNameField(resolved)
}

func hxrt_unserializerCreateClassInstance(className string) (any, bool) {
	switch className {
	case "haxe.Template":
		instance := &haxe__Template{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe._Int64.___Int64":
		instance := &haxe___Int64_____Int64{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe._Template.ExprCursor":
		instance := &haxe___Template__ExprCursor{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe._Template.TokenCursor":
		instance := &haxe___Template__TokenCursor{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe.ds.IntMap":
		instance := &haxe__ds__IntMap{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe.ds.List":
		instance := &haxe__ds__List{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe.ds.ObjectMap":
		instance := &haxe__ds__ObjectMap{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe.ds.StringMap":
		instance := &haxe__ds__StringMap{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe.ds._List.GoListIterator":
		instance := &haxe__ds___List__GoListIterator{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe.ds._List.GoListKeyValueIterator":
		instance := &haxe__ds___List__GoListKeyValueIterator{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe.iterators.MapKeyValueIterator":
		instance := &haxe__iterators__MapKeyValueIterator{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe.iterators.StringIterator":
		instance := &haxe__iterators__StringIterator{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe.iterators.StringKeyValueIterator":
		instance := &haxe__iterators__StringKeyValueIterator{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	default:
		return nil, false
	}
}

func hxrt_unserializerCreateEnumInstance(enumName string, constructorName string, constructorIndex int, hasConstructorIndex bool, args []any) (any, bool) {
	if hasConstructorIndex {
		if _, ok := hxrt_serializerLookupEnumConstructorByName(enumName, constructorIndex); !ok {
			return nil, false
		}
	} else {
		resolvedIndex, ok := hxrt_serializerLookupEnumIndexByName(enumName, constructorName)
		if !ok {
			return nil, false
		}
		constructorIndex = resolvedIndex
	}
	switch enumName {
	case "haxe._Template.TemplateExpr":
		enumValue := &haxe___Template__TemplateExpr{tag: constructorIndex}
		if len(args) > 0 {
			copied := make([]any, len(args))
			copy(copied, args)
			enumValue.params = copied
		}
		return enumValue, true
	default:
		return nil, false
	}
}

type haxe__Serializer struct {
	buf          *string
	useCache     bool
	useEnumIndex bool
	stringCache  map[string]int
	cacheRefs    map[uintptr]int
}

func New_haxe__Serializer() *haxe__Serializer {
	return &haxe__Serializer{buf: hxrt.StringFromLiteral(""), useCache: haxe__Serializer_USE_CACHE, useEnumIndex: haxe__Serializer_USE_ENUM_INDEX, stringCache: map[string]int{}, cacheRefs: map[uintptr]int{}}
}

func (self *haxe__Serializer) serialize(value any) {
	if self == nil {
		return
	}
	hxrt_serializerWriteValue(self, value)
}

func (self *haxe__Serializer) serializeException(value any) {
	if self == nil {
		return
	}
	hxrt_serializerAppend(self, "x")
	hxrt_serializerWriteValue(self, value)
}

func (self *haxe__Serializer) toString() *string {
	if self == nil || self.buf == nil {
		return hxrt.StringFromLiteral("")
	}
	return self.buf
}

func hxrt_serializerAppend(self *haxe__Serializer, chunk string) {
	if self == nil {
		return
	}
	if self.buf == nil {
		self.buf = hxrt.StringFromLiteral("")
	}
	self.buf = hxrt.StringFromLiteral(*hxrt.StdString(self.buf) + chunk)
}

func hxrt_serializerEscape(value *string) string {
	raw := *hxrt.StdString(value)
	var builder strings.Builder
	hex := "0123456789ABCDEF"
	for i := 0; i < len(raw); i++ {
		b := raw[i]
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '_' || b == '.' || b == '-' {
			builder.WriteByte(b)
			continue
		}
		builder.WriteByte('%')
		builder.WriteByte(hex[b>>4])
		builder.WriteByte(hex[b&0x0F])
	}
	return builder.String()
}

func hxrt_serializerWriteStringToken(self *haxe__Serializer, value string) {
	if self == nil {
		return
	}
	if self.stringCache == nil {
		self.stringCache = map[string]int{}
	}
	if index, ok := self.stringCache[value]; ok {
		hxrt_serializerAppend(self, "R"+strconv.Itoa(index))
		return
	}
	escaped := hxrt_serializerEscape(hxrt.StringFromLiteral(value))
	self.stringCache[value] = len(self.stringCache)
	hxrt_serializerAppend(self, "y"+strconv.Itoa(len(escaped))+":"+escaped)
}

func hxrt_serializerWriteIntToken(self *haxe__Serializer, value int64) {
	if value == 0 {
		hxrt_serializerAppend(self, "z")
		return
	}
	hxrt_serializerAppend(self, "i"+strconv.FormatInt(value, 10))
}

func hxrt_serializerWriteBytesToken(self *haxe__Serializer, raw []byte) {
	encoded := base64.StdEncoding.EncodeToString(raw)
	encoded = strings.TrimRight(encoded, "=")
	hxrt_serializerAppend(self, "s"+strconv.Itoa(len(encoded))+":"+encoded)
}

func hxrt_serializerWriteEnumToken(self *haxe__Serializer, enumName string, constructorName string, constructorIndex int, hasConstructorIndex bool, args []any) {
	if self != nil && self.useEnumIndex && hasConstructorIndex {
		hxrt_serializerAppend(self, "j")
		hxrt_serializerWriteStringToken(self, enumName)
		hxrt_serializerAppend(self, ":"+strconv.Itoa(constructorIndex))
	} else {
		hxrt_serializerAppend(self, "w")
		hxrt_serializerWriteStringToken(self, enumName)
		hxrt_serializerWriteStringToken(self, constructorName)
	}
	hxrt_serializerAppend(self, ":"+strconv.Itoa(len(args)))
	for _, arg := range args {
		hxrt_serializerWriteValue(self, arg)
	}
}

func hxrt_serializerWriteListToken(self *haxe__Serializer, items []any) {
	hxrt_serializerAppend(self, "l")
	for _, item := range items {
		hxrt_serializerWriteValue(self, item)
	}
	hxrt_serializerAppend(self, "h")
}

func hxrt_serializerWriteStringMapToken(self *haxe__Serializer, entries map[string]any) {
	hxrt_serializerAppend(self, "b")
	keys := make([]string, 0, len(entries))
	for key := range entries {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		hxrt_serializerWriteStringToken(self, key)
		hxrt_serializerWriteValue(self, entries[key])
	}
	hxrt_serializerAppend(self, "h")
}

func hxrt_serializerWriteIntMapToken(self *haxe__Serializer, entries map[int]any) {
	hxrt_serializerAppend(self, "q")
	keys := make([]int, 0, len(entries))
	for key := range entries {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		hxrt_serializerAppend(self, ":"+strconv.Itoa(key))
		hxrt_serializerWriteValue(self, entries[key])
	}
	hxrt_serializerAppend(self, "h")
}

func hxrt_serializerWriteObjectMapToken(self *haxe__Serializer, entries []hxrt.ObjectMapEntry) {
	hxrt_serializerAppend(self, "M")
	for _, entry := range entries {
		hxrt_serializerWriteValue(self, entry.Key)
		hxrt_serializerWriteValue(self, entry.Value)
	}
	hxrt_serializerAppend(self, "h")
}

func hxrt_serializerReflectAny(value reflect.Value) (any, bool) {
	defer func() {
		if recover() != nil {
		}
	}()
	if !value.IsValid() {
		return nil, false
	}
	if value.CanInterface() {
		return value.Interface(), true
	}
	if !value.CanAddr() {
		return nil, false
	}
	lifted := reflect.NewAt(value.Type(), unsafe.Pointer(value.UnsafeAddr())).Elem()
	if !lifted.IsValid() || !lifted.CanInterface() {
		return nil, false
	}
	return lifted.Interface(), true
}

func hxrt_serializerWriteSerializedClass(self *haxe__Serializer, serialized *haxe__SerializedClass) {
	if serialized == nil {
		hxrt_serializerAppend(self, "n")
		return
	}
	hxrt_serializerAppend(self, "c")
	hxrt_serializerWriteStringToken(self, serialized.name)
	limit := len(serialized.fieldNames)
	if len(serialized.fieldValues) < limit {
		limit = len(serialized.fieldValues)
	}
	for i := 0; i < limit; i++ {
		hxrt_serializerWriteStringToken(self, serialized.fieldNames[i])
		hxrt_serializerWriteValue(self, serialized.fieldValues[i])
	}
	hxrt_serializerAppend(self, "g")
}

func hxrt_serializerWriteSerializedEnum(self *haxe__Serializer, serialized *haxe__SerializedEnum) {
	if serialized == nil {
		hxrt_serializerAppend(self, "n")
		return
	}
	constructorIndex := serialized.constructorIndex
	hasConstructorIndex := serialized.hasConstructorIndex
	if !hasConstructorIndex {
		if resolvedIndex, ok := hxrt_serializerLookupEnumIndexByName(serialized.name, serialized.constructor); ok {
			constructorIndex = resolvedIndex
			hasConstructorIndex = true
		}
	}
	hxrt_serializerWriteEnumToken(self, serialized.name, serialized.constructor, constructorIndex, hasConstructorIndex, serialized.args)
}

func hxrt_serializerTryClassStruct(self *haxe__Serializer, value any, ref reflect.Value) bool {
	defer func() {
		if recover() != nil {
		}
	}()
	if !ref.IsValid() || ref.Kind() != reflect.Struct {
		return false
	}
	className, ok := hxrt_serializerLookupClassName(ref.Type().Name())
	if !ok {
		return false
	}
	if hxrt_serializerTrackRef(self, value) {
		return true
	}
	if custom, ok := value.(interface{ hxSerialize(*haxe__Serializer) }); ok {
		hxrt_serializerAppend(self, "C")
		hxrt_serializerWriteStringToken(self, className)
		custom.hxSerialize(self)
		hxrt_serializerAppend(self, "g")
		return true
	}
	if custom, ok := value.(interface{ HxSerialize(*haxe__Serializer) }); ok {
		hxrt_serializerAppend(self, "C")
		hxrt_serializerWriteStringToken(self, className)
		custom.HxSerialize(self)
		hxrt_serializerAppend(self, "g")
		return true
	}
	if ref.CanAddr() {
		addr := ref.Addr().Interface()
		if custom, ok := addr.(interface{ hxSerialize(*haxe__Serializer) }); ok {
			hxrt_serializerAppend(self, "C")
			hxrt_serializerWriteStringToken(self, className)
			custom.hxSerialize(self)
			hxrt_serializerAppend(self, "g")
			return true
		}
		if custom, ok := addr.(interface{ HxSerialize(*haxe__Serializer) }); ok {
			hxrt_serializerAppend(self, "C")
			hxrt_serializerWriteStringToken(self, className)
			custom.HxSerialize(self)
			hxrt_serializerAppend(self, "g")
			return true
		}
	}
	hxrt_serializerAppend(self, "c")
	hxrt_serializerWriteStringToken(self, className)
	refType := ref.Type()
	for i := 0; i < ref.NumField(); i++ {
		fieldInfo := refType.Field(i)
		fieldName := fieldInfo.Name
		if strings.HasPrefix(fieldName, "__hx_") {
			continue
		}
		fieldValue, ok := hxrt_serializerReflectAny(ref.Field(i))
		if !ok {
			return false
		}
		hxrt_serializerWriteStringToken(self, fieldName)
		hxrt_serializerWriteValue(self, fieldValue)
	}
	hxrt_serializerAppend(self, "g")
	return true
}

func hxrt_serializerTryEnumStruct(self *haxe__Serializer, value any, ref reflect.Value) bool {
	defer func() {
		if recover() != nil {
		}
	}()
	if !ref.IsValid() || ref.Kind() != reflect.Struct {
		return false
	}
	tagField := ref.FieldByName("tag")
	paramsField := ref.FieldByName("params")
	if !tagField.IsValid() || !paramsField.IsValid() || paramsField.Kind() != reflect.Slice {
		return false
	}
	var tag int
	switch tagField.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tag = int(tagField.Int())
	default:
		return false
	}
	enumName, constructorName, ok := hxrt_serializerLookupEnumConstructor(ref.Type().Name(), tag)
	if !ok {
		return false
	}
	if hxrt_serializerTrackRef(self, value) {
		return true
	}
	args := make([]any, 0, paramsField.Len())
	for i := 0; i < paramsField.Len(); i++ {
		value, ok := hxrt_serializerReflectAny(paramsField.Index(i))
		if !ok {
			return false
		}
		args = append(args, value)
	}
	hxrt_serializerWriteEnumToken(self, enumName, constructorName, tag, true, args)
	return true
}

func hxrt_serializerTryDateStruct(self *haxe__Serializer, ref reflect.Value) bool {
	defer func() {
		if recover() != nil {
		}
	}()
	if !ref.IsValid() || ref.Kind() != reflect.Struct {
		return false
	}
	msField := ref.FieldByName("ms")
	if !msField.IsValid() || msField.Kind() != reflect.Float64 {
		return false
	}
	ms := msField.Float()
	hxrt_serializerAppend(self, "v"+strconv.FormatFloat(ms, 'g', -1, 64))
	return true
}

func hxrt_serializerTrySpecialReflect(self *haxe__Serializer, value any) bool {
	ref := reflect.ValueOf(value)
	if !ref.IsValid() {
		return false
	}
	for ref.Kind() == reflect.Interface {
		if ref.IsNil() {
			return false
		}
		ref = ref.Elem()
	}
	for ref.Kind() == reflect.Pointer {
		if ref.IsNil() {
			return false
		}
		ref = ref.Elem()
	}
	if !ref.IsValid() || ref.Kind() != reflect.Struct {
		return false
	}
	typeName := ref.Type().Name()
	if typeName == "Date" {
		if hxrt_serializerTryDateStruct(self, ref) {
			return true
		}
	}
	if typeName == "haxe__io__Bytes" {
		bytesField := ref.FieldByName("b")
		if !bytesField.IsValid() || bytesField.Kind() != reflect.Slice {
			return false
		}
		raw := make([]byte, bytesField.Len())
		for i := 0; i < bytesField.Len(); i++ {
			entry := bytesField.Index(i)
			if !entry.IsValid() {
				return false
			}
			switch entry.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				raw[i] = byte(entry.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				raw[i] = byte(entry.Uint())
			default:
				return false
			}
		}
		hxrt_serializerWriteBytesToken(self, raw)
		return true
	}
	if hxrt_serializerTryEnumStruct(self, value, ref) {
		return true
	}
	if hxrt_serializerTryClassStruct(self, value, ref) {
		return true
	}
	return false
}

func hxrt_serializerTryTypeValueRef(self *haxe__Serializer, value any) bool {
	ref := reflect.ValueOf(value)
	if !ref.IsValid() {
		return false
	}
	for ref.Kind() == reflect.Interface || ref.Kind() == reflect.Pointer {
		if ref.IsNil() {
			return false
		}
		ref = ref.Elem()
	}
	if !ref.IsValid() || ref.Kind() != reflect.Struct {
		return false
	}
	typeName := ref.Type().Name()
	if typeName != "hxrt__TypeClassValue" && typeName != "hxrt__TypeEnumValue" {
		return false
	}
	nameField := ref.FieldByName("name")
	if !nameField.IsValid() {
		return false
	}
	for nameField.IsValid() && nameField.Kind() == reflect.Interface {
		if nameField.IsNil() {
			return false
		}
		nameField = nameField.Elem()
	}
	if !nameField.IsValid() {
		return false
	}
	resolvedName := ""
	if nameField.Kind() == reflect.String {
		resolvedName = nameField.String()
	} else if nameField.Kind() == reflect.Pointer {
		if nameField.IsNil() || nameField.Elem().Kind() != reflect.String {
			return false
		}
		resolvedName = nameField.Elem().String()
	} else {
		return false
	}
	if resolvedName == "" {
		return false
	}
	if typeName == "hxrt__TypeClassValue" {
		hxrt_serializerAppend(self, "A")
	} else {
		hxrt_serializerAppend(self, "B")
	}
	hxrt_serializerWriteStringToken(self, resolvedName)
	return true
}

func hxrt_serializerTrackRef(self *haxe__Serializer, value any) bool {
	if self == nil || !self.useCache {
		return false
	}
	if self.cacheRefs == nil {
		self.cacheRefs = map[uintptr]int{}
	}
	ref := reflect.ValueOf(value)
	if !ref.IsValid() {
		return false
	}
	for ref.Kind() == reflect.Interface {
		if ref.IsNil() {
			return false
		}
		ref = ref.Elem()
	}
	switch ref.Kind() {
	case reflect.Map, reflect.Slice, reflect.Pointer:
		if ref.IsNil() {
			return false
		}
		key := ref.Pointer()
		if index, ok := self.cacheRefs[key]; ok {
			hxrt_serializerAppend(self, "r"+strconv.Itoa(index))
			return true
		}
		self.cacheRefs[key] = len(self.cacheRefs)
	}
	return false
}

func hxrt_serializerWriteArray(self *haxe__Serializer, current *hxrt.Array) {
	if current == nil {
		hxrt_serializerAppend(self, "n")
		return
	}
	if hxrt_serializerTrackRef(self, current) {
		return
	}
	hxrt_serializerAppend(self, "a")
	nullRun := 0
	for _, item := range current.Values() {
		if item == nil {
			nullRun = (nullRun + 1)
			continue
		}
		if nullRun > 1 {
			hxrt_serializerAppend(self, ("u" + strconv.Itoa(nullRun)))
		} else {
			if nullRun == 1 {
				hxrt_serializerAppend(self, "n")
			}
		}
		nullRun = 0
		hxrt_serializerWriteValue(self, item)
	}
	if nullRun > 1 {
		hxrt_serializerAppend(self, ("u" + strconv.Itoa(nullRun)))
	} else {
		if nullRun == 1 {
			hxrt_serializerAppend(self, "n")
		}
	}
	hxrt_serializerAppend(self, "h")
}

func haxe__Serializer_run(value any) *string {
	serializer := New_haxe__Serializer()
	serializer.serialize(value)
	return serializer.toString()
}

func hxrt_serializerWriteValue(self *haxe__Serializer, value any) {
	if self == nil {
		return
	}
	if value == nil {
		hxrt_serializerAppend(self, "n")
		return
	}
	if hxrt_serializerTryTypeValueRef(self, value) {
		return
	}
	switch current := value.(type) {
	case bool:
		if current {
			hxrt_serializerAppend(self, "t")
		} else {
			hxrt_serializerAppend(self, "f")
		}
		return
	case string:
		hxrt_serializerWriteStringToken(self, current)
		return
	case *string:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			hxrt_serializerWriteStringToken(self, *current)
		}
		return
	case haxe__SerializedDate:
		hxrt_serializerAppend(self, "v"+strconv.FormatFloat(current.ms, 'g', -1, 64))
		return
	case *haxe__SerializedDate:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			hxrt_serializerAppend(self, "v"+strconv.FormatFloat(current.ms, 'g', -1, 64))
		}
		return
	case haxe__SerializedBytes:
		hxrt_serializerWriteBytesToken(self, current.data)
		return
	case *haxe__SerializedBytes:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			hxrt_serializerWriteBytesToken(self, current.data)
		}
		return
	case haxe__SerializedClass:
		hxrt_serializerWriteSerializedClass(self, &current)
		return
	case *haxe__SerializedClass:
		hxrt_serializerWriteSerializedClass(self, current)
		return
	case haxe__SerializedEnum:
		hxrt_serializerWriteSerializedEnum(self, &current)
		return
	case *haxe__SerializedEnum:
		hxrt_serializerWriteSerializedEnum(self, current)
		return
	case haxe__SerializedClassRef:
		hxrt_serializerAppend(self, "A")
		hxrt_serializerWriteStringToken(self, current.name)
		return
	case *haxe__SerializedClassRef:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			hxrt_serializerAppend(self, "A")
			hxrt_serializerWriteStringToken(self, current.name)
		}
		return
	case haxe__SerializedEnumRef:
		hxrt_serializerAppend(self, "B")
		hxrt_serializerWriteStringToken(self, current.name)
		return
	case *haxe__SerializedEnumRef:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			hxrt_serializerAppend(self, "B")
			hxrt_serializerWriteStringToken(self, current.name)
		}
		return
	case *hxrt.Array:
		hxrt_serializerWriteArray(self, current)
		return
	case *haxe__ds__List:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			if hxrt_serializerTrackRef(self, current) {
				return
			}
			iterator := current.iterator()
			items := make([]any, 0, current.length)
			for iterator.hasNext() {
				items = append(items, iterator.next())
			}
			hxrt_serializerWriteListToken(self, items)
		}
		return
	case *haxe__ds__StringMap:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			if hxrt_serializerTrackRef(self, current) {
				return
			}
			hxrt_serializerWriteStringMapToken(self, hxrt.StringMapSnapshot(current.h))
		}
		return
	case *haxe__ds__IntMap:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			if hxrt_serializerTrackRef(self, current) {
				return
			}
			hxrt_serializerWriteIntMapToken(self, hxrt.IntMapSnapshot(current.h))
		}
		return
	case *haxe__ds__ObjectMap:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			if hxrt_serializerTrackRef(self, current) {
				return
			}
			hxrt_serializerWriteObjectMapToken(self, hxrt.ObjectMapSnapshot(current.h))
		}
		return
	case int:
		hxrt_serializerWriteIntToken(self, int64(current))
		return
	case int8:
		hxrt_serializerWriteIntToken(self, int64(current))
		return
	case int16:
		hxrt_serializerWriteIntToken(self, int64(current))
		return
	case int32:
		hxrt_serializerWriteIntToken(self, int64(current))
		return
	case int64:
		hxrt_serializerWriteIntToken(self, current)
		return
	case uint:
		hxrt_serializerWriteIntToken(self, int64(current))
		return
	case uint8:
		hxrt_serializerWriteIntToken(self, int64(current))
		return
	case uint16:
		hxrt_serializerWriteIntToken(self, int64(current))
		return
	case uint32:
		hxrt_serializerWriteIntToken(self, int64(current))
		return
	case uint64:
		hxrt_serializerWriteIntToken(self, int64(current))
		return
	case float32:
		value64 := float64(current)
		if math.IsNaN(value64) {
			hxrt_serializerAppend(self, "k")
			return
		}
		if math.IsInf(value64, 1) {
			hxrt_serializerAppend(self, "p")
			return
		}
		if math.IsInf(value64, -1) {
			hxrt_serializerAppend(self, "m")
			return
		}
		hxrt_serializerAppend(self, "d"+strconv.FormatFloat(value64, 'g', -1, 64))
		return
	case float64:
		if math.IsNaN(current) {
			hxrt_serializerAppend(self, "k")
			return
		}
		if math.IsInf(current, 1) {
			hxrt_serializerAppend(self, "p")
			return
		}
		if math.IsInf(current, -1) {
			hxrt_serializerAppend(self, "m")
			return
		}
		hxrt_serializerAppend(self, "d"+strconv.FormatFloat(current, 'g', -1, 64))
		return
	}
	ref := reflect.ValueOf(value)
	if !ref.IsValid() {
		hxrt_serializerAppend(self, "n")
		return
	}
	if hxrt_serializerTrySpecialReflect(self, value) {
		return
	}
	switch ref.Kind() {
	case reflect.Slice, reflect.Array:
		if ref.Kind() == reflect.Slice && ref.IsNil() {
			hxrt_serializerAppend(self, "n")
			return
		}
		if hxrt_serializerTrackRef(self, value) {
			return
		}
		hxrt_serializerAppend(self, "a")
		nullRun := 0
		for i := 0; i < ref.Len(); i++ {
			item := ref.Index(i).Interface()
			if item == nil {
				nullRun++
				continue
			}
			if nullRun > 1 {
				hxrt_serializerAppend(self, "u"+strconv.Itoa(nullRun))
			} else if nullRun == 1 {
				hxrt_serializerAppend(self, "n")
			}
			nullRun = 0
			hxrt_serializerWriteValue(self, item)
		}
		if nullRun > 1 {
			hxrt_serializerAppend(self, "u"+strconv.Itoa(nullRun))
		} else if nullRun == 1 {
			hxrt_serializerAppend(self, "n")
		}
		hxrt_serializerAppend(self, "h")
		return
	case reflect.Map:
		if ref.IsNil() {
			hxrt_serializerAppend(self, "n")
			return
		}
		if ref.Type().Key().Kind() != reflect.String {
			hxrt.Throw(hxrt.StringFromLiteral("Serializer map keys must be strings"))
			hxrt_serializerAppend(self, "n")
			return
		}
		if hxrt_serializerTrackRef(self, value) {
			return
		}
		hxrt_serializerAppend(self, "o")
		keys := ref.MapKeys()
		sortedKeys := make([]string, 0, len(keys))
		for _, key := range keys {
			sortedKeys = append(sortedKeys, key.String())
		}
		sort.Strings(sortedKeys)
		for _, key := range sortedKeys {
			hxrt_serializerWriteStringToken(self, key)
			valueRef := ref.MapIndex(reflect.ValueOf(key))
			if valueRef.IsValid() {
				hxrt_serializerWriteValue(self, valueRef.Interface())
			} else {
				hxrt_serializerWriteValue(self, nil)
			}
		}
		hxrt_serializerAppend(self, "g")
		return
	case reflect.Pointer:
		if ref.IsNil() {
			hxrt_serializerAppend(self, "n")
			return
		}
		if hxrt_serializerTrackRef(self, value) {
			return
		}
		hxrt_serializerWriteValue(self, ref.Elem().Interface())
		return
	}
	hxrt.Throw(hxrt.StringFromLiteral("Unsupported serializer value type"))
	hxrt_serializerAppend(self, "n")
}

type haxe__Unserializer struct {
	buf         *string
	pos         int
	stringCache []*string
	cache       []any
	resolver    any
}

func New_haxe__Unserializer(buf *string) *haxe__Unserializer {
	resolver := haxe__Unserializer_DEFAULT_RESOLVER
	if resolver == nil {
		resolver = &haxe__Unserializer__DefaultResolver{}
		haxe__Unserializer_DEFAULT_RESOLVER = resolver
	}
	return &haxe__Unserializer{buf: buf, pos: 0, stringCache: []*string{}, cache: []any{}, resolver: resolver}
}

func (self *haxe__Unserializer) setResolver(resolver any) {
	if self == nil {
		return
	}
	if resolver == nil {
		self.resolver = haxe__Unserializer_NULL_RESOLVER
		return
	}
	self.resolver = resolver
}

func (self *haxe__Unserializer) getResolver() any {
	if self == nil {
		return nil
	}
	return self.resolver
}

func (self *haxe__Unserializer) unserialize() any {
	if self == nil || self.buf == nil {
		return nil
	}
	return haxe__Unserializer_readValue(self)
}

func haxe__Unserializer_readUInt(self *haxe__Unserializer) int {
	raw := *hxrt.StdString(self.buf)
	start := self.pos
	for self.pos < len(raw) && raw[self.pos] >= '0' && raw[self.pos] <= '9' {
		self.pos++
	}
	if self.pos == start {
		hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized integer"))
		return 0
	}
	parsed, err := strconv.Atoi(raw[start:self.pos])
	if err != nil {
		hxrt.Throw(err)
		return 0
	}
	return parsed
}

func haxe__Unserializer_readDigits(self *haxe__Unserializer) int {
	raw := *hxrt.StdString(self.buf)
	start := self.pos
	if self.pos < len(raw) && (raw[self.pos] == '-' || raw[self.pos] == '+') {
		self.pos++
	}
	digitStart := self.pos
	for self.pos < len(raw) && raw[self.pos] >= '0' && raw[self.pos] <= '9' {
		self.pos++
	}
	if self.pos == digitStart {
		hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized integer"))
		return 0
	}
	parsed, err := strconv.Atoi(raw[start:self.pos])
	if err != nil {
		hxrt.Throw(err)
		return 0
	}
	return parsed
}

func hxrt_unserializerSetField(target any, fieldName string, value any) {
	if target == nil {
		return
	}
	switch obj := target.(type) {
	case map[string]any:
		obj[fieldName] = value
		return
	case map[any]any:
		obj[fieldName] = value
		return
	case *map[string]any:
		if obj != nil {
			(*obj)[fieldName] = value
		}
		return
	case *map[any]any:
		if obj != nil {
			(*obj)[fieldName] = value
		}
		return
	}
	rv := reflect.ValueOf(target)
	if !rv.IsValid() || rv.Kind() != reflect.Pointer || rv.IsNil() {
		return
	}
	elem := rv.Elem()
	if !elem.IsValid() || elem.Kind() != reflect.Struct {
		return
	}
	field := elem.FieldByName(fieldName)
	if !field.IsValid() {
		return
	}
	targetField := field
	if !targetField.CanSet() {
		if !targetField.CanAddr() {
			return
		}
		targetField = reflect.NewAt(targetField.Type(), unsafe.Pointer(targetField.UnsafeAddr())).Elem()
	}
	if value == nil {
		targetField.Set(reflect.Zero(targetField.Type()))
		return
	}
	incoming := reflect.ValueOf(value)
	if incoming.IsValid() && incoming.Type().AssignableTo(targetField.Type()) {
		targetField.Set(incoming)
		return
	}
	if incoming.IsValid() && incoming.Type().ConvertibleTo(targetField.Type()) {
		targetField.Set(incoming.Convert(targetField.Type()))
		return
	}
	if targetField.Kind() == reflect.Interface && incoming.IsValid() {
		targetField.Set(incoming)
	}
}

func hxrt_unserializerReadObjectFields(self *haxe__Unserializer, target any, invalidMessage string) {
	raw := *hxrt.StdString(self.buf)
	for {
		if self.pos >= len(raw) {
			hxrt.Throw(hxrt.StringFromLiteral(invalidMessage))
			return
		}
		if raw[self.pos] == 'g' {
			self.pos++
			return
		}
		fieldNameAny := haxe__Unserializer_readValue(self)
		fieldName := *hxrt.StdString(fieldNameAny)
		fieldValue := haxe__Unserializer_readValue(self)
		hxrt_unserializerSetField(target, fieldName, fieldValue)
	}
}

func haxe__Unserializer_readHexNibble(ch byte) int {
	switch {
	case ch >= '0' && ch <= '9':
		return int(ch - '0')
	case ch >= 'A' && ch <= 'F':
		return int(ch-'A') + 10
	case ch >= 'a' && ch <= 'f':
		return int(ch-'a') + 10
	default:
		return -1
	}
}

func haxe__Unserializer_unescape(value string) *string {
	out := make([]byte, 0, len(value))
	for i := 0; i < len(value); i++ {
		if value[i] != '%' {
			out = append(out, value[i])
			continue
		}
		if i+2 >= len(value) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized string escape"))
			return hxrt.StringFromLiteral("")
		}
		high := haxe__Unserializer_readHexNibble(value[i+1])
		low := haxe__Unserializer_readHexNibble(value[i+2])
		if high < 0 || low < 0 {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized string escape"))
			return hxrt.StringFromLiteral("")
		}
		out = append(out, byte((high<<4)|low))
		i += 2
	}
	return hxrt.StringFromLiteral(string(out))
}

func haxe__Unserializer_readValue(self *haxe__Unserializer) any {
	if self == nil || self.buf == nil {
		return nil
	}
	raw := *hxrt.StdString(self.buf)
	if self.pos >= len(raw) {
		hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized string"))
		return nil
	}
	token := raw[self.pos]
	self.pos++
	switch token {
	case 'n':
		return nil
	case 't':
		return true
	case 'f':
		return false
	case 'z':
		return 0
	case 'k':
		return math.NaN()
	case 'p':
		return math.Inf(1)
	case 'm':
		return math.Inf(-1)
	case 'i':
		start := self.pos
		if self.pos < len(raw) && (raw[self.pos] == '-' || raw[self.pos] == '+') {
			self.pos++
		}
		for self.pos < len(raw) && raw[self.pos] >= '0' && raw[self.pos] <= '9' {
			self.pos++
		}
		if self.pos == start || (self.pos == start+1 && (raw[start] == '-' || raw[start] == '+')) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized integer"))
			return 0
		}
		parsed, err := strconv.Atoi(raw[start:self.pos])
		if err != nil {
			hxrt.Throw(err)
			return 0
		}
		return parsed
	case 'd':
		start := self.pos
		hasDigit := false
		for self.pos < len(raw) {
			ch := raw[self.pos]
			if ch >= '0' && ch <= '9' {
				hasDigit = true
				self.pos++
				continue
			}
			if ch == '+' || ch == '-' || ch == '.' || ch == 'e' || ch == 'E' {
				self.pos++
				continue
			}
			break
		}
		if !hasDigit {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized float"))
			return 0.0
		}
		parsed, err := strconv.ParseFloat(raw[start:self.pos], 64)
		if err != nil {
			hxrt.Throw(err)
			return 0.0
		}
		return parsed
	case 'v':
		start := self.pos
		hasDigit := false
		for self.pos < len(raw) {
			ch := raw[self.pos]
			if ch >= '0' && ch <= '9' {
				hasDigit = true
				self.pos++
				continue
			}
			if ch == '+' || ch == '-' || ch == '.' || ch == 'e' || ch == 'E' {
				self.pos++
				continue
			}
			break
		}
		if !hasDigit {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized date"))
			return &haxe__SerializedDate{ms: 0}
		}
		parsed, err := strconv.ParseFloat(raw[start:self.pos], 64)
		if err != nil {
			hxrt.Throw(err)
			return &haxe__SerializedDate{ms: 0}
		}
		return &haxe__SerializedDate{ms: parsed}
	case 's':
		length := haxe__Unserializer_readUInt(self)
		if self.pos >= len(raw) || raw[self.pos] != ':' {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized bytes"))
			return &haxe__SerializedBytes{data: []byte{}}
		}
		self.pos++
		if length < 0 || self.pos+length > len(raw) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized bytes length"))
			return &haxe__SerializedBytes{data: []byte{}}
		}
		encoded := raw[self.pos : self.pos+length]
		self.pos += length
		decoded, err := base64.RawStdEncoding.DecodeString(encoded)
		if err != nil {
			decoded, err = base64.StdEncoding.DecodeString(encoded)
			if err != nil {
				hxrt.Throw(err)
				return &haxe__SerializedBytes{data: []byte{}}
			}
		}
		out := make([]byte, len(decoded))
		copy(out, decoded)
		return &haxe__SerializedBytes{data: out}
	case 'y':
		length := haxe__Unserializer_readUInt(self)
		if self.pos >= len(raw) || raw[self.pos] != ':' {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized string"))
			return hxrt.StringFromLiteral("")
		}
		self.pos++
		if length < 0 || self.pos+length > len(raw) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized string length"))
			return hxrt.StringFromLiteral("")
		}
		decoded := haxe__Unserializer_unescape(raw[self.pos : self.pos+length])
		self.pos += length
		self.stringCache = append(self.stringCache, decoded)
		return decoded
	case 'R':
		index := haxe__Unserializer_readUInt(self)
		if index < 0 || index >= len(self.stringCache) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid string reference"))
			return hxrt.StringFromLiteral("")
		}
		return self.stringCache[index]
	case 'x':
		hxrt.Throw(haxe__Unserializer_readValue(self))
		return nil
	case 'l':
		list := New_haxe__ds__List()
		cacheIndex := len(self.cache)
		self.cache = append(self.cache, list)
		for {
			if self.pos >= len(raw) {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized list"))
				return list
			}
			if raw[self.pos] == 'h' {
				self.pos++
				break
			}
			list.add(haxe__Unserializer_readValue(self))
			self.cache[cacheIndex] = list
		}
		return list
	case 'b':
		stringMap := New_haxe__ds__StringMap()
		cacheIndex := len(self.cache)
		self.cache = append(self.cache, stringMap)
		for {
			if self.pos >= len(raw) {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized StringMap"))
				return stringMap
			}
			if raw[self.pos] == 'h' {
				self.pos++
				break
			}
			keyAny := haxe__Unserializer_readValue(self)
			key := *hxrt.StdString(keyAny)
			hxrt.StringMapSet(stringMap.h, hxrt.StringFromLiteral(key), haxe__Unserializer_readValue(self))
			self.cache[cacheIndex] = stringMap
		}
		return stringMap
	case 'q':
		intMap := New_haxe__ds__IntMap()
		cacheIndex := len(self.cache)
		self.cache = append(self.cache, intMap)
		for {
			if self.pos >= len(raw) {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized IntMap"))
				return intMap
			}
			if raw[self.pos] == 'h' {
				self.pos++
				break
			}
			if raw[self.pos] != ':' {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized IntMap format"))
				return intMap
			}
			self.pos++
			key := haxe__Unserializer_readDigits(self)
			hxrt.IntMapSet(intMap.h, key, haxe__Unserializer_readValue(self))
			self.cache[cacheIndex] = intMap
		}
		return intMap
	case 'M':
		objectMap := New_haxe__ds__ObjectMap()
		cacheIndex := len(self.cache)
		self.cache = append(self.cache, objectMap)
		for {
			if self.pos >= len(raw) {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized ObjectMap"))
				return objectMap
			}
			if raw[self.pos] == 'h' {
				self.pos++
				break
			}
			key := haxe__Unserializer_readValue(self)
			hxrt.ObjectMapSet(objectMap.h, key, haxe__Unserializer_readValue(self))
			self.cache[cacheIndex] = objectMap
		}
		return objectMap
	case 'a':
		arr := hxrt.NewArray()
		cacheIndex := len(self.cache)
		self.cache = append(self.cache, arr)
		for {
			if self.pos >= len(raw) {
				hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized array"))
				return arr
			}
			if raw[self.pos] == 'h' {
				self.pos++
				break
			}
			if raw[self.pos] == 'u' {
				self.pos++
				skip := haxe__Unserializer_readUInt(self)
				for i := 0; i < skip; i++ {
					arr.Push(nil)
				}
				self.cache[cacheIndex] = arr
				continue
			}
			arr.Push(haxe__Unserializer_readValue(self))
			self.cache[cacheIndex] = arr
		}
		return arr
	case 'o':
		obj := map[string]any{}
		cacheIndex := len(self.cache)
		self.cache = append(self.cache, obj)
		hxrt_unserializerReadObjectFields(self, obj, "Invalid serialized object")
		self.cache[cacheIndex] = obj
		return obj
	case 'C':
		classNameAny := haxe__Unserializer_readValue(self)
		requestedName := *hxrt.StdString(classNameAny)
		resolvedClass := hxrt_unserializerResolveClass(self, requestedName)
		if resolvedClass == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Class not found " + requestedName))
			return nil
		}
		className, ok := hxrt_unserializerResolvedClassName(resolvedClass)
		if !ok || className == "" {
			hxrt.Throw(hxrt.StringFromLiteral("Class not found " + requestedName))
			return nil
		}
		instance, ok := hxrt_unserializerCreateClassInstance(className)
		if !ok {
			hxrt.Throw(hxrt.StringFromLiteral("Class not found " + className))
			return nil
		}
		cacheIndex := len(self.cache)
		self.cache = append(self.cache, instance)
		if custom, ok := instance.(interface{ hxUnserialize(*haxe__Unserializer) }); ok {
			custom.hxUnserialize(self)
		} else if custom, ok := instance.(interface{ HxUnserialize(*haxe__Unserializer) }); ok {
			custom.HxUnserialize(self)
		} else {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid custom data"))
			return instance
		}
		if self.pos >= len(raw) || raw[self.pos] != 'g' {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid custom data"))
			return instance
		}
		self.pos++
		self.cache[cacheIndex] = instance
		return instance
	case 'A':
		nameAny := haxe__Unserializer_readValue(self)
		requestedName := *hxrt.StdString(nameAny)
		resolvedClass := hxrt_unserializerResolveClass(self, requestedName)
		if resolvedClass == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Class not found " + requestedName))
			return nil
		}
		return resolvedClass
	case 'B':
		nameAny := haxe__Unserializer_readValue(self)
		requestedName := *hxrt.StdString(nameAny)
		resolvedEnum := hxrt_unserializerResolveEnum(self, requestedName)
		if resolvedEnum == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Enum not found " + requestedName))
			return nil
		}
		return resolvedEnum
	case 'c':
		classNameAny := haxe__Unserializer_readValue(self)
		requestedName := *hxrt.StdString(classNameAny)
		resolvedClass := hxrt_unserializerResolveClass(self, requestedName)
		if resolvedClass == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Class not found " + requestedName))
			return nil
		}
		className, ok := hxrt_unserializerResolvedClassName(resolvedClass)
		if !ok || className == "" {
			hxrt.Throw(hxrt.StringFromLiteral("Class not found " + requestedName))
			return nil
		}
		instance, ok := hxrt_unserializerCreateClassInstance(className)
		if !ok {
			hxrt.Throw(hxrt.StringFromLiteral("Class not found " + className))
			return nil
		}
		cacheIndex := len(self.cache)
		self.cache = append(self.cache, instance)
		hxrt_unserializerReadObjectFields(self, instance, "Invalid serialized class")
		self.cache[cacheIndex] = instance
		return instance
	case 'j':
		enumNameAny := haxe__Unserializer_readValue(self)
		requestedEnumName := *hxrt.StdString(enumNameAny)
		resolvedEnum := hxrt_unserializerResolveEnum(self, requestedEnumName)
		if resolvedEnum == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Enum not found " + requestedEnumName))
			return nil
		}
		enumName, ok := hxrt_unserializerResolvedEnumName(resolvedEnum)
		if !ok || enumName == "" {
			hxrt.Throw(hxrt.StringFromLiteral("Enum not found " + requestedEnumName))
			return nil
		}
		if self.pos >= len(raw) || raw[self.pos] != ':' {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized enum index"))
			return nil
		}
		self.pos++
		enumIndex := haxe__Unserializer_readDigits(self)
		if self.pos >= len(raw) || raw[self.pos] != ':' {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized enum"))
			return nil
		}
		self.pos++
		argCount := haxe__Unserializer_readUInt(self)
		if argCount < 0 {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized enum arity"))
			return nil
		}
		args := make([]any, 0, argCount)
		for i := 0; i < argCount; i++ {
			args = append(args, haxe__Unserializer_readValue(self))
		}
		enumValue, ok := hxrt_unserializerCreateEnumInstance(enumName, "", enumIndex, true, args)
		if !ok {
			hxrt.Throw(hxrt.StringFromLiteral("Unknown enum index " + enumName + "@" + strconv.Itoa(enumIndex)))
			return nil
		}
		self.cache = append(self.cache, enumValue)
		return enumValue
	case 'w':
		enumNameAny := haxe__Unserializer_readValue(self)
		constructorAny := haxe__Unserializer_readValue(self)
		requestedEnumName := *hxrt.StdString(enumNameAny)
		resolvedEnum := hxrt_unserializerResolveEnum(self, requestedEnumName)
		if resolvedEnum == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Enum not found " + requestedEnumName))
			return nil
		}
		enumName, ok := hxrt_unserializerResolvedEnumName(resolvedEnum)
		if !ok || enumName == "" {
			hxrt.Throw(hxrt.StringFromLiteral("Enum not found " + requestedEnumName))
			return nil
		}
		constructorName := *hxrt.StdString(constructorAny)
		if self.pos >= len(raw) || raw[self.pos] != ':' {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized enum"))
			return nil
		}
		self.pos++
		argCount := haxe__Unserializer_readUInt(self)
		if argCount < 0 {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized enum arity"))
			return nil
		}
		args := make([]any, 0, argCount)
		for i := 0; i < argCount; i++ {
			args = append(args, haxe__Unserializer_readValue(self))
		}
		enumValue, ok := hxrt_unserializerCreateEnumInstance(enumName, constructorName, 0, false, args)
		if !ok {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized enum"))
			return nil
		}
		self.cache = append(self.cache, enumValue)
		return enumValue
	case 'r':
		index := haxe__Unserializer_readUInt(self)
		if index < 0 || index >= len(self.cache) {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid object reference"))
			return nil
		}
		return self.cache[index]
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid serialized token"))
		return nil
	}
}

func haxe__Unserializer_run(source *string) any {
	if source == nil {
		return nil
	}
	decoder := New_haxe__Unserializer(source)
	return decoder.unserialize()
}
