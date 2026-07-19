package main

import "examples_incident_api_metal/hxrt"

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
	}, func(hx_caught_47 any) {
		switch hx_typed_48 := hx_caught_47.(type) {
		case *app__core__IncidentRequestException:
			error := hx_typed_48
			response = app__http__HttpResponse_json(400, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{\"error\":\""), error.code), hxrt.StringFromLiteral("\"}")))
		default:
			hxrt.Throw(hx_caught_47)
		}
	})
	return response
}

func (self *app__core__IncidentApi) updateIncident(path *string, action *string) *app__http__HttpResponse {
	parts := hxrt.ArrayFromValues(func(hx_sort_src_49 []*string) []any {
		hx_sort_out_51 := make([]any, 0, len(hx_sort_src_49))
		for _, hx_sort_item_50 := range hx_sort_src_49 {
			hx_sort_out_51 = append(hx_sort_out_51, hx_sort_item_50)
		}
		return hx_sort_out_51
	}(hxrt.StringSplitStringPtr(path, hxrt.StringFromLiteral("/"))))
	if parts.Len() < 4 {
		return app__http__HttpResponse_json(404, hxrt.StringFromLiteral("{\"error\":\"not_found\"}"))
	}
	var id any = Std_parseInt(func(hx_value_52 any) *string {
		if hx_value_52 == nil {
			var hx_zero_53 *string
			return hx_zero_53
		}
		return hx_value_52.(*string)
	}(parts.Get(2)))
	if id == nil {
		return app__http__HttpResponse_json(400, hxrt.StringFromLiteral("{\"error\":\"invalid_id\"}"))
	}
	var hx_if_54 *app__core__Incident
	if hxrt.StringEqualStringPtr(action, hxrt.StringFromLiteral("ack")) {
		hx_if_54 = self.store.acknowledge(id.(int))
	} else {
		hx_if_54 = self.store.resolve(id.(int))
	}
	incident := hx_if_54
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
	var hx_if_55 *string
	if hxrt.AnyEqualsNull(value) {
		hx_if_55 = fallback
	} else {
		hx_if_55 = hxrt.StdString(value)
	}
	return hx_if_55
}

func app__core__IncidentApi_parseJsonBody(body *string) any {
	hx_try_return_56 := false
	var hx_try_value_57 any
	hxrt.TryCatch(func() {
		hx_try_value_57 = hxrt.JsonParse(func() *string {
			var hx_if_60 *string
			if hxrt.StringEqualStringPtr(body, hxrt.StringFromLiteral("")) {
				hx_if_60 = hxrt.StringFromLiteral("{}")
			} else {
				hx_if_60 = body
			}
			return hx_if_60
		}())
		hx_try_return_56 = true
		return
	}, func(hx_caught_58 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_58)
		_ = hx_tmp
		hxrt.Throw(New_app__core__IncidentRequestException(hxrt.StringFromLiteral("invalid_json")))
	})
	if hx_try_return_56 {
		return hx_try_value_57
	}
	return nil
}
