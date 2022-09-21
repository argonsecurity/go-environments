package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testUrl  = "https://www.google.com"
	testData = "data"
)

var (
	testHeaders Headers = Headers{"test": "test"}
	testParams  Params  = Params{"test": "test"}
)

type ClientMock struct{}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func GetFakeHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &ClientMock{},
	}
}

func testRequestWithData(t *testing.T, requestFunc func(url string, headers Headers, data interface{}) ([]byte, error)) {
	_, err := requestFunc(testUrl, testHeaders, testData)
	assert.NoError(t, err, "with headers and data")

	_, err = requestFunc(testUrl, nil, nil)
	assert.NoError(t, err, "without headers and data")

	_, err = requestFunc(testUrl, testHeaders, nil)
	assert.NoError(t, err, "without data")

	_, err = requestFunc(testUrl, nil, testData)
	assert.NoError(t, err, "without headers")
}

func Test_Get(t *testing.T) {
	client := GetFakeHTTPClient()
	_, err := client.Get(testUrl, nil, nil)
	assert.NoError(t, err)

	_, err = client.Get(testUrl, testHeaders, nil)
	assert.NoError(t, err)

	_, err = client.Get(testUrl, nil, testParams)
	assert.NoError(t, err)

	_, err = client.Get(testUrl, testHeaders, testParams)
	assert.NoError(t, err)
}

func Test_Post(t *testing.T) {
	client := GetFakeHTTPClient()
	testRequestWithData(t, client.Post)
}

func Test_Delete(t *testing.T) {
	client := GetFakeHTTPClient()
	testRequestWithData(t, client.Delete)
}

func Test_Put(t *testing.T) {
	client := GetFakeHTTPClient()
	testRequestWithData(t, client.Put)
}
