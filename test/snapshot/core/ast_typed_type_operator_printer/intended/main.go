package main

import (
	"snapshot/hxrt"
	"sync/atomic"
)

func main() {
	hxrt.Println(hxrt.StringFromLiteral("ok"))
}

type hxrt__test_ast_Local struct {
}

type hxrt__test_ast_type_shapes struct {
	Builtin    bool
	Named      hxrt__test_ast_Local
	Pointer    *string
	Slice      []*hxrt__test_ast_Local
	Array      [3]byte
	Map        map[string][]int
	Channel    chan int
	Receive    <-chan string
	Send       chan<- bool
	Callback   func(int, ...string) (bool, error)
	Constraint interface{ Apply(int) (string, error) }
	Empty      interface{}
	Generic    atomic.Pointer[int]
}

func hxrt__test_ast_typed_operator_smoke(left int, right int, boxed any) bool {
	_ = -left
	_ = ^right
	_ = []int{left, right}
	_ = boxed.(interface{ Apply(int) (string, error) })
	return ((left < right) && (left != right))
}
