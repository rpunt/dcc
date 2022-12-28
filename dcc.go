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
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func sendRequest(client *HttpClient, path string, method string) (HttpResponse, error) {
	requestData, _ := json.Marshal(client.Data)

	req, err := http.NewRequest(method, client.BaseURL, bytes.NewBuffer(requestData))
	if err != nil {
		log.Panicf("Error Occurred. %+v", err)
	}
	for k, v := range client.Headers {
		req.Header.Set(k, v)
	}

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
	response, err := sendRequest(client, path, http.MethodGet)
	if err != nil {
		fmt.Println("error:", err)
	}

	return response, err
}

func (client *HttpClient) Post(path string) (HttpResponse, error) {
	response, err := sendRequest(client, path, http.MethodPost)
	if err != nil {
		fmt.Println("error:", err)
	}

	return response, err
}
