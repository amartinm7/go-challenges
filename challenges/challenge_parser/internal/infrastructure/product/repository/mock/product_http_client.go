package mock

import (
	"bytes"
	"io"
	"net/http"
)

type HttpClientMock struct {
}

func (httpClientMock *HttpClientMock) Get(url string) (resp *http.Response, err error) {
	return &http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader([]byte(JsonProduct))),
	}, nil
}
