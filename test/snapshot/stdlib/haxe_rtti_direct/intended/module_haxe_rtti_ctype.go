package main

import "snapshot/hxrt"

func haxe__rtti__TypeApi_constructorEq(c1 map[string]any, c2 map[string]any) bool {
	if !hxrt.StringEqualStringPtr(func(hx_obj_535 map[string]any) *string {
		hx_field_536 := hx_obj_535["name"]
		if hx_field_536 == nil {
			var hx_zero_537 *string
			return hx_zero_537
		}
		return hx_field_536.(*string)
	}(c1), func(hx_obj_538 map[string]any) *string {
		hx_field_539 := hx_obj_538["name"]
		if hx_field_539 == nil {
			var hx_zero_540 *string
			return hx_zero_540
		}
		return hx_field_539.(*string)
	}(c2)) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_obj_541 map[string]any) *string {
		hx_field_542 := hx_obj_541["doc"]
		if hx_field_542 == nil {
			var hx_zero_543 *string
			return hx_zero_543
		}
		return hx_field_542.(*string)
	}(c1), func(hx_obj_544 map[string]any) *string {
		hx_field_545 := hx_obj_544["doc"]
		if hx_field_545 == nil {
			var hx_zero_546 *string
			return hx_zero_546
		}
		return hx_field_545.(*string)
	}(c2)) {
		return false
	}
	if (func(hx_obj_547 map[string]any) *hxrt.Array {
		hx_field_548 := hx_obj_547["args"]
		if hx_field_548 == nil {
			var hx_zero_549 *hxrt.Array
			return hx_zero_549
		}
		return hx_field_548.(*hxrt.Array)
	}(c1) == nil) != (func(hx_obj_550 map[string]any) *hxrt.Array {
		hx_field_551 := hx_obj_550["args"]
		if hx_field_551 == nil {
			var hx_zero_552 *hxrt.Array
			return hx_zero_552
		}
		return hx_field_551.(*hxrt.Array)
	}(c2) == nil) {
		return false
	}
	if (func(hx_obj_553 map[string]any) *hxrt.Array {
		hx_field_554 := hx_obj_553["args"]
		if hx_field_554 == nil {
			var hx_zero_555 *hxrt.Array
			return hx_zero_555
		}
		return hx_field_554.(*hxrt.Array)
	}(c1) != nil) && !haxe__rtti__TypeApi_sameConstructorArguments(func(hx_obj_556 map[string]any) *hxrt.Array {
		hx_field_557 := hx_obj_556["args"]
		if hx_field_557 == nil {
			var hx_zero_558 *hxrt.Array
			return hx_zero_558
		}
		return hx_field_557.(*hxrt.Array)
	}(c1), func(hx_obj_559 map[string]any) *hxrt.Array {
		hx_field_560 := hx_obj_559["args"]
		if hx_field_560 == nil {
			var hx_zero_561 *hxrt.Array
			return hx_zero_561
		}
		return hx_field_560.(*hxrt.Array)
	}(c2)) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_fieldEq(f1 map[string]any, f2 map[string]any) bool {
	if !hxrt.StringEqualStringPtr(func(hx_obj_562 map[string]any) *string {
		hx_field_563 := hx_obj_562["name"]
		if hx_field_563 == nil {
			var hx_zero_564 *string
			return hx_zero_564
		}
		return hx_field_563.(*string)
	}(f1), func(hx_obj_565 map[string]any) *string {
		hx_field_566 := hx_obj_565["name"]
		if hx_field_566 == nil {
			var hx_zero_567 *string
			return hx_zero_567
		}
		return hx_field_566.(*string)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_typeEq(func(hx_obj_568 map[string]any) *haxe__rtti__CType {
		hx_field_569 := hx_obj_568["type"]
		if hx_field_569 == nil {
			var hx_zero_570 *haxe__rtti__CType
			return hx_zero_570
		}
		return hx_field_569.(*haxe__rtti__CType)
	}(f1), func(hx_obj_571 map[string]any) *haxe__rtti__CType {
		hx_field_572 := hx_obj_571["type"]
		if hx_field_572 == nil {
			var hx_zero_573 *haxe__rtti__CType
			return hx_zero_573
		}
		return hx_field_572.(*haxe__rtti__CType)
	}(f2)) {
		return false
	}
	if func(hx_obj_574 map[string]any) bool {
		hx_field_575 := hx_obj_574["isPublic"]
		if hx_field_575 == nil {
			var hx_zero_576 bool
			return hx_zero_576
		}
		return hx_field_575.(bool)
	}(f1) != func(hx_obj_577 map[string]any) bool {
		hx_field_578 := hx_obj_577["isPublic"]
		if hx_field_578 == nil {
			var hx_zero_579 bool
			return hx_zero_579
		}
		return hx_field_578.(bool)
	}(f2) {
		return false
	}
	if !hxrt.StringEqualStringPtr(func(hx_obj_580 map[string]any) *string {
		hx_field_581 := hx_obj_580["doc"]
		if hx_field_581 == nil {
			var hx_zero_582 *string
			return hx_zero_582
		}
		return hx_field_581.(*string)
	}(f1), func(hx_obj_583 map[string]any) *string {
		hx_field_584 := hx_obj_583["doc"]
		if hx_field_584 == nil {
			var hx_zero_585 *string
			return hx_zero_585
		}
		return hx_field_584.(*string)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_obj_586 map[string]any) *haxe__rtti__Rights {
		hx_field_587 := hx_obj_586["get"]
		if hx_field_587 == nil {
			var hx_zero_588 *haxe__rtti__Rights
			return hx_zero_588
		}
		return hx_field_587.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_589 map[string]any) *haxe__rtti__Rights {
		hx_field_590 := hx_obj_589["get"]
		if hx_field_590 == nil {
			var hx_zero_591 *haxe__rtti__Rights
			return hx_zero_591
		}
		return hx_field_590.(*haxe__rtti__Rights)
	}(f2)) {
		return false
	}
	if !haxe__rtti__TypeApi_rightsEq(func(hx_obj_592 map[string]any) *haxe__rtti__Rights {
		hx_field_593 := hx_obj_592["set"]
		if hx_field_593 == nil {
			var hx_zero_594 *haxe__rtti__Rights
			return hx_zero_594
		}
		return hx_field_593.(*haxe__rtti__Rights)
	}(f1), func(hx_obj_595 map[string]any) *haxe__rtti__Rights {
		hx_field_596 := hx_obj_595["set"]
		if hx_field_596 == nil {
			var hx_zero_597 *haxe__rtti__Rights
			return hx_zero_597
		}
		return hx_field_596.(*haxe__rtti__Rights)
	}(f2)) {
		return false
	}
	if (func(hx_obj_598 map[string]any) *hxrt.Array {
		hx_field_599 := hx_obj_598["params"]
		if hx_field_599 == nil {
			var hx_zero_600 *hxrt.Array
			return hx_zero_600
		}
		return hx_field_599.(*hxrt.Array)
	}(f1) == nil) != (func(hx_obj_601 map[string]any) *hxrt.Array {
		hx_field_602 := hx_obj_601["params"]
		if hx_field_602 == nil {
			var hx_zero_603 *hxrt.Array
			return hx_zero_603
		}
		return hx_field_602.(*hxrt.Array)
	}(f2) == nil) {
		return false
	}
	if (func(hx_obj_604 map[string]any) *hxrt.Array {
		hx_field_605 := hx_obj_604["params"]
		if hx_field_605 == nil {
			var hx_zero_606 *hxrt.Array
			return hx_zero_606
		}
		return hx_field_605.(*hxrt.Array)
	}(f1) != nil) && !haxe__rtti__TypeApi_sameTypeParamNames(func(hx_obj_607 map[string]any) *hxrt.Array {
		hx_field_608 := hx_obj_607["params"]
		if hx_field_608 == nil {
			var hx_zero_609 *hxrt.Array
			return hx_zero_609
		}
		return hx_field_608.(*hxrt.Array)
	}(f1), func(hx_obj_610 map[string]any) *hxrt.Array {
		hx_field_611 := hx_obj_610["params"]
		if hx_field_611 == nil {
			var hx_zero_612 *hxrt.Array
			return hx_zero_612
		}
		return hx_field_611.(*hxrt.Array)
	}(f2)) {
		return false
	}
	return true
}

func haxe__rtti__TypeApi_isVar(t *haxe__rtti__CType) bool {
	var hx_if_613 bool
	if t.tag == 4 {
		_g := t.params[0].(*hxrt.Array)
		_ = _g
		_g_1 := t.params[1].(*haxe__rtti__CType)
		_ = _g_1
		hx_if_613 = false
	} else {
		hx_if_613 = true
	}
	return hx_if_613
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
		hx_post_614 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_614
		if !haxe__rtti__TypeApi_fieldEq(func(hx_value_615 any) map[string]any {
			if hx_value_615 == nil {
				var hx_zero_616 map[string]any
				return hx_zero_616
			}
			return hx_value_615.(map[string]any)
		}(l1.Get(i)), func(hx_value_617 any) map[string]any {
			if hx_value_617 == nil {
				var hx_zero_618 map[string]any
				return hx_zero_618
			}
			return hx_value_617.(map[string]any)
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
		hx_post_619 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_619
		a := func(hx_value_620 any) map[string]any {
			if hx_value_620 == nil {
				var hx_zero_621 map[string]any
				return hx_zero_621
			}
			return hx_value_620.(map[string]any)
		}(l1.Get(i))
		b := func(hx_value_622 any) map[string]any {
			if hx_value_622 == nil {
				var hx_zero_623 map[string]any
				return hx_zero_623
			}
			return hx_value_622.(map[string]any)
		}(l2.Get(i))
		if (!hxrt.StringEqualStringPtr(func(hx_obj_624 map[string]any) *string {
			hx_field_625 := hx_obj_624["name"]
			if hx_field_625 == nil {
				var hx_zero_626 *string
				return hx_zero_626
			}
			return hx_field_625.(*string)
		}(a), func(hx_obj_627 map[string]any) *string {
			hx_field_628 := hx_obj_627["name"]
			if hx_field_628 == nil {
				var hx_zero_629 *string
				return hx_zero_629
			}
			return hx_field_628.(*string)
		}(b)) || (func(hx_obj_630 map[string]any) bool {
			hx_field_631 := hx_obj_630["opt"]
			if hx_field_631 == nil {
				var hx_zero_632 bool
				return hx_zero_632
			}
			return hx_field_631.(bool)
		}(a) != func(hx_obj_633 map[string]any) bool {
			hx_field_634 := hx_obj_633["opt"]
			if hx_field_634 == nil {
				var hx_zero_635 bool
				return hx_zero_635
			}
			return hx_field_634.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_636 map[string]any) *haxe__rtti__CType {
			hx_field_637 := hx_obj_636["t"]
			if hx_field_637 == nil {
				var hx_zero_638 *haxe__rtti__CType
				return hx_zero_638
			}
			return hx_field_637.(*haxe__rtti__CType)
		}(a), func(hx_obj_639 map[string]any) *haxe__rtti__CType {
			hx_field_640 := hx_obj_639["t"]
			if hx_field_640 == nil {
				var hx_zero_641 *haxe__rtti__CType
				return hx_zero_641
			}
			return hx_field_640.(*haxe__rtti__CType)
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
		hx_post_642 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_642
		a := func(hx_value_643 any) map[string]any {
			if hx_value_643 == nil {
				var hx_zero_644 map[string]any
				return hx_zero_644
			}
			return hx_value_643.(map[string]any)
		}(l1.Get(i))
		b := func(hx_value_645 any) map[string]any {
			if hx_value_645 == nil {
				var hx_zero_646 map[string]any
				return hx_zero_646
			}
			return hx_value_645.(map[string]any)
		}(l2.Get(i))
		if (!hxrt.StringEqualStringPtr(func(hx_obj_647 map[string]any) *string {
			hx_field_648 := hx_obj_647["name"]
			if hx_field_648 == nil {
				var hx_zero_649 *string
				return hx_zero_649
			}
			return hx_field_648.(*string)
		}(a), func(hx_obj_650 map[string]any) *string {
			hx_field_651 := hx_obj_650["name"]
			if hx_field_651 == nil {
				var hx_zero_652 *string
				return hx_zero_652
			}
			return hx_field_651.(*string)
		}(b)) || (func(hx_obj_653 map[string]any) bool {
			hx_field_654 := hx_obj_653["opt"]
			if hx_field_654 == nil {
				var hx_zero_655 bool
				return hx_zero_655
			}
			return hx_field_654.(bool)
		}(a) != func(hx_obj_656 map[string]any) bool {
			hx_field_657 := hx_obj_656["opt"]
			if hx_field_657 == nil {
				var hx_zero_658 bool
				return hx_zero_658
			}
			return hx_field_657.(bool)
		}(b))) || !haxe__rtti__TypeApi_typeEq(func(hx_obj_659 map[string]any) *haxe__rtti__CType {
			hx_field_660 := hx_obj_659["t"]
			if hx_field_660 == nil {
				var hx_zero_661 *haxe__rtti__CType
				return hx_zero_661
			}
			return hx_field_660.(*haxe__rtti__CType)
		}(a), func(hx_obj_662 map[string]any) *haxe__rtti__CType {
			hx_field_663 := hx_obj_662["t"]
			if hx_field_663 == nil {
				var hx_zero_664 *haxe__rtti__CType
				return hx_zero_664
			}
			return hx_field_663.(*haxe__rtti__CType)
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
		hx_post_665 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_665
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
		hx_post_666 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_666
		if !haxe__rtti__TypeApi_typeEq(func(hx_value_667 any) *haxe__rtti__CType {
			if hx_value_667 == nil {
				var hx_zero_668 *haxe__rtti__CType
				return hx_zero_668
			}
			return hx_value_667.(*haxe__rtti__CType)
		}(l1.Get(i)), func(hx_value_669 any) *haxe__rtti__CType {
			if hx_value_669 == nil {
				var hx_zero_670 *haxe__rtti__CType
				return hx_zero_670
			}
			return hx_value_669.(*haxe__rtti__CType)
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
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func(hx_obj_674 map[string]any) *string {
		hx_field_675 := hx_obj_674["name"]
		if hx_field_675 == nil {
			var hx_zero_676 *string
			return hx_zero_676
		}
		return hx_field_675.(*string)
	}(cf), hxrt.StringFromLiteral(":")), haxe__rtti__CTypeTools_toString(func(hx_obj_677 map[string]any) *haxe__rtti__CType {
		hx_field_678 := hx_obj_677["type"]
		if hx_field_678 == nil {
			var hx_zero_679 *haxe__rtti__CType
			return hx_zero_679
		}
		return hx_field_678.(*haxe__rtti__CType)
	}(cf)))
}

func haxe__rtti__CTypeTools_functionArgumentName(arg map[string]any) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_if_683 *string
		if func(hx_obj_680 map[string]any) bool {
			hx_field_681 := hx_obj_680["opt"]
			if hx_field_681 == nil {
				var hx_zero_682 bool
				return hx_zero_682
			}
			return hx_field_681.(bool)
		}(arg) {
			hx_if_683 = hxrt.StringFromLiteral("?")
		} else {
			hx_if_683 = hxrt.StringFromLiteral("")
		}
		return hx_if_683
	}(), func() *string {
		var hx_if_690 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_684 map[string]any) *string {
			hx_field_685 := hx_obj_684["name"]
			if hx_field_685 == nil {
				var hx_zero_686 *string
				return hx_zero_686
			}
			return hx_field_685.(*string)
		}(arg), hxrt.StringFromLiteral("")) {
			hx_if_690 = hxrt.StringFromLiteral("")
		} else {
			hx_if_690 = hxrt.StringConcatStringPtr(func(hx_obj_687 map[string]any) *string {
				hx_field_688 := hx_obj_687["name"]
				if hx_field_688 == nil {
					var hx_zero_689 *string
					return hx_zero_689
				}
				return hx_field_688.(*string)
			}(arg), hxrt.StringFromLiteral(":"))
		}
		return hx_if_690
	}()), haxe__rtti__CTypeTools_toString(func(hx_obj_691 map[string]any) *haxe__rtti__CType {
		hx_field_692 := hx_obj_691["t"]
		if hx_field_692 == nil {
			var hx_zero_693 *haxe__rtti__CType
			return hx_zero_693
		}
		return hx_field_692.(*haxe__rtti__CType)
	}(arg))), func() *string {
		var hx_if_700 *string
		if hxrt.StringEqualStringPtr(func(hx_obj_694 map[string]any) *string {
			hx_field_695 := hx_obj_694["value"]
			if hx_field_695 == nil {
				var hx_zero_696 *string
				return hx_zero_696
			}
			return hx_field_695.(*string)
		}(arg), nil) {
			hx_if_700 = hxrt.StringFromLiteral("")
		} else {
			hx_if_700 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(" = "), func(hx_obj_697 map[string]any) *string {
				hx_field_698 := hx_obj_697["value"]
				if hx_field_698 == nil {
					var hx_zero_699 *string
					return hx_zero_699
				}
				return hx_field_698.(*string)
			}(arg))
		}
		return hx_if_700
	}())
}

func haxe__rtti__CTypeTools_joinClassFields(fields *hxrt.Array) *string {
	parts := hxrt.NewArray()
	_g := 0
	for _g < fields.Len() {
		field := func(hx_value_701 any) map[string]any {
			if hx_value_701 == nil {
				var hx_zero_702 map[string]any
				return hx_zero_702
			}
			return hx_value_701.(map[string]any)
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
		arg := func(hx_value_704 any) map[string]any {
			if hx_value_704 == nil {
				var hx_zero_705 map[string]any
				return hx_zero_705
			}
			return hx_value_704.(map[string]any)
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
		hx_post_707 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_707
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
		param := func(hx_value_708 any) *haxe__rtti__CType {
			if hx_value_708 == nil {
				var hx_zero_709 *haxe__rtti__CType
				return hx_zero_709
			}
			return hx_value_708.(*haxe__rtti__CType)
		}(params.Get(_g))
		_g = int(int32((_g + 1)))
		parts.Push(haxe__rtti__CTypeTools_toString(param))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(name, hxrt.StringFromLiteral("<")), haxe__rtti__CTypeTools_joinStringArray(parts, hxrt.StringFromLiteral(", "))), hxrt.StringFromLiteral(">"))
}

func haxe__rtti__CTypeTools_toString(t *haxe__rtti__CType) *string {
	var hx_switch_711 *string
	switch t.tag {
	case 0:
		hx_switch_711 = hxrt.StringFromLiteral("unknown")
	case 1:
		_g := t.params[0].(*string)
		_g1 := t.params[1].(*hxrt.Array)
		name := _g
		params := _g1
		hx_switch_711 = haxe__rtti__CTypeTools_nameWithParams(name, params)
	case 2:
		_g_1 := t.params[0].(*string)
		_g1_1 := t.params[1].(*hxrt.Array)
		name_1 := _g_1
		params_1 := _g1_1
		hx_switch_711 = haxe__rtti__CTypeTools_nameWithParams(name_1, params_1)
	case 3:
		_g_2 := t.params[0].(*string)
		_g1_2 := t.params[1].(*hxrt.Array)
		name_2 := _g_2
		params_2 := _g1_2
		hx_switch_711 = haxe__rtti__CTypeTools_nameWithParams(name_2, params_2)
	case 4:
		_g_3 := t.params[0].(*hxrt.Array)
		_g1_3 := t.params[1].(*haxe__rtti__CType)
		args := _g_3
		ret := _g1_3
		var hx_if_712 *string
		if args.Len() == 0 {
			hx_if_712 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Void -> "), haxe__rtti__CTypeTools_toString(ret))
		} else {
			hx_if_712 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(haxe__rtti__CTypeTools_joinFunctionArguments(args), hxrt.StringFromLiteral(" -> ")), haxe__rtti__CTypeTools_toString(ret))
		}
		hx_switch_711 = hx_if_712
	case 5:
		_g_4 := t.params[0].(*hxrt.Array)
		fields := _g_4
		hx_switch_711 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("{ "), haxe__rtti__CTypeTools_joinClassFields(fields)), hxrt.StringFromLiteral("}"))
	case 6:
		_g_5 := t.params[0].(*haxe__rtti__CType)
		d := _g_5
		var hx_if_713 *string
		if d == nil {
			hx_if_713 = hxrt.StringFromLiteral("Dynamic")
		} else {
			hx_if_713 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Dynamic<"), haxe__rtti__CTypeTools_toString(d)), hxrt.StringFromLiteral(">"))
		}
		hx_switch_711 = hx_if_713
	case 7:
		_g_6 := t.params[0].(*string)
		_g1_4 := t.params[1].(*hxrt.Array)
		name_3 := _g_6
		params_3 := _g1_4
		hx_switch_711 = haxe__rtti__CTypeTools_nameWithParams(name_3, params_3)
	}
	return hx_switch_711
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
