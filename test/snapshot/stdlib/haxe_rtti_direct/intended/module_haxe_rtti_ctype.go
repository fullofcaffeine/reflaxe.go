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
	var hx_if_83 bool
	if t.tag == 4 {
		_g := func(hx_value_79 any) *hxrt.Array {
			if hx_value_79 == nil {
				var hx_zero_80 *hxrt.Array
				return hx_zero_80
			}
			return hx_value_79.(*hxrt.Array)
		}(t.params[0])
		_ = _g
		_g_1 := func(hx_value_81 any) *haxe__rtti__CType {
			if hx_value_81 == nil {
				var hx_zero_82 *haxe__rtti__CType
				return hx_zero_82
			}
			return hx_value_81.(*haxe__rtti__CType)
		}(t.params[1])
		_ = _g_1
		hx_if_83 = false
	} else {
		hx_if_83 = true
	}
	return hx_if_83
}

func haxe__rtti__TypeApi_rightsEq(r1 *haxe__rtti__Rights, r2 *haxe__rtti__Rights) bool {
	if r1 == r2 {
		return true
	}
	if r1.tag == 2 {
		_g := func(hx_value_84 any) *string {
			if hx_value_84 == nil {
				var hx_zero_85 *string
				return hx_zero_85
			}
			return hx_value_84.(*string)
		}(r1.params[0])
		m1 := _g
		if r2.tag == 2 {
			_g_1 := func(hx_value_86 any) *string {
				if hx_value_86 == nil {
					var hx_zero_87 *string
					return hx_zero_87
				}
				return hx_value_86.(*string)
			}(r2.params[0])
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
		hx_post_88 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_88
		if !haxe__rtti__TypeApi_fieldEq(func(hx_value_89 any) map[string]any {
			if hx_value_89 == nil {
				var hx_zero_90 map[string]any
				return hx_zero_90
			}
			return hx_value_89.(map[string]any)
		}(l1.Get(i)), func(hx_value_91 any) map[string]any {
			if hx_value_91 == nil {
				var hx_zero_92 map[string]any
				return hx_zero_92
			}
			return hx_value_91.(map[string]any)
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
		hx_post_93 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_93
		a := func(hx_value_94 any) map[string]any {
			if hx_value_94 == nil {
				var hx_zero_95 map[string]any
				return hx_zero_95
			}
			return hx_value_94.(map[string]any)
		}(l1.Get(i))
		b := func(hx_value_96 any) map[string]any {
			if hx_value_96 == nil {
				var hx_zero_97 map[string]any
				return hx_zero_97
			}
			return hx_value_96.(map[string]any)
		}(l2.Get(i))
		if (!hxrt.StringEqualStringPtr(func(hx_obj_98 map[string]any) *string {
			hx_field_99 := hx_obj_98["name"]
			if hx_field_99 == nil {
				var hx_zero_100 *string
				return hx_zero_100
			}
			return hx_field_99.(*string)
		}(a), func(hx_obj_101 map[string]any) *string {
			hx_field_102 := hx_obj_101["name"]
			if hx_field_102 == nil {
				var hx_zero_103 *string
				return hx_zero_103
			}
			return hx_field_102.(*string)
		}(b)) || (func(hx_obj_104 map[string]any) bool {
			hx_field_105 := hx_obj_104["opt"]
			if hx_field_105 == nil {
				var hx_zero_106 bool
				return hx_zero_106
			}
			return hx_field_105.(bool)
		}(a) != func(hx_obj_107 map[string]any) bool {
			hx_field_108 := hx_obj_107["opt"]
			if hx_field_108 == nil {
				var hx_zero_109 bool
				return hx_zero_109
			}
			return hx_field_108.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_110 map[string]any) *haxe__rtti__CType {
			hx_field_111 := hx_obj_110["t"]
			if hx_field_111 == nil {
				var hx_zero_112 *haxe__rtti__CType
				return hx_zero_112
			}
			return hx_field_111.(*haxe__rtti__CType)
		}(a), func(hx_obj_113 map[string]any) *haxe__rtti__CType {
			hx_field_114 := hx_obj_113["t"]
			if hx_field_114 == nil {
				var hx_zero_115 *haxe__rtti__CType
				return hx_zero_115
			}
			return hx_field_114.(*haxe__rtti__CType)
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
		hx_post_116 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_116
		a := func(hx_value_117 any) map[string]any {
			if hx_value_117 == nil {
				var hx_zero_118 map[string]any
				return hx_zero_118
			}
			return hx_value_117.(map[string]any)
		}(l1.Get(i))
		b := func(hx_value_119 any) map[string]any {
			if hx_value_119 == nil {
				var hx_zero_120 map[string]any
				return hx_zero_120
			}
			return hx_value_119.(map[string]any)
		}(l2.Get(i))
		if (!hxrt.StringEqualStringPtr(func(hx_obj_121 map[string]any) *string {
			hx_field_122 := hx_obj_121["name"]
			if hx_field_122 == nil {
				var hx_zero_123 *string
				return hx_zero_123
			}
			return hx_field_122.(*string)
		}(a), func(hx_obj_124 map[string]any) *string {
			hx_field_125 := hx_obj_124["name"]
			if hx_field_125 == nil {
				var hx_zero_126 *string
				return hx_zero_126
			}
			return hx_field_125.(*string)
		}(b)) || (func(hx_obj_127 map[string]any) bool {
			hx_field_128 := hx_obj_127["opt"]
			if hx_field_128 == nil {
				var hx_zero_129 bool
				return hx_zero_129
			}
			return hx_field_128.(bool)
		}(a) != func(hx_obj_130 map[string]any) bool {
			hx_field_131 := hx_obj_130["opt"]
			if hx_field_131 == nil {
				var hx_zero_132 bool
				return hx_zero_132
			}
			return hx_field_131.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_133 map[string]any) *haxe__rtti__CType {
			hx_field_134 := hx_obj_133["t"]
			if hx_field_134 == nil {
				var hx_zero_135 *haxe__rtti__CType
				return hx_zero_135
			}
			return hx_field_134.(*haxe__rtti__CType)
		}(a), func(hx_obj_136 map[string]any) *haxe__rtti__CType {
			hx_field_137 := hx_obj_136["t"]
			if hx_field_137 == nil {
				var hx_zero_138 *haxe__rtti__CType
				return hx_zero_138
			}
			return hx_field_137.(*haxe__rtti__CType)
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
		hx_post_139 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_139
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
		hx_post_140 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_140
		if !haxe__rtti__TypeApi_typeEq(func(hx_value_141 any) *haxe__rtti__CType {
			if hx_value_141 == nil {
				var hx_zero_142 *haxe__rtti__CType
				return hx_zero_142
			}
			return hx_value_141.(*haxe__rtti__CType)
		}(l1.Get(i)), func(hx_value_143 any) *haxe__rtti__CType {
			if hx_value_143 == nil {
				var hx_zero_144 *haxe__rtti__CType
				return hx_zero_144
			}
			return hx_value_143.(*haxe__rtti__CType)
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
		_g := func(hx_value_145 any) *string {
			if hx_value_145 == nil {
				var hx_zero_146 *string
				return hx_zero_146
			}
			return hx_value_145.(*string)
		}(t1.params[0])
		_g1 := func(hx_value_147 any) *hxrt.Array {
			if hx_value_147 == nil {
				var hx_zero_148 *hxrt.Array
				return hx_zero_148
			}
			return hx_value_147.(*hxrt.Array)
		}(t1.params[1])
		name := _g
		params := _g1
		if t2.tag == 1 {
			_g_1 := func(hx_value_149 any) *string {
				if hx_value_149 == nil {
					var hx_zero_150 *string
					return hx_zero_150
				}
				return hx_value_149.(*string)
			}(t2.params[0])
			_g1_1 := func(hx_value_151 any) *hxrt.Array {
				if hx_value_151 == nil {
					var hx_zero_152 *hxrt.Array
					return hx_zero_152
				}
				return hx_value_151.(*hxrt.Array)
			}(t2.params[1])
			name2 := _g_1
			params2 := _g1_1
			return (hxrt.StringEqualStringPtr(name, name2) && haxe__rtti__TypeApi_sameTypes(params, params2))
		} else {
		}
	case 2:
		_g_2 := func(hx_value_153 any) *string {
			if hx_value_153 == nil {
				var hx_zero_154 *string
				return hx_zero_154
			}
			return hx_value_153.(*string)
		}(t1.params[0])
		_g1_2 := func(hx_value_155 any) *hxrt.Array {
			if hx_value_155 == nil {
				var hx_zero_156 *hxrt.Array
				return hx_zero_156
			}
			return hx_value_155.(*hxrt.Array)
		}(t1.params[1])
		name_1 := _g_2
		params_1 := _g1_2
		if t2.tag == 2 {
			_g_3 := func(hx_value_157 any) *string {
				if hx_value_157 == nil {
					var hx_zero_158 *string
					return hx_zero_158
				}
				return hx_value_157.(*string)
			}(t2.params[0])
			_g1_3 := func(hx_value_159 any) *hxrt.Array {
				if hx_value_159 == nil {
					var hx_zero_160 *hxrt.Array
					return hx_zero_160
				}
				return hx_value_159.(*hxrt.Array)
			}(t2.params[1])
			name2_1 := _g_3
			params2_1 := _g1_3
			return (hxrt.StringEqualStringPtr(name_1, name2_1) && haxe__rtti__TypeApi_sameTypes(params_1, params2_1))
		} else {
		}
	case 3:
		_g_4 := func(hx_value_161 any) *string {
			if hx_value_161 == nil {
				var hx_zero_162 *string
				return hx_zero_162
			}
			return hx_value_161.(*string)
		}(t1.params[0])
		_g1_4 := func(hx_value_163 any) *hxrt.Array {
			if hx_value_163 == nil {
				var hx_zero_164 *hxrt.Array
				return hx_zero_164
			}
			return hx_value_163.(*hxrt.Array)
		}(t1.params[1])
		name_2 := _g_4
		params_2 := _g1_4
		if t2.tag == 3 {
			_g_5 := func(hx_value_165 any) *string {
				if hx_value_165 == nil {
					var hx_zero_166 *string
					return hx_zero_166
				}
				return hx_value_165.(*string)
			}(t2.params[0])
			_g1_5 := func(hx_value_167 any) *hxrt.Array {
				if hx_value_167 == nil {
					var hx_zero_168 *hxrt.Array
					return hx_zero_168
				}
				return hx_value_167.(*hxrt.Array)
			}(t2.params[1])
			name2_2 := _g_5
			params2_2 := _g1_5
			return (hxrt.StringEqualStringPtr(name_2, name2_2) && haxe__rtti__TypeApi_sameTypes(params_2, params2_2))
		} else {
		}
	case 4:
		_g_6 := func(hx_value_169 any) *hxrt.Array {
			if hx_value_169 == nil {
				var hx_zero_170 *hxrt.Array
				return hx_zero_170
			}
			return hx_value_169.(*hxrt.Array)
		}(t1.params[0])
		_g1_6 := func(hx_value_171 any) *haxe__rtti__CType {
			if hx_value_171 == nil {
				var hx_zero_172 *haxe__rtti__CType
				return hx_zero_172
			}
			return hx_value_171.(*haxe__rtti__CType)
		}(t1.params[1])
		args := _g_6
		ret := _g1_6
		if t2.tag == 4 {
			_g_7 := func(hx_value_173 any) *hxrt.Array {
				if hx_value_173 == nil {
					var hx_zero_174 *hxrt.Array
					return hx_zero_174
				}
				return hx_value_173.(*hxrt.Array)
			}(t2.params[0])
			_g1_7 := func(hx_value_175 any) *haxe__rtti__CType {
				if hx_value_175 == nil {
					var hx_zero_176 *haxe__rtti__CType
					return hx_zero_176
				}
				return hx_value_175.(*haxe__rtti__CType)
			}(t2.params[1])
			args2 := _g_7
			ret2 := _g1_7
			return (haxe__rtti__TypeApi_sameFunctionArguments(args, args2) && haxe__rtti__TypeApi_typeEq(ret, ret2))
		} else {
		}
	case 5:
		_g_8 := func(hx_value_177 any) *hxrt.Array {
			if hx_value_177 == nil {
				var hx_zero_178 *hxrt.Array
				return hx_zero_178
			}
			return hx_value_177.(*hxrt.Array)
		}(t1.params[0])
		fields := _g_8
		if t2.tag == 5 {
			_g_9 := func(hx_value_179 any) *hxrt.Array {
				if hx_value_179 == nil {
					var hx_zero_180 *hxrt.Array
					return hx_zero_180
				}
				return hx_value_179.(*hxrt.Array)
			}(t2.params[0])
			fields2 := _g_9
			return haxe__rtti__TypeApi_sameClassFields(fields, fields2)
		} else {
		}
	case 6:
		_g_10 := func(hx_value_181 any) *haxe__rtti__CType {
			if hx_value_181 == nil {
				var hx_zero_182 *haxe__rtti__CType
				return hx_zero_182
			}
			return hx_value_181.(*haxe__rtti__CType)
		}(t1.params[0])
		t := _g_10
		if t2.tag == 6 {
			_g_11 := func(hx_value_183 any) *haxe__rtti__CType {
				if hx_value_183 == nil {
					var hx_zero_184 *haxe__rtti__CType
					return hx_zero_184
				}
				return hx_value_183.(*haxe__rtti__CType)
			}(t2.params[0])
			t2_1 := _g_11
			if (t == nil) != (t2_1 == nil) {
				return false
			}
			return ((t == nil) || haxe__rtti__TypeApi_typeEq(t, t2_1))
		} else {
		}
	case 7:
		_g_12 := func(hx_value_185 any) *string {
			if hx_value_185 == nil {
				var hx_zero_186 *string
				return hx_zero_186
			}
			return hx_value_185.(*string)
		}(t1.params[0])
		_g1_8 := func(hx_value_187 any) *hxrt.Array {
			if hx_value_187 == nil {
				var hx_zero_188 *hxrt.Array
				return hx_zero_188
			}
			return hx_value_187.(*hxrt.Array)
		}(t1.params[1])
		name_3 := _g_12
		params_3 := _g1_8
		if t2.tag == 7 {
			_g_13 := func(hx_value_189 any) *string {
				if hx_value_189 == nil {
					var hx_zero_190 *string
					return hx_zero_190
				}
				return hx_value_189.(*string)
			}(t2.params[0])
			_g1_9 := func(hx_value_191 any) *hxrt.Array {
				if hx_value_191 == nil {
					var hx_zero_192 *hxrt.Array
					return hx_zero_192
				}
				return hx_value_191.(*hxrt.Array)
			}(t2.params[1])
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
		_g := func(hx_value_193 any) *string {
			if hx_value_193 == nil {
				var hx_zero_194 *string
				return hx_zero_194
			}
			return hx_value_193.(*string)
		}(t.params[0])
		_ = _g
		_g_1 := func(hx_value_195 any) *string {
			if hx_value_195 == nil {
				var hx_zero_196 *string
				return hx_zero_196
			}
			return hx_value_195.(*string)
		}(t.params[1])
		_ = _g_1
		_g_2 := func(hx_value_197 any) *hxrt.Array {
			if hx_value_197 == nil {
				var hx_zero_198 *hxrt.Array
				return hx_zero_198
			}
			return hx_value_197.(*hxrt.Array)
		}(t.params[2])
		_ = _g_2
		hxrt.Throw(hxrt.StringFromLiteral("Unexpected Package"))
	case 1:
		_g_3 := func(hx_value_199 any) map[string]any {
			if hx_value_199 == nil {
				var hx_zero_200 map[string]any
				return hx_zero_200
			}
			return hx_value_199.(map[string]any)
		}(t.params[0])
		c := _g_3
		inf = c
	case 2:
		_g_4 := func(hx_value_201 any) map[string]any {
			if hx_value_201 == nil {
				var hx_zero_202 map[string]any
				return hx_zero_202
			}
			return hx_value_201.(map[string]any)
		}(t.params[0])
		e := _g_4
		inf = e
	case 3:
		_g_5 := func(hx_value_203 any) map[string]any {
			if hx_value_203 == nil {
				var hx_zero_204 map[string]any
				return hx_zero_204
			}
			return hx_value_203.(map[string]any)
		}(t.params[0])
		t_1 := _g_5
		inf = t_1
	case 4:
		_g_6 := func(hx_value_205 any) map[string]any {
			if hx_value_205 == nil {
				var hx_zero_206 map[string]any
				return hx_zero_206
			}
			return hx_value_205.(map[string]any)
		}(t.params[0])
		a := _g_6
		inf = a
	}
	return inf
}

func haxe__rtti__CTypeTools_classField(cf map[string]any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func(hx_obj_207 map[string]any) *string {
		hx_field_208 := hx_obj_207["name"]
		if hx_field_208 == nil {
			var hx_zero_209 *string
			return hx_zero_209
		}
		return hx_field_208.(*string)
	}(cf), hxrt.StringFromLiteral(":")), haxe__rtti__CTypeTools_toString(func(hx_obj_210 map[string]any) *haxe__rtti__CType {
		hx_field_211 := hx_obj_210["type"]
		if hx_field_211 == nil {
			var hx_zero_212 *haxe__rtti__CType
			return hx_zero_212
		}
		return hx_field_211.(*haxe__rtti__CType)
	}(cf)))
}

func haxe__rtti__CTypeTools_functionArgumentName(arg map[string]any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_if_216 *string
		if func(hx_obj_213 map[string]any) bool {
			hx_field_214 := hx_obj_213["opt"]
			if hx_field_214 == nil {
				var hx_zero_215 bool
				return hx_zero_215
			}
			return hx_field_214.(bool)
		}(arg) {
			hx_if_216 = hxrt.StringFromLiteral("?")
		} else {
			hx_if_216 = hxrt.StringFromLiteral("")
		}
		return hx_if_216
	}(), func() *string {
		var hx_if_223 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_217 map[string]any) *string {
			hx_field_218 := hx_obj_217["name"]
			if hx_field_218 == nil {
				var hx_zero_219 *string
				return hx_zero_219
			}
			return hx_field_218.(*string)
		}(arg), hxrt.StringFromLiteral("")) {
			hx_if_223 = hxrt.StringFromLiteral("")
		} else {
			hx_if_223 = hxrt.StringConcatStringPtr(func(hx_obj_220 map[string]any) *string {
				hx_field_221 := hx_obj_220["name"]
				if hx_field_221 == nil {
					var hx_zero_222 *string
					return hx_zero_222
				}
				return hx_field_221.(*string)
			}(arg), hxrt.StringFromLiteral(":"))
		}
		return hx_if_223
	}()), haxe__rtti__CTypeTools_toString(func(hx_obj_224 map[string]any) *haxe__rtti__CType {
		hx_field_225 := hx_obj_224["t"]
		if hx_field_225 == nil {
			var hx_zero_226 *haxe__rtti__CType
			return hx_zero_226
		}
		return hx_field_225.(*haxe__rtti__CType)
	}(arg))), func() *string {
		var hx_if_233 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_227 map[string]any) *string {
			hx_field_228 := hx_obj_227["value"]
			if hx_field_228 == nil {
				var hx_zero_229 *string
				return hx_zero_229
			}
			return hx_field_228.(*string)
		}(arg), nil) {
			hx_if_233 = hxrt.StringFromLiteral("")
		} else {
			hx_if_233 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" = "), func(hx_obj_230 map[string]any) *string {
				hx_field_231 := hx_obj_230["value"]
				if hx_field_231 == nil {
					var hx_zero_232 *string
					return hx_zero_232
				}
				return hx_field_231.(*string)
			}(arg))
		}
		return hx_if_233
	}())
}

func haxe__rtti__CTypeTools_joinClassFields(fields *hxrt.Array) *string {
	parts := hxrt.NewArray()
	_g := 0
	for _g < fields.Len() {
		field := func(hx_value_234 any) map[string]any {
			if hx_value_234 == nil {
				var hx_zero_235 map[string]any
				return hx_zero_235
			}
			return hx_value_234.(map[string]any)
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
		arg := func(hx_value_237 any) map[string]any {
			if hx_value_237 == nil {
				var hx_zero_238 map[string]any
				return hx_zero_238
			}
			return hx_value_237.(map[string]any)
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
		hx_post_240 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_240
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
		param := func(hx_value_241 any) *haxe__rtti__CType {
			if hx_value_241 == nil {
				var hx_zero_242 *haxe__rtti__CType
				return hx_zero_242
			}
			return hx_value_241.(*haxe__rtti__CType)
		}(params.Get(_g))
		_g = int(int32((_g + 1)))
		parts.Push(haxe__rtti__CTypeTools_toString(param))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(name, hxrt.StringFromLiteral("<")), haxe__rtti__CTypeTools_joinStringArray(parts, hxrt.StringFromLiteral(", "))), hxrt.StringFromLiteral(">"))
}

func haxe__rtti__CTypeTools_toString(t *haxe__rtti__CType) *string {
	var hx_switch_244 *string
	switch t.tag {
	case 0:
		hx_switch_244 = hxrt.StringFromLiteral("unknown")
	case 1:
		_g := func(hx_value_245 any) *string {
			if hx_value_245 == nil {
				var hx_zero_246 *string
				return hx_zero_246
			}
			return hx_value_245.(*string)
		}(t.params[0])
		_g1 := func(hx_value_247 any) *hxrt.Array {
			if hx_value_247 == nil {
				var hx_zero_248 *hxrt.Array
				return hx_zero_248
			}
			return hx_value_247.(*hxrt.Array)
		}(t.params[1])
		name := _g
		params := _g1
		hx_switch_244 = haxe__rtti__CTypeTools_nameWithParams(name, params)
	case 2:
		_g_1 := func(hx_value_249 any) *string {
			if hx_value_249 == nil {
				var hx_zero_250 *string
				return hx_zero_250
			}
			return hx_value_249.(*string)
		}(t.params[0])
		_g1_1 := func(hx_value_251 any) *hxrt.Array {
			if hx_value_251 == nil {
				var hx_zero_252 *hxrt.Array
				return hx_zero_252
			}
			return hx_value_251.(*hxrt.Array)
		}(t.params[1])
		name_1 := _g_1
		params_1 := _g1_1
		hx_switch_244 = haxe__rtti__CTypeTools_nameWithParams(name_1, params_1)
	case 3:
		_g_2 := func(hx_value_253 any) *string {
			if hx_value_253 == nil {
				var hx_zero_254 *string
				return hx_zero_254
			}
			return hx_value_253.(*string)
		}(t.params[0])
		_g1_2 := func(hx_value_255 any) *hxrt.Array {
			if hx_value_255 == nil {
				var hx_zero_256 *hxrt.Array
				return hx_zero_256
			}
			return hx_value_255.(*hxrt.Array)
		}(t.params[1])
		name_2 := _g_2
		params_2 := _g1_2
		hx_switch_244 = haxe__rtti__CTypeTools_nameWithParams(name_2, params_2)
	case 4:
		_g_3 := func(hx_value_257 any) *hxrt.Array {
			if hx_value_257 == nil {
				var hx_zero_258 *hxrt.Array
				return hx_zero_258
			}
			return hx_value_257.(*hxrt.Array)
		}(t.params[0])
		_g1_3 := func(hx_value_259 any) *haxe__rtti__CType {
			if hx_value_259 == nil {
				var hx_zero_260 *haxe__rtti__CType
				return hx_zero_260
			}
			return hx_value_259.(*haxe__rtti__CType)
		}(t.params[1])
		args := _g_3
		ret := _g1_3
		var hx_if_261 *string
		if args.Len() == 0 {
			hx_if_261 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Void -> "), haxe__rtti__CTypeTools_toString(ret))
		} else {
			hx_if_261 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(haxe__rtti__CTypeTools_joinFunctionArguments(args), hxrt.StringFromLiteral(" -> ")), haxe__rtti__CTypeTools_toString(ret))
		}
		hx_switch_244 = hx_if_261
	case 5:
		_g_4 := func(hx_value_262 any) *hxrt.Array {
			if hx_value_262 == nil {
				var hx_zero_263 *hxrt.Array
				return hx_zero_263
			}
			return hx_value_262.(*hxrt.Array)
		}(t.params[0])
		fields := _g_4
		hx_switch_244 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{ "), haxe__rtti__CTypeTools_joinClassFields(fields)), hxrt.StringFromLiteral("}"))
	case 6:
		_g_5 := func(hx_value_264 any) *haxe__rtti__CType {
			if hx_value_264 == nil {
				var hx_zero_265 *haxe__rtti__CType
				return hx_zero_265
			}
			return hx_value_264.(*haxe__rtti__CType)
		}(t.params[0])
		d := _g_5
		var hx_if_266 *string
		if d == nil {
			hx_if_266 = hxrt.StringFromLiteral("Dynamic")
		} else {
			hx_if_266 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Dynamic<"), haxe__rtti__CTypeTools_toString(d)), hxrt.StringFromLiteral(">"))
		}
		hx_switch_244 = hx_if_266
	case 7:
		_g_6 := func(hx_value_267 any) *string {
			if hx_value_267 == nil {
				var hx_zero_268 *string
				return hx_zero_268
			}
			return hx_value_267.(*string)
		}(t.params[0])
		_g1_4 := func(hx_value_269 any) *hxrt.Array {
			if hx_value_269 == nil {
				var hx_zero_270 *hxrt.Array
				return hx_zero_270
			}
			return hx_value_269.(*hxrt.Array)
		}(t.params[1])
		name_3 := _g_6
		params_3 := _g1_4
		hx_switch_244 = haxe__rtti__CTypeTools_nameWithParams(name_3, params_3)
	}
	return hx_switch_244
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
