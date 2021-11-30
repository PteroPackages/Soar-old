package soar

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SoarSession struct {
	URL        string
	Key        string
	Client     *http.Client
	RetryLimit int32 // TODO: implement this
}

func (SoarSession) New(url, key string) *SoarSession {
	return &SoarSession{
		URL:        url,
		Key:        key,
		Client:     &http.Client{},
		RetryLimit: 3,
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
