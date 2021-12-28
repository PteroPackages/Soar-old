package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pteropackages/soar/config"
)

type SoarSession struct {
	Config     *config.Config
	URL        string
	Key        string
	Client     *http.Client
	RetryLimit int32 // TODO: implement this
}

func NewSession(api string) *SoarSession {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Printf("soar: %s", err.Error())
		return nil
	}

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
