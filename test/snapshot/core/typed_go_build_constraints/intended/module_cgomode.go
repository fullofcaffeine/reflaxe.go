//go:build cgo

package main

type CgoMode struct {
	tag    int
	params []any
}

var CgoMode_Enabled *CgoMode = &CgoMode{tag: 0}
