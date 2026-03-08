package main

import (
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"io"
	"math"
	"reflect"
	"regexp"
	"snapshot/hxrt"
	"sort"
	"strconv"
	"strings"
	"time"
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
	hxrt.Println(basic.execute(func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["name"] = hxrt.StringFromLiteral("ok")
		return hx_obj_1
	}(), nil))
	cond := New_haxe__Template(hxrt.StringFromLiteral("::if enabled::yes::else::no::end::"))
	hxrt.Println(cond.execute(func() map[string]any {
		hx_obj_2 := map[string]any{}
		hx_obj_2["enabled"] = true
		return hx_obj_2
	}(), nil))
	hxrt.Println(cond.execute(func() map[string]any {
		hx_obj_3 := map[string]any{}
		hx_obj_3["enabled"] = false
		return hx_obj_3
	}(), nil))
	loop := New_haxe__Template(hxrt.StringFromLiteral("::foreach items::::__current__::::end::"))
	hxrt.Println(loop.execute(func() map[string]any {
		hx_obj_4 := map[string]any{}
		hx_obj_4["items"] = []*string{hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("c")}
		return hx_obj_4
	}(), nil))
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

type haxe__ds__IntMap struct {
	h map[int]any
}

type haxe__ds__StringMap struct {
	h map[string]any
}

type haxe__ds__ObjectMap struct {
	h map[any]any
}

type haxe__ds__EnumValueMap struct {
	h map[any]any
}

type haxe__ds__List struct {
	items  []any
	length int
}

func New_haxe__ds__IntMap() *haxe__ds__IntMap {
	return &haxe__ds__IntMap{h: map[int]any{}}
}

func (self *haxe__ds__IntMap) set(key any, value any) {
	resolvedKey := hxrt.IntFromNullableAny(key)
	self.h[resolvedKey] = value
}

func (self *haxe__ds__IntMap) get(key any) any {
	resolvedKey := hxrt.IntFromNullableAny(key)
	value := self.h[resolvedKey]
	return value
}

func (self *haxe__ds__IntMap) exists(key any) bool {
	resolvedKey := hxrt.IntFromNullableAny(key)
	_, ok := self.h[resolvedKey]
	return ok
}

func (self *haxe__ds__IntMap) remove(key any) bool {
	resolvedKey := hxrt.IntFromNullableAny(key)
	_, ok := self.h[resolvedKey]
	delete(self.h, resolvedKey)
	return ok
}

func (self *haxe__ds__IntMap) keys() map[string]any {
	keys := make([]int, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() int { key := keys[index]; index++; return key }
	return iter
}

func (self *haxe__ds__IntMap) iterator() map[string]any {
	keys := make([]int, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__IntMap) keyValueIterator() map[string]any {
	keys := make([]int, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": key, "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__IntMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__IntMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__IntMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__IntMap) clear() {
	self.h = map[int]any{}
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	return &haxe__ds__StringMap{h: map[string]any{}}
}

func (self *haxe__ds__StringMap) set(key any, value any) {
	self.h[*hxrt.StdString(key)] = value
}

func (self *haxe__ds__StringMap) get(key any) any {
	value := self.h[*hxrt.StdString(key)]
	return value
}

func (self *haxe__ds__StringMap) exists(key any) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	return ok
}

func (self *haxe__ds__StringMap) remove(key any) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	delete(self.h, *hxrt.StdString(key))
	return ok
}

func (self *haxe__ds__StringMap) keys() map[string]any {
	keys := make([]string, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() *string { key := keys[index]; index++; return hxrt.StringFromLiteral(key) }
	return iter
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	keys := make([]string, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	keys := make([]string, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": hxrt.StringFromLiteral(key), "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__StringMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__StringMap) clear() {
	self.h = map[string]any{}
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	return &haxe__ds__ObjectMap{h: map[any]any{}}
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func (self *haxe__ds__ObjectMap) keys() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return key }
	return iter
}

func (self *haxe__ds__ObjectMap) iterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__ObjectMap) keyValueIterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": key, "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__ObjectMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__ObjectMap) clear() {
	self.h = map[any]any{}
}

func New_haxe__ds__EnumValueMap() *haxe__ds__EnumValueMap {
	return &haxe__ds__EnumValueMap{h: map[any]any{}}
}

func (self *haxe__ds__EnumValueMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__EnumValueMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__EnumValueMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__EnumValueMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func (self *haxe__ds__EnumValueMap) keys() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return key }
	return iter
}

func (self *haxe__ds__EnumValueMap) iterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() any { key := keys[index]; index++; return self.h[key] }
	return iter
}

func (self *haxe__ds__EnumValueMap) keyValueIterator() map[string]any {
	keys := make([]any, 0, len(self.h))
	for key := range self.h {
		keys = append(keys, key)
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(keys) }
	iter["next"] = func() map[string]any {
		key := keys[index]
		index++
		return map[string]any{"key": key, "value": self.h[key]}
	}
	return iter
}

func (self *haxe__ds__EnumValueMap) copyIMap() haxe__IMap {
	copied := New_haxe__ds__EnumValueMap()
	for key, value := range self.h {
		copied.h[key] = value
	}
	return copied
}

func (self *haxe__ds__EnumValueMap) toString() *string {
	return hxrt.StringFromLiteral("{}")
}

func (self *haxe__ds__EnumValueMap) clear() {
	self.h = map[any]any{}
}

func New_haxe__ds__List() *haxe__ds__List {
	return &haxe__ds__List{items: []any{}, length: 0}
}

func (self *haxe__ds__List) add(item any) {
	self.items = append(self.items, item)
	self.length = len(self.items)
}

func (self *haxe__ds__List) push(item any) {
	self.items = append([]any{item}, self.items...)
	self.length = len(self.items)
}

func (self *haxe__ds__List) pop() any {
	if len(self.items) == 0 {
		return nil
	}
	head := self.items[0]
	self.items = self.items[1:]
	self.length = len(self.items)
	return head
}

func (self *haxe__ds__List) first() any {
	if len(self.items) == 0 {
		return nil
	}
	return self.items[0]
}

func (self *haxe__ds__List) last() any {
	size := len(self.items)
	if size == 0 {
		return nil
	}
	return self.items[(size - 1)]
}

type Std struct {
}

func _UnicodeString__UnicodeString_Impl__get_length(value any) int {
	return len([]rune(*hxrt.StdString(value)))
}

func _UnicodeString__UnicodeString_Impl__charAt(value any, index int) *string {
	if index < 0 {
		return hxrt.StringFromLiteral("")
	}
	runes := []rune(*hxrt.StdString(value))
	if index >= len(runes) {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(string(runes[index]))
}

func _UnicodeString__UnicodeString_Impl__charCodeAt(value any, index int) any {
	if index < 0 {
		return nil
	}
	runes := []rune(*hxrt.StdString(value))
	if index >= len(runes) {
		return nil
	}
	return int(runes[index])
}

func _UnicodeString__UnicodeString_Impl__substring(value any, startIndex int, endIndex ...int) *string {
	runes := []rune(*hxrt.StdString(value))
	end := len(runes)
	if len(endIndex) > 0 {
		end = endIndex[0]
	}
	if startIndex < 0 {
		startIndex = 0
	}
	if end < 0 {
		end = 0
	}
	if startIndex == end {
		return hxrt.StringFromLiteral("")
	}
	if startIndex > end {
		startIndex, end = end, startIndex
	}
	if startIndex > len(runes) {
		return hxrt.StringFromLiteral("")
	}
	if end > len(runes) {
		end = len(runes)
	}
	return hxrt.StringFromLiteral(string(runes[startIndex:end]))
}

func _UnicodeString__UnicodeString_Impl__substr(value any, pos int, lengthArgs ...int) *string {
	runes := []rune(*hxrt.StdString(value))
	unicodeLength := len(runes)
	if pos < 0 {
		pos = unicodeLength + pos
		if pos < 0 {
			pos = 0
		}
	}
	if pos > unicodeLength {
		return hxrt.StringFromLiteral("")
	}
	if len(lengthArgs) == 0 {
		return hxrt.StringFromLiteral(string(runes[pos:]))
	}
	lengthValue := lengthArgs[0]
	end := unicodeLength
	if lengthValue < 0 {
		end = unicodeLength + lengthValue
	} else {
		end = pos + lengthValue
	}
	if end < pos {
		return hxrt.StringFromLiteral("")
	}
	if end > unicodeLength {
		end = unicodeLength
	}
	return hxrt.StringFromLiteral(string(runes[pos:end]))
}

func _UnicodeString__UnicodeString_Impl__indexOf(value any, str *string, startIndex ...int) int {
	runes := []rune(*hxrt.StdString(value))
	needle := []rune(*hxrt.StdString(str))
	start := 0
	if len(startIndex) > 0 {
		start = startIndex[0]
		if start < 0 {
			start = len(runes) + start
		}
	}
	if start < 0 {
		start = 0
	}
	if len(needle) == 0 {
		if start > len(runes) {
			return len(runes)
		}
		return start
	}
	if start > len(runes) || len(needle) > len(runes) {
		return -1
	}
	for i := start; i+len(needle) <= len(runes); i++ {
		matched := true
		for j := 0; j < len(needle); j++ {
			if runes[i+j] != needle[j] {
				matched = false
				break
			}
		}
		if matched {
			return i
		}
	}
	return -1
}

func _UnicodeString__UnicodeString_Impl__lastIndexOf(value any, str *string, startIndex ...int) int {
	runes := []rune(*hxrt.StdString(value))
	needle := []rune(*hxrt.StdString(str))
	if len(needle) == 0 {
		if len(startIndex) == 0 {
			return len(runes)
		}
		start := startIndex[0]
		if start < 0 {
			start = 0
		}
		if start > len(runes) {
			return len(runes)
		}
		return start
	}
	start := len(runes)
	if len(startIndex) > 0 {
		start = startIndex[0]
		if start < 0 {
			start = 0
		}
	}
	limit := start + len(needle)
	if limit > len(runes) {
		limit = len(runes)
	}
	for i := limit - len(needle); i >= 0; i-- {
		matched := true
		for j := 0; j < len(needle); j++ {
			if runes[i+j] != needle[j] {
				matched = false
				break
			}
		}
		if matched {
			return i
		}
	}
	return -1
}

func _UnicodeString__UnicodeString_Impl__validate(value *haxe__io__Bytes, encoding *haxe__io__Encoding) bool {
	if haxe__io__resolveEncodingTag(encoding) == 1 {
		hxrt.Throw(hxrt.StringFromLiteral("UnicodeString.validate: RawNative encoding is not supported"))
		return false
	}
	raw := hxrt_haxeBytesToRaw(value)
	pos := 0
	max := len(raw)
	for pos < max {
		c := int(raw[pos])
		pos++
		if c < 0x80 {
			continue
		} else if c < 0xC2 {
			return false
		} else if c < 0xE0 {
			if pos+1 > max {
				return false
			}
			c2 := int(raw[pos])
			pos++
			if c2 < 0x80 || c2 > 0xBF {
				return false
			}
		} else if c < 0xF0 {
			if pos+2 > max {
				return false
			}
			c2 := int(raw[pos])
			pos++
			if c == 0xE0 {
				if c2 < 0xA0 || c2 > 0xBF {
					return false
				}
			} else if c2 < 0x80 || c2 > 0xBF {
				return false
			}
			c3 := int(raw[pos])
			pos++
			if c3 < 0x80 || c3 > 0xBF {
				return false
			}
			c = (c << 16) | (c2 << 8) | c3
			if 0xEDA080 <= c && c <= 0xEDBFBF {
				return false
			}
		} else if c > 0xF4 {
			return false
		} else {
			if pos+3 > max {
				return false
			}
			c2 := int(raw[pos])
			pos++
			if c == 0xF0 {
				if c2 < 0x90 || c2 > 0xBF {
					return false
				}
			} else if c == 0xF4 {
				if c2 < 0x80 || c2 > 0x8F {
					return false
				}
			} else if c2 < 0x80 || c2 > 0xBF {
				return false
			}
			c3 := int(raw[pos])
			pos++
			if c3 < 0x80 || c3 > 0xBF {
				return false
			}
			c4 := int(raw[pos])
			pos++
			if c4 < 0x80 || c4 > 0xBF {
				return false
			}
		}
	}
	return true
}

type Date struct {
	value time.Time
}

func Date_fromString(source *string) *Date {
	raw := *hxrt.StdString(source)
	parsed, err := time.ParseInLocation("2006-01-02 15:04:05", raw, time.Local)
	if err != nil {
		parsedDateOnly, errDateOnly := time.ParseInLocation("2006-01-02", raw, time.Local)
		if errDateOnly == nil {
			parsed = parsedDateOnly
		} else {
			parsed = time.Unix(0, 0)
		}
	}
	return &Date{value: parsed}
}

func Date_now() *Date {
	return &Date{value: time.Now()}
}

func Date_fromTime(ms float64) *Date {
	nanos := int64(ms * 1e6)
	return &Date{value: time.Unix(0, nanos).In(time.Local)}
}

func (self *Date) getFullYear() int {
	return self.value.Year()
}

func (self *Date) getMonth() int {
	return int(self.value.Month()) - 1
}

func (self *Date) getDate() int {
	return self.value.Day()
}

func (self *Date) getDay() int {
	return int(self.value.Weekday())
}

func (self *Date) getHours() int {
	return self.value.Hour()
}

func (self *Date) getMinutes() int {
	return self.value.Minute()
}

func (self *Date) getSeconds() int {
	return self.value.Second()
}

func (self *Date) getTime() float64 {
	return float64(self.value.UnixNano()) / 1e6
}

type Math struct {
}

func Math_floor(value float64) int {
	return int(math.Floor(value))
}

func Math_ceil(value float64) int {
	return int(math.Ceil(value))
}

func Math_round(value float64) int {
	return int(math.Floor(value + 0.5))
}

func Math_abs(value float64) float64 {
	return math.Abs(value)
}

func Math_isNaN(value float64) bool {
	return math.IsNaN(value)
}

func Math_isFinite(value float64) bool {
	return !math.IsInf(value, 0)
}

func Math_min(a float64, b float64) float64 {
	return math.Min(a, b)
}

func Math_max(a float64, b float64) float64 {
	return math.Max(a, b)
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

func Reflect_field(obj any, field *string) any {
	if obj == nil {
		return nil
	}
	key := *hxrt.StdString(field)
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

var Xml_Element int = 0

var Xml_PCData int = 1

var Xml_CData int = 2

var Xml_Comment int = 3

var Xml_DocType int = 4

var Xml_ProcessingInstruction int = 5

var Xml_Document int = 6

type Xml struct {
	nodeType       int
	nodeName       *string
	nodeValue      *string
	parent         *Xml
	children       []*Xml
	attributeMap   map[string]string
	attributeOrder []string
}

func _Xml__XmlType_Impl__toString(value int) *string {
	switch value {
	case Xml_Element:
		return hxrt.StringFromLiteral("Element")
	case Xml_PCData:
		return hxrt.StringFromLiteral("PCData")
	case Xml_CData:
		return hxrt.StringFromLiteral("CData")
	case Xml_Comment:
		return hxrt.StringFromLiteral("Comment")
	case Xml_DocType:
		return hxrt.StringFromLiteral("DocType")
	case Xml_ProcessingInstruction:
		return hxrt.StringFromLiteral("ProcessingInstruction")
	case Xml_Document:
		return hxrt.StringFromLiteral("Document")
	default:
		return hxrt.StringFromLiteral("XmlType")
	}
}

func New_Xml(nodeType int) *Xml {
	return &Xml{nodeType: nodeType, children: []*Xml{}, attributeMap: map[string]string{}, attributeOrder: []string{}}
}

func Xml_createElement(name *string) *Xml {
	xml := New_Xml(Xml_Element)
	xml.nodeName = name
	return xml
}

func Xml_createPCData(data *string) *Xml {
	xml := New_Xml(Xml_PCData)
	xml.nodeValue = data
	return xml
}

func Xml_createCData(data *string) *Xml {
	xml := New_Xml(Xml_CData)
	xml.nodeValue = data
	return xml
}

func Xml_createComment(data *string) *Xml {
	xml := New_Xml(Xml_Comment)
	xml.nodeValue = data
	return xml
}

func Xml_createDocType(data *string) *Xml {
	xml := New_Xml(Xml_DocType)
	xml.nodeValue = data
	return xml
}

func Xml_createProcessingInstruction(data *string) *Xml {
	xml := New_Xml(Xml_ProcessingInstruction)
	xml.nodeValue = data
	return xml
}

func Xml_createDocument() *Xml {
	return New_Xml(Xml_Document)
}

func Xml_ensureElementType(self *Xml) {
	if self.nodeType != Xml_Document && self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element or Document but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
	}
}

func Xml_parse(source *string) *Xml {
	return haxe__xml__Parser_parse(source)
}

func (self *Xml) get(att *string) *string {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return nil
	}
	key := *hxrt.StdString(att)
	value, ok := self.attributeMap[key]
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(value)
}

func (self *Xml) set(att *string, value *string) {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return
	}
	key := *hxrt.StdString(att)
	if _, ok := self.attributeMap[key]; !ok {
		self.attributeOrder = append(self.attributeOrder, key)
	}
	self.attributeMap[key] = *hxrt.StdString(value)
}

func (self *Xml) remove(att *string) {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return
	}
	key := *hxrt.StdString(att)
	delete(self.attributeMap, key)
	filtered := make([]string, 0, len(self.attributeOrder))
	for _, existing := range self.attributeOrder {
		if existing != key {
			filtered = append(filtered, existing)
		}
	}
	self.attributeOrder = filtered
}

func (self *Xml) exists(att *string) bool {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return false
	}
	_, ok := self.attributeMap[*hxrt.StdString(att)]
	return ok
}

func (self *Xml) attributes() map[string]any {
	if self.nodeType != Xml_Element {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Bad node type, expected Element but found "), _Xml__XmlType_Impl__toString(self.nodeType)))
		return map[string]any{}
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(self.attributeOrder) }
	iter["next"] = func() *string { key := self.attributeOrder[index]; index++; return hxrt.StringFromLiteral(key) }
	return iter
}

func (self *Xml) iterator() map[string]any {
	Xml_ensureElementType(self)
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(self.children) }
	iter["next"] = func() *Xml { child := self.children[index]; index++; return child }
	return iter
}

func (self *Xml) elements() map[string]any {
	Xml_ensureElementType(self)
	matches := make([]*Xml, 0, len(self.children))
	for _, child := range self.children {
		if child.nodeType == Xml_Element {
			matches = append(matches, child)
		}
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(matches) }
	iter["next"] = func() *Xml { child := matches[index]; index++; return child }
	return iter
}

func (self *Xml) elementsNamed(name *string) map[string]any {
	Xml_ensureElementType(self)
	wanted := *hxrt.StdString(name)
	matches := make([]*Xml, 0, len(self.children))
	for _, child := range self.children {
		if child.nodeType == Xml_Element && *hxrt.StdString(child.nodeName) == wanted {
			matches = append(matches, child)
		}
	}
	index := 0
	iter := map[string]any{}
	iter["hasNext"] = func() bool { return index < len(matches) }
	iter["next"] = func() *Xml { child := matches[index]; index++; return child }
	return iter
}

func (self *Xml) firstChild() *Xml {
	Xml_ensureElementType(self)
	if len(self.children) == 0 {
		return nil
	}
	return self.children[0]
}

func (self *Xml) firstElement() *Xml {
	Xml_ensureElementType(self)
	for _, child := range self.children {
		if child.nodeType == Xml_Element {
			return child
		}
	}
	return nil
}

func (self *Xml) addChild(x *Xml) {
	Xml_ensureElementType(self)
	if x == nil {
		return
	}
	if x.parent != nil {
		x.parent.removeChild(x)
	}
	self.children = append(self.children, x)
	x.parent = self
}

func (self *Xml) removeChild(x *Xml) bool {
	Xml_ensureElementType(self)
	for i, child := range self.children {
		if child == x {
			self.children = append(self.children[:i], self.children[i+1:]...)
			x.parent = nil
			return true
		}
	}
	return false
}

func (self *Xml) insertChild(x *Xml, pos int) {
	Xml_ensureElementType(self)
	if x == nil {
		return
	}
	if x.parent != nil {
		x.parent.removeChild(x)
	}
	if pos < 0 {
		pos = 0
	}
	if pos > len(self.children) {
		pos = len(self.children)
	}
	self.children = append(self.children, nil)
	copy(self.children[pos+1:], self.children[pos:])
	self.children[pos] = x
	x.parent = self
}

func (self *Xml) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return haxe__xml__Printer_print(self)
}

type haxe__crypto__Base64 struct {
}

type haxe__crypto__Md5 struct {
}

type haxe__crypto__Sha1 struct {
}

type haxe__crypto__Sha224 struct {
}

type haxe__crypto__Sha256 struct {
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

func haxe__crypto__Base64_encode(bytes *haxe__io__Bytes, complement ...bool) *string {
	useComplement := true
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	encoded := base64.StdEncoding.EncodeToString(hxrt_haxeBytesToRaw(bytes))
	if !useComplement {
		encoded = strings.TrimRight(encoded, "=")
	}
	return hxrt.StringFromLiteral(encoded)
}

func haxe__crypto__Base64_decode(value *string, complement ...bool) *haxe__io__Bytes {
	useComplement := true
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	rawValue := *hxrt.StdString(value)
	if useComplement {
		rawValue = strings.TrimRight(rawValue, "=")
	}
	decoded, err := base64.RawStdEncoding.DecodeString(rawValue)
	if err != nil {
		decoded, err = base64.StdEncoding.DecodeString(*hxrt.StdString(value))
		if err != nil {
			hxrt.Throw(err)
			return &haxe__io__Bytes{b: []int{}, length: 0}
		}
	}
	return hxrt_rawToHaxeBytes(decoded)
}

func haxe__crypto__Base64_urlEncode(bytes *haxe__io__Bytes, complement ...bool) *string {
	useComplement := false
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	encoded := base64.RawURLEncoding.EncodeToString(hxrt_haxeBytesToRaw(bytes))
	if useComplement {
		missing := len(encoded) % 4
		if missing != 0 {
			encoded = (encoded + strings.Repeat("=", (4-missing)))
		}
	}
	return hxrt.StringFromLiteral(encoded)
}

func haxe__crypto__Base64_urlDecode(value *string, complement ...bool) *haxe__io__Bytes {
	rawValue := *hxrt.StdString(value)
	decoded, err := base64.RawURLEncoding.DecodeString(strings.TrimRight(rawValue, "="))
	if err != nil {
		hxrt.Throw(err)
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return hxrt_rawToHaxeBytes(decoded)
}

func haxe__crypto__Md5_encode(value *string) *string {
	sum := md5.Sum([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Md5_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := md5.Sum(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha1_encode(value *string) *string {
	sum := sha1.Sum([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha1_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha1.Sum(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha224_encode(value *string) *string {
	sum := sha256.Sum224([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha224_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha256.Sum224(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha256_encode(value *string) *string {
	sum := sha256.Sum256([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha256_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha256.Sum256(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

type haxe__ds__BalancedTree struct {
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}

type haxe__io__StringInput struct {
}

type haxe__xml__Parser struct {
}

type haxe__xml__Printer struct {
}

func haxe__xml__Parser_parse(source *string, strict ...bool) *Xml {
	raw := *hxrt.StdString(source)
	doc := Xml_createDocument()
	stack := []*Xml{doc}
	decoder := xml.NewDecoder(strings.NewReader(raw))
	for {
		tokenStart := decoder.InputOffset()
		token, err := decoder.Token()
		tokenEnd := decoder.InputOffset()
		if err == io.EOF {
			break
		}
		if err != nil {
			hxrt.Throw(err)
			return Xml_createDocument()
		}
		current := stack[len(stack)-1]
		switch value := token.(type) {
		case xml.StartElement:
			node := Xml_createElement(hxrt.StringFromLiteral(value.Name.Local))
			for _, attr := range value.Attr {
				node.set(hxrt.StringFromLiteral(attr.Name.Local), hxrt.StringFromLiteral(attr.Value))
			}
			current.addChild(node)
			stack = append(stack, node)
		case xml.EndElement:
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		case xml.CharData:
			text := string([]byte(value))
			if len(text) != 0 {
				tokenSource := raw[tokenStart:tokenEnd]
				if strings.HasPrefix(tokenSource, "<![CDATA[") && strings.HasSuffix(tokenSource, "]]>") {
					current.addChild(Xml_createCData(hxrt.StringFromLiteral(text)))
				} else {
					current.addChild(Xml_createPCData(hxrt.StringFromLiteral(text)))
				}
			}
		case xml.Comment:
			current.addChild(Xml_createComment(hxrt.StringFromLiteral(string([]byte(value)))))
		case xml.Directive:
			directive := strings.TrimSpace(string([]byte(value)))
			upper := strings.ToUpper(directive)
			if strings.HasPrefix(upper, "DOCTYPE") {
				directive = strings.TrimSpace(directive[len("DOCTYPE"):])
			}
			current.addChild(Xml_createDocType(hxrt.StringFromLiteral(directive)))
		case xml.ProcInst:
			payload := value.Target
			if len(value.Inst) != 0 {
				payload += " " + string(value.Inst)
			}
			current.addChild(Xml_createProcessingInstruction(hxrt.StringFromLiteral(strings.TrimSpace(payload))))
		}
	}
	return doc
}

func haxe__xml__Printer_escapeText(value string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;").Replace(value)
}

func haxe__xml__Printer_escapeAttr(value string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "\"", "&quot;").Replace(value)
}

func haxe__xml__Printer_hasChildren(value *Xml) bool {
	for _, child := range value.children {
		switch child.nodeType {
		case Xml_Element, Xml_PCData:
			return true
		case Xml_CData, Xml_Comment:
			if len(strings.TrimLeft(*hxrt.StdString(child.nodeValue), " \n\r\t")) != 0 {
				return true
			}
		}
	}
	return false
}

func haxe__xml__Printer_writeNode(output *strings.Builder, value *Xml, tabs string, pretty bool) {
	newline := func() {
		if pretty {
			output.WriteString("\n")
		}
	}
	switch value.nodeType {
	case Xml_CData:
		output.WriteString(tabs + "<![CDATA[")
		output.WriteString(*hxrt.StdString(value.nodeValue))
		output.WriteString("]]>")
		newline()
	case Xml_Comment:
		commentContent := strings.NewReplacer("\n", "", "\r", "", "\t", "").Replace(*hxrt.StdString(value.nodeValue))
		output.WriteString(tabs)
		output.WriteString(strings.TrimSpace("<!--" + commentContent + "-->"))
		newline()
	case Xml_Document:
		for _, child := range value.children {
			haxe__xml__Printer_writeNode(output, child, tabs, pretty)
		}
	case Xml_Element:
		output.WriteString(tabs + "<")
		output.WriteString(*hxrt.StdString(value.nodeName))
		for _, attribute := range value.attributeOrder {
			output.WriteString(" " + attribute + "=\"")
			output.WriteString(haxe__xml__Printer_escapeAttr(value.attributeMap[attribute]))
			output.WriteString("\"")
		}
		if haxe__xml__Printer_hasChildren(value) {
			output.WriteString(">")
			newline()
			childTabs := tabs
			if pretty {
				childTabs = tabs + "\t"
			}
			for _, child := range value.children {
				haxe__xml__Printer_writeNode(output, child, childTabs, pretty)
			}
			output.WriteString(tabs + "</")
			output.WriteString(*hxrt.StdString(value.nodeName))
			output.WriteString(">")
			newline()
		} else {
			output.WriteString("/>")
			newline()
		}
	case Xml_PCData:
		nodeValue := *hxrt.StdString(value.nodeValue)
		if len(nodeValue) != 0 {
			output.WriteString(tabs + haxe__xml__Printer_escapeText(nodeValue))
			newline()
		}
	case Xml_ProcessingInstruction:
		output.WriteString("<?" + *hxrt.StdString(value.nodeValue) + "?>")
		newline()
	case Xml_DocType:
		output.WriteString("<!DOCTYPE " + *hxrt.StdString(value.nodeValue) + ">")
		newline()
	}
}

func haxe__xml__Printer_print(value *Xml, pretty ...bool) *string {
	if value == nil {
		return hxrt.StringFromLiteral("")
	}
	usePretty := false
	if len(pretty) > 0 {
		usePretty = pretty[0]
	}
	var output strings.Builder
	haxe__xml__Printer_writeNode(&output, value, "", usePretty)
	return hxrt.StringFromLiteral(output.String())
}

type haxe__zip__Compress struct {
}

type haxe__zip__Uncompress struct {
}

func haxe__zip__Compress_run(src *haxe__io__Bytes, level int) *haxe__io__Bytes {
	raw := hxrt_haxeBytesToRaw(src)
	var buffer bytes.Buffer
	writer, err := zlib.NewWriterLevel(&buffer, level)
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	if _, err := writer.Write(raw); err != nil {
		_ = writer.Close()
		hxrt.Throw(err)
		return nil
	}
	if err := writer.Close(); err != nil {
		hxrt.Throw(err)
		return nil
	}
	return hxrt_rawToHaxeBytes(buffer.Bytes())
}

func haxe__zip__Uncompress_run(src *haxe__io__Bytes, bufsize ...int) *haxe__io__Bytes {
	raw := hxrt_haxeBytesToRaw(src)
	reader, err := zlib.NewReader(bytes.NewReader(raw))
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	defer reader.Close()
	decoded, err := io.ReadAll(reader)
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	return hxrt_rawToHaxeBytes(decoded)
}

type sys__FileSystem struct {
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
	case "haxe.Template":
		return hxrt_typeCallAny(New_haxe__Template, args)
	case "haxe._Template.ExprCursor":
		return hxrt_typeCallAny(New_haxe___Template__ExprCursor, args)
	case "haxe._Template.TokenCursor":
		return hxrt_typeCallAny(New_haxe___Template__TokenCursor, args)
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
	case "haxe._Template.ExprCursor":
		return &haxe___Template__ExprCursor{}, true
	case "haxe._Template.TokenCursor":
		return &haxe___Template__TokenCursor{}, true
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
	case *haxe__Template:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Template")}
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
	case "haxe.Template":
		return nil
	case "haxe._Template.ExprCursor":
		return nil
	case "haxe._Template.TokenCursor":
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

func Type_getClassFields(c any) []*string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return []*string{}
	}
	switch className {
	case "haxe.Template":
		return []*string{hxrt.StringFromLiteral("addValues"), hxrt.StringFromLiteral("anyArrayToSlice"), hxrt.StringFromLiteral("compareValues"), hxrt.StringFromLiteral("divideValues"), hxrt.StringFromLiteral("expr_float"), hxrt.StringFromLiteral("expr_int"), hxrt.StringFromLiteral("expr_splitter"), hxrt.StringFromLiteral("expr_trim"), hxrt.StringFromLiteral("globals"), hxrt.StringFromLiteral("isSpaceOnly"), hxrt.StringFromLiteral("joinDynamicArgs"), hxrt.StringFromLiteral("kwdEnd"), hxrt.StringFromLiteral("multiplyValues"), hxrt.StringFromLiteral("parseFloatLiteral"), hxrt.StringFromLiteral("parseIntLiteral"), hxrt.StringFromLiteral("peekExprToken"), hxrt.StringFromLiteral("peekToken"), hxrt.StringFromLiteral("popExprToken"), hxrt.StringFromLiteral("popToken"), hxrt.StringFromLiteral("splitter"), hxrt.StringFromLiteral("subtractValues"), hxrt.StringFromLiteral("trimExprToken"), hxrt.StringFromLiteral("valueAsBool"), hxrt.StringFromLiteral("valueAsFloat")}
	case "haxe._Template.ExprCursor":
		return []*string{}
	case "haxe._Template.TokenCursor":
		return []*string{}
	case "haxe.iterators.StringIterator":
		return []*string{}
	case "haxe.iterators.StringKeyValueIterator":
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
	case "haxe.Template":
		return []*string{hxrt.StringFromLiteral("context"), hxrt.StringFromLiteral("execute"), hxrt.StringFromLiteral("expr"), hxrt.StringFromLiteral("macros"), hxrt.StringFromLiteral("makeConst"), hxrt.StringFromLiteral("makeExpr"), hxrt.StringFromLiteral("makeExpr2"), hxrt.StringFromLiteral("makePath"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("parse"), hxrt.StringFromLiteral("parseBlock"), hxrt.StringFromLiteral("parseExpr"), hxrt.StringFromLiteral("parseTokens"), hxrt.StringFromLiteral("popStackValue"), hxrt.StringFromLiteral("resolve"), hxrt.StringFromLiteral("run"), hxrt.StringFromLiteral("skipSpaces"), hxrt.StringFromLiteral("stack")}
	case "haxe._Template.ExprCursor":
		return []*string{hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("tokens")}
	case "haxe._Template.TokenCursor":
		return []*string{hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("tokens")}
	case "haxe.iterators.StringIterator":
		return []*string{hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s")}
	case "haxe.iterators.StringKeyValueIterator":
		return []*string{hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s")}
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
	case "haxe.Template":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Template.ExprCursor":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Template.TokenCursor":
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

func Type_getEnumConstructs(e any) []*string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return []*string{}
	}
	switch enumName {
	case "ValueType":
		return []*string{hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown")}
	case "haxe._Template.TemplateExpr":
		return []*string{hxrt.StringFromLiteral("OpVar"), hxrt.StringFromLiteral("OpExpr"), hxrt.StringFromLiteral("OpIf"), hxrt.StringFromLiteral("OpStr"), hxrt.StringFromLiteral("OpBlock"), hxrt.StringFromLiteral("OpForeach"), hxrt.StringFromLiteral("OpMacro")}
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
	case *haxe___Template__TemplateExpr:
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
	case "haxe._Template.TemplateExpr":
		return []any{}
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

func haxe__Template_anyArrayToSlice_runtime(value any) []any {
	if value == nil {
		return nil
	}
	rv := reflect.ValueOf(value)
	if !rv.IsValid() {
		return nil
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil
	}
	out := make([]any, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		item := rv.Index(i)
		if item.CanInterface() {
			out[i] = item.Interface()
		}
	}
	return out
}

func Reflect_getProperty(obj any, field *string) any {
	return Reflect_field(obj, field)
}

func Reflect_isObject(obj any) bool {
	if obj == nil {
		return false
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return false
	}
	switch rv.Kind() {
	case reflect.Pointer, reflect.Interface:
		if rv.IsNil() {
			return false
		}
		return Reflect_isObject(rv.Elem().Interface())
	case reflect.Struct, reflect.Map:
		return true
	default:
		return false
	}
}

func Reflect_callMethod(obj any, funcValue any, args []any) any {
	_ = obj
	if funcValue == nil {
		return nil
	}
	fn := reflect.ValueOf(funcValue)
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil
	}
	callArgs := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		callArgs = append(callArgs, reflect.ValueOf(arg))
	}
	results := fn.Call(callArgs)
	if len(results) == 0 {
		return nil
	}
	return results[0].Interface()
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

func (self *EReg) split(source *string) []*string {
	raw := *hxrt.StdString(source)
	if self == nil || self.regex == nil {
		return []*string{hxrt.StringFromLiteral(raw)}
	}
	parts := self.regex.Split(raw, -1)
	out := make([]*string, 0, len(parts))
	for _, part := range parts {
		out = append(out, hxrt.StringFromLiteral(part))
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
	case "haxe___Template__ExprCursor":
		return "haxe._Template.ExprCursor", true
	case "haxe___Template__TokenCursor":
		return "haxe._Template.TokenCursor", true
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
	case "haxe._Template.ExprCursor":
		return true
	case "haxe._Template.TokenCursor":
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
	case "haxe._Template.ExprCursor":
		instance := &haxe___Template__ExprCursor{}
		hxrt_unserializerBindSelf(instance)
		return instance, true
	case "haxe._Template.TokenCursor":
		instance := &haxe___Template__TokenCursor{}
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

func hxrt_serializerWriteObjectMapToken(self *haxe__Serializer, entries map[any]any) {
	hxrt_serializerAppend(self, "M")
	for key, value := range entries {
		hxrt_serializerWriteValue(self, key)
		hxrt_serializerWriteValue(self, value)
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

func hxrt_serializerTryDsListStruct(self *haxe__Serializer, ref reflect.Value) bool {
	if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != "haxe__ds__List" {
		return false
	}
	itemsField := ref.FieldByName("items")
	if !itemsField.IsValid() || itemsField.Kind() != reflect.Slice {
		return false
	}
	items := make([]any, 0, itemsField.Len())
	for i := 0; i < itemsField.Len(); i++ {
		item, ok := hxrt_serializerReflectAny(itemsField.Index(i))
		if !ok {
			return false
		}
		items = append(items, item)
	}
	hxrt_serializerWriteListToken(self, items)
	return true
}

func hxrt_serializerTryDsStringMapStruct(self *haxe__Serializer, ref reflect.Value) bool {
	if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != "haxe__ds__StringMap" {
		return false
	}
	mapField := ref.FieldByName("h")
	if !mapField.IsValid() || mapField.Kind() != reflect.Map {
		return false
	}
	entries := map[string]any{}
	for _, key := range mapField.MapKeys() {
		if key.Kind() != reflect.String {
			return false
		}
		value, ok := hxrt_serializerReflectAny(mapField.MapIndex(key))
		if !ok {
			return false
		}
		entries[key.String()] = value
	}
	hxrt_serializerWriteStringMapToken(self, entries)
	return true
}

func hxrt_serializerTryDsIntMapStruct(self *haxe__Serializer, ref reflect.Value) bool {
	if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != "haxe__ds__IntMap" {
		return false
	}
	mapField := ref.FieldByName("h")
	if !mapField.IsValid() || mapField.Kind() != reflect.Map {
		return false
	}
	entries := map[int]any{}
	for _, key := range mapField.MapKeys() {
		var intKey int
		switch key.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intKey = int(key.Int())
		default:
			return false
		}
		value, ok := hxrt_serializerReflectAny(mapField.MapIndex(key))
		if !ok {
			return false
		}
		entries[intKey] = value
	}
	hxrt_serializerWriteIntMapToken(self, entries)
	return true
}

func hxrt_serializerTryDsObjectMapStruct(self *haxe__Serializer, ref reflect.Value) bool {
	if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != "haxe__ds__ObjectMap" {
		return false
	}
	mapField := ref.FieldByName("h")
	if !mapField.IsValid() || mapField.Kind() != reflect.Map {
		return false
	}
	entries := map[any]any{}
	for _, key := range mapField.MapKeys() {
		keyAny, ok := hxrt_serializerReflectAny(key)
		if !ok {
			return false
		}
		valueAny, ok := hxrt_serializerReflectAny(mapField.MapIndex(key))
		if !ok {
			return false
		}
		entries[keyAny] = valueAny
	}
	hxrt_serializerWriteObjectMapToken(self, entries)
	return true
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
	valueField := ref.FieldByName("value")
	if !valueField.IsValid() || !valueField.CanAddr() {
		return false
	}
	fieldType := valueField.Type()
	if fieldType.PkgPath() != "time" || fieldType.Name() != "Time" {
		return false
	}
	timeAny := reflect.NewAt(fieldType, unsafe.Pointer(valueField.UnsafeAddr())).Elem().Interface()
	timeValue, ok := timeAny.(time.Time)
	if !ok {
		return false
	}
	ms := float64(timeValue.UnixNano()) / 1000000.0
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
	if hxrt_serializerTryDsListStruct(self, ref) {
		return true
	}
	if hxrt_serializerTryDsStringMapStruct(self, ref) {
		return true
	}
	if hxrt_serializerTryDsIntMapStruct(self, ref) {
		return true
	}
	if hxrt_serializerTryDsObjectMapStruct(self, ref) {
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
	case *haxe__ds__List:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			if hxrt_serializerTrackRef(self, current) {
				return
			}
			hxrt_serializerWriteListToken(self, current.items)
		}
		return
	case *haxe__ds__StringMap:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			if hxrt_serializerTrackRef(self, current) {
				return
			}
			hxrt_serializerWriteStringMapToken(self, current.h)
		}
		return
	case *haxe__ds__IntMap:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			if hxrt_serializerTrackRef(self, current) {
				return
			}
			hxrt_serializerWriteIntMapToken(self, current.h)
		}
		return
	case *haxe__ds__ObjectMap:
		if current == nil {
			hxrt_serializerAppend(self, "n")
		} else {
			if hxrt_serializerTrackRef(self, current) {
				return
			}
			hxrt_serializerWriteObjectMapToken(self, current.h)
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
			list.items = append(list.items, haxe__Unserializer_readValue(self))
			list.length = len(list.items)
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
			stringMap.h[key] = haxe__Unserializer_readValue(self)
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
			intMap.h[key] = haxe__Unserializer_readValue(self)
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
			objectMap.h[key] = haxe__Unserializer_readValue(self)
			self.cache[cacheIndex] = objectMap
		}
		return objectMap
	case 'a':
		arr := make([]any, 0)
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
					arr = append(arr, nil)
				}
				self.cache[cacheIndex] = arr
				continue
			}
			arr = append(arr, haxe__Unserializer_readValue(self))
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
