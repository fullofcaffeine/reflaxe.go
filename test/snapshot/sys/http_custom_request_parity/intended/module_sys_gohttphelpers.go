package main

import "snapshot/hxrt"

func sys__GoHttpHelpers_captureApi(api any, payload *haxe__io__Bytes) {
	if hxrt.AnyEqualsNull(api) || (payload == nil) {
		return
	}
	_ = func() int {
		switch out := api.(type) {
		case *haxe__io__BytesBuffer:
			out.add(payload)
		case interface{ add(*haxe__io__Bytes) }:
			out.add(payload)
		case interface {
			writeBytes(*haxe__io__Bytes, int, int) int
		}:
			out.writeBytes(payload, 0, payload.length)
		case interface {
			writeFullBytes(*haxe__io__Bytes, int, int)
		}:
			out.writeFullBytes(payload, 0, payload.length)
		case interface{ writeString(*string) }:
			out.writeString(payload.toString())
		}
		return 0
	}()
}

func sys__GoHttpHelpers_getResponseHeaderValues(self *sys__Http, key *string) *hxrt.Array {
	if self == nil {
		return nil
	}
	normalized := func() *string {
		raw := *hxrt.StdString(key)
		out := make([]byte, len(raw))
		for i := 0; i < len(raw); i++ {
			c := raw[i]
			if c >= 'A' && c <= 'Z' {
				c += 32
			}
			out[i] = c
		}
		return hxrt.StringFromLiteral(string(out))
	}()
	_ = normalized
	nativeValues := func() []*string {
		if self.responseHeadersSameKey != nil {
			if values, ok := self.responseHeadersSameKey[*hxrt.StdString(key)]; ok {
				return values
			}
			if values, ok := self.responseHeadersSameKey[*hxrt.StdString(normalized)]; ok {
				return values
			}
		}
		if self.responseHeaders == nil {
			return nil
		}
		single := self.responseHeaders.get(key)
		if single == nil && *hxrt.StdString(key) != *hxrt.StdString(normalized) {
			single = self.responseHeaders.get(normalized)
		}
		if single == nil {
			return nil
		}
		return []*string{hxrt.StdString(single)}
	}()
	var hx_if_57 *hxrt.Array
	if nativeValues == nil {
		hx_if_57 = nil
	} else {
		hx_if_57 = hxrt.ArrayFromValues(func(hx_sort_src_54 []*string) []any {
			hx_sort_out_56 := make([]any, 0, len(hx_sort_src_54))
			for _, hx_sort_item_55 := range hx_sort_src_54 {
				hx_sort_out_56 = append(hx_sort_out_56, hx_sort_item_55)
			}
			return hx_sort_out_56
		}(nativeValues))
	}
	return hx_if_57
}
