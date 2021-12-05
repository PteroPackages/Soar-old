package request

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

func toPascalCase(key string) string {
	res := strings.Builder{}
	start := false
	nextUpper := false

	for _, char := range strings.Split(key, "") {
		if char == "_" {
			nextUpper = true
			continue
		}

		if nextUpper {
			nextUpper = false
			res.WriteString(strings.ToUpper(char))
			continue
		}

		if start {
			start = false
			res.WriteString(strings.ToUpper(char))
			continue
		}

		res.WriteString(char)
	}

	return res.String()
}

func ParseToJSON(data []byte) (parsed map[string]interface{}) {
	json.Unmarshal(data, &parsed)
	return
}

func ParseToYAML(data []byte) (parsed map[string]interface{}) {
	yaml.Unmarshal(data, &parsed)
	return
}

func ParseToString(data []byte) string {
	parsed := ParseToJSON(data)
	var res *strings.Builder

	for k, v := range parsed {
		res.WriteString(fmt.Sprintf("%s:\t%v\n", toPascalCase(k), v))
	}

	return res.String()
}
