package main

import "snapshot/hxrt"

type I_sys__ssl__Certificate interface {
	subject(field *string) *string
	issuer(field *string) *string
	next() *sys__ssl__Certificate
	add(pem *string)
	addDER(der *haxe__io__Bytes)
	get_commonName() *string
	get_altNames() *hxrt.Array
	get_notBefore() *Date
	get_notAfter() *Date
}

type sys__ssl__Certificate struct {
	__hx_this  I_sys__ssl__Certificate
	handle     *hxrt.SslCertificate
	commonName *string
	altNames   *hxrt.Array
	notBefore  *Date
	notAfter   *Date
}

func New_sys__ssl__Certificate(handle *hxrt.SslCertificate) *sys__ssl__Certificate {
	self := &sys__ssl__Certificate{}
	self.__hx_this = self
	self.handle = handle
	return self
}

func (self *sys__ssl__Certificate) subject(field *string) *string {
	return hxrt.StdString(hxrt.SslCertSubject(self.handle, field))
}

func (self *sys__ssl__Certificate) issuer(field *string) *string {
	return hxrt.StdString(hxrt.SslCertIssuer(self.handle, field))
}

func (self *sys__ssl__Certificate) next() *sys__ssl__Certificate {
	nextHandle := hxrt.SslCertNext(self.handle)
	var hx_if_44 *sys__ssl__Certificate
	if nextHandle == nil {
		hx_if_44 = nil
	} else {
		hx_if_44 = New_sys__ssl__Certificate(nextHandle)
	}
	return hx_if_44
}

func (self *sys__ssl__Certificate) add(pem *string) {
	hxrt.SslCertAddPEM(self.handle, pem)
}

func (self *sys__ssl__Certificate) addDER(der *haxe__io__Bytes) {
	values := hxrt.NewArray()
	_g := 0
	_g1 := der.length
	for _g < _g1 {
		hx_post_45 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_45
		values.Push(der.b[index])
	}
	hxrt.SslCertAddDERValues(self.handle, func(hx_lambda_raw_47 []any) []int {
		hx_lambda_out_48 := make([]int, 0, len(hx_lambda_raw_47))
		for _, hx_lambda_item_49 := range hx_lambda_raw_47 {
			hx_lambda_out_48 = append(hx_lambda_out_48, func(hx_value_50 any) int {
				if hx_value_50 == nil {
					var hx_zero_51 int
					return hx_zero_51
				}
				return hx_value_50.(int)
			}(hx_lambda_item_49))
		}
		return hx_lambda_out_48
	}(values.Values()))
}

func (self *sys__ssl__Certificate) get_commonName() *string {
	return hxrt.StdString(hxrt.SslCertCommonName(self.handle))
}

func (self *sys__ssl__Certificate) get_altNames() *hxrt.Array {
	return hxrt.ArrayFromValues(func(hx_sort_src_52 []*string) []any {
		hx_sort_out_54 := make([]any, 0, len(hx_sort_src_52))
		for _, hx_sort_item_53 := range hx_sort_src_52 {
			hx_sort_out_54 = append(hx_sort_out_54, hx_sort_item_53)
		}
		return hx_sort_out_54
	}(hxrt.SslCertAltNames(self.handle)))
}

func (self *sys__ssl__Certificate) get_notBefore() *Date {
	return Date_fromTime(hxrt.SslCertNotBeforeMs(self.handle))
}

func (self *sys__ssl__Certificate) get_notAfter() *Date {
	return Date_fromTime(hxrt.SslCertNotAfterMs(self.handle))
}

func sys__ssl__Certificate_fromString(value *string) *sys__ssl__Certificate {
	return New_sys__ssl__Certificate(hxrt.SslCertFromString(value))
}

func sys__ssl__Certificate_loadDefaults() *sys__ssl__Certificate {
	return New_sys__ssl__Certificate(hxrt.SslCertLoadDefaults())
}

func sys__ssl__Certificate_loadFile(file *string) *sys__ssl__Certificate {
	return New_sys__ssl__Certificate(hxrt.SslCertLoadFile(file))
}

func sys__ssl__Certificate_loadPath(path *string) *sys__ssl__Certificate {
	return New_sys__ssl__Certificate(hxrt.SslCertLoadPath(path))
}
