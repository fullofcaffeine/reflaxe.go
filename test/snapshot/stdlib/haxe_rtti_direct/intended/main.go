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

var Demo___rtti *string = hxrt.StringFromLiteral("<class path=\"Demo\" params=\"\" module=\"Main\">\n\t<field public=\"1\" expr=\"&quot;value&quot;\" line=\"6\" static=\"1\">\n\t\t<c path=\"String\"/>\n\t\t<meta><m n=\":value\"><e>\"value\"</e></m></meta>\n\t</field>\n\t<meta>\n\t\t<m n=\":directlyUsed\"/>\n\t\t<m n=\":rtti\"/>\n\t</meta>\n</class>")

var Demo_field *string = hxrt.StringFromLiteral("value")

func main() {
	info := haxe__rtti__Rtti_getRtti(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Demo")})
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("path="), func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(hxrt.AnonymousGet(info, "path"))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("staticType="), haxe__rtti__CTypeTools_toString(func(hx_value_5 any) *haxe__rtti__CType {
		if hx_value_5 == nil {
			var hx_zero_6 *haxe__rtti__CType
			return hx_zero_6
		}
		return hx_value_5.(*haxe__rtti__CType)
	}(hxrt.AnonymousGet(hx_structural_recover(func(hx_value_3 any) *hxrt.Array {
		if hx_value_3 == nil {
			var hx_zero_4 *hxrt.Array
			return hx_zero_4
		}
		return hx_value_3.(*hxrt.Array)
	}(hxrt.AnonymousGet(info, "statics")).Get(0)), "type")))))
	hxrt.Println(v_1)
	var rawRtti any = Reflect_field(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Demo")}, hxrt.StringFromLiteral("__rtti"))
	parsed := New_haxe__rtti__XmlParser().__hx_this.processElement(Xml_parse(hxrt.StdString(rawRtti)).__hx_this.firstElement())
	if parsed.tag == 1 {
		_g := hx_structural_recover(parsed.params[0])
		c := _g
		var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("parsedPath="), func(hx_value_7 any) *string {
			if hx_value_7 == nil {
				var hx_zero_8 *string
				return hx_zero_8
			}
			return hx_value_7.(*string)
		}(hxrt.AnonymousGet(c, "path"))))
		hxrt.Println(v_2)
	} else {
		hxrt.Println(any(hxrt.StringFromLiteral("parsed=unexpected")))
	}
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_1 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_1.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_1.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_1.high = func(hx_value_2 any) int {
						if hx_value_2 == nil {
							var hx_zero_3 int
							return hx_zero_3
						}
						return hx_value_2.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_1.low = func(hx_value_4 any) int {
						if hx_value_4 == nil {
							var hx_zero_5 int
							return hx_zero_5
						}
						return hx_value_4.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__rtti__XmlParser:
		return func(hx_structural_source_6 *haxe__rtti__XmlParser) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_6, func(name string) hxrt.AnonymousField {
				switch name {
				case "root":
					return hxrt.AnonymousField{Value: hx_structural_source_6.root, Found: true}
				case "curplatform":
					return hxrt.AnonymousField{Value: hx_structural_source_6.curplatform, Found: true}
				case "sort":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array)] {
						hx_receiver_11 := hx_structural_source_6
						return hxrt.BoundFunction(func(l *hxrt.Array) {
							hx_receiver_11.__hx_this.sort(l)
						}, func() any {
							return hx_receiver_11.__hx_this
						}, "sort")
					}(), Found: true}
				case "sortFields":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array)] {
						hx_receiver_12 := hx_structural_source_6
						return hxrt.BoundFunction(func(a *hxrt.Array) {
							hx_receiver_12.__hx_this.sortFields(a)
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "sortFields")
					}(), Found: true}
				case "process":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml, *string)] {
						hx_receiver_13 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml, platform *string) {
							hx_receiver_13.__hx_this.process(x, platform)
						}, func() any {
							return hx_receiver_13.__hx_this
						}, "process")
					}(), Found: true}
				case "mergeRights":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
						hx_receiver_14 := hx_structural_source_6
						return hxrt.BoundFunction(func(f1 hxrt.Anonymous, f2 hxrt.Anonymous) bool {
							return hx_receiver_14.__hx_this.mergeRights(f1, f2)
						}, func() any {
							return hx_receiver_14.__hx_this
						}, "mergeRights")
					}(), Found: true}
				case "mergeDoc":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
						hx_receiver_15 := hx_structural_source_6
						return hxrt.BoundFunction(func(f1 hxrt.Anonymous, f2 hxrt.Anonymous) bool {
							return hx_receiver_15.__hx_this.mergeDoc(f1, f2)
						}, func() any {
							return hx_receiver_15.__hx_this
						}, "mergeDoc")
					}(), Found: true}
				case "mergeFields":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
						hx_receiver_16 := hx_structural_source_6
						return hxrt.BoundFunction(func(f hxrt.Anonymous, f2 hxrt.Anonymous) bool {
							return hx_receiver_16.__hx_this.mergeFields(f, f2)
						}, func() any {
							return hx_receiver_16.__hx_this
						}, "mergeFields")
					}(), Found: true}
				case "newField":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_6.newField), Found: true}
				case "mergeClasses":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
						hx_receiver_17 := hx_structural_source_6
						return hxrt.BoundFunction(func(c hxrt.Anonymous, c2 hxrt.Anonymous) bool {
							return hx_receiver_17.__hx_this.mergeClasses(c, c2)
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "mergeClasses")
					}(), Found: true}
				case "mergeEnums":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
						hx_receiver_18 := hx_structural_source_6
						return hxrt.BoundFunction(func(e hxrt.Anonymous, e2 hxrt.Anonymous) bool {
							return hx_receiver_18.__hx_this.mergeEnums(e, e2)
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "mergeEnums")
					}(), Found: true}
				case "mergeTypedefs":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
						hx_receiver_19 := hx_structural_source_6
						return hxrt.BoundFunction(func(t hxrt.Anonymous, t2 hxrt.Anonymous) bool {
							return hx_receiver_19.__hx_this.mergeTypedefs(t, t2)
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "mergeTypedefs")
					}(), Found: true}
				case "mergeAbstracts":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
						hx_receiver_20 := hx_structural_source_6
						return hxrt.BoundFunction(func(a hxrt.Anonymous, a2 hxrt.Anonymous) bool {
							return hx_receiver_20.__hx_this.mergeAbstracts(a, a2)
						}, func() any {
							return hx_receiver_20.__hx_this
						}, "mergeAbstracts")
					}(), Found: true}
				case "merge":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__rtti__TypeTree)] {
						hx_receiver_21 := hx_structural_source_6
						return hxrt.BoundFunction(func(t *haxe__rtti__TypeTree) {
							hx_receiver_21.__hx_this.merge(t)
						}, func() any {
							return hx_receiver_21.__hx_this
						}, "merge")
					}(), Found: true}
				case "mkPath":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *string] {
						hx_receiver_22 := hx_structural_source_6
						return hxrt.BoundFunction(func(p *string) *string {
							return hx_receiver_22.__hx_this.mkPath(p)
						}, func() any {
							return hx_receiver_22.__hx_this
						}, "mkPath")
					}(), Found: true}
				case "mkTypeParams":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_23 := hx_structural_source_6
						return hxrt.BoundFunction(func(p *string) *hxrt.Array {
							return hx_receiver_23.__hx_this.mkTypeParams(p)
						}, func() any {
							return hx_receiver_23.__hx_this
						}, "mkTypeParams")
					}(), Found: true}
				case "mkRights":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *haxe__rtti__Rights] {
						hx_receiver_24 := hx_structural_source_6
						return hxrt.BoundFunction(func(r *string) *haxe__rtti__Rights {
							return hx_receiver_24.__hx_this.mkRights(r)
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "mkRights")
					}(), Found: true}
				case "xroot":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml)] {
						hx_receiver_25 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) {
							hx_receiver_25.__hx_this.xroot(x)
						}, func() any {
							return hx_receiver_25.__hx_this
						}, "xroot")
					}(), Found: true}
				case "processElement":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) *haxe__rtti__TypeTree] {
						hx_receiver_26 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) *haxe__rtti__TypeTree {
							return hx_receiver_26.__hx_this.processElement(x)
						}, func() any {
							return hx_receiver_26.__hx_this
						}, "processElement")
					}(), Found: true}
				case "xmeta":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) *hxrt.Array] {
						hx_receiver_27 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) *hxrt.Array {
							return hx_receiver_27.__hx_this.xmeta(x)
						}, func() any {
							return hx_receiver_27.__hx_this
						}, "xmeta")
					}(), Found: true}
				case "xoverloads":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) *hxrt.Array] {
						hx_receiver_28 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) *hxrt.Array {
							return hx_receiver_28.__hx_this.xoverloads(x)
						}, func() any {
							return hx_receiver_28.__hx_this
						}, "xoverloads")
					}(), Found: true}
				case "xpath":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
						hx_receiver_29 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
							return hx_receiver_29.__hx_this.xpath(x)
						}, func() any {
							return hx_receiver_29.__hx_this
						}, "xpath")
					}(), Found: true}
				case "xclass":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
						hx_receiver_30 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
							return hx_receiver_30.__hx_this.xclass(x)
						}, func() any {
							return hx_receiver_30.__hx_this
						}, "xclass")
					}(), Found: true}
				case "xclassfield":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml, any) hxrt.Anonymous] {
						hx_receiver_31 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml, defPublic any) hxrt.Anonymous {
							if defPublic == nil {
								defPublic = false
							}
							return hx_receiver_31.__hx_this.xclassfield(x, func(hx_value_32 any) bool {
								if hx_value_32 == nil {
									var hx_zero_33 bool
									return hx_zero_33
								}
								return hx_value_32.(bool)
							}(defPublic))
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "xclassfield")
					}(), Found: true}
				case "xenum":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
						hx_receiver_34 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
							return hx_receiver_34.__hx_this.xenum(x)
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "xenum")
					}(), Found: true}
				case "xenumfield":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
						hx_receiver_35 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
							return hx_receiver_35.__hx_this.xenumfield(x)
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "xenumfield")
					}(), Found: true}
				case "xabstract":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
						hx_receiver_36 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
							return hx_receiver_36.__hx_this.xabstract(x)
						}, func() any {
							return hx_receiver_36.__hx_this
						}, "xabstract")
					}(), Found: true}
				case "xtypedef":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
						hx_receiver_37 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
							return hx_receiver_37.__hx_this.xtypedef(x)
						}, func() any {
							return hx_receiver_37.__hx_this
						}, "xtypedef")
					}(), Found: true}
				case "xtype":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) *haxe__rtti__CType] {
						hx_receiver_38 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) *haxe__rtti__CType {
							return hx_receiver_38.__hx_this.xtype(x)
						}, func() any {
							return hx_receiver_38.__hx_this
						}, "xtype")
					}(), Found: true}
				case "xtypeparams":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) *hxrt.Array] {
						hx_receiver_39 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) *hxrt.Array {
							return hx_receiver_39.__hx_this.xtypeparams(x)
						}, func() any {
							return hx_receiver_39.__hx_this
						}, "xtypeparams")
					}(), Found: true}
				case "defplat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.Array] {
						hx_receiver_40 := hx_structural_source_6
						return hxrt.BoundFunction(func() *hxrt.Array {
							return hx_receiver_40.__hx_this.defplat()
						}, func() any {
							return hx_receiver_40.__hx_this
						}, "defplat")
					}(), Found: true}
				case "joinStringArray":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Array, *string) *string] {
						hx_receiver_41 := hx_structural_source_6
						return hxrt.BoundFunction(func(values *hxrt.Array, separator *string) *string {
							return hx_receiver_41.__hx_this.joinStringArray(values, separator)
						}, func() any {
							return hx_receiver_41.__hx_this
						}, "joinStringArray")
					}(), Found: true}
				case "splitString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string) *hxrt.Array] {
						hx_receiver_42 := hx_structural_source_6
						return hxrt.BoundFunction(func(value *string, separator *string) *hxrt.Array {
							return hx_receiver_42.__hx_this.splitString(value, separator)
						}, func() any {
							return hx_receiver_42.__hx_this
						}, "splitString")
					}(), Found: true}
				case "findSeparator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, int) int] {
						hx_receiver_43 := hx_structural_source_6
						return hxrt.BoundFunction(func(value *string, separator *string, start int) int {
							return hx_receiver_43.__hx_this.findSeparator(value, separator, start)
						}, func() any {
							return hx_receiver_43.__hx_this
						}, "findSeparator")
					}(), Found: true}
				case "requireAttr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml, *string) *string] {
						hx_receiver_44 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml, name *string) *string {
							return hx_receiver_44.__hx_this.requireAttr(x, name)
						}, func() any {
							return hx_receiver_44.__hx_this
						}, "requireAttr")
					}(), Found: true}
				case "hasNamedElement":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml, *string) bool] {
						hx_receiver_45 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml, name *string) bool {
							return hx_receiver_45.__hx_this.hasNamedElement(x, name)
						}, func() any {
							return hx_receiver_45.__hx_this
						}, "hasNamedElement")
					}(), Found: true}
				case "requireNamedElement":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml, *string) *Xml] {
						hx_receiver_46 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml, name *string) *Xml {
							return hx_receiver_46.__hx_this.requireNamedElement(x, name)
						}, func() any {
							return hx_receiver_46.__hx_this
						}, "requireNamedElement")
					}(), Found: true}
				case "requireFirstElement":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) *Xml] {
						hx_receiver_47 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) *Xml {
							return hx_receiver_47.__hx_this.requireFirstElement(x)
						}, func() any {
							return hx_receiver_47.__hx_this
						}, "requireFirstElement")
					}(), Found: true}
				case "nodeDisplayName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) *string] {
						hx_receiver_48 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) *string {
							return hx_receiver_48.__hx_this.nodeDisplayName(x)
						}, func() any {
							return hx_receiver_48.__hx_this
						}, "nodeDisplayName")
					}(), Found: true}
				case "elementName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) *string] {
						hx_receiver_49 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) *string {
							return hx_receiver_49.__hx_this.elementName(x)
						}, func() any {
							return hx_receiver_49.__hx_this
						}, "elementName")
					}(), Found: true}
				case "innerData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) *string] {
						hx_receiver_50 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) *string {
							return hx_receiver_50.__hx_this.innerData(x)
						}, func() any {
							return hx_receiver_50.__hx_this
						}, "innerData")
					}(), Found: true}
				case "innerHTML":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) *string] {
						hx_receiver_51 := hx_structural_source_6
						return hxrt.BoundFunction(func(x *Xml) *string {
							return hx_receiver_51.__hx_this.innerHTML(x)
						}, func() any {
							return hx_receiver_51.__hx_this
						}, "innerHTML")
					}(), Found: true}
				case "parseIntString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) int] {
						hx_receiver_52 := hx_structural_source_6
						return hxrt.BoundFunction(func(value *string) int {
							return hx_receiver_52.__hx_this.parseIntString(value)
						}, func() any {
							return hx_receiver_52.__hx_this
						}, "parseIntString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "root":
					hx_structural_source_6.root = func(hx_value_7 any) *hxrt.Array {
						if hx_value_7 == nil {
							var hx_zero_8 *hxrt.Array
							return hx_zero_8
						}
						return hx_value_7.(*hxrt.Array)
					}(incoming)
					return true
				case "curplatform":
					hx_structural_source_6.curplatform = func(hx_value_9 any) *string {
						if hx_value_9 == nil {
							var hx_zero_10 *string
							return hx_zero_10
						}
						return hx_value_9.(*string)
					}(incoming)
					return true
				case "newField":
					hx_structural_source_6.newField = hx_recover_callable_40a93da1(incoming)
					return true
				}
				return false
			})
		}(source)
	case *Xml:
		return func(hx_structural_source_53 *Xml) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_53, func(name string) hxrt.AnonymousField {
				switch name {
				case "nodeType":
					return hxrt.AnonymousField{Value: hx_structural_source_53.nodeType, Found: true}
				case "parent":
					return hxrt.AnonymousField{Value: hx_structural_source_53.parent, Found: true}
				case "children":
					return hxrt.AnonymousField{Value: hx_structural_source_53.children, Found: true}
				case "attributeMap":
					return hxrt.AnonymousField{Value: hx_structural_source_53.attributeMap, Found: true}
				case "get_nodeName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_58 := hx_structural_source_53
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_58.__hx_this.get_nodeName()
						}, func() any {
							return hx_receiver_58.__hx_this
						}, "get_nodeName")
					}(), Found: true}
				case "set_nodeName":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *string] {
						hx_receiver_59 := hx_structural_source_53
						return hxrt.BoundFunction(func(v *string) *string {
							return hx_receiver_59.__hx_this.set_nodeName(v)
						}, func() any {
							return hx_receiver_59.__hx_this
						}, "set_nodeName")
					}(), Found: true}
				case "get_nodeValue":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_60 := hx_structural_source_53
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_60.__hx_this.get_nodeValue()
						}, func() any {
							return hx_receiver_60.__hx_this
						}, "get_nodeValue")
					}(), Found: true}
				case "set_nodeValue":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *string] {
						hx_receiver_61 := hx_structural_source_53
						return hxrt.BoundFunction(func(v *string) *string {
							return hx_receiver_61.__hx_this.set_nodeValue(v)
						}, func() any {
							return hx_receiver_61.__hx_this
						}, "set_nodeValue")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *string] {
						hx_receiver_62 := hx_structural_source_53
						return hxrt.BoundFunction(func(att *string) *string {
							return hx_receiver_62.__hx_this.get(att)
						}, func() any {
							return hx_receiver_62.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_63 := hx_structural_source_53
						return hxrt.BoundFunction(func(att *string, value *string) {
							hx_receiver_63.__hx_this.set(att, value)
						}, func() any {
							return hx_receiver_63.__hx_this
						}, "set")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_64 := hx_structural_source_53
						return hxrt.BoundFunction(func(att *string) {
							hx_receiver_64.__hx_this.remove(att)
						}, func() any {
							return hx_receiver_64.__hx_this
						}, "remove")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_65 := hx_structural_source_53
						return hxrt.BoundFunction(func(att *string) bool {
							return hx_receiver_65.__hx_this.exists(att)
						}, func() any {
							return hx_receiver_65.__hx_this
						}, "exists")
					}(), Found: true}
				case "attributes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_66 := hx_structural_source_53
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_66.__hx_this.attributes()
						}, func() any {
							return hx_receiver_66.__hx_this
						}, "attributes")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_67 := hx_structural_source_53
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_67.__hx_this.iterator()
						}, func() any {
							return hx_receiver_67.__hx_this
						}, "iterator")
					}(), Found: true}
				case "elements":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_68 := hx_structural_source_53
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_68.__hx_this.elements()
						}, func() any {
							return hx_receiver_68.__hx_this
						}, "elements")
					}(), Found: true}
				case "elementsNamed":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) hxrt.Anonymous] {
						hx_receiver_69 := hx_structural_source_53
						return hxrt.BoundFunction(func(name *string) hxrt.Anonymous {
							return hx_receiver_69.__hx_this.elementsNamed(name)
						}, func() any {
							return hx_receiver_69.__hx_this
						}, "elementsNamed")
					}(), Found: true}
				case "firstChild":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *Xml] {
						hx_receiver_70 := hx_structural_source_53
						return hxrt.BoundFunction(func() *Xml {
							return hx_receiver_70.__hx_this.firstChild()
						}, func() any {
							return hx_receiver_70.__hx_this
						}, "firstChild")
					}(), Found: true}
				case "firstElement":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *Xml] {
						hx_receiver_71 := hx_structural_source_53
						return hxrt.BoundFunction(func() *Xml {
							return hx_receiver_71.__hx_this.firstElement()
						}, func() any {
							return hx_receiver_71.__hx_this
						}, "firstElement")
					}(), Found: true}
				case "addChild":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml)] {
						hx_receiver_72 := hx_structural_source_53
						return hxrt.BoundFunction(func(x *Xml) {
							hx_receiver_72.__hx_this.addChild(x)
						}, func() any {
							return hx_receiver_72.__hx_this
						}, "addChild")
					}(), Found: true}
				case "removeChild":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) bool] {
						hx_receiver_73 := hx_structural_source_53
						return hxrt.BoundFunction(func(x *Xml) bool {
							return hx_receiver_73.__hx_this.removeChild(x)
						}, func() any {
							return hx_receiver_73.__hx_this
						}, "removeChild")
					}(), Found: true}
				case "insertChild":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml, int)] {
						hx_receiver_74 := hx_structural_source_53
						return hxrt.BoundFunction(func(x *Xml, pos int) {
							hx_receiver_74.__hx_this.insertChild(x, pos)
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "insertChild")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_75 := hx_structural_source_53
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_75.__hx_this.toString()
						}, func() any {
							return hx_receiver_75.__hx_this
						}, "toString")
					}(), Found: true}
				case "ensureElementType":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_76 := hx_structural_source_53
						return hxrt.BoundFunction(func() {
							hx_receiver_76.__hx_this.ensureElementType()
						}, func() any {
							return hx_receiver_76.__hx_this
						}, "ensureElementType")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "children":
					hx_structural_source_53.children = func(hx_value_54 any) *hxrt.Array {
						if hx_value_54 == nil {
							var hx_zero_55 *hxrt.Array
							return hx_zero_55
						}
						return hx_value_54.(*hxrt.Array)
					}(incoming)
					return true
				case "attributeMap":
					hx_structural_source_53.attributeMap = func(hx_value_56 any) *haxe__ds__StringMap {
						if hx_value_56 == nil {
							var hx_zero_57 *haxe__ds__StringMap
							return hx_zero_57
						}
						return hx_value_56.(*haxe__ds__StringMap)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_77 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_77, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_77.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_77.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_80 := hx_structural_source_77
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_80.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_80.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_81 := hx_structural_source_77
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_81.__hx_this.next()
						}, func() any {
							return hx_receiver_81.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_77.map_ = func(hx_value_78 any) haxe__IMap {
						if hx_value_78 == nil {
							var hx_zero_79 haxe__IMap
							return hx_zero_79
						}
						return hx_value_78.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_77.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__StringMap:
		return func(hx_structural_source_82 *haxe__ds__StringMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_82, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_82.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, any)] {
						hx_receiver_85 := hx_structural_source_82
						return hxrt.BoundFunction(func(key *string, value any) {
							hx_receiver_85.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_85.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_86 := hx_structural_source_82
						return hxrt.BoundFunction(func(key *string) any {
							return hx_receiver_86.__hx_this.get(key)
						}, func() any {
							return hx_receiver_86.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_87 := hx_structural_source_82
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_87.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_87.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_88 := hx_structural_source_82
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_88.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_88.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_89 := hx_structural_source_82
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_89.__hx_this.keys()
						}, func() any {
							return hx_receiver_89.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_90 := hx_structural_source_82
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_90.__hx_this.iterator()
						}, func() any {
							return hx_receiver_90.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_91 := hx_structural_source_82
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_91.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_91.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_92 := hx_structural_source_82
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_92.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_92.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_93 := hx_structural_source_82
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_93.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_93.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_94 := hx_structural_source_82
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_94.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_94.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_95 := hx_structural_source_82
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_95.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_95.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_96 := hx_structural_source_82
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_96.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_96.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__StringMap] {
						hx_receiver_97 := hx_structural_source_82
						return hxrt.BoundFunction(func() *haxe__ds__StringMap {
							return hx_receiver_97.__hx_this.copy()
						}, func() any {
							return hx_receiver_97.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_98 := hx_structural_source_82
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_98.__hx_this.toString()
						}, func() any {
							return hx_receiver_98.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_99 := hx_structural_source_82
						return hxrt.BoundFunction(func() {
							hx_receiver_99.__hx_this.clear()
						}, func() any {
							return hx_receiver_99.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_82.h = func(hx_value_83 any) *hxrt.StringMapCell {
						if hx_value_83 == nil {
							var hx_zero_84 *hxrt.StringMapCell
							return hx_zero_84
						}
						return hx_value_83.(*hxrt.StringMapCell)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_100 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_100, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_100.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_100.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_100.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_100.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_103 := hx_structural_source_100
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_103.__hx_this.get_message()
						}, func() any {
							return hx_receiver_103.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_104 := hx_structural_source_100
						return hxrt.BoundFunction(func() any {
							return hx_receiver_104.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_104.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_105 := hx_structural_source_100
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_105.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_105.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_106 := hx_structural_source_100
						return hxrt.BoundFunction(func() any {
							return hx_receiver_106.__hx_this.get_native()
						}, func() any {
							return hx_receiver_106.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_107 := hx_structural_source_100
						return hxrt.BoundFunction(func() any {
							return hx_receiver_107.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_107.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_108 := hx_structural_source_100
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_108.__hx_this.toString()
						}, func() any {
							return hx_receiver_108.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_109 := hx_structural_source_100
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_109.__hx_this.details()
						}, func() any {
							return hx_receiver_109.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_100.savedMessage = func(hx_value_101 any) *string {
						if hx_value_101 == nil {
							var hx_zero_102 *string
							return hx_zero_102
						}
						return hx_value_101.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_100.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_100.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_100.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__xml__XmlParserException:
		return func(hx_structural_source_110 *haxe__xml__XmlParserException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_110, func(name string) hxrt.AnonymousField {
				switch name {
				case "message":
					return hxrt.AnonymousField{Value: hx_structural_source_110.message, Found: true}
				case "lineNumber":
					return hxrt.AnonymousField{Value: hx_structural_source_110.lineNumber, Found: true}
				case "positionAtLine":
					return hxrt.AnonymousField{Value: hx_structural_source_110.positionAtLine, Found: true}
				case "position":
					return hxrt.AnonymousField{Value: hx_structural_source_110.position, Found: true}
				case "xml":
					return hxrt.AnonymousField{Value: hx_structural_source_110.xml, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_121 := hx_structural_source_110
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_121.__hx_this.toString()
						}, func() any {
							return hx_receiver_121.__hx_this
						}, "toString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "message":
					hx_structural_source_110.message = func(hx_value_111 any) *string {
						if hx_value_111 == nil {
							var hx_zero_112 *string
							return hx_zero_112
						}
						return hx_value_111.(*string)
					}(incoming)
					return true
				case "lineNumber":
					hx_structural_source_110.lineNumber = func(hx_value_113 any) int {
						if hx_value_113 == nil {
							var hx_zero_114 int
							return hx_zero_114
						}
						return hx_value_113.(int)
					}(incoming)
					return true
				case "positionAtLine":
					hx_structural_source_110.positionAtLine = func(hx_value_115 any) int {
						if hx_value_115 == nil {
							var hx_zero_116 int
							return hx_zero_116
						}
						return hx_value_115.(int)
					}(incoming)
					return true
				case "position":
					hx_structural_source_110.position = func(hx_value_117 any) int {
						if hx_value_117 == nil {
							var hx_zero_118 int
							return hx_zero_118
						}
						return hx_value_117.(int)
					}(incoming)
					return true
				case "xml":
					hx_structural_source_110.xml = func(hx_value_119 any) *string {
						if hx_value_119 == nil {
							var hx_zero_120 *string
							return hx_zero_120
						}
						return hx_value_119.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__xml__Printer:
		return func(hx_structural_source_122 *haxe__xml__Printer) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_122, func(name string) hxrt.AnonymousField {
				switch name {
				case "output":
					return hxrt.AnonymousField{Value: hx_structural_source_122.output, Found: true}
				case "pretty":
					return hxrt.AnonymousField{Value: hx_structural_source_122.pretty, Found: true}
				case "writeNode":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml, *string)] {
						hx_receiver_125 := hx_structural_source_122
						return hxrt.BoundFunction(func(value *Xml, tabs *string) {
							hx_receiver_125.__hx_this.writeNode(value, tabs)
						}, func() any {
							return hx_receiver_125.__hx_this
						}, "writeNode")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_126 := hx_structural_source_122
						return hxrt.BoundFunction(func(input *string) {
							hx_receiver_126.__hx_this.write(input)
						}, func() any {
							return hx_receiver_126.__hx_this
						}, "write")
					}(), Found: true}
				case "newline":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_127 := hx_structural_source_122
						return hxrt.BoundFunction(func() {
							hx_receiver_127.__hx_this.newline()
						}, func() any {
							return hx_receiver_127.__hx_this
						}, "newline")
					}(), Found: true}
				case "hasChildren":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*Xml) bool] {
						hx_receiver_128 := hx_structural_source_122
						return hxrt.BoundFunction(func(value *Xml) bool {
							return hx_receiver_128.__hx_this.hasChildren(value)
						}, func() any {
							return hx_receiver_128.__hx_this
						}, "hasChildren")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "output":
					hx_structural_source_122.output = hx_recover_class_StringBuf(incoming)
					return true
				case "pretty":
					hx_structural_source_122.pretty = func(hx_value_123 any) bool {
						if hx_value_123 == nil {
							var hx_zero_124 bool
							return hx_zero_124
						}
						return hx_value_123.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__ArrayIterator:
		return func(hx_structural_source_129 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_129, func(name string) hxrt.AnonymousField {
				switch name {
				case "array":
					return hxrt.AnonymousField{Value: hx_structural_source_129.array, Found: true}
				case "current":
					return hxrt.AnonymousField{Value: hx_structural_source_129.current, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_132 := hx_structural_source_129
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_132.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_132.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_133 := hx_structural_source_129
						return hxrt.BoundFunction(func() any {
							return hx_receiver_133.__hx_this.next()
						}, func() any {
							return hx_receiver_133.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "current":
					hx_structural_source_129.current = func(hx_value_130 any) int {
						if hx_value_130 == nil {
							var hx_zero_131 int
							return hx_zero_131
						}
						return hx_value_130.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *StringBuf:
		return func(hx_structural_source_134 *StringBuf) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_134, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_134.b, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_134.b = func(hx_value_135 any) *string {
						if hx_value_135 == nil {
							var hx_zero_136 *string
							return hx_zero_136
						}
						return hx_value_135.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_137 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_137, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_137.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_138 := hx_structural_source_137
						return hxrt.BoundFunction(func() any {
							return hx_receiver_138.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_138.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_137.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_137.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_137.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_137.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_141 := hx_structural_source_137
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_141.__hx_this.get_message()
						}, func() any {
							return hx_receiver_141.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_142 := hx_structural_source_137
						return hxrt.BoundFunction(func() any {
							return hx_receiver_142.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_142.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_143 := hx_structural_source_137
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_143.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_143.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_144 := hx_structural_source_137
						return hxrt.BoundFunction(func() any {
							return hx_receiver_144.__hx_this.get_native()
						}, func() any {
							return hx_receiver_144.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_145 := hx_structural_source_137
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_145.__hx_this.toString()
						}, func() any {
							return hx_receiver_145.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_146 := hx_structural_source_137
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_146.__hx_this.details()
						}, func() any {
							return hx_receiver_146.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_137.savedMessage = func(hx_value_139 any) *string {
						if hx_value_139 == nil {
							var hx_zero_140 *string
							return hx_zero_140
						}
						return hx_value_139.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_137.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_137.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_137.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_147 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_147, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_147.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_147.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_152 := hx_structural_source_147
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_152.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_152.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_153 := hx_structural_source_147
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_153.__hx_this.next()
						}, func() any {
							return hx_receiver_153.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_147.offset = func(hx_value_148 any) int {
						if hx_value_148 == nil {
							var hx_zero_149 int
							return hx_zero_149
						}
						return hx_value_148.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_147.s = func(hx_value_150 any) *string {
						if hx_value_150 == nil {
							var hx_zero_151 *string
							return hx_zero_151
						}
						return hx_value_150.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_154 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_154, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_154.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_154.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_159 := hx_structural_source_154
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_159.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_159.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_160 := hx_structural_source_154
						return hxrt.BoundFunction(func() int {
							return hx_receiver_160.__hx_this.next()
						}, func() any {
							return hx_receiver_160.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_154.offset = func(hx_value_155 any) int {
						if hx_value_155 == nil {
							var hx_zero_156 int
							return hx_zero_156
						}
						return hx_value_155.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_154.s = func(hx_value_157 any) *string {
						if hx_value_157 == nil {
							var hx_zero_158 *string
							return hx_zero_158
						}
						return hx_value_157.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_161 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_161, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_161.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_161.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_161.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_161.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_161.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_170 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_170.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_170.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_171 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_171.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_171.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_172 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len_local int) {
							hx_receiver_172.__hx_this.blit(pos, src, srcpos, len_local)
						}, func() any {
							return hx_receiver_172.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_173 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int, len_local int, value int) {
							hx_receiver_173.__hx_this.fill(pos, len_local, value)
						}, func() any {
							return hx_receiver_173.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_174 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int, len_local int) *haxe__io__Bytes {
							return hx_receiver_174.__hx_this.sub(pos, len_local)
						}, func() any {
							return hx_receiver_174.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_175 := hx_structural_source_161
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_175.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_175.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_176 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_176.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_176.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_177 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_177.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_177.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_178 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_178.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_178.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_179 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_179.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_179.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_180 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_180.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_180.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_181 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_181.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_181.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_182 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_182.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_182.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_183 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_183.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_183.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_184 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_184.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_184.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_185 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_185.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_185.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_186 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int, len_local int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_186.__hx_this.getString(pos, len_local, encoding)
						}, func() any {
							return hx_receiver_186.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_187 := hx_structural_source_161
						return hxrt.BoundFunction(func(pos int, len_local int) *string {
							return hx_receiver_187.__hx_this.readString(pos, len_local)
						}, func() any {
							return hx_receiver_187.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_188 := hx_structural_source_161
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_188.__hx_this.toString()
						}, func() any {
							return hx_receiver_188.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_189 := hx_structural_source_161
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_189.__hx_this.toHex()
						}, func() any {
							return hx_receiver_189.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_190 := hx_structural_source_161
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_190.__hx_this.getData()
						}, func() any {
							return hx_receiver_190.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_191 := hx_structural_source_161
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_191.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_191.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_161.b = func(hx_value_162 any) []int {
						if hx_value_162 == nil {
							var hx_zero_163 []int
							return hx_zero_163
						}
						return hx_value_162.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_161.__hx_raw = func(hx_value_164 any) *hxrt.ByteView {
						if hx_value_164 == nil {
							var hx_zero_165 *hxrt.ByteView
							return hx_zero_165
						}
						return hx_value_164.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_161.__hx_rawValid = func(hx_value_166 any) bool {
						if hx_value_166 == nil {
							var hx_zero_167 bool
							return hx_zero_167
						}
						return hx_value_166.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_161.__hx_dataExposed = func(hx_value_168 any) bool {
						if hx_value_168 == nil {
							var hx_zero_169 bool
							return hx_zero_169
						}
						return hx_value_168.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *Xml:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__ArrayIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__rtti__XmlParser:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__xml__Printer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__xml__XmlParserException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Xml:
		return hxrt__generated_method_field__Xml(value, key)
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_method_field__haxe__ds__StringMap(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_method_field__haxe__io__Bytes(value, key)
	case *haxe__iterators__ArrayIterator:
		return hxrt__generated_method_field__haxe__iterators__ArrayIterator(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_method_field__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value, key)
	case *haxe__rtti__XmlParser:
		return hxrt__generated_method_field__haxe__rtti__XmlParser(value, key)
	case *haxe__xml__Printer:
		return hxrt__generated_method_field__haxe__xml__Printer(value, key)
	case *haxe__xml__XmlParserException:
		return hxrt__generated_method_field__haxe__xml__XmlParserException(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Xml(value *Xml, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "addChild":
		return func() *hxrt.Function[func(*Xml)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) {
				hx_receiver_1.__hx_this.addChild(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "addChild")
		}()
	case "attributes":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.attributes()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "attributes")
		}()
	case "elements":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.elements()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "elements")
		}()
	case "elementsNamed":
		return func() *hxrt.Function[func(*string) hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(name *string) hxrt.Anonymous {
				return hx_receiver_1.__hx_this.elementsNamed(name)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "elementsNamed")
		}()
	case "ensureElementType":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.ensureElementType()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "ensureElementType")
		}()
	case "exists":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(att *string) bool {
				return hx_receiver_1.__hx_this.exists(att)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "exists")
		}()
	case "firstChild":
		return func() *hxrt.Function[func() *Xml] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *Xml {
				return hx_receiver_1.__hx_this.firstChild()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "firstChild")
		}()
	case "firstElement":
		return func() *hxrt.Function[func() *Xml] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *Xml {
				return hx_receiver_1.__hx_this.firstElement()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "firstElement")
		}()
	case "get":
		return func() *hxrt.Function[func(*string) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(att *string) *string {
				return hx_receiver_1.__hx_this.get(att)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get")
		}()
	case "get_nodeName":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.get_nodeName()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_nodeName")
		}()
	case "get_nodeValue":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.get_nodeValue()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_nodeValue")
		}()
	case "insertChild":
		return func() *hxrt.Function[func(*Xml, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml, pos int) {
				hx_receiver_1.__hx_this.insertChild(x, pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "insertChild")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.iterator()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "iterator")
		}()
	case "remove":
		return func() *hxrt.Function[func(*string)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(att *string) {
				hx_receiver_1.__hx_this.remove(att)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "remove")
		}()
	case "removeChild":
		return func() *hxrt.Function[func(*Xml) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) bool {
				return hx_receiver_1.__hx_this.removeChild(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "removeChild")
		}()
	case "set":
		return func() *hxrt.Function[func(*string, *string)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(att *string, value *string) {
				hx_receiver_1.__hx_this.set(att, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set")
		}()
	case "set_nodeName":
		return func() *hxrt.Function[func(*string) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(v *string) *string {
				return hx_receiver_1.__hx_this.set_nodeName(v)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set_nodeName")
		}()
	case "set_nodeValue":
		return func() *hxrt.Function[func(*string) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(v *string) *string {
				return hx_receiver_1.__hx_this.set_nodeValue(v)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set_nodeValue")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.details()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.get_message()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_native()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_1.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.clear()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "clear")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__StringMap] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *haxe__ds__StringMap {
				return hx_receiver_1.__hx_this.copy()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_1.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key *string) bool {
				return hx_receiver_1.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_1.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(*string) any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key *string) any {
				return hx_receiver_1.__hx_this.get(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_1.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.iterator()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.keys()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "keys")
		}()
	case "remove":
		return func() *hxrt.Function[func(*string) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key *string) bool {
				return hx_receiver_1.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_1.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "removeIMap")
		}()
	case "set":
		return func() *hxrt.Function[func(*string, any)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key *string, value any) {
				hx_receiver_1.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_1.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setIMap")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_nativeView":
		return func() *hxrt.Function[func() *hxrt.ByteView] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *hxrt.ByteView {
				return hx_receiver_1.__hx_this.__hx_nativeView()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "__hx_nativeView")
		}()
	case "blit":
		return func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len_local int) {
				hx_receiver_1.__hx_this.blit(pos, src, srcpos, len_local)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "blit")
		}()
	case "compare":
		return func() *hxrt.Function[func(*haxe__io__Bytes) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
				return hx_receiver_1.__hx_this.compare(other)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "compare")
		}()
	case "fill":
		return func() *hxrt.Function[func(int, int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int, value int) {
				hx_receiver_1.__hx_this.fill(pos, len_local, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "fill")
		}()
	case "get":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.get(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get")
		}()
	case "getData":
		return func() *hxrt.Function[func() []int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() []int {
				return hx_receiver_1.__hx_this.getData()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getData")
		}()
	case "getDouble":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_1.__hx_this.getDouble(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getDouble")
		}()
	case "getFloat":
		return func() *hxrt.Function[func(int) float64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) float64 {
				return hx_receiver_1.__hx_this.getFloat(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getFloat")
		}()
	case "getInt32":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.getInt32(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getInt32")
		}()
	case "getInt64":
		return func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
				return hx_receiver_1.__hx_this.getInt64(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getInt64")
		}()
	case "getString":
		return func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int, encoding *haxe__io__Encoding) *string {
				return hx_receiver_1.__hx_this.getString(pos, len_local, encoding)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getString")
		}()
	case "getUInt16":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int) int {
				return hx_receiver_1.__hx_this.getUInt16(pos)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "getUInt16")
		}()
	case "readString":
		return func() *hxrt.Function[func(int, int) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int) *string {
				return hx_receiver_1.__hx_this.readString(pos, len_local)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "readString")
		}()
	case "set":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.set(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "set")
		}()
	case "setDouble":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_1.__hx_this.setDouble(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setDouble")
		}()
	case "setFloat":
		return func() *hxrt.Function[func(int, float64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value float64) {
				hx_receiver_1.__hx_this.setFloat(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setFloat")
		}()
	case "setInt32":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.setInt32(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setInt32")
		}()
	case "setInt64":
		return func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
				hx_receiver_1.__hx_this.setInt64(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setInt64")
		}()
	case "setUInt16":
		return func() *hxrt.Function[func(int, int)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, value int) {
				hx_receiver_1.__hx_this.setUInt16(pos, value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "setUInt16")
		}()
	case "sub":
		return func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(pos int, len_local int) *haxe__io__Bytes {
				return hx_receiver_1.__hx_this.sub(pos, len_local)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "sub")
		}()
	case "toHex":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toHex()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toHex")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__ArrayIterator(value *haxe__iterators__ArrayIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_1.__hx_this.next()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__rtti__XmlParser(value *haxe__rtti__XmlParser, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "defplat":
		return func() *hxrt.Function[func() *hxrt.Array] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *hxrt.Array {
				return hx_receiver_1.__hx_this.defplat()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "defplat")
		}()
	case "elementName":
		return func() *hxrt.Function[func(*Xml) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) *string {
				return hx_receiver_1.__hx_this.elementName(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "elementName")
		}()
	case "findSeparator":
		return func() *hxrt.Function[func(*string, *string, int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value *string, separator *string, start int) int {
				return hx_receiver_1.__hx_this.findSeparator(value, separator, start)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "findSeparator")
		}()
	case "hasNamedElement":
		return func() *hxrt.Function[func(*Xml, *string) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml, name *string) bool {
				return hx_receiver_1.__hx_this.hasNamedElement(x, name)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNamedElement")
		}()
	case "innerData":
		return func() *hxrt.Function[func(*Xml) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) *string {
				return hx_receiver_1.__hx_this.innerData(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "innerData")
		}()
	case "innerHTML":
		return func() *hxrt.Function[func(*Xml) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) *string {
				return hx_receiver_1.__hx_this.innerHTML(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "innerHTML")
		}()
	case "joinStringArray":
		return func() *hxrt.Function[func(*hxrt.Array, *string) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(values *hxrt.Array, separator *string) *string {
				return hx_receiver_1.__hx_this.joinStringArray(values, separator)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "joinStringArray")
		}()
	case "merge":
		return func() *hxrt.Function[func(*haxe__rtti__TypeTree)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(t *haxe__rtti__TypeTree) {
				hx_receiver_1.__hx_this.merge(t)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "merge")
		}()
	case "mergeAbstracts":
		return func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(a hxrt.Anonymous, a2 hxrt.Anonymous) bool {
				return hx_receiver_1.__hx_this.mergeAbstracts(a, a2)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "mergeAbstracts")
		}()
	case "mergeClasses":
		return func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(c hxrt.Anonymous, c2 hxrt.Anonymous) bool {
				return hx_receiver_1.__hx_this.mergeClasses(c, c2)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "mergeClasses")
		}()
	case "mergeDoc":
		return func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(f1 hxrt.Anonymous, f2 hxrt.Anonymous) bool {
				return hx_receiver_1.__hx_this.mergeDoc(f1, f2)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "mergeDoc")
		}()
	case "mergeEnums":
		return func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(e hxrt.Anonymous, e2 hxrt.Anonymous) bool {
				return hx_receiver_1.__hx_this.mergeEnums(e, e2)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "mergeEnums")
		}()
	case "mergeFields":
		return func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(f hxrt.Anonymous, f2 hxrt.Anonymous) bool {
				return hx_receiver_1.__hx_this.mergeFields(f, f2)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "mergeFields")
		}()
	case "mergeRights":
		return func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(f1 hxrt.Anonymous, f2 hxrt.Anonymous) bool {
				return hx_receiver_1.__hx_this.mergeRights(f1, f2)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "mergeRights")
		}()
	case "mergeTypedefs":
		return func() *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(t hxrt.Anonymous, t2 hxrt.Anonymous) bool {
				return hx_receiver_1.__hx_this.mergeTypedefs(t, t2)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "mergeTypedefs")
		}()
	case "mkPath":
		return func() *hxrt.Function[func(*string) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(p *string) *string {
				return hx_receiver_1.__hx_this.mkPath(p)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "mkPath")
		}()
	case "mkRights":
		return func() *hxrt.Function[func(*string) *haxe__rtti__Rights] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(r *string) *haxe__rtti__Rights {
				return hx_receiver_1.__hx_this.mkRights(r)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "mkRights")
		}()
	case "mkTypeParams":
		return func() *hxrt.Function[func(*string) *hxrt.Array] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(p *string) *hxrt.Array {
				return hx_receiver_1.__hx_this.mkTypeParams(p)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "mkTypeParams")
		}()
	case "nodeDisplayName":
		return func() *hxrt.Function[func(*Xml) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) *string {
				return hx_receiver_1.__hx_this.nodeDisplayName(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "nodeDisplayName")
		}()
	case "parseIntString":
		return func() *hxrt.Function[func(*string) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value *string) int {
				return hx_receiver_1.__hx_this.parseIntString(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "parseIntString")
		}()
	case "process":
		return func() *hxrt.Function[func(*Xml, *string)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml, platform *string) {
				hx_receiver_1.__hx_this.process(x, platform)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "process")
		}()
	case "processElement":
		return func() *hxrt.Function[func(*Xml) *haxe__rtti__TypeTree] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) *haxe__rtti__TypeTree {
				return hx_receiver_1.__hx_this.processElement(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "processElement")
		}()
	case "requireAttr":
		return func() *hxrt.Function[func(*Xml, *string) *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml, name *string) *string {
				return hx_receiver_1.__hx_this.requireAttr(x, name)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "requireAttr")
		}()
	case "requireFirstElement":
		return func() *hxrt.Function[func(*Xml) *Xml] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) *Xml {
				return hx_receiver_1.__hx_this.requireFirstElement(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "requireFirstElement")
		}()
	case "requireNamedElement":
		return func() *hxrt.Function[func(*Xml, *string) *Xml] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml, name *string) *Xml {
				return hx_receiver_1.__hx_this.requireNamedElement(x, name)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "requireNamedElement")
		}()
	case "sort":
		return func() *hxrt.Function[func(*hxrt.Array)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(l *hxrt.Array) {
				hx_receiver_1.__hx_this.sort(l)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "sort")
		}()
	case "sortFields":
		return func() *hxrt.Function[func(*hxrt.Array)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(a *hxrt.Array) {
				hx_receiver_1.__hx_this.sortFields(a)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "sortFields")
		}()
	case "splitString":
		return func() *hxrt.Function[func(*string, *string) *hxrt.Array] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value *string, separator *string) *hxrt.Array {
				return hx_receiver_1.__hx_this.splitString(value, separator)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "splitString")
		}()
	case "xabstract":
		return func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
				return hx_receiver_1.__hx_this.xabstract(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xabstract")
		}()
	case "xclass":
		return func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
				return hx_receiver_1.__hx_this.xclass(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xclass")
		}()
	case "xclassfield":
		return func() *hxrt.Function[func(*Xml, any) hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml, defPublic any) hxrt.Anonymous {
				if defPublic == nil {
					defPublic = false
				}
				return hx_receiver_1.__hx_this.xclassfield(x, func(hx_value_2 any) bool {
					if hx_value_2 == nil {
						var hx_zero_3 bool
						return hx_zero_3
					}
					return hx_value_2.(bool)
				}(defPublic))
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xclassfield")
		}()
	case "xenum":
		return func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
				return hx_receiver_1.__hx_this.xenum(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xenum")
		}()
	case "xenumfield":
		return func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
				return hx_receiver_1.__hx_this.xenumfield(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xenumfield")
		}()
	case "xmeta":
		return func() *hxrt.Function[func(*Xml) *hxrt.Array] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) *hxrt.Array {
				return hx_receiver_1.__hx_this.xmeta(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xmeta")
		}()
	case "xoverloads":
		return func() *hxrt.Function[func(*Xml) *hxrt.Array] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) *hxrt.Array {
				return hx_receiver_1.__hx_this.xoverloads(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xoverloads")
		}()
	case "xpath":
		return func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
				return hx_receiver_1.__hx_this.xpath(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xpath")
		}()
	case "xroot":
		return func() *hxrt.Function[func(*Xml)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) {
				hx_receiver_1.__hx_this.xroot(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xroot")
		}()
	case "xtype":
		return func() *hxrt.Function[func(*Xml) *haxe__rtti__CType] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) *haxe__rtti__CType {
				return hx_receiver_1.__hx_this.xtype(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xtype")
		}()
	case "xtypedef":
		return func() *hxrt.Function[func(*Xml) hxrt.Anonymous] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) hxrt.Anonymous {
				return hx_receiver_1.__hx_this.xtypedef(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xtypedef")
		}()
	case "xtypeparams":
		return func() *hxrt.Function[func(*Xml) *hxrt.Array] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(x *Xml) *hxrt.Array {
				return hx_receiver_1.__hx_this.xtypeparams(x)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "xtypeparams")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__xml__Printer(value *haxe__xml__Printer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasChildren":
		return func() *hxrt.Function[func(*Xml) bool] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value *Xml) bool {
				return hx_receiver_1.__hx_this.hasChildren(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasChildren")
		}()
	case "newline":
		return func() *hxrt.Function[func()] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_1.__hx_this.newline()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "newline")
		}()
	case "write":
		return func() *hxrt.Function[func(*string)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(input *string) {
				hx_receiver_1.__hx_this.write(input)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "write")
		}()
	case "writeNode":
		return func() *hxrt.Function[func(*Xml, *string)] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value *Xml, tabs *string) {
				hx_receiver_1.__hx_this.writeNode(value, tabs)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "writeNode")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__xml__XmlParserException(value *haxe__xml__XmlParserException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	}
	return nil
}

type Type struct {
}

func hxrt_typeCreateEmpty__StringBuf() *StringBuf {
	instance := &StringBuf{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__Xml() *Xml {
	instance := &Xml{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__Exception() *haxe__Exception {
	instance := &haxe__Exception{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ValueException() *haxe__ValueException {
	instance := &haxe__ValueException{}
	instance.haxe__Exception = &haxe__Exception{}
	instance.haxe__Exception.__hx_this = instance
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe___Int64_____Int64() *haxe___Int64_____Int64 {
	instance := &haxe___Int64_____Int64{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ds__StringMap() *haxe__ds__StringMap {
	instance := &haxe__ds__StringMap{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__io__Bytes() *haxe__io__Bytes {
	instance := &haxe__io__Bytes{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__iterators__ArrayIterator() *haxe__iterators__ArrayIterator {
	instance := &haxe__iterators__ArrayIterator{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__iterators__MapKeyValueIterator() *haxe__iterators__MapKeyValueIterator {
	instance := &haxe__iterators__MapKeyValueIterator{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__iterators__StringIterator() *haxe__iterators__StringIterator {
	instance := &haxe__iterators__StringIterator{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__iterators__StringKeyValueIterator() *haxe__iterators__StringKeyValueIterator {
	instance := &haxe__iterators__StringKeyValueIterator{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__rtti__XmlParser() *haxe__rtti__XmlParser {
	instance := &haxe__rtti__XmlParser{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__xml__Printer() *haxe__xml__Printer {
	instance := &haxe__xml__Printer{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__xml__XmlParserException() *haxe__xml__XmlParserException {
	instance := &haxe__xml__XmlParserException{}
	instance.__hx_this = instance
	return instance
}

type ValueType struct {
	tag    int
	params []any
}

func (self *ValueType) HxEnumValue() (string, []any) {
	return [9]string{"TNull", "TInt", "TFloat", "TBool", "TObject", "TFunction", "TClass", "TEnum", "TUnknown"}[self.tag], self.params
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
	case "Demo":
		return nil, false
	case "Main":
		return nil, false
	case "Reflect":
		return nil, false
	case "Std":
		return nil, false
	case "StringBuf":
		return hxrt_typeCallAny(New_StringBuf, args)
	case "StringTools":
		return nil, false
	case "Xml":
		return hxrt_typeCallAny(New_Xml, args)
	case "_Xml.XmlType_Impl_":
		return nil, false
	case "haxe.Exception":
		return hxrt_typeCallAny(New_haxe__Exception, args)
	case "haxe.Int64Helper":
		return nil, false
	case "haxe.ValueException":
		return hxrt_typeCallAny(New_haxe__ValueException, args)
	case "haxe._CallStack.CallStack_Impl_":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.ds.ArraySort":
		return nil, false
	case "haxe.ds.StringMap":
		return hxrt_typeCallAny(New_haxe__ds__StringMap, args)
	case "haxe.io.Bytes":
		return hxrt_typeCallAny(New_haxe__io__Bytes, args)
	case "haxe.io.FPHelper":
		return nil, false
	case "haxe.iterators.ArrayIterator":
		return hxrt_typeCallAny(New_haxe__iterators__ArrayIterator, args)
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__MapKeyValueIterator, args)
	case "haxe.iterators.StringIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringIterator, args)
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__StringKeyValueIterator, args)
	case "haxe.rtti.CTypeTools":
		return nil, false
	case "haxe.rtti.Rtti":
		return nil, false
	case "haxe.rtti.TypeApi":
		return nil, false
	case "haxe.rtti.XmlParser":
		return hxrt_typeCallAny(New_haxe__rtti__XmlParser, args)
	case "haxe.xml.Parser":
		return nil, false
	case "haxe.xml.Printer":
		return hxrt_typeCallAny(New_haxe__xml__Printer, args)
	case "haxe.xml.XmlParserException":
		return hxrt_typeCallAny(New_haxe__xml__XmlParserException, args)
	case "haxe.xml._Parser.S_Impl_":
		return nil, false
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "StringBuf":
		return hxrt_typeCreateEmpty__StringBuf(), true
	case "Xml":
		return hxrt_typeCreateEmpty__Xml(), true
	case "haxe.Exception":
		return hxrt_typeCreateEmpty__haxe__Exception(), true
	case "haxe.ValueException":
		return hxrt_typeCreateEmpty__haxe__ValueException(), true
	case "haxe._Int64.___Int64":
		return hxrt_typeCreateEmpty__haxe___Int64_____Int64(), true
	case "haxe.ds.StringMap":
		return hxrt_typeCreateEmpty__haxe__ds__StringMap(), true
	case "haxe.io.Bytes":
		return hxrt_typeCreateEmpty__haxe__io__Bytes(), true
	case "haxe.iterators.ArrayIterator":
		return hxrt_typeCreateEmpty__haxe__iterators__ArrayIterator(), true
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCreateEmpty__haxe__iterators__MapKeyValueIterator(), true
	case "haxe.iterators.StringIterator":
		return hxrt_typeCreateEmpty__haxe__iterators__StringIterator(), true
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt_typeCreateEmpty__haxe__iterators__StringKeyValueIterator(), true
	case "haxe.rtti.XmlParser":
		return hxrt_typeCreateEmpty__haxe__rtti__XmlParser(), true
	case "haxe.xml.Printer":
		return hxrt_typeCreateEmpty__haxe__xml__Printer(), true
	case "haxe.xml.XmlParserException":
		return hxrt_typeCreateEmpty__haxe__xml__XmlParserException(), true
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
	case "haxe.StackItem":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__StackItem_CFunction, true
			case 1:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_Module, args)
			case 2:
				if len(args) != 4 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_FilePos, args)
			case 3:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_Method, args)
			case 4:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_LocalFunction, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "CFunction":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__StackItem_CFunction, true
		case "Module":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_Module, args)
		case "FilePos":
			if len(args) != 4 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_FilePos, args)
		case "Method":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_Method, args)
		case "LocalFunction":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_LocalFunction, args)
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
	case "haxe.rtti.CType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__CType_CUnknown, true
			case 1:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CEnum, args)
			case 2:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CClass, args)
			case 3:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CTypedef, args)
			case 4:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CFunction, args)
			case 5:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CAnonymous, args)
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CDynamic, args)
			case 7:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__CType_CAbstract, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "CUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__CType_CUnknown, true
		case "CEnum":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CEnum, args)
		case "CClass":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CClass, args)
		case "CTypedef":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CTypedef, args)
		case "CFunction":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CFunction, args)
		case "CAnonymous":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CAnonymous, args)
		case "CDynamic":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CDynamic, args)
		case "CAbstract":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__CType_CAbstract, args)
		default:
			return nil, false
		}
	case "haxe.rtti.Rights":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__Rights_RNormal, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__Rights_RNo, true
			case 2:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__Rights_RCall, args)
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__Rights_RMethod, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__Rights_RDynamic, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__rtti__Rights_RInline, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "RNormal":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__Rights_RNormal, true
		case "RNo":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__Rights_RNo, true
		case "RCall":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__Rights_RCall, args)
		case "RMethod":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__Rights_RMethod, true
		case "RDynamic":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__Rights_RDynamic, true
		case "RInline":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__rtti__Rights_RInline, true
		default:
			return nil, false
		}
	case "haxe.rtti.TypeTree":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 3 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__TypeTree_TPackage, args)
			case 1:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__TypeTree_TClassdecl, args)
			case 2:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__TypeTree_TEnumdecl, args)
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__TypeTree_TTypedecl, args)
			case 4:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__rtti__TypeTree_TAbstractdecl, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TPackage":
			if len(args) != 3 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__TypeTree_TPackage, args)
		case "TClassdecl":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__TypeTree_TClassdecl, args)
		case "TEnumdecl":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__TypeTree_TEnumdecl, args)
		case "TTypedecl":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__TypeTree_TTypedecl, args)
		case "TAbstractdecl":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__rtti__TypeTree_TAbstractdecl, args)
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func Type_getClass(o any) any {
	o = hxrt.ObjectSource(o)
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
	case *StringBuf:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("StringBuf")}
	case *Xml:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Xml")}
	case *haxe__Exception:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case *haxe__ValueException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ValueException")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	case *haxe__ds__StringMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.StringMap")}
	case *haxe__io__Bytes:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Bytes")}
	case *haxe__iterators__ArrayIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.ArrayIterator")}
	case *haxe__iterators__MapKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.MapKeyValueIterator")}
	case *haxe__iterators__StringIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.StringIterator")}
	case *haxe__iterators__StringKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.StringKeyValueIterator")}
	case *haxe__rtti__XmlParser:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.rtti.XmlParser")}
	case *haxe__xml__Printer:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.xml.Printer")}
	case *haxe__xml__XmlParserException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.xml.XmlParserException")}
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
	case *haxe__StackItem:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.StackItem")}
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
	case *haxe__rtti__CType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.rtti.CType")}
	case *haxe__rtti__Rights:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.rtti.Rights")}
	case *haxe__rtti__TypeTree:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.rtti.TypeTree")}
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
	case "Demo":
		return nil
	case "Main":
		return nil
	case "Reflect":
		return nil
	case "Std":
		return nil
	case "StringBuf":
		return nil
	case "StringTools":
		return nil
	case "Xml":
		return nil
	case "_Xml.XmlType_Impl_":
		return nil
	case "haxe.Exception":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe.ValueException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "haxe._CallStack.CallStack_Impl_":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.ds.ArraySort":
		return nil
	case "haxe.ds.StringMap":
		return nil
	case "haxe.io.Bytes":
		return nil
	case "haxe.io.FPHelper":
		return nil
	case "haxe.iterators.ArrayIterator":
		return nil
	case "haxe.iterators.MapKeyValueIterator":
		return nil
	case "haxe.iterators.StringIterator":
		return nil
	case "haxe.iterators.StringKeyValueIterator":
		return nil
	case "haxe.rtti.CTypeTools":
		return nil
	case "haxe.rtti.Rtti":
		return nil
	case "haxe.rtti.TypeApi":
		return nil
	case "haxe.rtti.XmlParser":
		return nil
	case "haxe.xml.Parser":
		return nil
	case "haxe.xml.Printer":
		return nil
	case "haxe.xml.XmlParserException":
		return nil
	case "haxe.xml._Parser.S_Impl_":
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
	case "Demo":
		return hxrt.NewArray(hxrt.StringFromLiteral("__rtti"), hxrt.StringFromLiteral("field"))
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("main"))
	case "Reflect":
		return hxrt.NewArray(hxrt.StringFromLiteral("callMethod"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("compareMethods"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("deleteField"), hxrt.StringFromLiteral("field"), hxrt.StringFromLiteral("fields"), hxrt.StringFromLiteral("getProperty"), hxrt.StringFromLiteral("hasField"), hxrt.StringFromLiteral("isEnumValue"), hxrt.StringFromLiteral("isFunction"), hxrt.StringFromLiteral("isObject"), hxrt.StringFromLiteral("makeVarArgs"), hxrt.StringFromLiteral("setField"), hxrt.StringFromLiteral("setProperty"))
	case "Std":
		return hxrt.NewArray(hxrt.StringFromLiteral("digitValue"), hxrt.StringFromLiteral("downcast"), hxrt.StringFromLiteral("instance"), hxrt.StringFromLiteral("int"), hxrt.StringFromLiteral("invalidFloat"), hxrt.StringFromLiteral("is"), hxrt.StringFromLiteral("isDecimalDigit"), hxrt.StringFromLiteral("isSpaceCode"), hxrt.StringFromLiteral("parseFloat"), hxrt.StringFromLiteral("parseInt"), hxrt.StringFromLiteral("random"))
	case "StringBuf":
		return hxrt.NewArray()
	case "StringTools":
		return hxrt.NewArray(hxrt.StringFromLiteral("MAX_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_HIGH_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("MIN_SURROGATE_CODE_POINT"), hxrt.StringFromLiteral("contains"), hxrt.StringFromLiteral("containsImpl"), hxrt.StringFromLiteral("endsWith"), hxrt.StringFromLiteral("endsWithImpl"), hxrt.StringFromLiteral("fastCodeAt"), hxrt.StringFromLiteral("hex"), hxrt.StringFromLiteral("hexDigitValue"), hxrt.StringFromLiteral("htmlEscape"), hxrt.StringFromLiteral("htmlUnescape"), hxrt.StringFromLiteral("isEof"), hxrt.StringFromLiteral("isSpace"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("lpad"), hxrt.StringFromLiteral("ltrim"), hxrt.StringFromLiteral("replace"), hxrt.StringFromLiteral("rpad"), hxrt.StringFromLiteral("rtrim"), hxrt.StringFromLiteral("startsWith"), hxrt.StringFromLiteral("startsWithImpl"), hxrt.StringFromLiteral("trim"), hxrt.StringFromLiteral("unsafeCodeAt"), hxrt.StringFromLiteral("urlDecode"), hxrt.StringFromLiteral("urlEncode"), hxrt.StringFromLiteral("utf16CodePointAt"))
	case "Xml":
		return hxrt.NewArray(hxrt.StringFromLiteral("CData"), hxrt.StringFromLiteral("Comment"), hxrt.StringFromLiteral("DocType"), hxrt.StringFromLiteral("Document"), hxrt.StringFromLiteral("Element"), hxrt.StringFromLiteral("PCData"), hxrt.StringFromLiteral("ProcessingInstruction"), hxrt.StringFromLiteral("createCData"), hxrt.StringFromLiteral("createComment"), hxrt.StringFromLiteral("createDocType"), hxrt.StringFromLiteral("createDocument"), hxrt.StringFromLiteral("createElement"), hxrt.StringFromLiteral("createPCData"), hxrt.StringFromLiteral("createProcessingInstruction"), hxrt.StringFromLiteral("parse"))
	case "_Xml.XmlType_Impl_":
		return hxrt.NewArray(hxrt.StringFromLiteral("CData"), hxrt.StringFromLiteral("Comment"), hxrt.StringFromLiteral("DocType"), hxrt.StringFromLiteral("Document"), hxrt.StringFromLiteral("Element"), hxrt.StringFromLiteral("PCData"), hxrt.StringFromLiteral("ProcessingInstruction"), hxrt.StringFromLiteral("toString"))
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("catchValue"), hxrt.StringFromLiteral("caught"), hxrt.StringFromLiteral("thrown"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.ValueException":
		return hxrt.NewArray()
	case "haxe._CallStack.CallStack_Impl_":
		return hxrt.NewArray(hxrt.StringFromLiteral("asArray"), hxrt.StringFromLiteral("callStack"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("equalItems"), hxrt.StringFromLiteral("exceptionStack"), hxrt.StringFromLiteral("exceptionToString"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("get_length"), hxrt.StringFromLiteral("itemToString"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("subtract"), hxrt.StringFromLiteral("toString"))
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.ds.ArraySort":
		return hxrt.NewArray(hxrt.StringFromLiteral("doMerge"), hxrt.StringFromLiteral("gcd"), hxrt.StringFromLiteral("lower"), hxrt.StringFromLiteral("rec"), hxrt.StringFromLiteral("rotate"), hxrt.StringFromLiteral("sort"), hxrt.StringFromLiteral("swap"), hxrt.StringFromLiteral("upper"))
	case "haxe.ds.StringMap":
		return hxrt.NewArray()
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_fromNativeView"), hxrt.StringFromLiteral("alloc"), hxrt.StringFromLiteral("fastGet"), hxrt.StringFromLiteral("ofData"), hxrt.StringFromLiteral("ofHex"), hxrt.StringFromLiteral("ofString"), hxrt.StringFromLiteral("rawNativeUsesUtf16LE"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray(hxrt.StringFromLiteral("doubleToI64"), hxrt.StringFromLiteral("floatToI32"), hxrt.StringFromLiteral("i32ToFloat"), hxrt.StringFromLiteral("i64ToDouble"))
	case "haxe.iterators.ArrayIterator":
		return hxrt.NewArray()
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray()
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt.NewArray()
	case "haxe.rtti.CTypeTools":
		return hxrt.NewArray(hxrt.StringFromLiteral("classField"), hxrt.StringFromLiteral("functionArgumentName"), hxrt.StringFromLiteral("joinClassFields"), hxrt.StringFromLiteral("joinFunctionArguments"), hxrt.StringFromLiteral("joinStringArray"), hxrt.StringFromLiteral("nameWithParams"), hxrt.StringFromLiteral("toString"))
	case "haxe.rtti.Rtti":
		return hxrt.NewArray(hxrt.StringFromLiteral("getRtti"), hxrt.StringFromLiteral("hasRtti"))
	case "haxe.rtti.TypeApi":
		return hxrt.NewArray(hxrt.StringFromLiteral("constructorEq"), hxrt.StringFromLiteral("fieldEq"), hxrt.StringFromLiteral("isVar"), hxrt.StringFromLiteral("rightsEq"), hxrt.StringFromLiteral("sameClassFields"), hxrt.StringFromLiteral("sameConstructorArguments"), hxrt.StringFromLiteral("sameFunctionArguments"), hxrt.StringFromLiteral("sameTypeParamNames"), hxrt.StringFromLiteral("sameTypes"), hxrt.StringFromLiteral("typeEq"), hxrt.StringFromLiteral("typeInfos"))
	case "haxe.rtti.XmlParser":
		return hxrt.NewArray()
	case "haxe.xml.Parser":
		return hxrt.NewArray(hxrt.StringFromLiteral("doParse"), hxrt.StringFromLiteral("escapes"), hxrt.StringFromLiteral("parse"))
	case "haxe.xml.Printer":
		return hxrt.NewArray(hxrt.StringFromLiteral("print"))
	case "haxe.xml.XmlParserException":
		return hxrt.NewArray()
	case "haxe.xml._Parser.S_Impl_":
		return hxrt.NewArray()
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
	case "Demo":
		return hxrt.NewArray()
	case "Main":
		return hxrt.NewArray()
	case "Reflect":
		return hxrt.NewArray()
	case "Std":
		return hxrt.NewArray()
	case "StringBuf":
		return hxrt.NewArray(hxrt.StringFromLiteral("b"))
	case "StringTools":
		return hxrt.NewArray()
	case "Xml":
		return hxrt.NewArray(hxrt.StringFromLiteral("addChild"), hxrt.StringFromLiteral("attributeMap"), hxrt.StringFromLiteral("attributes"), hxrt.StringFromLiteral("children"), hxrt.StringFromLiteral("elements"), hxrt.StringFromLiteral("elementsNamed"), hxrt.StringFromLiteral("ensureElementType"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("firstChild"), hxrt.StringFromLiteral("firstElement"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("get_nodeName"), hxrt.StringFromLiteral("get_nodeValue"), hxrt.StringFromLiteral("insertChild"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("nodeName"), hxrt.StringFromLiteral("nodeType"), hxrt.StringFromLiteral("nodeValue"), hxrt.StringFromLiteral("parent"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeChild"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("set_nodeName"), hxrt.StringFromLiteral("set_nodeValue"), hxrt.StringFromLiteral("toString"))
	case "_Xml.XmlType_Impl_":
		return hxrt.NewArray()
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.ValueException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"), hxrt.StringFromLiteral("value"))
	case "haxe._CallStack.CallStack_Impl_":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.ds.ArraySort":
		return hxrt.NewArray()
	case "haxe.ds.StringMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_dataExposed"), hxrt.StringFromLiteral("__hx_nativeView"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("blit"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("fill"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getData"), hxrt.StringFromLiteral("getDouble"), hxrt.StringFromLiteral("getFloat"), hxrt.StringFromLiteral("getInt32"), hxrt.StringFromLiteral("getInt64"), hxrt.StringFromLiteral("getString"), hxrt.StringFromLiteral("getUInt16"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setDouble"), hxrt.StringFromLiteral("setFloat"), hxrt.StringFromLiteral("setInt32"), hxrt.StringFromLiteral("setInt64"), hxrt.StringFromLiteral("setUInt16"), hxrt.StringFromLiteral("sub"), hxrt.StringFromLiteral("toHex"), hxrt.StringFromLiteral("toString"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray()
	case "haxe.iterators.ArrayIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("array"), hxrt.StringFromLiteral("current"), hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"))
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("next"))
	case "haxe.iterators.StringIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case "haxe.iterators.StringKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case "haxe.rtti.CTypeTools":
		return hxrt.NewArray()
	case "haxe.rtti.Rtti":
		return hxrt.NewArray()
	case "haxe.rtti.TypeApi":
		return hxrt.NewArray()
	case "haxe.rtti.XmlParser":
		return hxrt.NewArray(hxrt.StringFromLiteral("curplatform"), hxrt.StringFromLiteral("defplat"), hxrt.StringFromLiteral("elementName"), hxrt.StringFromLiteral("findSeparator"), hxrt.StringFromLiteral("hasNamedElement"), hxrt.StringFromLiteral("innerData"), hxrt.StringFromLiteral("innerHTML"), hxrt.StringFromLiteral("joinStringArray"), hxrt.StringFromLiteral("merge"), hxrt.StringFromLiteral("mergeAbstracts"), hxrt.StringFromLiteral("mergeClasses"), hxrt.StringFromLiteral("mergeDoc"), hxrt.StringFromLiteral("mergeEnums"), hxrt.StringFromLiteral("mergeFields"), hxrt.StringFromLiteral("mergeRights"), hxrt.StringFromLiteral("mergeTypedefs"), hxrt.StringFromLiteral("mkPath"), hxrt.StringFromLiteral("mkRights"), hxrt.StringFromLiteral("mkTypeParams"), hxrt.StringFromLiteral("newField"), hxrt.StringFromLiteral("nodeDisplayName"), hxrt.StringFromLiteral("parseIntString"), hxrt.StringFromLiteral("process"), hxrt.StringFromLiteral("processElement"), hxrt.StringFromLiteral("requireAttr"), hxrt.StringFromLiteral("requireFirstElement"), hxrt.StringFromLiteral("requireNamedElement"), hxrt.StringFromLiteral("root"), hxrt.StringFromLiteral("sort"), hxrt.StringFromLiteral("sortFields"), hxrt.StringFromLiteral("splitString"), hxrt.StringFromLiteral("xabstract"), hxrt.StringFromLiteral("xclass"), hxrt.StringFromLiteral("xclassfield"), hxrt.StringFromLiteral("xenum"), hxrt.StringFromLiteral("xenumfield"), hxrt.StringFromLiteral("xmeta"), hxrt.StringFromLiteral("xoverloads"), hxrt.StringFromLiteral("xpath"), hxrt.StringFromLiteral("xroot"), hxrt.StringFromLiteral("xtype"), hxrt.StringFromLiteral("xtypedef"), hxrt.StringFromLiteral("xtypeparams"))
	case "haxe.xml.Parser":
		return hxrt.NewArray()
	case "haxe.xml.Printer":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasChildren"), hxrt.StringFromLiteral("newline"), hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("pretty"), hxrt.StringFromLiteral("write"), hxrt.StringFromLiteral("writeNode"))
	case "haxe.xml.XmlParserException":
		return hxrt.NewArray(hxrt.StringFromLiteral("lineNumber"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("position"), hxrt.StringFromLiteral("positionAtLine"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("xml"))
	case "haxe.xml._Parser.S_Impl_":
		return hxrt.NewArray()
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
	case "Demo":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Reflect":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Std":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringBuf":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringTools":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Xml":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "_Xml.XmlType_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Exception":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ValueException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._CallStack.CallStack_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.ArraySort":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.StringMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Bytes":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.FPHelper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.ArrayIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.MapKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.StringKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.CTypeTools":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.Rtti":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.TypeApi":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.XmlParser":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml.Parser":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml.Printer":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml.XmlParserException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.xml._Parser.S_Impl_":
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
	case "haxe.StackItem":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Encoding":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Error":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.CType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.Rights":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.rtti.TypeTree":
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
	case *haxe__StackItem:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("CFunction")
		case 1:
			return hxrt.StringFromLiteral("Module")
		case 2:
			return hxrt.StringFromLiteral("FilePos")
		case 3:
			return hxrt.StringFromLiteral("Method")
		case 4:
			return hxrt.StringFromLiteral("LocalFunction")
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
	case *haxe__rtti__CType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("CUnknown")
		case 1:
			return hxrt.StringFromLiteral("CEnum")
		case 2:
			return hxrt.StringFromLiteral("CClass")
		case 3:
			return hxrt.StringFromLiteral("CTypedef")
		case 4:
			return hxrt.StringFromLiteral("CFunction")
		case 5:
			return hxrt.StringFromLiteral("CAnonymous")
		case 6:
			return hxrt.StringFromLiteral("CDynamic")
		case 7:
			return hxrt.StringFromLiteral("CAbstract")
		default:
			return nil
		}
	case *haxe__rtti__Rights:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("RNormal")
		case 1:
			return hxrt.StringFromLiteral("RNo")
		case 2:
			return hxrt.StringFromLiteral("RCall")
		case 3:
			return hxrt.StringFromLiteral("RMethod")
		case 4:
			return hxrt.StringFromLiteral("RDynamic")
		case 5:
			return hxrt.StringFromLiteral("RInline")
		default:
			return nil
		}
	case *haxe__rtti__TypeTree:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TPackage")
		case 1:
			return hxrt.StringFromLiteral("TClassdecl")
		case 2:
			return hxrt.StringFromLiteral("TEnumdecl")
		case 3:
			return hxrt.StringFromLiteral("TTypedecl")
		case 4:
			return hxrt.StringFromLiteral("TAbstractdecl")
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
	case *haxe__StackItem:
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
	case *haxe__rtti__CType:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__rtti__Rights:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__rtti__TypeTree:
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
	case "haxe.StackItem":
		return hxrt.NewArray(hxrt.StringFromLiteral("CFunction"), hxrt.StringFromLiteral("Module"), hxrt.StringFromLiteral("FilePos"), hxrt.StringFromLiteral("Method"), hxrt.StringFromLiteral("LocalFunction"))
	case "haxe.io.Encoding":
		return hxrt.NewArray(hxrt.StringFromLiteral("UTF8"), hxrt.StringFromLiteral("RawNative"))
	case "haxe.io.Error":
		return hxrt.NewArray(hxrt.StringFromLiteral("Blocked"), hxrt.StringFromLiteral("Overflow"), hxrt.StringFromLiteral("OutsideBounds"), hxrt.StringFromLiteral("Custom"))
	case "haxe.rtti.CType":
		return hxrt.NewArray(hxrt.StringFromLiteral("CUnknown"), hxrt.StringFromLiteral("CEnum"), hxrt.StringFromLiteral("CClass"), hxrt.StringFromLiteral("CTypedef"), hxrt.StringFromLiteral("CFunction"), hxrt.StringFromLiteral("CAnonymous"), hxrt.StringFromLiteral("CDynamic"), hxrt.StringFromLiteral("CAbstract"))
	case "haxe.rtti.Rights":
		return hxrt.NewArray(hxrt.StringFromLiteral("RNormal"), hxrt.StringFromLiteral("RNo"), hxrt.StringFromLiteral("RCall"), hxrt.StringFromLiteral("RMethod"), hxrt.StringFromLiteral("RDynamic"), hxrt.StringFromLiteral("RInline"))
	case "haxe.rtti.TypeTree":
		return hxrt.NewArray(hxrt.StringFromLiteral("TPackage"), hxrt.StringFromLiteral("TClassdecl"), hxrt.StringFromLiteral("TEnumdecl"), hxrt.StringFromLiteral("TTypedecl"), hxrt.StringFromLiteral("TAbstractdecl"))
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
	case *haxe__StackItem:
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
	case *haxe__rtti__CType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__rtti__Rights:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__rtti__TypeTree:
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
	case "haxe.StackItem":
		return hxrt.NewArray(haxe__StackItem_CFunction)
	case "haxe.io.Encoding":
		return hxrt.NewArray(haxe__io__Encoding_UTF8, haxe__io__Encoding_RawNative)
	case "haxe.io.Error":
		return hxrt.NewArray(haxe__io__Error_Blocked, haxe__io__Error_Overflow, haxe__io__Error_OutsideBounds)
	case "haxe.rtti.CType":
		return hxrt.NewArray(haxe__rtti__CType_CUnknown)
	case "haxe.rtti.Rights":
		return hxrt.NewArray(haxe__rtti__Rights_RNormal, haxe__rtti__Rights_RNo, haxe__rtti__Rights_RMethod, haxe__rtti__Rights_RDynamic, haxe__rtti__Rights_RInline)
	case "haxe.rtti.TypeTree":
		return hxrt.NewArray()
	default:
		return hxrt.NewArray()
	}
}

func Type_typeof(v any) *ValueType {
	v = hxrt.ObjectSource(v)
	if hxrt.ReflectIsFunction(v) {
		return ValueType_TFunction
	}
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
	classValue, ok := value.(*hxrt__TypeClassValue)
	if !ok || classValue == nil {
		return nil, false
	}
	className := *hxrt.StdString(classValue.name)
	switch className {
	case "Demo":
		switch key {
		case "__rtti":
			return Demo___rtti, true
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedField(object any, field *string) any {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *StringBuf:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Xml:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__ArrayIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__rtti__XmlParser:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__xml__Printer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__xml__XmlParserException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *StringBuf:
		return hxrt__generated_field_lookup__StringBuf(value, key)
	case *Xml:
		return hxrt__generated_field_lookup__Xml(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_lookup__haxe__ds__StringMap(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_lookup__haxe__io__Bytes(value, key)
	case *haxe__iterators__ArrayIterator:
		return hxrt__generated_field_lookup__haxe__iterators__ArrayIterator(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value, key)
	case *haxe__rtti__XmlParser:
		return hxrt__generated_field_lookup__haxe__rtti__XmlParser(value, key)
	case *haxe__xml__Printer:
		return hxrt__generated_field_lookup__haxe__xml__Printer(value, key)
	case *haxe__xml__XmlParserException:
		return hxrt__generated_field_lookup__haxe__xml__XmlParserException(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *StringBuf:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Xml:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__ArrayIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__rtti__XmlParser:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__xml__Printer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__xml__XmlParserException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *StringBuf:
		return hxrt__generated_field_has__StringBuf(value, key)
	case *Xml:
		return hxrt__generated_field_has__Xml(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_has__haxe__ds__StringMap(value, key)
	case *haxe__io__Bytes:
		return hxrt__generated_field_has__haxe__io__Bytes(value, key)
	case *haxe__iterators__ArrayIterator:
		return hxrt__generated_field_has__haxe__iterators__ArrayIterator(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value, key)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_has__haxe__iterators__StringIterator(value, key)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value, key)
	case *haxe__rtti__XmlParser:
		return hxrt__generated_field_has__haxe__rtti__XmlParser(value, key)
	case *haxe__xml__Printer:
		return hxrt__generated_field_has__haxe__xml__Printer(value, key)
	case *haxe__xml__XmlParserException:
		return hxrt__generated_field_has__haxe__xml__XmlParserException(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *StringBuf:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Xml:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__ArrayIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__rtti__XmlParser:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__xml__Printer:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__xml__XmlParserException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *StringBuf:
		return hxrt__generated_field_set__StringBuf(value, key, incoming)
	case *Xml:
		return hxrt__generated_field_set__Xml(value, key, incoming)
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe__ds__StringMap:
		return hxrt__generated_field_set__haxe__ds__StringMap(value, key, incoming)
	case *haxe__io__Bytes:
		return hxrt__generated_field_set__haxe__io__Bytes(value, key, incoming)
	case *haxe__iterators__ArrayIterator:
		return hxrt__generated_field_set__haxe__iterators__ArrayIterator(value, key, incoming)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value, key, incoming)
	case *haxe__iterators__StringIterator:
		return hxrt__generated_field_set__haxe__iterators__StringIterator(value, key, incoming)
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value, key, incoming)
	case *haxe__rtti__XmlParser:
		return hxrt__generated_field_set__haxe__rtti__XmlParser(value, key, incoming)
	case *haxe__xml__Printer:
		return hxrt__generated_field_set__haxe__xml__Printer(value, key, incoming)
	case *haxe__xml__XmlParserException:
		return hxrt__generated_field_set__haxe__xml__XmlParserException(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *StringBuf:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Xml:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__StringMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__io__Bytes:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__ArrayIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__StringKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__rtti__XmlParser:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__xml__Printer:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__xml__XmlParserException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *StringBuf:
		return hxrt.NewArray(hxrt.StringFromLiteral("b"))
	case *Xml:
		return hxrt.NewArray(hxrt.StringFromLiteral("nodeType"), hxrt.StringFromLiteral("parent"), hxrt.StringFromLiteral("children"), hxrt.StringFromLiteral("attributeMap"))
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__StringMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__io__Bytes:
		return hxrt.NewArray(hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("__hx_dataExposed"))
	case *haxe__iterators__ArrayIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("array"), hxrt.StringFromLiteral("current"))
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("keys"))
	case *haxe__iterators__StringIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *haxe__iterators__StringKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("offset"), hxrt.StringFromLiteral("s"))
	case *haxe__rtti__XmlParser:
		return hxrt.NewArray(hxrt.StringFromLiteral("root"), hxrt.StringFromLiteral("curplatform"), hxrt.StringFromLiteral("newField"))
	case *haxe__xml__Printer:
		return hxrt.NewArray(hxrt.StringFromLiteral("output"), hxrt.StringFromLiteral("pretty"))
	case *haxe__xml__XmlParserException:
		return hxrt.NewArray(hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("lineNumber"), hxrt.StringFromLiteral("positionAtLine"), hxrt.StringFromLiteral("position"), hxrt.StringFromLiteral("xml"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__StringBuf(value *StringBuf, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "b":
		return hxrt.ReadStoredField(value.b)
	}
	return nil
}

func hxrt__generated_field_has__StringBuf(value *StringBuf, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		return true
	}
	return false
}

func hxrt__generated_field_set__StringBuf(value *StringBuf, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "b":
		if incoming == nil {
			var zero *string
			value.b = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.b = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__Xml(value *Xml, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "attributeMap":
		return hxrt.ReadStoredField(value.attributeMap)
	case "children":
		return hxrt.ReadStoredField(value.children)
	case "nodeType":
		return hxrt.ReadStoredField(value.nodeType)
	case "parent":
		return hxrt.ReadStoredField(value.parent)
	}
	return nil
}

func hxrt__generated_field_has__Xml(value *Xml, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "attributeMap":
		return true
	case "children":
		return true
	case "nodeType":
		return true
	case "parent":
		return true
	}
	return false
}

func hxrt__generated_field_set__Xml(value *Xml, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "attributeMap":
		if incoming == nil {
			var zero *haxe__ds__StringMap
			value.attributeMap = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__ds__StringMap:
			value.attributeMap = typed
			return true
		default:
			return false
		}
	case "children":
		if incoming == nil {
			var zero *hxrt.Array
			value.children = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.children = typed
			return true
		default:
			return false
		}
	case "nodeType":
		if incoming == nil {
			var zero int
			value.nodeType = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.nodeType = typed
			return true
		default:
			return false
		}
	case "parent":
		if incoming == nil {
			var zero *Xml
			value.parent = zero
			return true
		}
		switch typed := incoming.(type) {
		case *Xml:
			value.parent = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "savedMessage":
		return hxrt.ReadStoredField(value.savedMessage)
	case "savedNative":
		return hxrt.ReadStoredField(value.savedNative)
	case "savedPrevious":
		return hxrt.ReadStoredField(value.savedPrevious)
	case "savedStack":
		return hxrt.ReadStoredField(value.savedStack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Exception(value *haxe__Exception, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		return true
	case "savedNative":
		return true
	case "savedPrevious":
		return true
	case "savedStack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Exception(value *haxe__Exception, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		if incoming == nil {
			var zero *string
			value.savedMessage = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.savedMessage = typed
			return true
		default:
			return false
		}
	case "savedNative":
		if incoming == nil {
			var zero any
			value.savedNative = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedNative = typed
			return true
		default:
			return false
		}
	case "savedPrevious":
		if incoming == nil {
			var zero *haxe__Exception
			value.savedPrevious = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__Exception:
			value.savedPrevious = typed
			return true
		default:
			return false
		}
	case "savedStack":
		if incoming == nil {
			var zero any
			value.savedStack = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedStack = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__ValueException(value *haxe__ValueException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__ValueException(value *haxe__ValueException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		if incoming == nil {
			var zero any
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.value = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return hxrt.ReadStoredField(value.high)
	case "low":
		return hxrt.ReadStoredField(value.low)
	}
	return nil
}

func hxrt__generated_field_has__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		return true
	case "low":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		if incoming == nil {
			var zero int
			value.high = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.high = typed
			return true
		default:
			return false
		}
	case "low":
		if incoming == nil {
			var zero int
			value.low = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.low = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "h":
		return hxrt.ReadStoredField(value.h)
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__StringMap(value *haxe__ds__StringMap, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__StringMap(value *haxe__ds__StringMap, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		if incoming == nil {
			var zero *hxrt.StringMapCell
			value.h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.StringMapCell:
			value.h = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__io__Bytes(value *haxe__io__Bytes, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "__hx_dataExposed":
		return hxrt.ReadStoredField(value.__hx_dataExposed)
	case "__hx_raw":
		return hxrt.ReadStoredField(value.__hx_raw)
	case "__hx_rawValid":
		return hxrt.ReadStoredField(value.__hx_rawValid)
	case "b":
		return hxrt.ReadStoredField(value.b)
	case "length":
		return hxrt.ReadStoredField(value.length)
	}
	return nil
}

func hxrt__generated_field_has__haxe__io__Bytes(value *haxe__io__Bytes, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__hx_dataExposed":
		return true
	case "__hx_raw":
		return true
	case "__hx_rawValid":
		return true
	case "b":
		return true
	case "length":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__io__Bytes(value *haxe__io__Bytes, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "__hx_dataExposed":
		if incoming == nil {
			var zero bool
			value.__hx_dataExposed = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.__hx_dataExposed = typed
			return true
		default:
			return false
		}
	case "__hx_raw":
		if incoming == nil {
			var zero *hxrt.ByteView
			value.__hx_raw = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.ByteView:
			value.__hx_raw = typed
			return true
		default:
			return false
		}
	case "__hx_rawValid":
		if incoming == nil {
			var zero bool
			value.__hx_rawValid = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.__hx_rawValid = typed
			return true
		default:
			return false
		}
	case "b":
		if incoming == nil {
			var zero []int
			value.b = zero
			return true
		}
		switch typed := incoming.(type) {
		case []int:
			value.b = typed
			return true
		default:
			return false
		}
	case "length":
		if incoming == nil {
			var zero int
			value.length = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.length = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__ArrayIterator(value *haxe__iterators__ArrayIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "array":
		return hxrt.ReadStoredField(value.array)
	case "current":
		return hxrt.ReadStoredField(value.current)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__ArrayIterator(value *haxe__iterators__ArrayIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "array":
		return true
	case "current":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__ArrayIterator(value *haxe__iterators__ArrayIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "array":
		if incoming == nil {
			var zero *hxrt.Array
			value.array = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.array = typed
			return true
		default:
			return false
		}
	case "current":
		if incoming == nil {
			var zero int
			value.current = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.current = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "keys":
		return hxrt.ReadStoredField(value.keys)
	case "map":
		return hxrt.ReadStoredField(value.map_)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		return true
	case "map":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		if incoming == nil {
			var zero hxrt.Anonymous
			value.keys = zero
			return true
		}
		switch typed := incoming.(type) {
		case hxrt.Anonymous:
			value.keys = typed
			return true
		default:
			return false
		}
	case "map":
		if incoming == nil {
			var zero haxe__IMap
			value.map_ = zero
			return true
		}
		switch typed := incoming.(type) {
		case haxe__IMap:
			value.map_ = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "offset":
		return hxrt.ReadStoredField(value.offset)
	case "s":
		return hxrt.ReadStoredField(value.s)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		return true
	case "s":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__StringIterator(value *haxe__iterators__StringIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		if incoming == nil {
			var zero int
			value.offset = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.offset = typed
			return true
		default:
			return false
		}
	case "s":
		if incoming == nil {
			var zero *string
			value.s = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.s = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "offset":
		return hxrt.ReadStoredField(value.offset)
	case "s":
		return hxrt.ReadStoredField(value.s)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		return true
	case "s":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__StringKeyValueIterator(value *haxe__iterators__StringKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "offset":
		if incoming == nil {
			var zero int
			value.offset = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.offset = typed
			return true
		default:
			return false
		}
	case "s":
		if incoming == nil {
			var zero *string
			value.s = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.s = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__rtti__XmlParser(value *haxe__rtti__XmlParser, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "curplatform":
		return hxrt.ReadStoredField(value.curplatform)
	case "newField":
		return hxrt.ReadStoredField(value.newField)
	case "root":
		return hxrt.ReadStoredField(value.root)
	}
	return nil
}

func hxrt__generated_field_has__haxe__rtti__XmlParser(value *haxe__rtti__XmlParser, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "curplatform":
		return true
	case "newField":
		return true
	case "root":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__rtti__XmlParser(value *haxe__rtti__XmlParser, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "curplatform":
		if incoming == nil {
			var zero *string
			value.curplatform = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.curplatform = typed
			return true
		default:
			return false
		}
	case "newField":
		if (hxrt.FunctionOrigin(incoming) == nil) || hxrt.ReflectIsFunction(incoming) {
			value.newField = hx_recover_callable_40a93da1(incoming)
			return true
		}
		return false
	case "root":
		if incoming == nil {
			var zero *hxrt.Array
			value.root = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.root = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__xml__Printer(value *haxe__xml__Printer, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "output":
		return hxrt.ReadStoredField(value.output)
	case "pretty":
		return hxrt.ReadStoredField(value.pretty)
	}
	return nil
}

func hxrt__generated_field_has__haxe__xml__Printer(value *haxe__xml__Printer, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "output":
		return true
	case "pretty":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__xml__Printer(value *haxe__xml__Printer, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "output":
		if incoming == nil {
			var zero *StringBuf
			value.output = zero
			return true
		}
		switch typed := incoming.(type) {
		case *StringBuf:
			value.output = typed
			return true
		default:
			return false
		}
	case "pretty":
		if incoming == nil {
			var zero bool
			value.pretty = zero
			return true
		}
		switch typed := incoming.(type) {
		case bool:
			value.pretty = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__xml__XmlParserException(value *haxe__xml__XmlParserException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "lineNumber":
		return hxrt.ReadStoredField(value.lineNumber)
	case "message":
		return hxrt.ReadStoredField(value.message)
	case "position":
		return hxrt.ReadStoredField(value.position)
	case "positionAtLine":
		return hxrt.ReadStoredField(value.positionAtLine)
	case "xml":
		return hxrt.ReadStoredField(value.xml)
	}
	return nil
}

func hxrt__generated_field_has__haxe__xml__XmlParserException(value *haxe__xml__XmlParserException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "lineNumber":
		return true
	case "message":
		return true
	case "position":
		return true
	case "positionAtLine":
		return true
	case "xml":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__xml__XmlParserException(value *haxe__xml__XmlParserException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "lineNumber":
		if incoming == nil {
			var zero int
			value.lineNumber = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.lineNumber = typed
			return true
		default:
			return false
		}
	case "message":
		if incoming == nil {
			var zero *string
			value.message = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.message = typed
			return true
		default:
			return false
		}
	case "position":
		if incoming == nil {
			var zero int
			value.position = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.position = typed
			return true
		default:
			return false
		}
	case "positionAtLine":
		if incoming == nil {
			var zero int
			value.positionAtLine = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.positionAtLine = typed
			return true
		default:
			return false
		}
	case "xml":
		if incoming == nil {
			var zero *string
			value.xml = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.xml = typed
			return true
		default:
			return false
		}
	}
	return false
}

func reflaxe__go___internal__CompilerReflect_typeField(object any, field *string) any {
	key := *hxrt.StdString(field)
	value, found := hxrt_typeClassMetadataField(object, key)
	if !found {
		return nil
	}
	return value
}

func reflaxe__go___internal__CompilerReflect_hasTypeField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	_, found := hxrt_typeClassMetadataField(object, key)
	return found
}

func reflaxe__go___internal__CompilerReflect_generatedMethod(object any, field *string) any {
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	switch enumValue := value.(type) {
	case *haxe__StackItem:
		return (enumValue != nil)
	case *haxe__io__Encoding:
		return (enumValue != nil)
	case *haxe__io__Error:
		return (enumValue != nil)
	case *haxe__rtti__CType:
		return (enumValue != nil)
	case *haxe__rtti__Rights:
		return (enumValue != nil)
	case *haxe__rtti__TypeTree:
		return (enumValue != nil)
	default:
		return false
	}
}

func hx_is_of_type_token(hx_value any, hx_type any) bool {
	switch hx_type_marker := hx_type.(type) {
	case *hxrt__TypeClassValue:
		if hx_type_marker == nil {
			return false
		}
		if hx_type_marker.name == nil {
			return false
		}
		switch *hx_type_marker.name {
		case "Array":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *hxrt.Array:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Bool":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case bool:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Class":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *hxrt__TypeClassValue:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Dynamic":
			return (hx_value != nil)
		case "Enum":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *hxrt__TypeEnumValue:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Float":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case int:
					return true
				case int8:
					return true
				case int16:
					return true
				case int32:
					return true
				case int64:
					return true
				case uint:
					return true
				case uint8:
					return true
				case uint16:
					return true
				case uint32:
					return true
				case uint64:
					return true
				case uintptr:
					return true
				case float32:
					return true
				case float64:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Int":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case int:
					return true
				case int8:
					return true
				case int16:
					return true
				case int32:
					return true
				case int64:
					return true
				case uint:
					return true
				case uint8:
					return true
				case uint16:
					return true
				case uint32:
					return true
				case uint64:
					return true
				case uintptr:
					return true
				default:
					return false
				}
			}(hx_value)
		case "String":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *string:
					return true
				case string:
					return true
				default:
					return false
				}
			}(hx_value)
		case "StringBuf":
			return hx_is_nominal_StringBuf(hx_value)
		case "Xml":
			return hx_is_nominal_Xml(hx_value)
		case "haxe.Exception":
			return hx_is_nominal_haxe__Exception(hx_value)
		case "haxe.ValueException":
			return hx_is_nominal_haxe__ValueException(hx_value)
		case "haxe._Int64.___Int64":
			return hx_is_nominal_haxe___Int64_____Int64(hx_value)
		case "haxe.ds.StringMap":
			return hx_is_nominal_haxe__ds__StringMap(hx_value)
		case "haxe.io.Bytes":
			return hx_is_nominal_haxe__io__Bytes(hx_value)
		case "haxe.iterators.ArrayIterator":
			return hx_is_nominal_haxe__iterators__ArrayIterator(hx_value)
		case "haxe.iterators.MapKeyValueIterator":
			return hx_is_nominal_haxe__iterators__MapKeyValueIterator(hx_value)
		case "haxe.iterators.StringIterator":
			return hx_is_nominal_haxe__iterators__StringIterator(hx_value)
		case "haxe.iterators.StringKeyValueIterator":
			return hx_is_nominal_haxe__iterators__StringKeyValueIterator(hx_value)
		case "haxe.rtti.XmlParser":
			return hx_is_nominal_haxe__rtti__XmlParser(hx_value)
		case "haxe.xml.Printer":
			return hx_is_nominal_haxe__xml__Printer(hx_value)
		case "haxe.xml.XmlParserException":
			return hx_is_nominal_haxe__xml__XmlParserException(hx_value)
		default:
			return false
		}
	case *hxrt__TypeEnumValue:
		if hx_type_marker == nil {
			return false
		}
		if hx_type_marker.name == nil {
			return false
		}
		switch *hx_type_marker.name {
		case "haxe.StackItem":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__StackItem:
					return true
				default:
					return false
				}
			}(hx_value)
		case "haxe.io.Encoding":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__io__Encoding:
					return true
				default:
					return false
				}
			}(hx_value)
		case "haxe.io.Error":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__io__Error:
					return true
				default:
					return false
				}
			}(hx_value)
		case "haxe.rtti.CType":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__rtti__CType:
					return true
				default:
					return false
				}
			}(hx_value)
		case "haxe.rtti.Rights":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__rtti__Rights:
					return true
				default:
					return false
				}
			}(hx_value)
		case "haxe.rtti.TypeTree":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__rtti__TypeTree:
					return true
				default:
					return false
				}
			}(hx_value)
		default:
			return false
		}
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(hx_structural_recover(argument_0), hx_structural_recover(argument_1))
		return nil
	case *hxrt.Function[func(*haxe__rtti__TypeTree, *haxe__rtti__TypeTree) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_1 any) *haxe__rtti__TypeTree {
			if hx_value_1 == nil {
				var hx_zero_2 *haxe__rtti__TypeTree
				return hx_zero_2
			}
			return hx_value_1.(*haxe__rtti__TypeTree)
		}(argument_0), func(hx_value_3 any) *haxe__rtti__TypeTree {
			if hx_value_3 == nil {
				var hx_zero_4 *haxe__rtti__TypeTree
				return hx_zero_4
			}
			return hx_value_3.(*haxe__rtti__TypeTree)
		}(argument_1))
	case *hxrt.Function[func(any, any) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(argument_0, argument_1)
	case *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_structural_recover(argument_0), hx_structural_recover(argument_1))
	case *hxrt.Function[func() bool]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*hxrt.Array)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_5 any) *hxrt.Array {
			if hx_value_5 == nil {
				var hx_zero_6 *hxrt.Array
				return hx_zero_6
			}
			return hx_value_5.(*hxrt.Array)
		}(argument_0))
		return nil
	case *hxrt.Function[func(*Xml, *string)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0), func(hx_value_7 any) *string {
			if hx_value_7 == nil {
				var hx_zero_8 *string
				return hx_zero_8
			}
			return hx_value_7.(*string)
		}(argument_1))
		return nil
	case *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_structural_recover(argument_0), hx_structural_recover(argument_1))
	case *hxrt.Function[func(*haxe__rtti__TypeTree)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_9 any) *haxe__rtti__TypeTree {
			if hx_value_9 == nil {
				var hx_zero_10 *haxe__rtti__TypeTree
				return hx_zero_10
			}
			return hx_value_9.(*haxe__rtti__TypeTree)
		}(argument_0))
		return nil
	case *hxrt.Function[func(*string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_11 any) *string {
			if hx_value_11 == nil {
				var hx_zero_12 *string
				return hx_zero_12
			}
			return hx_value_11.(*string)
		}(argument_0))
	case *hxrt.Function[func(*string) *hxrt.Array]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
		}(argument_0))
	case *hxrt.Function[func(*string) *haxe__rtti__Rights]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_15 any) *string {
			if hx_value_15 == nil {
				var hx_zero_16 *string
				return hx_zero_16
			}
			return hx_value_15.(*string)
		}(argument_0))
	case *hxrt.Function[func(*Xml)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0))
		return nil
	case *hxrt.Function[func(*Xml) *haxe__rtti__TypeTree]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0))
	case *hxrt.Function[func(*Xml) *hxrt.Array]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0))
	case *hxrt.Function[func(*Xml) hxrt.Anonymous]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0))
	case *hxrt.Function[func(*Xml, any) hxrt.Anonymous]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0), argument_1)
	case *hxrt.Function[func(*Xml) *haxe__rtti__CType]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0))
	case *hxrt.Function[func() *hxrt.Array]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*hxrt.Array, *string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_17 any) *hxrt.Array {
			if hx_value_17 == nil {
				var hx_zero_18 *hxrt.Array
				return hx_zero_18
			}
			return hx_value_17.(*hxrt.Array)
		}(argument_0), func(hx_value_19 any) *string {
			if hx_value_19 == nil {
				var hx_zero_20 *string
				return hx_zero_20
			}
			return hx_value_19.(*string)
		}(argument_1))
	case *hxrt.Function[func(*string, *string) *hxrt.Array]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_21 any) *string {
			if hx_value_21 == nil {
				var hx_zero_22 *string
				return hx_zero_22
			}
			return hx_value_21.(*string)
		}(argument_0), func(hx_value_23 any) *string {
			if hx_value_23 == nil {
				var hx_zero_24 *string
				return hx_zero_24
			}
			return hx_value_23.(*string)
		}(argument_1))
	case *hxrt.Function[func(*string, *string, int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(func(hx_value_25 any) *string {
			if hx_value_25 == nil {
				var hx_zero_26 *string
				return hx_zero_26
			}
			return hx_value_25.(*string)
		}(argument_0), func(hx_value_27 any) *string {
			if hx_value_27 == nil {
				var hx_zero_28 *string
				return hx_zero_28
			}
			return hx_value_27.(*string)
		}(argument_1), func(hx_value_29 any) int {
			if hx_value_29 == nil {
				var hx_zero_30 int
				return hx_zero_30
			}
			return hx_value_29.(int)
		}(argument_2))
	case *hxrt.Function[func(*Xml, *string) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0), func(hx_value_31 any) *string {
			if hx_value_31 == nil {
				var hx_zero_32 *string
				return hx_zero_32
			}
			return hx_value_31.(*string)
		}(argument_1))
	case *hxrt.Function[func(*Xml, *string) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0), func(hx_value_33 any) *string {
			if hx_value_33 == nil {
				var hx_zero_34 *string
				return hx_zero_34
			}
			return hx_value_33.(*string)
		}(argument_1))
	case *hxrt.Function[func(*Xml, *string) *Xml]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0), func(hx_value_35 any) *string {
			if hx_value_35 == nil {
				var hx_zero_36 *string
				return hx_zero_36
			}
			return hx_value_35.(*string)
		}(argument_1))
	case *hxrt.Function[func(*Xml) *Xml]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0))
	case *hxrt.Function[func(*Xml) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0))
	case *hxrt.Function[func(*string) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_37 any) *string {
			if hx_value_37 == nil {
				var hx_zero_38 *string
				return hx_zero_38
			}
			return hx_value_37.(*string)
		}(argument_0))
	case *hxrt.Function[func(*string, *string)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_39 any) *string {
			if hx_value_39 == nil {
				var hx_zero_40 *string
				return hx_zero_40
			}
			return hx_value_39.(*string)
		}(argument_0), func(hx_value_41 any) *string {
			if hx_value_41 == nil {
				var hx_zero_42 *string
				return hx_zero_42
			}
			return hx_value_41.(*string)
		}(argument_1))
		return nil
	case *hxrt.Function[func(*string)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		hxrt.FunctionCall(source)(func(hx_value_43 any) *string {
			if hx_value_43 == nil {
				var hx_zero_44 *string
				return hx_zero_44
			}
			return hx_value_43.(*string)
		}(argument_0))
		return nil
	case *hxrt.Function[func(*string) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_45 any) *string {
			if hx_value_45 == nil {
				var hx_zero_46 *string
				return hx_zero_46
			}
			return hx_value_45.(*string)
		}(argument_0))
	case *hxrt.Function[func(*string) hxrt.Anonymous]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_47 any) *string {
			if hx_value_47 == nil {
				var hx_zero_48 *string
				return hx_zero_48
			}
			return hx_value_47.(*string)
		}(argument_0))
	case *hxrt.Function[func() *Xml]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(*Xml) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0))
	case *hxrt.Function[func(*Xml, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(hx_recover_class_Xml(argument_0), func(hx_value_49 any) int {
			if hx_value_49 == nil {
				var hx_zero_50 int
				return hx_zero_50
			}
			return hx_value_49.(int)
		}(argument_1))
		return nil
	case *hxrt.Function[func()]:
		hxrt.FunctionCall(source)()
		return nil
	case *hxrt.Function[func(*string, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_51 any) *string {
			if hx_value_51 == nil {
				var hx_zero_52 *string
				return hx_zero_52
			}
			return hx_value_51.(*string)
		}(argument_0), argument_1)
		return nil
	case *hxrt.Function[func(*string) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_53 any) *string {
			if hx_value_53 == nil {
				var hx_zero_54 *string
				return hx_zero_54
			}
			return hx_value_53.(*string)
		}(argument_0))
	case *hxrt.Function[func(any) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(any, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(argument_0, argument_1)
		return nil
	case *hxrt.Function[func(any) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_55 any) int {
			if hx_value_55 == nil {
				var hx_zero_56 int
				return hx_zero_56
			}
			return hx_value_55.(int)
		}(argument_0))
	case *hxrt.Function[func(int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_57 any) int {
			if hx_value_57 == nil {
				var hx_zero_58 int
				return hx_zero_58
			}
			return hx_value_57.(int)
		}(argument_0), func(hx_value_59 any) int {
			if hx_value_59 == nil {
				var hx_zero_60 int
				return hx_zero_60
			}
			return hx_value_59.(int)
		}(argument_1))
		return nil
	case *hxrt.Function[func(int, *haxe__io__Bytes, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		var argument_3 any = nil
		if 3 < len(arguments) {
			argument_3 = arguments[3]
		}
		hxrt.FunctionCall(source)(func(hx_value_61 any) int {
			if hx_value_61 == nil {
				var hx_zero_62 int
				return hx_zero_62
			}
			return hx_value_61.(int)
		}(argument_0), hx_recover_class_haxe__io__Bytes(argument_1), func(hx_value_63 any) int {
			if hx_value_63 == nil {
				var hx_zero_64 int
				return hx_zero_64
			}
			return hx_value_63.(int)
		}(argument_2), func(hx_value_65 any) int {
			if hx_value_65 == nil {
				var hx_zero_66 int
				return hx_zero_66
			}
			return hx_value_65.(int)
		}(argument_3))
		return nil
	case *hxrt.Function[func(int, int, int)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		hxrt.FunctionCall(source)(func(hx_value_67 any) int {
			if hx_value_67 == nil {
				var hx_zero_68 int
				return hx_zero_68
			}
			return hx_value_67.(int)
		}(argument_0), func(hx_value_69 any) int {
			if hx_value_69 == nil {
				var hx_zero_70 int
				return hx_zero_70
			}
			return hx_value_69.(int)
		}(argument_1), func(hx_value_71 any) int {
			if hx_value_71 == nil {
				var hx_zero_72 int
				return hx_zero_72
			}
			return hx_value_71.(int)
		}(argument_2))
		return nil
	case *hxrt.Function[func(int, int) *haxe__io__Bytes]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_73 any) int {
			if hx_value_73 == nil {
				var hx_zero_74 int
				return hx_zero_74
			}
			return hx_value_73.(int)
		}(argument_0), func(hx_value_75 any) int {
			if hx_value_75 == nil {
				var hx_zero_76 int
				return hx_zero_76
			}
			return hx_value_75.(int)
		}(argument_1))
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(argument_0))
	case *hxrt.Function[func(int) float64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_77 any) int {
			if hx_value_77 == nil {
				var hx_zero_78 int
				return hx_zero_78
			}
			return hx_value_77.(int)
		}(argument_0))
	case *hxrt.Function[func(int, float64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_79 any) int {
			if hx_value_79 == nil {
				var hx_zero_80 int
				return hx_zero_80
			}
			return hx_value_79.(int)
		}(argument_0), func(hx_value_81 any) float64 {
			if hx_value_81 == nil {
				var hx_zero_82 float64
				return hx_zero_82
			}
			return hx_value_81.(float64)
		}(argument_1))
		return nil
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_83 any) int {
			if hx_value_83 == nil {
				var hx_zero_84 int
				return hx_zero_84
			}
			return hx_value_83.(int)
		}(argument_0))
	case *hxrt.Function[func(int, *haxe___Int64_____Int64)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(func(hx_value_85 any) int {
			if hx_value_85 == nil {
				var hx_zero_86 int
				return hx_zero_86
			}
			return hx_value_85.(int)
		}(argument_0), hx_recover_class_haxe___Int64_____Int64(argument_1))
		return nil
	case *hxrt.Function[func(int, int, *haxe__io__Encoding) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		var argument_2 any = nil
		if 2 < len(arguments) {
			argument_2 = arguments[2]
		}
		return hxrt.FunctionCall(source)(func(hx_value_87 any) int {
			if hx_value_87 == nil {
				var hx_zero_88 int
				return hx_zero_88
			}
			return hx_value_87.(int)
		}(argument_0), func(hx_value_89 any) int {
			if hx_value_89 == nil {
				var hx_zero_90 int
				return hx_zero_90
			}
			return hx_value_89.(int)
		}(argument_1), func(hx_value_91 any) *haxe__io__Encoding {
			if hx_value_91 == nil {
				var hx_zero_92 *haxe__io__Encoding
				return hx_zero_92
			}
			return hx_value_91.(*haxe__io__Encoding)
		}(argument_2))
	case *hxrt.Function[func(int, int) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(func(hx_value_93 any) int {
			if hx_value_93 == nil {
				var hx_zero_94 int
				return hx_zero_94
			}
			return hx_value_93.(int)
		}(argument_0), func(hx_value_95 any) int {
			if hx_value_95 == nil {
				var hx_zero_96 int
				return hx_zero_96
			}
			return hx_value_95.(int)
		}(argument_1))
	case *hxrt.Function[func() []int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_2f3c5984(value any) *hxrt.Function[func() *Xml] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *Xml]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.Array]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *Xml]:
		return source
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *Xml {
			hxrt.FunctionCall(source)()
			return hx_recover_class_Xml(nil)
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() *Xml {
			return hx_recover_class_Xml(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_40a93da1(value any) *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous)]:
		return source
	case *hxrt.Function[func(*haxe__rtti__TypeTree, *haxe__rtti__TypeTree) int]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) *haxe__rtti__TypeTree {
				if hx_value_1 == nil {
					var hx_zero_2 *haxe__rtti__TypeTree
					return hx_zero_2
				}
				return hx_value_1.(*haxe__rtti__TypeTree)
			}(arg_0), func(hx_value_3 any) *haxe__rtti__TypeTree {
				if hx_value_3 == nil {
					var hx_zero_4 *haxe__rtti__TypeTree
					return hx_zero_4
				}
				return hx_value_3.(*haxe__rtti__TypeTree)
			}(arg_1))
		})
	case *hxrt.Function[func(any, any) int]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(arg_0, arg_1)
		})
	case *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) int]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(arg_0, arg_1)
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(arg_0, arg_1)
		})
	case *hxrt.Function[func(*Xml, *string)]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(hx_recover_class_Xml(arg_0), func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(arg_1))
		})
	case *hxrt.Function[func(hxrt.Anonymous, hxrt.Anonymous) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(arg_0, arg_1)
		})
	case *hxrt.Function[func(*Xml, any) hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(hx_recover_class_Xml(arg_0), arg_1)
		})
	case *hxrt.Function[func(*hxrt.Array, *string) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(func(hx_value_7 any) *hxrt.Array {
				if hx_value_7 == nil {
					var hx_zero_8 *hxrt.Array
					return hx_zero_8
				}
				return hx_value_7.(*hxrt.Array)
			}(arg_0), func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(arg_1))
		})
	case *hxrt.Function[func(*string, *string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(arg_0), func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(arg_1))
		})
	case *hxrt.Function[func(*Xml, *string) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(hx_recover_class_Xml(arg_0), func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(arg_1))
		})
	case *hxrt.Function[func(*Xml, *string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(hx_recover_class_Xml(arg_0), func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(arg_1))
		})
	case *hxrt.Function[func(*Xml, *string) *Xml]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(hx_recover_class_Xml(arg_0), func(hx_value_19 any) *string {
				if hx_value_19 == nil {
					var hx_zero_20 *string
					return hx_zero_20
				}
				return hx_value_19.(*string)
			}(arg_1))
		})
	case *hxrt.Function[func(*string, *string)]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(func(hx_value_21 any) *string {
				if hx_value_21 == nil {
					var hx_zero_22 *string
					return hx_zero_22
				}
				return hx_value_21.(*string)
			}(arg_0), func(hx_value_23 any) *string {
				if hx_value_23 == nil {
					var hx_zero_24 *string
					return hx_zero_24
				}
				return hx_value_23.(*string)
			}(arg_1))
		})
	case *hxrt.Function[func(*Xml, int)]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(hx_recover_class_Xml(arg_0), func(hx_value_25 any) int {
				if hx_value_25 == nil {
					var hx_zero_26 int
					return hx_zero_26
				}
				return hx_value_25.(int)
			}(arg_1))
		})
	case *hxrt.Function[func(*string, any)]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(func(hx_value_27 any) *string {
				if hx_value_27 == nil {
					var hx_zero_28 *string
					return hx_zero_28
				}
				return hx_value_27.(*string)
			}(arg_0), arg_1)
		})
	case *hxrt.Function[func(any, any)]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(arg_0, arg_1)
		})
	case *hxrt.Function[func(int, int)]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(func(hx_value_29 any) int {
				if hx_value_29 == nil {
					var hx_zero_30 int
					return hx_zero_30
				}
				return hx_value_29.(int)
			}(arg_0), func(hx_value_31 any) int {
				if hx_value_31 == nil {
					var hx_zero_32 int
					return hx_zero_32
				}
				return hx_value_31.(int)
			}(arg_1))
		})
	case *hxrt.Function[func(int, int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(func(hx_value_33 any) int {
				if hx_value_33 == nil {
					var hx_zero_34 int
					return hx_zero_34
				}
				return hx_value_33.(int)
			}(arg_0), func(hx_value_35 any) int {
				if hx_value_35 == nil {
					var hx_zero_36 int
					return hx_zero_36
				}
				return hx_value_35.(int)
			}(arg_1))
		})
	case *hxrt.Function[func(int, float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(func(hx_value_37 any) int {
				if hx_value_37 == nil {
					var hx_zero_38 int
					return hx_zero_38
				}
				return hx_value_37.(int)
			}(arg_0), func(hx_value_39 any) float64 {
				if hx_value_39 == nil {
					var hx_zero_40 float64
					return hx_zero_40
				}
				return hx_value_39.(float64)
			}(arg_1))
		})
	case *hxrt.Function[func(int, *haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(func(hx_value_41 any) int {
				if hx_value_41 == nil {
					var hx_zero_42 int
					return hx_zero_42
				}
				return hx_value_41.(int)
			}(arg_0), hx_recover_class_haxe___Int64_____Int64(arg_1))
		})
	case *hxrt.Function[func(int, int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 hxrt.Anonymous, arg_1 hxrt.Anonymous) {
			hxrt.FunctionCall(source)(func(hx_value_43 any) int {
				if hx_value_43 == nil {
					var hx_zero_44 int
					return hx_zero_44
				}
				return hx_value_43.(int)
			}(arg_0), func(hx_value_45 any) int {
				if hx_value_45 == nil {
					var hx_zero_46 int
					return hx_zero_46
				}
				return hx_value_45.(int)
			}(arg_1))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_b9d17ba0(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.Array]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *Xml]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(nil)
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_19 any) *string {
				if hx_value_19 == nil {
					var hx_zero_20 *string
					return hx_zero_20
				}
				return hx_value_19.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_21 any) *string {
				if hx_value_21 == nil {
					var hx_zero_22 *string
					return hx_zero_22
				}
				return hx_value_21.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_23 any) *string {
				if hx_value_23 == nil {
					var hx_zero_24 *string
					return hx_zero_24
				}
				return hx_value_23.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_25 any) *string {
				if hx_value_25 == nil {
					var hx_zero_26 *string
					return hx_zero_26
				}
				return hx_value_25.(*string)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_c5309b43(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.Array]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *Xml]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_cf6928b1(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.Array]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_9 any) bool {
				if hx_value_9 == nil {
					var hx_zero_10 bool
					return hx_zero_10
				}
				return hx_value_9.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *Xml]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_11 any) bool {
				if hx_value_11 == nil {
					var hx_zero_12 bool
					return hx_zero_12
				}
				return hx_value_11.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_13 any) bool {
				if hx_value_13 == nil {
					var hx_zero_14 bool
					return hx_zero_14
				}
				return hx_value_13.(bool)
			}(nil)
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_15 any) bool {
				if hx_value_15 == nil {
					var hx_zero_16 bool
					return hx_zero_16
				}
				return hx_value_15.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_17 any) bool {
				if hx_value_17 == nil {
					var hx_zero_18 bool
					return hx_zero_18
				}
				return hx_value_17.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_19 any) bool {
				if hx_value_19 == nil {
					var hx_zero_20 bool
					return hx_zero_20
				}
				return hx_value_19.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_21 any) bool {
				if hx_value_21 == nil {
					var hx_zero_22 bool
					return hx_zero_22
				}
				return hx_value_21.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_23 any) bool {
				if hx_value_23 == nil {
					var hx_zero_24 bool
					return hx_zero_24
				}
				return hx_value_23.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_25 any) bool {
				if hx_value_25 == nil {
					var hx_zero_26 bool
					return hx_zero_26
				}
				return hx_value_25.(bool)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_StringBuf(value any) *StringBuf {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *StringBuf:
		return source
	default:
		return value.(*StringBuf)
	}
}

func hx_recover_class_Xml(value any) *Xml {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Xml:
		return source
	default:
		return value.(*Xml)
	}
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

func hx_recover_class_haxe___Int64_____Int64(value any) *haxe___Int64_____Int64 {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return source
	default:
		return value.(*haxe___Int64_____Int64)
	}
}

func hx_recover_class_haxe__ds__StringMap(value any) *haxe__ds__StringMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return source
	default:
		return value.(*haxe__ds__StringMap)
	}
}

func hx_recover_class_haxe__io__Bytes(value any) *haxe__io__Bytes {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return source
	default:
		return value.(*haxe__io__Bytes)
	}
}

func hx_is_nominal_StringBuf(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *StringBuf:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_Xml(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *Xml:
		return (source != nil)
	default:
		return false
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

func hx_is_nominal_haxe___Int64_____Int64(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ds__StringMap(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Bytes(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__iterators__ArrayIterator(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__ArrayIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__iterators__MapKeyValueIterator(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__MapKeyValueIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__iterators__StringIterator(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__StringIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__iterators__StringKeyValueIterator(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__StringKeyValueIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__rtti__XmlParser(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__rtti__XmlParser:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__xml__Printer(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__xml__Printer:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__xml__XmlParserException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__xml__XmlParserException:
		return (source != nil)
	default:
		return false
	}
}
