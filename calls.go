// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
)

// ListCalls -
func (c *Client) ListCalls(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Calls", c.Config.AccountSid),
		params,
	)
	return
}

// GetCall -
func (c *Client) GetCall(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Calls/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// CreateCall -
func (c *Client) CreateCall(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Calls", c.Config.AccountSid),
		params,
	)
	return
}

// InterruptCall -
func (c *Client) InterruptCall(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/Calls/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// SendDigitsCall -
func (c *Client) SendDigitsCall(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/Calls/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// RecordCall -
func (c *Client) RecordCall(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Calls/%s/Recordings", c.Config.AccountSid, sid),
		params,
	)
	return
}

// PlayAudioCall -
func (c *Client) PlayAudioCall(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Calls/%s/Play", c.Config.AccountSid, sid),
		params,
	)
	return
}

// ApplyEffectsCall -
func (c *Client) ApplyEffectsCall(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Calls/%s/Effect", c.Config.AccountSid, sid),
		params,
	)
	return
}
