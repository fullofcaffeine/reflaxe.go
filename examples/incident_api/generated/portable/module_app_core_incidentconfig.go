package main

import "examples_incident_api_portable/hxrt"

type I_app__core__IncidentConfig interface {
}

type app__core__IncidentConfig struct {
	__hx_this   I_app__core__IncidentConfig
	serviceName *string
	host        *string
	port        int
	statePath   *string
}

func New_app__core__IncidentConfig(serviceName *string, host *string, port int, statePath *string) *app__core__IncidentConfig {
	self := &app__core__IncidentConfig{}
	self.__hx_this = self
	self.serviceName = serviceName
	self.host = host
	self.port = port
	self.statePath = statePath
	return self
}

func app__core__IncidentConfig_defaults() *app__core__IncidentConfig {
	return New_app__core__IncidentConfig(hxrt.StringFromLiteral("incident-api"), hxrt.StringFromLiteral("127.0.0.1"), 0, hxrt.StringFromLiteral(".incident_api_state.json"))
}

func app__core__IncidentConfig_intField(raw any, name *string, fallback int) int {
	if hxrt.AnyEqualsNull(raw) || !Reflect_hasField(raw, name) {
		return fallback
	}
	var value any = Reflect_field(raw, name)
	if hxrt.AnyEqualsNull(value) {
		return fallback
	}
	var parsed any = hxrt.StdParseInt(hxrt.StdString(value))
	var hx_if_37 int
	if parsed == nil {
		hx_if_37 = fallback
	} else {
		hx_if_37 = parsed.(int)
	}
	return hx_if_37
}

func app__core__IncidentConfig_load(path *string) *app__core__IncidentConfig {
	config := app__core__IncidentConfig_defaults()
	if !sys__FileSystem_exists(path) {
		return config
	}
	text := sys__io__File_getContent(path)
	var raw any = hxrt.JsonParse(text)
	config.serviceName = app__core__IncidentConfig_stringField(raw, hxrt.StringFromLiteral("serviceName"), config.serviceName)
	config.host = app__core__IncidentConfig_stringField(raw, hxrt.StringFromLiteral("host"), config.host)
	config.port = app__core__IncidentConfig_intField(raw, hxrt.StringFromLiteral("port"), config.port)
	config.statePath = app__core__IncidentConfig_stringField(raw, hxrt.StringFromLiteral("statePath"), config.statePath)
	return config
}

func app__core__IncidentConfig_saveExample(path *string) {
	sys__io__File_saveContent(path, hxrt.StringFromLiteral("{\"serviceName\":\"incident-api\",\"host\":\"127.0.0.1\",\"port\":0,\"statePath\":\".incident_api_state.json\"}\n"))
}

func app__core__IncidentConfig_stringField(raw any, name *string, fallback *string) *string {
	if hxrt.AnyEqualsNull(raw) || !Reflect_hasField(raw, name) {
		return fallback
	}
	var value any = Reflect_field(raw, name)
	if hxrt.AnyEqualsNull(value) {
		return fallback
	}
	return hxrt.StdString(value)
}
