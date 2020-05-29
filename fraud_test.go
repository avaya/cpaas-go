// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListFraud(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List fraud retrieval successful", t, func() {
		response, err := client.ListFraud(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("frauds")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for application := range subResource {
			So(subResource[application]["date_updated"], ShouldNotBeBlank)
		}
	})

}

func TestBlockFraud(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Will disable sms but leave voice enabled (fraud)", t, func() {
		response, err := client.BlockDestination("US", map[string]string{
			"MobileEnabled": "true",
			"SmsEnabled":    "false",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSingleResource("blocked")
		So(subResource["date_updated"], ShouldNotBeBlank)
	})
}

func TestAuthorizeFraud(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Will authorize US", t, func() {
		response, err := client.AuthorizeDestination("US", map[string]string{
			"MobileEnabled": "true",
			"SmsEnabled":    "true",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSingleResource("authorized")
		So(subResource["date_updated"], ShouldNotBeBlank)
	})
}

func TestExtendFraud(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Will authorize US", t, func() {
		response, err := client.ExtendDestination("US")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSingleResource("authorized")
		So(subResource["date_updated"], ShouldNotBeBlank)
	})
}

func TestWhitelistFraud(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Will whitelist US", t, func() {
		response, err := client.WhitelistDestination("US", map[string]string{
			"MobileEnabled": "true",
			"SmsEnabled":    "true",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSingleResource("whitelisted")
		So(subResource["date_updated"], ShouldNotBeBlank)
	})
}
