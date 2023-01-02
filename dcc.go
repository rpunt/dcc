package dcc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type HttpClient struct {
	BaseURL    string
	Headers    map[string]string
	Data       map[string]string
	Params     map[string]string
	HTTPClient *http.Client
}

type HttpResponse struct {
	Body    string
	Code    int
	Headers map[string]string
}

func NewClient() *HttpClient {
	return &HttpClient{
		BaseURL: "",
		Headers: make(map[string]string),
		Data:    make(map[string]string),
		Params:  make(map[string]string),
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func sendRequest(client *HttpClient, path string, method string) (HttpResponse, error) {
	// url, err := url.Parse(fmt.Sprintf("%s%s", client.BaseURL, path))
	// if err != nil {
	// 	log.Panicf("Error occurred. %+v", err)
	// }

	var requestData []byte
	if len(client.Data) > 0 {
		requestData, _ = json.Marshal(client.Data)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", client.BaseURL, path), bytes.NewBuffer(requestData))
	if err != nil {
		log.Panicf("Error Occurred. %+v", err)
	}
	for k, v := range client.Headers {
		req.Header.Set(k, v)
	}

	// Query params
	q := req.URL.Query()
	for k, v := range client.Params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	response, err := client.HTTPClient.Do(req)
	if err != nil {
		log.Panicf("Error sending request to API endpoint. %+v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Panicf("Couldn't parse response body. %+v", err)
	}

	resp := HttpResponse{
		Body: string(body),
		Code: response.StatusCode,
	}
	return resp, nil
}

func (client *HttpClient) Get(path string) (HttpResponse, error) {
	return sendRequest(client, path, http.MethodGet)
}

func (client *HttpClient) Post(path string) (HttpResponse, error) {
	return sendRequest(client, path, http.MethodPost)
}
