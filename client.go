// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

// Client -
type Client struct {
	*Config
	*Request
}

// NewClient -
func NewClient() (*Client, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

	return &Client{
		Config:  config,
		Request: &Request{Config: config},
	}, nil
}
