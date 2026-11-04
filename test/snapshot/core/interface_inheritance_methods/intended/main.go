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

func (self *QueryService) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_QueryService() *QueryService {
	return hx_init_QueryService(hx_alloc_QueryService(nil))
}

func hx_alloc_QueryService(owner any) *QueryService {
	self := &QueryService{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_QueryService)
	}
	return self
}

func hx_init_QueryService(self *QueryService) *QueryService {
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
