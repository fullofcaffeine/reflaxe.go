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
	handle     any
	commonName *string
	altNames   *hxrt.Array
	notBefore  *Date
	notAfter   *Date
}

func New_sys__ssl__Certificate(handle any) *sys__ssl__Certificate {
	self := &sys__ssl__Certificate{}
	self.__hx_this = self
	self.handle = handle
	return self
}

func (self *sys__ssl__Certificate) subject(field *string) *string {
	return hxrt.SslCertSubject(self.handle, field)
}

func (self *sys__ssl__Certificate) issuer(field *string) *string {
	return hxrt.SslCertIssuer(self.handle, field)
}

func (self *sys__ssl__Certificate) next() *sys__ssl__Certificate {
	var nextHandle any = hxrt.SslCertNext(self.handle)
	var hx_if_12 *sys__ssl__Certificate
	if hxrt.AnyEqualsNull(nextHandle) {
		hx_if_12 = nil
	} else {
		hx_if_12 = New_sys__ssl__Certificate(nextHandle)
	}
	return hx_if_12
}

func (self *sys__ssl__Certificate) add(pem *string) {
	_ = func() int { hxrt.SslCertAddPEM(self.handle, pem); return 0 }()
}

func (self *sys__ssl__Certificate) addDER(der *haxe__io__Bytes) {
	_ = func() int { hxrt.SslCertAddDER(self.handle, hxrt_haxeBytesToRaw(der)); return 0 }()
}

func (self *sys__ssl__Certificate) get_commonName() *string {
	return hxrt.SslCertCommonName(self.handle)
}

func (self *sys__ssl__Certificate) get_altNames() *hxrt.Array {
	return hxrt.ArrayFromValues(func(hx_sort_src_13 []*string) []any {
		hx_sort_out_15 := make([]any, 0, len(hx_sort_src_13))
		for _, hx_sort_item_14 := range hx_sort_src_13 {
			hx_sort_out_15 = append(hx_sort_out_15, hx_sort_item_14)
		}
		return hx_sort_out_15
	}(hxrt.SslCertAltNames(self.handle)))
}

func (self *sys__ssl__Certificate) get_notBefore() *Date {
	return Date_fromTime(hxrt.SslCertNotBeforeMs(self.handle))
}

func (self *sys__ssl__Certificate) get_notAfter() *Date {
	return Date_fromTime(hxrt.SslCertNotAfterMs(self.handle))
}

func sys__ssl__Certificate_fromString(str *string) *sys__ssl__Certificate {
	return New_sys__ssl__Certificate(hxrt.SslCertFromString(str))
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
