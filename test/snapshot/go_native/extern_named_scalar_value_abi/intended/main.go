package main

import (
	"fmt"
	"snapshot/hxrt"
	"time"
)

func main() {
	duration := func() *ParseDurationResult {
		hx_tuple_1, hx_tuple_2 := time.ParseDuration(*hxrt.StdString(hxrt.StringFromLiteral("1s")))
		return New_ParseDurationResult(&hx_tuple_1, func(err error) *go___Error {
			if err == nil {
				return nil
			}
			return New_go___Error(hxrt.StringFromLiteral(err.Error()))
		}(hx_tuple_2))
	}()
	start := func() *time.Time {
		hx_extern_value_3 := time.Unix(41, 0)
		return &hx_extern_value_3
	}()
	fmt.Println(*hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ok="), hxrt.StdString(((duration.error == nil) && (func() *time.Time {
		hx_extern_value_4 := start.Add(*duration.value)
		return &hx_extern_value_4
	}() != nil))))))
}

type I_ParseDurationResult interface {
}

type ParseDurationResult struct {
	__hx_this I_ParseDurationResult
	value     *time.Duration
	error     *go___Error
}

func New_ParseDurationResult(value *time.Duration, error *go___Error) *ParseDurationResult {
	self := &ParseDurationResult{}
	self.__hx_this = self
	self.value = value
	self.error = error
	return self
}
