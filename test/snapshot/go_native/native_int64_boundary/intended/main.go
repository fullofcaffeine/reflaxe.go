package main

import (
	"database/sql"
	"errors"
	"snapshot/hxrt"
	"strconv"
)

func main() {
	text := hxrt.StringFromLiteral("-9223372036854775808")
	parsed := go__result_fromValueError(strconv.ParseInt(*hxrt.StdString(text), 10, 64))
	if !go__result_isOk__int64_03d22364(parsed) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("ParseInt failed")))
	}
	parsedTuple := func() *ParsedInt64 {
		hx_tuple_1, hx_tuple_2 := strconv.ParseInt(*hxrt.StdString(text), 10, 64)
		return New_ParsedInt64(hx_tuple_1, func(err interface{ Error() string }) *go___Error {
			if err == nil {
				return nil
			}
			return New_go___Error(hxrt.StringFromLiteral(err.Error()))
		}(hx_tuple_2))
	}()
	if parsedTuple.error != nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("ParseInt tuple failed")))
	}
	if !hxrt.StringEqualStringPtr(hxrt.StdString(strconv.FormatInt(parsedTuple.value, 10)), text) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("tuple value changed")))
	}
	row := &sql.NullInt64{}
	row.Int64 = go__result_unwrap__int64_03d22364(parsed)
	row.Valid = true
	actual := hxrt.StdString(strconv.FormatInt(row.Int64, 10))
	if !row.Valid || !hxrt.StringEqualStringPtr(actual, text) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("int64 field round trip failed")))
	}
	hxrt.Println(any(actual))
	text_1 := hxrt.StringFromLiteral("9223372036854775807")
	parsed_1 := go__result_fromValueError(strconv.ParseInt(*hxrt.StdString(text_1), 10, 64))
	if !go__result_isOk__int64_03d22364(parsed_1) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("ParseInt failed")))
	}
	parsedTuple_1 := func() *ParsedInt64 {
		hx_tuple_3, hx_tuple_4 := strconv.ParseInt(*hxrt.StdString(text_1), 10, 64)
		return New_ParsedInt64(hx_tuple_3, func(err interface{ Error() string }) *go___Error {
			if err == nil {
				return nil
			}
			return New_go___Error(hxrt.StringFromLiteral(err.Error()))
		}(hx_tuple_4))
	}()
	if parsedTuple_1.error != nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("ParseInt tuple failed")))
	}
	if !hxrt.StringEqualStringPtr(hxrt.StdString(strconv.FormatInt(parsedTuple_1.value, 10)), text_1) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("tuple value changed")))
	}
	row_1 := &sql.NullInt64{}
	row_1.Int64 = go__result_unwrap__int64_03d22364(parsed_1)
	row_1.Valid = true
	actual_1 := hxrt.StdString(strconv.FormatInt(row_1.Int64, 10))
	if !row_1.Valid || !hxrt.StringEqualStringPtr(actual_1, text_1) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("int64 field round trip failed")))
	}
	hxrt.Println(any(actual_1))
	text_2 := hxrt.StringFromLiteral("9007199254740993")
	parsed_2 := go__result_fromValueError(strconv.ParseInt(*hxrt.StdString(text_2), 10, 64))
	if !go__result_isOk__int64_03d22364(parsed_2) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("ParseInt failed")))
	}
	parsedTuple_2 := func() *ParsedInt64 {
		hx_tuple_5, hx_tuple_6 := strconv.ParseInt(*hxrt.StdString(text_2), 10, 64)
		return New_ParsedInt64(hx_tuple_5, func(err interface{ Error() string }) *go___Error {
			if err == nil {
				return nil
			}
			return New_go___Error(hxrt.StringFromLiteral(err.Error()))
		}(hx_tuple_6))
	}()
	if parsedTuple_2.error != nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("ParseInt tuple failed")))
	}
	if !hxrt.StringEqualStringPtr(hxrt.StdString(strconv.FormatInt(parsedTuple_2.value, 10)), text_2) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("tuple value changed")))
	}
	row_2 := &sql.NullInt64{}
	row_2.Int64 = go__result_unwrap__int64_03d22364(parsed_2)
	row_2.Valid = true
	actual_2 := hxrt.StdString(strconv.FormatInt(row_2.Int64, 10))
	if !row_2.Valid || !hxrt.StringEqualStringPtr(actual_2, text_2) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("int64 field round trip failed")))
	}
	hxrt.Println(any(actual_2))
}

type I_ParsedInt64 interface {
}

type ParsedInt64 struct {
	__hx_this I_ParsedInt64
	value     int64
	error     *go___Error
}

func (self *ParsedInt64) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_ParsedInt64(value int64, error *go___Error) *ParsedInt64 {
	return hx_init_ParsedInt64(hx_alloc_ParsedInt64(nil), value, error)
}

func hx_alloc_ParsedInt64(owner any) *ParsedInt64 {
	self := &ParsedInt64{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_ParsedInt64)
	}
	return self
}

func hx_init_ParsedInt64(self *ParsedInt64, value int64, error *go___Error) *ParsedInt64 {
	self.value = value
	self.error = error
	return self
}

func (self *ParsedInt64) HxClassName() string {
	return "ParsedInt64"
}

func go__result_fromValueError(value any, err error) *go___Result {
	if err != nil {
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral(err.Error())))
	}
	return New_go___Result(value, nil)
}

func go__result_ok__int64_03d22364(value int64) *go___Result {
	return New_go___Result(value, nil)
}

func go__result_failure__int64_03d22364(message *string) *go___Result {
	return New_go___Result(nil, New_go___Error(message))
}

func go__result_valueError__int64_03d22364(result *go___Result) (int64, error) {
	var zero int64
	if result == nil {
		return zero, errors.New("nil go.Result")
	}
	if result.errorValue != nil {
		return zero, errors.New(*hxrt.StdString(result.errorValue.message))
	}
	if result.value == nil {
		return zero, nil
	}
	return result.value.(int64), nil
}

func go__result_isOk__int64_03d22364(result *go___Result) bool {
	_, err := go__result_valueError__int64_03d22364(result)
	return (err == nil)
}

func go__result_isErr__int64_03d22364(result *go___Result) bool {
	_, err := go__result_valueError__int64_03d22364(result)
	return (err != nil)
}

func go__result_unwrap__int64_03d22364(result *go___Result) int64 {
	value, err := go__result_valueError__int64_03d22364(result)
	if err != nil {
		hxrt.Throw(hxrt.StringFromLiteral(err.Error()))
		var zero int64
		return zero
	}
	return value
}

func go__result_error__int64_03d22364(result *go___Result) *string {
	_, err := go__result_valueError__int64_03d22364(result)
	if err == nil {
		return nil
	}
	return hxrt.StringFromLiteral(err.Error())
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
