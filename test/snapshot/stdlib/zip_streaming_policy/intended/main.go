package main

import "snapshot/hxrt"

func flushPolicy() *string {
	fullCompressor := New_haxe__zip__Compress(6)
	full := throws(func() {
		fullCompressor.__hx_this.setFlushMode(haxe__zip__FlushMode_FULL)
	})
	fullCompressor.__hx_this.close()
	blockCompressor := New_haxe__zip__Compress(6)
	block := throws(func() {
		blockCompressor.__hx_this.setFlushMode(haxe__zip__FlushMode_BLOCK)
	})
	blockCompressor.__hx_this.close()
	fullInflater := New_haxe__zip__Uncompress(nil)
	inflateFull := throws(func() {
		fullInflater.__hx_this.setFlushMode(haxe__zip__FlushMode_FULL)
	})
	fullInflater.__hx_this.close()
	blockInflater := New_haxe__zip__Uncompress(nil)
	inflateBlock := throws(func() {
		blockInflater.__hx_this.setFlushMode(haxe__zip__FlushMode_BLOCK)
	})
	blockInflater.__hx_this.close()
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString(full)), hxrt.StringFromLiteral(":")), hxrt.StdString(block)), hxrt.StringFromLiteral(":")), hxrt.StdString(inflateFull)), hxrt.StringFromLiteral(":")), hxrt.StdString(inflateBlock))
}

func lifecycleAndBounds(payload *haxe__io__Bytes) *string {
	closedCompressor := New_haxe__zip__Compress(6)
	closedCompressor.__hx_this.close()
	closedCompressor.__hx_this.close()
	compressClosed := throws(func() {
		closedCompressor.__hx_this.execute(payload, 0, haxe__io__Bytes_alloc(16), 0)
	})
	closedInflater := New_haxe__zip__Uncompress(nil)
	closedInflater.__hx_this.close()
	closedInflater.__hx_this.close()
	inflateClosed := throws(func() {
		closedInflater.__hx_this.execute(payload, 0, haxe__io__Bytes_alloc(16), 0)
	})
	compressor := New_haxe__zip__Compress(6)
	compressBounds := throws(func() {
		compressor.__hx_this.execute(payload, -1, haxe__io__Bytes_alloc(16), 0)
	})
	zeroCompress := compressor.__hx_this.execute(payload, 0, haxe__io__Bytes_alloc(0), 0)
	compressor.__hx_this.close()
	inflater := New_haxe__zip__Uncompress(nil)
	inflateBounds := throws(func() {
		inflater.__hx_this.execute(payload, 0, haxe__io__Bytes_alloc(16), 17)
	})
	zeroInflate := inflater.__hx_this.execute(payload, 0, haxe__io__Bytes_alloc(0), 0)
	inflater.__hx_this.close()
	zeroCapacity := ((((((func(hx_obj_1 map[string]any) int {
		hx_field_2 := hx_obj_1["read"]
		if hx_field_2 == nil {
			var hx_zero_3 int
			return hx_zero_3
		}
		return hx_field_2.(int)
	}(zeroCompress) == 0) && (func(hx_obj_4 map[string]any) int {
		hx_field_5 := hx_obj_4["write"]
		if hx_field_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_field_5.(int)
	}(zeroCompress) == 0)) && !func(hx_obj_7 map[string]any) bool {
		hx_field_8 := hx_obj_7["done"]
		if hx_field_8 == nil {
			var hx_zero_9 bool
			return hx_zero_9
		}
		return hx_field_8.(bool)
	}(zeroCompress)) && (func(hx_obj_10 map[string]any) int {
		hx_field_11 := hx_obj_10["read"]
		if hx_field_11 == nil {
			var hx_zero_12 int
			return hx_zero_12
		}
		return hx_field_11.(int)
	}(zeroInflate) == 0)) && (func(hx_obj_13 map[string]any) int {
		hx_field_14 := hx_obj_13["write"]
		if hx_field_14 == nil {
			var hx_zero_15 int
			return hx_zero_15
		}
		return hx_field_14.(int)
	}(zeroInflate) == 0)) && !func(hx_obj_16 map[string]any) bool {
		hx_field_17 := hx_obj_16["done"]
		if hx_field_17 == nil {
			var hx_zero_18 bool
			return hx_zero_18
		}
		return hx_field_17.(bool)
	}(zeroInflate))
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString(compressClosed)), hxrt.StringFromLiteral(":")), hxrt.StdString(inflateClosed)), hxrt.StringFromLiteral(":")), hxrt.StdString(compressBounds)), hxrt.StringFromLiteral(":")), hxrt.StdString(inflateBounds)), hxrt.StringFromLiteral(":")), hxrt.StdString(zeroCapacity))
}

func main() {
	payload := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("stream-safe-zip"), nil)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("stream="), streamRoundTrip(payload)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("flush="), flushPolicy()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("lifecycle="), lifecycleAndBounds(payload)))
	hxrt.Println(v_2)
}

func streamRoundTrip(payload *haxe__io__Bytes) *string {
	compressor := New_haxe__zip__Compress(6)
	compressor.__hx_this.setFlushMode(haxe__zip__FlushMode_FINISH)
	compressedOutput := New_haxe__io__BytesBuffer()
	sourcePosition := 0
	compressCalls := 0
	compressDone := false
	for !compressDone {
		destination := haxe__io__Bytes_alloc(4)
		result := compressor.__hx_this.execute(payload, sourcePosition, destination, 0)
		sourcePosition = int((hxrt.Int32Wrap(sourcePosition) + hxrt.Int32Wrap(func(hx_obj_19 map[string]any) int {
			hx_field_20 := hx_obj_19["read"]
			if hx_field_20 == nil {
				var hx_zero_21 int
				return hx_zero_21
			}
			return hx_field_20.(int)
		}(result))))
		len := func(hx_obj_22 map[string]any) int {
			hx_field_23 := hx_obj_22["write"]
			if hx_field_23 == nil {
				var hx_zero_24 int
				return hx_zero_24
			}
			return hx_field_23.(int)
		}(result)
		if (len < 0) || (len > destination.length) {
			hxrt.Throw(haxe__io__Error_OutsideBounds)
		}
		compressedOutput.b = hxrt.BytesBufferAddSlice(compressedOutput.b, destination.__hx_this.getData(), 0, len)
		compressDone = func(hx_obj_25 map[string]any) bool {
			hx_field_26 := hx_obj_25["done"]
			if hx_field_26 == nil {
				var hx_zero_27 bool
				return hx_zero_27
			}
			return hx_field_26.(bool)
		}(result)
		compressCalls = int(int32((compressCalls + 1)))
	}
	compressor.__hx_this.close()
	compressed := compressedOutput.__hx_this.getBytes()
	inflater := New_haxe__zip__Uncompress(nil)
	inflater.__hx_this.setFlushMode(haxe__zip__FlushMode_SYNC)
	restoredOutput := New_haxe__io__BytesBuffer()
	compressedPosition := 0
	inflateCalls := 0
	inflateDone := false
	for !inflateDone {
		destination_1 := haxe__io__Bytes_alloc(3)
		result_1 := inflater.__hx_this.execute(compressed, compressedPosition, destination_1, 0)
		compressedPosition = int((hxrt.Int32Wrap(compressedPosition) + hxrt.Int32Wrap(func(hx_obj_28 map[string]any) int {
			hx_field_29 := hx_obj_28["read"]
			if hx_field_29 == nil {
				var hx_zero_30 int
				return hx_zero_30
			}
			return hx_field_29.(int)
		}(result_1))))
		len_1 := func(hx_obj_31 map[string]any) int {
			hx_field_32 := hx_obj_31["write"]
			if hx_field_32 == nil {
				var hx_zero_33 int
				return hx_zero_33
			}
			return hx_field_32.(int)
		}(result_1)
		if (len_1 < 0) || (len_1 > destination_1.length) {
			hxrt.Throw(haxe__io__Error_OutsideBounds)
		}
		restoredOutput.b = hxrt.BytesBufferAddSlice(restoredOutput.b, destination_1.__hx_this.getData(), 0, len_1)
		inflateDone = func(hx_obj_34 map[string]any) bool {
			hx_field_35 := hx_obj_34["done"]
			if hx_field_35 == nil {
				var hx_zero_36 bool
				return hx_zero_36
			}
			return hx_field_35.(bool)
		}(result_1)
		inflateCalls = int(int32((inflateCalls + 1)))
	}
	inflater.__hx_this.close()
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString(hxrt.StringEqualStringPtr(restoredOutput.__hx_this.getBytes().__hx_this.toString(), payload.__hx_this.toString()))), hxrt.StringFromLiteral(":")), hxrt.StdString((compressCalls > 1))), hxrt.StringFromLiteral(":")), hxrt.StdString((inflateCalls > 1)))
}

func throws(block func()) bool {
	hx_try_return_37 := false
	var hx_try_value_38 bool
	hxrt.TryCatch(func() {
		block()
	}, func(hx_caught_39 any) {
		hx_tmp := hx_caught_39
		_ = hx_tmp
		hx_try_value_38 = true
		hx_try_return_37 = true
		return
	})
	if hx_try_return_37 {
		return hx_try_value_38
	}
	return false
}
