// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListSms(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List sms retrieval successful", t, func() {
		response, err := client.ListSms(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("sms_messages")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for application := range subResource {
			So(subResource[application]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestGetSms(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Sms retrieval successful", t, func() {
		response, err := client.GetSms("SM24b6d24c57c94c58b3d6f35e0ca9bd71")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "SMO07889084c63ecd8c3c3f49c7b552ba6f")
	})
}

func TestSendSms(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Send mms", t, func() {
		response, err := client.SendSms(map[string]string{
			"From": "+xxx",
			"To":   "+xxx",
			"Body": "body",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "SM07889084c63ecd8c3c3f49c7b552ba6f")
	})
}
