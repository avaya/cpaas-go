// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
)

// ListConferences -
func (c *Client) ListConferences(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Conferences", c.Config.AccountSid),
		params,
	)
	return
}

// GetConference -
func (c *Client) GetConference(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Conferences/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// ListParticipants -
func (c *Client) ListParticipants(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Conferences/%s/Participants", c.Config.AccountSid, sid),
		params,
	)
	return
}

// GetParticipant -
func (c *Client) GetParticipant(conferenceSid string, sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Conferences/%s/Participants/%s", c.Config.AccountSid, conferenceSid, sid),
		map[string]string{},
	)
	return
}

// MuteOrDeafParticipant -
func (c *Client) MuteOrDeafParticipant(conferenceSid string, sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/Conferences/%s/Participants/%s", c.Config.AccountSid, conferenceSid, sid),
		params,
	)
	return
}

// PlayParticipant -
func (c *Client) PlayParticipant(conferenceSid string, sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/Conferences/%s/Participants/%s/Play", c.Config.AccountSid, conferenceSid, sid),
		params,
	)
	return
}

// HangupParticipant -
func (c *Client) HangupParticipant(conferenceSid string, sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/Conferences/%s/Participants/%s", c.Config.AccountSid, conferenceSid, sid),
		map[string]string{},
	)
	return
}
