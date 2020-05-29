// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
)

// ListSipDomains -
func (c *Client) ListSipDomains(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/Domains", c.Config.AccountSid),
		params,
	)
	return

}

// GetSipDomain -
func (c *Client) GetSipDomain(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/Domains/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// CreateSipDomain -
func (c *Client) CreateSipDomain(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/SIP/Domains", c.Config.AccountSid),
		params,
	)
	return
}

// UpdateSipDomain -
func (c *Client) UpdateSipDomain(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/SIP/Domains/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// DeleteSipDomain -
func (c *Client) DeleteSipDomain(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/SIP/Domains/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}
