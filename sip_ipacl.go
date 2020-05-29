// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
)

// ListIpAccessLists -
func (c *Client) ListIpAccessLists(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/IpAccessControlLists", c.Config.AccountSid),
		params,
	)
	return
}

// GetIpAccessList -
func (c *Client) GetIpAccessList(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/IpAccessControlLists/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// CreateIpAccessList -
func (c *Client) CreateIpAccessList(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/SIP/IpAccessControlLists", c.Config.AccountSid),
		params,
	)
	return
}

// UpdateIpAccessList -
func (c *Client) UpdateIpAccessList(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/SIP/IpAccessControlLists/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// DeleteIpAccessList -
func (c *Client) DeleteIpAccessList(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/SIP/IpAccessControlLists/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// ListSipIpAddresses -
func (c *Client) ListSipIpAddresses(ipAccessListSid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/IpAccessControlLists/%s/IpAddresses", c.Config.AccountSid, ipAccessListSid),
		params,
	)
	return
}

// GetSipIpAddress -
func (c *Client) GetSipIpAddress(ipAccessListSid string, sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/IpAccessControlLists/%s/IpAddresses/%s", c.Config.AccountSid, ipAccessListSid, sid),
		map[string]string{},
	)
	return
}

// CreateSipIpAddress -
func (c *Client) CreateSipIpAddress(ipAccessListSid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/SIP/IpAccessControlLists/%s/IpAddresses", c.Config.AccountSid, ipAccessListSid),
		params,
	)
	return
}

// UpdateSipIpAddresses -
func (c *Client) UpdateSipIpAddress(ipAccessListSid string, sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/SIP/IpAccessControlLists/%s/IpAddresses/%s", c.Config.AccountSid, ipAccessListSid, sid),
		params,
	)
	return
}

// DeleteSipIpAddress -
func (c *Client) DeleteSipIpAddress(ipAccessListSid string, sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/SIP/IpAccessControlLists/%s/IpAddresses/%s", c.Config.AccountSid, ipAccessListSid, sid),
		map[string]string{},
	)
	return
}
