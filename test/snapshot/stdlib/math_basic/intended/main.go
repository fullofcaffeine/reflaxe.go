package main

import (
	"math"
	"path/filepath"
	"snapshot/hxrt"
	"strings"
	"time"
)

func main() {
	x := 3.8
	_ = x
	hxrt.Println(Math_floor(x))
	hxrt.Println(Math_ceil(x))
	hxrt.Println(Math_round(x))
	hxrt.Println(Math_round(x))
	hxrt.Println(Math_abs(x))
	hxrt.Println(Math_min(x, 2.1))
	hxrt.Println(Math_max(x, 2.1))
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

type Xml struct {
}

type EReg struct {
}

type haxe__Serializer struct {
}

type haxe__Unserializer struct {
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

type haxe__zip__Compress struct {
}

type haxe__zip__Uncompress struct {
}

type sys__FileSystem struct {
}

type sys__net__Host struct {
}

type sys__net__Socket struct {
}
