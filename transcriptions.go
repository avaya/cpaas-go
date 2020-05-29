// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
)

// ListTranscriptions -
func (c *Client) ListTranscriptions(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Transcriptions", c.Config.AccountSid),
		params,
	)
	return
}

// GetTranscription -
func (c *Client) GetTranscription(sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Transcriptions/%s", c.Config.AccountSid, sid),
		map[string]string{},
	)
	return
}

// TranscribeRecording -
func (c *Client) TranscribeRecording(recordingSid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(recordingSid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Recordings/%s/Transcriptions", c.Config.AccountSid, recordingSid),
		params,
	)
	return
}

// ListTranscriptions -
func (c *Client) TranscribeAudioUrl(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/Transcriptions", c.Config.AccountSid),
		params,
	)
	return
}
