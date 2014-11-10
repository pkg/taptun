package taptun

import (
	"unsafe"
)

const (
	FIODGNAME     = 0x80106678
	SIOCIFDESTROY = 0x80206979
)

type fiodgnameArg struct {
	length int32
	_pad   [4]byte
	buf    unsafe.Pointer
}
