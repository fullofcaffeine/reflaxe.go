package main

import "snapshot/hxrt"

type I_haxe__ds__GenericStack interface {
	add(item any)
	first() any
	pop() any
	isEmpty() bool
	remove(v any) bool
	iterator() map[string]any
	toString() *string
}

type haxe__ds__GenericStack struct {
	__hx_this I_haxe__ds__GenericStack
	head      *haxe__ds__GenericCell
}

func New_haxe__ds__GenericStack() *haxe__ds__GenericStack {
	self := &haxe__ds__GenericStack{}
	self.__hx_this = self
	return self
}

func (self *haxe__ds__GenericStack) add(item any) {
	self.head = New_haxe__ds__GenericCell(item, self.head)
}

func (self *haxe__ds__GenericStack) first() any {
	var hx_if_1 any
	if self.head == nil {
		hx_if_1 = nil
	} else {
		hx_if_1 = self.head.elt
	}
	return hx_if_1
}

func (self *haxe__ds__GenericStack) pop() any {
	current := self.head
	if current == nil {
		return nil
	}
	self.head = current.next
	return current.elt
}

func (self *haxe__ds__GenericStack) isEmpty() bool {
	return (self.head == nil)
}

func (self *haxe__ds__GenericStack) remove(v any) bool {
	var prev *haxe__ds__GenericCell = nil
	current := self.head
	for current != nil {
		if haxe__ds__GenericStack_sameValue(current.elt, v) {
			if prev == nil {
				self.head = current.next
			} else {
				prev.next = current.next
			}
			return true
		}
		prev = current
		current = current.next
	}
	return false
}

func (self *haxe__ds__GenericStack) iterator() map[string]any {
	current := self.head
	hx_obj_2 := map[string]any{}
	hx_obj_2["hasNext"] = func() bool {
		return (current != nil)
	}
	hx_obj_2["next"] = func() any {
		cell := current
		current = cell.next
		return cell.elt
	}
	return hx_obj_2
}

func (self *haxe__ds__GenericStack) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("{"))
	current := self.head
	firstItem := true
	for current != nil {
		if !firstItem {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(","))
		}
		x := hxrt.StdString(current.elt)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		firstItem = false
		current = current.next
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("}"))
	return out_b
}

func (self *haxe__ds__GenericStack) String() string {
	return *self.__hx_this.toString()
}

func haxe__ds__GenericStack_sameValue(left any, right any) bool {
	if hxrt.AnyEqualsNull(left) || hxrt.AnyEqualsNull(right) {
		return hxrt.HaxeEqual(left, right)
	}
	if func(hx_value any) bool {
		switch hx_value.(type) {
		case *string:
			return true
		case string:
			return true
		default:
			return false
		}
	}(any(left)) || func(hx_value any) bool {
		switch hx_value.(type) {
		case *string:
			return true
		case string:
			return true
		default:
			return false
		}
	}(any(right)) {
		return hxrt.StringEqualStringPtr(hxrt.StdString(left), hxrt.StdString(right))
	}
	return hxrt.HaxeEqual(left, right)
}

type I_haxe__ds__GenericCell interface {
}

type haxe__ds__GenericCell struct {
	__hx_this I_haxe__ds__GenericCell
	elt       any
	next      *haxe__ds__GenericCell
}

func New_haxe__ds__GenericCell(elt any, next *haxe__ds__GenericCell) *haxe__ds__GenericCell {
	self := &haxe__ds__GenericCell{}
	self.__hx_this = self
	self.elt = elt
	self.next = next
	return self
}
