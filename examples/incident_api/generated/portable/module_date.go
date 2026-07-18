package main

import "examples_incident_api_portable/hxrt"

type I_Date interface {
	getTime() float64
	getHours() int
	getMinutes() int
	getSeconds() int
	getFullYear() int
	getMonth() int
	getDate() int
	getDay() int
	getUTCHours() int
	getUTCMinutes() int
	getUTCSeconds() int
	getUTCFullYear() int
	getUTCMonth() int
	getUTCDate() int
	getUTCDay() int
	getTimezoneOffset() int
	toString() *string
	localParts() *hxrt.DateParts
	utcParts() *hxrt.DateParts
}

type Date struct {
	__hx_this I_Date
	ms        float64
}

func New_Date(year int, month int, day int, hour int, min int, sec int) *Date {
	self := &Date{}
	self.__hx_this = self
	self.ms = hxrt.DateLocalTime(year, month, day, hour, min, sec)
	return self
}

func (self *Date) getTime() float64 {
	return self.ms
}

func (self *Date) getHours() int {
	return hxrt.DateLocalParts(self.ms).Hours
}

func (self *Date) getMinutes() int {
	return hxrt.DateLocalParts(self.ms).Minutes
}

func (self *Date) getSeconds() int {
	return hxrt.DateLocalParts(self.ms).Seconds
}

func (self *Date) getFullYear() int {
	return hxrt.DateLocalParts(self.ms).FullYear
}

func (self *Date) getMonth() int {
	return hxrt.DateLocalParts(self.ms).Month
}

func (self *Date) getDate() int {
	return hxrt.DateLocalParts(self.ms).Date
}

func (self *Date) getDay() int {
	return hxrt.DateLocalParts(self.ms).Day
}

func (self *Date) getUTCHours() int {
	return hxrt.DateUTCParts(self.ms).Hours
}

func (self *Date) getUTCMinutes() int {
	return hxrt.DateUTCParts(self.ms).Minutes
}

func (self *Date) getUTCSeconds() int {
	return hxrt.DateUTCParts(self.ms).Seconds
}

func (self *Date) getUTCFullYear() int {
	return hxrt.DateUTCParts(self.ms).FullYear
}

func (self *Date) getUTCMonth() int {
	return hxrt.DateUTCParts(self.ms).Month
}

func (self *Date) getUTCDate() int {
	return hxrt.DateUTCParts(self.ms).Date
}

func (self *Date) getUTCDay() int {
	return hxrt.DateUTCParts(self.ms).Day
}

func (self *Date) getTimezoneOffset() int {
	return hxrt.DateTimezoneOffset(self.ms)
}

func (self *Date) toString() *string {
	return hxrt.StdString(hxrt.DateFormatLocal(self.ms))
}

func (self *Date) localParts() *hxrt.DateParts {
	return hxrt.DateLocalParts(self.ms)
}

func (self *Date) utcParts() *hxrt.DateParts {
	return hxrt.DateUTCParts(self.ms)
}

func (self *Date) String() string {
	return *self.__hx_this.toString()
}

func Date_fromMilliseconds(value float64) *Date {
	result := New_Date(1970, 0, 1, 0, 0, 0)
	result.ms = value
	return result
}

func Date_fromString(s *string) *Date {
	return Date_fromMilliseconds(hxrt.DateParse(s))
}

func Date_fromTime(t float64) *Date {
	return Date_fromMilliseconds(t)
}

func Date_now() *Date {
	return Date_fromMilliseconds(hxrt.DateNow())
}
