// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
)

// ListApplications -
func (c *Client) ListApplications(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Applications", c.Config.AccountSid),
		params,
	)
	return
}

// GetApplication -
func (c *Client) GetApplication(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Applications/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// CreateApplication -
func (c *Client) CreateApplication(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Applications", c.Config.AccountSid),
		params,
	)
	return
}

// UpdateApplication -
func (c *Client) UpdateApplication(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/Applications/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// DeleteApplication -
func (c *Client) DeleteApplication(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/Applications/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// GetApplicationClient -
func (c *Client) GetApplicationClient(applicationSid string, sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Applications/%s/Clients/%s", c.Config.AccountSid, applicationSid, sid),
		map[string]string{},
	)
	return
}

// ListApplicationClients -
func (c *Client) ListApplicationClients(applicationSid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Applications/%s/Clients", c.Config.AccountSid, applicationSid),
		params,
	)
	return
}
