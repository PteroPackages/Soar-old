package soar

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func WriteLocalFile(data, ext string) error {
	dir, _ := os.Getwd()
	name := fmt.Sprintf("%s/out_%d.%s", dir, time.Now().Unix(), ext)
	fp, err := filepath.Abs(name)
	if err != nil {
		return err
	}

	var _file *os.File
	if exists(fp) {
		_file, err = os.Open(fp)
	} else {
		_file, err = os.Create(fp)
	}
	if err != nil {
		return err
	}

	_file.WriteString(data)
	_file.Close()
	return nil
}

func WriteLogFile(data string) error {
	file, err := os.Create(fmt.Sprintf("/bin/logs/log_%d.log", time.Now().Unix()))
	if err != nil {
		return err
	}
	file.WriteString(data)
	file.Close()
	return nil
}
