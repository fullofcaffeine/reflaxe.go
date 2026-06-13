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
		var raw any = hxrt.JsonParse(func() *string {
			var hx_if_30 *string
			if hxrt.StringEqualStringPtr(body, hxrt.StringFromLiteral("")) {
				hx_if_30 = hxrt.StringFromLiteral("{}")
			} else {
				hx_if_30 = body
			}
			return hx_if_30
		}())
		title := app__core__IncidentApi_fieldString(raw, hxrt.StringFromLiteral("title"), hxrt.StringFromLiteral(""))
		if hxrt.StringEqualStringPtr(title, hxrt.StringFromLiteral("")) {
			response = app__http__HttpResponse_json(400, hxrt.StringFromLiteral("{\"error\":\"missing_title\"}"))
		} else {
			severity := app__core__IncidentApi_fieldString(raw, hxrt.StringFromLiteral("severity"), hxrt.StringFromLiteral("low"))
			incident := self.store.create(title, severity)
			response = app__http__HttpResponse_json(201, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{\"incident\":"), incident.toJson()), hxrt.StringFromLiteral("}")))
		}
	}, func(hx_caught_28 any) {
		hx_tmp := hx_caught_28
		_ = hx_tmp
		response = app__http__HttpResponse_json(400, hxrt.StringFromLiteral("{\"error\":\"invalid_json\"}"))
	})
	return response
}

func (self *app__core__IncidentApi) updateIncident(path *string, action *string) *app__http__HttpResponse {
	parts := hxrt.StringSplitStringPtr(path, hxrt.StringFromLiteral("/"))
	if len(parts) < 4 {
		return app__http__HttpResponse_json(404, hxrt.StringFromLiteral("{\"error\":\"not_found\"}"))
	}
	var id any = hxrt.StdParseInt(parts[2])
	if id == nil {
		return app__http__HttpResponse_json(400, hxrt.StringFromLiteral("{\"error\":\"invalid_id\"}"))
	}
	var hx_if_31 *app__core__Incident
	if hxrt.StringEqualStringPtr(action, hxrt.StringFromLiteral("ack")) {
		hx_if_31 = self.store.acknowledge(id.(int))
	} else {
		hx_if_31 = self.store.resolve(id.(int))
	}
	incident := hx_if_31
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
	var hx_if_32 *string
	if hxrt.AnyEqualsNull(value) {
		hx_if_32 = fallback
	} else {
		hx_if_32 = hxrt.StdString(value)
	}
	return hx_if_32
}
