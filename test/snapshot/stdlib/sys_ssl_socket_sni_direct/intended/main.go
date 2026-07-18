package main

import (
	"reflect"
	"snapshot/hxrt"
)

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

var DEFAULT_CERT_PEM *string = hxrt.StringFromLiteral("-----BEGIN CERTIFICATE-----\nMIIDETCCAfmgAwIBAgIUVftG+IyWIEqIQE+K9ztyXCDYIcswDQYJKoZIhvcNAQEL\nBQAwGDEWMBQGA1UEAwwNZGVmYXVsdC5sb2NhbDAeFw0yNjAzMDcwNjA1NTBaFw0y\nNzAzMDcwNjA1NTBaMBgxFjAUBgNVBAMMDWRlZmF1bHQubG9jYWwwggEiMA0GCSqG\nSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCktatoj06k/dqldeSzjPUnPxeMC/WpprMz\n7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqaQDkLF0Zy+Q2K+GWLRCux\n7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVKT3jjcc3ivv6j/kOsbE7j\njELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLeoT+UU7O9VKsEIQHCGt5P\np7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubpnx+I/wUmLK4Aqh2VuuZz\nmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi9/IqJ3NpAgMBAAGjUzBR\nMB0GA1UdDgQWBBR7eOzK13X2mO5Kpc9MZOFGJ+MdlDAfBgNVHSMEGDAWgBR7eOzK\n13X2mO5Kpc9MZOFGJ+MdlDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUA\nA4IBAQAxilSFs9ZEEI5RFGVpWUTCH8Iuewc4K0JH2LbGqExgLu5MQJF5xsGoEBRe\nmDB3duaMDwnoDY19hdoClIz/Z5IO0wPEcny3hTb582W8+cRDiCQx0Qz5g2NpsfGH\nwZaGyVzSS2xbt0F6TPRQarXtzV97J067j7bRuMbD4fFYb6iqF1GbaaQsP89sA9b4\nx8yT0TGREdx2Dw29lokd4H16b3O9aYX274qZk9qpgE3oYpCeNsjHCSdTWRDw2+pg\ncKP2FsxcbB4qChjPXhhE5I3zDMTI1/4r4wGbnVIud6TmibElpAU6hwbpo1onyxKw\nVCpI190WWaS+Cz+9vGVNBSsRYjgN\n-----END CERTIFICATE-----\n")

var DEFAULT_KEY_PEM *string = hxrt.StringFromLiteral("-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCktatoj06k/dql\ndeSzjPUnPxeMC/WpprMz7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqa\nQDkLF0Zy+Q2K+GWLRCux7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVK\nT3jjcc3ivv6j/kOsbE7jjELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLe\noT+UU7O9VKsEIQHCGt5Pp7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubp\nnx+I/wUmLK4Aqh2VuuZzmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi\n9/IqJ3NpAgMBAAECggEAECDG2jvwGjlOO9XtsUQr7C4iHuE76qMLWJo5vI5GfteQ\ndZUHVuA1Fh+NC+3z2N+uHIukuWz9G+wWGuEPhN3AVPk8oX89qDOiaK90XV6Cwt3s\n0ZU6gW/nz+qHmvdXru60nCrLephnEr+cP0TFZFYMMDf+BK5cz4kid2cQUmItbKBy\nE9k1vHvyyLONK18y72IXsMEA572V+AhtvWVar5qUJiY2+dWtxmWh9dzUvWCUJ/Gj\niqyrJdNxHJfXdteT8QpJEndsX19MWNAMZY+e10XOVzs2Rd4RdnlLwUE37HbPz73V\nC1z2h6tshSOYi9ZwlNdYEjfWbuiQimwgljwc9/NagQKBgQDg4JkcchEOmntEXPz0\n/HTEmAd7Sl75V0ozj7xPTYcaSvdoEtLEbUJbGoVfLKJM17jLuEjpisS1IHa29kkG\n2PACbbIRzOLlCA2Spkzm9R2Rni9qbBCwYbnYbdhV+hshcnr6LwUt8KBAFHBCk2lW\nMfiCsBE1KzBeye1JFxIZ9+dyKQKBgQC7gVSvPDIqPRYP8l8sLbONMHTNhbZpMh23\npdD1i0YpYTIuvrxZMXwY6gt71SJcHIwyh/6HEg8mAIMZutdRzpZc4NWQw98yG9Zo\nkpY7S2uLZ8YnARps+NprdWiARkY3PJjIiK7tieNo+dQtfBeuL3Ox7s/XR0nDUXMP\nxkuAnByfQQKBgQCks9twehsEFyExcOnUhRMA6liQdGgbN1OhcCT78EyDdWS/VQoJ\n0/xFvabxjj9RCK7QhqjgZEKuZpiMaNYTrdAb9zv0zZthJATM5ABvKBgAD1urFnsi\ntHDpk4pfbk9wr+hiVQ32F8dHJ7EREeaUuwTIsyvnRTqoMj0Yy0z2uBtMAQKBgQC2\nEnG+7z7vEP4ZYgrUhVQyp3jkERD9uTJuH892f1UT3VOzXHbcTVbpgmrARkflFbt1\nXeTkF78p8ZlcJLfssiQD8DaxKeHTcICUbrL+xM+bQJuDSGj2o/bEHe/pj1OjU24w\nW7kw45I1X1KPEE6WT3GSuAiOTKTtymtmR/EM44pPgQKBgQC52gWDbpJb1UREu9Sb\naIy1zgqGhSxg1DLChu4vz3t5ZBhWoiw6wG7vQDrtrn9Jl61mVtog0rSTSiIj726W\nLDxSrTdYBhcgkbJf/fw1N6Uk+FfJcVS/zuoJwogkiP6tmircL0qA9sYR4RZo2ubE\nAg6A+vHIw54BclYKSczsXnCe+A==\n-----END PRIVATE KEY-----\n")

var SNI_CERT_PEM *string = hxrt.StringFromLiteral("-----BEGIN CERTIFICATE-----\nMIIDHzCCAgegAwIBAgIUMCFv2IwHqj9Wy6ycvAibkSutMKQwDQYJKoZIhvcNAQEL\nBQAwFDESMBAGA1UEAwwJc25pLmxvY2FsMB4XDTI2MDYxMDA1MTUwMloXDTI3MDYx\nMDA1MTUwMlowFDESMBAGA1UEAwwJc25pLmxvY2FsMIIBIjANBgkqhkiG9w0BAQEF\nAAOCAQ8AMIIBCgKCAQEAsQVlch80ZQ4hg3Hb3XMPgk+XPsJuoIdE0rn78PGM/O6W\n1iQocq9hYRosp/aT6itBnpvxh+tCt9D7deoEDBzwU6CkJu7N6CGY0yWWrFkits4f\n8KiTbkWZA3qKfreUOiDkiU3DgM+b9GUyBkgUXxBxW1ql+mFAX6l3eRDK5q9dxxme\nmouBL4yQxMa1bgunsYoZQhV/0RoH0gUuDUTPr12+YEA/XyCJC82p3v50mIcURjnX\nHl26xMRv5O73mJDUqshn3Z0SMmXhHfk6JFsacc2sB7f8FCOAAFsjjhSKHEQ8OTTP\nzcXAw6fcqxF1WNMrsIfZJ4Dzjt2kOYO9KxiLR3+0dwIDAQABo2kwZzAdBgNVHQ4E\nFgQUnbhLMtV250tds1NaZBDZyIvDXfwwHwYDVR0jBBgwFoAUnbhLMtV250tds1Na\nZBDZyIvDXfwwDwYDVR0TAQH/BAUwAwEB/zAUBgNVHREEDTALgglzbmkubG9jYWww\nDQYJKoZIhvcNAQELBQADggEBADHJ3F34+3iZrJNgxuFbr5IOsm4GDCZxd8g0BmWD\n+JxabwGYMk+xGOL/dn7jPwhodHcFbJXrYzT2giGMG9EdLLExqwnpQb7Z+uqga4Ej\nfyQVZZ61Np8wSGJQ7C03Yu6A0jYljW8KWM3UeiX6ee/d8/JXK8Nts8qsKXe6SFJO\nF1vblGLBZyb/DZ4GKD5cp3RZdCQxwlvLmbYKrKDwySJYwWQEx8bD5UUMYPr0Y593\nn6ZbEIWl+17PkP0kQRloeGUOxogZaI9jM/fGalKf52SbipDQV4P8NXqsdW0Im/LA\nCW6HU+6ozncjLj3SyLuwLvd54+YEPhv/cuFgP6/WWqyCx0A=\n-----END CERTIFICATE-----\n")

var SNI_KEY_PEM *string = hxrt.StringFromLiteral("-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCxBWVyHzRlDiGD\ncdvdcw+CT5c+wm6gh0TSufvw8Yz87pbWJChyr2FhGiyn9pPqK0Gem/GH60K30Pt1\n6gQMHPBToKQm7s3oIZjTJZasWSK2zh/wqJNuRZkDeop+t5Q6IOSJTcOAz5v0ZTIG\nSBRfEHFbWqX6YUBfqXd5EMrmr13HGZ6ai4EvjJDExrVuC6exihlCFX/RGgfSBS4N\nRM+vXb5gQD9fIIkLzane/nSYhxRGOdceXbrExG/k7veYkNSqyGfdnRIyZeEd+Tok\nWxpxzawHt/wUI4AAWyOOFIocRDw5NM/NxcDDp9yrEXVY0yuwh9kngPOO3aQ5g70r\nGItHf7R3AgMBAAECggEANjMyxmhrgG19MWPRL9Kk8v8vjdW2TYxdNDAhxboPsvnS\nUSqs/8BXDoYXGi5TR6WK5+dTYoxT1zgzZf0K1DKgGtrap9kCTorK4gtmQMrh6Brg\niKz0xxSkLv58HSRUTB/6GVgn/e6TD5dUY7v6EMlWC+SLYUgZj7Cxle3gUhVrnyPk\noaEOMDU8w9iiPoCK0E60DReKAmIcXigGqMNlkFjADjfOtr5GDpercL+mvwDmfLmB\ndzYcXuHM3fm1pgoDXN5zrpO7mFSp0VG/O697nZzY18A25AMMNiBUjuk7/u4KJEcA\nMmOxVAU6yjHAKO+YRS98pKsxkvZKNpj69lXtBv6agQKBgQDVnTDnIudJwbaAvfBf\n7ZZ3l46Ojm2s3qDR0Sp+W0tTc5cVhJiJaShzpnYKPbZP6R8EnDrdtldw8sYRRB6R\nPfM3/DX87z793+pKB5WDJNzevO81ecpMhiYgUxlKzzkz7nfDZC8tBQ3YLLhUX5JL\nv7REeUskkmW+xq9e9+ecK/RpVwKBgQDUJWrvJ2tEf2JJ+qdYZBHJXw5+ajIZeUs+\nCh3cNtOrUGDW/y9SKvrRHzUP0PvRle21O4oUimEb1SqdNG2nFB3f4BVUB8X6DcMz\nM16kyNBdk4q2k3xITw2OdErlYZU2WAsLUrOxurDK9ng6YO1fJwz3zrGiNDjbdebY\n1m4NbyN54QKBgQCE0akvfh9LV/wPDoqgSszs7TpBX0PIYeCitShzynYKnGuLgJeL\nkOwLBKyOb5KlGzEjH7TmWFMEMp9+6tkKu/c3j0VOUL/dANXfU9nd4hTHFbiyiliD\nvkGEhcbLIg/SP2sN/YPrvSG/kQbHx2jiWn9OuBBF3BURSt6N8Rx8mUPuHwKBgQCG\nAqj9L8Jz/5/gKaVCkdwmf5SRSJYjP1rHcu6P6FZntpul1IdY+Wt9ZKBJQHOCXppN\nTLIZ7ZwQT+Teb3sA+xUwEcaHUW2/Wqg/FKkpoOz237fVQ29T4hQnM9EH+0+dh5pa\nacC3eb4qR+2EuyvXWry3YWsWkrSD9YOA4FuewuD/IQKBgDUet4ixooKltLLhW6RT\nB+8+u2hKmlctueypoiB9+yiCt/oEiZEHYX9IwlGaax6bpDg6vO77Ay7JwGaggM+K\nhhIsqAuszWLO7bdjdS4JZ1YshP5YusdPs3m9OGeTjnO/+goL+tECrZmzupwg4BTA\nKOG6OBx2h7HDGnJcWJNLuBbJ\n-----END PRIVATE KEY-----\n")

func connectAndPrint(server *sys__ssl__Socket, port int, hostname *string) {
	acceptedSockets := New_sys__thread__Deque()
	sys__thread__Thread_create(func() {
		acceptedSockets.__hx_this.add(server.__hx_this.accept())
	})
	client := New_sys__ssl__Socket()
	client.verifyCert = false
	client.__hx_this.setHostname(hostname)
	client.__hx_this.connect(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), port)
	client.__hx_this.handshake()
	peer := client.__hx_this.peerCertificate()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hostname, hxrt.StringFromLiteral("=")), peer.__hx_this.get_commonName()))
	hxrt.Println(v)
	func(hx_value_1 any) *sys__net__Socket {
		if hx_value_1 == nil {
			var hx_zero_2 *sys__net__Socket
			return hx_zero_2
		}
		return hx_value_1.(*sys__net__Socket)
	}(acceptedSockets.__hx_this.pop(true)).__hx_this.close()
	client.__hx_this.close()
}

func main() {
	defer hxrt.ThreadWaitForAll()
	defaultCert := sys__ssl__Certificate_fromString(DEFAULT_CERT_PEM)
	defaultKey := sys__ssl__Key_readPEM(DEFAULT_KEY_PEM, false, nil)
	sniCert := sys__ssl__Certificate_fromString(SNI_CERT_PEM)
	sniKey := sys__ssl__Key_readPEM(SNI_KEY_PEM, false, nil)
	server := New_sys__ssl__Socket()
	server.__hx_this.setCertificate(defaultCert, defaultKey)
	server.__hx_this.addSNICertificate(func(name *string) bool {
		return hxrt.StringEqualStringPtr(name, hxrt.StringFromLiteral("sni.local"))
	}, sniCert, sniKey)
	server.__hx_this.bind(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), 0)
	server.__hx_this.listen(2)
	bound := server.__hx_this.host()
	connectAndPrint(server, func(hx_obj_3 map[string]any) int {
		hx_field_4 := hx_obj_3["port"]
		if hx_field_4 == nil {
			var hx_zero_5 int
			return hx_zero_5
		}
		return hx_field_4.(int)
	}(bound), hxrt.StringFromLiteral("default.local"))
	connectAndPrint(server, func(hx_obj_6 map[string]any) int {
		hx_field_7 := hx_obj_6["port"]
		if hx_field_7 == nil {
			var hx_zero_8 int
			return hx_zero_8
		}
		return hx_field_7.(int)
	}(bound), hxrt.StringFromLiteral("sni.local"))
	server.__hx_this.close()
}

type Std struct {
}

type Type struct {
}

type Reflect struct {
}

func Reflect_compare(a any, b any) int {
	toFloat := func(value any) (float64, bool) {
		switch v := value.(type) {
		case int:
			return float64(v), true
		case int8:
			return float64(v), true
		case int16:
			return float64(v), true
		case int32:
			return float64(v), true
		case int64:
			return float64(v), true
		case uint:
			return float64(v), true
		case uint8:
			return float64(v), true
		case uint16:
			return float64(v), true
		case uint32:
			return float64(v), true
		case uint64:
			return float64(v), true
		case float32:
			return float64(v), true
		case float64:
			return v, true
		default:
			return 0, false
		}
	}
	if af, ok := toFloat(a); ok {
		if bf, okB := toFloat(b); okB {
			if af < bf {
				return -1
			}
			if af > bf {
				return 1
			}
			return 0
		}
	}
	aStr := *hxrt.StdString(a)
	bStr := *hxrt.StdString(b)
	if aStr < bStr {
		return -1
	}
	if aStr > bStr {
		return 1
	}
	return 0
}

func Reflect_compareMethods(a any, b any) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	if !av.IsValid() || !bv.IsValid() {
		return !av.IsValid() && !bv.IsValid()
	}
	if av.Kind() == reflect.Func && bv.Kind() == reflect.Func {
		if av.IsNil() || bv.IsNil() {
			return av.IsNil() && bv.IsNil()
		}
		return av.Pointer() == bv.Pointer()
	}
	return reflect.DeepEqual(a, b)
}

func Reflect_field(obj any, field *string) any {
	if obj == nil {
		return nil
	}
	key := *hxrt.StdString(field)
	if metadataValue, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return metadataValue
	}
	switch value := obj.(type) {
	case map[string]any:
		return value[key]
	case map[any]any:
		return value[key]
	case *map[string]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	case *map[any]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return nil
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if fieldValue := rv.FieldByName(key); fieldValue.IsValid() && fieldValue.CanInterface() {
			return fieldValue.Interface()
		}
	}
	method := reflect.ValueOf(obj).MethodByName(key)
	if method.IsValid() {
		return method.Interface()
	}
	return nil
}

func Reflect_hasField(obj any, field *string) bool {
	if obj == nil {
		return false
	}
	key := *hxrt.StdString(field)
	if _, ok := hxrt_typeClassMetadataField(obj, key); ok {
		return true
	}
	switch value := obj.(type) {
	case map[string]any:
		_, ok := value[key]
		return ok
	case map[any]any:
		_, ok := value[key]
		return ok
	case *map[string]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	case *map[any]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return false
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if rv.FieldByName(key).IsValid() {
			return true
		}
	}
	return reflect.ValueOf(obj).MethodByName(key).IsValid()
}

func Reflect_setField(obj any, field *string, value any) {
	if obj == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	key := *hxrt.StdString(field)
	switch target := obj.(type) {
	case map[string]any:
		target[key] = value
		return
	case map[any]any:
		target[key] = value
		return
	case *map[string]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	case *map[any]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() || rv.Kind() != reflect.Pointer {
		return
	}
	if rv.IsNil() {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return
	}
	fieldValue := rv.FieldByName(key)
	if !fieldValue.IsValid() || !fieldValue.CanSet() {
		return
	}
	if value == nil {
		fieldValue.Set(reflect.Zero(fieldValue.Type()))
		return
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(fieldValue.Type()) {
		fieldValue.Set(incoming)
		return
	}
	if incoming.Type().ConvertibleTo(fieldValue.Type()) {
		fieldValue.Set(incoming.Convert(fieldValue.Type()))
		return
	}
	if fieldValue.Kind() == reflect.Interface {
		fieldValue.Set(incoming)
	}
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}

type ValueType struct {
	tag    int
	params []any
}

var ValueType_TNull *ValueType = &ValueType{tag: 0, params: []any{}}

var ValueType_TInt *ValueType = &ValueType{tag: 1, params: []any{}}

var ValueType_TFloat *ValueType = &ValueType{tag: 2, params: []any{}}

var ValueType_TBool *ValueType = &ValueType{tag: 3, params: []any{}}

var ValueType_TObject *ValueType = &ValueType{tag: 4, params: []any{}}

var ValueType_TFunction *ValueType = &ValueType{tag: 5, params: []any{}}

var ValueType_TUnknown *ValueType = &ValueType{tag: 8, params: []any{}}

func ValueType_TClass(c any) *ValueType {
	return &ValueType{tag: 6, params: []any{c}}
}

func ValueType_TEnum(e any) *ValueType {
	return &ValueType{tag: 7, params: []any{e}}
}

func hxrt_typeCallAny(callable any, args []any) (any, bool) {
	result := any(nil)
	ok := false
	defer func() {
		if recover() != nil {
			result = nil
			ok = false
		}
	}()
	if callable == nil {
		return nil, false
	}
	fn := reflect.ValueOf(callable)
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil, false
	}
	fnType := fn.Type()
	if fnType.NumIn() != len(args) {
		return nil, false
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		paramType := fnType.In(i)
		arg := args[i]
		if arg == nil {
			in[i] = reflect.Zero(paramType)
			continue
		}
		v := reflect.ValueOf(arg)
		if v.IsValid() && v.Type().AssignableTo(paramType) {
			in[i] = v
			continue
		}
		if v.IsValid() && v.Type().ConvertibleTo(paramType) {
			in[i] = v.Convert(paramType)
			continue
		}
		if paramType.Kind() == reflect.Interface && v.IsValid() {
			in[i] = v
			continue
		}
		return nil, false
	}
	out := fn.Call(in)
	if len(out) == 0 {
		return nil, true
	}
	first := out[0]
	if !first.IsValid() {
		return nil, true
	}
	result = first.Interface()
	ok = true
	return result, ok
}

func hxrt_typeArrayValues(value *hxrt.Array) []any {
	if value == nil {
		return []any{}
	}
	return value.Values()
}

func hxrt_typeResolvedClassName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeClassValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeClassValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeResolvedEnumName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeEnumValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeEnumValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeCreateClassInstance(className string, args []any) (any, bool) {
	switch className {
	case "Date":
		return hxrt_typeCallAny(New_Date, args)
	case "Main":
		return nil, false
	case "StringBuf":
		return nil, false
	case "haxe.Int64Helper":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.ds.List":
		return hxrt_typeCallAny(New_haxe__ds__List, args)
	case "haxe.ds._List.GoListIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListIterator, args)
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__ds___List__GoListKeyValueIterator, args)
	case "haxe.exceptions.NotImplementedException":
		return hxrt_typeCallAny(New_haxe__exceptions__NotImplementedException, args)
	case "haxe.exceptions.PosException":
		return hxrt_typeCallAny(New_haxe__exceptions__PosException, args)
	case "haxe.io.Bytes":
		return hxrt_typeCallAny(New_haxe__io__Bytes, args)
	case "haxe.io.BytesBuffer":
		return hxrt_typeCallAny(New_haxe__io__BytesBuffer, args)
	case "haxe.io.Eof":
		return hxrt_typeCallAny(New_haxe__io__Eof, args)
	case "haxe.io.FPHelper":
		return nil, false
	case "haxe.io.Input":
		return hxrt_typeCallAny(New_haxe__io__Input, args)
	case "haxe.io.Output":
		return hxrt_typeCallAny(New_haxe__io__Output, args)
	case "sys.net.Host":
		return hxrt_typeCallAny(New_sys__net__Host, args)
	case "sys.net.Socket":
		return hxrt_typeCallAny(New_sys__net__Socket, args)
	case "sys.net.SocketInput":
		return hxrt_typeCallAny(New_sys__net__SocketInput, args)
	case "sys.net.SocketOutput":
		return hxrt_typeCallAny(New_sys__net__SocketOutput, args)
	case "sys.ssl.Certificate":
		return hxrt_typeCallAny(New_sys__ssl__Certificate, args)
	case "sys.ssl.Key":
		return hxrt_typeCallAny(New_sys__ssl__Key, args)
	case "sys.ssl.Socket":
		return hxrt_typeCallAny(New_sys__ssl__Socket, args)
	case "sys.thread.Deque":
		return hxrt_typeCallAny(New_sys__thread__Deque, args)
	case "sys.thread.EventLoop":
		return hxrt_typeCallAny(New_sys__thread__EventLoop, args)
	case "sys.thread.Lock":
		return hxrt_typeCallAny(New_sys__thread__Lock, args)
	case "sys.thread.Mutex":
		return hxrt_typeCallAny(New_sys__thread__Mutex, args)
	case "sys.thread.NoEventLoopException":
		return hxrt_typeCallAny(New_sys__thread__NoEventLoopException, args)
	case "sys.thread.Thread":
		return hxrt_typeCallAny(New_sys__thread__Thread, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "Date":
		return &Date{}, true
	case "haxe._Int64.___Int64":
		return &haxe___Int64_____Int64{}, true
	case "haxe.ds.List":
		return &haxe__ds__List{}, true
	case "haxe.ds._List.GoListIterator":
		return &haxe__ds___List__GoListIterator{}, true
	case "haxe.ds._List.GoListKeyValueIterator":
		return &haxe__ds___List__GoListKeyValueIterator{}, true
	case "haxe.exceptions.NotImplementedException":
		return &haxe__exceptions__NotImplementedException{}, true
	case "haxe.exceptions.PosException":
		return &haxe__exceptions__PosException{}, true
	case "haxe.io.Bytes":
		return &haxe__io__Bytes{}, true
	case "haxe.io.BytesBuffer":
		return &haxe__io__BytesBuffer{}, true
	case "haxe.io.Eof":
		return &haxe__io__Eof{}, true
	case "haxe.io.Input":
		return &haxe__io__Input{}, true
	case "haxe.io.Output":
		return &haxe__io__Output{}, true
	case "sys.net.Host":
		return &sys__net__Host{}, true
	case "sys.net.Socket":
		return &sys__net__Socket{}, true
	case "sys.net.SocketInput":
		return &sys__net__SocketInput{}, true
	case "sys.net.SocketOutput":
		return &sys__net__SocketOutput{}, true
	case "sys.ssl.Certificate":
		return &sys__ssl__Certificate{}, true
	case "sys.ssl.Key":
		return &sys__ssl__Key{}, true
	case "sys.ssl.Socket":
		return &sys__ssl__Socket{}, true
	case "sys.thread.Deque":
		return &sys__thread__Deque{}, true
	case "sys.thread.EventLoop":
		return &sys__thread__EventLoop{}, true
	case "sys.thread.Lock":
		return &sys__thread__Lock{}, true
	case "sys.thread.Mutex":
		return &sys__thread__Mutex{}, true
	case "sys.thread.NoEventLoopException":
		return &sys__thread__NoEventLoopException{}, true
	case "sys.thread.Thread":
		return &sys__thread__Thread{}, true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
	case "ValueType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TNull, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TInt, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFloat, true
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TBool, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TObject, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFunction, true
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TClass, args)
			case 7:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TEnum, args)
			case 8:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TUnknown, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TNull":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TNull, true
		case "TInt":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TInt, true
		case "TFloat":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFloat, true
		case "TBool":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TBool, true
		case "TObject":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TObject, true
		case "TFunction":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFunction, true
		case "TClass":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TClass, args)
		case "TEnum":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TEnum, args)
		case "TUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TUnknown, true
		default:
			return nil, false
		}
	case "haxe.io.Encoding":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Encoding_UTF8, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Encoding_RawNative, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "UTF8":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Encoding_UTF8, true
		case "RawNative":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Encoding_RawNative, true
		default:
			return nil, false
		}
	case "haxe.io.Error":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_Blocked, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_Overflow, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_OutsideBounds, true
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__io__Error_Custom, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "Blocked":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_Blocked, true
		case "Overflow":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_Overflow, true
		case "OutsideBounds":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_OutsideBounds, true
		case "Custom":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__io__Error_Custom, args)
		default:
			return nil, false
		}
	case "sys.thread.NextEventTime":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return sys__thread__NextEventTime_Now, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return sys__thread__NextEventTime_Never, true
			case 2:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(sys__thread__NextEventTime_AnyTime, args)
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(sys__thread__NextEventTime_At, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "Now":
			if len(args) != 0 {
				return nil, false
			}
			return sys__thread__NextEventTime_Now, true
		case "Never":
			if len(args) != 0 {
				return nil, false
			}
			return sys__thread__NextEventTime_Never, true
		case "AnyTime":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(sys__thread__NextEventTime_AnyTime, args)
		case "At":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(sys__thread__NextEventTime_At, args)
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func Type_getClass(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeClassValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeClassValue:
		copyValue := value
		return &copyValue
	case *hxrt.Array:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")}
	case *Date:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Date")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	case *haxe__ds__List:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.List")}
	case *haxe__ds___List__GoListIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._List.GoListIterator")}
	case *haxe__ds___List__GoListKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds._List.GoListKeyValueIterator")}
	case *haxe__exceptions__NotImplementedException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.NotImplementedException")}
	case *haxe__exceptions__PosException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.PosException")}
	case *haxe__io__Bytes:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Bytes")}
	case *haxe__io__BytesBuffer:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.BytesBuffer")}
	case *haxe__io__Eof:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Eof")}
	case *haxe__io__Input:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Input")}
	case *haxe__io__Output:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Output")}
	case *sys__net__Host:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.Host")}
	case *sys__net__Socket:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.Socket")}
	case *sys__net__SocketInput:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.SocketInput")}
	case *sys__net__SocketOutput:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.SocketOutput")}
	case *sys__ssl__Certificate:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.ssl.Certificate")}
	case *sys__ssl__Key:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.ssl.Key")}
	case *sys__ssl__Socket:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.ssl.Socket")}
	case *sys__thread__Deque:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Deque")}
	case *sys__thread__EventLoop:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.EventLoop")}
	case *sys__thread__Lock:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Lock")}
	case *sys__thread__Mutex:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Mutex")}
	case *sys__thread__NoEventLoopException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.NoEventLoopException")}
	case *sys__thread__Thread:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.thread.Thread")}
	default:
		return nil
	}
}

func Type_getEnum(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeEnumValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeEnumValue:
		copyValue := value
		return &copyValue
	case *ValueType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("ValueType")}
	case *haxe__io__Encoding:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.io.Encoding")}
	case *haxe__io__Error:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.io.Error")}
	case *sys__thread__NextEventTime:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("sys.thread.NextEventTime")}
	default:
		return nil
	}
}

func Type_getSuperClass(c any) any {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	switch className {
	case "Date":
		return nil
	case "Main":
		return nil
	case "StringBuf":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.ds.List":
		return nil
	case "haxe.ds._List.GoListIterator":
		return nil
	case "haxe.ds._List.GoListKeyValueIterator":
		return nil
	case "haxe.exceptions.NotImplementedException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.PosException")}
	case "haxe.exceptions.PosException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "haxe.io.Bytes":
		return nil
	case "haxe.io.BytesBuffer":
		return nil
	case "haxe.io.Eof":
		return nil
	case "haxe.io.FPHelper":
		return nil
	case "haxe.io.Input":
		return nil
	case "haxe.io.Output":
		return nil
	case "sys.net.Host":
		return nil
	case "sys.net.Socket":
		return nil
	case "sys.net.SocketInput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Input")}
	case "sys.net.SocketOutput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Output")}
	case "sys.ssl.Certificate":
		return nil
	case "sys.ssl.Key":
		return nil
	case "sys.ssl.Socket":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("sys.net.Socket")}
	case "sys.thread.Deque":
		return nil
	case "sys.thread.EventLoop":
		return nil
	case "sys.thread.Lock":
		return nil
	case "sys.thread.Mutex":
		return nil
	case "sys.thread.NoEventLoopException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "sys.thread.Thread":
		return nil
	default:
		return nil
	}
}

func Type_getClassName(c any) *string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(className)
}

func Type_getClassFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Date":
		return hxrt.NewArray(hxrt.StringFromLiteral("fromMilliseconds"), hxrt.StringFromLiteral("fromString"), hxrt.StringFromLiteral("fromTime"), hxrt.StringFromLiteral("now"))
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("DEFAULT_CERT_PEM"), hxrt.StringFromLiteral("DEFAULT_KEY_PEM"), hxrt.StringFromLiteral("SNI_CERT_PEM"), hxrt.StringFromLiteral("SNI_KEY_PEM"), hxrt.StringFromLiteral("connectAndPrint"), hxrt.StringFromLiteral("main"))
	case "StringBuf":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("sameValue"))
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray()
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.exceptions.NotImplementedException":
		return hxrt.NewArray()
	case "haxe.exceptions.PosException":
		return hxrt.NewArray()
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_fromNativeView"), hxrt.StringFromLiteral("alloc"), hxrt.StringFromLiteral("fastGet"), hxrt.StringFromLiteral("ofData"), hxrt.StringFromLiteral("ofHex"), hxrt.StringFromLiteral("ofString"), hxrt.StringFromLiteral("rawNativeUsesUtf16LE"))
	case "haxe.io.BytesBuffer":
		return hxrt.NewArray()
	case "haxe.io.Eof":
		return hxrt.NewArray()
	case "haxe.io.FPHelper":
		return hxrt.NewArray(hxrt.StringFromLiteral("doubleToI64"), hxrt.StringFromLiteral("floatToI32"), hxrt.StringFromLiteral("i32ToFloat"), hxrt.StringFromLiteral("i64ToDouble"))
	case "haxe.io.Input":
		return hxrt.NewArray()
	case "haxe.io.Output":
		return hxrt.NewArray()
	case "sys.net.Host":
		return hxrt.NewArray(hxrt.StringFromLiteral("fromIPv4"), hxrt.StringFromLiteral("localhost"))
	case "sys.net.Socket":
		return hxrt.NewArray(hxrt.StringFromLiteral("pick"), hxrt.StringFromLiteral("publicAddress"), hxrt.StringFromLiteral("select"))
	case "sys.net.SocketInput":
		return hxrt.NewArray(hxrt.StringFromLiteral("translateReadStatus"))
	case "sys.net.SocketOutput":
		return hxrt.NewArray(hxrt.StringFromLiteral("translateWriteStatus"))
	case "sys.ssl.Certificate":
		return hxrt.NewArray(hxrt.StringFromLiteral("fromString"), hxrt.StringFromLiteral("loadDefaults"), hxrt.StringFromLiteral("loadFile"), hxrt.StringFromLiteral("loadPath"))
	case "sys.ssl.Key":
		return hxrt.NewArray(hxrt.StringFromLiteral("loadFile"), hxrt.StringFromLiteral("readDER"), hxrt.StringFromLiteral("readPEM"))
	case "sys.ssl.Socket":
		return hxrt.NewArray(hxrt.StringFromLiteral("DEFAULT_CA"), hxrt.StringFromLiteral("DEFAULT_VERIFY_CERT"))
	case "sys.thread.Deque":
		return hxrt.NewArray()
	case "sys.thread.EventLoop":
		return hxrt.NewArray(hxrt.StringFromLiteral("__fromHandle"))
	case "sys.thread.Lock":
		return hxrt.NewArray()
	case "sys.thread.Mutex":
		return hxrt.NewArray()
	case "sys.thread.NoEventLoopException":
		return hxrt.NewArray()
	case "sys.thread.Thread":
		return hxrt.NewArray(hxrt.StringFromLiteral("create"), hxrt.StringFromLiteral("createWithEventLoop"), hxrt.StringFromLiteral("current"), hxrt.StringFromLiteral("processEvents"), hxrt.StringFromLiteral("readMessage"), hxrt.StringFromLiteral("runWithEventLoop"))
	default:
		return hxrt.NewArray()
	}
}

func Type_getInstanceFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Date":
		return hxrt.NewArray(hxrt.StringFromLiteral("getDate"), hxrt.StringFromLiteral("getDay"), hxrt.StringFromLiteral("getFullYear"), hxrt.StringFromLiteral("getHours"), hxrt.StringFromLiteral("getMinutes"), hxrt.StringFromLiteral("getMonth"), hxrt.StringFromLiteral("getSeconds"), hxrt.StringFromLiteral("getTime"), hxrt.StringFromLiteral("getTimezoneOffset"), hxrt.StringFromLiteral("getUTCDate"), hxrt.StringFromLiteral("getUTCDay"), hxrt.StringFromLiteral("getUTCFullYear"), hxrt.StringFromLiteral("getUTCHours"), hxrt.StringFromLiteral("getUTCMinutes"), hxrt.StringFromLiteral("getUTCMonth"), hxrt.StringFromLiteral("getUTCSeconds"), hxrt.StringFromLiteral("localParts"), hxrt.StringFromLiteral("ms"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("utcParts"))
	case "Main":
		return hxrt.NewArray()
	case "StringBuf":
		return hxrt.NewArray()
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.ds.List":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("filter"), hxrt.StringFromLiteral("first"), hxrt.StringFromLiteral("isEmpty"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("join"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("last"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("pop"), hxrt.StringFromLiteral("push"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("toString"))
	case "haxe.ds._List.GoListIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "haxe.ds._List.GoListKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("index"), hxrt.StringFromLiteral("items"), hxrt.StringFromLiteral("next"))
	case "haxe.exceptions.NotImplementedException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("posInfos"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.exceptions.PosException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("posInfos"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_dataExposed"), hxrt.StringFromLiteral("__hx_nativeView"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("blit"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("fill"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getData"), hxrt.StringFromLiteral("getDouble"), hxrt.StringFromLiteral("getFloat"), hxrt.StringFromLiteral("getInt32"), hxrt.StringFromLiteral("getInt64"), hxrt.StringFromLiteral("getString"), hxrt.StringFromLiteral("getUInt16"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setDouble"), hxrt.StringFromLiteral("setFloat"), hxrt.StringFromLiteral("setInt32"), hxrt.StringFromLiteral("setInt64"), hxrt.StringFromLiteral("setUInt16"), hxrt.StringFromLiteral("sub"), hxrt.StringFromLiteral("toHex"), hxrt.StringFromLiteral("toString"))
	case "haxe.io.BytesBuffer":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("addByte"), hxrt.StringFromLiteral("addBytes"), hxrt.StringFromLiteral("addDouble"), hxrt.StringFromLiteral("addFloat"), hxrt.StringFromLiteral("addInt32"), hxrt.StringFromLiteral("addInt64"), hxrt.StringFromLiteral("addString"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("getBytes"), hxrt.StringFromLiteral("get_length"), hxrt.StringFromLiteral("length"))
	case "haxe.io.Eof":
		return hxrt.NewArray(hxrt.StringFromLiteral("toString"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray()
	case "haxe.io.Input":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("readAll"), hxrt.StringFromLiteral("readByte"), hxrt.StringFromLiteral("readBytes"), hxrt.StringFromLiteral("readDouble"), hxrt.StringFromLiteral("readFloat"), hxrt.StringFromLiteral("readFullBytes"), hxrt.StringFromLiteral("readInt16"), hxrt.StringFromLiteral("readInt24"), hxrt.StringFromLiteral("readInt32"), hxrt.StringFromLiteral("readInt8"), hxrt.StringFromLiteral("readLine"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("readUInt16"), hxrt.StringFromLiteral("readUInt24"), hxrt.StringFromLiteral("readUntil"), hxrt.StringFromLiteral("set_bigEndian"))
	case "haxe.io.Output":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("flush"), hxrt.StringFromLiteral("prepare"), hxrt.StringFromLiteral("set_bigEndian"), hxrt.StringFromLiteral("write"), hxrt.StringFromLiteral("writeByte"), hxrt.StringFromLiteral("writeBytes"), hxrt.StringFromLiteral("writeDouble"), hxrt.StringFromLiteral("writeFloat"), hxrt.StringFromLiteral("writeFullBytes"), hxrt.StringFromLiteral("writeInput"), hxrt.StringFromLiteral("writeInt16"), hxrt.StringFromLiteral("writeInt24"), hxrt.StringFromLiteral("writeInt32"), hxrt.StringFromLiteral("writeInt8"), hxrt.StringFromLiteral("writeString"), hxrt.StringFromLiteral("writeUInt16"), hxrt.StringFromLiteral("writeUInt24"))
	case "sys.net.Host":
		return hxrt.NewArray(hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("ip"), hxrt.StringFromLiteral("reverse"), hxrt.StringFromLiteral("toString"))
	case "sys.net.Socket":
		return hxrt.NewArray(hxrt.StringFromLiteral("accept"), hxrt.StringFromLiteral("bind"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("connect"), hxrt.StringFromLiteral("custom"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("input"), hxrt.StringFromLiteral("listen"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("peer"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("replaceHandle"), hxrt.StringFromLiteral("setBlocking"), hxrt.StringFromLiteral("setFastSend"), hxrt.StringFromLiteral("setTimeout"), hxrt.StringFromLiteral("shutdown"), hxrt.StringFromLiteral("waitForRead"), hxrt.StringFromLiteral("write"))
	case "sys.net.SocketInput":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("readAll"), hxrt.StringFromLiteral("readByte"), hxrt.StringFromLiteral("readBytes"), hxrt.StringFromLiteral("readDouble"), hxrt.StringFromLiteral("readFloat"), hxrt.StringFromLiteral("readFullBytes"), hxrt.StringFromLiteral("readInt16"), hxrt.StringFromLiteral("readInt24"), hxrt.StringFromLiteral("readInt32"), hxrt.StringFromLiteral("readInt8"), hxrt.StringFromLiteral("readLine"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("readUInt16"), hxrt.StringFromLiteral("readUInt24"), hxrt.StringFromLiteral("readUntil"), hxrt.StringFromLiteral("set_bigEndian"))
	case "sys.net.SocketOutput":
		return hxrt.NewArray(hxrt.StringFromLiteral("bigEndian"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("flush"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("prepare"), hxrt.StringFromLiteral("set_bigEndian"), hxrt.StringFromLiteral("write"), hxrt.StringFromLiteral("writeByte"), hxrt.StringFromLiteral("writeBytes"), hxrt.StringFromLiteral("writeDouble"), hxrt.StringFromLiteral("writeFloat"), hxrt.StringFromLiteral("writeFullBytes"), hxrt.StringFromLiteral("writeInput"), hxrt.StringFromLiteral("writeInt16"), hxrt.StringFromLiteral("writeInt24"), hxrt.StringFromLiteral("writeInt32"), hxrt.StringFromLiteral("writeInt8"), hxrt.StringFromLiteral("writeString"), hxrt.StringFromLiteral("writeUInt16"), hxrt.StringFromLiteral("writeUInt24"))
	case "sys.ssl.Certificate":
		return hxrt.NewArray(hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("addDER"), hxrt.StringFromLiteral("altNames"), hxrt.StringFromLiteral("commonName"), hxrt.StringFromLiteral("get_altNames"), hxrt.StringFromLiteral("get_commonName"), hxrt.StringFromLiteral("get_notAfter"), hxrt.StringFromLiteral("get_notBefore"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("issuer"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("notAfter"), hxrt.StringFromLiteral("notBefore"), hxrt.StringFromLiteral("subject"))
	case "sys.ssl.Key":
		return hxrt.NewArray(hxrt.StringFromLiteral("handle"))
	case "sys.ssl.Socket":
		return hxrt.NewArray(hxrt.StringFromLiteral("accept"), hxrt.StringFromLiteral("addSNICertificate"), hxrt.StringFromLiteral("bind"), hxrt.StringFromLiteral("caCert"), hxrt.StringFromLiteral("close"), hxrt.StringFromLiteral("connect"), hxrt.StringFromLiteral("custom"), hxrt.StringFromLiteral("handle"), hxrt.StringFromLiteral("handshake"), hxrt.StringFromLiteral("host"), hxrt.StringFromLiteral("hostname"), hxrt.StringFromLiteral("input"), hxrt.StringFromLiteral("listen"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("ownCert"), hxrt.StringFromLiteral("ownKey"), hxrt.StringFromLiteral("peer"), hxrt.StringFromLiteral("peerCertificate"), hxrt.StringFromLiteral("read"), hxrt.StringFromLiteral("replaceHandle"), hxrt.StringFromLiteral("setBlocking"), hxrt.StringFromLiteral("setCA"), hxrt.StringFromLiteral("setCertificate"), hxrt.StringFromLiteral("setFastSend"), hxrt.StringFromLiteral("setHostname"), hxrt.StringFromLiteral("setTimeout"), hxrt.StringFromLiteral("shutdown"), hxrt.StringFromLiteral("sniConfig"), hxrt.StringFromLiteral("verifyCert"), hxrt.StringFromLiteral("waitForRead"), hxrt.StringFromLiteral("write"))
	case "sys.thread.Deque":
		return hxrt.NewArray(hxrt.StringFromLiteral("__available"), hxrt.StringFromLiteral("__items"), hxrt.StringFromLiteral("__mutex"), hxrt.StringFromLiteral("add"), hxrt.StringFromLiteral("pop"), hxrt.StringFromLiteral("push"))
	case "sys.thread.EventLoop":
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"), hxrt.StringFromLiteral("cancel"), hxrt.StringFromLiteral("loop"), hxrt.StringFromLiteral("progress"), hxrt.StringFromLiteral("promise"), hxrt.StringFromLiteral("repeat"), hxrt.StringFromLiteral("run"), hxrt.StringFromLiteral("runPromised"), hxrt.StringFromLiteral("wait"))
	case "sys.thread.Lock":
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"), hxrt.StringFromLiteral("release"), hxrt.StringFromLiteral("wait"))
	case "sys.thread.Mutex":
		return hxrt.NewArray(hxrt.StringFromLiteral("__h"), hxrt.StringFromLiteral("acquire"), hxrt.StringFromLiteral("release"), hxrt.StringFromLiteral("tryAcquire"))
	case "sys.thread.NoEventLoopException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "sys.thread.Thread":
		return hxrt.NewArray(hxrt.StringFromLiteral("__id"), hxrt.StringFromLiteral("events"), hxrt.StringFromLiteral("get_events"), hxrt.StringFromLiteral("sendMessage"))
	default:
		return hxrt.NewArray()
	}
}

func Type_getEnumName(e any) *string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(enumName)
}

func Type_resolveClass(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "Date":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringBuf":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.List":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds._List.GoListKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.exceptions.NotImplementedException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.exceptions.PosException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Bytes":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.BytesBuffer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Eof":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.FPHelper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Input":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Output":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.Host":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.Socket":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.SocketInput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.net.SocketOutput":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.ssl.Certificate":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.ssl.Key":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.ssl.Socket":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Deque":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.EventLoop":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Lock":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Mutex":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.NoEventLoopException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.Thread":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_resolveEnum(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "ValueType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Encoding":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Error":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "sys.thread.NextEventTime":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_createInstance(cl any, args *hxrt.Array) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, hxrt_typeArrayValues(args))
	if !ok {
		return nil
	}
	return instance
}

func Type_createEmptyInstance(cl any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassEmptyInstance(className)
	if !ok {
		return nil
	}
	return instance
}

func Type_createEnum(e any, constr *string, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_enumConstructor(e any) *string {
	if hxrt.AnyEqualsNull(e) {
		return nil
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TNull")
		case 1:
			return hxrt.StringFromLiteral("TInt")
		case 2:
			return hxrt.StringFromLiteral("TFloat")
		case 3:
			return hxrt.StringFromLiteral("TBool")
		case 4:
			return hxrt.StringFromLiteral("TObject")
		case 5:
			return hxrt.StringFromLiteral("TFunction")
		case 6:
			return hxrt.StringFromLiteral("TClass")
		case 7:
			return hxrt.StringFromLiteral("TEnum")
		case 8:
			return hxrt.StringFromLiteral("TUnknown")
		default:
			return nil
		}
	case *haxe__io__Encoding:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("UTF8")
		case 1:
			return hxrt.StringFromLiteral("RawNative")
		default:
			return nil
		}
	case *haxe__io__Error:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("Blocked")
		case 1:
			return hxrt.StringFromLiteral("Overflow")
		case 2:
			return hxrt.StringFromLiteral("OutsideBounds")
		case 3:
			return hxrt.StringFromLiteral("Custom")
		default:
			return nil
		}
	case *sys__thread__NextEventTime:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("Now")
		case 1:
			return hxrt.StringFromLiteral("Never")
		case 2:
			return hxrt.StringFromLiteral("AnyTime")
		case 3:
			return hxrt.StringFromLiteral("At")
		default:
			return nil
		}
	default:
		return nil
	}
}

func Type_enumIndex(e any) int {
	if hxrt.AnyEqualsNull(e) {
		return -1
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__io__Encoding:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__io__Error:
		if value == nil {
			return -1
		}
		return value.tag
	case *sys__thread__NextEventTime:
		if value == nil {
			return -1
		}
		return value.tag
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown"))
	case "haxe.io.Encoding":
		return hxrt.NewArray(hxrt.StringFromLiteral("UTF8"), hxrt.StringFromLiteral("RawNative"))
	case "haxe.io.Error":
		return hxrt.NewArray(hxrt.StringFromLiteral("Blocked"), hxrt.StringFromLiteral("Overflow"), hxrt.StringFromLiteral("OutsideBounds"), hxrt.StringFromLiteral("Custom"))
	case "sys.thread.NextEventTime":
		return hxrt.NewArray(hxrt.StringFromLiteral("Now"), hxrt.StringFromLiteral("Never"), hxrt.StringFromLiteral("AnyTime"), hxrt.StringFromLiteral("At"))
	default:
		return hxrt.NewArray()
	}
}

func Type_enumParameters(e any) *hxrt.Array {
	if hxrt.AnyEqualsNull(e) {
		return hxrt.NewArray()
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__io__Encoding:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__io__Error:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *sys__thread__NextEventTime:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	default:
		return hxrt.NewArray()
	}
}

func Type_allEnums(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown)
	case "haxe.io.Encoding":
		return hxrt.NewArray(haxe__io__Encoding_UTF8, haxe__io__Encoding_RawNative)
	case "haxe.io.Error":
		return hxrt.NewArray(haxe__io__Error_Blocked, haxe__io__Error_Overflow, haxe__io__Error_OutsideBounds)
	case "sys.thread.NextEventTime":
		return hxrt.NewArray(sys__thread__NextEventTime_Now, sys__thread__NextEventTime_Never)
	default:
		return hxrt.NewArray()
	}
}

func Type_typeof(v any) *ValueType {
	if hxrt.AnyEqualsNull(v) {
		return ValueType_TNull
	}
	switch v.(type) {
	case *hxrt__TypeClassValue, hxrt__TypeClassValue, *hxrt__TypeEnumValue, hxrt__TypeEnumValue:
		return ValueType_TObject
	}
	if enumValue := Type_getEnum(v); enumValue != nil {
		return ValueType_TEnum(enumValue)
	}
	if classValue := Type_getClass(v); classValue != nil {
		return ValueType_TClass(classValue)
	}
	switch v.(type) {
	case bool:
		return ValueType_TBool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return ValueType_TInt
	case float32, float64:
		return ValueType_TFloat
	case string, *string:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("String")})
	case *hxrt.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	}
	ref := reflect.ValueOf(v)
	if !ref.IsValid() {
		return ValueType_TNull
	}
	switch ref.Kind() {
	case reflect.Func:
		return ValueType_TFunction
	case reflect.Slice, reflect.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:
		return ValueType_TObject
	default:
		return ValueType_TUnknown
	}
}

func Type_enumEq(a any, b any) bool {
	return reflect.DeepEqual(a, b)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	className, ok := hxrt_typeResolvedClassName(value)
	if !ok {
		return nil, false
	}
	switch className {
	default:
		return nil, false
	}
}
