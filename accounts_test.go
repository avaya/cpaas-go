// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAccount(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)

		Convey("Accounts retrieval successful", func() {
			response, err := client.GetAccounts(nil)
			So(response, ShouldNotBeNil)
			So(err, ShouldBeNil)

			So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
			So(response.Resource, ShouldNotBeNil)

			subResource := response.GetSubResource("accounts")

			So(len(subResource), ShouldBeGreaterThan, 0)

			for account := range subResource {
				So(subResource[account]["sid"], ShouldEqual, os.Getenv("ZANG_CLOUD_ACCOUNT_SID"))
			}
		})

		Convey("Account retrieval successful", func() {
			response, err := client.GetAccount()
			So(response, ShouldNotBeNil)
			So(err, ShouldBeNil)

			So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
			So(response.Resource, ShouldNotBeNil)

			So(response.Resource["sid"], ShouldEqual, os.Getenv("ZANG_CLOUD_ACCOUNT_SID"))
		})
	})

}

func TestUpdateAccount(t *testing.T) {
	Convey("New Client With Defaults", t, func() {
		client, err := NewClient()

		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)

		Convey("Account retrieval successful", func() {
			response, err := client.UpdateAccount(map[string]string{
				"FriendlyName": "Nevio testing helper",
			})
			So(response, ShouldNotBeNil)
			So(err, ShouldBeNil)

			So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
			So(response.Resource, ShouldNotBeNil)

			So(response.Resource["sid"], ShouldEqual, os.Getenv("ZANG_CLOUD_ACCOUNT_SID"))
			So(response.Resource["friendly_name"], ShouldEqual, "Nevio testing helper")
		})

	})
}
