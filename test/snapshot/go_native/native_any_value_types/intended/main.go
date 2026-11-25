package main

import (
	"encoding/json"
	"os"
	"regexp/syntax"
	"snapshot/hxrt"
	"text/template"
	"time"
)

func main() {
	created := func() *time.Time {
		hx_extern_value_2 := func() *time.Time {
			hx_extern_value_1 := time.Unix(1795608000, 0)
			return &hx_extern_value_1
		}().UTC()
		return &hx_extern_value_2
	}()
	timeout := func() *ParseDurationResult {
		hx_tuple_3, hx_tuple_4 := time.ParseDuration(*hxrt.StdString(hxrt.StringFromLiteral("1m30s")))
		return New_ParseDurationResult(&hx_tuple_3, func(err interface{ Error() string }) *go___Error {
			if err == nil {
				return nil
			}
			return New_go___Error(hxrt.StringFromLiteral(err.Error()))
		}(hx_tuple_4))
	}().value
	issue := map[string]any{}
	issue["Title"] = "Fix the parser"
	issue["CreatedAt"] = *created
	issue["Month"] = created.Month()
	data := map[string]any{}
	data["Issue"] = issue
	data["Timeout"] = *timeout
	data["Code"] = syntax.ErrMissingBracket
	data["Zero"] = time.Time{}
	values := hxrt.NewArray(created.Month(), syntax.ErrMissingBracket)
	data["Values"] = values.ValuesCopy()
	render(hxrt.StringFromLiteral("{{.Issue.CreatedAt.Year}} {{.Issue.CreatedAt.Format \"2006-01-02\"}} {{.Issue.CreatedAt}}"), data)
	render(hxrt.StringFromLiteral("{{printf \"%T %T %T %T %T\" .Issue.CreatedAt .Issue.Month .Timeout .Code .Values}}"), data)
	render(hxrt.StringFromLiteral("{{.Issue.Month}} {{.Timeout}} {{.Code}} {{if eq .Code \"missing closing ]\"}}eq{{end}} {{range .Values}}[{{printf \"%T\" .}}]{{end}} {{.Zero.IsZero}}"), data)
	error_local := hxrt.ImportNativeError(json.NewEncoder(os.Stdout).Encode(data))
	if error_local != nil {
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("json: "), hxrt.StdString(error_local.Text())))
		hxrt.Println(v)
	}
}

func render(source *string, data map[string]any) {
	parsed := go__result_fromValueError(template.New(*hxrt.StdString(hxrt.StringFromLiteral("format"))).Parse(*hxrt.StdString(source)))
	if !func(hx_value_11 any) bool {
		if hx_value_11 == nil {
			var hx_zero_12 bool
			return hx_zero_12
		}
		return hx_value_11.(bool)
	}(parsed.isOk()) {
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("invalid format template: "), func(hx_value_9 any) *string {
			if hx_value_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_value_9.(*string)
		}(parsed.error())))
		hxrt.Println(v)
		return
	}
	error_local := hxrt.ImportNativeError(func(hx_value_13 any) *template.Template {
		if hx_value_13 == nil {
			var hx_zero_14 *template.Template
			return hx_zero_14
		}
		return hx_value_13.(*template.Template)
	}(parsed.unwrap()).Execute(os.Stdout, data))
	if error_local != nil {
		var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("template execution error: "), hxrt.StdString(error_local.Text())))
		hxrt.Println(v_1)
		return
	}
	hxrt.Println(any(hxrt.StringFromLiteral("")))
}

type I_ParseDurationResult interface {
}

type ParseDurationResult struct {
	__hx_this I_ParseDurationResult
	value     *time.Duration
	error     *go___Error
}

func (self *ParseDurationResult) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_ParseDurationResult(value *time.Duration, error_local *go___Error) *ParseDurationResult {
	return hx_init_ParseDurationResult(hx_alloc_ParseDurationResult(nil), value, error_local)
}

func hx_alloc_ParseDurationResult(owner any) *ParseDurationResult {
	self := &ParseDurationResult{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_ParseDurationResult)
	}
	return self
}

func hx_init_ParseDurationResult(self *ParseDurationResult, value *time.Duration, error_local *go___Error) *ParseDurationResult {
	self.value = value
	self.error = error_local
	return self
}

func (self *ParseDurationResult) HxClassName() string {
	return "ParseDurationResult"
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
