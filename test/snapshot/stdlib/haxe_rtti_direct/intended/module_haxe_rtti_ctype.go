package main

import "snapshot/hxrt"

func haxe__rtti__TypeApi_constructorEq(c1 map[string]any, c2 map[string]any) bool {
	if !hxrt.StringEqualStringPtr(func(hx_obj_545 map[string]any) *string {
		hx_field_546 := hx_obj_545["name"]
		if hx_field_546 == nil {
			var hx_zero_547 *string
			return hx_zero_547
		}
		return hx_field_546.(*string)
	}(c1), func(hx_obj_548 map[string]any) *string {
		hx_field_549 := hx_obj_548["name"]
		if hx_field_549 == nil {
			var hx_zero_550 *string
			return hx_zero_550
		}
		return hx_field_549.(*string)
	}(c2)) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_obj_551 map[string]any) *string {
		hx_field_552 := hx_obj_551["doc"]
		if hx_field_552 == nil {
			var hx_zero_553 *string
			return hx_zero_553
		}
		return hx_field_552.(*string)
	}(c1), func(hx_obj_554 map[string]any) *string {
		hx_field_555 := hx_obj_554["doc"]
		if hx_field_555 == nil {
			var hx_zero_556 *string
			return hx_zero_556
		}
		return hx_field_555.(*string)
	}(c2)) {
		return false
	}
	if (func(hx_obj_557 map[string]any) []map[string]any {
		hx_field_558 := hx_obj_557["args"]
		if hx_field_558 == nil {
			var hx_zero_559 []map[string]any
			return hx_zero_559
		}
		return hx_field_558.([]map[string]any)
	}(c1) == nil) != (func(hx_obj_560 map[string]any) []map[string]any {
		hx_field_561 := hx_obj_560["args"]
		if hx_field_561 == nil {
			var hx_zero_562 []map[string]any
			return hx_zero_562
		}
		return hx_field_561.([]map[string]any)
	}(c2) == nil) {
		return false
	}
	if (func(hx_obj_563 map[string]any) []map[string]any {
		hx_field_564 := hx_obj_563["args"]
		if hx_field_564 == nil {
			var hx_zero_565 []map[string]any
			return hx_zero_565
		}
		return hx_field_564.([]map[string]any)
	}(c1) != nil) && !haxe__rtti__TypeApi_sameConstructorArguments(func(hx_obj_566 map[string]any) []map[string]any {
		hx_field_567 := hx_obj_566["args"]
		if hx_field_567 == nil {
			var hx_zero_568 []map[string]any
			return hx_zero_568
		}
		return hx_field_567.([]map[string]any)
	}(c1), func(hx_obj_569 map[string]any) []map[string]any {
		hx_field_570 := hx_obj_569["args"]
		if hx_field_570 == nil {
			var hx_zero_571 []map[string]any
			return hx_zero_571
		}
		return hx_field_570.([]map[string]any)
	}(c2)) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_fieldEq(f1 map[string]any, f2 map[string]any) bool {
	if !hxrt.StringEqualStringPtr(func(hx_obj_572 map[string]any) *string {
		hx_field_573 := hx_obj_572["name"]
		if hx_field_573 == nil {
			var hx_zero_574 *string
			return hx_zero_574
		}
		return hx_field_573.(*string)
	}(f1), func(hx_obj_575 map[string]any) *string {
		hx_field_576 := hx_obj_575["name"]
		if hx_field_576 == nil {
			var hx_zero_577 *string
			return hx_zero_577
		}
		return hx_field_576.(*string)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_typeEq(func(hx_obj_578 map[string]any) *haxe__rtti__CType {
		hx_field_579 := hx_obj_578["type"]
		if hx_field_579 == nil {
			var hx_zero_580 *haxe__rtti__CType
			return hx_zero_580
		}
		return hx_field_579.(*haxe__rtti__CType)
	}(f1), func(hx_obj_581 map[string]any) *haxe__rtti__CType {
		hx_field_582 := hx_obj_581["type"]
		if hx_field_582 == nil {
			var hx_zero_583 *haxe__rtti__CType
			return hx_zero_583
		}
		return hx_field_582.(*haxe__rtti__CType)
	}(f2)) {
		return false
	}
	if func(hx_obj_584 map[string]any) bool {
		hx_field_585 := hx_obj_584["isPublic"]
		if hx_field_585 == nil {
			var hx_zero_586 bool
			return hx_zero_586
		}
		return hx_field_585.(bool)
	}(f1) != func(hx_obj_587 map[string]any) bool {
		hx_field_588 := hx_obj_587["isPublic"]
		if hx_field_588 == nil {
			var hx_zero_589 bool
			return hx_zero_589
		}
		return hx_field_588.(bool)
	}(f2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_obj_590 map[string]any) *string {
		hx_field_591 := hx_obj_590["doc"]
		if hx_field_591 == nil {
			var hx_zero_592 *string
			return hx_zero_592
		}
		return hx_field_591.(*string)
	}(f1), func(hx_obj_593 map[string]any) *string {
		hx_field_594 := hx_obj_593["doc"]
		if hx_field_594 == nil {
			var hx_zero_595 *string
			return hx_zero_595
		}
		return hx_field_594.(*string)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_obj_596 map[string]any) *haxe__rtti__Rights {
		hx_field_597 := hx_obj_596["get"]
		if hx_field_597 == nil {
			var hx_zero_598 *haxe__rtti__Rights
			return hx_zero_598
		}
		return hx_field_597.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_599 map[string]any) *haxe__rtti__Rights {
		hx_field_600 := hx_obj_599["get"]
		if hx_field_600 == nil {
			var hx_zero_601 *haxe__rtti__Rights
			return hx_zero_601
		}
		return hx_field_600.(*haxe__rtti__Rights)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_obj_602 map[string]any) *haxe__rtti__Rights {
		hx_field_603 := hx_obj_602["set"]
		if hx_field_603 == nil {
			var hx_zero_604 *haxe__rtti__Rights
			return hx_zero_604
		}
		return hx_field_603.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_605 map[string]any) *haxe__rtti__Rights {
		hx_field_606 := hx_obj_605["set"]
		if hx_field_606 == nil {
			var hx_zero_607 *haxe__rtti__Rights
			return hx_zero_607
		}
		return hx_field_606.(*haxe__rtti__Rights)
	}(f2)) {
		return false
	}
	if (func(hx_obj_608 map[string]any) []*string {
		hx_field_609 := hx_obj_608["params"]
		if hx_field_609 == nil {
			var hx_zero_610 []*string
			return hx_zero_610
		}
		return hx_field_609.([]*string)
	}(f1) == nil) != (func(hx_obj_611 map[string]any) []*string {
		hx_field_612 := hx_obj_611["params"]
		if hx_field_612 == nil {
			var hx_zero_613 []*string
			return hx_zero_613
		}
		return hx_field_612.([]*string)
	}(f2) == nil) {
		return false
	}
	if (func(hx_obj_614 map[string]any) []*string {
		hx_field_615 := hx_obj_614["params"]
		if hx_field_615 == nil {
			var hx_zero_616 []*string
			return hx_zero_616
		}
		return hx_field_615.([]*string)
	}(f1) != nil) && !haxe__rtti__TypeApi_sameTypeParamNames(func(hx_obj_617 map[string]any) []*string {
		hx_field_618 := hx_obj_617["params"]
		if hx_field_618 == nil {
			var hx_zero_619 []*string
			return hx_zero_619
		}
		return hx_field_618.([]*string)
	}(f1), func(hx_obj_620 map[string]any) []*string {
		hx_field_621 := hx_obj_620["params"]
		if hx_field_621 == nil {
			var hx_zero_622 []*string
			return hx_zero_622
		}
		return hx_field_621.([]*string)
	}(f2)) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_isVar(t *haxe__rtti__CType) bool {
	var hx_if_623 bool
	if t.tag == 4 {
		_g := t.params[0].([]map[string]any)
		_ = _g
		_g_1 := t.params[1].(*haxe__rtti__CType)
		_ = _g_1
		hx_if_623 = false
	} else {
		hx_if_623 = true
	}
	return hx_if_623
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
		hx_post_624 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_624
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
		hx_post_625 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_625
		a := l1[i]
		b := l2[i]
		if (!hxrt.StringEqualStringPtr(func(hx_obj_626 map[string]any) *string {
			hx_field_627 := hx_obj_626["name"]
			if hx_field_627 == nil {
				var hx_zero_628 *string
				return hx_zero_628
			}
			return hx_field_627.(*string)
		}(a), func(hx_obj_629 map[string]any) *string {
			hx_field_630 := hx_obj_629["name"]
			if hx_field_630 == nil {
				var hx_zero_631 *string
				return hx_zero_631
			}
			return hx_field_630.(*string)
		}(b)) || (func(hx_obj_632 map[string]any) bool {
			hx_field_633 := hx_obj_632["opt"]
			if hx_field_633 == nil {
				var hx_zero_634 bool
				return hx_zero_634
			}
			return hx_field_633.(bool)
		}(a) != func(hx_obj_635 map[string]any) bool {
			hx_field_636 := hx_obj_635["opt"]
			if hx_field_636 == nil {
				var hx_zero_637 bool
				return hx_zero_637
			}
			return hx_field_636.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_638 map[string]any) *haxe__rtti__CType {
			hx_field_639 := hx_obj_638["t"]
			if hx_field_639 == nil {
				var hx_zero_640 *haxe__rtti__CType
				return hx_zero_640
			}
			return hx_field_639.(*haxe__rtti__CType)
		}(a), func(hx_obj_641 map[string]any) *haxe__rtti__CType {
			hx_field_642 := hx_obj_641["t"]
			if hx_field_642 == nil {
				var hx_zero_643 *haxe__rtti__CType
				return hx_zero_643
			}
			return hx_field_642.(*haxe__rtti__CType)
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
		hx_post_644 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_644
		a := l1[i]
		b := l2[i]
		if (!hxrt.StringEqualStringPtr(func(hx_obj_645 map[string]any) *string {
			hx_field_646 := hx_obj_645["name"]
			if hx_field_646 == nil {
				var hx_zero_647 *string
				return hx_zero_647
			}
			return hx_field_646.(*string)
		}(a), func(hx_obj_648 map[string]any) *string {
			hx_field_649 := hx_obj_648["name"]
			if hx_field_649 == nil {
				var hx_zero_650 *string
				return hx_zero_650
			}
			return hx_field_649.(*string)
		}(b)) || (func(hx_obj_651 map[string]any) bool {
			hx_field_652 := hx_obj_651["opt"]
			if hx_field_652 == nil {
				var hx_zero_653 bool
				return hx_zero_653
			}
			return hx_field_652.(bool)
		}(a) != func(hx_obj_654 map[string]any) bool {
			hx_field_655 := hx_obj_654["opt"]
			if hx_field_655 == nil {
				var hx_zero_656 bool
				return hx_zero_656
			}
			return hx_field_655.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_657 map[string]any) *haxe__rtti__CType {
			hx_field_658 := hx_obj_657["t"]
			if hx_field_658 == nil {
				var hx_zero_659 *haxe__rtti__CType
				return hx_zero_659
			}
			return hx_field_658.(*haxe__rtti__CType)
		}(a), func(hx_obj_660 map[string]any) *haxe__rtti__CType {
			hx_field_661 := hx_obj_660["t"]
			if hx_field_661 == nil {
				var hx_zero_662 *haxe__rtti__CType
				return hx_zero_662
			}
			return hx_field_661.(*haxe__rtti__CType)
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
		hx_post_663 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_663
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
		hx_post_664 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_664
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
		var hx_throw_zero_665 map[string]any
		return hx_throw_zero_665
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
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func(hx_obj_668 map[string]any) *string {
		hx_field_669 := hx_obj_668["name"]
		if hx_field_669 == nil {
			var hx_zero_670 *string
			return hx_zero_670
		}
		return hx_field_669.(*string)
	}(cf), hxrt.StringFromLiteral(":")), haxe__rtti__CTypeTools_toString(func(hx_obj_671 map[string]any) *haxe__rtti__CType {
		hx_field_672 := hx_obj_671["type"]
		if hx_field_672 == nil {
			var hx_zero_673 *haxe__rtti__CType
			return hx_zero_673
		}
		return hx_field_672.(*haxe__rtti__CType)
	}(cf)))
}

func haxe__rtti__CTypeTools_functionArgumentName(arg map[string]any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_if_677 *string
		if func(hx_obj_674 map[string]any) bool {
			hx_field_675 := hx_obj_674["opt"]
			if hx_field_675 == nil {
				var hx_zero_676 bool
				return hx_zero_676
			}
			return hx_field_675.(bool)
		}(arg) {
			hx_if_677 = hxrt.StringFromLiteral("?")
		} else {
			hx_if_677 = hxrt.StringFromLiteral("")
		}
		return hx_if_677
	}(), func() *string {
		var hx_if_684 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_678 map[string]any) *string {
			hx_field_679 := hx_obj_678["name"]
			if hx_field_679 == nil {
				var hx_zero_680 *string
				return hx_zero_680
			}
			return hx_field_679.(*string)
		}(arg), hxrt.StringFromLiteral("")) {
			hx_if_684 = hxrt.StringFromLiteral("")
		} else {
			hx_if_684 = hxrt.StringConcatStringPtr(func(hx_obj_681 map[string]any) *string {
				hx_field_682 := hx_obj_681["name"]
				if hx_field_682 == nil {
					var hx_zero_683 *string
					return hx_zero_683
				}
				return hx_field_682.(*string)
			}(arg), hxrt.StringFromLiteral(":"))
		}
		return hx_if_684
	}()), haxe__rtti__CTypeTools_toString(func(hx_obj_685 map[string]any) *haxe__rtti__CType {
		hx_field_686 := hx_obj_685["t"]
		if hx_field_686 == nil {
			var hx_zero_687 *haxe__rtti__CType
			return hx_zero_687
		}
		return hx_field_686.(*haxe__rtti__CType)
	}(arg))), func() *string {
		var hx_if_694 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_688 map[string]any) *string {
			hx_field_689 := hx_obj_688["value"]
			if hx_field_689 == nil {
				var hx_zero_690 *string
				return hx_zero_690
			}
			return hx_field_689.(*string)
		}(arg), nil) {
			hx_if_694 = hxrt.StringFromLiteral("")
		} else {
			hx_if_694 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" = "), func(hx_obj_691 map[string]any) *string {
				hx_field_692 := hx_obj_691["value"]
				if hx_field_692 == nil {
					var hx_zero_693 *string
					return hx_zero_693
				}
				return hx_field_692.(*string)
			}(arg))
		}
		return hx_if_694
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
		hx_post_697 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_697
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
	var hx_switch_699 *string
	switch t.tag {
	case 0:
		hx_switch_699 = hxrt.StringFromLiteral("unknown")
	case 1:
		_g := t.params[0].(*string)
		_g1 := t.params[1].([]*haxe__rtti__CType)
		name := _g
		params := _g1
		hx_switch_699 = haxe__rtti__CTypeTools_nameWithParams(name, params)
	case 2:
		_g_1 := t.params[0].(*string)
		_g1_1 := t.params[1].([]*haxe__rtti__CType)
		name_1 := _g_1
		params_1 := _g1_1
		hx_switch_699 = haxe__rtti__CTypeTools_nameWithParams(name_1, params_1)
	case 3:
		_g_2 := t.params[0].(*string)
		_g1_2 := t.params[1].([]*haxe__rtti__CType)
		name_2 := _g_2
		params_2 := _g1_2
		hx_switch_699 = haxe__rtti__CTypeTools_nameWithParams(name_2, params_2)
	case 4:
		_g_3 := t.params[0].([]map[string]any)
		_g1_3 := t.params[1].(*haxe__rtti__CType)
		args := _g_3
		ret := _g1_3
		var hx_if_700 *string
		if len(args) == 0 {
			hx_if_700 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Void -> "), haxe__rtti__CTypeTools_toString(ret))
		} else {
			hx_if_700 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(haxe__rtti__CTypeTools_joinFunctionArguments(args), hxrt.StringFromLiteral(" -> ")), haxe__rtti__CTypeTools_toString(ret))
		}
		hx_switch_699 = hx_if_700
	case 5:
		_g_4 := t.params[0].([]map[string]any)
		fields := _g_4
		hx_switch_699 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{ "), haxe__rtti__CTypeTools_joinClassFields(fields)), hxrt.StringFromLiteral("}"))
	case 6:
		_g_5 := t.params[0].(*haxe__rtti__CType)
		d := _g_5
		var hx_if_701 *string
		if d == nil {
			hx_if_701 = hxrt.StringFromLiteral("Dynamic")
		} else {
			hx_if_701 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Dynamic<"), haxe__rtti__CTypeTools_toString(d)), hxrt.StringFromLiteral(">"))
		}
		hx_switch_699 = hx_if_701
	case 7:
		_g_6 := t.params[0].(*string)
		_g1_4 := t.params[1].([]*haxe__rtti__CType)
		name_3 := _g_6
		params_3 := _g1_4
		hx_switch_699 = haxe__rtti__CTypeTools_nameWithParams(name_3, params_3)
	}
	return hx_switch_699
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
