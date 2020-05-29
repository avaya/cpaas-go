// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListSipDomains(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List sip domains retrieval successful", t, func() {
		response, err := client.ListSipDomains(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("domains")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for transcription := range subResource {
			So(subResource[transcription]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestGetSipDomain(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Transcription retrieval successful", t, func() {
		response, err := client.GetSipDomain("SD8f889084962719d55a674fd7b84f4402")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "SD8f889084962719d55a674fd7b84f4402")
	})
}

func TestCreateSipDomain(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Sip domain creation successful", t, func() {
		response, err := client.CreateSipDomain(map[string]string{
			"FriendlyName": "Testing",
			"DomainName":   "testing",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "Testing")
	})
}

func TestUpdateSipDomain(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Transcription retrieval successful", t, func() {
		response, err := client.UpdateSipDomain("SD8f889084962719d55a674fd7b84f4402", map[string]string{
			"FriendlyName": "testing-update",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "testing-update")
	})
}

func TestDeleteSipDomain(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Transcription retrieval successful", t, func() {
		response, err := client.DeleteSipDomain("SD8f889084962719d55a674fd7b84f4402")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "SD8f889084962719d55a674fd7b84f4402")
	})
}
