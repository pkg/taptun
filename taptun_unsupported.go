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

func openTun(_ string) (string, *os.File, error) {
	return createInterface(0)
}

func openTap(_ string) (string, *os.File, error) {
	return createInterface(0)
}
