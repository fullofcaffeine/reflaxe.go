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
	"path/filepath"
	"reflect"
	"snapshot/hxrt"
	"strings"
	"time"
)

func main() {
	d := Date_fromString(hxrt.StringFromLiteral("2024-02-03 04:05:06"))
	_ = d
	hxrt.Println(d.getFullYear())
	hxrt.Println(d.getMonth())
	hxrt.Println(d.getDate())
	hxrt.Println(d.getHours())
	hxrt.Println(haxe__io__Path_join([]*string{hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("c.txt")}))
	p := New_haxe__io__Path(hxrt.StringFromLiteral("/tmp/demo.txt"))
	_ = p
	hxrt.Println(p.dir)
	hxrt.Println(p.file)
	hxrt.Println(p.ext)
}

type haxe__io__Encoding struct {
}

type haxe__io__Input struct {
}

type haxe__io__Output struct {
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

func New_haxe__io__Input() *haxe__io__Input {
	return &haxe__io__Input{}
}

func New_haxe__io__Output() *haxe__io__Output {
	return &haxe__io__Output{}
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
	raw := []byte(*hxrt.StdString(value))
	converted := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		converted[i] = int(raw[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}
}

func (self *haxe__io__Bytes) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.BytesToString(self.b)
}

func (self *haxe__io__Bytes) get(pos int) int {
	return self.b[pos]
}

func (self *haxe__io__Bytes) set(pos int, value int) {
	self.b[pos] = value
	self.__hx_rawValid = false
}

func New_haxe__io__BytesBuffer() *haxe__io__BytesBuffer {
	return &haxe__io__BytesBuffer{b: []int{}}
}

func (self *haxe__io__BytesBuffer) addByte(value int) {
	self.b = append(self.b, (value & 255))
}

func (self *haxe__io__BytesBuffer) add(src *haxe__io__Bytes) {
	if src == nil {
		return
	}
	self.b = append(self.b, src.b...)
}

func (self *haxe__io__BytesBuffer) addString(value *string, encoding ...*haxe__io__Encoding) {
	self.add(haxe__io__Bytes_ofString(value))
}

func (self *haxe__io__BytesBuffer) getBytes() *haxe__io__Bytes {
	copied := hxrt.BytesClone(self.b)
	return &haxe__io__Bytes{b: copied, length: len(copied)}
}

func (self *haxe__io__BytesBuffer) get_length() int {
	return len(self.b)
}

type Std struct {
}

type StringTools struct {
}

func StringTools_trim(value *string) *string {
	return hxrt.StringFromLiteral(strings.TrimSpace(*hxrt.StdString(value)))
}

func StringTools_startsWith(value *string, prefix *string) bool {
	return strings.HasPrefix(*hxrt.StdString(value), *hxrt.StdString(prefix))
}

func StringTools_replace(value *string, sub *string, by *string) *string {
	return hxrt.StringFromLiteral(strings.ReplaceAll(*hxrt.StdString(value), *hxrt.StdString(sub), *hxrt.StdString(by)))
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

func (self *Date) getFullYear() int {
	return self.value.Year()
}

func (self *Date) getMonth() int {
	return int(self.value.Month()) - 1
}

func (self *Date) getDate() int {
	return self.value.Day()
}

func (self *Date) getHours() int {
	return self.value.Hour()
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

type Xml struct {
	raw *string
}

func Xml_parse(source *string) *Xml {
	return haxe__xml__Parser_parse(source)
}

func (self *Xml) toString() *string {
	if self == nil || self.raw == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(*self.raw)
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

type haxe__io__BytesInput struct {
}

type haxe__io__BytesOutput struct {
}

type haxe__io__Eof struct {
}

type haxe__io__Error struct {
}

type haxe__io__Path struct {
	dir       *string
	file      *string
	ext       *string
	backslash bool
}

func New_haxe__io__Path(path *string) *haxe__io__Path {
	raw := *hxrt.StdString(path)
	dir := filepath.Dir(raw)
	if dir == "." {
		dir = ""
	}
	base := filepath.Base(raw)
	dotExt := filepath.Ext(base)
	file := base
	if dotExt != "" {
		file = strings.TrimSuffix(base, dotExt)
	}
	ext := strings.TrimPrefix(dotExt, ".")
	return &haxe__io__Path{dir: hxrt.StringFromLiteral(dir), file: hxrt.StringFromLiteral(file), ext: hxrt.StringFromLiteral(ext), backslash: strings.Contains(raw, "\\")}
}

func haxe__io__Path_join(parts []*string) *string {
	if len(parts) == 0 {
		return hxrt.StringFromLiteral("")
	}
	joined := filepath.ToSlash(filepath.Join(hxrt.StringSlice(parts)...))
	return hxrt.StringFromLiteral(joined)
}

type haxe__io__StringInput struct {
}

type haxe__xml__Parser struct {
}

type haxe__xml__Printer struct {
}

func haxe__xml__Parser_parse(source *string, strict ...bool) *Xml {
	raw := *hxrt.StdString(source)
	decoder := xml.NewDecoder(strings.NewReader(raw))
	for {
		_, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			hxrt.Throw(err)
			return &Xml{raw: hxrt.StringFromLiteral("")}
		}
	}
	return &Xml{raw: hxrt.StringFromLiteral(raw)}
}

func haxe__xml__Printer_print(value *Xml, pretty ...bool) *string {
	if value == nil || value.raw == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(*value.raw)
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
