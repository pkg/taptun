package taptun

import (
	"bytes"
	"os"
	"syscall"
	"unsafe"
)

type ifreq struct {
	name  [syscall.IFNAMSIZ]byte // c string
	flags uint16                 // c short
	_pad  [24 - unsafe.Sizeof(uint16(0))]byte
}

func createInterface(flags uint16) (string, *os.File, error) {
	f, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0600)
	if err != nil {
		return "", nil, err
	}

	fd := f.Fd()

	ifr := ifreq{flags: flags}
	if err := ioctl(fd, syscall.TUNSETIFF, unsafe.Pointer(&ifr)); err != nil {
		return "", nil, err
	}
	return cstringToGoString(ifr.name[:]), f, nil
}

func ioctl(fd, request uintptr, argp unsafe.Pointer) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, fd, request, uintptr(argp), 0, 0, 0); e != 0 {
		return e
	}
	return nil
}

func cstringToGoString(cstring []byte) string {
	strs := bytes.Split(cstring, []byte{0x00})
	return string(strs[0])
}
