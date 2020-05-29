// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
)

// ListMms -
func (c *Client) ListMms(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/MMS/Messages", c.Config.AccountSid),
		params,
	)
	return
}

// GetMms -
func (c *Client) GetMms(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/MMS/Messages/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// SendMms -
func (c *Client) SendMms(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/MMS/Messages", c.Config.AccountSid),
		params,
	)
	return
}
