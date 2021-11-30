package soar

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"path/filepath"
)

func writeJSON(data []byte, fp string) (string, error) {
	var res *map[string]interface{}
	json.Unmarshal(data, res)
	m, _ := json.Marshal(res)

	if fp != "" {
		abspath, err := filepath.Abs(fp)
		if err != nil {
			return "", err
		}
		if !path.IsAbs(abspath) {
			return "", errors.New("could not resolve path, aborting")
		}
		file, err := os.Open(abspath)
		if err != nil {
			return "", err
		}
		file.Write(m)
	}

	return string(m), nil
}
