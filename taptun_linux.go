package taptun

import (
	"errors"
	"os"
	"syscall"
	"unsafe"
)

type ifreq struct {
	name  [syscall.IFNAMSIZ]byte // c string
	flags uint16                 // c short
	_pad  [24 - unsafe.Sizeof(uint16(0))]byte
}

func createInterface(flags uint16, name string) (string, *os.File, error) {
	// Last byte of name must be nil for C string, so name must be
	// short enough to allow that
	if len(name) > syscall.IFNAMSIZ-1 {
		return "", nil, errors.New("device name too long")
	}

	f, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0600)
	if err != nil {
		return "", nil, err
	}

	var nbuf [syscall.IFNAMSIZ]byte
	copy(nbuf[:], []byte(name))

	fd := f.Fd()

	ifr := ifreq{
		name:  nbuf,
		flags: flags,
	}
	if err := ioctl(fd, syscall.TUNSETIFF, unsafe.Pointer(&ifr)); err != nil {
		return "", nil, err
	}
	return cstringToGoString(ifr.name[:]), f, nil
}

func destroyInterface(name string) error {
	return nil
}

func openTun(name string) (string, *os.File, error) {
	return createInterface(syscall.IFF_TUN|syscall.IFF_NO_PI, name)
}

func openTap(name string) (string, *os.File, error) {
	return createInterface(syscall.IFF_TAP|syscall.IFF_NO_PI, name)
}
