package http

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/exp/maps"
)

type HTTPClient struct {
	client  GoHTTPClient
	headers Headers
}

func NewHTTPClient(headers Headers) *HTTPClient {
	return &HTTPClient{
		client:  http.DefaultClient,
		headers: headers,
	}
}

func (httpClient *HTTPClient) createHttpRequest(method string, url string, headers Headers, data interface{}) (*http.Request, error) {
	var err error

	allHeaders := (map[string]string)(httpClient.headers)
	if allHeaders == nil {
		allHeaders = make(map[string]string)
	}
	maps.Copy(allHeaders, headers)

	buf, contentType, err := parseData(data, allHeaders)
	if err != nil {
		return nil, err
	}

	if contentType != "" {
		allHeaders[ContentTypeHeader] = contentType
	}

	var request *http.Request
	if buf == nil { // Bug in net/http when passing an empty buffer
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, buf)
	}

	if err != nil {
		return nil, err
	}

	for k, v := range allHeaders {
		request.Header.Add(k, v)
	}

	return request, nil
}

func (httpClient *HTTPClient) sendRequest(req *http.Request) ([]byte, error) {

	resp, err := httpClient.client.Do(req)
	if err != nil {
		return nil, err
	}

	var body []byte
	if resp.Body != nil {
		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("got a response with status code %d, message: %s", resp.StatusCode, string(body))
	}
	return body, nil
}

func (httpClient *HTTPClient) Get(url string, headers Headers, params Params) ([]byte, error) {
	req, err := httpClient.createHttpRequest(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}

	// Append query params
	query := req.URL.Query()
	for k, v := range params {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()

	return httpClient.sendRequest(req)
}

func (httpClient *HTTPClient) Post(url string, headers Headers, data interface{}) ([]byte, error) {
	req, err := httpClient.createHttpRequest(http.MethodPost, url, headers, data)
	if err != nil {
		return nil, err
	}

	return httpClient.sendRequest(req)
}

func (httpClient *HTTPClient) Put(url string, headers Headers, data interface{}) ([]byte, error) {
	req, err := httpClient.createHttpRequest(http.MethodPut, url, headers, data)
	if err != nil {
		return nil, err
	}

	return httpClient.sendRequest(req)
}

func (httpClient *HTTPClient) Delete(url string, headers Headers, data interface{}) ([]byte, error) {
	req, err := httpClient.createHttpRequest(http.MethodDelete, url, headers, data)
	if err != nil {
		return nil, err
	}

	return httpClient.sendRequest(req)
}
