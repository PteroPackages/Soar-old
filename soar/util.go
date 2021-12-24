package soar

import "os"

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func OpenOrCreate(path string) (*os.File, error) {
	if Exists(path) {
		return os.Open(path)
	} else {
		return os.Create(path)
	}
}
