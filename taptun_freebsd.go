package taptun

import (
	"os"
	"syscall"
	"unsafe"
)

type ifreq struct {
	name [syscall.IFNAMSIZ]byte
	_    [16]byte
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

func destroyInterface(name string) error {
	s, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_IP)
	if err != nil {
		return err
	}
	defer syscall.Close(s)

	ifreq := ifreq{}
	copy(ifreq.name[:], []byte(name))

	return ioctl(uintptr(s), syscall.SIOCIFDESTROY, unsafe.Pointer(&ifreq))
}

func openTun() (string, *os.File, error) {
	return createInterface("/dev/tun")
}

func openTap() (string, *os.File, error) {
	return createInterface("/dev/tap")
}
