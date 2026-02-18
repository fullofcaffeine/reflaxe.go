package main

import "examples_profile_storyboard_portable/hxrt"

func Harness_assertContract(runtime profile__StoryboardRuntime) *string {
	cards := Harness_buildCards()
	_ = cards
	baseline := Harness_baselineSummary(cards)
	_ = baseline
	if !hxrt.StringEqualAny(baseline, hxrt.StringFromLiteral("cards=3,points=13,open=3")) {
		hxrt.Throw(hxrt.StringConcatAny(hxrt.StringFromLiteral("baseline drift: "), baseline))
	}
	extra := runtime.extraSignal(cards)
	_ = extra
	if hxrt.StringEqualAny(extra, nil) || hxrt.StringEqualAny(extra, hxrt.StringFromLiteral("")) {
		hxrt.Throw(hxrt.StringFromLiteral("missing extra signal"))
	}
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("OK "), runtime.profileId())
}

func Harness_baselineSummary(cards *haxe__ds__List) *string {
	totalPoints := 0
	_ = totalPoints
	open := 0
	_ = open
	count := cards.length
	_ = count
	i := 0
	_ = i
	for i < count {
		value := cards.pop().(*domain__StoryCard)
		_ = value
		if hxrt.StringEqualAny(value, nil) {
			break
		}
		card := value
		_ = card
		totalPoints = (totalPoints + card.points)
		open = (open + 1)
		cards.add(card)
		i = (i + 1)
	}
	return hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("cards="), open), hxrt.StringFromLiteral(",points=")), totalPoints), hxrt.StringFromLiteral(",open=")), open)
}

func Harness_buildCards() *haxe__ds__List {
	cards := New_haxe__ds__List()
	_ = cards
	cards.add(Harness_card(1, hxrt.StringFromLiteral("Ship profile docs"), 3, Harness_makeTags(hxrt.StringFromLiteral("docs"), hxrt.StringFromLiteral("profiles"))))
	cards.add(Harness_card(2, hxrt.StringFromLiteral("Backfill regression snapshots"), 5, Harness_makeTags(hxrt.StringFromLiteral("tests"), nil)))
	cards.add(Harness_card(3, hxrt.StringFromLiteral("Wire release artifacts"), 5, Harness_makeTags(hxrt.StringFromLiteral("ci"), hxrt.StringFromLiteral("release"))))
	return cards
}

func Harness_card(id int, title *string, points int, tags *haxe__ds__List) *domain__StoryCard {
	return New_domain__StoryCard(id, title, points, tags)
}

func Harness_formatCards(cards *haxe__ds__List, runtime profile__StoryboardRuntime) *string {
	out := hxrt.StringFromLiteral("")
	_ = out
	firstCard := true
	_ = firstCard
	cardCount := cards.length
	_ = cardCount
	i := 0
	_ = i
	for i < cardCount {
		cardValue := cards.pop().(*domain__StoryCard)
		_ = cardValue
		if hxrt.StringEqualAny(cardValue, nil) {
			break
		}
		card := cardValue
		_ = card
		tags := New_haxe__ds__List()
		_ = tags
		tagCount := card.tags.length
		_ = tagCount
		j := 0
		_ = j
		for j < tagCount {
			tagValue := card.tags.pop().(*string)
			_ = tagValue
			if hxrt.StringEqualAny(tagValue, nil) {
				break
			}
			tag := tagValue
			_ = tag
			tags.add(runtime.highlightTag(tag))
			card.tags.add(tag)
			j = (j + 1)
		}
		if !firstCard {
			out = hxrt.StringConcatAny(out, hxrt.StringFromLiteral(";"))
		}
		out = hxrt.StringConcatAny(out, hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("#"), card.id), hxrt.StringFromLiteral(":")), runtime.decorateTitle(card.title)), hxrt.StringFromLiteral(":p")), card.points), hxrt.StringFromLiteral(":")), Harness_joinStringList(tags, hxrt.StringFromLiteral("|"))))
		firstCard = false
		cards.add(card)
		i = (i + 1)
	}
	return out
}

func Harness_joinStringList(values *haxe__ds__List, separator *string) *string {
	out := hxrt.StringFromLiteral("")
	_ = out
	first := true
	_ = first
	count := values.length
	_ = count
	i := 0
	_ = i
	for i < count {
		raw := values.pop().(*string)
		_ = raw
		if hxrt.StringEqualAny(raw, nil) {
			break
		}
		value := raw
		_ = value
		if !first {
			out = hxrt.StringConcatAny(out, separator)
		}
		out = hxrt.StringConcatAny(out, value)
		values.add(value)
		first = false
		i = (i + 1)
	}
	return out
}

func Harness_makeTags(a *string, b *string) *haxe__ds__List {
	tags := New_haxe__ds__List()
	_ = tags
	tags.add(a)
	if !hxrt.StringEqualAny(b, nil) {
		tags.add(b)
	}
	return tags
}

func Harness_render(runtime profile__StoryboardRuntime) *string {
	cards := Harness_buildCards()
	_ = cards
	velocity := hxrt.StringFromLiteral("off")
	_ = velocity
	if runtime.supportsVelocityHint() {
		velocity = hxrt.StringFromLiteral("on")
	}
	return hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringConcatAny(hxrt.StringFromLiteral("profile="), runtime.profileId()), hxrt.StringFromLiteral("\nbaseline=")), Harness_baselineSummary(cards)), hxrt.StringFromLiteral("\nview=")), Harness_formatCards(cards, runtime)), hxrt.StringFromLiteral("\nextra=")), runtime.extraSignal(cards)), hxrt.StringFromLiteral("\nvelocity_hint=")), velocity)
}

func main() {
	var runtime profile__StoryboardRuntime = profile__RuntimeFactory_create()
	_ = runtime
	hxrt.Println(Harness_render(runtime))
}

type I_domain__StoryCard interface {
}

type domain__StoryCard struct {
	__hx_this I_domain__StoryCard
	id        int
	title     *string
	points    int
	tags      *haxe__ds__List
}

func New_domain__StoryCard(id int, title *string, points int, tags *haxe__ds__List) *domain__StoryCard {
	self := &domain__StoryCard{}
	self.__hx_this = self
	self.id = id
	self.title = title
	self.points = points
	self.tags = tags
	return self
}

type I_profile__PortableRuntime interface {
	profileId() *string
	decorateTitle(title *string) *string
	highlightTag(tag *string) *string
	extraSignal(cards *haxe__ds__List) *string
	supportsVelocityHint() bool
}

type profile__PortableRuntime struct {
	__hx_this I_profile__PortableRuntime
}

func New_profile__PortableRuntime() *profile__PortableRuntime {
	self := &profile__PortableRuntime{}
	self.__hx_this = self
	return self
}

func (self *profile__PortableRuntime) profileId() *string {
	return hxrt.StringFromLiteral("portable")
}

func (self *profile__PortableRuntime) decorateTitle(title *string) *string {
	return title
}

func (self *profile__PortableRuntime) highlightTag(tag *string) *string {
	return tag
}

func (self *profile__PortableRuntime) extraSignal(cards *haxe__ds__List) *string {
	return hxrt.StringFromLiteral("interop_lane=off")
}

func (self *profile__PortableRuntime) supportsVelocityHint() bool {
	return false
}

func profile__RuntimeFactory_create() profile__StoryboardRuntime {
	return New_profile__PortableRuntime()
}

type profile__StoryboardRuntime interface {
	profileId() *string
	decorateTitle(title *string) *string
	highlightTag(tag *string) *string
	extraSignal(cards *haxe__ds__List) *string
	supportsVelocityHint() bool
}

type haxe__ds__IntMap struct {
	h map[int]any
}

type haxe__ds__StringMap struct {
	h map[string]any
}

type haxe__ds__ObjectMap struct {
	h map[any]any
}

type haxe__ds__EnumValueMap struct {
	h map[any]any
}

type haxe__ds__List struct {
	items  []any
	length int
}

func New_haxe__ds__IntMap() *haxe__ds__IntMap {
	return &haxe__ds__IntMap{h: map[int]any{}}
}

func (self *haxe__ds__IntMap) set(key int, value any) {
	self.h[key] = value
}

func (self *haxe__ds__IntMap) get(key int) any {
	value := self.h[key]
	return value
}

func (self *haxe__ds__IntMap) exists(key int) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__IntMap) remove(key int) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	return &haxe__ds__StringMap{h: map[string]any{}}
}

func (self *haxe__ds__StringMap) set(key *string, value any) {
	self.h[*hxrt.StdString(key)] = value
}

func (self *haxe__ds__StringMap) get(key *string) any {
	value := self.h[*hxrt.StdString(key)]
	return value
}

func (self *haxe__ds__StringMap) exists(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	return ok
}

func (self *haxe__ds__StringMap) remove(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	delete(self.h, *hxrt.StdString(key))
	return ok
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	return &haxe__ds__ObjectMap{h: map[any]any{}}
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__EnumValueMap() *haxe__ds__EnumValueMap {
	return &haxe__ds__EnumValueMap{h: map[any]any{}}
}

func (self *haxe__ds__EnumValueMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__EnumValueMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__EnumValueMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__EnumValueMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__List() *haxe__ds__List {
	return &haxe__ds__List{items: []any{}, length: 0}
}

func (self *haxe__ds__List) add(item any) {
	self.items = append(self.items, item)
	self.length = len(self.items)
}

func (self *haxe__ds__List) push(item any) {
	self.add(item)
}

func (self *haxe__ds__List) pop() any {
	if len(self.items) == 0 {
		return nil
	}
	head := self.items[0]
	self.items = self.items[1:]
	self.length = len(self.items)
	return head
}

func (self *haxe__ds__List) first() any {
	if len(self.items) == 0 {
		return nil
	}
	return self.items[0]
}

func (self *haxe__ds__List) last() any {
	size := len(self.items)
	if size == 0 {
		return nil
	}
	return self.items[(size - 1)]
}
