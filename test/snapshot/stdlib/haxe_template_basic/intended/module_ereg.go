package main

import "snapshot/hxrt"

type I_EReg interface {
	match(s *string) bool
	matched(n int) *string
	matchedLeft() *string
	matchedRight() *string
	matchedPos() map[string]any
	matchSub(s *string, pos int, len int) bool
	split(s *string) *hxrt.Array
	replace(s *string, by *string) *string
	map_(s *string, f func(*EReg) *string) *string
	remember(source *string, match *hxrt.RegexMatch)
	requireMatch() *hxrt.RegexMatch
	expandReplacement(by *string, source *string, currentMatch *hxrt.RegexMatch) *string
}

type EReg struct {
	__hx_this  I_EReg
	handle     *hxrt.RegexHandle
	global     bool
	lastSource *string
	lastMatch  *hxrt.RegexMatch
}

func New_EReg(r *string, opt *string) *EReg {
	self := &EReg{}
	self.__hx_this = self
	self.handle = hxrt.RegexCompile(r, opt)
	self.global = StringTools_contains(opt, hxrt.StringFromLiteral("g"))
	self.lastSource = nil
	self.lastMatch = nil
	return self
}

func (self *EReg) match(s *string) bool {
	found := hxrt.RegexFind(self.handle, s, 0)
	if found == nil {
		self.lastSource = nil
		self.lastMatch = nil
		return false
	}
	self.lastSource = s
	self.lastMatch = found
	return true
}

func (self *EReg) matched(n int) *string {
	current := self.__hx_this.requireMatch()
	if n < 0 {
		hxrt.Throw(hxrt.StringFromLiteral("Invalid group"))
	}
	offset := int(int32((hxrt.Int32Wrap(n) * hxrt.Int32Wrap(2))))
	if int(int32((hxrt.Int32Wrap(offset) + hxrt.Int32Wrap(1)))) >= len(current.Indices) {
		hxrt.Throw(hxrt.StringFromLiteral("Invalid group"))
	}
	start := current.Indices[offset]
	end := current.Indices[int(int32((hxrt.Int32Wrap(offset) + hxrt.Int32Wrap(1))))]
	if (start < 0) || (end < start) {
		return nil
	}
	return hxrt.StringSubstrStringPtr(self.lastSource, start, int(int32((hxrt.Int32Wrap(end) - hxrt.Int32Wrap(start)))), true)
}

func (self *EReg) matchedLeft() *string {
	current := self.__hx_this.requireMatch()
	return hxrt.StringSubstrStringPtr(self.lastSource, 0, current.Indices[0], true)
}

func (self *EReg) matchedRight() *string {
	current := self.__hx_this.requireMatch()
	return hxrt.StringSubstrStringPtr(self.lastSource, current.Indices[1], 0, false)
}

func (self *EReg) matchedPos() map[string]any {
	current := self.__hx_this.requireMatch()
	start := current.Indices[0]
	hx_obj_276 := map[string]any{}
	hx_obj_276["pos"] = start
	hx_obj_276["len"] = int(int32((hxrt.Int32Wrap(current.Indices[1]) - hxrt.Int32Wrap(start))))
	return hx_obj_276
}

func (self *EReg) matchSub(s *string, pos int, len int) bool {
	var hx_if_277 int
	if pos < 0 {
		hx_if_277 = 0
	} else {
		hx_if_277 = pos
	}
	start := hx_if_277
	var hx_if_278 int
	if len < 0 {
		hx_if_278 = hxrt.StringLengthStringPtr(s)
	} else {
		hx_if_278 = int(int32((hxrt.Int32Wrap(start) + hxrt.Int32Wrap(len))))
	}
	limit := hx_if_278
	if limit > hxrt.StringLengthStringPtr(s) {
		limit = hxrt.StringLengthStringPtr(s)
	}
	if start > limit {
		return false
	}
	var hx_if_279 *string
	if limit == hxrt.StringLengthStringPtr(s) {
		hx_if_279 = s
	} else {
		hx_if_279 = hxrt.StringSubstrStringPtr(s, 0, limit, true)
	}
	searched := hx_if_279
	found := hxrt.RegexFind(self.handle, searched, start)
	if found == nil {
		return false
	}
	self.lastSource = s
	self.lastMatch = found
	return true
}

func (self *EReg) split(s *string) *hxrt.Array {
	if hxrt.StringLengthStringPtr(s) == 0 {
		return hxrt.NewArray(s)
	}
	if !self.global {
		first := hxrt.RegexFind(self.handle, s, 0)
		if first == nil {
			return hxrt.NewArray(s)
		}
		return hxrt.NewArray(hxrt.StringSubstrStringPtr(s, 0, first.Indices[0], true), hxrt.StringSubstrStringPtr(s, first.Indices[1], 0, false))
	}
	parts := hxrt.NewArray()
	copyOffset := 0
	searchStart := 0
	for true {
		current := hxrt.RegexFind(self.handle, s, searchStart)
		if current == nil {
			parts.Push(hxrt.StringSubstrStringPtr(s, copyOffset, 0, false))
			break
		}
		start := current.Indices[0]
		end := current.Indices[1]
		parts.Push(hxrt.StringSubstrStringPtr(s, copyOffset, int(int32((hxrt.Int32Wrap(start) - hxrt.Int32Wrap(copyOffset)))), true))
		copyOffset = end
		nextStart := end
		if (start == end) && (nextStart == searchStart) {
			nextStart = int(int32((nextStart + 1)))
		}
		if nextStart >= hxrt.StringLengthStringPtr(s) {
			parts.Push(hxrt.StringFromLiteral(""))
			break
		}
		searchStart = nextStart
	}
	return parts
}

func (self *EReg) replace(s *string, by *string) *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	copyOffset := 0
	searchStart := 0
	for true {
		current := hxrt.RegexFind(self.handle, s, searchStart)
		if current == nil {
			break
		}
		start := current.Indices[0]
		end := current.Indices[1]
		x := hxrt.StringSubstrStringPtr(s, copyOffset, int(int32((hxrt.Int32Wrap(start) - hxrt.Int32Wrap(copyOffset)))), true)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		x_1 := self.__hx_this.expandReplacement(by, s, current)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		copyOffset = end
		if !self.global {
			break
		}
		if start == end {
			if end >= hxrt.StringLengthStringPtr(s) {
				break
			}
			searchStart = int(int32((hxrt.Int32Wrap(end) + hxrt.Int32Wrap(1))))
		} else {
			searchStart = end
		}
	}
	x_2 := hxrt.StringSubstrStringPtr(s, copyOffset, 0, false)
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_2))
	return out_b
}

func (self *EReg) map_(s *string, f func(*EReg) *string) *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	offset := 0
	for offset < hxrt.StringLengthStringPtr(s) {
		current := hxrt.RegexFind(self.handle, s, offset)
		if current == nil {
			x := hxrt.StringSubstrStringPtr(s, offset, 0, false)
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
			break
		}
		start := current.Indices[0]
		end := current.Indices[1]
		x_1 := hxrt.StringSubstrStringPtr(s, offset, int(int32((hxrt.Int32Wrap(start) - hxrt.Int32Wrap(offset)))), true)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		self.lastSource = s
		self.lastMatch = current
		x_2 := f(self)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_2))
		if start == end {
			x_3 := hxrt.StringSubstrStringPtr(s, start, 1, true)
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_3))
			offset = int(int32((hxrt.Int32Wrap(start) + hxrt.Int32Wrap(1))))
		} else {
			offset = end
		}
		if !self.global {
			if offset < hxrt.StringLengthStringPtr(s) {
				x_4 := hxrt.StringSubstrStringPtr(s, offset, 0, false)
				out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_4))
			}
			break
		}
	}
	return out_b
}

func (self *EReg) remember(source *string, match *hxrt.RegexMatch) {
	self.lastSource = source
	self.lastMatch = match
}

func (self *EReg) requireMatch() *hxrt.RegexMatch {
	if (hxrt.StringEqualStringPtr(self.lastSource, nil) || (self.lastMatch == nil)) || (len(self.lastMatch.Indices) < 2) {
		hxrt.Throw(hxrt.StringFromLiteral("Invalid regex operation because no match was made"))
	}
	return self.lastMatch
}

func (self *EReg) expandReplacement(by *string, source *string, currentMatch *hxrt.RegexMatch) *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	index := 0
	for index < hxrt.StringLengthStringPtr(by) {
		current := hxrt.StringCharAtStringPtr(by, index)
		if hxrt.StringEqualStringPtr(current, hxrt.StringFromLiteral("$")) && (int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(1)))) < hxrt.StringLengthStringPtr(by)) {
			next := hxrt.StringCharAtStringPtr(by, int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(1)))))
			if hxrt.StringEqualStringPtr(next, hxrt.StringFromLiteral("$")) {
				out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("$"))
				index = int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(2))))
				continue
			}
			var hx_switch_283 int
			switch *hxrt.StdString(next) {
			case *hxrt.StdString(hxrt.StringFromLiteral("1")):
				hx_switch_283 = 1
			case *hxrt.StdString(hxrt.StringFromLiteral("2")):
				hx_switch_283 = 2
			case *hxrt.StdString(hxrt.StringFromLiteral("3")):
				hx_switch_283 = 3
			case *hxrt.StdString(hxrt.StringFromLiteral("4")):
				hx_switch_283 = 4
			case *hxrt.StdString(hxrt.StringFromLiteral("5")):
				hx_switch_283 = 5
			case *hxrt.StdString(hxrt.StringFromLiteral("6")):
				hx_switch_283 = 6
			case *hxrt.StdString(hxrt.StringFromLiteral("7")):
				hx_switch_283 = 7
			case *hxrt.StdString(hxrt.StringFromLiteral("8")):
				hx_switch_283 = 8
			case *hxrt.StdString(hxrt.StringFromLiteral("9")):
				hx_switch_283 = 9
			default:
				hx_switch_283 = 0
			}
			group := hx_switch_283
			if group != 0 {
				offset := int(int32((hxrt.Int32Wrap(group) * hxrt.Int32Wrap(2))))
				if int(int32((hxrt.Int32Wrap(offset) + hxrt.Int32Wrap(1)))) >= len(currentMatch.Indices) {
					hxrt.Throw(hxrt.StringFromLiteral("Invalid group"))
				}
				start := currentMatch.Indices[offset]
				end := currentMatch.Indices[int(int32((hxrt.Int32Wrap(offset) + hxrt.Int32Wrap(1))))]
				var hx_if_284 *string
				if (start < 0) || (end < start) {
					hx_if_284 = nil
				} else {
					hx_if_284 = hxrt.StringSubstrStringPtr(source, start, int(int32((hxrt.Int32Wrap(end) - hxrt.Int32Wrap(start)))), true)
				}
				value := hx_if_284
				if !hxrt.StringEqualStringPtr(value, nil) {
					out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(value))
				}
				index = int(int32((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(2))))
				continue
			}
		}
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(current))
		index = int(int32((index + 1)))
	}
	return out_b
}

func EReg_escape(s *string) *string {
	return hxrt.StdString(hxrt.RegexEscape(s))
}
