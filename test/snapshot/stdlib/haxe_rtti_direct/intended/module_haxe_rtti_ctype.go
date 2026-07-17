package main

import "snapshot/hxrt"

func haxe__rtti__TypeApi_constructorEq(c1 map[string]any, c2 map[string]any) bool {
	if !hxrt.StringEqualStringPtr(func(hx_obj_525 map[string]any) *string {
		hx_field_526 := hx_obj_525["name"]
		if hx_field_526 == nil {
			var hx_zero_527 *string
			return hx_zero_527
		}
		return hx_field_526.(*string)
	}(c1), func(hx_obj_528 map[string]any) *string {
		hx_field_529 := hx_obj_528["name"]
		if hx_field_529 == nil {
			var hx_zero_530 *string
			return hx_zero_530
		}
		return hx_field_529.(*string)
	}(c2)) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_obj_531 map[string]any) *string {
		hx_field_532 := hx_obj_531["doc"]
		if hx_field_532 == nil {
			var hx_zero_533 *string
			return hx_zero_533
		}
		return hx_field_532.(*string)
	}(c1), func(hx_obj_534 map[string]any) *string {
		hx_field_535 := hx_obj_534["doc"]
		if hx_field_535 == nil {
			var hx_zero_536 *string
			return hx_zero_536
		}
		return hx_field_535.(*string)
	}(c2)) {
		return false
	}
	if (func(hx_obj_537 map[string]any) []map[string]any {
		hx_field_538 := hx_obj_537["args"]
		if hx_field_538 == nil {
			var hx_zero_539 []map[string]any
			return hx_zero_539
		}
		return hx_field_538.([]map[string]any)
	}(c1) == nil) != (func(hx_obj_540 map[string]any) []map[string]any {
		hx_field_541 := hx_obj_540["args"]
		if hx_field_541 == nil {
			var hx_zero_542 []map[string]any
			return hx_zero_542
		}
		return hx_field_541.([]map[string]any)
	}(c2) == nil) {
		return false
	}
	if (func(hx_obj_543 map[string]any) []map[string]any {
		hx_field_544 := hx_obj_543["args"]
		if hx_field_544 == nil {
			var hx_zero_545 []map[string]any
			return hx_zero_545
		}
		return hx_field_544.([]map[string]any)
	}(c1) != nil) && !haxe__rtti__TypeApi_sameConstructorArguments(func(hx_obj_546 map[string]any) []map[string]any {
		hx_field_547 := hx_obj_546["args"]
		if hx_field_547 == nil {
			var hx_zero_548 []map[string]any
			return hx_zero_548
		}
		return hx_field_547.([]map[string]any)
	}(c1), func(hx_obj_549 map[string]any) []map[string]any {
		hx_field_550 := hx_obj_549["args"]
		if hx_field_550 == nil {
			var hx_zero_551 []map[string]any
			return hx_zero_551
		}
		return hx_field_550.([]map[string]any)
	}(c2)) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_fieldEq(f1 map[string]any, f2 map[string]any) bool {
	if !hxrt.StringEqualStringPtr(func(hx_obj_552 map[string]any) *string {
		hx_field_553 := hx_obj_552["name"]
		if hx_field_553 == nil {
			var hx_zero_554 *string
			return hx_zero_554
		}
		return hx_field_553.(*string)
	}(f1), func(hx_obj_555 map[string]any) *string {
		hx_field_556 := hx_obj_555["name"]
		if hx_field_556 == nil {
			var hx_zero_557 *string
			return hx_zero_557
		}
		return hx_field_556.(*string)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_typeEq(func(hx_obj_558 map[string]any) *haxe__rtti__CType {
		hx_field_559 := hx_obj_558["type"]
		if hx_field_559 == nil {
			var hx_zero_560 *haxe__rtti__CType
			return hx_zero_560
		}
		return hx_field_559.(*haxe__rtti__CType)
	}(f1), func(hx_obj_561 map[string]any) *haxe__rtti__CType {
		hx_field_562 := hx_obj_561["type"]
		if hx_field_562 == nil {
			var hx_zero_563 *haxe__rtti__CType
			return hx_zero_563
		}
		return hx_field_562.(*haxe__rtti__CType)
	}(f2)) {
		return false
	}
	if func(hx_obj_564 map[string]any) bool {
		hx_field_565 := hx_obj_564["isPublic"]
		if hx_field_565 == nil {
			var hx_zero_566 bool
			return hx_zero_566
		}
		return hx_field_565.(bool)
	}(f1) != func(hx_obj_567 map[string]any) bool {
		hx_field_568 := hx_obj_567["isPublic"]
		if hx_field_568 == nil {
			var hx_zero_569 bool
			return hx_zero_569
		}
		return hx_field_568.(bool)
	}(f2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_obj_570 map[string]any) *string {
		hx_field_571 := hx_obj_570["doc"]
		if hx_field_571 == nil {
			var hx_zero_572 *string
			return hx_zero_572
		}
		return hx_field_571.(*string)
	}(f1), func(hx_obj_573 map[string]any) *string {
		hx_field_574 := hx_obj_573["doc"]
		if hx_field_574 == nil {
			var hx_zero_575 *string
			return hx_zero_575
		}
		return hx_field_574.(*string)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_obj_576 map[string]any) *haxe__rtti__Rights {
		hx_field_577 := hx_obj_576["get"]
		if hx_field_577 == nil {
			var hx_zero_578 *haxe__rtti__Rights
			return hx_zero_578
		}
		return hx_field_577.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_579 map[string]any) *haxe__rtti__Rights {
		hx_field_580 := hx_obj_579["get"]
		if hx_field_580 == nil {
			var hx_zero_581 *haxe__rtti__Rights
			return hx_zero_581
		}
		return hx_field_580.(*haxe__rtti__Rights)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_obj_582 map[string]any) *haxe__rtti__Rights {
		hx_field_583 := hx_obj_582["set"]
		if hx_field_583 == nil {
			var hx_zero_584 *haxe__rtti__Rights
			return hx_zero_584
		}
		return hx_field_583.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_585 map[string]any) *haxe__rtti__Rights {
		hx_field_586 := hx_obj_585["set"]
		if hx_field_586 == nil {
			var hx_zero_587 *haxe__rtti__Rights
			return hx_zero_587
		}
		return hx_field_586.(*haxe__rtti__Rights)
	}(f2)) {
		return false
	}
	if (func(hx_obj_588 map[string]any) []*string {
		hx_field_589 := hx_obj_588["params"]
		if hx_field_589 == nil {
			var hx_zero_590 []*string
			return hx_zero_590
		}
		return hx_field_589.([]*string)
	}(f1) == nil) != (func(hx_obj_591 map[string]any) []*string {
		hx_field_592 := hx_obj_591["params"]
		if hx_field_592 == nil {
			var hx_zero_593 []*string
			return hx_zero_593
		}
		return hx_field_592.([]*string)
	}(f2) == nil) {
		return false
	}
	if (func(hx_obj_594 map[string]any) []*string {
		hx_field_595 := hx_obj_594["params"]
		if hx_field_595 == nil {
			var hx_zero_596 []*string
			return hx_zero_596
		}
		return hx_field_595.([]*string)
	}(f1) != nil) && !haxe__rtti__TypeApi_sameTypeParamNames(func(hx_obj_597 map[string]any) []*string {
		hx_field_598 := hx_obj_597["params"]
		if hx_field_598 == nil {
			var hx_zero_599 []*string
			return hx_zero_599
		}
		return hx_field_598.([]*string)
	}(f1), func(hx_obj_600 map[string]any) []*string {
		hx_field_601 := hx_obj_600["params"]
		if hx_field_601 == nil {
			var hx_zero_602 []*string
			return hx_zero_602
		}
		return hx_field_601.([]*string)
	}(f2)) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_isVar(t *haxe__rtti__CType) bool {
	var hx_if_603 bool
	if t.tag == 4 {
		_g := t.params[0].([]map[string]any)
		_ = _g
		_g_1 := t.params[1].(*haxe__rtti__CType)
		_ = _g_1
		hx_if_603 = false
	} else {
		hx_if_603 = true
	}
	return hx_if_603
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
		hx_post_604 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_604
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
		hx_post_605 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_605
		a := l1[i]
		b := l2[i]
		if (!hxrt.StringEqualStringPtr(func(hx_obj_606 map[string]any) *string {
			hx_field_607 := hx_obj_606["name"]
			if hx_field_607 == nil {
				var hx_zero_608 *string
				return hx_zero_608
			}
			return hx_field_607.(*string)
		}(a), func(hx_obj_609 map[string]any) *string {
			hx_field_610 := hx_obj_609["name"]
			if hx_field_610 == nil {
				var hx_zero_611 *string
				return hx_zero_611
			}
			return hx_field_610.(*string)
		}(b)) || (func(hx_obj_612 map[string]any) bool {
			hx_field_613 := hx_obj_612["opt"]
			if hx_field_613 == nil {
				var hx_zero_614 bool
				return hx_zero_614
			}
			return hx_field_613.(bool)
		}(a) != func(hx_obj_615 map[string]any) bool {
			hx_field_616 := hx_obj_615["opt"]
			if hx_field_616 == nil {
				var hx_zero_617 bool
				return hx_zero_617
			}
			return hx_field_616.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_618 map[string]any) *haxe__rtti__CType {
			hx_field_619 := hx_obj_618["t"]
			if hx_field_619 == nil {
				var hx_zero_620 *haxe__rtti__CType
				return hx_zero_620
			}
			return hx_field_619.(*haxe__rtti__CType)
		}(a), func(hx_obj_621 map[string]any) *haxe__rtti__CType {
			hx_field_622 := hx_obj_621["t"]
			if hx_field_622 == nil {
				var hx_zero_623 *haxe__rtti__CType
				return hx_zero_623
			}
			return hx_field_622.(*haxe__rtti__CType)
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
		hx_post_624 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_624
		a := l1[i]
		b := l2[i]
		if (!hxrt.StringEqualStringPtr(func(hx_obj_625 map[string]any) *string {
			hx_field_626 := hx_obj_625["name"]
			if hx_field_626 == nil {
				var hx_zero_627 *string
				return hx_zero_627
			}
			return hx_field_626.(*string)
		}(a), func(hx_obj_628 map[string]any) *string {
			hx_field_629 := hx_obj_628["name"]
			if hx_field_629 == nil {
				var hx_zero_630 *string
				return hx_zero_630
			}
			return hx_field_629.(*string)
		}(b)) || (func(hx_obj_631 map[string]any) bool {
			hx_field_632 := hx_obj_631["opt"]
			if hx_field_632 == nil {
				var hx_zero_633 bool
				return hx_zero_633
			}
			return hx_field_632.(bool)
		}(a) != func(hx_obj_634 map[string]any) bool {
			hx_field_635 := hx_obj_634["opt"]
			if hx_field_635 == nil {
				var hx_zero_636 bool
				return hx_zero_636
			}
			return hx_field_635.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_637 map[string]any) *haxe__rtti__CType {
			hx_field_638 := hx_obj_637["t"]
			if hx_field_638 == nil {
				var hx_zero_639 *haxe__rtti__CType
				return hx_zero_639
			}
			return hx_field_638.(*haxe__rtti__CType)
		}(a), func(hx_obj_640 map[string]any) *haxe__rtti__CType {
			hx_field_641 := hx_obj_640["t"]
			if hx_field_641 == nil {
				var hx_zero_642 *haxe__rtti__CType
				return hx_zero_642
			}
			return hx_field_641.(*haxe__rtti__CType)
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
		hx_post_643 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_643
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
		hx_post_644 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_644
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
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func(hx_obj_646 map[string]any) *string {
		hx_field_647 := hx_obj_646["name"]
		if hx_field_647 == nil {
			var hx_zero_648 *string
			return hx_zero_648
		}
		return hx_field_647.(*string)
	}(cf), hxrt.StringFromLiteral(":")), haxe__rtti__CTypeTools_toString(func(hx_obj_649 map[string]any) *haxe__rtti__CType {
		hx_field_650 := hx_obj_649["type"]
		if hx_field_650 == nil {
			var hx_zero_651 *haxe__rtti__CType
			return hx_zero_651
		}
		return hx_field_650.(*haxe__rtti__CType)
	}(cf)))
}

func haxe__rtti__CTypeTools_functionArgumentName(arg map[string]any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_if_655 *string
		if func(hx_obj_652 map[string]any) bool {
			hx_field_653 := hx_obj_652["opt"]
			if hx_field_653 == nil {
				var hx_zero_654 bool
				return hx_zero_654
			}
			return hx_field_653.(bool)
		}(arg) {
			hx_if_655 = hxrt.StringFromLiteral("?")
		} else {
			hx_if_655 = hxrt.StringFromLiteral("")
		}
		return hx_if_655
	}(), func() *string {
		var hx_if_662 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_656 map[string]any) *string {
			hx_field_657 := hx_obj_656["name"]
			if hx_field_657 == nil {
				var hx_zero_658 *string
				return hx_zero_658
			}
			return hx_field_657.(*string)
		}(arg), hxrt.StringFromLiteral("")) {
			hx_if_662 = hxrt.StringFromLiteral("")
		} else {
			hx_if_662 = hxrt.StringConcatStringPtr(func(hx_obj_659 map[string]any) *string {
				hx_field_660 := hx_obj_659["name"]
				if hx_field_660 == nil {
					var hx_zero_661 *string
					return hx_zero_661
				}
				return hx_field_660.(*string)
			}(arg), hxrt.StringFromLiteral(":"))
		}
		return hx_if_662
	}()), haxe__rtti__CTypeTools_toString(func(hx_obj_663 map[string]any) *haxe__rtti__CType {
		hx_field_664 := hx_obj_663["t"]
		if hx_field_664 == nil {
			var hx_zero_665 *haxe__rtti__CType
			return hx_zero_665
		}
		return hx_field_664.(*haxe__rtti__CType)
	}(arg))), func() *string {
		var hx_if_672 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_666 map[string]any) *string {
			hx_field_667 := hx_obj_666["value"]
			if hx_field_667 == nil {
				var hx_zero_668 *string
				return hx_zero_668
			}
			return hx_field_667.(*string)
		}(arg), nil) {
			hx_if_672 = hxrt.StringFromLiteral("")
		} else {
			hx_if_672 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" = "), func(hx_obj_669 map[string]any) *string {
				hx_field_670 := hx_obj_669["value"]
				if hx_field_670 == nil {
					var hx_zero_671 *string
					return hx_zero_671
				}
				return hx_field_670.(*string)
			}(arg))
		}
		return hx_if_672
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
		hx_post_675 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_675
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
	var hx_switch_677 *string
	switch t.tag {
	case 0:
		hx_switch_677 = hxrt.StringFromLiteral("unknown")
	case 1:
		_g := t.params[0].(*string)
		_g1 := t.params[1].([]*haxe__rtti__CType)
		name := _g
		params := _g1
		hx_switch_677 = haxe__rtti__CTypeTools_nameWithParams(name, params)
	case 2:
		_g_1 := t.params[0].(*string)
		_g1_1 := t.params[1].([]*haxe__rtti__CType)
		name_1 := _g_1
		params_1 := _g1_1
		hx_switch_677 = haxe__rtti__CTypeTools_nameWithParams(name_1, params_1)
	case 3:
		_g_2 := t.params[0].(*string)
		_g1_2 := t.params[1].([]*haxe__rtti__CType)
		name_2 := _g_2
		params_2 := _g1_2
		hx_switch_677 = haxe__rtti__CTypeTools_nameWithParams(name_2, params_2)
	case 4:
		_g_3 := t.params[0].([]map[string]any)
		_g1_3 := t.params[1].(*haxe__rtti__CType)
		args := _g_3
		ret := _g1_3
		var hx_if_678 *string
		if len(args) == 0 {
			hx_if_678 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Void -> "), haxe__rtti__CTypeTools_toString(ret))
		} else {
			hx_if_678 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(haxe__rtti__CTypeTools_joinFunctionArguments(args), hxrt.StringFromLiteral(" -> ")), haxe__rtti__CTypeTools_toString(ret))
		}
		hx_switch_677 = hx_if_678
	case 5:
		_g_4 := t.params[0].([]map[string]any)
		fields := _g_4
		hx_switch_677 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{ "), haxe__rtti__CTypeTools_joinClassFields(fields)), hxrt.StringFromLiteral("}"))
	case 6:
		_g_5 := t.params[0].(*haxe__rtti__CType)
		d := _g_5
		var hx_if_679 *string
		if d == nil {
			hx_if_679 = hxrt.StringFromLiteral("Dynamic")
		} else {
			hx_if_679 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Dynamic<"), haxe__rtti__CTypeTools_toString(d)), hxrt.StringFromLiteral(">"))
		}
		hx_switch_677 = hx_if_679
	case 7:
		_g_6 := t.params[0].(*string)
		_g1_4 := t.params[1].([]*haxe__rtti__CType)
		name_3 := _g_6
		params_3 := _g1_4
		hx_switch_677 = haxe__rtti__CTypeTools_nameWithParams(name_3, params_3)
	}
	return hx_switch_677
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
