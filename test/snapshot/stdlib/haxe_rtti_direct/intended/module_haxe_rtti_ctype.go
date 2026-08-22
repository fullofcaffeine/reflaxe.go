package main

import "snapshot/hxrt"

func haxe__rtti__TypeApi_constructorEq(c1 map[string]any, c2 map[string]any) bool {
	if !hxrt.StringEqualStringPtr(func(hx_obj_1 map[string]any) *string {
		hx_field_2 := hx_obj_1["name"]
		if hx_field_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_field_2.(*string)
	}(c1), func(hx_obj_4 map[string]any) *string {
		hx_field_5 := hx_obj_4["name"]
		if hx_field_5 == nil {
			var hx_zero_6 *string
			return hx_zero_6
		}
		return hx_field_5.(*string)
	}(c2)) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_obj_7 map[string]any) *string {
		hx_field_8 := hx_obj_7["doc"]
		if hx_field_8 == nil {
			var hx_zero_9 *string
			return hx_zero_9
		}
		return hx_field_8.(*string)
	}(c1), func(hx_obj_10 map[string]any) *string {
		hx_field_11 := hx_obj_10["doc"]
		if hx_field_11 == nil {
			var hx_zero_12 *string
			return hx_zero_12
		}
		return hx_field_11.(*string)
	}(c2)) {
		return false
	}
	if (func(hx_obj_13 map[string]any) *hxrt.Array {
		hx_field_14 := hx_obj_13["args"]
		if hx_field_14 == nil {
			var hx_zero_15 *hxrt.Array
			return hx_zero_15
		}
		return hx_field_14.(*hxrt.Array)
	}(c1) == nil) != (func(hx_obj_16 map[string]any) *hxrt.Array {
		hx_field_17 := hx_obj_16["args"]
		if hx_field_17 == nil {
			var hx_zero_18 *hxrt.Array
			return hx_zero_18
		}
		return hx_field_17.(*hxrt.Array)
	}(c2) == nil) {
		return false
	}
	if (func(hx_obj_19 map[string]any) *hxrt.Array {
		hx_field_20 := hx_obj_19["args"]
		if hx_field_20 == nil {
			var hx_zero_21 *hxrt.Array
			return hx_zero_21
		}
		return hx_field_20.(*hxrt.Array)
	}(c1) != nil) && !haxe__rtti__TypeApi_sameConstructorArguments(func(hx_obj_22 map[string]any) *hxrt.Array {
		hx_field_23 := hx_obj_22["args"]
		if hx_field_23 == nil {
			var hx_zero_24 *hxrt.Array
			return hx_zero_24
		}
		return hx_field_23.(*hxrt.Array)
	}(c1), func(hx_obj_25 map[string]any) *hxrt.Array {
		hx_field_26 := hx_obj_25["args"]
		if hx_field_26 == nil {
			var hx_zero_27 *hxrt.Array
			return hx_zero_27
		}
		return hx_field_26.(*hxrt.Array)
	}(c2)) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_fieldEq(f1 map[string]any, f2 map[string]any) bool {
	if !hxrt.StringEqualStringPtr(func(hx_obj_28 map[string]any) *string {
		hx_field_29 := hx_obj_28["name"]
		if hx_field_29 == nil {
			var hx_zero_30 *string
			return hx_zero_30
		}
		return hx_field_29.(*string)
	}(f1), func(hx_obj_31 map[string]any) *string {
		hx_field_32 := hx_obj_31["name"]
		if hx_field_32 == nil {
			var hx_zero_33 *string
			return hx_zero_33
		}
		return hx_field_32.(*string)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_typeEq(func(hx_obj_34 map[string]any) *haxe__rtti__CType {
		hx_field_35 := hx_obj_34["type"]
		if hx_field_35 == nil {
			var hx_zero_36 *haxe__rtti__CType
			return hx_zero_36
		}
		return hx_field_35.(*haxe__rtti__CType)
	}(f1), func(hx_obj_37 map[string]any) *haxe__rtti__CType {
		hx_field_38 := hx_obj_37["type"]
		if hx_field_38 == nil {
			var hx_zero_39 *haxe__rtti__CType
			return hx_zero_39
		}
		return hx_field_38.(*haxe__rtti__CType)
	}(f2)) {
		return false
	}
	if func(hx_obj_40 map[string]any) bool {
		hx_field_41 := hx_obj_40["isPublic"]
		if hx_field_41 == nil {
			var hx_zero_42 bool
			return hx_zero_42
		}
		return hx_field_41.(bool)
	}(f1) != func(hx_obj_43 map[string]any) bool {
		hx_field_44 := hx_obj_43["isPublic"]
		if hx_field_44 == nil {
			var hx_zero_45 bool
			return hx_zero_45
		}
		return hx_field_44.(bool)
	}(f2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_obj_46 map[string]any) *string {
		hx_field_47 := hx_obj_46["doc"]
		if hx_field_47 == nil {
			var hx_zero_48 *string
			return hx_zero_48
		}
		return hx_field_47.(*string)
	}(f1), func(hx_obj_49 map[string]any) *string {
		hx_field_50 := hx_obj_49["doc"]
		if hx_field_50 == nil {
			var hx_zero_51 *string
			return hx_zero_51
		}
		return hx_field_50.(*string)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_obj_52 map[string]any) *haxe__rtti__Rights {
		hx_field_53 := hx_obj_52["get"]
		if hx_field_53 == nil {
			var hx_zero_54 *haxe__rtti__Rights
			return hx_zero_54
		}
		return hx_field_53.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_55 map[string]any) *haxe__rtti__Rights {
		hx_field_56 := hx_obj_55["get"]
		if hx_field_56 == nil {
			var hx_zero_57 *haxe__rtti__Rights
			return hx_zero_57
		}
		return hx_field_56.(*haxe__rtti__Rights)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_obj_58 map[string]any) *haxe__rtti__Rights {
		hx_field_59 := hx_obj_58["set"]
		if hx_field_59 == nil {
			var hx_zero_60 *haxe__rtti__Rights
			return hx_zero_60
		}
		return hx_field_59.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_61 map[string]any) *haxe__rtti__Rights {
		hx_field_62 := hx_obj_61["set"]
		if hx_field_62 == nil {
			var hx_zero_63 *haxe__rtti__Rights
			return hx_zero_63
		}
		return hx_field_62.(*haxe__rtti__Rights)
	}(f2)) {
		return false
	}
	if (func(hx_obj_64 map[string]any) *hxrt.Array {
		hx_field_65 := hx_obj_64["params"]
		if hx_field_65 == nil {
			var hx_zero_66 *hxrt.Array
			return hx_zero_66
		}
		return hx_field_65.(*hxrt.Array)
	}(f1) == nil) != (func(hx_obj_67 map[string]any) *hxrt.Array {
		hx_field_68 := hx_obj_67["params"]
		if hx_field_68 == nil {
			var hx_zero_69 *hxrt.Array
			return hx_zero_69
		}
		return hx_field_68.(*hxrt.Array)
	}(f2) == nil) {
		return false
	}
	if (func(hx_obj_70 map[string]any) *hxrt.Array {
		hx_field_71 := hx_obj_70["params"]
		if hx_field_71 == nil {
			var hx_zero_72 *hxrt.Array
			return hx_zero_72
		}
		return hx_field_71.(*hxrt.Array)
	}(f1) != nil) && !haxe__rtti__TypeApi_sameTypeParamNames(func(hx_obj_73 map[string]any) *hxrt.Array {
		hx_field_74 := hx_obj_73["params"]
		if hx_field_74 == nil {
			var hx_zero_75 *hxrt.Array
			return hx_zero_75
		}
		return hx_field_74.(*hxrt.Array)
	}(f1), func(hx_obj_76 map[string]any) *hxrt.Array {
		hx_field_77 := hx_obj_76["params"]
		if hx_field_77 == nil {
			var hx_zero_78 *hxrt.Array
			return hx_zero_78
		}
		return hx_field_77.(*hxrt.Array)
	}(f2)) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_isVar(t *haxe__rtti__CType) bool {
	var hx_if_79 bool
	if t.tag == 4 {
		_g := t.params[0].(*hxrt.Array)
		_ = _g
		_g_1 := t.params[1].(*haxe__rtti__CType)
		_ = _g_1
		hx_if_79 = false
	} else {
		hx_if_79 = true
	}
	return hx_if_79
}

func haxe__rtti__TypeApi_rightsEq(r1 *haxe__rtti__Rights, r2 *haxe__rtti__Rights) bool {
	if r1 == r2 {
		return true
	}
	if r1.tag == 2 {
		_g := r1.params[0].(*string)
		m1 := _g
		if r2.tag == 2 {
			_g_1 := r2.params[0].(*string)
			m2 := _g_1
			return hxrt.StringEqualStringPtr(m1, m2)
		} else {
		}
	} else {
	}
	return false
}

func haxe__rtti__TypeApi_sameClassFields(l1 *hxrt.Array, l2 *hxrt.Array) bool {
	if l1.Len() != l2.Len() {
		return false
	}
	_g := 0
	_g1 := l1.Len()
	for _g < _g1 {
		hx_post_80 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_80
		if !haxe__rtti__TypeApi_fieldEq(func(hx_value_81 any) map[string]any {
			if hx_value_81 == nil {
				var hx_zero_82 map[string]any
				return hx_zero_82
			}
			return hx_value_81.(map[string]any)
		}(l1.Get(i)), func(hx_value_83 any) map[string]any {
			if hx_value_83 == nil {
				var hx_zero_84 map[string]any
				return hx_zero_84
			}
			return hx_value_83.(map[string]any)
		}(l2.Get(i))) {
			return false
		}
	}
	return true
}

func haxe__rtti__TypeApi_sameConstructorArguments(l1 *hxrt.Array, l2 *hxrt.Array) bool {
	if l1.Len() != l2.Len() {
		return false
	}
	_g := 0
	_g1 := l1.Len()
	for _g < _g1 {
		hx_post_85 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_85
		a := func(hx_value_86 any) map[string]any {
			if hx_value_86 == nil {
				var hx_zero_87 map[string]any
				return hx_zero_87
			}
			return hx_value_86.(map[string]any)
		}(l1.Get(i))
		b := func(hx_value_88 any) map[string]any {
			if hx_value_88 == nil {
				var hx_zero_89 map[string]any
				return hx_zero_89
			}
			return hx_value_88.(map[string]any)
		}(l2.Get(i))
		if (!hxrt.StringEqualStringPtr(func(hx_obj_90 map[string]any) *string {
			hx_field_91 := hx_obj_90["name"]
			if hx_field_91 == nil {
				var hx_zero_92 *string
				return hx_zero_92
			}
			return hx_field_91.(*string)
		}(a), func(hx_obj_93 map[string]any) *string {
			hx_field_94 := hx_obj_93["name"]
			if hx_field_94 == nil {
				var hx_zero_95 *string
				return hx_zero_95
			}
			return hx_field_94.(*string)
		}(b)) || (func(hx_obj_96 map[string]any) bool {
			hx_field_97 := hx_obj_96["opt"]
			if hx_field_97 == nil {
				var hx_zero_98 bool
				return hx_zero_98
			}
			return hx_field_97.(bool)
		}(a) != func(hx_obj_99 map[string]any) bool {
			hx_field_100 := hx_obj_99["opt"]
			if hx_field_100 == nil {
				var hx_zero_101 bool
				return hx_zero_101
			}
			return hx_field_100.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_102 map[string]any) *haxe__rtti__CType {
			hx_field_103 := hx_obj_102["t"]
			if hx_field_103 == nil {
				var hx_zero_104 *haxe__rtti__CType
				return hx_zero_104
			}
			return hx_field_103.(*haxe__rtti__CType)
		}(a), func(hx_obj_105 map[string]any) *haxe__rtti__CType {
			hx_field_106 := hx_obj_105["t"]
			if hx_field_106 == nil {
				var hx_zero_107 *haxe__rtti__CType
				return hx_zero_107
			}
			return hx_field_106.(*haxe__rtti__CType)
		}(b)) {
			return false
		}
	}
	return true
}

func haxe__rtti__TypeApi_sameFunctionArguments(l1 *hxrt.Array, l2 *hxrt.Array) bool {
	if l1.Len() != l2.Len() {
		return false
	}
	_g := 0
	_g1 := l1.Len()
	for _g < _g1 {
		hx_post_108 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_108
		a := func(hx_value_109 any) map[string]any {
			if hx_value_109 == nil {
				var hx_zero_110 map[string]any
				return hx_zero_110
			}
			return hx_value_109.(map[string]any)
		}(l1.Get(i))
		b := func(hx_value_111 any) map[string]any {
			if hx_value_111 == nil {
				var hx_zero_112 map[string]any
				return hx_zero_112
			}
			return hx_value_111.(map[string]any)
		}(l2.Get(i))
		if (!hxrt.StringEqualStringPtr(func(hx_obj_113 map[string]any) *string {
			hx_field_114 := hx_obj_113["name"]
			if hx_field_114 == nil {
				var hx_zero_115 *string
				return hx_zero_115
			}
			return hx_field_114.(*string)
		}(a), func(hx_obj_116 map[string]any) *string {
			hx_field_117 := hx_obj_116["name"]
			if hx_field_117 == nil {
				var hx_zero_118 *string
				return hx_zero_118
			}
			return hx_field_117.(*string)
		}(b)) || (func(hx_obj_119 map[string]any) bool {
			hx_field_120 := hx_obj_119["opt"]
			if hx_field_120 == nil {
				var hx_zero_121 bool
				return hx_zero_121
			}
			return hx_field_120.(bool)
		}(a) != func(hx_obj_122 map[string]any) bool {
			hx_field_123 := hx_obj_122["opt"]
			if hx_field_123 == nil {
				var hx_zero_124 bool
				return hx_zero_124
			}
			return hx_field_123.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_125 map[string]any) *haxe__rtti__CType {
			hx_field_126 := hx_obj_125["t"]
			if hx_field_126 == nil {
				var hx_zero_127 *haxe__rtti__CType
				return hx_zero_127
			}
			return hx_field_126.(*haxe__rtti__CType)
		}(a), func(hx_obj_128 map[string]any) *haxe__rtti__CType {
			hx_field_129 := hx_obj_128["t"]
			if hx_field_129 == nil {
				var hx_zero_130 *haxe__rtti__CType
				return hx_zero_130
			}
			return hx_field_129.(*haxe__rtti__CType)
		}(b)) {
			return false
		}
	}
	return true
}

func haxe__rtti__TypeApi_sameTypeParamNames(p1 *hxrt.Array, p2 *hxrt.Array) bool {
	if p1.Len() != p2.Len() {
		return false
	}
	_g := 0
	_g1 := p1.Len()
	for _g < _g1 {
		hx_post_131 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_131
		if !hxrt.StringEqualAny(p1.Get(i), p2.Get(i)) {
			return false
		}
	}
	return true
}

func haxe__rtti__TypeApi_sameTypes(l1 *hxrt.Array, l2 *hxrt.Array) bool {
	if l1.Len() != l2.Len() {
		return false
	}
	_g := 0
	_g1 := l1.Len()
	for _g < _g1 {
		hx_post_132 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_132
		if !haxe__rtti__TypeApi_typeEq(func(hx_value_133 any) *haxe__rtti__CType {
			if hx_value_133 == nil {
				var hx_zero_134 *haxe__rtti__CType
				return hx_zero_134
			}
			return hx_value_133.(*haxe__rtti__CType)
		}(l1.Get(i)), func(hx_value_135 any) *haxe__rtti__CType {
			if hx_value_135 == nil {
				var hx_zero_136 *haxe__rtti__CType
				return hx_zero_136
			}
			return hx_value_135.(*haxe__rtti__CType)
		}(l2.Get(i))) {
			return false
		}
	}
	return true
}

func haxe__rtti__TypeApi_typeEq(t1 *haxe__rtti__CType, t2 *haxe__rtti__CType) bool {
	switch t1.tag {
	case 0:
		return (t2 == haxe__rtti__CType_CUnknown)
	case 1:
		_g := t1.params[0].(*string)
		_g1 := t1.params[1].(*hxrt.Array)
		name := _g
		params := _g1
		if t2.tag == 1 {
			_g_1 := t2.params[0].(*string)
			_g1_1 := t2.params[1].(*hxrt.Array)
			name2 := _g_1
			params2 := _g1_1
			return (hxrt.StringEqualStringPtr(name, name2) && haxe__rtti__TypeApi_sameTypes(params, params2))
		} else {
		}
	case 2:
		_g_2 := t1.params[0].(*string)
		_g1_2 := t1.params[1].(*hxrt.Array)
		name_1 := _g_2
		params_1 := _g1_2
		if t2.tag == 2 {
			_g_3 := t2.params[0].(*string)
			_g1_3 := t2.params[1].(*hxrt.Array)
			name2_1 := _g_3
			params2_1 := _g1_3
			return (hxrt.StringEqualStringPtr(name_1, name2_1) && haxe__rtti__TypeApi_sameTypes(params_1, params2_1))
		} else {
		}
	case 3:
		_g_4 := t1.params[0].(*string)
		_g1_4 := t1.params[1].(*hxrt.Array)
		name_2 := _g_4
		params_2 := _g1_4
		if t2.tag == 3 {
			_g_5 := t2.params[0].(*string)
			_g1_5 := t2.params[1].(*hxrt.Array)
			name2_2 := _g_5
			params2_2 := _g1_5
			return (hxrt.StringEqualStringPtr(name_2, name2_2) && haxe__rtti__TypeApi_sameTypes(params_2, params2_2))
		} else {
		}
	case 4:
		_g_6 := t1.params[0].(*hxrt.Array)
		_g1_6 := t1.params[1].(*haxe__rtti__CType)
		args := _g_6
		ret := _g1_6
		if t2.tag == 4 {
			_g_7 := t2.params[0].(*hxrt.Array)
			_g1_7 := t2.params[1].(*haxe__rtti__CType)
			args2 := _g_7
			ret2 := _g1_7
			return (haxe__rtti__TypeApi_sameFunctionArguments(args, args2) && haxe__rtti__TypeApi_typeEq(ret, ret2))
		} else {
		}
	case 5:
		_g_8 := t1.params[0].(*hxrt.Array)
		fields := _g_8
		if t2.tag == 5 {
			_g_9 := t2.params[0].(*hxrt.Array)
			fields2 := _g_9
			return haxe__rtti__TypeApi_sameClassFields(fields, fields2)
		} else {
		}
	case 6:
		_g_10 := t1.params[0].(*haxe__rtti__CType)
		t := _g_10
		if t2.tag == 6 {
			_g_11 := t2.params[0].(*haxe__rtti__CType)
			t2_1 := _g_11
			if (t == nil) != (t2_1 == nil) {
				return false
			}
			return ((t == nil) || haxe__rtti__TypeApi_typeEq(t, t2_1))
		} else {
		}
	case 7:
		_g_12 := t1.params[0].(*string)
		_g1_8 := t1.params[1].(*hxrt.Array)
		name_3 := _g_12
		params_3 := _g1_8
		if t2.tag == 7 {
			_g_13 := t2.params[0].(*string)
			_g1_9 := t2.params[1].(*hxrt.Array)
			name2_3 := _g_13
			params2_3 := _g1_9
			return (hxrt.StringEqualStringPtr(name_3, name2_3) && haxe__rtti__TypeApi_sameTypes(params_3, params2_3))
		} else {
		}
	}
	return false
}

func haxe__rtti__TypeApi_typeInfos(t *haxe__rtti__TypeTree) map[string]any {
	var inf map[string]any
	switch t.tag {
	case 0:
		_g := t.params[0].(*string)
		_ = _g
		_g_1 := t.params[1].(*string)
		_ = _g_1
		_g_2 := t.params[2].(*hxrt.Array)
		_ = _g_2
		hxrt.Throw(hxrt.StringFromLiteral("Unexpected Package"))
	case 1:
		_g_3 := t.params[0].(map[string]any)
		c := _g_3
		inf = c
	case 2:
		_g_4 := t.params[0].(map[string]any)
		e := _g_4
		inf = e
	case 3:
		_g_5 := t.params[0].(map[string]any)
		t_1 := _g_5
		inf = t_1
	case 4:
		_g_6 := t.params[0].(map[string]any)
		a := _g_6
		inf = a
	}
	return inf
}

func haxe__rtti__CTypeTools_classField(cf map[string]any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func(hx_obj_137 map[string]any) *string {
		hx_field_138 := hx_obj_137["name"]
		if hx_field_138 == nil {
			var hx_zero_139 *string
			return hx_zero_139
		}
		return hx_field_138.(*string)
	}(cf), hxrt.StringFromLiteral(":")), haxe__rtti__CTypeTools_toString(func(hx_obj_140 map[string]any) *haxe__rtti__CType {
		hx_field_141 := hx_obj_140["type"]
		if hx_field_141 == nil {
			var hx_zero_142 *haxe__rtti__CType
			return hx_zero_142
		}
		return hx_field_141.(*haxe__rtti__CType)
	}(cf)))
}

func haxe__rtti__CTypeTools_functionArgumentName(arg map[string]any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_if_146 *string
		if func(hx_obj_143 map[string]any) bool {
			hx_field_144 := hx_obj_143["opt"]
			if hx_field_144 == nil {
				var hx_zero_145 bool
				return hx_zero_145
			}
			return hx_field_144.(bool)
		}(arg) {
			hx_if_146 = hxrt.StringFromLiteral("?")
		} else {
			hx_if_146 = hxrt.StringFromLiteral("")
		}
		return hx_if_146
	}(), func() *string {
		var hx_if_153 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_147 map[string]any) *string {
			hx_field_148 := hx_obj_147["name"]
			if hx_field_148 == nil {
				var hx_zero_149 *string
				return hx_zero_149
			}
			return hx_field_148.(*string)
		}(arg), hxrt.StringFromLiteral("")) {
			hx_if_153 = hxrt.StringFromLiteral("")
		} else {
			hx_if_153 = hxrt.StringConcatStringPtr(func(hx_obj_150 map[string]any) *string {
				hx_field_151 := hx_obj_150["name"]
				if hx_field_151 == nil {
					var hx_zero_152 *string
					return hx_zero_152
				}
				return hx_field_151.(*string)
			}(arg), hxrt.StringFromLiteral(":"))
		}
		return hx_if_153
	}()), haxe__rtti__CTypeTools_toString(func(hx_obj_154 map[string]any) *haxe__rtti__CType {
		hx_field_155 := hx_obj_154["t"]
		if hx_field_155 == nil {
			var hx_zero_156 *haxe__rtti__CType
			return hx_zero_156
		}
		return hx_field_155.(*haxe__rtti__CType)
	}(arg))), func() *string {
		var hx_if_163 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_157 map[string]any) *string {
			hx_field_158 := hx_obj_157["value"]
			if hx_field_158 == nil {
				var hx_zero_159 *string
				return hx_zero_159
			}
			return hx_field_158.(*string)
		}(arg), nil) {
			hx_if_163 = hxrt.StringFromLiteral("")
		} else {
			hx_if_163 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" = "), func(hx_obj_160 map[string]any) *string {
				hx_field_161 := hx_obj_160["value"]
				if hx_field_161 == nil {
					var hx_zero_162 *string
					return hx_zero_162
				}
				return hx_field_161.(*string)
			}(arg))
		}
		return hx_if_163
	}())
}

func haxe__rtti__CTypeTools_joinClassFields(fields *hxrt.Array) *string {
	parts := hxrt.NewArray()
	_g := 0
	for _g < fields.Len() {
		field := func(hx_value_164 any) map[string]any {
			if hx_value_164 == nil {
				var hx_zero_165 map[string]any
				return hx_zero_165
			}
			return hx_value_164.(map[string]any)
		}(fields.Get(_g))
		_g = int(int32((_g + 1)))
		parts.Push(haxe__rtti__CTypeTools_classField(field))
	}
	return haxe__rtti__CTypeTools_joinStringArray(parts, hxrt.StringFromLiteral(", "))
}

func haxe__rtti__CTypeTools_joinFunctionArguments(args *hxrt.Array) *string {
	parts := hxrt.NewArray()
	_g := 0
	for _g < args.Len() {
		arg := func(hx_value_167 any) map[string]any {
			if hx_value_167 == nil {
				var hx_zero_168 map[string]any
				return hx_zero_168
			}
			return hx_value_167.(map[string]any)
		}(args.Get(_g))
		_g = int(int32((_g + 1)))
		parts.Push(haxe__rtti__CTypeTools_functionArgumentName(arg))
	}
	return haxe__rtti__CTypeTools_joinStringArray(parts, hxrt.StringFromLiteral(" -> "))
}

func haxe__rtti__CTypeTools_joinStringArray(parts *hxrt.Array, separator *string) *string {
	var buf_b *string
	buf_b = hxrt.StringFromLiteral("")
	_g := 0
	_g1 := parts.Len()
	for _g < _g1 {
		hx_post_170 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_170
		if i > 0 {
			buf_b = hxrt.StringConcatStringPtr(buf_b, hxrt.StdString(separator))
		}
		buf_b = hxrt.StringConcatStringPtr(buf_b, hxrt.StdString(parts.Get(i)))
	}
	return buf_b
}

func haxe__rtti__CTypeTools_nameWithParams(name *string, params *hxrt.Array) *string {
	if params.Len() == 0 {
		return name
	}
	parts := hxrt.NewArray()
	_g := 0
	for _g < params.Len() {
		param := func(hx_value_171 any) *haxe__rtti__CType {
			if hx_value_171 == nil {
				var hx_zero_172 *haxe__rtti__CType
				return hx_zero_172
			}
			return hx_value_171.(*haxe__rtti__CType)
		}(params.Get(_g))
		_g = int(int32((_g + 1)))
		parts.Push(haxe__rtti__CTypeTools_toString(param))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(name, hxrt.StringFromLiteral("<")), haxe__rtti__CTypeTools_joinStringArray(parts, hxrt.StringFromLiteral(", "))), hxrt.StringFromLiteral(">"))
}

func haxe__rtti__CTypeTools_toString(t *haxe__rtti__CType) *string {
	var hx_switch_174 *string
	switch t.tag {
	case 0:
		hx_switch_174 = hxrt.StringFromLiteral("unknown")
	case 1:
		_g := t.params[0].(*string)
		_g1 := t.params[1].(*hxrt.Array)
		name := _g
		params := _g1
		hx_switch_174 = haxe__rtti__CTypeTools_nameWithParams(name, params)
	case 2:
		_g_1 := t.params[0].(*string)
		_g1_1 := t.params[1].(*hxrt.Array)
		name_1 := _g_1
		params_1 := _g1_1
		hx_switch_174 = haxe__rtti__CTypeTools_nameWithParams(name_1, params_1)
	case 3:
		_g_2 := t.params[0].(*string)
		_g1_2 := t.params[1].(*hxrt.Array)
		name_2 := _g_2
		params_2 := _g1_2
		hx_switch_174 = haxe__rtti__CTypeTools_nameWithParams(name_2, params_2)
	case 4:
		_g_3 := t.params[0].(*hxrt.Array)
		_g1_3 := t.params[1].(*haxe__rtti__CType)
		args := _g_3
		ret := _g1_3
		var hx_if_175 *string
		if args.Len() == 0 {
			hx_if_175 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Void -> "), haxe__rtti__CTypeTools_toString(ret))
		} else {
			hx_if_175 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(haxe__rtti__CTypeTools_joinFunctionArguments(args), hxrt.StringFromLiteral(" -> ")), haxe__rtti__CTypeTools_toString(ret))
		}
		hx_switch_174 = hx_if_175
	case 5:
		_g_4 := t.params[0].(*hxrt.Array)
		fields := _g_4
		hx_switch_174 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{ "), haxe__rtti__CTypeTools_joinClassFields(fields)), hxrt.StringFromLiteral("}"))
	case 6:
		_g_5 := t.params[0].(*haxe__rtti__CType)
		d := _g_5
		var hx_if_176 *string
		if d == nil {
			hx_if_176 = hxrt.StringFromLiteral("Dynamic")
		} else {
			hx_if_176 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Dynamic<"), haxe__rtti__CTypeTools_toString(d)), hxrt.StringFromLiteral(">"))
		}
		hx_switch_174 = hx_if_176
	case 7:
		_g_6 := t.params[0].(*string)
		_g1_4 := t.params[1].(*hxrt.Array)
		name_3 := _g_6
		params_3 := _g1_4
		hx_switch_174 = haxe__rtti__CTypeTools_nameWithParams(name_3, params_3)
	}
	return hx_switch_174
}

type haxe__rtti__TypeTree struct {
	tag    int
	params []any
}

func haxe__rtti__TypeTree_TPackage(name *string, full *string, subs *hxrt.Array) *haxe__rtti__TypeTree {
	enumValue := &haxe__rtti__TypeTree{tag: 0}
	enumValue.params = []any{name, full, subs}
	return enumValue
}

func haxe__rtti__TypeTree_TClassdecl(c map[string]any) *haxe__rtti__TypeTree {
	enumValue := &haxe__rtti__TypeTree{tag: 1}
	enumValue.params = []any{c}
	return enumValue
}

func haxe__rtti__TypeTree_TEnumdecl(e map[string]any) *haxe__rtti__TypeTree {
	enumValue := &haxe__rtti__TypeTree{tag: 2}
	enumValue.params = []any{e}
	return enumValue
}

func haxe__rtti__TypeTree_TTypedecl(t map[string]any) *haxe__rtti__TypeTree {
	enumValue := &haxe__rtti__TypeTree{tag: 3}
	enumValue.params = []any{t}
	return enumValue
}

func haxe__rtti__TypeTree_TAbstractdecl(a map[string]any) *haxe__rtti__TypeTree {
	enumValue := &haxe__rtti__TypeTree{tag: 4}
	enumValue.params = []any{a}
	return enumValue
}

type haxe__rtti__Rights struct {
	tag    int
	params []any
}

var haxe__rtti__Rights_RNormal *haxe__rtti__Rights = &haxe__rtti__Rights{tag: 0}

var haxe__rtti__Rights_RNo *haxe__rtti__Rights = &haxe__rtti__Rights{tag: 1}

func haxe__rtti__Rights_RCall(m *string) *haxe__rtti__Rights {
	enumValue := &haxe__rtti__Rights{tag: 2}
	enumValue.params = []any{m}
	return enumValue
}

var haxe__rtti__Rights_RMethod *haxe__rtti__Rights = &haxe__rtti__Rights{tag: 3}

var haxe__rtti__Rights_RDynamic *haxe__rtti__Rights = &haxe__rtti__Rights{tag: 4}

var haxe__rtti__Rights_RInline *haxe__rtti__Rights = &haxe__rtti__Rights{tag: 5}

type haxe__rtti__CType struct {
	tag    int
	params []any
}

var haxe__rtti__CType_CUnknown *haxe__rtti__CType = &haxe__rtti__CType{tag: 0}

func haxe__rtti__CType_CEnum(name *string, params *hxrt.Array) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 1}
	enumValue.params = []any{name, params}
	return enumValue
}

func haxe__rtti__CType_CClass(name *string, params *hxrt.Array) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 2}
	enumValue.params = []any{name, params}
	return enumValue
}

func haxe__rtti__CType_CTypedef(name *string, params *hxrt.Array) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 3}
	enumValue.params = []any{name, params}
	return enumValue
}

func haxe__rtti__CType_CFunction(args *hxrt.Array, ret *haxe__rtti__CType) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 4}
	enumValue.params = []any{args, ret}
	return enumValue
}

func haxe__rtti__CType_CAnonymous(fields *hxrt.Array) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 5}
	enumValue.params = []any{fields}
	return enumValue
}

func haxe__rtti__CType_CDynamic(t *haxe__rtti__CType) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 6}
	enumValue.params = []any{t}
	return enumValue
}

func haxe__rtti__CType_CAbstract(name *string, params *hxrt.Array) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 7}
	enumValue.params = []any{name, params}
	return enumValue
}
