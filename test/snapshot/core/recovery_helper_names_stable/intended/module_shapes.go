package main

type I_Circle interface {
}

type Circle struct {
	*Shape
	__hx_this I_Circle
}

func (self *Circle) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Circle(size int) *Circle {
	return hx_init_Circle(hx_alloc_Circle(nil), size)
}

func hx_alloc_Circle(owner any) *Circle {
	self := &Circle{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Circle)
	}
	self.Shape = hx_alloc_Shape(self.__hx_this)
	return self
}

func hx_init_Circle(self *Circle, size int) *Circle {
	hx_init_Shape(self.Shape, size)
	return self
}

func (self *Circle) HxClassName() string {
	return "Circle"
}

type I_Shape interface {
}

type Shape struct {
	__hx_this I_Shape
	size      int
}

func (self *Shape) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Shape(size int) *Shape {
	return hx_init_Shape(hx_alloc_Shape(nil), size)
}

func hx_alloc_Shape(owner any) *Shape {
	self := &Shape{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Shape)
	}
	return self
}

func hx_init_Shape(self *Shape, size int) *Shape {
	self.size = size
	return self
}

func (self *Shape) HxClassName() string {
	return "Shape"
}

type I_Square interface {
}

type Square struct {
	*Shape
	__hx_this I_Square
}

func (self *Square) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Square(size int) *Square {
	return hx_init_Square(hx_alloc_Square(nil), size)
}

func hx_alloc_Square(owner any) *Square {
	self := &Square{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Square)
	}
	self.Shape = hx_alloc_Shape(self.__hx_this)
	return self
}

func hx_init_Square(self *Square, size int) *Square {
	hx_init_Shape(self.Shape, size)
	return self
}

func (self *Square) HxClassName() string {
	return "Square"
}
