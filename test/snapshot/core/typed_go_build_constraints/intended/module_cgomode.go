//go:build cgo

package main

type CgoMode struct {
	tag    int
	params []any
}

func (self *CgoMode) HxEnumValue() (string, []any) {
	return [1]string{"Enabled"}[self.tag], self.params
}

var CgoMode_Enabled *CgoMode = &CgoMode{tag: 0}
