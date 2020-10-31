package rest

import (
	json2 "encoding/json"
	"net/http"
	"strings"
)

// TODO: Do we have some sort of system to register endpoints as JSON or TOML or through Swagger docs which can then be referenced by name in the scenarios?
type (
	Request struct {
		BaseUrl string
		Method  string
		Path    string
		Body    string
		Headers map[string]string
		// TODO:
		QueryParameters map[string]string
	}
)

func NewRequest(baseUrl string) Request {
	return Request{BaseUrl: baseUrl}
}

func (request Request) WithBaseUrl(baseUrl string) Request {
	request.BaseUrl = baseUrl
	return request
}

func (request Request) WithMethod(method string) Request {
	request.Method = method
	return request
}

// TODO: Append path?
// TODO: Nice way to handle path variables here too? -> Then we can just set those to specific values from examples
func (request Request) WithPath(path string) Request {
	request.Path = path
	return request
}

func (request Request) WithBody(body string) Request {
	request.Body = body
	return request
}

func (request Request) WithJsonBody(body interface{}) (Request, error) {
	json, err := json2.Marshal(body)
	if err != nil {
		return request, err
	}

	request.Body = string(json)
	return request, nil
}

func (request Request) WithHeader(headerName string, headerValue string) Request {
	// TODO: Check this doesn't change the original header
	newRequest := request
	newRequest.Headers[headerName] = headerValue
	return newRequest
}

func (request Request) buildQueryParameters() string {
	return ""
}

func (request Request) buildUrl() string {
	return ""
}

func (request Request) SendRequest() (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(request.Method, request.buildUrl(), strings.NewReader(request.Body))

	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
