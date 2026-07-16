package main

import "snapshot/hxrt"

func haxe__rtti__TypeApi_constructorEq(c1 map[string]any, c2 map[string]any) bool {
	if !hxrt.StringEqualStringPtr(func(hx_obj_532 map[string]any) *string {
		hx_field_533 := hx_obj_532["name"]
		if hx_field_533 == nil {
			var hx_zero_534 *string
			return hx_zero_534
		}
		return hx_field_533.(*string)
	}(c1), func(hx_obj_535 map[string]any) *string {
		hx_field_536 := hx_obj_535["name"]
		if hx_field_536 == nil {
			var hx_zero_537 *string
			return hx_zero_537
		}
		return hx_field_536.(*string)
	}(c2)) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_obj_538 map[string]any) *string {
		hx_field_539 := hx_obj_538["doc"]
		if hx_field_539 == nil {
			var hx_zero_540 *string
			return hx_zero_540
		}
		return hx_field_539.(*string)
	}(c1), func(hx_obj_541 map[string]any) *string {
		hx_field_542 := hx_obj_541["doc"]
		if hx_field_542 == nil {
			var hx_zero_543 *string
			return hx_zero_543
		}
		return hx_field_542.(*string)
	}(c2)) {
		return false
	}
	if (func(hx_obj_544 map[string]any) []map[string]any {
		hx_field_545 := hx_obj_544["args"]
		if hx_field_545 == nil {
			var hx_zero_546 []map[string]any
			return hx_zero_546
		}
		return hx_field_545.([]map[string]any)
	}(c1) == nil) != (func(hx_obj_547 map[string]any) []map[string]any {
		hx_field_548 := hx_obj_547["args"]
		if hx_field_548 == nil {
			var hx_zero_549 []map[string]any
			return hx_zero_549
		}
		return hx_field_548.([]map[string]any)
	}(c2) == nil) {
		return false
	}
	if (func(hx_obj_550 map[string]any) []map[string]any {
		hx_field_551 := hx_obj_550["args"]
		if hx_field_551 == nil {
			var hx_zero_552 []map[string]any
			return hx_zero_552
		}
		return hx_field_551.([]map[string]any)
	}(c1) != nil) && !haxe__rtti__TypeApi_sameConstructorArguments(func(hx_obj_553 map[string]any) []map[string]any {
		hx_field_554 := hx_obj_553["args"]
		if hx_field_554 == nil {
			var hx_zero_555 []map[string]any
			return hx_zero_555
		}
		return hx_field_554.([]map[string]any)
	}(c1), func(hx_obj_556 map[string]any) []map[string]any {
		hx_field_557 := hx_obj_556["args"]
		if hx_field_557 == nil {
			var hx_zero_558 []map[string]any
			return hx_zero_558
		}
		return hx_field_557.([]map[string]any)
	}(c2)) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_fieldEq(f1 map[string]any, f2 map[string]any) bool {
	if !hxrt.StringEqualStringPtr(func(hx_obj_559 map[string]any) *string {
		hx_field_560 := hx_obj_559["name"]
		if hx_field_560 == nil {
			var hx_zero_561 *string
			return hx_zero_561
		}
		return hx_field_560.(*string)
	}(f1), func(hx_obj_562 map[string]any) *string {
		hx_field_563 := hx_obj_562["name"]
		if hx_field_563 == nil {
			var hx_zero_564 *string
			return hx_zero_564
		}
		return hx_field_563.(*string)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_typeEq(func(hx_obj_565 map[string]any) *haxe__rtti__CType {
		hx_field_566 := hx_obj_565["type"]
		if hx_field_566 == nil {
			var hx_zero_567 *haxe__rtti__CType
			return hx_zero_567
		}
		return hx_field_566.(*haxe__rtti__CType)
	}(f1), func(hx_obj_568 map[string]any) *haxe__rtti__CType {
		hx_field_569 := hx_obj_568["type"]
		if hx_field_569 == nil {
			var hx_zero_570 *haxe__rtti__CType
			return hx_zero_570
		}
		return hx_field_569.(*haxe__rtti__CType)
	}(f2)) {
		return false
	}
	if func(hx_obj_571 map[string]any) bool {
		hx_field_572 := hx_obj_571["isPublic"]
		if hx_field_572 == nil {
			var hx_zero_573 bool
			return hx_zero_573
		}
		return hx_field_572.(bool)
	}(f1) != func(hx_obj_574 map[string]any) bool {
		hx_field_575 := hx_obj_574["isPublic"]
		if hx_field_575 == nil {
			var hx_zero_576 bool
			return hx_zero_576
		}
		return hx_field_575.(bool)
	}(f2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_obj_577 map[string]any) *string {
		hx_field_578 := hx_obj_577["doc"]
		if hx_field_578 == nil {
			var hx_zero_579 *string
			return hx_zero_579
		}
		return hx_field_578.(*string)
	}(f1), func(hx_obj_580 map[string]any) *string {
		hx_field_581 := hx_obj_580["doc"]
		if hx_field_581 == nil {
			var hx_zero_582 *string
			return hx_zero_582
		}
		return hx_field_581.(*string)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_obj_583 map[string]any) *haxe__rtti__Rights {
		hx_field_584 := hx_obj_583["get"]
		if hx_field_584 == nil {
			var hx_zero_585 *haxe__rtti__Rights
			return hx_zero_585
		}
		return hx_field_584.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_586 map[string]any) *haxe__rtti__Rights {
		hx_field_587 := hx_obj_586["get"]
		if hx_field_587 == nil {
			var hx_zero_588 *haxe__rtti__Rights
			return hx_zero_588
		}
		return hx_field_587.(*haxe__rtti__Rights)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_obj_589 map[string]any) *haxe__rtti__Rights {
		hx_field_590 := hx_obj_589["set"]
		if hx_field_590 == nil {
			var hx_zero_591 *haxe__rtti__Rights
			return hx_zero_591
		}
		return hx_field_590.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_592 map[string]any) *haxe__rtti__Rights {
		hx_field_593 := hx_obj_592["set"]
		if hx_field_593 == nil {
			var hx_zero_594 *haxe__rtti__Rights
			return hx_zero_594
		}
		return hx_field_593.(*haxe__rtti__Rights)
	}(f2)) {
		return false
	}
	if (func(hx_obj_595 map[string]any) []*string {
		hx_field_596 := hx_obj_595["params"]
		if hx_field_596 == nil {
			var hx_zero_597 []*string
			return hx_zero_597
		}
		return hx_field_596.([]*string)
	}(f1) == nil) != (func(hx_obj_598 map[string]any) []*string {
		hx_field_599 := hx_obj_598["params"]
		if hx_field_599 == nil {
			var hx_zero_600 []*string
			return hx_zero_600
		}
		return hx_field_599.([]*string)
	}(f2) == nil) {
		return false
	}
	if (func(hx_obj_601 map[string]any) []*string {
		hx_field_602 := hx_obj_601["params"]
		if hx_field_602 == nil {
			var hx_zero_603 []*string
			return hx_zero_603
		}
		return hx_field_602.([]*string)
	}(f1) != nil) && !haxe__rtti__TypeApi_sameTypeParamNames(func(hx_obj_604 map[string]any) []*string {
		hx_field_605 := hx_obj_604["params"]
		if hx_field_605 == nil {
			var hx_zero_606 []*string
			return hx_zero_606
		}
		return hx_field_605.([]*string)
	}(f1), func(hx_obj_607 map[string]any) []*string {
		hx_field_608 := hx_obj_607["params"]
		if hx_field_608 == nil {
			var hx_zero_609 []*string
			return hx_zero_609
		}
		return hx_field_608.([]*string)
	}(f2)) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_isVar(t *haxe__rtti__CType) bool {
	var hx_if_610 bool
	if t.tag == 4 {
		_g := t.params[0].([]map[string]any)
		_ = _g
		_g_1 := t.params[1].(*haxe__rtti__CType)
		_ = _g_1
		hx_if_610 = false
	} else {
		hx_if_610 = true
	}
	return hx_if_610
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

func haxe__rtti__TypeApi_sameClassFields(l1 []map[string]any, l2 []map[string]any) bool {
	if len(l1) != len(l2) {
		return false
	}
	_g := 0
	_g1 := len(l1)
	for _g < _g1 {
		hx_post_611 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_611
		if !haxe__rtti__TypeApi_fieldEq(l1[i], l2[i]) {
			return false
		}
	}
	return true
}

func haxe__rtti__TypeApi_sameConstructorArguments(l1 []map[string]any, l2 []map[string]any) bool {
	if len(l1) != len(l2) {
		return false
	}
	_g := 0
	_g1 := len(l1)
	for _g < _g1 {
		hx_post_612 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_612
		a := l1[i]
		b := l2[i]
		if (!hxrt.StringEqualStringPtr(func(hx_obj_613 map[string]any) *string {
			hx_field_614 := hx_obj_613["name"]
			if hx_field_614 == nil {
				var hx_zero_615 *string
				return hx_zero_615
			}
			return hx_field_614.(*string)
		}(a), func(hx_obj_616 map[string]any) *string {
			hx_field_617 := hx_obj_616["name"]
			if hx_field_617 == nil {
				var hx_zero_618 *string
				return hx_zero_618
			}
			return hx_field_617.(*string)
		}(b)) || (func(hx_obj_619 map[string]any) bool {
			hx_field_620 := hx_obj_619["opt"]
			if hx_field_620 == nil {
				var hx_zero_621 bool
				return hx_zero_621
			}
			return hx_field_620.(bool)
		}(a) != func(hx_obj_622 map[string]any) bool {
			hx_field_623 := hx_obj_622["opt"]
			if hx_field_623 == nil {
				var hx_zero_624 bool
				return hx_zero_624
			}
			return hx_field_623.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_625 map[string]any) *haxe__rtti__CType {
			hx_field_626 := hx_obj_625["t"]
			if hx_field_626 == nil {
				var hx_zero_627 *haxe__rtti__CType
				return hx_zero_627
			}
			return hx_field_626.(*haxe__rtti__CType)
		}(a), func(hx_obj_628 map[string]any) *haxe__rtti__CType {
			hx_field_629 := hx_obj_628["t"]
			if hx_field_629 == nil {
				var hx_zero_630 *haxe__rtti__CType
				return hx_zero_630
			}
			return hx_field_629.(*haxe__rtti__CType)
		}(b)) {
			return false
		}
	}
	return true
}

func haxe__rtti__TypeApi_sameFunctionArguments(l1 []map[string]any, l2 []map[string]any) bool {
	if len(l1) != len(l2) {
		return false
	}
	_g := 0
	_g1 := len(l1)
	for _g < _g1 {
		hx_post_631 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_631
		a := l1[i]
		b := l2[i]
		if (!hxrt.StringEqualStringPtr(func(hx_obj_632 map[string]any) *string {
			hx_field_633 := hx_obj_632["name"]
			if hx_field_633 == nil {
				var hx_zero_634 *string
				return hx_zero_634
			}
			return hx_field_633.(*string)
		}(a), func(hx_obj_635 map[string]any) *string {
			hx_field_636 := hx_obj_635["name"]
			if hx_field_636 == nil {
				var hx_zero_637 *string
				return hx_zero_637
			}
			return hx_field_636.(*string)
		}(b)) || (func(hx_obj_638 map[string]any) bool {
			hx_field_639 := hx_obj_638["opt"]
			if hx_field_639 == nil {
				var hx_zero_640 bool
				return hx_zero_640
			}
			return hx_field_639.(bool)
		}(a) != func(hx_obj_641 map[string]any) bool {
			hx_field_642 := hx_obj_641["opt"]
			if hx_field_642 == nil {
				var hx_zero_643 bool
				return hx_zero_643
			}
			return hx_field_642.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_644 map[string]any) *haxe__rtti__CType {
			hx_field_645 := hx_obj_644["t"]
			if hx_field_645 == nil {
				var hx_zero_646 *haxe__rtti__CType
				return hx_zero_646
			}
			return hx_field_645.(*haxe__rtti__CType)
		}(a), func(hx_obj_647 map[string]any) *haxe__rtti__CType {
			hx_field_648 := hx_obj_647["t"]
			if hx_field_648 == nil {
				var hx_zero_649 *haxe__rtti__CType
				return hx_zero_649
			}
			return hx_field_648.(*haxe__rtti__CType)
		}(b)) {
			return false
		}
	}
	return true
}

func haxe__rtti__TypeApi_sameTypeParamNames(p1 []*string, p2 []*string) bool {
	if len(p1) != len(p2) {
		return false
	}
	_g := 0
	_g1 := len(p1)
	for _g < _g1 {
		hx_post_650 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_650
		if !hxrt.StringEqualStringPtr(p1[i], p2[i]) {
			return false
		}
	}
	return true
}

func haxe__rtti__TypeApi_sameTypes(l1 []*haxe__rtti__CType, l2 []*haxe__rtti__CType) bool {
	if len(l1) != len(l2) {
		return false
	}
	_g := 0
	_g1 := len(l1)
	for _g < _g1 {
		hx_post_651 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_651
		if !haxe__rtti__TypeApi_typeEq(l1[i], l2[i]) {
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
		_g1 := t1.params[1].([]*haxe__rtti__CType)
		name := _g
		params := _g1
		if t2.tag == 1 {
			_g_1 := t2.params[0].(*string)
			_g1_1 := t2.params[1].([]*haxe__rtti__CType)
			name2 := _g_1
			params2 := _g1_1
			return (hxrt.StringEqualStringPtr(name, name2) && haxe__rtti__TypeApi_sameTypes(params, params2))
		} else {
		}
	case 2:
		_g_2 := t1.params[0].(*string)
		_g1_2 := t1.params[1].([]*haxe__rtti__CType)
		name_1 := _g_2
		params_1 := _g1_2
		if t2.tag == 2 {
			_g_3 := t2.params[0].(*string)
			_g1_3 := t2.params[1].([]*haxe__rtti__CType)
			name2_1 := _g_3
			params2_1 := _g1_3
			return (hxrt.StringEqualStringPtr(name_1, name2_1) && haxe__rtti__TypeApi_sameTypes(params_1, params2_1))
		} else {
		}
	case 3:
		_g_4 := t1.params[0].(*string)
		_g1_4 := t1.params[1].([]*haxe__rtti__CType)
		name_2 := _g_4
		params_2 := _g1_4
		if t2.tag == 3 {
			_g_5 := t2.params[0].(*string)
			_g1_5 := t2.params[1].([]*haxe__rtti__CType)
			name2_2 := _g_5
			params2_2 := _g1_5
			return (hxrt.StringEqualStringPtr(name_2, name2_2) && haxe__rtti__TypeApi_sameTypes(params_2, params2_2))
		} else {
		}
	case 4:
		_g_6 := t1.params[0].([]map[string]any)
		_g1_6 := t1.params[1].(*haxe__rtti__CType)
		args := _g_6
		ret := _g1_6
		if t2.tag == 4 {
			_g_7 := t2.params[0].([]map[string]any)
			_g1_7 := t2.params[1].(*haxe__rtti__CType)
			args2 := _g_7
			ret2 := _g1_7
			return (haxe__rtti__TypeApi_sameFunctionArguments(args, args2) && haxe__rtti__TypeApi_typeEq(ret, ret2))
		} else {
		}
	case 5:
		_g_8 := t1.params[0].([]map[string]any)
		fields := _g_8
		if t2.tag == 5 {
			_g_9 := t2.params[0].([]map[string]any)
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
		_g1_8 := t1.params[1].([]*haxe__rtti__CType)
		name_3 := _g_12
		params_3 := _g1_8
		if t2.tag == 7 {
			_g_13 := t2.params[0].(*string)
			_g1_9 := t2.params[1].([]*haxe__rtti__CType)
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
		_g_2 := t.params[2].([]*haxe__rtti__TypeTree)
		_ = _g_2
		hxrt.Throw(hxrt.StringFromLiteral("Unexpected Package"))
		var hx_throw_zero_652 map[string]any
		return hx_throw_zero_652
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
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func(hx_obj_655 map[string]any) *string {
		hx_field_656 := hx_obj_655["name"]
		if hx_field_656 == nil {
			var hx_zero_657 *string
			return hx_zero_657
		}
		return hx_field_656.(*string)
	}(cf), hxrt.StringFromLiteral(":")), haxe__rtti__CTypeTools_toString(func(hx_obj_658 map[string]any) *haxe__rtti__CType {
		hx_field_659 := hx_obj_658["type"]
		if hx_field_659 == nil {
			var hx_zero_660 *haxe__rtti__CType
			return hx_zero_660
		}
		return hx_field_659.(*haxe__rtti__CType)
	}(cf)))
}

func haxe__rtti__CTypeTools_functionArgumentName(arg map[string]any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_if_664 *string
		if func(hx_obj_661 map[string]any) bool {
			hx_field_662 := hx_obj_661["opt"]
			if hx_field_662 == nil {
				var hx_zero_663 bool
				return hx_zero_663
			}
			return hx_field_662.(bool)
		}(arg) {
			hx_if_664 = hxrt.StringFromLiteral("?")
		} else {
			hx_if_664 = hxrt.StringFromLiteral("")
		}
		return hx_if_664
	}(), func() *string {
		var hx_if_671 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_665 map[string]any) *string {
			hx_field_666 := hx_obj_665["name"]
			if hx_field_666 == nil {
				var hx_zero_667 *string
				return hx_zero_667
			}
			return hx_field_666.(*string)
		}(arg), hxrt.StringFromLiteral("")) {
			hx_if_671 = hxrt.StringFromLiteral("")
		} else {
			hx_if_671 = hxrt.StringConcatStringPtr(func(hx_obj_668 map[string]any) *string {
				hx_field_669 := hx_obj_668["name"]
				if hx_field_669 == nil {
					var hx_zero_670 *string
					return hx_zero_670
				}
				return hx_field_669.(*string)
			}(arg), hxrt.StringFromLiteral(":"))
		}
		return hx_if_671
	}()), haxe__rtti__CTypeTools_toString(func(hx_obj_672 map[string]any) *haxe__rtti__CType {
		hx_field_673 := hx_obj_672["t"]
		if hx_field_673 == nil {
			var hx_zero_674 *haxe__rtti__CType
			return hx_zero_674
		}
		return hx_field_673.(*haxe__rtti__CType)
	}(arg))), func() *string {
		var hx_if_681 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_675 map[string]any) *string {
			hx_field_676 := hx_obj_675["value"]
			if hx_field_676 == nil {
				var hx_zero_677 *string
				return hx_zero_677
			}
			return hx_field_676.(*string)
		}(arg), nil) {
			hx_if_681 = hxrt.StringFromLiteral("")
		} else {
			hx_if_681 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" = "), func(hx_obj_678 map[string]any) *string {
				hx_field_679 := hx_obj_678["value"]
				if hx_field_679 == nil {
					var hx_zero_680 *string
					return hx_zero_680
				}
				return hx_field_679.(*string)
			}(arg))
		}
		return hx_if_681
	}())
}

func haxe__rtti__CTypeTools_joinClassFields(fields []map[string]any) *string {
	parts := []*string{}
	_g := 0
	for _g < len(fields) {
		field := fields[_g]
		_g = int(int32((_g + 1)))
		parts = append(parts, haxe__rtti__CTypeTools_classField(field))
	}
	return haxe__rtti__CTypeTools_joinStringArray(parts, hxrt.StringFromLiteral(", "))
}

func haxe__rtti__CTypeTools_joinFunctionArguments(args []map[string]any) *string {
	parts := []*string{}
	_g := 0
	for _g < len(args) {
		arg := args[_g]
		_g = int(int32((_g + 1)))
		parts = append(parts, haxe__rtti__CTypeTools_functionArgumentName(arg))
	}
	return haxe__rtti__CTypeTools_joinStringArray(parts, hxrt.StringFromLiteral(" -> "))
}

func haxe__rtti__CTypeTools_joinStringArray(parts []*string, separator *string) *string {
	var buf_b *string
	buf_b = hxrt.StringFromLiteral("")
	_g := 0
	_g1 := len(parts)
	for _g < _g1 {
		hx_post_684 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_684
		if i > 0 {
			buf_b = hxrt.StringConcatStringPtr(buf_b, hxrt.StdString(separator))
		}
		buf_b = hxrt.StringConcatStringPtr(buf_b, hxrt.StdString(parts[i]))
	}
	return buf_b
}

func haxe__rtti__CTypeTools_nameWithParams(name *string, params []*haxe__rtti__CType) *string {
	if len(params) == 0 {
		return name
	}
	parts := []*string{}
	_g := 0
	for _g < len(params) {
		param := params[_g]
		_g = int(int32((_g + 1)))
		parts = append(parts, haxe__rtti__CTypeTools_toString(param))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(name, hxrt.StringFromLiteral("<")), haxe__rtti__CTypeTools_joinStringArray(parts, hxrt.StringFromLiteral(", "))), hxrt.StringFromLiteral(">"))
}

func haxe__rtti__CTypeTools_toString(t *haxe__rtti__CType) *string {
	var hx_switch_686 *string
	switch t.tag {
	case 0:
		hx_switch_686 = hxrt.StringFromLiteral("unknown")
	case 1:
		_g := t.params[0].(*string)
		_g1 := t.params[1].([]*haxe__rtti__CType)
		name := _g
		params := _g1
		hx_switch_686 = haxe__rtti__CTypeTools_nameWithParams(name, params)
	case 2:
		_g_1 := t.params[0].(*string)
		_g1_1 := t.params[1].([]*haxe__rtti__CType)
		name_1 := _g_1
		params_1 := _g1_1
		hx_switch_686 = haxe__rtti__CTypeTools_nameWithParams(name_1, params_1)
	case 3:
		_g_2 := t.params[0].(*string)
		_g1_2 := t.params[1].([]*haxe__rtti__CType)
		name_2 := _g_2
		params_2 := _g1_2
		hx_switch_686 = haxe__rtti__CTypeTools_nameWithParams(name_2, params_2)
	case 4:
		_g_3 := t.params[0].([]map[string]any)
		_g1_3 := t.params[1].(*haxe__rtti__CType)
		args := _g_3
		ret := _g1_3
		var hx_if_687 *string
		if len(args) == 0 {
			hx_if_687 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Void -> "), haxe__rtti__CTypeTools_toString(ret))
		} else {
			hx_if_687 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(haxe__rtti__CTypeTools_joinFunctionArguments(args), hxrt.StringFromLiteral(" -> ")), haxe__rtti__CTypeTools_toString(ret))
		}
		hx_switch_686 = hx_if_687
	case 5:
		_g_4 := t.params[0].([]map[string]any)
		fields := _g_4
		hx_switch_686 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{ "), haxe__rtti__CTypeTools_joinClassFields(fields)), hxrt.StringFromLiteral("}"))
	case 6:
		_g_5 := t.params[0].(*haxe__rtti__CType)
		d := _g_5
		var hx_if_688 *string
		if d == nil {
			hx_if_688 = hxrt.StringFromLiteral("Dynamic")
		} else {
			hx_if_688 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Dynamic<"), haxe__rtti__CTypeTools_toString(d)), hxrt.StringFromLiteral(">"))
		}
		hx_switch_686 = hx_if_688
	case 7:
		_g_6 := t.params[0].(*string)
		_g1_4 := t.params[1].([]*haxe__rtti__CType)
		name_3 := _g_6
		params_3 := _g1_4
		hx_switch_686 = haxe__rtti__CTypeTools_nameWithParams(name_3, params_3)
	}
	return hx_switch_686
}

type haxe__rtti__TypeTree struct {
	tag    int
	params []any
}

func haxe__rtti__TypeTree_TPackage(name *string, full *string, subs []*haxe__rtti__TypeTree) *haxe__rtti__TypeTree {
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

func haxe__rtti__CType_CEnum(name *string, params []*haxe__rtti__CType) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 1}
	enumValue.params = []any{name, params}
	return enumValue
}

func haxe__rtti__CType_CClass(name *string, params []*haxe__rtti__CType) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 2}
	enumValue.params = []any{name, params}
	return enumValue
}

func haxe__rtti__CType_CTypedef(name *string, params []*haxe__rtti__CType) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 3}
	enumValue.params = []any{name, params}
	return enumValue
}

func haxe__rtti__CType_CFunction(args []map[string]any, ret *haxe__rtti__CType) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 4}
	enumValue.params = []any{args, ret}
	return enumValue
}

func haxe__rtti__CType_CAnonymous(fields []map[string]any) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 5}
	enumValue.params = []any{fields}
	return enumValue
}

func haxe__rtti__CType_CDynamic(t *haxe__rtti__CType) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 6}
	enumValue.params = []any{t}
	return enumValue
}

func haxe__rtti__CType_CAbstract(name *string, params []*haxe__rtti__CType) *haxe__rtti__CType {
	enumValue := &haxe__rtti__CType{tag: 7}
	enumValue.params = []any{name, params}
	return enumValue
}
