package main

import "examples_incident_api_portable/hxrt"

type I_app__core__IncidentStore interface {
	create(title *string, severity *string) *app__core__Incident
	acknowledge(id int) *app__core__Incident
	resolve(id int) *app__core__Incident
	listJson() *string
	metricsJson(serviceName *string, requests int) *string
	find(id int) *app__core__Incident
	load()
	save()
}

type app__core__IncidentStore struct {
	__hx_this I_app__core__IncidentStore
	statePath *string
	incidents *hxrt.Array
	nextId    int
}

func New_app__core__IncidentStore(statePath *string) *app__core__IncidentStore {
	self := &app__core__IncidentStore{}
	self.__hx_this = self
	self.statePath = statePath
	self.incidents = hxrt.NewArray()
	self.nextId = 1
	self.load()
	return self
}

func (self *app__core__IncidentStore) create(title *string, severity *string) *app__core__Incident {
	incident := New_app__core__Incident(self.nextId, title, app__core__IncidentStore_normalizeSeverity(severity), false, false, hxrt.StringFromLiteral("2026-06-12T00:00:00Z"))
	self.nextId = int(int32((self.nextId + 1)))
	hx_arr_62 := self.incidents
	hx_arr_62.Push(incident)
	self.save()
	return incident
}

func (self *app__core__IncidentStore) acknowledge(id int) *app__core__Incident {
	incident := self.find(id)
	if incident == nil {
		return nil
	}
	incident.acknowledged = true
	self.save()
	return incident
}

func (self *app__core__IncidentStore) resolve(id int) *app__core__Incident {
	incident := self.find(id)
	if incident == nil {
		return nil
	}
	incident.acknowledged = true
	incident.resolved = true
	self.save()
	return incident
}

func (self *app__core__IncidentStore) listJson() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	i := 0
	for i < self.incidents.Len() {
		if i > 0 {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(","))
		}
		x := func(hx_value_63 any) *app__core__Incident {
			if hx_value_63 == nil {
				var hx_zero_64 *app__core__Incident
				return hx_zero_64
			}
			return hx_value_63.(*app__core__Incident)
		}(self.incidents.Get(i)).toJson()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		i = int(int32((i + 1)))
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("]"))
	return out_b
}

func (self *app__core__IncidentStore) metricsJson(serviceName *string, requests int) *string {
	open := 0
	acked := 0
	resolved := 0
	_g := 0
	_g1 := self.incidents
	for _g < _g1.Len() {
		incident := func(hx_value_65 any) *app__core__Incident {
			if hx_value_65 == nil {
				var hx_zero_66 *app__core__Incident
				return hx_zero_66
			}
			return hx_value_65.(*app__core__Incident)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if incident.resolved {
			resolved = int(int32((resolved + 1)))
		} else {
			open = int(int32((open + 1)))
		}
		if incident.acknowledged {
			acked = int(int32((acked + 1)))
		}
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{\"service\":\""), app__core__Incident_jsonEscape(serviceName)), hxrt.StringFromLiteral("\",\"requests\":")), requests), hxrt.StringFromLiteral(",\"open\":")), open), hxrt.StringFromLiteral(",\"acknowledged\":")), acked), hxrt.StringFromLiteral(",\"resolved\":")), resolved), hxrt.StringFromLiteral("}"))
}

func (self *app__core__IncidentStore) find(id int) *app__core__Incident {
	_g := 0
	_g1 := self.incidents
	for _g < _g1.Len() {
		incident := func(hx_value_67 any) *app__core__Incident {
			if hx_value_67 == nil {
				var hx_zero_68 *app__core__Incident
				return hx_zero_68
			}
			return hx_value_67.(*app__core__Incident)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if incident.id == id {
			return incident
		}
	}
	return nil
}

func (self *app__core__IncidentStore) load() {
	if !sys__FileSystem_exists(self.statePath) {
		return
	}
	content := StringTools_trim(sys__io__File_getContent(self.statePath))
	if hxrt.StringEqualStringPtr(content, hxrt.StringFromLiteral("")) {
		return
	}
	var raw any = hxrt.JsonParse(content)
	self.nextId = app__core__IncidentStore_intField(raw, hxrt.StringFromLiteral("nextId"), 1)
	loaded := hxrt.NewArray()
	if !hxrt.AnyEqualsNull(raw) && Reflect_hasField(raw, hxrt.StringFromLiteral("incidents")) {
		values := func(hx_value_69 any) *hxrt.Array {
			if hx_value_69 == nil {
				var hx_zero_70 *hxrt.Array
				return hx_zero_70
			}
			return hx_value_69.(*hxrt.Array)
		}(Reflect_field(raw, hxrt.StringFromLiteral("incidents")))
		_g := 0
		for _g < values.Len() {
			var value any = values.Get(_g)
			_g = int(int32((_g + 1)))
			loaded.Push(New_app__core__Incident(app__core__IncidentStore_intField(value, hxrt.StringFromLiteral("id"), self.nextId), app__core__IncidentStore_stringField(value, hxrt.StringFromLiteral("title"), hxrt.StringFromLiteral("untitled")), app__core__IncidentStore_normalizeSeverity(app__core__IncidentStore_stringField(value, hxrt.StringFromLiteral("severity"), hxrt.StringFromLiteral("low"))), app__core__IncidentStore_boolField(value, hxrt.StringFromLiteral("acknowledged"), false), app__core__IncidentStore_boolField(value, hxrt.StringFromLiteral("resolved"), false), app__core__IncidentStore_stringField(value, hxrt.StringFromLiteral("createdAt"), hxrt.StringFromLiteral("2026-06-12T00:00:00Z"))))
		}
	}
	self.incidents = loaded
}

func (self *app__core__IncidentStore) save() {
	sys__io__File_saveContent(self.statePath, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("{\"nextId\":"), self.nextId), hxrt.StringFromLiteral(",\"incidents\":")), self.listJson()), hxrt.StringFromLiteral("}\n")))
}

func app__core__IncidentStore_boolField(raw any, name *string, fallback bool) bool {
	if hxrt.AnyEqualsNull(raw) || !Reflect_hasField(raw, name) {
		return fallback
	}
	return hxrt.StringEqualStringPtr(hxrt.StdString(Reflect_field(raw, name)), hxrt.StringFromLiteral("true"))
}

func app__core__IncidentStore_intField(raw any, name *string, fallback int) int {
	if hxrt.AnyEqualsNull(raw) || !Reflect_hasField(raw, name) {
		return fallback
	}
	var parsed any = Std_parseInt(hxrt.StdString(Reflect_field(raw, name)))
	var hx_if_72 int
	if parsed == nil {
		hx_if_72 = fallback
	} else {
		hx_if_72 = parsed.(int)
	}
	return hx_if_72
}

func app__core__IncidentStore_normalizeSeverity(raw *string) *string {
	value := hxrt.StringToLowerCaseStringPtr(raw)
	if ((hxrt.StringEqualStringPtr(value, hxrt.StringFromLiteral("critical")) || hxrt.StringEqualStringPtr(value, hxrt.StringFromLiteral("high"))) || hxrt.StringEqualStringPtr(value, hxrt.StringFromLiteral("medium"))) || hxrt.StringEqualStringPtr(value, hxrt.StringFromLiteral("low")) {
		return value
	}
	return hxrt.StringFromLiteral("low")
}

func app__core__IncidentStore_stringField(raw any, name *string, fallback *string) *string {
	if hxrt.AnyEqualsNull(raw) || !Reflect_hasField(raw, name) {
		return fallback
	}
	var value any = Reflect_field(raw, name)
	var hx_if_73 *string
	if hxrt.AnyEqualsNull(value) {
		hx_if_73 = fallback
	} else {
		hx_if_73 = hxrt.StdString(value)
	}
	return hx_if_73
}
