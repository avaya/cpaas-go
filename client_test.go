// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func newTestClient() (*Client, error) {
	config, err := NewCustomConfig(&Config{
		ApiUrl:              ZangApiUrl,
		ApiVersion:          ZangApiVersion,
		AccountSid:          "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		AuthToken:           "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		ResponseContentType: "json",
	})

	if err != nil {
		return nil, err
	}

	return &Client{Config: config}, nil
}

func TestNewClient(t *testing.T) {
	//Convey("New Client With Defaults (No AccountSid)", t, func() {
	//	client, err := NewClient()
	//	So(client, ShouldBeNil)
	//	So(err.Error(), ShouldContainSubstring, "Make sure to provide valid AccountSid")
	//})

	Convey("New custom client config expects valid api url", t, func() {
		config, err := NewCustomConfig(&Config{})
		So(config, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Make sure to provide valid Api URL.")
	})

	Convey("New custom client config expects valid api version", t, func() {
		config, err := NewCustomConfig(&Config{
			ApiUrl: ZangApiUrl,
		})
		So(config, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Make sure to provide valid api version.")
	})

	Convey("New custom client config expects valid auth token", t, func() {
		config, err := NewCustomConfig(&Config{
			ApiUrl:     ZangApiUrl,
			ApiVersion: ZangApiVersion,
			AccountSid: "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		})
		So(config, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Make sure to provide valid AuthToken.")
	})

	Convey("New custom client config expects response content type", t, func() {
		config, err := NewCustomConfig(&Config{
			ApiUrl:     ZangApiUrl,
			ApiVersion: ZangApiVersion,
			AccountSid: "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			AuthToken:  "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		})
		So(config, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Make sure to provide valid response content type.")
	})

	Convey("New custom client config expects response valid content type", t, func() {
		config, err := NewCustomConfig(&Config{
			ApiUrl:              ZangApiUrl,
			ApiVersion:          ZangApiVersion,
			AccountSid:          "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			AuthToken:           "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			ResponseContentType: "csv",
		})
		So(config, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Make sure to provide valid response content type.")
	})

	Convey("New custom client config should pass", t, func() {
		config, err := NewCustomConfig(&Config{
			ApiUrl:              ZangApiUrl,
			ApiVersion:          ZangApiVersion,
			AccountSid:          "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			AuthToken:           "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			ResponseContentType: "json",
		})
		So(config, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("New Client with proper access credentials validates fine", t, func() {
		os.Setenv("ZANG_CLOUD_ACCOUNT_SID", "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		os.Setenv("ZANG_CLOUD_AUTH_TOKEN", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		defer os.Setenv("ZANG_CLOUD_ACCOUNT_SID", "")
		defer os.Setenv("ZANG_CLOUD_AUTH_TOKEN", "")

		client, err := NewClient()
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}
