package lynx

import (
	"os"
)

// Makes a directory if none exists
func mkdirIfNone(dirname string) error {
	err := os.MkdirAll(dirname, os.ModePerm)
	if err == os.ErrInvalid || err == os.ErrPermission {
		return err
	}
	return nil
}