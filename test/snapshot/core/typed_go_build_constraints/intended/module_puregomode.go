//go:build !cgo

package main

type PureGoMode struct {
	tag    int
	params []any
}

func (self *PureGoMode) HxEnumValue() (string, []any) {
	return [1]string{"Enabled"}[self.tag], self.params
}

var PureGoMode_Enabled *PureGoMode = &PureGoMode{tag: 0}
