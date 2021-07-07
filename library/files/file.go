package files

import (
	"os"
)

// FileExists to check if file by path are exists
func FileExists(path string) bool {
	i, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return !i.IsDir()
}
