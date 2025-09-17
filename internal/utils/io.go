package utils

import (
	"errors"
	"os"
)

func WriteNewFile(filename, content string) error {
	if _, err := os.Stat(filename); err != nil {
		return WriteFile(filename, content)
	}
	return errors.New("file already exists")
}

func WriteFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

