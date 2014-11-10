package taptun

import (
	"unsafe"
)

const (
	FIODGNAME     = 0x80086678
	SIOCIFDESTROY = 0x80206979
)

type fiodgnameArg struct {
	length int32
	buf    unsafe.Pointer
}
