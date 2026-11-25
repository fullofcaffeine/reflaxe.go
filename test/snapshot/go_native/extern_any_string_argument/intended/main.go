package main

import (
	"os"
	"snapshot/hxrt"
	"sync/atomic"
	"text/template"
)

func compile(source *string) *template.Template {
	return func(hx_value_1 any) *template.Template {
		if hx_value_1 == nil {
			var hx_zero_2 *template.Template
			return hx_zero_2
		}
		return hx_value_1.(*template.Template)
	}(go__result_fromValueError(template.New(*hxrt.StdString(hxrt.StringFromLiteral("t"))).Parse(*hxrt.StdString(source))).unwrap())
}

func main() {
	kind := compile(hxrt.StringFromLiteral("{{printf \"%T|%v|%-6s|\" . . .}}{{if eq . \"42\"}}eq{{else}}ne{{end}}\n"))
	plain := hxrt.StringFromLiteral("plain")
	number := hxrt.StringFromLiteral("42")
	label := hxrt.StringFromLiteral("label")
	var missing *string = nil
	report(hxrt.ImportNativeError(kind.Execute(os.Stdout, *hxrt.StdString(plain))))
	report(hxrt.ImportNativeError(kind.Execute(os.Stdout, *hxrt.StdString(number))))
	report(hxrt.ImportNativeError(kind.Execute(os.Stdout, *hxrt.StdString(label))))
	report(hxrt.ImportNativeError(kind.Execute(os.Stdout, *hxrt.StdString(hxrt.StringFromLiteral("literal")))))
	report(hxrt.ImportNativeError(kind.Execute(os.Stdout, *hxrt.StdString(text()))))
	report(hxrt.ImportNativeError(kind.Execute(os.Stdout, *hxrt.StdString(missing))))
	report(hxrt.ImportNativeError(kind.Execute(os.Stdout, *hxrt.StdString(number))))
	scalar := compile(hxrt.StringFromLiteral("{{printf \"%T|%v\" . .}}\n"))
	report(hxrt.ImportNativeError(scalar.Execute(os.Stdout, 7)))
	report(hxrt.ImportNativeError(scalar.Execute(os.Stdout, 2.5)))
	report(hxrt.ImportNativeError(scalar.Execute(os.Stdout, true)))
	box := &atomic.Value{}
	box.Store(*hxrt.StdString(plain))
	back := hxrt.DynamicToString(box.Load())
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(back, hxrt.StringFromLiteral("/")), hxrt.StringLengthStringPtr(back)), hxrt.StringFromLiteral("/")), hxrt.StdString(hxrt.StringEqualStringPtr(back, hxrt.StringFromLiteral("plain")))))
	hxrt.Println(v)
}

func report(error_local *hxrt.NativeErrorValue) {
	if error_local != nil {
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error: "), hxrt.StdString(error_local.Text())))
		hxrt.Println(v)
	}
}

func text() *string {
	return hxrt.StringFromLiteral("call")
}

func go__result_fromValueError(value any, err error) *go___Result {
	if err != nil {
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral(err.Error())))
	}
	return New_go___Result(value, nil)
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
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

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
