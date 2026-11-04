package main

import "examples_incident_api_portable/hxrt"

type I_app__core__IncidentApi interface {
	handle(request *app__http__HttpRequest) *app__http__HttpResponse
	createIncident(body *string) *app__http__HttpResponse
	updateIncident(path *string, action *string) *app__http__HttpResponse
}

type app__core__IncidentApi struct {
	__hx_this I_app__core__IncidentApi
	config    *app__core__IncidentConfig
	store     *app__core__IncidentStore
	requests  int
}

func (self *app__core__IncidentApi) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__core__IncidentApi(config *app__core__IncidentConfig, store *app__core__IncidentStore) *app__core__IncidentApi {
	return hx_init_app__core__IncidentApi(hx_alloc_app__core__IncidentApi(nil), config, store)
}

func hx_alloc_app__core__IncidentApi(owner any) *app__core__IncidentApi {
	self := &app__core__IncidentApi{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__core__IncidentApi)
	}
	return self
}

func hx_init_app__core__IncidentApi(self *app__core__IncidentApi, config *app__core__IncidentConfig, store *app__core__IncidentStore) *app__core__IncidentApi {
	self.config = config
	self.store = store
	self.requests = 0
	return self
}

func (self *app__core__IncidentApi) handle(request *app__http__HttpRequest) *app__http__HttpResponse {
	self.requests = int(int32((self.requests + 1)))
	if hxrt.StringEqualStringPtr(request.method, hxrt.StringFromLiteral("GET")) && hxrt.StringEqualStringPtr(request.path, hxrt.StringFromLiteral("/health")) {
		return app__http__HttpResponse_json(200, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{\"ok\":true,\"service\":\""), app__core__Incident_jsonEscape(self.config.serviceName)), hxrt.StringFromLiteral("\"}")))
	}
	if hxrt.StringEqualStringPtr(request.method, hxrt.StringFromLiteral("GET")) && hxrt.StringEqualStringPtr(request.path, hxrt.StringFromLiteral("/incidents")) {
		return app__http__HttpResponse_json(200, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{\"incidents\":"), self.store.listJson()), hxrt.StringFromLiteral("}")))
	}
	if hxrt.StringEqualStringPtr(request.method, hxrt.StringFromLiteral("POST")) && hxrt.StringEqualStringPtr(request.path, hxrt.StringFromLiteral("/incidents")) {
		return self.createIncident(request.body)
	}
	if (hxrt.StringEqualStringPtr(request.method, hxrt.StringFromLiteral("POST")) && StringTools_startsWith(request.path, hxrt.StringFromLiteral("/incidents/"))) && StringTools_endsWith(request.path, hxrt.StringFromLiteral("/ack")) {
		return self.updateIncident(request.path, hxrt.StringFromLiteral("ack"))
	}
	if (hxrt.StringEqualStringPtr(request.method, hxrt.StringFromLiteral("POST")) && StringTools_startsWith(request.path, hxrt.StringFromLiteral("/incidents/"))) && StringTools_endsWith(request.path, hxrt.StringFromLiteral("/resolve")) {
		return self.updateIncident(request.path, hxrt.StringFromLiteral("resolve"))
	}
	if hxrt.StringEqualStringPtr(request.method, hxrt.StringFromLiteral("GET")) && hxrt.StringEqualStringPtr(request.path, hxrt.StringFromLiteral("/metrics")) {
		return app__http__HttpResponse_json(200, self.store.metricsJson(self.config.serviceName, self.requests))
	}
	return app__http__HttpResponse_json(404, hxrt.StringFromLiteral("{\"error\":\"not_found\"}"))
}

func (self *app__core__IncidentApi) createIncident(body *string) *app__http__HttpResponse {
	response := app__http__HttpResponse_json(400, hxrt.StringFromLiteral("{\"error\":\"invalid_json\"}"))
	hxrt.TryCatch(func() {
		var raw any = app__core__IncidentApi_parseJsonBody(body)
		title := app__core__IncidentApi_fieldString(raw, hxrt.StringFromLiteral("title"), hxrt.StringFromLiteral(""))
		if hxrt.StringEqualStringPtr(title, hxrt.StringFromLiteral("")) {
			hxrt.Throw(haxe__Exception_thrown(New_app__core__IncidentRequestException(hxrt.StringFromLiteral("missing_title"))))
		}
		severity := app__core__IncidentApi_fieldString(raw, hxrt.StringFromLiteral("severity"), hxrt.StringFromLiteral("low"))
		incident := self.store.create(title, severity)
		response = app__http__HttpResponse_json(201, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{\"incident\":"), incident.toJson()), hxrt.StringFromLiteral("}")))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_payload_3
		if hx_is_nominal_0(hx_exception_2) {
			var error *app__core__IncidentRequestException = hx_recover_class_1(hx_exception_2)
			response = app__http__HttpResponse_json(400, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{\"error\":\""), error.code), hxrt.StringFromLiteral("\"}")))
		} else {
			hxrt.Throw(hx_caught_1)
		}
	})
	return response
}

func (self *app__core__IncidentApi) updateIncident(path *string, action *string) *app__http__HttpResponse {
	parts := hxrt.ArrayFromValues(func(hx_sort_src_4 []*string) []any {
		hx_sort_out_6 := make([]any, 0, len(hx_sort_src_4))
		for _, hx_sort_item_5 := range hx_sort_src_4 {
			hx_sort_out_6 = append(hx_sort_out_6, hx_sort_item_5)
		}
		return hx_sort_out_6
	}(hxrt.StringSplitStringPtr(path, hxrt.StringFromLiteral("/"))))
	if parts.Len() < 4 {
		return app__http__HttpResponse_json(404, hxrt.StringFromLiteral("{\"error\":\"not_found\"}"))
	}
	var id any = Std_parseInt(func(hx_value_7 any) *string {
		if hx_value_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_value_7.(*string)
	}(parts.Get(2)))
	if id == nil {
		return app__http__HttpResponse_json(400, hxrt.StringFromLiteral("{\"error\":\"invalid_id\"}"))
	}
	var hx_if_9 *app__core__Incident
	if hxrt.StringEqualStringPtr(action, hxrt.StringFromLiteral("ack")) {
		hx_if_9 = self.store.acknowledge(id.(int))
	} else {
		hx_if_9 = self.store.resolve(id.(int))
	}
	incident := hx_if_9
	if incident == nil {
		return app__http__HttpResponse_json(404, hxrt.StringFromLiteral("{\"error\":\"incident_not_found\"}"))
	}
	return app__http__HttpResponse_json(200, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{\"incident\":"), incident.toJson()), hxrt.StringFromLiteral("}")))
}

func app__core__IncidentApi_fieldString(raw any, name *string, fallback *string) *string {
	if hxrt.AnyEqualsNull(raw) || !Reflect_hasField(raw, name) {
		return fallback
	}
	var value any = Reflect_field(raw, name)
	var hx_if_10 *string
	if hxrt.AnyEqualsNull(value) {
		hx_if_10 = fallback
	} else {
		hx_if_10 = hxrt.StdString(value)
	}
	return hx_if_10
}

func app__core__IncidentApi_parseJsonBody(body *string) any {
	hx_try_return_11 := false
	var hx_try_value_12 any
	hxrt.TryCatch(func() {
		hx_try_value_12 = hxrt.JsonParse(func() *string {
			var hx_if_16 *string
			if hxrt.StringEqualStringPtr(body, hxrt.StringFromLiteral("")) {
				hx_if_16 = hxrt.StringFromLiteral("{}")
			} else {
				hx_if_16 = body
			}
			return hx_if_16
		}())
		hx_try_return_11 = true
		return
	}, func(hx_caught_13 any) {
		hx_exception_14 := haxe__Exception_caught(hx_caught_13)
		hx_payload_15 := haxe__Exception_catchValue(hx_caught_13)
		_ = hx_payload_15
		var hx_tmp *haxe__Exception = hx_recover_class_0(hx_exception_14)
		_ = hx_tmp
		hxrt.Throw(haxe__Exception_thrown(New_app__core__IncidentRequestException(hxrt.StringFromLiteral("invalid_json"))))
	})
	if hx_try_return_11 {
		return hx_try_value_12
	}
	return nil
}
