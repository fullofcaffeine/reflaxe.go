package hxrt

import "testing"

type serializationProbe struct {
	__hx_this *serializationProbe
	z         int
	a         *string
}

func TestSerializationFieldsIncludesPrivateFieldsInDeclarationOrder(t *testing.T) {
	probe := &serializationProbe{z: 7, a: StringFromLiteral("ok")}
	fields := SerializationFields(probe)
	if len(fields) != 2 {
		t.Fatalf("fields length = %d, want 2", len(fields))
	}
	if *fields[0].Name != "z" || *fields[1].Name != "a" {
		t.Fatalf("field order = [%s %s], want [z a]", *fields[0].Name, *fields[1].Name)
	}
	if fields[0].Value != 7 {
		t.Fatalf("z value = %#v, want 7", fields[0].Value)
	}
}

func TestSerializationSetFieldAndBindSelfReachPrivateFields(t *testing.T) {
	probe := &serializationProbe{}
	SerializationSetField(probe, StringFromLiteral("z"), 9)
	SerializationBindSelf(probe)
	if probe.z != 9 {
		t.Fatalf("z = %d, want 9", probe.z)
	}
	if probe.__hx_this != probe {
		t.Fatal("hidden self pointer was not rebound")
	}
}
