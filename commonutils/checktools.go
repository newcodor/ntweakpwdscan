package commonutils

import (
	"os"
)

func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil || os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
