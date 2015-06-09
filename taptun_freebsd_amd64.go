package taptun

import (
	"unsafe"
)

const (
	FIODGNAME     = 0x80106678
)

type fiodgnameArg struct {
	length int32
	_pad   [4]byte
	buf    unsafe.Pointer
}
