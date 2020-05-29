// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

// Config - Configuration struct, designed to be
type Config struct {
	ApiUrl              string
	ApiVersion          string
	AccountSid          string
	AuthToken           string
	ResponseContentType string
}

// OptionString - Will return ENV value back based on provided name casted as string
func OptionString(name, def string) string {
	if res := os.Getenv(name); res != "" {
		return string(res)
	}

	return def
}

// Validate - Will validate client configuration and return error if any issues are found
func (c *Config) Validate() error {
	if _, err := url.ParseRequestURI(c.ApiUrl); err != nil {
		return fmt.Errorf("Make sure to provide valid Api URL. You've provided: %s. Error: %s", c.ApiUrl, err)
	}

	if c.ApiVersion != "v2" {
		return fmt.Errorf("Make sure to provide valid api version. You've provided: %s.", c.ApiVersion)
	}

	if len(c.AccountSid) != 34 || !strings.HasPrefix(c.AccountSid, "AC") {
		return fmt.Errorf("Make sure to provide valid AccountSid. You've provided: %s", c.AccountSid)
	}

	if len(c.AuthToken) != 32 {
		return fmt.Errorf("Make sure to provide valid AuthToken. You've provided: %s", c.AuthToken)
	}

	if !StringInSlice(c.ResponseContentType, AvailableResponseContentTypes) {
		return fmt.Errorf("Make sure to provide valid response content type. You've provided: %s. Available are: %v", c.AuthToken, AvailableResponseContentTypes)
	}

	return nil
}

// SetCredentials - Will set account sid and auth token to desired account
func (c *Config) SetCredentials(accountSid, authToken string) error {
	if len(accountSid) != 34 || !strings.HasPrefix(accountSid, "AC") {
		return fmt.Errorf("Make sure to provide valid AccountSid. You've provided: %s", accountSid)
	}

	if len(authToken) != 32 {
		return fmt.Errorf("Make sure to provide valid AuthToken. You've provided: %s", authToken)
	}

	c.AccountSid = accountSid
	c.AuthToken = authToken
	return nil
}

// NewConfig -
func NewConfig() (*Config, error) {
	cfg := &Config{
		ApiUrl:              OptionString("ZANG_CLOUD_API_URL", ZangApiUrl),
		ApiVersion:          OptionString("ZANG_CLOUD_API_VERSION", ZangApiVersion),
		AccountSid:          OptionString("ZANG_CLOUD_ACCOUNT_SID", ""),
		AuthToken:           OptionString("ZANG_CLOUD_AUTH_TOKEN", ""),
		ResponseContentType: OptionString("ZANG_CLOUD_RESPONSE_TYPE", DefaultResponseContentType),
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// NewCustomConfig -
func NewCustomConfig(cfg *Config) (*Config, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}
