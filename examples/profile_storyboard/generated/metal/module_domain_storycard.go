package main

type I_domain__StoryCard interface {
}

type domain__StoryCard struct {
	__hx_this I_domain__StoryCard
	id        int
	title     *string
	points    int
	tags      []*string
	state     *string
	owner     *string
}

func New_domain__StoryCard(id int, title *string, points int, tags []*string, state *string, owner *string) *domain__StoryCard {
	self := &domain__StoryCard{}
	self.__hx_this = self
	self.id = id
	self.title = title
	self.points = points
	self.tags = tags
	self.state = state
	self.owner = owner
	return self
}
