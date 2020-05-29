// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
)

// ListSipDomainCredentialMappings -
func (c *Client) ListSipDomainCredentialMappings(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/Domains/%s/CredentialListMappings", c.Config.AccountSid, sid),
		params,
	)
	return
}

// MapSipDomainCredentialList -
func (c *Client) MapSipDomainCredentialList(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/SIP/Domains/%s/CredentialListMappings", c.Config.AccountSid, sid),
		params,
	)
	return
}

// DeleteSipDomainCredentialListMapping -
func (c *Client) DeleteSipDomainCredentialListMapping(domainSid string, sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/SIP/Domains/%s/CredentialListMappings/%s", c.Config.AccountSid, domainSid, sid),
		map[string]string{},
	)
	return
}

// ListSipDomainIpAclMappings -
func (c *Client) ListSipDomainIpAclMappings(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/Domains/%s/IpAccessControlListMappings", c.Config.AccountSid, sid),
		params,
	)
	return
}

// MapSipDomainIpAcl -
func (c *Client) MapSipDomainIpAcl(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/SIP/Domains/%s/IpAccessControlListMappings", c.Config.AccountSid, sid),
		params,
	)
	return
}

// DeleteSipDomainIpAclMapping -
func (c *Client) DeleteSipDomainIpAclMapping(domainSid string, sid string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/SIP/Domains/%s/IpAccessControlListMappings/%s", c.Config.AccountSid, domainSid, sid),
		map[string]string{},
	)
	return
}
