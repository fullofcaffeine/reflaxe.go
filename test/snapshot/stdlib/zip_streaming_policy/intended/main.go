package main

import "snapshot/hxrt"

func flushPolicy() *string {
	fullCompressor := New_haxe__zip__Compress(6)
	full := throws(hxrt.NewFunction(func() {
		fullCompressor.__hx_this.setFlushMode(haxe__zip__FlushMode_FULL)
	}))
	fullCompressor.__hx_this.close()
	blockCompressor := New_haxe__zip__Compress(6)
	block := throws(hxrt.NewFunction(func() {
		blockCompressor.__hx_this.setFlushMode(haxe__zip__FlushMode_BLOCK)
	}))
	blockCompressor.__hx_this.close()
	fullInflater := New_haxe__zip__Uncompress(nil)
	inflateFull := throws(hxrt.NewFunction(func() {
		fullInflater.__hx_this.setFlushMode(haxe__zip__FlushMode_FULL)
	}))
	fullInflater.__hx_this.close()
	blockInflater := New_haxe__zip__Uncompress(nil)
	inflateBlock := throws(hxrt.NewFunction(func() {
		blockInflater.__hx_this.setFlushMode(haxe__zip__FlushMode_BLOCK)
	}))
	blockInflater.__hx_this.close()
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString(full)), hxrt.StringFromLiteral(":")), hxrt.StdString(block)), hxrt.StringFromLiteral(":")), hxrt.StdString(inflateFull)), hxrt.StringFromLiteral(":")), hxrt.StdString(inflateBlock))
}

func lifecycleAndBounds(payload *haxe__io__Bytes) *string {
	closedCompressor := New_haxe__zip__Compress(6)
	closedCompressor.__hx_this.close()
	closedCompressor.__hx_this.close()
	compressClosed := throws(hxrt.NewFunction(func() {
		closedCompressor.__hx_this.execute(payload, 0, haxe__io__Bytes_alloc(16), 0)
	}))
	closedInflater := New_haxe__zip__Uncompress(nil)
	closedInflater.__hx_this.close()
	closedInflater.__hx_this.close()
	inflateClosed := throws(hxrt.NewFunction(func() {
		closedInflater.__hx_this.execute(payload, 0, haxe__io__Bytes_alloc(16), 0)
	}))
	compressor := New_haxe__zip__Compress(6)
	compressBounds := throws(hxrt.NewFunction(func() {
		compressor.__hx_this.execute(payload, -1, haxe__io__Bytes_alloc(16), 0)
	}))
	zeroCompress := compressor.__hx_this.execute(payload, 0, haxe__io__Bytes_alloc(0), 0)
	compressor.__hx_this.close()
	inflater := New_haxe__zip__Uncompress(nil)
	inflateBounds := throws(hxrt.NewFunction(func() {
		inflater.__hx_this.execute(payload, 0, haxe__io__Bytes_alloc(16), 17)
	}))
	zeroInflate := inflater.__hx_this.execute(payload, 0, haxe__io__Bytes_alloc(0), 0)
	inflater.__hx_this.close()
	zeroCapacity := ((((((func(hx_value_1 any) int {
		if hx_value_1 == nil {
			var hx_zero_2 int
			return hx_zero_2
		}
		return hx_value_1.(int)
	}(hxrt.AnonymousGet(zeroCompress, "read")) == 0) && (func(hx_value_3 any) int {
		if hx_value_3 == nil {
			var hx_zero_4 int
			return hx_zero_4
		}
		return hx_value_3.(int)
	}(hxrt.AnonymousGet(zeroCompress, "write")) == 0)) && !func(hx_value_5 any) bool {
		if hx_value_5 == nil {
			var hx_zero_6 bool
			return hx_zero_6
		}
		return hx_value_5.(bool)
	}(hxrt.AnonymousGet(zeroCompress, "done"))) && (func(hx_value_7 any) int {
		if hx_value_7 == nil {
			var hx_zero_8 int
			return hx_zero_8
		}
		return hx_value_7.(int)
	}(hxrt.AnonymousGet(zeroInflate, "read")) == 0)) && (func(hx_value_9 any) int {
		if hx_value_9 == nil {
			var hx_zero_10 int
			return hx_zero_10
		}
		return hx_value_9.(int)
	}(hxrt.AnonymousGet(zeroInflate, "write")) == 0)) && !func(hx_value_11 any) bool {
		if hx_value_11 == nil {
			var hx_zero_12 bool
			return hx_zero_12
		}
		return hx_value_11.(bool)
	}(hxrt.AnonymousGet(zeroInflate, "done")))
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
		sourcePosition = int((hxrt.Int32Wrap(sourcePosition) + hxrt.Int32Wrap(func(hx_value_13 any) int {
			if hx_value_13 == nil {
				var hx_zero_14 int
				return hx_zero_14
			}
			return hx_value_13.(int)
		}(hxrt.AnonymousGet(result, "read")))))
		len := func(hx_value_15 any) int {
			if hx_value_15 == nil {
				var hx_zero_16 int
				return hx_zero_16
			}
			return hx_value_15.(int)
		}(hxrt.AnonymousGet(result, "write"))
		if (len < 0) || (len > destination.length) {
			hxrt.Throw(haxe__Exception_thrown(haxe__io__Error_OutsideBounds))
		}
		compressedOutput.b = hxrt.BytesBufferAddSlice(compressedOutput.b, destination.__hx_this.getData(), 0, len)
		compressDone = func(hx_value_17 any) bool {
			if hx_value_17 == nil {
				var hx_zero_18 bool
				return hx_zero_18
			}
			return hx_value_17.(bool)
		}(hxrt.AnonymousGet(result, "done"))
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
		compressedPosition = int((hxrt.Int32Wrap(compressedPosition) + hxrt.Int32Wrap(func(hx_value_19 any) int {
			if hx_value_19 == nil {
				var hx_zero_20 int
				return hx_zero_20
			}
			return hx_value_19.(int)
		}(hxrt.AnonymousGet(result_1, "read")))))
		len_1 := func(hx_value_21 any) int {
			if hx_value_21 == nil {
				var hx_zero_22 int
				return hx_zero_22
			}
			return hx_value_21.(int)
		}(hxrt.AnonymousGet(result_1, "write"))
		if (len_1 < 0) || (len_1 > destination_1.length) {
			hxrt.Throw(haxe__Exception_thrown(haxe__io__Error_OutsideBounds))
		}
		restoredOutput.b = hxrt.BytesBufferAddSlice(restoredOutput.b, destination_1.__hx_this.getData(), 0, len_1)
		inflateDone = func(hx_value_23 any) bool {
			if hx_value_23 == nil {
				var hx_zero_24 bool
				return hx_zero_24
			}
			return hx_value_23.(bool)
		}(hxrt.AnonymousGet(result_1, "done"))
		inflateCalls = int(int32((inflateCalls + 1)))
	}
	inflater.__hx_this.close()
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), hxrt.StdString(hxrt.StringEqualStringPtr(restoredOutput.__hx_this.getBytes().__hx_this.toString(), payload.__hx_this.toString()))), hxrt.StringFromLiteral(":")), hxrt.StdString((compressCalls > 1))), hxrt.StringFromLiteral(":")), hxrt.StdString((inflateCalls > 1)))
}

func throws(block *hxrt.Function[func()]) bool {
	hx_try_return_25 := false
	var hx_try_value_26 bool
	hxrt.TryCatch(func() {
		func() {
			var hx_callable_30 *hxrt.Function[func()] = block
			hxrt.FunctionCall(hx_callable_30)()
		}()
	}, func(hx_caught_27 any) {
		hx_exception_28 := haxe__Exception_caught(hx_caught_27)
		hx_payload_29 := haxe__Exception_catchValue(hx_caught_27)
		_ = hx_exception_28
		var hx_tmp any = hx_payload_29
		_ = hx_tmp
		hx_try_value_26 = true
		hx_try_return_25 = true
		return
	})
	if hx_try_return_25 {
		return hx_try_value_26
	}
	return false
}

func hx_recover_class_0(value any) *haxe__ValueException {
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

func hx_recover_class_1(value any) *haxe__Exception {
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
