// +build !linux,!freebsd

package taptun

import (
	"fmt"
	"os"
	"runtime"
)

func createInterface(flags uint16) (string, *os.File, error) {
	return "", nil, fmt.Errorf("%s is unsupported", runtime.GOOS)
}

func destroyInterface(name string) error {
	return fmt.Errorf("%s is unsupported", runtime.GOOS)
}

func openTun() (string, *os.File, error) {
	return createInterface(0)
}

func openTap() (string, *os.File, error) {
	return createInterface(0)
}
