package main

type I_profile__TodoRuntimeMetrics interface {
}

type profile__TodoRuntimeMetrics struct {
	__hx_this I_profile__TodoRuntimeMetrics
	total     int
	open      int
	done      int
	p1        int
}

func (self *profile__TodoRuntimeMetrics) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_profile__TodoRuntimeMetrics(total int, open int, done int, p1 int) *profile__TodoRuntimeMetrics {
	return hx_init_profile__TodoRuntimeMetrics(hx_alloc_profile__TodoRuntimeMetrics(nil), total, open, done, p1)
}

func hx_alloc_profile__TodoRuntimeMetrics(owner any) *profile__TodoRuntimeMetrics {
	self := &profile__TodoRuntimeMetrics{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_profile__TodoRuntimeMetrics)
	}
	return self
}

func hx_init_profile__TodoRuntimeMetrics(self *profile__TodoRuntimeMetrics, total int, open int, done int, p1 int) *profile__TodoRuntimeMetrics {
	self.total = total
	self.open = open
	self.done = done
	self.p1 = p1
	return self
}
