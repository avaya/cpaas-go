// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import "fmt"

// ListFraud -
func (c *Client) ListFraud(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Fraud", c.Config.AccountSid),
		params,
	)
	return
}

// BlockDestination -
func (c *Client) BlockDestination(countryCode string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Fraud/Block/%s", c.Config.AccountSid, countryCode),
		params,
	)
	return
}

// AuthorizeDestination -
func (c *Client) AuthorizeDestination(countryCode string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Fraud/Authorize/%s", c.Config.AccountSid, countryCode),
		params,
	)
	return
}

// ExtendDestination -
func (c *Client) ExtendDestination(countryCode string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Fraud/Extend/%s", c.Config.AccountSid, countryCode),
		map[string]string{},
	)
	return
}

// WhitelistDestination -
func (c *Client) WhitelistDestination(countryCode string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Fraud/Whitelist/%s", c.Config.AccountSid, countryCode),
		params,
	)
	return
}
