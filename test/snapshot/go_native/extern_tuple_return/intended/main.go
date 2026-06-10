package main

import (
	"snapshot/hxrt"
	"strconv"
	"time"
)

type I_AtoiResult interface {
}

type AtoiResult struct {
	__hx_this I_AtoiResult
	n         int
	err       *go___Error
}

func New_AtoiResult(n int, err *go___Error) *AtoiResult {
	self := &AtoiResult{}
	self.__hx_this = self
	self.n = n
	self.err = err
	return self
}

func main() {
	ok := func() *AtoiResult {
		hx_tuple_1, hx_tuple_2 := strconv.Atoi(*hxrt.StdString(hxrt.StringFromLiteral("12")))
		return New_AtoiResult(hx_tuple_1, func(err error) *go___Error {
			if err == nil {
				return nil
			}
			return New_go___Error(hxrt.StringFromLiteral(err.Error()))
		}(hx_tuple_2))
	}()
	bad := func() *AtoiResult {
		hx_tuple_3, hx_tuple_4 := strconv.Atoi(*hxrt.StdString(hxrt.StringFromLiteral("nope")))
		return New_AtoiResult(hx_tuple_3, func(err error) *go___Error {
			if err == nil {
				return nil
			}
			return New_go___Error(hxrt.StringFromLiteral(err.Error()))
		}(hx_tuple_4))
	}()
	zone := func() *TimeZoneResult {
		hx_tuple_5, hx_tuple_6 := time.Now().Zone()
		return New_TimeZoneResult(hxrt.StdString(hx_tuple_5), hx_tuple_6)
	}()
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("atoi.ok="), ok.n), hxrt.StringFromLiteral(":")), hxrt.StdString((ok.err == nil))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("atoi.err="), hxrt.StdString((bad.err != nil))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("zone.typed="), hxrt.StdString((hxrt.StringLengthStringPtr(zone.name) >= 0))), hxrt.StringFromLiteral(":")), hxrt.StdString((zone.offset >= -86400))))
}

type I_TimeZoneResult interface {
}

type TimeZoneResult struct {
	__hx_this I_TimeZoneResult
	name      *string
	offset    int
}

func New_TimeZoneResult(name *string, offset int) *TimeZoneResult {
	self := &TimeZoneResult{}
	self.__hx_this = self
	self.name = name
	self.offset = offset
	return self
}
