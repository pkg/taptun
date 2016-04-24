package taptun

import (
	"errors"
	"os"
	"syscall"
	"unsafe"
	"fmt"
)

// #cgo CFLAGS: -I./
// #include <taptun_darwin.h>
import "C"

type ifreq struct {
	name  [syscall.IFNAMSIZ]byte // c string
	flags uint16                 // c short
	_pad  [24 - unsafe.Sizeof(uint16(0))]byte
}

func createInterface(name string) (string, *os.File, error) {
	var fd, unit C.int
	var error *C.char
	C.osxtun_open(&fd, &unit, &error)
	if fd < 0 {
		return "", nil, errors.New(C.GoString(error))
	}
	tunName := fmt.Sprintf("utun%d", unit)
	return tunName, os.NewFile(uintptr(fd), tunName), nil
}

func destroyInterface(name string) error {
	return nil
}

func openTun(name string) (string, *os.File, error) {
	return createInterface(name)
}

func openTap(name string) (string, *os.File, error) {
	// not support yet
	return "", nil, errors.New("tap not support yet.")
}
