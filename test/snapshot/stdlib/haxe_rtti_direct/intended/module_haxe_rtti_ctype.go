package main

import "snapshot/hxrt"

func haxe__rtti__TypeApi_constructorEq(c1 hxrt.Anonymous, c2 hxrt.Anonymous) bool {
	if !hxrt.StringEqualStringPtr(func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(hxrt.AnonymousGet(c1, "name")), func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(hxrt.AnonymousGet(c2, "name"))) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_value_5 any) *string {
		if hx_value_5 == nil {
			var hx_zero_6 *string
			return hx_zero_6
		}
		return hx_value_5.(*string)
	}(hxrt.AnonymousGet(c1, "doc")), func(hx_value_7 any) *string {
		if hx_value_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_value_7.(*string)
	}(hxrt.AnonymousGet(c2, "doc"))) {
		return false
	}
	if (func(hx_value_9 any) *hxrt.Array {
		if hx_value_9 == nil {
			var hx_zero_10 *hxrt.Array
			return hx_zero_10
		}
		return hx_value_9.(*hxrt.Array)
	}(hxrt.AnonymousGet(c1, "args")) == nil) != (func(hx_value_11 any) *hxrt.Array {
		if hx_value_11 == nil {
			var hx_zero_12 *hxrt.Array
			return hx_zero_12
		}
		return hx_value_11.(*hxrt.Array)
	}(hxrt.AnonymousGet(c2, "args")) == nil) {
		return false
	}
	if (func(hx_value_13 any) *hxrt.Array {
		if hx_value_13 == nil {
			var hx_zero_14 *hxrt.Array
			return hx_zero_14
		}
		return hx_value_13.(*hxrt.Array)
	}(hxrt.AnonymousGet(c1, "args")) != nil) && !haxe__rtti__TypeApi_sameConstructorArguments(func(hx_value_15 any) *hxrt.Array {
		if hx_value_15 == nil {
			var hx_zero_16 *hxrt.Array
			return hx_zero_16
		}
		return hx_value_15.(*hxrt.Array)
	}(hxrt.AnonymousGet(c1, "args")), func(hx_value_17 any) *hxrt.Array {
		if hx_value_17 == nil {
			var hx_zero_18 *hxrt.Array
			return hx_zero_18
		}
		return hx_value_17.(*hxrt.Array)
	}(hxrt.AnonymousGet(c2, "args"))) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_fieldEq(f1 hxrt.Anonymous, f2 hxrt.Anonymous) bool {
	if !hxrt.StringEqualStringPtr(func(hx_value_19 any) *string {
		if hx_value_19 == nil {
			var hx_zero_20 *string
			return hx_zero_20
		}
		return hx_value_19.(*string)
	}(hxrt.AnonymousGet(f1, "name")), func(hx_value_21 any) *string {
		if hx_value_21 == nil {
			var hx_zero_22 *string
			return hx_zero_22
		}
		return hx_value_21.(*string)
	}(hxrt.AnonymousGet(f2, "name"))) {
		return false
	}
	if !haxe__rtti__TypeApi_typeEq(func(hx_value_23 any) *haxe__rtti__CType {
		if hx_value_23 == nil {
			var hx_zero_24 *haxe__rtti__CType
			return hx_zero_24
		}
		return hx_value_23.(*haxe__rtti__CType)
	}(hxrt.AnonymousGet(f1, "type")), func(hx_value_25 any) *haxe__rtti__CType {
		if hx_value_25 == nil {
			var hx_zero_26 *haxe__rtti__CType
			return hx_zero_26
		}
		return hx_value_25.(*haxe__rtti__CType)
	}(hxrt.AnonymousGet(f2, "type"))) {
		return false
	}
	if func(hx_value_27 any) bool {
		if hx_value_27 == nil {
			var hx_zero_28 bool
			return hx_zero_28
		}
		return hx_value_27.(bool)
	}(hxrt.AnonymousGet(f1, "isPublic")) != func(hx_value_29 any) bool {
		if hx_value_29 == nil {
			var hx_zero_30 bool
			return hx_zero_30
		}
		return hx_value_29.(bool)
	}(hxrt.AnonymousGet(f2, "isPublic")) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_value_31 any) *string {
		if hx_value_31 == nil {
			var hx_zero_32 *string
			return hx_zero_32
		}
		return hx_value_31.(*string)
	}(hxrt.AnonymousGet(f1, "doc")), func(hx_value_33 any) *string {
		if hx_value_33 == nil {
			var hx_zero_34 *string
			return hx_zero_34
		}
		return hx_value_33.(*string)
	}(hxrt.AnonymousGet(f2, "doc"))) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_value_35 any) *haxe__rtti__Rights {
		if hx_value_35 == nil {
			var hx_zero_36 *haxe__rtti__Rights
			return hx_zero_36
		}
		return hx_value_35.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f1, "get")), func(hx_value_37 any) *haxe__rtti__Rights {
		if hx_value_37 == nil {
			var hx_zero_38 *haxe__rtti__Rights
			return hx_zero_38
		}
		return hx_value_37.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f2, "get"))) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_value_39 any) *haxe__rtti__Rights {
		if hx_value_39 == nil {
			var hx_zero_40 *haxe__rtti__Rights
			return hx_zero_40
		}
		return hx_value_39.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f1, "set")), func(hx_value_41 any) *haxe__rtti__Rights {
		if hx_value_41 == nil {
			var hx_zero_42 *haxe__rtti__Rights
			return hx_zero_42
		}
		return hx_value_41.(*haxe__rtti__Rights)
	}(hxrt.AnonymousGet(f2, "set"))) {
		return false
	}
	if (func(hx_value_43 any) *hxrt.Array {
		if hx_value_43 == nil {
			var hx_zero_44 *hxrt.Array
			return hx_zero_44
		}
		return hx_value_43.(*hxrt.Array)
	}(hxrt.AnonymousGet(f1, "params")) == nil) != (func(hx_value_45 any) *hxrt.Array {
		if hx_value_45 == nil {
			var hx_zero_46 *hxrt.Array
			return hx_zero_46
		}
		return hx_value_45.(*hxrt.Array)
	}(hxrt.AnonymousGet(f2, "params")) == nil) {
		return false
	}
	if (func(hx_value_47 any) *hxrt.Array {
		if hx_value_47 == nil {
			var hx_zero_48 *hxrt.Array
			return hx_zero_48
		}
		return hx_value_47.(*hxrt.Array)
	}(hxrt.AnonymousGet(f1, "params")) != nil) && !haxe__rtti__TypeApi_sameTypeParamNames(func(hx_value_49 any) *hxrt.Array {
		if hx_value_49 == nil {
			var hx_zero_50 *hxrt.Array
			return hx_zero_50
		}
		return hx_value_49.(*hxrt.Array)
	}(hxrt.AnonymousGet(f1, "params")), func(hx_value_51 any) *hxrt.Array {
		if hx_value_51 == nil {
			var hx_zero_52 *hxrt.Array
			return hx_zero_52
		}
		return hx_value_51.(*hxrt.Array)
	}(hxrt.AnonymousGet(f2, "params"))) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_isVar(t *haxe__rtti__CType) bool {
	var hx_if_57 bool
	if t.tag == 4 {
		_g := func(hx_value_53 any) *hxrt.Array {
			if hx_value_53 == nil {
				var hx_zero_54 *hxrt.Array
				return hx_zero_54
			}
			return hx_value_53.(*hxrt.Array)
		}(t.params[0])
		_ = _g
		_g_1 := func(hx_value_55 any) *haxe__rtti__CType {
			if hx_value_55 == nil {
				var hx_zero_56 *haxe__rtti__CType
				return hx_zero_56
			}
			return hx_value_55.(*haxe__rtti__CType)
		}(t.params[1])
		_ = _g_1
		hx_if_57 = false
	} else {
		hx_if_57 = true
	}
	return hx_if_57
}

func haxe__rtti__TypeApi_rightsEq(r1 *haxe__rtti__Rights, r2 *haxe__rtti__Rights) bool {
	if r1 == r2 {
		return true
	}
	if r1.tag == 2 {
		_g := func(hx_value_58 any) *string {
			if hx_value_58 == nil {
				var hx_zero_59 *string
				return hx_zero_59
			}
			return hx_value_58.(*string)
		}(r1.params[0])
		m1 := _g
		if r2.tag == 2 {
			_g_1 := func(hx_value_60 any) *string {
				if hx_value_60 == nil {
					var hx_zero_61 *string
					return hx_zero_61
				}
				return hx_value_60.(*string)
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
		hx_post_62 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_62
		if !haxe__rtti__TypeApi_fieldEq(hx_structural_recover(l1.Get(i)), hx_structural_recover(l2.Get(i))) {
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
		hx_post_63 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_63
		a := hx_structural_recover(l1.Get(i))
		b := hx_structural_recover(l2.Get(i))
		if (!hxrt.StringEqualStringPtr(func(hx_value_64 any) *string {
			if hx_value_64 == nil {
				var hx_zero_65 *string
				return hx_zero_65
			}
			return hx_value_64.(*string)
		}(hxrt.AnonymousGet(a, "name")), func(hx_value_66 any) *string {
			if hx_value_66 == nil {
				var hx_zero_67 *string
				return hx_zero_67
			}
			return hx_value_66.(*string)
		}(hxrt.AnonymousGet(b, "name"))) || (func(hx_value_68 any) bool {
			if hx_value_68 == nil {
				var hx_zero_69 bool
				return hx_zero_69
			}
			return hx_value_68.(bool)
		}(hxrt.AnonymousGet(a, "opt")) != func(hx_value_70 any) bool {
			if hx_value_70 == nil {
				var hx_zero_71 bool
				return hx_zero_71
			}
			return hx_value_70.(bool)
		}(hxrt.AnonymousGet(b, "opt")))) || !haxe__rtti__TypeApi_typeEq(func(hx_value_72 any) *haxe__rtti__CType {
			if hx_value_72 == nil {
				var hx_zero_73 *haxe__rtti__CType
				return hx_zero_73
			}
			return hx_value_72.(*haxe__rtti__CType)
		}(hxrt.AnonymousGet(a, "t")), func(hx_value_74 any) *haxe__rtti__CType {
			if hx_value_74 == nil {
				var hx_zero_75 *haxe__rtti__CType
				return hx_zero_75
			}
			return hx_value_74.(*haxe__rtti__CType)
		}(hxrt.AnonymousGet(b, "t"))) {
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
		hx_post_76 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_76
		a := hx_structural_recover(l1.Get(i))
		b := hx_structural_recover(l2.Get(i))
		if (!hxrt.StringEqualStringPtr(func(hx_value_77 any) *string {
			if hx_value_77 == nil {
				var hx_zero_78 *string
				return hx_zero_78
			}
			return hx_value_77.(*string)
		}(hxrt.AnonymousGet(a, "name")), func(hx_value_79 any) *string {
			if hx_value_79 == nil {
				var hx_zero_80 *string
				return hx_zero_80
			}
			return hx_value_79.(*string)
		}(hxrt.AnonymousGet(b, "name"))) || (func(hx_value_81 any) bool {
			if hx_value_81 == nil {
				var hx_zero_82 bool
				return hx_zero_82
			}
			return hx_value_81.(bool)
		}(hxrt.AnonymousGet(a, "opt")) != func(hx_value_83 any) bool {
			if hx_value_83 == nil {
				var hx_zero_84 bool
				return hx_zero_84
			}
			return hx_value_83.(bool)
		}(hxrt.AnonymousGet(b, "opt")))) || !haxe__rtti__TypeApi_typeEq(func(hx_value_85 any) *haxe__rtti__CType {
			if hx_value_85 == nil {
				var hx_zero_86 *haxe__rtti__CType
				return hx_zero_86
			}
			return hx_value_85.(*haxe__rtti__CType)
		}(hxrt.AnonymousGet(a, "t")), func(hx_value_87 any) *haxe__rtti__CType {
			if hx_value_87 == nil {
				var hx_zero_88 *haxe__rtti__CType
				return hx_zero_88
			}
			return hx_value_87.(*haxe__rtti__CType)
		}(hxrt.AnonymousGet(b, "t"))) {
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
		hx_post_89 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_89
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
		hx_post_90 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_90
		if !haxe__rtti__TypeApi_typeEq(func(hx_value_91 any) *haxe__rtti__CType {
			if hx_value_91 == nil {
				var hx_zero_92 *haxe__rtti__CType
				return hx_zero_92
			}
			return hx_value_91.(*haxe__rtti__CType)
		}(l1.Get(i)), func(hx_value_93 any) *haxe__rtti__CType {
			if hx_value_93 == nil {
				var hx_zero_94 *haxe__rtti__CType
				return hx_zero_94
			}
			return hx_value_93.(*haxe__rtti__CType)
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
		_g := func(hx_value_95 any) *string {
			if hx_value_95 == nil {
				var hx_zero_96 *string
				return hx_zero_96
			}
			return hx_value_95.(*string)
		}(t1.params[0])
		_g1 := func(hx_value_97 any) *hxrt.Array {
			if hx_value_97 == nil {
				var hx_zero_98 *hxrt.Array
				return hx_zero_98
			}
			return hx_value_97.(*hxrt.Array)
		}(t1.params[1])
		name := _g
		params := _g1
		if t2.tag == 1 {
			_g_1 := func(hx_value_99 any) *string {
				if hx_value_99 == nil {
					var hx_zero_100 *string
					return hx_zero_100
				}
				return hx_value_99.(*string)
			}(t2.params[0])
			_g1_1 := func(hx_value_101 any) *hxrt.Array {
				if hx_value_101 == nil {
					var hx_zero_102 *hxrt.Array
					return hx_zero_102
				}
				return hx_value_101.(*hxrt.Array)
			}(t2.params[1])
			name2 := _g_1
			params2 := _g1_1
			return (hxrt.StringEqualStringPtr(name, name2) && haxe__rtti__TypeApi_sameTypes(params, params2))
		} else {
		}
	case 2:
		_g_2 := func(hx_value_103 any) *string {
			if hx_value_103 == nil {
				var hx_zero_104 *string
				return hx_zero_104
			}
			return hx_value_103.(*string)
		}(t1.params[0])
		_g1_2 := func(hx_value_105 any) *hxrt.Array {
			if hx_value_105 == nil {
				var hx_zero_106 *hxrt.Array
				return hx_zero_106
			}
			return hx_value_105.(*hxrt.Array)
		}(t1.params[1])
		name_1 := _g_2
		params_1 := _g1_2
		if t2.tag == 2 {
			_g_3 := func(hx_value_107 any) *string {
				if hx_value_107 == nil {
					var hx_zero_108 *string
					return hx_zero_108
				}
				return hx_value_107.(*string)
			}(t2.params[0])
			_g1_3 := func(hx_value_109 any) *hxrt.Array {
				if hx_value_109 == nil {
					var hx_zero_110 *hxrt.Array
					return hx_zero_110
				}
				return hx_value_109.(*hxrt.Array)
			}(t2.params[1])
			name2_1 := _g_3
			params2_1 := _g1_3
			return (hxrt.StringEqualStringPtr(name_1, name2_1) && haxe__rtti__TypeApi_sameTypes(params_1, params2_1))
		} else {
		}
	case 3:
		_g_4 := func(hx_value_111 any) *string {
			if hx_value_111 == nil {
				var hx_zero_112 *string
				return hx_zero_112
			}
			return hx_value_111.(*string)
		}(t1.params[0])
		_g1_4 := func(hx_value_113 any) *hxrt.Array {
			if hx_value_113 == nil {
				var hx_zero_114 *hxrt.Array
				return hx_zero_114
			}
			return hx_value_113.(*hxrt.Array)
		}(t1.params[1])
		name_2 := _g_4
		params_2 := _g1_4
		if t2.tag == 3 {
			_g_5 := func(hx_value_115 any) *string {
				if hx_value_115 == nil {
					var hx_zero_116 *string
					return hx_zero_116
				}
				return hx_value_115.(*string)
			}(t2.params[0])
			_g1_5 := func(hx_value_117 any) *hxrt.Array {
				if hx_value_117 == nil {
					var hx_zero_118 *hxrt.Array
					return hx_zero_118
				}
				return hx_value_117.(*hxrt.Array)
			}(t2.params[1])
			name2_2 := _g_5
			params2_2 := _g1_5
			return (hxrt.StringEqualStringPtr(name_2, name2_2) && haxe__rtti__TypeApi_sameTypes(params_2, params2_2))
		} else {
		}
	case 4:
		_g_6 := func(hx_value_119 any) *hxrt.Array {
			if hx_value_119 == nil {
				var hx_zero_120 *hxrt.Array
				return hx_zero_120
			}
			return hx_value_119.(*hxrt.Array)
		}(t1.params[0])
		_g1_6 := func(hx_value_121 any) *haxe__rtti__CType {
			if hx_value_121 == nil {
				var hx_zero_122 *haxe__rtti__CType
				return hx_zero_122
			}
			return hx_value_121.(*haxe__rtti__CType)
		}(t1.params[1])
		args := _g_6
		ret := _g1_6
		if t2.tag == 4 {
			_g_7 := func(hx_value_123 any) *hxrt.Array {
				if hx_value_123 == nil {
					var hx_zero_124 *hxrt.Array
					return hx_zero_124
				}
				return hx_value_123.(*hxrt.Array)
			}(t2.params[0])
			_g1_7 := func(hx_value_125 any) *haxe__rtti__CType {
				if hx_value_125 == nil {
					var hx_zero_126 *haxe__rtti__CType
					return hx_zero_126
				}
				return hx_value_125.(*haxe__rtti__CType)
			}(t2.params[1])
			args2 := _g_7
			ret2 := _g1_7
			return (haxe__rtti__TypeApi_sameFunctionArguments(args, args2) && haxe__rtti__TypeApi_typeEq(ret, ret2))
		} else {
		}
	case 5:
		_g_8 := func(hx_value_127 any) *hxrt.Array {
			if hx_value_127 == nil {
				var hx_zero_128 *hxrt.Array
				return hx_zero_128
			}
			return hx_value_127.(*hxrt.Array)
		}(t1.params[0])
		fields := _g_8
		if t2.tag == 5 {
			_g_9 := func(hx_value_129 any) *hxrt.Array {
				if hx_value_129 == nil {
					var hx_zero_130 *hxrt.Array
					return hx_zero_130
				}
				return hx_value_129.(*hxrt.Array)
			}(t2.params[0])
			fields2 := _g_9
			return haxe__rtti__TypeApi_sameClassFields(fields, fields2)
		} else {
		}
	case 6:
		_g_10 := func(hx_value_131 any) *haxe__rtti__CType {
			if hx_value_131 == nil {
				var hx_zero_132 *haxe__rtti__CType
				return hx_zero_132
			}
			return hx_value_131.(*haxe__rtti__CType)
		}(t1.params[0])
		t := _g_10
		if t2.tag == 6 {
			_g_11 := func(hx_value_133 any) *haxe__rtti__CType {
				if hx_value_133 == nil {
					var hx_zero_134 *haxe__rtti__CType
					return hx_zero_134
				}
				return hx_value_133.(*haxe__rtti__CType)
			}(t2.params[0])
			t2_1 := _g_11
			if (t == nil) != (t2_1 == nil) {
				return false
			}
			return ((t == nil) || haxe__rtti__TypeApi_typeEq(t, t2_1))
		} else {
		}
	case 7:
		_g_12 := func(hx_value_135 any) *string {
			if hx_value_135 == nil {
				var hx_zero_136 *string
				return hx_zero_136
			}
			return hx_value_135.(*string)
		}(t1.params[0])
		_g1_8 := func(hx_value_137 any) *hxrt.Array {
			if hx_value_137 == nil {
				var hx_zero_138 *hxrt.Array
				return hx_zero_138
			}
			return hx_value_137.(*hxrt.Array)
		}(t1.params[1])
		name_3 := _g_12
		params_3 := _g1_8
		if t2.tag == 7 {
			_g_13 := func(hx_value_139 any) *string {
				if hx_value_139 == nil {
					var hx_zero_140 *string
					return hx_zero_140
				}
				return hx_value_139.(*string)
			}(t2.params[0])
			_g1_9 := func(hx_value_141 any) *hxrt.Array {
				if hx_value_141 == nil {
					var hx_zero_142 *hxrt.Array
					return hx_zero_142
				}
				return hx_value_141.(*hxrt.Array)
			}(t2.params[1])
			name2_3 := _g_13
			params2_3 := _g1_9
			return (hxrt.StringEqualStringPtr(name_3, name2_3) && haxe__rtti__TypeApi_sameTypes(params_3, params2_3))
		} else {
		}
	}
	return false
}

func haxe__rtti__TypeApi_typeInfos(t *haxe__rtti__TypeTree) hxrt.Anonymous {
	var inf hxrt.Anonymous
	switch t.tag {
	case 0:
		_g := func(hx_value_143 any) *string {
			if hx_value_143 == nil {
				var hx_zero_144 *string
				return hx_zero_144
			}
			return hx_value_143.(*string)
		}(t.params[0])
		_ = _g
		_g_1 := func(hx_value_145 any) *string {
			if hx_value_145 == nil {
				var hx_zero_146 *string
				return hx_zero_146
			}
			return hx_value_145.(*string)
		}(t.params[1])
		_ = _g_1
		_g_2 := func(hx_value_147 any) *hxrt.Array {
			if hx_value_147 == nil {
				var hx_zero_148 *hxrt.Array
				return hx_zero_148
			}
			return hx_value_147.(*hxrt.Array)
		}(t.params[2])
		_ = _g_2
		hxrt.Throw(hxrt.StringFromLiteral("Unexpected Package"))
	case 1:
		_g_3 := hx_structural_recover(t.params[0])
		c := _g_3
		inf = c
	case 2:
		_g_4 := hx_structural_recover(t.params[0])
		e := _g_4
		inf = e
	case 3:
		_g_5 := hx_structural_recover(t.params[0])
		t_1 := _g_5
		inf = t_1
	case 4:
		_g_6 := hx_structural_recover(t.params[0])
		a := _g_6
		inf = a
	}
	return inf
}

func haxe__rtti__CTypeTools_classField(cf hxrt.Anonymous) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func(hx_value_149 any) *string {
		if hx_value_149 == nil {
			var hx_zero_150 *string
			return hx_zero_150
		}
		return hx_value_149.(*string)
	}(hxrt.AnonymousGet(cf, "name")), hxrt.StringFromLiteral(":")), haxe__rtti__CTypeTools_toString(func(hx_value_151 any) *haxe__rtti__CType {
		if hx_value_151 == nil {
			var hx_zero_152 *haxe__rtti__CType
			return hx_zero_152
		}
		return hx_value_151.(*haxe__rtti__CType)
	}(hxrt.AnonymousGet(cf, "type"))))
}

func haxe__rtti__CTypeTools_functionArgumentName(arg hxrt.Anonymous) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_if_155 *string
		if func(hx_value_153 any) bool {
			if hx_value_153 == nil {
				var hx_zero_154 bool
				return hx_zero_154
			}
			return hx_value_153.(bool)
		}(hxrt.AnonymousGet(arg, "opt")) {
			hx_if_155 = hxrt.StringFromLiteral("?")
		} else {
			hx_if_155 = hxrt.StringFromLiteral("")
		}
		return hx_if_155
	}(), func() *string {
		var hx_if_160 *string
		if hxrt.StringEqualStringPtr(func(hx_value_156 any) *string {
			if hx_value_156 == nil {
				var hx_zero_157 *string
				return hx_zero_157
			}
			return hx_value_156.(*string)
		}(hxrt.AnonymousGet(arg, "name")), hxrt.StringFromLiteral("")) {
			hx_if_160 = hxrt.StringFromLiteral("")
		} else {
			hx_if_160 = hxrt.StringConcatStringPtr(func(hx_value_158 any) *string {
				if hx_value_158 == nil {
					var hx_zero_159 *string
					return hx_zero_159
				}
				return hx_value_158.(*string)
			}(hxrt.AnonymousGet(arg, "name")), hxrt.StringFromLiteral(":"))
		}
		return hx_if_160
	}()), haxe__rtti__CTypeTools_toString(func(hx_value_161 any) *haxe__rtti__CType {
		if hx_value_161 == nil {
			var hx_zero_162 *haxe__rtti__CType
			return hx_zero_162
		}
		return hx_value_161.(*haxe__rtti__CType)
	}(hxrt.AnonymousGet(arg, "t")))), func() *string {
		var hx_if_167 *string
		if hxrt.StringEqualStringPtr(func(hx_value_163 any) *string {
			if hx_value_163 == nil {
				var hx_zero_164 *string
				return hx_zero_164
			}
			return hx_value_163.(*string)
		}(hxrt.AnonymousGet(arg, "value")), nil) {
			hx_if_167 = hxrt.StringFromLiteral("")
		} else {
			hx_if_167 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" = "), func(hx_value_165 any) *string {
				if hx_value_165 == nil {
					var hx_zero_166 *string
					return hx_zero_166
				}
				return hx_value_165.(*string)
			}(hxrt.AnonymousGet(arg, "value")))
		}
		return hx_if_167
	}())
}

func haxe__rtti__CTypeTools_joinClassFields(fields *hxrt.Array) *string {
	parts := hxrt.NewArray()
	_g := 0
	for _g < fields.Len() {
		field := hx_structural_recover(fields.Get(_g))
		_g = int(int32((_g + 1)))
		parts.Push(haxe__rtti__CTypeTools_classField(field))
	}
	return haxe__rtti__CTypeTools_joinStringArray(parts, hxrt.StringFromLiteral(", "))
}

func haxe__rtti__CTypeTools_joinFunctionArguments(args *hxrt.Array) *string {
	parts := hxrt.NewArray()
	_g := 0
	for _g < args.Len() {
		arg := hx_structural_recover(args.Get(_g))
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
		_g := func(hx_value_175 any) *string {
			if hx_value_175 == nil {
				var hx_zero_176 *string
				return hx_zero_176
			}
			return hx_value_175.(*string)
		}(t.params[0])
		_g1 := func(hx_value_177 any) *hxrt.Array {
			if hx_value_177 == nil {
				var hx_zero_178 *hxrt.Array
				return hx_zero_178
			}
			return hx_value_177.(*hxrt.Array)
		}(t.params[1])
		name := _g
		params := _g1
		hx_switch_174 = haxe__rtti__CTypeTools_nameWithParams(name, params)
	case 2:
		_g_1 := func(hx_value_179 any) *string {
			if hx_value_179 == nil {
				var hx_zero_180 *string
				return hx_zero_180
			}
			return hx_value_179.(*string)
		}(t.params[0])
		_g1_1 := func(hx_value_181 any) *hxrt.Array {
			if hx_value_181 == nil {
				var hx_zero_182 *hxrt.Array
				return hx_zero_182
			}
			return hx_value_181.(*hxrt.Array)
		}(t.params[1])
		name_1 := _g_1
		params_1 := _g1_1
		hx_switch_174 = haxe__rtti__CTypeTools_nameWithParams(name_1, params_1)
	case 3:
		_g_2 := func(hx_value_183 any) *string {
			if hx_value_183 == nil {
				var hx_zero_184 *string
				return hx_zero_184
			}
			return hx_value_183.(*string)
		}(t.params[0])
		_g1_2 := func(hx_value_185 any) *hxrt.Array {
			if hx_value_185 == nil {
				var hx_zero_186 *hxrt.Array
				return hx_zero_186
			}
			return hx_value_185.(*hxrt.Array)
		}(t.params[1])
		name_2 := _g_2
		params_2 := _g1_2
		hx_switch_174 = haxe__rtti__CTypeTools_nameWithParams(name_2, params_2)
	case 4:
		_g_3 := func(hx_value_187 any) *hxrt.Array {
			if hx_value_187 == nil {
				var hx_zero_188 *hxrt.Array
				return hx_zero_188
			}
			return hx_value_187.(*hxrt.Array)
		}(t.params[0])
		_g1_3 := func(hx_value_189 any) *haxe__rtti__CType {
			if hx_value_189 == nil {
				var hx_zero_190 *haxe__rtti__CType
				return hx_zero_190
			}
			return hx_value_189.(*haxe__rtti__CType)
		}(t.params[1])
		args := _g_3
		ret := _g1_3
		var hx_if_191 *string
		if args.Len() == 0 {
			hx_if_191 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Void -> "), haxe__rtti__CTypeTools_toString(ret))
		} else {
			hx_if_191 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(haxe__rtti__CTypeTools_joinFunctionArguments(args), hxrt.StringFromLiteral(" -> ")), haxe__rtti__CTypeTools_toString(ret))
		}
		hx_switch_174 = hx_if_191
	case 5:
		_g_4 := func(hx_value_192 any) *hxrt.Array {
			if hx_value_192 == nil {
				var hx_zero_193 *hxrt.Array
				return hx_zero_193
			}
			return hx_value_192.(*hxrt.Array)
		}(t.params[0])
		fields := _g_4
		hx_switch_174 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{ "), haxe__rtti__CTypeTools_joinClassFields(fields)), hxrt.StringFromLiteral("}"))
	case 6:
		_g_5 := func(hx_value_194 any) *haxe__rtti__CType {
			if hx_value_194 == nil {
				var hx_zero_195 *haxe__rtti__CType
				return hx_zero_195
			}
			return hx_value_194.(*haxe__rtti__CType)
		}(t.params[0])
		d := _g_5
		var hx_if_196 *string
		if d == nil {
			hx_if_196 = hxrt.StringFromLiteral("Dynamic")
		} else {
			hx_if_196 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Dynamic<"), haxe__rtti__CTypeTools_toString(d)), hxrt.StringFromLiteral(">"))
		}
		hx_switch_174 = hx_if_196
	case 7:
		_g_6 := func(hx_value_197 any) *string {
			if hx_value_197 == nil {
				var hx_zero_198 *string
				return hx_zero_198
			}
			return hx_value_197.(*string)
		}(t.params[0])
		_g1_4 := func(hx_value_199 any) *hxrt.Array {
			if hx_value_199 == nil {
				var hx_zero_200 *hxrt.Array
				return hx_zero_200
			}
			return hx_value_199.(*hxrt.Array)
		}(t.params[1])
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

func haxe__rtti__TypeTree_TClassdecl(c hxrt.Anonymous) *haxe__rtti__TypeTree {
	enumValue := &haxe__rtti__TypeTree{tag: 1}
	enumValue.params = []any{c}
	return enumValue
}

func haxe__rtti__TypeTree_TEnumdecl(e hxrt.Anonymous) *haxe__rtti__TypeTree {
	enumValue := &haxe__rtti__TypeTree{tag: 2}
	enumValue.params = []any{e}
	return enumValue
}

func haxe__rtti__TypeTree_TTypedecl(t hxrt.Anonymous) *haxe__rtti__TypeTree {
	enumValue := &haxe__rtti__TypeTree{tag: 3}
	enumValue.params = []any{t}
	return enumValue
}

func haxe__rtti__TypeTree_TAbstractdecl(a hxrt.Anonymous) *haxe__rtti__TypeTree {
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
