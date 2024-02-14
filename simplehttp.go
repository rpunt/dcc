package simplehttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	BaseURL    string
	// Headers    map[string][]string
	Headers    map[string]string
	Data       map[string]string
	Params     map[string]string
	HTTPClient *http.Client
}

type HttpResponse struct {
	Body    string
	Code    int
	Headers map[string][]string
}

func New(baseURL string) *HttpClient {
	return &HttpClient{
		BaseURL: baseURL,
		// Headers: make(map[string][]string),
		Headers: make(map[string]string),
		Data:    make(map[string]string),
		Params:  make(map[string]string),
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func sendRequest(client *HttpClient, path string, method string) (HttpResponse, error) {
	// create the request body, as appropriate
	var requestData []byte
	if len(client.Data) > 0 {
		var err error
		requestData, err = json.Marshal(client.Data)
		if err != nil {
			return HttpResponse{}, err
		}
	}

	// construct the request
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", client.BaseURL, path), bytes.NewBuffer(requestData))
	if err != nil {
		return HttpResponse{}, err
	}
	for k, v := range client.Headers {
		// req.Header.Set(k, v[0])
		req.Header.Set(k, v)
	}

	// add query params, if any
	q := req.URL.Query()
	for k, v := range client.Params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// do :allthethings:
	response, err := client.HTTPClient.Do(req)
	if err != nil {
		return HttpResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return HttpResponse{}, err
	}

	response_headers := make(map[string][]string)
	for k, v := range response.Header {
		response_headers[k] = v
	}

	resp := HttpResponse{
		Body:    string(body),
		Code:    response.StatusCode,
		Headers: response_headers,
	}
	return resp, nil
}

func (client *HttpClient) Get(path string) (HttpResponse, error) {
	return sendRequest(client, path, http.MethodGet)
}

func (client *HttpClient) Post(path string) (HttpResponse, error) {
	return sendRequest(client, path, http.MethodPost)
}

func (client *HttpClient) Patch(path string) (HttpResponse, error) {
	return sendRequest(client, path, http.MethodPatch)
}

func (client *HttpClient) Put(path string) (HttpResponse, error) {
	return sendRequest(client, path, http.MethodPut)
}

func (client *HttpClient) Delete(path string) (HttpResponse, error) {
	return sendRequest(client, path, http.MethodDelete)
}

func (client *HttpClient) Head(path string) (HttpResponse, error) {
	return sendRequest(client, path, http.MethodHead)
}
