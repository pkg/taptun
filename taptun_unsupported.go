// +build !linux

package taptun

import (
	"fmt"
	"runtime"
)

func createInterface(flags uint16) (string, *os.File, error) {
	return "", nil, fmt.Errorf("%s is unsupported", runtime.GOOS)
}
