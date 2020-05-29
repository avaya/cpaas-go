// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

// GetAccount -
func (c *Client) GetAccount() (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(c.Config.AccountSid, nil)
	return
}

// GetAccounts -
func (c *Client) GetAccounts(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get("", params)
	return
}

// UpdateAccount -
func (c *Client) UpdateAccount(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Update(c.Config.AccountSid, params)
	return
}
