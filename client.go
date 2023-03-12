package core

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	baseUrl string
	ctx     context.Context
	Options map[string]string

	client HttpClient
}

func NewClient(baseUrl string, opts map[string]string) *Client {
	return &Client{
		Options: opts,
		ctx:     context.Background(),
		client:  http.DefaultClient,
		baseUrl: baseUrl,
	}
}

func NewMockClient(handler http.HandlerFunc) *Client {
	return &Client{
		ctx:    context.Background(),
		client: NewMockHttpClient(handler),
	}
}

func (c Client) newRequest(method string, url string, body interface{}) (*http.Request, error) {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	url = c.baseUrl + url
	req, err := http.NewRequestWithContext(c.ctx, method, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, err
	}

	for k, v := range c.Options {
		req.Header.Add(k, v)
	}

	return req, nil
}

func (c Client) Request(method string, path string, req, resp interface{}) (*Response, error) {
	httpReq, err := c.newRequest(method, path, req)
	if err != nil {
		return nil, err
	}
	rawQuery, err := buildRawQuery(httpReq, req)
	if err != nil {
		return nil, err
	}
	httpReq.URL.RawQuery = rawQuery

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	response := &Response{}
	response.Data = resp
	response.StatusCode = httpResp.StatusCode
	if httpResp.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, &response)
	} else {
		err = json.Unmarshal(body, &response.Data)
	}
	if err != nil {
		return nil, err
	}

	return response, nil
}

type MockHttpClient struct {
	handler http.HandlerFunc
}

func NewMockHttpClient(handler http.HandlerFunc) *MockHttpClient {
	return &MockHttpClient{
		handler: handler,
	}
}

func (c MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	rr := httptest.NewRecorder()
	c.handler.ServeHTTP(rr, req)

	return rr.Result(), nil
}

func NewMockHttpHandler(statusCode int, json string, headers map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(headers) > 0 {
			for key, value := range headers {
				w.Header().Add(key, value)
			}
		}

		w.WriteHeader(statusCode)
		w.Write([]byte(json))
	}
}
