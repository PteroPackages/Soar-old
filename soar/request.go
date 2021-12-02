package soar

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/pteropackages/soar/cmd"
	"gopkg.in/yaml.v2"
)

type SoarSession struct {
	Config     *cmd.SoarConfig
	URL        string
	Key        string
	Client     *http.Client
	RetryLimit int32 // TODO: implement this
}

func GetConfig() *cmd.SoarConfig {
	file, err := os.ReadFile("/bin/config.yml")
	if err != nil {
		panic(err)
	}

	var config *cmd.SoarConfig
	yaml.Unmarshal(file, &config)
	return config
}

func NewSession(api string) *SoarSession {
	config := GetConfig()

	var auth []string
	switch api {
	case "application":
		auth = []string{config.Application.URL, config.Application.Key}
	case "client":
		auth = []string{config.Client.URL, config.Client.Key}
	}

	return &SoarSession{
		Config:     config,
		URL:        auth[0],
		Key:        auth[1],
		RetryLimit: 1,
	}
}

func (ctx *SoarSession) GetHeaders() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": "Bearer " + ctx.Key,
	}
}

func (ctx *SoarSession) Request(path, method string, payload interface{}) ([]byte, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, ctx.URL+path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	for k, v := range ctx.GetHeaders() {
		req.Header.Set(k, v)
	}

	res, err := ctx.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var data []byte
	res.Body.Read(data)

	return data, nil
}
