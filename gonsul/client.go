package gonsul

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type consulResponse struct {
	Value string `json:"Value"`
}

type apiParams struct {
	method string
	body   []byte
}

func (g Gonsul) callAPI(params apiParams) (httpResponse *http.Response, err error) {
	request, err := http.NewRequest(params.method, fmt.Sprintf("%s/%s", g.Host, g.Path), bytes.NewBuffer(params.body))
	if err != nil {
		err = errUnableToCreateHttpRequest
		return
	}

	request.Header.Set("Content-Type", "application/json")

	httpResponse, err = g.HttpClient.Do(request)
	if err != nil {
		err = errUnableToCallConsulAPI
	}
	return
}

func (g Gonsul) getValues() (resp consulResponse, err error) {
	response, err := g.callAPI(apiParams{method: http.MethodGet, body: nil})
	if err != nil {
		return
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		err = errUnableToReadConsulResponse
		return
	}

	if len(bodyBytes) == 0 {
		err = errConsulKeyNotFound
		return
	}
	var r []consulResponse
	err = json.Unmarshal(bodyBytes, &r)
	if err != nil {
		err = errUnableToReadConsulResponse
		return
	}
	if len(r) > 0 {
		return r[0], nil
	}
	return
}

func (g Gonsul) putValues(value []byte) (err error) {
	response, err := g.callAPI(apiParams{method: http.MethodPut, body: value})
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		err = errConsulUpdateValueFailed
		return
	}
	return
}
