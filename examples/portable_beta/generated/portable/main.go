package main

import "examples_portable_beta/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type _Main__DeliveryState struct {
	tag    int
	params []any
}

var _Main__DeliveryState_Queued *_Main__DeliveryState = &_Main__DeliveryState{tag: 0}

var _Main__DeliveryState_Delivered *_Main__DeliveryState = &_Main__DeliveryState{tag: 1}

var SCRATCH_FILE *string = hxrt.StringFromLiteral(".portable_beta_contract.txt")

func checkExceptionContract() {
	observed := hxrt.StringFromLiteral("")
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("portable-beta-error")))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_payload_3
		var error *haxe__Exception = hx_recover_class_0(hx_exception_2)
		observed = error.__hx_this.get_message()
	})
	require(hxrt.StringEqualStringPtr(observed, hxrt.StringFromLiteral("portable-beta-error")), hxrt.StringFromLiteral("typed exception message contract"))
}

func checkFileLifecycle() {
	if sys__FileSystem_exists(hxrt.StringFromLiteral(".portable_beta_contract.txt")) {
		sys__FileSystem_deleteFile(hxrt.StringFromLiteral(".portable_beta_contract.txt"))
	}
	hxrt.TryCatch(func() {
		sys__io__File_saveContent(hxrt.StringFromLiteral(".portable_beta_contract.txt"), hxrt.StringFromLiteral("portable-beta-ok"))
		require(sys__FileSystem_exists(hxrt.StringFromLiteral(".portable_beta_contract.txt")), hxrt.StringFromLiteral("file existence contract"))
		require(hxrt.StringEqualStringPtr(sys__io__File_getContent(hxrt.StringFromLiteral(".portable_beta_contract.txt")), hxrt.StringFromLiteral("portable-beta-ok")), hxrt.StringFromLiteral("file content contract"))
		sys__FileSystem_deleteFile(hxrt.StringFromLiteral(".portable_beta_contract.txt"))
	}, func(hx_caught_4 any) {
		hx_exception_5 := haxe__Exception_caught(hx_caught_4)
		hx_payload_6 := haxe__Exception_catchValue(hx_caught_4)
		_ = hx_payload_6
		var error *haxe__Exception = hx_recover_class_0(hx_exception_5)
		if sys__FileSystem_exists(hxrt.StringFromLiteral(".portable_beta_contract.txt")) {
			sys__FileSystem_deleteFile(hxrt.StringFromLiteral(".portable_beta_contract.txt"))
		}
		hxrt.Throw(haxe__Exception_thrown(error))
	})
	require(!sys__FileSystem_exists(hxrt.StringFromLiteral(".portable_beta_contract.txt")), hxrt.StringFromLiteral("file cleanup contract"))
}

func checkLanguageAndCollections() {
	deliveries := hxrt.NewArray()
	deliveries.Push(New__Main__Delivery(hxrt.StringFromLiteral("alpha"), _Main__DeliveryState_Delivered))
	deliveries.Push(New__Main__Delivery(hxrt.StringFromLiteral("beta"), _Main__DeliveryState_Queued))
	deliveries.Push(New__Main__Delivery(hxrt.StringFromLiteral("gamma"), _Main__DeliveryState_Delivered))
	delivered := 0
	_g := 0
	_g1 := deliveries.Len()
	for _g < _g1 {
		hx_post_10 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_10
		_g_1 := hx_recover_class_1(deliveries.Get(index)).state
		switch _g_1.tag {
		case 0:
		case 1:
			delivered = int(int32((delivered + 1)))
		}
	}
	require((delivered == 2), hxrt.StringFromLiteral("enum and numeric-for contract"))
	removed := func() *_Main__Delivery {
		return hx_recover_class_1(deliveries.Pop())
	}()
	require((removed != nil), hxrt.StringFromLiteral("array pop contract"))
	if removed != nil {
		require(hxrt.StringEqualStringPtr(removed.name, hxrt.StringFromLiteral("gamma")), hxrt.StringFromLiteral("array order contract"))
	}
	require((deliveries.Len() == 2), hxrt.StringFromLiteral("array length contract"))
	decorate := hxrt.NewFunction(func(value *string) *string {
		return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), value), hxrt.StringFromLiteral("]"))
	})
	require(hxrt.StringEqualStringPtr(func() *string {
		var hx_callable_12 *hxrt.Function[func(*string) *string] = decorate
		var hx_argument_13 *string = hxrt.StringFromLiteral("portable")
		return hxrt.FunctionCall(hx_callable_12)(hx_argument_13)
	}(), hxrt.StringFromLiteral("[portable]")), hxrt.StringFromLiteral("closure contract"))
}

func checkTextAndData() {
	cleaned := StringTools_trim(hxrt.StringFromLiteral("  alpha-go  "))
	require((hxrt.StringLengthStringPtr(cleaned) == 8), hxrt.StringFromLiteral("string length contract"))
	require(hxrt.StringEqualStringPtr(hxrt.StringCharAtStringPtr(cleaned, 5), hxrt.StringFromLiteral("-")), hxrt.StringFromLiteral("string charAt contract"))
	require(hxrt.StringEqualStringPtr(hxrt.StringSubstringStringPtr(cleaned, 0, 5), hxrt.StringFromLiteral("alpha")), hxrt.StringFromLiteral("string substring contract"))
	require(StringTools_startsWith(cleaned, hxrt.StringFromLiteral("alpha")), hxrt.StringFromLiteral("startsWith contract"))
	require(StringTools_endsWith(cleaned, hxrt.StringFromLiteral("go")), hxrt.StringFromLiteral("endsWith contract"))
	require(StringTools_contains(cleaned, hxrt.StringFromLiteral("ha-g")), hxrt.StringFromLiteral("contains contract"))
	require(hxrt.StringEqualStringPtr(StringTools_replace(cleaned, hxrt.StringFromLiteral("-"), hxrt.StringFromLiteral(" ")), hxrt.StringFromLiteral("alpha go")), hxrt.StringFromLiteral("replace contract"))
	bytes := haxe__io__Bytes_ofString(cleaned, nil)
	require(hxrt.StringEqualStringPtr(bytes.__hx_this.toString(), cleaned), hxrt.StringFromLiteral("bytes round-trip contract"))
	names := New_haxe__ds__StringMap()
	names.__hx_this.set(hxrt.StringFromLiteral("primary"), hxrt.StringFromLiteral("alpha"))
	require(func(hx_value_14 any) bool {
		if hx_value_14 == nil {
			var hx_zero_15 bool
			return hx_zero_15
		}
		return hx_value_14.(bool)
	}(names.__hx_this.exists(hxrt.StringFromLiteral("primary"))), hxrt.StringFromLiteral("StringMap exists contract"))
	require(hxrt.StringEqualStringPtr(func(hx_value_16 any) *string {
		if hx_value_16 == nil {
			var hx_zero_17 *string
			return hx_zero_17
		}
		return hx_value_16.(*string)
	}(names.__hx_this.get(hxrt.StringFromLiteral("primary"))), hxrt.StringFromLiteral("alpha")), hxrt.StringFromLiteral("StringMap get contract"))
	require(func(hx_value_18 any) bool {
		if hx_value_18 == nil {
			var hx_zero_19 bool
			return hx_zero_19
		}
		return hx_value_18.(bool)
	}(names.__hx_this.remove(hxrt.StringFromLiteral("primary"))), hxrt.StringFromLiteral("StringMap remove contract"))
	require(!func(hx_value_20 any) bool {
		if hx_value_20 == nil {
			var hx_zero_21 bool
			return hx_zero_21
		}
		return hx_value_20.(bool)
	}(names.__hx_this.exists(hxrt.StringFromLiteral("primary"))), hxrt.StringFromLiteral("StringMap removal contract"))
	var parsed any = hxrt.JsonParse(hxrt.StringFromLiteral("{\"name\":\"alpha\",\"ready\":true}"))
	require(Reflect_hasField(parsed, hxrt.StringFromLiteral("name")), hxrt.StringFromLiteral("JSON field contract"))
	parsedName := func(hx_value_22 any) *string {
		if hx_value_22 == nil {
			var hx_zero_23 *string
			return hx_zero_23
		}
		return hx_value_22.(*string)
	}(Reflect_field(parsed, hxrt.StringFromLiteral("name")))
	require(hxrt.StringEqualStringPtr(parsedName, hxrt.StringFromLiteral("alpha")), hxrt.StringFromLiteral("JSON value contract"))
	var space *string = nil
	encoded := hxrt.StdString(hxrt.JsonStringify(any(func() hxrt.Anonymous {
		hx_obj_24 := map[string]any{}
		hx_obj_24["name"] = hxrt.StringFromLiteral("alpha")
		return hxrt.AnonymousFromMap(hx_obj_24)
	}()), space))
	require(StringTools_contains(encoded, hxrt.StringFromLiteral("\"name\":\"alpha\"")), hxrt.StringFromLiteral("JSON stringify contract"))
	path := New_haxe__io__Path(haxe__io__Path_join(hxrt.NewArray(hxrt.StringFromLiteral("reports"), hxrt.StringFromLiteral("result.json"))))
	require(hxrt.StringEqualStringPtr(path.dir, hxrt.StringFromLiteral("reports")), hxrt.StringFromLiteral("Path directory contract"))
	require(hxrt.StringEqualStringPtr(path.file, hxrt.StringFromLiteral("result")), hxrt.StringFromLiteral("Path file contract"))
	require(hxrt.StringEqualStringPtr(path.ext, hxrt.StringFromLiteral("json")), hxrt.StringFromLiteral("Path extension contract"))
}

func main() {
	checkLanguageAndCollections()
	checkTextAndData()
	checkFileLifecycle()
	checkExceptionContract()
}

func require(condition bool, message *string) {
	if !condition {
		hxrt.Throw(haxe__Exception_thrown(message))
	}
}

type I__Main__Delivery interface {
}

type _Main__Delivery struct {
	__hx_this I__Main__Delivery
	name      *string
	state     *_Main__DeliveryState
}

func (self *_Main__Delivery) HxFunctionReceiver() any {
	return self.__hx_this
}

func New__Main__Delivery(name *string, state *_Main__DeliveryState) *_Main__Delivery {
	return hx_init__Main__Delivery(hx_alloc__Main__Delivery(nil), name, state)
}

func hx_alloc__Main__Delivery(owner any) *_Main__Delivery {
	self := &_Main__Delivery{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I__Main__Delivery)
	}
	return self
}

func hx_init__Main__Delivery(self *_Main__Delivery, name *string, state *_Main__DeliveryState) *_Main__Delivery {
	self.name = name
	self.state = state
	return self
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__NotImplementedException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__PosException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__BytesBuffer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Eof:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Input:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Output:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__io__FileInput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__io__FileOutput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_method_field__Date(value, key)
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_method_field__haxe__ds__StringMap(value, key)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_method_field__haxe__exceptions__NotImplementedException(value, key)
	case *haxe__exceptions__PosException:
		return hxrt__generated_method_field__haxe__exceptions__PosException(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_method_field__haxe__io__Bytes(value, key)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_method_field__haxe__io__BytesBuffer(value, key)
	case *haxe__io__Eof:
		return hxrt__generated_method_field__haxe__io__Eof(value, key)
	case *haxe__io__Input:
		return hxrt__generated_method_field__haxe__io__Input(value, key)
	case *haxe__io__Output:
		return hxrt__generated_method_field__haxe__io__Output(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_method_field__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value, key)
	case *sys__io__FileInput:
		return hxrt__generated_method_field__sys__io__FileInput(value, key)
	case *sys__io__FileOutput:
		return hxrt__generated_method_field__sys__io__FileOutput(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Date(value *Date, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "getDate":
		return func() *hxrt.Function[func() int] {
			hx_receiver_267 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_267.__hx_this.getDate()
			}, func() any {
				return hx_receiver_267.__hx_this
			}, "getDate")
		}()
	case "getDay":
		return func() *hxrt.Function[func() int] {
			hx_receiver_268 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_268.__hx_this.getDay()
			}, func() any {
				return hx_receiver_268.__hx_this
			}, "getDay")
		}()
	case "getFullYear":
		return func() *hxrt.Function[func() int] {
			hx_receiver_269 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_269.__hx_this.getFullYear()
			}, func() any {
				return hx_receiver_269.__hx_this
			}, "getFullYear")
		}()
	case "getHours":
		return func() *hxrt.Function[func() int] {
			hx_receiver_270 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_270.__hx_this.getHours()
			}, func() any {
				return hx_receiver_270.__hx_this
			}, "getHours")
		}()
	case "getMinutes":
		return func() *hxrt.Function[func() int] {
			hx_receiver_271 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_271.__hx_this.getMinutes()
			}, func() any {
				return hx_receiver_271.__hx_this
			}, "getMinutes")
		}()
	case "getMonth":
		return func() *hxrt.Function[func() int] {
			hx_receiver_272 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_272.__hx_this.getMonth()
			}, func() any {
				return hx_receiver_272.__hx_this
			}, "getMonth")
		}()
	case "getSeconds":
		return func() *hxrt.Function[func() int] {
			hx_receiver_273 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_273.__hx_this.getSeconds()
			}, func() any {
				return hx_receiver_273.__hx_this
			}, "getSeconds")
		}()
	case "getTime":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_274 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_274.__hx_this.getTime()
			}, func() any {
				return hx_receiver_274.__hx_this
			}, "getTime")
		}()
	case "getTimezoneOffset":
		return func() *hxrt.Function[func() int] {
			hx_receiver_275 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_275.__hx_this.getTimezoneOffset()
			}, func() any {
				return hx_receiver_275.__hx_this
			}, "getTimezoneOffset")
		}()
	case "getUTCDate":
		return func() *hxrt.Function[func() int] {
			hx_receiver_276 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_276.__hx_this.getUTCDate()
			}, func() any {
				return hx_receiver_276.__hx_this
			}, "getUTCDate")
		}()
	case "getUTCDay":
		return func() *hxrt.Function[func() int] {
			hx_receiver_277 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_277.__hx_this.getUTCDay()
			}, func() any {
				return hx_receiver_277.__hx_this
			}, "getUTCDay")
		}()
	case "getUTCFullYear":
		return func() *hxrt.Function[func() int] {
			hx_receiver_278 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_278.__hx_this.getUTCFullYear()
			}, func() any {
				return hx_receiver_278.__hx_this
			}, "getUTCFullYear")
		}()
	case "getUTCHours":
		return func() *hxrt.Function[func() int] {
			hx_receiver_279 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_279.__hx_this.getUTCHours()
			}, func() any {
				return hx_receiver_279.__hx_this
			}, "getUTCHours")
		}()
	case "getUTCMinutes":
		return func() *hxrt.Function[func() int] {
			hx_receiver_280 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_280.__hx_this.getUTCMinutes()
			}, func() any {
				return hx_receiver_280.__hx_this
			}, "getUTCMinutes")
		}()
	case "getUTCMonth":
		return func() *hxrt.Function[func() int] {
			hx_receiver_281 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_281.__hx_this.getUTCMonth()
			}, func() any {
				return hx_receiver_281.__hx_this
			}, "getUTCMonth")
		}()
	case "getUTCSeconds":
		return func() *hxrt.Function[func() int] {
			hx_receiver_282 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_282.__hx_this.getUTCSeconds()
			}, func() any {
				return hx_receiver_282.__hx_this
			}, "getUTCSeconds")
		}()
	case "localParts":
		return func() *hxrt.Function[func() *hxrt.DateParts] {
			hx_receiver_283 := value
			return hxrt.BoundFunction(func() *hxrt.DateParts {
				return hx_receiver_283.__hx_this.localParts()
			}, func() any {
				return hx_receiver_283.__hx_this
			}, "localParts")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_284 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_284.__hx_this.toString()
			}, func() any {
				return hx_receiver_284.__hx_this
			}, "toString")
		}()
	case "utcParts":
		return func() *hxrt.Function[func() *hxrt.DateParts] {
			hx_receiver_285 := value
			return hxrt.BoundFunction(func() *hxrt.DateParts {
				return hx_receiver_285.__hx_this.utcParts()
			}, func() any {
				return hx_receiver_285.__hx_this
			}, "utcParts")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_286 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_286.__hx_this.details()
			}, func() any {
				return hx_receiver_286.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_287 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_287.__hx_this.get_message()
			}, func() any {
				return hx_receiver_287.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_288 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_288.__hx_this.get_native()
			}, func() any {
				return hx_receiver_288.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_289 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_289.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_289.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_290 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_290.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_290.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_291 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_291.__hx_this.toString()
			}, func() any {
				return hx_receiver_291.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_292 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_292.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_292.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_293 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_293.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_293.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_294 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_294.__hx_this.clear()
			}, func() any {
				return hx_receiver_294.__hx_this
			}, "clear")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__StringMap] {
			hx_receiver_295 := value
			return hxrt.BoundFunction(func() *haxe__ds__StringMap {
				return hx_receiver_295.__hx_this.copy()
			}, func() any {
				return hx_receiver_295.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_296 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_296.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_296.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_297 := value
			return hxrt.BoundFunction(func(key *string) bool {
				return hx_receiver_297.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_297.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_298 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_298.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_298.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(*string) any] {
			hx_receiver_299 := value
			return hxrt.BoundFunction(func(key *string) any {
				return hx_receiver_299.__hx_this.get(key)
			}, func() any {
				return hx_receiver_299.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_300 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_300.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_300.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_301 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_301.__hx_this.iterator()
			}, func() any {
				return hx_receiver_301.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_302 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_302.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_302.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_303 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_303.__hx_this.keys()
			}, func() any {
				return hx_receiver_303.__hx_this
			}, "keys")
		}()
	case "remove":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_304 := value
			return hxrt.BoundFunction(func(key *string) bool {
				return hx_receiver_304.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_304.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_305 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_305.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_305.__hx_this
			}, "removeIMap")
		}()
	case "set":
		return func() *hxrt.Function[func(*string, any)] {
			hx_receiver_306 := value
			return hxrt.BoundFunction(func(key *string, value any) {
				hx_receiver_306.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_306.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_307 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_307.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_307.__hx_this
			}, "setIMap")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_308 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_308.__hx_this.toString()
			}, func() any {
				return hx_receiver_308.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__exceptions__NotImplementedException(value *haxe__exceptions__NotImplementedException, key string) any {
	if value == nil {
		return nil
	}
	if value.haxe__exceptions__PosException == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__exceptions__PosException(value.haxe__exceptions__PosException, key)
}

func hxrt__generated_method_field__haxe__exceptions__PosException(value *haxe__exceptions__PosException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_309 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_309.__hx_this.toString()
			}, func() any {
				return hx_receiver_309.__hx_this
			}, "toString")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_nativeView":
		return func() *hxrt.Function[func() *hxrt.ByteView] {
			hx_receiver_310 := value
			return hxrt.BoundFunction(func() *hxrt.ByteView {
				return hx_receiver_310.__hx_this.__hx_nativeView()
			}, func() any {
				return hx_receiver_310.__hx_this
			}, "__hx_nativeView")
		}()
	case "blit":
		return func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
			hx_receiver_311 := value
			return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
				hx_receiver_311.__hx_this.blit(pos, src, srcpos, len)
			}, func() any {
				return hx_receiver_311.__hx_this
			}, "blit")
		}()
	case "compare":
		return func() *hxrt.Function[func(*haxe__io__Bytes) int] {
			hx_receiver_312 := value
			return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
				return hx_receiver_312.__hx_this.compare(other)
			}, func() any {
				return hx_receiver_312.__hx_this
			}, "compare")
		}()
	case "fill":
		return func() *hxrt.Function[func(int, int, int)] {
			hx_receiver_313 := value
			return hxrt.BoundFunction(func(pos int, len int, value int) {
				hx_receiver_313.__hx_this.fill(pos, len, value)
			}, func() any {
				return hx_receiver_313.__hx_this
			}, "fill")
		}()
	case "get":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_314 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_314.__hx_this.get(pos)
			}, func() any {
				return hx_receiver_314.__hx_this
			}, "get")
		}()
	case "getData":
		return func() *hxrt.Function[func() []int] {
			hx_receiver_315 := value
			return hxrt.BoundFunction(func() []int {
				return hx_receiver_315.__hx_this.getData()
			}, func() any {
				return hx_receiver_315.__hx_this
			}, "getData")
		}()
	case "getDouble":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_316 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_316.__hx_this.getDouble(pos)
			}, func() any {
				return hx_receiver_316.__hx_this
			}, "getDouble")
		}()
	case "getFloat":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_317 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_317.__hx_this.getFloat(pos)
			}, func() any {
				return hx_receiver_317.__hx_this
			}, "getFloat")
		}()
	case "getInt32":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_318 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_318.__hx_this.getInt32(pos)
			}, func() any {
				return hx_receiver_318.__hx_this
			}, "getInt32")
		}()
	case "getInt64":
		return func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
			hx_receiver_319 := value
			return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
				return hx_receiver_319.__hx_this.getInt64(pos)
			}, func() any {
				return hx_receiver_319.__hx_this
			}, "getInt64")
		}()
	case "getString":
		return func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
			hx_receiver_320 := value
			return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_320.__hx_this.getString(pos, len, encoding)
			}, func() any {
				return hx_receiver_320.__hx_this
			}, "getString")
		}()
	case "getUInt16":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_321 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_321.__hx_this.getUInt16(pos)
			}, func() any {
				return hx_receiver_321.__hx_this
			}, "getUInt16")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, int) *string] {
			hx_receiver_322 := value
			return hxrt.BoundFunction(func(pos int, len int) *string {
				return hx_receiver_322.__hx_this.readString(pos, len)
			}, func() any {
				return hx_receiver_322.__hx_this
			}, "readString")
		}()
	case "set":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_323 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_323.__hx_this.set(pos, value)
			}, func() any {
				return hx_receiver_323.__hx_this
			}, "set")
		}()
	case "setDouble":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_324 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_324.__hx_this.setDouble(pos, value)
			}, func() any {
				return hx_receiver_324.__hx_this
			}, "setDouble")
		}()
	case "setFloat":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_325 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_325.__hx_this.setFloat(pos, value)
			}, func() any {
				return hx_receiver_325.__hx_this
			}, "setFloat")
		}()
	case "setInt32":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_326 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_326.__hx_this.setInt32(pos, value)
			}, func() any {
				return hx_receiver_326.__hx_this
			}, "setInt32")
		}()
	case "setInt64":
		return func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
			hx_receiver_327 := value
			return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
				hx_receiver_327.__hx_this.setInt64(pos, value)
			}, func() any {
				return hx_receiver_327.__hx_this
			}, "setInt64")
		}()
	case "setUInt16":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_328 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_328.__hx_this.setUInt16(pos, value)
			}, func() any {
				return hx_receiver_328.__hx_this
			}, "setUInt16")
		}()
	case "sub":
		return func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
			hx_receiver_329 := value
			return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
				return hx_receiver_329.__hx_this.sub(pos, len)
			}, func() any {
				return hx_receiver_329.__hx_this
			}, "sub")
		}()
	case "toHex":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_330 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_330.__hx_this.toHex()
			}, func() any {
				return hx_receiver_330.__hx_this
			}, "toHex")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_331 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_331.__hx_this.toString()
			}, func() any {
				return hx_receiver_331.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__BytesBuffer(value *haxe__io__BytesBuffer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return func() *hxrt.Function[func(*haxe__io__Bytes)] {
			hx_receiver_332 := value
			return hxrt.BoundFunction(func(source *haxe__io__Bytes) {
				hx_receiver_332.__hx_this.add(source)
			}, func() any {
				return hx_receiver_332.__hx_this
			}, "add")
		}()
	case "addByte":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_333 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_333.__hx_this.addByte(value)
			}, func() any {
				return hx_receiver_333.__hx_this
			}, "addByte")
		}()
	case "addBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
			hx_receiver_334 := value
			return hxrt.BoundFunction(func(source *haxe__io__Bytes, pos int, len int) {
				hx_receiver_334.__hx_this.addBytes(source, pos, len)
			}, func() any {
				return hx_receiver_334.__hx_this
			}, "addBytes")
		}()
	case "addDouble":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_335 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_335.__hx_this.addDouble(value)
			}, func() any {
				return hx_receiver_335.__hx_this
			}, "addDouble")
		}()
	case "addFloat":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_336 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_336.__hx_this.addFloat(value)
			}, func() any {
				return hx_receiver_336.__hx_this
			}, "addFloat")
		}()
	case "addInt32":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_337 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_337.__hx_this.addInt32(value)
			}, func() any {
				return hx_receiver_337.__hx_this
			}, "addInt32")
		}()
	case "addInt64":
		return func() *hxrt.Function[func(*haxe___Int64_____Int64)] {
			hx_receiver_338 := value
			return hxrt.BoundFunction(func(value *haxe___Int64_____Int64) {
				hx_receiver_338.__hx_this.addInt64(value)
			}, func() any {
				return hx_receiver_338.__hx_this
			}, "addInt64")
		}()
	case "addString":
		return func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
			hx_receiver_339 := value
			return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
				hx_receiver_339.__hx_this.addString(value, encoding)
			}, func() any {
				return hx_receiver_339.__hx_this
			}, "addString")
		}()
	case "getBytes":
		return func() *hxrt.Function[func() *haxe__io__Bytes] {
			hx_receiver_340 := value
			return hxrt.BoundFunction(func() *haxe__io__Bytes {
				return hx_receiver_340.__hx_this.getBytes()
			}, func() any {
				return hx_receiver_340.__hx_this
			}, "getBytes")
		}()
	case "get_length":
		return func() *hxrt.Function[func() int] {
			hx_receiver_341 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_341.__hx_this.get_length()
			}, func() any {
				return hx_receiver_341.__hx_this
			}, "get_length")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Eof(value *haxe__io__Eof, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_342 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_342.__hx_this.toString()
			}, func() any {
				return hx_receiver_342.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Input(value *haxe__io__Input, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return func() *hxrt.Function[func()] {
			hx_receiver_343 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_343.__hx_this.close()
			}, func() any {
				return hx_receiver_343.__hx_this
			}, "close")
		}()
	case "read":
		return func() *hxrt.Function[func(int) *haxe__io__Bytes] {
			hx_receiver_344 := value
			return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
				return hx_receiver_344.__hx_this.read(nbytes)
			}, func() any {
				return hx_receiver_344.__hx_this
			}, "read")
		}()
	case "readAll":
		return func() *hxrt.Function[func(any) *haxe__io__Bytes] {
			hx_receiver_345 := value
			return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
				return hx_receiver_345.__hx_this.readAll(bufsize)
			}, func() any {
				return hx_receiver_345.__hx_this
			}, "readAll")
		}()
	case "readByte":
		return func() *hxrt.Function[func() int] {
			hx_receiver_346 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_346.__hx_this.readByte()
			}, func() any {
				return hx_receiver_346.__hx_this
			}, "readByte")
		}()
	case "readBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_347 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
				return hx_receiver_347.__hx_this.readBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_347.__hx_this
			}, "readBytes")
		}()
	case "readDouble":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_348 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_348.__hx_this.readDouble()
			}, func() any {
				return hx_receiver_348.__hx_this
			}, "readDouble")
		}()
	case "readFloat":
		return func() *hxrt.Function[func() float64] {
			hx_receiver_349 := value
			return hxrt.BoundFunction(func() float64 {
				return hx_receiver_349.__hx_this.readFloat()
			}, func() any {
				return hx_receiver_349.__hx_this
			}, "readFloat")
		}()
	case "readFullBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
			hx_receiver_350 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
				hx_receiver_350.__hx_this.readFullBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_350.__hx_this
			}, "readFullBytes")
		}()
	case "readInt16":
		return func() *hxrt.Function[func() int] {
			hx_receiver_351 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_351.__hx_this.readInt16()
			}, func() any {
				return hx_receiver_351.__hx_this
			}, "readInt16")
		}()
	case "readInt24":
		return func() *hxrt.Function[func() int] {
			hx_receiver_352 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_352.__hx_this.readInt24()
			}, func() any {
				return hx_receiver_352.__hx_this
			}, "readInt24")
		}()
	case "readInt32":
		return func() *hxrt.Function[func() int] {
			hx_receiver_353 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_353.__hx_this.readInt32()
			}, func() any {
				return hx_receiver_353.__hx_this
			}, "readInt32")
		}()
	case "readInt8":
		return func() *hxrt.Function[func() int] {
			hx_receiver_354 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_354.__hx_this.readInt8()
			}, func() any {
				return hx_receiver_354.__hx_this
			}, "readInt8")
		}()
	case "readLine":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_355 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_355.__hx_this.readLine()
			}, func() any {
				return hx_receiver_355.__hx_this
			}, "readLine")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
			hx_receiver_356 := value
			return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_356.__hx_this.readString(len, encoding)
			}, func() any {
				return hx_receiver_356.__hx_this
			}, "readString")
		}()
	case "readUInt16":
		return func() *hxrt.Function[func() int] {
			hx_receiver_357 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_357.__hx_this.readUInt16()
			}, func() any {
				return hx_receiver_357.__hx_this
			}, "readUInt16")
		}()
	case "readUInt24":
		return func() *hxrt.Function[func() int] {
			hx_receiver_358 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_358.__hx_this.readUInt24()
			}, func() any {
				return hx_receiver_358.__hx_this
			}, "readUInt24")
		}()
	case "readUntil":
		return func() *hxrt.Function[func(int) *string] {
			hx_receiver_359 := value
			return hxrt.BoundFunction(func(end int) *string {
				return hx_receiver_359.__hx_this.readUntil(end)
			}, func() any {
				return hx_receiver_359.__hx_this
			}, "readUntil")
		}()
	case "set_bigEndian":
		return func() *hxrt.Function[func(bool) bool] {
			hx_receiver_360 := value
			return hxrt.BoundFunction(func(value bool) bool {
				return hx_receiver_360.__hx_this.set_bigEndian(value)
			}, func() any {
				return hx_receiver_360.__hx_this
			}, "set_bigEndian")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Output(value *haxe__io__Output, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return func() *hxrt.Function[func()] {
			hx_receiver_361 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_361.__hx_this.close()
			}, func() any {
				return hx_receiver_361.__hx_this
			}, "close")
		}()
	case "flush":
		return func() *hxrt.Function[func()] {
			hx_receiver_362 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_362.__hx_this.flush()
			}, func() any {
				return hx_receiver_362.__hx_this
			}, "flush")
		}()
	case "prepare":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_363 := value
			return hxrt.BoundFunction(func(nbytes int) {
				hx_receiver_363.__hx_this.prepare(nbytes)
			}, func() any {
				return hx_receiver_363.__hx_this
			}, "prepare")
		}()
	case "set_bigEndian":
		return func() *hxrt.Function[func(bool) bool] {
			hx_receiver_364 := value
			return hxrt.BoundFunction(func(value bool) bool {
				return hx_receiver_364.__hx_this.set_bigEndian(value)
			}, func() any {
				return hx_receiver_364.__hx_this
			}, "set_bigEndian")
		}()
	case "write":
		return func() *hxrt.Function[func(*haxe__io__Bytes)] {
			hx_receiver_365 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
				hx_receiver_365.__hx_this.write(bytes)
			}, func() any {
				return hx_receiver_365.__hx_this
			}, "write")
		}()
	case "writeByte":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_366 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_366.__hx_this.writeByte(value)
			}, func() any {
				return hx_receiver_366.__hx_this
			}, "writeByte")
		}()
	case "writeBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_367 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
				return hx_receiver_367.__hx_this.writeBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_367.__hx_this
			}, "writeBytes")
		}()
	case "writeDouble":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_368 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_368.__hx_this.writeDouble(value)
			}, func() any {
				return hx_receiver_368.__hx_this
			}, "writeDouble")
		}()
	case "writeFloat":
		return func() *hxrt.Function[func(float64)] {
			hx_receiver_369 := value
			return hxrt.BoundFunction(func(value float64) {
				hx_receiver_369.__hx_this.writeFloat(value)
			}, func() any {
				return hx_receiver_369.__hx_this
			}, "writeFloat")
		}()
	case "writeFullBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
			hx_receiver_370 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
				hx_receiver_370.__hx_this.writeFullBytes(bytes, pos, len)
			}, func() any {
				return hx_receiver_370.__hx_this
			}, "writeFullBytes")
		}()
	case "writeInput":
		return func() *hxrt.Function[func(*haxe__io__Input, any)] {
			hx_receiver_371 := value
			return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
				hx_receiver_371.__hx_this.writeInput(input, bufsize)
			}, func() any {
				return hx_receiver_371.__hx_this
			}, "writeInput")
		}()
	case "writeInt16":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_372 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_372.__hx_this.writeInt16(value)
			}, func() any {
				return hx_receiver_372.__hx_this
			}, "writeInt16")
		}()
	case "writeInt24":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_373 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_373.__hx_this.writeInt24(value)
			}, func() any {
				return hx_receiver_373.__hx_this
			}, "writeInt24")
		}()
	case "writeInt32":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_374 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_374.__hx_this.writeInt32(value)
			}, func() any {
				return hx_receiver_374.__hx_this
			}, "writeInt32")
		}()
	case "writeInt8":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_375 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_375.__hx_this.writeInt8(value)
			}, func() any {
				return hx_receiver_375.__hx_this
			}, "writeInt8")
		}()
	case "writeString":
		return func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
			hx_receiver_376 := value
			return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
				hx_receiver_376.__hx_this.writeString(value, encoding)
			}, func() any {
				return hx_receiver_376.__hx_this
			}, "writeString")
		}()
	case "writeUInt16":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_377 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_377.__hx_this.writeUInt16(value)
			}, func() any {
				return hx_receiver_377.__hx_this
			}, "writeUInt16")
		}()
	case "writeUInt24":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_378 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_378.__hx_this.writeUInt24(value)
			}, func() any {
				return hx_receiver_378.__hx_this
			}, "writeUInt24")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_379 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_379.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_379.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_380 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_380.__hx_this.next()
			}, func() any {
				return hx_receiver_380.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_381 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_381.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_381.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() int] {
			hx_receiver_382 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_382.__hx_this.next()
			}, func() any {
				return hx_receiver_382.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_383 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_383.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_383.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_384 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_384.__hx_this.next()
			}, func() any {
				return hx_receiver_384.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__sys__io__FileInput(value *sys__io__FileInput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return func() *hxrt.Function[func()] {
			hx_receiver_385 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_385.__hx_this.close()
			}, func() any {
				return hx_receiver_385.__hx_this
			}, "close")
		}()
	case "eof":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_386 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_386.__hx_this.eof()
			}, func() any {
				return hx_receiver_386.__hx_this
			}, "eof")
		}()
	case "readByte":
		return func() *hxrt.Function[func() int] {
			hx_receiver_387 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_387.__hx_this.readByte()
			}, func() any {
				return hx_receiver_387.__hx_this
			}, "readByte")
		}()
	case "readBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_388 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
				return hx_receiver_388.__hx_this.readBytes(bytes, pos, length)
			}, func() any {
				return hx_receiver_388.__hx_this
			}, "readBytes")
		}()
	case "seek":
		return func() *hxrt.Function[func(int, *sys__io__FileSeek)] {
			hx_receiver_389 := value
			return hxrt.BoundFunction(func(p int, pos *sys__io__FileSeek) {
				hx_receiver_389.__hx_this.seek(p, pos)
			}, func() any {
				return hx_receiver_389.__hx_this
			}, "seek")
		}()
	case "tell":
		return func() *hxrt.Function[func() int] {
			hx_receiver_390 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_390.__hx_this.tell()
			}, func() any {
				return hx_receiver_390.__hx_this
			}, "tell")
		}()
	}
	if value.haxe__io__Input == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__io__Input(value.haxe__io__Input, key)
}

func hxrt__generated_method_field__sys__io__FileOutput(value *sys__io__FileOutput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "close":
		return func() *hxrt.Function[func()] {
			hx_receiver_391 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_391.__hx_this.close()
			}, func() any {
				return hx_receiver_391.__hx_this
			}, "close")
		}()
	case "flush":
		return func() *hxrt.Function[func()] {
			hx_receiver_392 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_392.__hx_this.flush()
			}, func() any {
				return hx_receiver_392.__hx_this
			}, "flush")
		}()
	case "seek":
		return func() *hxrt.Function[func(int, *sys__io__FileSeek)] {
			hx_receiver_393 := value
			return hxrt.BoundFunction(func(p int, pos *sys__io__FileSeek) {
				hx_receiver_393.__hx_this.seek(p, pos)
			}, func() any {
				return hx_receiver_393.__hx_this
			}, "seek")
		}()
	case "tell":
		return func() *hxrt.Function[func() int] {
			hx_receiver_394 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_394.__hx_this.tell()
			}, func() any {
				return hx_receiver_394.__hx_this
			}, "tell")
		}()
	case "writeByte":
		return func() *hxrt.Function[func(int)] {
			hx_receiver_395 := value
			return hxrt.BoundFunction(func(value int) {
				hx_receiver_395.__hx_this.writeByte(value)
			}, func() any {
				return hx_receiver_395.__hx_this
			}, "writeByte")
		}()
	case "writeBytes":
		return func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
			hx_receiver_396 := value
			return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
				return hx_receiver_396.__hx_this.writeBytes(bytes, pos, length)
			}, func() any {
				return hx_receiver_396.__hx_this
			}, "writeBytes")
		}()
	}
	if value.haxe__io__Output == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__io__Output(value.haxe__io__Output, key)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	classValue, ok := value.(*hxrt__TypeClassValue)
	if !ok || classValue == nil {
		return nil, false
	}
	className := *hxrt.StdString(classValue.name)
	switch className {
	default:
		return nil, false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedField(object any, field *string) any {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *_Main__Delivery:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__NotImplementedException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__PosException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__BytesBuffer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Input:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Output:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Path:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__io__FileInput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__io__FileOutput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_field_lookup__Date(value, key)
	case *_Main__Delivery:
		return hxrt__generated_field_lookup___Main__Delivery(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_lookup__haxe__ds__StringMap(value, key)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_field_lookup__haxe__exceptions__NotImplementedException(value, key)
	case *haxe__exceptions__PosException:
		return hxrt__generated_field_lookup__haxe__exceptions__PosException(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_lookup__haxe__io__Bytes(value, key)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_field_lookup__haxe__io__BytesBuffer(value, key)
	case *haxe__io__Input:
		return hxrt__generated_field_lookup__haxe__io__Input(value, key)
	case *haxe__io__Output:
		return hxrt__generated_field_lookup__haxe__io__Output(value, key)
	case *haxe__io__Path:
		return hxrt__generated_field_lookup__haxe__io__Path(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value, key)
	case *sys__io__FileInput:
		return hxrt__generated_field_lookup__sys__io__FileInput(value, key)
	case *sys__io__FileOutput:
		return hxrt__generated_field_lookup__sys__io__FileOutput(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *_Main__Delivery:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__exceptions__NotImplementedException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__exceptions__PosException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__BytesBuffer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Input:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Output:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Path:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__io__FileInput:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__io__FileOutput:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_field_has__Date(value, key)
	case *_Main__Delivery:
		return hxrt__generated_field_has___Main__Delivery(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_has__haxe__ds__StringMap(value, key)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_field_has__haxe__exceptions__NotImplementedException(value, key)
	case *haxe__exceptions__PosException:
		return hxrt__generated_field_has__haxe__exceptions__PosException(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_has__haxe__io__Bytes(value, key)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_field_has__haxe__io__BytesBuffer(value, key)
	case *haxe__io__Input:
		return hxrt__generated_field_has__haxe__io__Input(value, key)
	case *haxe__io__Output:
		return hxrt__generated_field_has__haxe__io__Output(value, key)
	case *haxe__io__Path:
		return hxrt__generated_field_has__haxe__io__Path(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_has__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value, key)
	case *sys__io__FileInput:
		return hxrt__generated_field_has__sys__io__FileInput(value, key)
	case *sys__io__FileOutput:
		return hxrt__generated_field_has__sys__io__FileOutput(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *_Main__Delivery:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__exceptions__NotImplementedException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__exceptions__PosException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__BytesBuffer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Input:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Output:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Path:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__io__FileInput:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *sys__io__FileOutput:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Date:
		return hxrt__generated_field_set__Date(value, key, incoming)
	case *_Main__Delivery:
		return hxrt__generated_field_set___Main__Delivery(value, key, incoming)
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_set__haxe__ds__StringMap(value, key, incoming)
	case *haxe__exceptions__NotImplementedException:
		return hxrt__generated_field_set__haxe__exceptions__NotImplementedException(value, key, incoming)
	case *haxe__exceptions__PosException:
		return hxrt__generated_field_set__haxe__exceptions__PosException(value, key, incoming)
	case *haxe__io__Bytes:
		return hxrt__generated_field_set__haxe__io__Bytes(value, key, incoming)
	case *haxe__io__BytesBuffer:
		return hxrt__generated_field_set__haxe__io__BytesBuffer(value, key, incoming)
	case *haxe__io__Input:
		return hxrt__generated_field_set__haxe__io__Input(value, key, incoming)
	case *haxe__io__Output:
		return hxrt__generated_field_set__haxe__io__Output(value, key, incoming)
	case *haxe__io__Path:
		return hxrt__generated_field_set__haxe__io__Path(value, key, incoming)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value, key, incoming)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_set__haxe__iterators__StringIterator(value, key, incoming)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value, key, incoming)
	case *sys__io__FileInput:
		return hxrt__generated_field_set__sys__io__FileInput(value, key, incoming)
	case *sys__io__FileOutput:
		return hxrt__generated_field_set__sys__io__FileOutput(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Date:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *_Main__Delivery:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__NotImplementedException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__exceptions__PosException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__BytesBuffer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Eof:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Input:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Output:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Path:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__io__FileInput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *sys__io__FileOutput:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *Date:
		return hxrt.NewArray(hxrt.StringFromLiteral("ms"))
	case *_Main__Delivery:
		return hxrt.NewArray(hxrt.StringFromLiteral("name"), hxrt.StringFromLiteral("state"))
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__StringMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__exceptions__NotImplementedException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("posInfos"))
	case *haxe__exceptions__PosException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("posInfos"))
	case *haxe__io__Bytes:
		return hxrt.NewArray(hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("__hx_dataExposed"))
	case *haxe__io__BytesBuffer:
		return hxrt.NewArray(hxrt.StringFromLiteral("b"))
	case *haxe__io__Eof:
		return hxrt.NewArray()
	case *haxe__io__Input:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"))
	case *haxe__io__Output:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"))
	case *haxe__io__Path:
		return hxrt.NewArray(hxrt.StringFromLiteral("dir"), hxrt.StringFromLiteral("file"), hxrt.StringFromLiteral("ext"), hxrt.StringFromLiteral("backslash"))
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("keys"))
	case *haxe__iterators__StringIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *sys__io__FileInput:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("handle"))
	case *sys__io__FileOutput:
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("handle"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__Date(value *Date, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "ms":
		return hxrt.ReadStoredField(value.ms)
	}
	return nil
}

func hxrt__generated_field_has__Date(value *Date, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "ms":
		return true
	}
	return false
}

func hxrt__generated_field_set__Date(value *Date, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "ms":
		if incoming == nil {
			var zero float64
			value.ms = zero
			return true
		}
		switch typed := incoming.(type) {
		case float64:
			value.ms = typed
			return true
		case int:
			value.ms = float64(typed)
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup___Main__Delivery(value *_Main__Delivery, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "name":
		return hxrt.ReadStoredField(value.name)
	case "state":
		return hxrt.ReadStoredField(value.state)
	}
	return nil
}

func hxrt__generated_field_has___Main__Delivery(value *_Main__Delivery, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "name":
		return true
	case "state":
		return true
	}
	return false
}

func hxrt__generated_field_set___Main__Delivery(value *_Main__Delivery, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "name":
		if incoming == nil {
			var zero *string
			value.name = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.name = typed
			return true
		default:
			return false
		}
	case "state":
		if incoming == nil {
			var zero *_Main__DeliveryState
			value.state = zero
			return true
		}
		switch typed := incoming.(type) {
		case *_Main__DeliveryState:
			value.state = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "savedMessage":
		return hxrt.ReadStoredField(value.savedMessage)
	case "savedNative":
		return hxrt.ReadStoredField(value.savedNative)
	case "savedPrevious":
		return hxrt.ReadStoredField(value.savedPrevious)
	case "savedStack":
		return hxrt.ReadStoredField(value.savedStack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Exception(value *haxe__Exception, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		return true
	case "savedNative":
		return true
	case "savedPrevious":
		return true
	case "savedStack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Exception(value *haxe__Exception, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		if incoming == nil {
			var zero *string
			value.savedMessage = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.savedMessage = typed
			return true
		default:
			return false
		}
	case "savedNative":
		if incoming == nil {
			var zero any
			value.savedNative = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedNative = typed
			return true
		default:
			return false
		}
	case "savedPrevious":
		if incoming == nil {
			var zero *haxe__Exception
			value.savedPrevious = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__Exception:
			value.savedPrevious = typed
			return true
		default:
			return false
		}
	case "savedStack":
		if incoming == nil {
			var zero any
			value.savedStack = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedStack = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__ValueException(value *haxe__ValueException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__ValueException(value *haxe__ValueException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		if incoming == nil {
			var zero any
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.value = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return hxrt.ReadStoredField(value.high)
	case "low":
		return hxrt.ReadStoredField(value.low)
	}
	return nil
}

func hxrt__generated_field_has__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		return true
	case "low":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		if incoming == nil {
			var zero int
			value.high = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.high = typed
			return true
		default:
			return false
		}
	case "low":
		if incoming == nil {
			var zero int
			value.low = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.low = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "h":
		return hxrt.ReadStoredField(value.h)
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__StringMap(value *haxe__ds__StringMap, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		if incoming == nil {
			var zero *hxrt.StringMapCell
			value.h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.StringMapCell:
			value.h = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__exceptions__NotImplementedException(value *haxe__exceptions__NotImplementedException, key string) any {
	if value == nil {
		return nil
	}
	if value.haxe__exceptions__PosException == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__exceptions__PosException(value.haxe__exceptions__PosException, key)
}

func hxrt__generated_field_has__haxe__exceptions__NotImplementedException(value *haxe__exceptions__NotImplementedException, key string) bool {
	if value == nil {
		return false
	}
	if value.haxe__exceptions__PosException == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__exceptions__PosException(value.haxe__exceptions__PosException, key)
}

func hxrt__generated_field_set__haxe__exceptions__NotImplementedException(value *haxe__exceptions__NotImplementedException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	if value.haxe__exceptions__PosException == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__exceptions__PosException(value.haxe__exceptions__PosException, key, incoming)
}

func hxrt__generated_field_lookup__haxe__exceptions__PosException(value *haxe__exceptions__PosException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "posInfos":
		return hxrt.ReadStoredField(value.posInfos)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__exceptions__PosException(value *haxe__exceptions__PosException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "posInfos":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__exceptions__PosException(value *haxe__exceptions__PosException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "posInfos":
		if incoming == nil {
			var zero hxrt.Anonymous
			value.posInfos = zero
			return true
		}
		switch typed := incoming.(type) {
		case hxrt.Anonymous:
			value.posInfos = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_dataExposed":
		return hxrt.ReadStoredField(value.__hx_dataExposed)
	case "__hx_raw":
		return hxrt.ReadStoredField(value.__hx_raw)
	case "__hx_rawValid":
		return hxrt.ReadStoredField(value.__hx_rawValid)
	case "b":
		return hxrt.ReadStoredField(value.b)
	case "length":
		return hxrt.ReadStoredField(value.length)
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__Bytes(value *haxe__io__Bytes, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__hx_dataExposed":
		return true
	case "__hx_raw":
		return true
	case "__hx_rawValid":
		return true
	case "b":
		return true
	case "length":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__Bytes(value *haxe__io__Bytes, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__hx_dataExposed":
		if incoming == nil {
			var zero bool
			value.__hx_dataExposed = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.__hx_dataExposed = typed
			return true
		default:
			return false
		}
	case "__hx_raw":
		if incoming == nil {
			var zero *hxrt.ByteView
			value.__hx_raw = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.ByteView:
			value.__hx_raw = typed
			return true
		default:
			return false
		}
	case "__hx_rawValid":
		if incoming == nil {
			var zero bool
			value.__hx_rawValid = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.__hx_rawValid = typed
			return true
		default:
			return false
		}
	case "b":
		if incoming == nil {
			var zero []int
			value.b = zero
			return true
		}
		switch typed := incoming.(type) {
		case []int:
			value.b = typed
			return true
		default:
			return false
		}
	case "length":
		if incoming == nil {
			var zero int
			value.length = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.length = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__io__BytesBuffer(value *haxe__io__BytesBuffer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "b":
		return hxrt.ReadStoredField(value.b)
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__BytesBuffer(value *haxe__io__BytesBuffer, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__BytesBuffer(value *haxe__io__BytesBuffer, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		if incoming == nil {
			var zero []int
			value.b = zero
			return true
		}
		switch typed := incoming.(type) {
		case []int:
			value.b = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__io__Input(value *haxe__io__Input, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "bigEndian":
		return hxrt.ReadStoredField(value.bigEndian)
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__Input(value *haxe__io__Input, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "bigEndian":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__Input(value *haxe__io__Input, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "bigEndian":
		if incoming == nil {
			var zero bool
			value.bigEndian = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.bigEndian = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__io__Output(value *haxe__io__Output, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "bigEndian":
		return hxrt.ReadStoredField(value.bigEndian)
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__Output(value *haxe__io__Output, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "bigEndian":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__Output(value *haxe__io__Output, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "bigEndian":
		if incoming == nil {
			var zero bool
			value.bigEndian = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.bigEndian = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__io__Path(value *haxe__io__Path, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "backslash":
		return hxrt.ReadStoredField(value.backslash)
	case "dir":
		return hxrt.ReadStoredField(value.dir)
	case "ext":
		return hxrt.ReadStoredField(value.ext)
	case "file":
		return hxrt.ReadStoredField(value.file)
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__Path(value *haxe__io__Path, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "backslash":
		return true
	case "dir":
		return true
	case "ext":
		return true
	case "file":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__Path(value *haxe__io__Path, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "backslash":
		if incoming == nil {
			var zero bool
			value.backslash = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.backslash = typed
			return true
		default:
			return false
		}
	case "dir":
		if incoming == nil {
			var zero *string
			value.dir = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.dir = typed
			return true
		default:
			return false
		}
	case "ext":
		if incoming == nil {
			var zero *string
			value.ext = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.ext = typed
			return true
		default:
			return false
		}
	case "file":
		if incoming == nil {
			var zero *string
			value.file = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.file = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "keys":
		return hxrt.ReadStoredField(value.keys)
	case "map":
		return hxrt.ReadStoredField(value.map_)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		return true
	case "map":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		if incoming == nil {
			var zero hxrt.Anonymous
			value.keys = zero
			return true
		}
		switch typed := incoming.(type) {
		case hxrt.Anonymous:
			value.keys = typed
			return true
		default:
			return false
		}
	case "map":
		if incoming == nil {
			var zero haxe__IMap
			value.map_ = zero
			return true
		}
		switch typed := incoming.(type) {
		case haxe__IMap:
			value.map_ = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "offset":
		return hxrt.ReadStoredField(value.offset)
	case "s":
		return hxrt.ReadStoredField(value.s)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		return true
	case "s":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		if incoming == nil {
			var zero int
			value.offset = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.offset = typed
			return true
		default:
			return false
		}
	case "s":
		if incoming == nil {
			var zero *string
			value.s = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.s = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "offset":
		return hxrt.ReadStoredField(value.offset)
	case "s":
		return hxrt.ReadStoredField(value.s)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		return true
	case "s":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		if incoming == nil {
			var zero int
			value.offset = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.offset = typed
			return true
		default:
			return false
		}
	case "s":
		if incoming == nil {
			var zero *string
			value.s = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.s = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__sys__io__FileInput(value *sys__io__FileInput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "handle":
		return hxrt.ReadStoredField(value.handle)
	}
	if value.haxe__io__Input == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__io__Input(value.haxe__io__Input, key)
}

func hxrt__generated_field_has__sys__io__FileInput(value *sys__io__FileInput, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "handle":
		return true
	}
	if value.haxe__io__Input == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__io__Input(value.haxe__io__Input, key)
}

func hxrt__generated_field_set__sys__io__FileInput(value *sys__io__FileInput, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "handle":
		if incoming == nil {
			var zero *hxrt.FileInput
			value.handle = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.FileInput:
			value.handle = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__io__Input == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__io__Input(value.haxe__io__Input, key, incoming)
}

func hxrt__generated_field_lookup__sys__io__FileOutput(value *sys__io__FileOutput, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "handle":
		return hxrt.ReadStoredField(value.handle)
	}
	if value.haxe__io__Output == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__io__Output(value.haxe__io__Output, key)
}

func hxrt__generated_field_has__sys__io__FileOutput(value *sys__io__FileOutput, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "handle":
		return true
	}
	if value.haxe__io__Output == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__io__Output(value.haxe__io__Output, key)
}

func hxrt__generated_field_set__sys__io__FileOutput(value *sys__io__FileOutput, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "handle":
		if incoming == nil {
			var zero *hxrt.FileOutput
			value.handle = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.FileOutput:
			value.handle = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__io__Output == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__io__Output(value.haxe__io__Output, key, incoming)
}

func reflaxe__go___internal__CompilerReflect_typeField(object any, field *string) any {
	key := *hxrt.StdString(field)
	value, found := hxrt_typeClassMetadataField(object, key)
	if !found {
		return nil
	}
	return value
}

func reflaxe__go___internal__CompilerReflect_hasTypeField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	_, found := hxrt_typeClassMetadataField(object, key)
	return found
}

func reflaxe__go___internal__CompilerReflect_generatedMethod(object any, field *string) any {
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	switch enumValue := value.(type) {
	case *_Main__DeliveryState:
		return (enumValue != nil)
	case *haxe__StackItem:
		return (enumValue != nil)
	case *haxe__io__Encoding:
		return (enumValue != nil)
	case *haxe__io__Error:
		return (enumValue != nil)
	case *sys__io__FileSeek:
		return (enumValue != nil)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_397 any) *string {
			if hx_value_397 == nil {
				var hx_zero_398 *string
				return hx_zero_398
			}
			return hx_value_397.(*string)
		}(any(argument_0)))
	case *hxrt.Function[func() bool]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() float64]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *hxrt.DateParts]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func()]:
		hxrt.FunctionCall(source)()
		return nil
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*string) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_399 any) *string {
			if hx_value_399 == nil {
				var hx_zero_400 *string
				return hx_zero_400
			}
			return hx_value_399.(*string)
		}(any(argument_0)))
	case *hxrt.Function[func(any) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(*string) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_401 any) *string {
			if hx_value_401 == nil {
				var hx_zero_402 *string
				return hx_zero_402
			}
			return hx_value_401.(*string)
		}(any(argument_0)))
	case *hxrt.Function[func(any) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(*string, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_403 any) *string {
			if hx_value_403 == nil {
				var hx_zero_404 *string
				return hx_zero_404
			}
			return hx_value_403.(*string)
		}(any(argument_0)), argument_1)
		return nil
	case *hxrt.Function[func(any, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(argument_0, argument_1)
		return nil
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int, *haxe__io__Bytes, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		var argument_3 any = nil
		if 3 < len(arguments) {
			argument_3 = arguments[3]
		}
		hxrt.FunctionCall(source)(func(hx_value_405 any) int {
			if hx_value_405 == nil {
				var hx_zero_406 int
				return hx_zero_406
			}
			return hx_value_405.(int)
		}(any(argument_0)), hx_recover_class_5(argument_1), func(hx_value_407 any) int {
			if hx_value_407 == nil {
				var hx_zero_408 int
				return hx_zero_408
			}
			return hx_value_407.(int)
		}(any(argument_2)), func(hx_value_409 any) int {
			if hx_value_409 == nil {
				var hx_zero_410 int
				return hx_zero_410
			}
			return hx_value_409.(int)
		}(any(argument_3)))
		return nil
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_5(argument_0))
	case *hxrt.Function[func(int, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		hxrt.FunctionCall(source)(func(hx_value_411 any) int {
			if hx_value_411 == nil {
				var hx_zero_412 int
				return hx_zero_412
			}
			return hx_value_411.(int)
		}(any(argument_0)), func(hx_value_413 any) int {
			if hx_value_413 == nil {
				var hx_zero_414 int
				return hx_zero_414
			}
			return hx_value_413.(int)
		}(any(argument_1)), func(hx_value_415 any) int {
			if hx_value_415 == nil {
				var hx_zero_416 int
				return hx_zero_416
			}
			return hx_value_415.(int)
		}(any(argument_2)))
		return nil
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_417 any) int {
			if hx_value_417 == nil {
				var hx_zero_418 int
				return hx_zero_418
			}
			return hx_value_417.(int)
		}(any(argument_0)))
	case *hxrt.Function[func() []int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) float64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_419 any) int {
			if hx_value_419 == nil {
				var hx_zero_420 int
				return hx_zero_420
			}
			return hx_value_419.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_421 any) int {
			if hx_value_421 == nil {
				var hx_zero_422 int
				return hx_zero_422
			}
			return hx_value_421.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(int, int, *haxe__io__Encoding) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(func(hx_value_423 any) int {
			if hx_value_423 == nil {
				var hx_zero_424 int
				return hx_zero_424
			}
			return hx_value_423.(int)
		}(any(argument_0)), func(hx_value_425 any) int {
			if hx_value_425 == nil {
				var hx_zero_426 int
				return hx_zero_426
			}
			return hx_value_425.(int)
		}(any(argument_1)), func(hx_value_427 any) *haxe__io__Encoding {
			if hx_value_427 == nil {
				var hx_zero_428 *haxe__io__Encoding
				return hx_zero_428
			}
			return hx_value_427.(*haxe__io__Encoding)
		}(any(argument_2)))
	case *hxrt.Function[func(int, int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_429 any) int {
			if hx_value_429 == nil {
				var hx_zero_430 int
				return hx_zero_430
			}
			return hx_value_429.(int)
		}(any(argument_0)), func(hx_value_431 any) int {
			if hx_value_431 == nil {
				var hx_zero_432 int
				return hx_zero_432
			}
			return hx_value_431.(int)
		}(any(argument_1)))
	case *hxrt.Function[func(int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_433 any) int {
			if hx_value_433 == nil {
				var hx_zero_434 int
				return hx_zero_434
			}
			return hx_value_433.(int)
		}(any(argument_0)), func(hx_value_435 any) int {
			if hx_value_435 == nil {
				var hx_zero_436 int
				return hx_zero_436
			}
			return hx_value_435.(int)
		}(any(argument_1)))
		return nil
	case *hxrt.Function[func(int, float64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_437 any) int {
			if hx_value_437 == nil {
				var hx_zero_438 int
				return hx_zero_438
			}
			return hx_value_437.(int)
		}(any(argument_0)), func(hx_value_439 any) float64 {
			if hx_value_439 == nil {
				var hx_zero_440 float64
				return hx_zero_440
			}
			return hx_value_439.(float64)
		}(any(argument_1)))
		return nil
	case *hxrt.Function[func(int, *haxe___Int64_____Int64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_441 any) int {
			if hx_value_441 == nil {
				var hx_zero_442 int
				return hx_zero_442
			}
			return hx_value_441.(int)
		}(any(argument_0)), hx_recover_class_6(argument_1))
		return nil
	case *hxrt.Function[func(int, int) *haxe__io__Bytes]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_443 any) int {
			if hx_value_443 == nil {
				var hx_zero_444 int
				return hx_zero_444
			}
			return hx_value_443.(int)
		}(any(argument_0)), func(hx_value_445 any) int {
			if hx_value_445 == nil {
				var hx_zero_446 int
				return hx_zero_446
			}
			return hx_value_445.(int)
		}(any(argument_1)))
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_5(argument_0))
		return nil
	case *hxrt.Function[func(int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_447 any) int {
			if hx_value_447 == nil {
				var hx_zero_448 int
				return hx_zero_448
			}
			return hx_value_447.(int)
		}(any(argument_0)))
		return nil
	case *hxrt.Function[func(*haxe__io__Bytes, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		hxrt.FunctionCall(source)(hx_recover_class_5(argument_0), func(hx_value_449 any) int {
			if hx_value_449 == nil {
				var hx_zero_450 int
				return hx_zero_450
			}
			return hx_value_449.(int)
		}(any(argument_1)), func(hx_value_451 any) int {
			if hx_value_451 == nil {
				var hx_zero_452 int
				return hx_zero_452
			}
			return hx_value_451.(int)
		}(any(argument_2)))
		return nil
	case *hxrt.Function[func(float64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_453 any) float64 {
			if hx_value_453 == nil {
				var hx_zero_454 float64
				return hx_zero_454
			}
			return hx_value_453.(float64)
		}(any(argument_0)))
		return nil
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_6(argument_0))
		return nil
	case *hxrt.Function[func(*string, *haxe__io__Encoding)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_455 any) *string {
			if hx_value_455 == nil {
				var hx_zero_456 *string
				return hx_zero_456
			}
			return hx_value_455.(*string)
		}(any(argument_0)), func(hx_value_457 any) *haxe__io__Encoding {
			if hx_value_457 == nil {
				var hx_zero_458 *haxe__io__Encoding
				return hx_zero_458
			}
			return hx_value_457.(*haxe__io__Encoding)
		}(any(argument_1)))
		return nil
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_459 any) int {
			if hx_value_459 == nil {
				var hx_zero_460 int
				return hx_zero_460
			}
			return hx_value_459.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(*haxe__io__Bytes, int, int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_5(argument_0), func(hx_value_461 any) int {
			if hx_value_461 == nil {
				var hx_zero_462 int
				return hx_zero_462
			}
			return hx_value_461.(int)
		}(any(argument_1)), func(hx_value_463 any) int {
			if hx_value_463 == nil {
				var hx_zero_464 int
				return hx_zero_464
			}
			return hx_value_463.(int)
		}(any(argument_2)))
	case *hxrt.Function[func(int, *haxe__io__Encoding) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_465 any) int {
			if hx_value_465 == nil {
				var hx_zero_466 int
				return hx_zero_466
			}
			return hx_value_465.(int)
		}(any(argument_0)), func(hx_value_467 any) *haxe__io__Encoding {
			if hx_value_467 == nil {
				var hx_zero_468 *haxe__io__Encoding
				return hx_zero_468
			}
			return hx_value_467.(*haxe__io__Encoding)
		}(any(argument_1)))
	case *hxrt.Function[func(int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_469 any) int {
			if hx_value_469 == nil {
				var hx_zero_470 int
				return hx_zero_470
			}
			return hx_value_469.(int)
		}(any(argument_0)))
	case *hxrt.Function[func(bool) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_471 any) bool {
			if hx_value_471 == nil {
				var hx_zero_472 bool
				return hx_zero_472
			}
			return hx_value_471.(bool)
		}(any(argument_0)))
	case *hxrt.Function[func(*haxe__io__Input, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(hx_recover_class_7(argument_0), argument_1)
		return nil
	case *hxrt.Function[func(int, *sys__io__FileSeek)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_473 any) int {
			if hx_value_473 == nil {
				var hx_zero_474 int
				return hx_zero_474
			}
			return hx_value_473.(int)
		}(any(argument_0)), func(hx_value_475 any) *sys__io__FileSeek {
			if hx_value_475 == nil {
				var hx_zero_476 *sys__io__FileSeek
				return hx_zero_476
			}
			return hx_value_475.(*sys__io__FileSeek)
		}(any(argument_1)))
		return nil
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_477 any) bool {
				if hx_value_477 == nil {
					var hx_zero_478 bool
					return hx_zero_478
				}
				return hx_value_477.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_479 any) bool {
				if hx_value_479 == nil {
					var hx_zero_480 bool
					return hx_zero_480
				}
				return hx_value_479.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_481 any) bool {
				if hx_value_481 == nil {
					var hx_zero_482 bool
					return hx_zero_482
				}
				return hx_value_481.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_483 any) bool {
				if hx_value_483 == nil {
					var hx_zero_484 bool
					return hx_zero_484
				}
				return hx_value_483.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_485 any) bool {
				if hx_value_485 == nil {
					var hx_zero_486 bool
					return hx_zero_486
				}
				return hx_value_485.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_487 any) bool {
				if hx_value_487 == nil {
					var hx_zero_488 bool
					return hx_zero_488
				}
				return hx_value_487.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *hxrt.DateParts]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_489 any) bool {
				if hx_value_489 == nil {
					var hx_zero_490 bool
					return hx_zero_490
				}
				return hx_value_489.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_491 any) bool {
				if hx_value_491 == nil {
					var hx_zero_492 bool
					return hx_zero_492
				}
				return hx_value_491.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_493 any) bool {
				if hx_value_493 == nil {
					var hx_zero_494 bool
					return hx_zero_494
				}
				return hx_value_493.(bool)
			}(any(nil))
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_495 any) bool {
				if hx_value_495 == nil {
					var hx_zero_496 bool
					return hx_zero_496
				}
				return hx_value_495.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_497 any) bool {
				if hx_value_497 == nil {
					var hx_zero_498 bool
					return hx_zero_498
				}
				return hx_value_497.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_499 any) bool {
				if hx_value_499 == nil {
					var hx_zero_500 bool
					return hx_zero_500
				}
				return hx_value_499.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_501 any) bool {
				if hx_value_501 == nil {
					var hx_zero_502 bool
					return hx_zero_502
				}
				return hx_value_501.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_503 any) bool {
				if hx_value_503 == nil {
					var hx_zero_504 bool
					return hx_zero_504
				}
				return hx_value_503.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.DateParts]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_505 any) *string {
				if hx_value_505 == nil {
					var hx_zero_506 *string
					return hx_zero_506
				}
				return hx_value_505.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_507 any) *string {
				if hx_value_507 == nil {
					var hx_zero_508 *string
					return hx_zero_508
				}
				return hx_value_507.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_509 any) *string {
				if hx_value_509 == nil {
					var hx_zero_510 *string
					return hx_zero_510
				}
				return hx_value_509.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_511 any) *string {
				if hx_value_511 == nil {
					var hx_zero_512 *string
					return hx_zero_512
				}
				return hx_value_511.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_513 any) *string {
				if hx_value_513 == nil {
					var hx_zero_514 *string
					return hx_zero_514
				}
				return hx_value_513.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_515 any) *string {
				if hx_value_515 == nil {
					var hx_zero_516 *string
					return hx_zero_516
				}
				return hx_value_515.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *hxrt.DateParts]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_517 any) *string {
				if hx_value_517 == nil {
					var hx_zero_518 *string
					return hx_zero_518
				}
				return hx_value_517.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_519 any) *string {
				if hx_value_519 == nil {
					var hx_zero_520 *string
					return hx_zero_520
				}
				return hx_value_519.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_521 any) *string {
				if hx_value_521 == nil {
					var hx_zero_522 *string
					return hx_zero_522
				}
				return hx_value_521.(*string)
			}(any(nil))
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_523 any) *string {
				if hx_value_523 == nil {
					var hx_zero_524 *string
					return hx_zero_524
				}
				return hx_value_523.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_525 any) *string {
				if hx_value_525 == nil {
					var hx_zero_526 *string
					return hx_zero_526
				}
				return hx_value_525.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_527 any) *string {
				if hx_value_527 == nil {
					var hx_zero_528 *string
					return hx_zero_528
				}
				return hx_value_527.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_529 any) *string {
				if hx_value_529 == nil {
					var hx_zero_530 *string
					return hx_zero_530
				}
				return hx_value_529.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_531 any) *string {
				if hx_value_531 == nil {
					var hx_zero_532 *string
					return hx_zero_532
				}
				return hx_value_531.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_1(value any) *_Main__Delivery {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *_Main__Delivery:
		return source
	default:
		return value.(*_Main__Delivery)
	}
}

func hx_recover_class_2(value any) *haxe__io__Eof {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Eof:
		return source
	default:
		return value.(*haxe__io__Eof)
	}
}

func hx_recover_class_3(value any) *haxe__ds__StringMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return source
	default:
		return value.(*haxe__ds__StringMap)
	}
}

func hx_recover_class_4(value any) *haxe__ValueException {
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

func hx_recover_class_5(value any) *haxe__io__Bytes {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return source
	default:
		return value.(*haxe__io__Bytes)
	}
}

func hx_recover_class_6(value any) *haxe___Int64_____Int64 {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return source
	default:
		return value.(*haxe___Int64_____Int64)
	}
}

func hx_recover_class_7(value any) *haxe__io__Input {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__io__FileInput:
		return source.haxe__io__Input
	case *haxe__io__Input:
		return source
	default:
		return value.(*haxe__io__Input)
	}
}
