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

func destroyInterface(name string) error {
	return nil
}

func openTun() (string, *os.File, error) {
	return createInterface(syscall.IFF_TUN | syscall.IFF_NO_PI)
}

func openTap() (string, *os.File, error) {
	return createInterface(syscall.IFF_TAP | syscall.IFF_NO_PI)
}
