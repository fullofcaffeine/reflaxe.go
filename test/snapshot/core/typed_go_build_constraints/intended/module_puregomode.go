//go:build !cgo

package main

type PureGoMode struct {
	tag    int
	params []any
}

var PureGoMode_Enabled *PureGoMode = &PureGoMode{tag: 0}
