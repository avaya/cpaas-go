// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import "fmt"

// ListCarrierLookups -
func (c *Client) ListCarrierLookups(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Lookups/Carrier", c.Config.AccountSid),
		params,
	)
	return
}

// CarrierLookup -
func (c *Client) CarrierLookup(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Lookups/Carrier", c.Config.AccountSid),
		params,
	)
	return
}

// ListCnamLookups -
func (c *Client) ListCnamLookups(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Lookups/Cnam", c.Config.AccountSid),
		params,
	)
	return
}

// CnamLookup -
func (c *Client) CnamLookup(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Lookups/Cnam", c.Config.AccountSid),
		params,
	)
	return
}

// ListBnaLookups -
func (c *Client) ListBnaLookups(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Lookups/Bna", c.Config.AccountSid),
		params,
	)
	return
}

// BnaLookup -
func (c *Client) BnaLookup(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Lookups/Bna", c.Config.AccountSid),
		params,
	)
	return
}
