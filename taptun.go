// Package taptun provides an interface to the user level network
// TAP / TUN device.
//
// https://www.kernel.org/doc/Documentation/networking/tuntap.txt
package taptun

import (
	"bytes"
	"fmt"
	"io"
	"syscall"
	"unsafe"
)

// NewTun creates a *Tun device with the specified name and returns the
// device connected to the tun interface.
//
// If an empty string ("") is specified for name, a tunN interface is
// created.
func NewTun(name string) (*Tun, error) {
	n, f, err := openTun(name)
	return &Tun{
		ReadWriteCloser: f,
		name:            n,
	}, err
}

// OpenTun creates a tunN interface and returns a *Tun device connected to
// the tun interface.
func OpenTun() (*Tun, error) {
	return NewTun("")
}

// Tun represents a TUN Virtual Point-to-Point network device.
type Tun struct {
	io.ReadWriteCloser
	name string
}

func (t *Tun) String() string {
	return t.name
}

func (t *Tun) Close() error {
	if err := t.ReadWriteCloser.Close(); err != nil {
		return err
	}
	return destroyInterface(t.name)
}

// NewTap creates a *Tap device with the specified name and returns the
// device connected to the tap interface.
//
// If an empty string ("") is specified for name, a tapN interface is
// created.
func NewTap(name string) (*Tap, error) {
	n, f, err := openTap(name)
	return &Tap{
		ReadWriteCloser: f,
		name:            n,
	}, err
}

// OpenTap creates a tapN interface and returns a *Tap device connected to
// the tap interface.
func OpenTap() (*Tap, error) {
	return NewTap("")
}

// Tap represents a TAP Virtual Ethernet network device.
type Tap struct {
	io.ReadWriteCloser
	name string
}

func (t *Tap) String() string {
	return t.name
}

func (t *Tap) Close() error {
	if err := t.ReadWriteCloser.Close(); err != nil {
		return err
	}
	return destroyInterface(t.name)
}

// ErrTruncated indicates the buffer supplied to ReadFrame was insufficient
// to hold the ingress frame.
type ErrTruncated struct {
	length int
}

func (e ErrTruncated) Error() string {
	return fmt.Sprintf("supplied buffer was not large enough, frame truncated at %v", e.length)
}

// ReadFrame reads a single frame from the tap device.
// The buffer supplied must be large enough to hold the whole frame including a 4 byte header returned by the kernel.
// If the buffer is not large enough to hold the entire frame and error of type ErrTruncated will be returned.
func ReadFrame(tap *Tap, buf []byte) ([]byte, error) {
	n, err := tap.Read(buf)
	return buf[:n], err
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
