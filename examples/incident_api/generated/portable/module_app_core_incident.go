package main

import "examples_incident_api_portable/hxrt"

type I_app__core__Incident interface {
	toJson() *string
}

type app__core__Incident struct {
	__hx_this    I_app__core__Incident
	id           int
	title        *string
	severity     *string
	acknowledged bool
	resolved     bool
	createdAt    *string
}

func New_app__core__Incident(id int, title *string, severity *string, acknowledged bool, resolved bool, createdAt *string) *app__core__Incident {
	self := &app__core__Incident{}
	self.__hx_this = self
	self.id = id
	self.title = title
	self.severity = severity
	self.acknowledged = acknowledged
	self.resolved = resolved
	self.createdAt = createdAt
	return self
}

func (self *app__core__Incident) toJson() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("{\"id\":"), self.id), hxrt.StringFromLiteral(",\"title\":\"")), app__core__Incident_jsonEscape(self.title)), hxrt.StringFromLiteral("\"")), hxrt.StringFromLiteral(",\"severity\":\"")), app__core__Incident_jsonEscape(self.severity)), hxrt.StringFromLiteral("\"")), hxrt.StringFromLiteral(",\"acknowledged\":")), app__core__Incident_boolJson(self.acknowledged)), hxrt.StringFromLiteral(",\"resolved\":")), app__core__Incident_boolJson(self.resolved)), hxrt.StringFromLiteral(",\"createdAt\":\"")), app__core__Incident_jsonEscape(self.createdAt)), hxrt.StringFromLiteral("\"}"))
}

func app__core__Incident_boolJson(value bool) *string {
	var hx_if_46 *string
	if value {
		hx_if_46 = hxrt.StringFromLiteral("true")
	} else {
		hx_if_46 = hxrt.StringFromLiteral("false")
	}
	return hx_if_46
}

func app__core__Incident_jsonEscape(value *string) *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	i := 0
	for i < hxrt.StringLengthStringPtr(value) {
		var code any = hxrt.StringCharCodeAtAnyStringPtr(value, i)
		if code == 34 {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("\\\""))
		} else {
			if code == 92 {
				out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("\\\\"))
			} else {
				if code == 10 {
					out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("\\n"))
				} else {
					if code == 13 {
						out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("\\r"))
					} else {
						if code == 9 {
							out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("\\t"))
						} else {
							c := hxrt.IntFromNullableAny(code)
							out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromCharCode(c))
						}
					}
				}
			}
		}
		i = int(int32((i + 1)))
	}
	return out_b
}
