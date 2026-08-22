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

func New_app__core__IncidentApi(config *app__core__IncidentConfig, store *app__core__IncidentStore) *app__core__IncidentApi {
	self := &app__core__IncidentApi{}
	self.__hx_this = self
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
			hxrt.Throw(New_app__core__IncidentRequestException(hxrt.StringFromLiteral("missing_title")))
		}
		severity := app__core__IncidentApi_fieldString(raw, hxrt.StringFromLiteral("severity"), hxrt.StringFromLiteral("low"))
		incident := self.store.create(title, severity)
		response = app__http__HttpResponse_json(201, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{\"incident\":"), incident.toJson()), hxrt.StringFromLiteral("}")))
	}, func(hx_caught_1 any) {
		switch hx_typed_2 := hx_caught_1.(type) {
		case *app__core__IncidentRequestException:
			error := hx_typed_2
			response = app__http__HttpResponse_json(400, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{\"error\":\""), error.code), hxrt.StringFromLiteral("\"}")))
		default:
			hxrt.Throw(hx_caught_1)
		}
	})
	return response
}

func (self *app__core__IncidentApi) updateIncident(path *string, action *string) *app__http__HttpResponse {
	parts := hxrt.ArrayFromValues(func(hx_sort_src_3 []*string) []any {
		hx_sort_out_5 := make([]any, 0, len(hx_sort_src_3))
		for _, hx_sort_item_4 := range hx_sort_src_3 {
			hx_sort_out_5 = append(hx_sort_out_5, hx_sort_item_4)
		}
		return hx_sort_out_5
	}(hxrt.StringSplitStringPtr(path, hxrt.StringFromLiteral("/"))))
	if parts.Len() < 4 {
		return app__http__HttpResponse_json(404, hxrt.StringFromLiteral("{\"error\":\"not_found\"}"))
	}
	var id any = Std_parseInt(func(hx_value_6 any) *string {
		if hx_value_6 == nil {
			var hx_zero_7 *string
			return hx_zero_7
		}
		return hx_value_6.(*string)
	}(parts.Get(2)))
	if id == nil {
		return app__http__HttpResponse_json(400, hxrt.StringFromLiteral("{\"error\":\"invalid_id\"}"))
	}
	var hx_if_8 *app__core__Incident
	if hxrt.StringEqualStringPtr(action, hxrt.StringFromLiteral("ack")) {
		hx_if_8 = self.store.acknowledge(id.(int))
	} else {
		hx_if_8 = self.store.resolve(id.(int))
	}
	incident := hx_if_8
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
	var hx_if_9 *string
	if hxrt.AnyEqualsNull(value) {
		hx_if_9 = fallback
	} else {
		hx_if_9 = hxrt.StdString(value)
	}
	return hx_if_9
}

func app__core__IncidentApi_parseJsonBody(body *string) any {
	hx_try_return_10 := false
	var hx_try_value_11 any
	hxrt.TryCatch(func() {
		hx_try_value_11 = hxrt.JsonParse(func() *string {
			var hx_if_14 *string
			if hxrt.StringEqualStringPtr(body, hxrt.StringFromLiteral("")) {
				hx_if_14 = hxrt.StringFromLiteral("{}")
			} else {
				hx_if_14 = body
			}
			return hx_if_14
		}())
		hx_try_return_10 = true
		return
	}, func(hx_caught_12 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_12)
		_ = hx_tmp
		hxrt.Throw(New_app__core__IncidentRequestException(hxrt.StringFromLiteral("invalid_json")))
	})
	if hx_try_return_10 {
		return hx_try_value_11
	}
	return nil
}
