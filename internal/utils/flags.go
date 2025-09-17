package utils

import (
	"errors"
	"os"
	"time"
)

func ParseArgs() (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New("usage error")
	}
	if os.Args[1] != "new" {
		return "", errors.New("unimplemented error")
	}
	if len(os.Args) > 2 {
		return os.Args[2], nil
	}
	return time.Now().Format("02-01-2006.txt"), nil
}
