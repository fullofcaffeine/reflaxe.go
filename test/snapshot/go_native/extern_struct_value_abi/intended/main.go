package main

import (
	"fmt"
	"image"
	"snapshot/hxrt"
	"time"
)

func main() {
	left := &image.Point{}
	left.X = 20
	right := &image.Point{}
	right.Y = 22
	sum := func() *image.Point {
		hx_extern_value_1 := left.Add(*right)
		return &hx_extern_value_1
	}()
	fmt.Println(*hxrt.StdString(hxrt.StringConcatAny(hxrt.StringFromLiteral("point="), int((hxrt.Int32Wrap(sum.X) + hxrt.Int32Wrap(sum.Y))))))
	first := func() *ParseResult {
		hx_tuple_2, hx_tuple_3 := time.Parse(*hxrt.StdString(hxrt.StringFromLiteral("2006-01-02")), *hxrt.StdString(hxrt.StringFromLiteral("2026-08-24")))
		return New_ParseResult(&hx_tuple_2, func(err interface{ Error() string }) *go___Error {
			if err == nil {
				return nil
			}
			return New_go___Error(hxrt.StringFromLiteral(err.Error()))
		}(hx_tuple_3))
	}()
	second := func() *ParseResult {
		hx_tuple_4, hx_tuple_5 := time.Parse(*hxrt.StdString(hxrt.StringFromLiteral("2006-01-02")), *hxrt.StdString(hxrt.StringFromLiteral("2026-08-24")))
		return New_ParseResult(&hx_tuple_4, func(err interface{ Error() string }) *go___Error {
			if err == nil {
				return nil
			}
			return New_go___Error(hxrt.StringFromLiteral(err.Error()))
		}(hx_tuple_5))
	}()
	fmt.Println(*hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("time.equal="), hxrt.StdString((((first.error == nil) && (second.error == nil)) && first.value.Equal(*second.value))))))
}

type I_ParseResult interface {
}

type ParseResult struct {
	__hx_this I_ParseResult
	value     *time.Time
	error     *go___Error
}

func New_ParseResult(value *time.Time, error *go___Error) *ParseResult {
	self := &ParseResult{}
	self.__hx_this = self
	self.value = value
	self.error = error
	return self
}
