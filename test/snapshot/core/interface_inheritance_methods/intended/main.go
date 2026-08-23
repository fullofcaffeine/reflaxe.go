package main

import "snapshot/hxrt"

type CombinedQuery interface {
	root(value *string) *string
	right() *string
	left() *string
	local() *string
}

type LeftQuery interface {
	root(value *string) *string
	left() *string
}

func main() {
	printQuery(New_QueryService())
}

func printQuery(query CombinedQuery) {
	func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(query.root(hxrt.StringFromLiteral("root")))
	query.left()
	query.right()
	query.local()
}

type I_QueryService interface {
	root(value *string) *string
	left() *string
	right() *string
	local() *string
}

type QueryService struct {
	__hx_this I_QueryService
}

func New_QueryService() *QueryService {
	self := &QueryService{}
	self.__hx_this = self
	return self
}

func (self *QueryService) root(value *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("root:"), value)
}

func (self *QueryService) left() *string {
	return hxrt.StringFromLiteral("left")
}

func (self *QueryService) right() *string {
	return hxrt.StringFromLiteral("right")
}

func (self *QueryService) local() *string {
	return hxrt.StringFromLiteral("local")
}

type RightQuery interface {
	root(value *string) *string
	right() *string
}

type RootQuery interface {
	root(value any) any
}
