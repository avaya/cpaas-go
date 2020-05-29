// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
)

// ListSms -
func (c *Client) ListSms(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SMS/Messages", c.Config.AccountSid),
		params,
	)
	return
}

// GetSms -
func (c *Client) GetSms(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SMS/Messages/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// SendSms -
func (c *Client) SendSms(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/SMS/Messages", c.Config.AccountSid),
		params,
	)
	return
}
