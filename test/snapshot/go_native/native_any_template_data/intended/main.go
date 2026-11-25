package main

import (
	"encoding/json"
	"os"
	"snapshot/hxrt"
	"text/template"
)

func issue(id *string, title *string, priority int, labels *hxrt.Array) map[string]any {
	data := map[string]any{}
	data["ID"] = *hxrt.StdString(id)
	data["Title"] = *hxrt.StdString(title)
	data["Priority"] = priority
	data["Labels"] = func(hx_lambda_raw_1 []any) []string {
		hx_lambda_out_2 := make([]string, 0, len(hx_lambda_raw_1))
		for _, hx_lambda_item_3 := range hx_lambda_raw_1 {
			hx_lambda_out_2 = append(hx_lambda_out_2, *hxrt.StdString(hx_lambda_item_3))
		}
		return hx_lambda_out_2
	}(labels.Values())
	return data
}

func main() {
	parent := issue(hxrt.StringFromLiteral("bd-1"), hxrt.StringFromLiteral("Fix the parser"), 1, hxrt.NewArray(hxrt.StringFromLiteral("parser"), hxrt.StringFromLiteral("p1")))
	dependency := map[string]any{}
	dependency["IssueID"] = "bd-1"
	dependency["DependsOnID"] = "bd-22"
	dependency["Type"] = "blocks"
	data := map[string]any{}
	data["IssueID"] = "bd-1"
	data["DependsOnID"] = "bd-22"
	data["Type"] = "blocks"
	data["Issue"] = parent
	data["Dependency"] = dependency
	data["Score"] = 2.0
	data["Ready"] = true
	data["Missing"] = nil
	nested := hxrt.NewArray("nested")
	items := hxrt.NewArray("one", 2, 3.5, false, nil, nested.ValuesCopy())
	data["Items"] = items.ValuesCopy()
	render(hxrt.StringFromLiteral("{{.IssueID}} {{.DependsOnID}}"), data)
	render(hxrt.StringFromLiteral("{{printf \"%-14s\" .IssueID}}|{{if eq .Type \"blocks\"}}blocks{{else}}other{{end}}|{{.Issue.Title}}|{{.Dependency.DependsOnID}}"), data)
	render(hxrt.StringFromLiteral("{{range $i, $label := .Issue.Labels}}{{if $i}},{{end}}{{$label}}{{end}} {{len .Issue.Labels}} {{index .Issue.Labels 1}}"), data)
	render(hxrt.StringFromLiteral("{{printf \"%T %T %T %T %T %T %T %T %T\" . .IssueID .Issue.Priority .Score .Ready .Missing .Issue.Labels .Items .Issue}}"), data)
	render(hxrt.StringFromLiteral("{{range .Items}}[{{printf \"%T\" .}} {{.}}]{{end}}"), data)
	render(hxrt.StringFromLiteral("{{.Missing}} {{.Score}} {{printf \"%.1f\" .Score}} {{.Issue.Priority | printf \"P%d\"}} {{if .Ready}}ready{{end}} {{.Nope}}"), data)
	render(hxrt.StringFromLiteral("{{.Issue.Title.Missing}}"), data)
	encoder := json.NewEncoder(os.Stdout)
	error_local := hxrt.ImportNativeError(encoder.Encode(data))
	if error_local != nil {
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("json: "), hxrt.StdString(error_local.Text())))
		hxrt.Println(v)
	}
	ordered := map[string]any{}
	hx_native_map_5 := target(ordered)
	hx_native_key_6 := *hxrt.StdString(note(hxrt.StringFromLiteral("key")))
	var hx_if_4 any
	if trail.Len() > 0 {
		hx_if_4 = *hxrt.StdString(note(hxrt.StringFromLiteral("value")))
	} else {
		hx_if_4 = "other"
	}
	hx_native_map_5[hx_native_key_6] = hx_if_4
	var v_1 any = any(hxrt.StringJoinAny(trail.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_1)
	hxrt.ImportNativeError(encoder.Encode(ordered))
}

func note(step *string) *string {
	hx_arr_7 := trail
	hx_arr_7.Push(step)
	return step
}

func render(source *string, data map[string]any) {
	parsed := go__result_fromValueError(template.New(*hxrt.StdString(hxrt.StringFromLiteral("format"))).Parse(*hxrt.StdString(source)))
	if !func(hx_value_10 any) bool {
		if hx_value_10 == nil {
			var hx_zero_11 bool
			return hx_zero_11
		}
		return hx_value_10.(bool)
	}(parsed.isOk()) {
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("invalid format template: "), func(hx_value_8 any) *string {
			if hx_value_8 == nil {
				var hx_zero_9 *string
				return hx_zero_9
			}
			return hx_value_8.(*string)
		}(parsed.error())))
		hxrt.Println(v)
		return
	}
	error_local := hxrt.ImportNativeError(func(hx_value_12 any) *template.Template {
		if hx_value_12 == nil {
			var hx_zero_13 *template.Template
			return hx_zero_13
		}
		return hx_value_12.(*template.Template)
	}(parsed.unwrap()).Execute(os.Stdout, data))
	if error_local != nil {
		var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("template execution error: "), hxrt.StdString(error_local.Text())))
		hxrt.Println(v_1)
		return
	}
	hxrt.Println(any(hxrt.StringFromLiteral("")))
}

func target(map_ map[string]any) map[string]any {
	hx_arr_14 := trail
	hx_arr_14.Push(hxrt.StringFromLiteral("map"))
	return map_
}

var trail *hxrt.Array = hxrt.NewArray()

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
