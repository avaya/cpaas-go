// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Response - generic response interface. one day we might wanna change this to static
type Response struct {
	*http.Response
	Resource map[string]interface{}
}

// GetSubResource - Get subresource (useful if you're listing for exmple, accounts)
func (r *Response) GetSubResource(jsonKey string) []map[string]interface{} {
	if r.Resource[jsonKey] == nil {
		return nil
	}

	resource := r.Resource[jsonKey].([]interface{})

	response := []map[string]interface{}{}

	for index := range resource {
		response = append(response, resource[index].(map[string]interface{}))
	}

	return response
}

// GetSingleResource - Get single subresource (useful if you're listing for exmple, accounts)
func (r *Response) GetSingleResource(jsonKey string) map[string]interface{} {
	if r.Resource[jsonKey] == nil {
		return nil
	}

	return r.Resource[jsonKey].(map[string]interface{})
}

// Request - helper designed to handle generic requests towards api interface
type Request struct {
	*Config
}

// GetApiUrl - Will return prepared full api url
func (r *Request) GetApiUrl(uri string) string {
	if uri != "" {
		uri = "/" + uri
	}

	return fmt.Sprintf(
		"%s/%s/Accounts%s.%s",
		r.Config.ApiUrl, r.Config.ApiVersion, uri,
		r.Config.ResponseContentType,
	)
}

// Get - Will initiate GET request towards api
func (r *Request) Get(uri string, params map[string]string) (resp *Response, err error) {
	fmt.Printf("[get] received params: %+v", params)
	if resp, err = r.request("GET", r.GetApiUrl(uri), params); err != nil {
		return nil, err
	}
	return
}

// Update - Will initiate UPDATE request towards api
func (r *Request) Update(uri string, params map[string]string) (resp *Response, err error) {
	if resp, err = r.request("POST", r.GetApiUrl(uri), params); err != nil {
		return nil, err
	}
	return
}

// Create - Will initiate UPDATE request towards api
func (r *Request) Create(uri string, params map[string]string) (resp *Response, err error) {
	if resp, err = r.request("POST", r.GetApiUrl(uri), params); err != nil {
		return nil, err
	}
	return
}

// Delete - Will initiate DELETE request towards api
func (r *Request) Delete(uri string, params map[string]string) (resp *Response, err error) {
	if resp, err = r.request("DELETE", r.GetApiUrl(uri), params); err != nil {
		return nil, err
	}
	return
}

// request - internal method designed to
func (r *Request) request(method string, url string, params map[string]string) (*Response, error) {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	var buffer *bytes.Buffer

	if method == "POST" || method == "PUT" {
		buffer = bytes.NewBufferString(MapToUrlValues(params).Encode())
	} else {
		if len(params) > 0 {
			url = fmt.Sprintf("%s?%s", url, MapToUrlValues(params).Encode())
		}
		buffer = bytes.NewBufferString("")
	}

	fmt.Printf("[request] encoded params: %+v \n", MapToUrlValues(params).Encode())
	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, err
	}

	fmt.Printf("[request] full request url: %s \n", req.URL.String())

	req.SetBasicAuth(r.Config.AccountSid, r.Config.AuthToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Response body bytes are: %s", bodyBytes)

	var jsonData map[string]interface{}

	if err := json.Unmarshal(bodyBytes, &jsonData); err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("An unexpected response returned from Zang API (Status Code: %d): %s", resp.StatusCode, bodyBytes)
	}

	return &Response{
		Response: resp,
		Resource: jsonData,
	}, nil
}
