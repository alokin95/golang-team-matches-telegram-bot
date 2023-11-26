package client

import (
	"bytes"
	"io"
	"net/http"
)

type HttpClient struct {
	Client        *http.Client
	DefaultHeader map[string]string
	ApiKey        string
}

func NewHttpClient(apiKey string) *HttpClient {
	return &HttpClient{
		Client: &http.Client{},
		DefaultHeader: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + apiKey,
			"OpenAI-Beta":   "assistants=v1",
		},
	}
}

func (client *HttpClient) SendPost(url string, data []byte) (*http.Response, error) {
	req, err := client.createRequest("POST", url, data)
	if err != nil {
		return nil, err
	}

	return client.Client.Do(req)
}

func (client *HttpClient) SendGet(url string) (*http.Response, error) {
	req, err := client.createRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return client.Client.Do(req)
}

func (client *HttpClient) createRequest(method, url string, data []byte) (*http.Request, error) {
	var body io.Reader

	if data != nil {
		body = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range client.DefaultHeader {
		req.Header.Set(key, value)
	}

	return req, nil
}
