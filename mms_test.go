// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListMms(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List mms retrieval successful", t, func() {
		response, err := client.ListMms(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("mms_messages")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for application := range subResource {
			So(subResource[application]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestGetMms(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Conference retrieval successful", t, func() {
		response, err := client.GetMms("MM24b6d24c57c94c58b3d6f35e0ca9bd71")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		res := response.GetSingleResource("mms_row")
		So(res["mms_sid"], ShouldEqual, "MM24b6d24c57c94c58b3d6f35e0ca9bd71")
	})
}

func TestSendMms(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Send mms", t, func() {
		response, err := client.SendMms(map[string]string{
			"From":     "+xxx",
			"To":       "+xxx",
			"Body":     "body",
			"MediaUrl": "http://full-url-to-media.jpeg",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		res := response.GetSingleResource("mms_row")
		So(res["mms_sid"], ShouldNotBeEmpty)
	})
}
