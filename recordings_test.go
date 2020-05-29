// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListRecordings(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List recordings retrieval successful", t, func() {
		response, err := client.ListRecordings(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("recordings")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for recording := range subResource {
			So(subResource[recording]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestGetRecording(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Get recording successful", t, func() {
		response, err := client.GetRecording("RE65889084ddbf05d8f49a44d3ac173731")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "RE65889084ddbf05d8f49a44d3ac173731")
	})
}

func TestDeleteRecording(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Recording deletion successful", t, func() {
		response, err := client.DeleteRecording("RE§e1889084349f3883ba094fea8f252572")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "REe1889084349f3883ba094fea8f252572")
	})
}
