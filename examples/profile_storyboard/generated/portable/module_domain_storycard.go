package main

import "examples_profile_storyboard_portable/hxrt"

type I_domain__StoryCard interface {
}

type domain__StoryCard struct {
	__hx_this I_domain__StoryCard
	id        int
	title     *string
	points    int
	tags      *hxrt.Array
	state     *string
	owner     *string
}

func (self *domain__StoryCard) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_domain__StoryCard(id int, title *string, points int, tags *hxrt.Array, state *string, owner *string) *domain__StoryCard {
	return hx_init_domain__StoryCard(hx_alloc_domain__StoryCard(nil), id, title, points, tags, state, owner)
}

func hx_alloc_domain__StoryCard(owner any) *domain__StoryCard {
	self := &domain__StoryCard{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_domain__StoryCard)
	}
	return self
}

func hx_init_domain__StoryCard(self *domain__StoryCard, id int, title *string, points int, tags *hxrt.Array, state *string, owner *string) *domain__StoryCard {
	self.id = id
	self.title = title
	self.points = points
	self.tags = tags
	self.state = state
	self.owner = owner
	return self
}
