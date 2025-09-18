package utils

import (
	"errors"
	"os"
	"time"
)

var ErrUsage = errors.New("usage error")
var ErrNotImplemented = errors.New("unimplemented error")

func ParseArgs() (string, error) {
	if len(os.Args) == 1 {
		return "", ErrUsage
	}
	if os.Args[1] != "new" {
		return "", ErrNotImplemented
	}
	if len(os.Args) > 2 {
		return os.Args[2], nil
	}
	return time.Now().Format("02-01-2006.txt"), nil
}
