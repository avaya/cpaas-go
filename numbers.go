// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
)

// ListNumbers -
func (c *Client) ListNumbers(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/IncomingPhoneNumbers", c.Config.AccountSid),
		params,
	)
	return
}

// GetNumber -
func (c *Client) GetNumber(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/IncomingPhoneNumbers/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// PurchaseNumber -
func (c *Client) PurchaseNumber(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/IncomingPhoneNumbers", c.Config.AccountSid),
		params,
	)
	return
}

// UpdateNumber -
func (c *Client) UpdateNumber(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/IncomingPhoneNumbers/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// DeleteNumber -
func (c *Client) DeleteNumber(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/IncomingPhoneNumbers/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// ListAvailableNumbers -
func (c *Client) ListAvailableNumbers(country string, numberType string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/AvailablePhoneNumbers/%s/%s", c.Config.AccountSid, country, numberType),
		params,
	)
	return
}
