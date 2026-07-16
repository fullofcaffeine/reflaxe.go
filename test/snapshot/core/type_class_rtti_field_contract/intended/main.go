package main

import (
	"reflect"
	"snapshot/hxrt"
)

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

var Demo___rtti *string = hxrt.StringFromLiteral("<class path=\"Demo\" params=\"\" module=\"Main\">\n\t<value public=\"1\" expr=\"1\" line=\"4\" static=\"1\">\n\t\t<x path=\"Int\"/>\n\t\t<meta><m n=\":value\"><e>1</e></m></meta>\n\t</value>\n\t<meta>\n\t\t<m n=\":directlyUsed\"/>\n\t\t<m n=\":keep\"/>\n\t\t<m n=\":rtti\"/>\n\t</meta>\n</class>")

var Demo_value int = 1

func main() {
	fields := Type_getClassFields(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Demo")})
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("fields.len="), len(fields)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fields.has_rtti="), hxrt.StdString(Lambda_has(func() map[string]any {
		hx_lambda_source_10 := fields
		hx_lambda_wrapped_11 := map[string]any{}
		hx_lambda_wrapped_11["iterator"] = func() map[string]any {
			hx_lambda_index_14 := 0
			hx_lambda_iterator_map_12 := map[string]any{}
			hx_lambda_iterator_map_12["hasNext"] = func() bool {
				return (hx_lambda_index_14 < len(hx_lambda_source_10))
			}
			hx_lambda_iterator_map_12["next"] = func() any {
				hx_lambda_value_13 := hx_lambda_source_10[hx_lambda_index_14]
				hx_lambda_index_14 = (hx_lambda_index_14 + 1)
				return hx_lambda_value_13
			}
			return hx_lambda_iterator_map_12
		}
		return hx_lambda_wrapped_11
	}(), hxrt.StringFromLiteral("__rtti")))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw.rtti.null="), hxrt.StdString(hxrt.AnyEqualsNull(Reflect_field(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Demo")}, hxrt.StringFromLiteral("__rtti"))))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw.meta.null="), hxrt.StdString(hxrt.AnyEqualsNull(Reflect_field(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Demo")}, hxrt.StringFromLiteral("__meta__"))))))
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
	case "Demo":
		return nil, false
	case "Lambda":
		return nil, false
	case "Main":
		return nil, false
	case "haxe.Int64Helper":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
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
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
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
	case "Demo":
		return nil
	case "Lambda":
		return nil
	case "Main":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
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

func Type_getClassFields(c any) []*string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return []*string{}
	}
	switch className {
	case "Demo":
		return []*string{hxrt.StringFromLiteral("__rtti"), hxrt.StringFromLiteral("value")}
	case "Lambda":
		return []*string{hxrt.StringFromLiteral("has")}
	case "Main":
		return []*string{hxrt.StringFromLiteral("main")}
	case "haxe.Int64Helper":
		return []*string{}
	case "haxe._Int32.Int32_Impl_":
		return []*string{}
	case "haxe._Int64.Int64_Impl_":
		return []*string{}
	case "haxe._Int64.___Int64":
		return []*string{}
	default:
		return []*string{}
	}
}

func Type_getInstanceFields(c any) []*string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return []*string{}
	}
	switch className {
	case "Demo":
		return []*string{}
	case "Lambda":
		return []*string{}
	case "Main":
		return []*string{}
	case "haxe.Int64Helper":
		return []*string{}
	case "haxe._Int32.Int32_Impl_":
		return []*string{}
	case "haxe._Int64.Int64_Impl_":
		return []*string{}
	case "haxe._Int64.___Int64":
		return []*string{hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low")}
	default:
		return []*string{}
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
	case "Demo":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Lambda":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
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
	default:
		return nil
	}
}

func Type_createInstance(cl any, args []any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, args)
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

func Type_createEnum(e any, constr *string, params []any) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, params)
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params []any) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, params)
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
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) []*string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return []*string{}
	}
	switch enumName {
	case "ValueType":
		return []*string{hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown")}
	default:
		return []*string{}
	}
}

func Type_enumParameters(e any) []any {
	if hxrt.AnyEqualsNull(e) {
		return []any{}
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return []any{}
		}
		out := make([]any, len(value.params))
		copy(out, value.params)
		return out
	default:
		return []any{}
	}
}

func Type_allEnums(e any) []any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return []any{}
	}
	switch enumName {
	case "ValueType":
		return []any{ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown}
	default:
		return []any{}
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
	case "Demo":
		switch key {
		case "__rtti":
			return Demo___rtti, true
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}
