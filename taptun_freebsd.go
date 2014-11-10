package taptun

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	FIODGNAME = 0x80106678
)

type ifreq struct {
	name [syscall.IFNAMSIZ]byte
	_    [16]byte
}

type fiodgnameArg struct {
	length int32
	_pad   [4]byte
	buf    unsafe.Pointer
}

func interfaceName(fd uintptr) (string, error) {
	var name [syscall.IFNAMSIZ]byte

	arg := fiodgnameArg{length: syscall.IFNAMSIZ, buf: unsafe.Pointer(&name)}
	if err := ioctl(fd, FIODGNAME, unsafe.Pointer(&arg)); err != nil {
		return "", err
	}
	return cstringToGoString(name[:]), nil
}

func createInterface(clonefile string) (string, *os.File, error) {
	f, err := os.OpenFile(clonefile, os.O_RDWR, 0600)
	if err != nil {
		return "", nil, err
	}

	fd := f.Fd()
	ifname, err := interfaceName(fd)
	if err != nil {
		f.Close()
		return "", nil, err
	}

	return ifname, f, nil
}

func openTun() (string, *os.File, error) {
	return createInterface("/dev/tun")
}

func openTap() (string, *os.File, error) {
	return createInterface("/dev/tap")
}
