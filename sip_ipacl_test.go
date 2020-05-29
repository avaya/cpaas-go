// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListIpAccessLists(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("ListIpAccessLists retrieval successful", t, func() {
		response, err := client.ListIpAccessLists(map[string]string{
			"Page":     "1",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("ip_access_control")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for transcription := range subResource {
			So(subResource[transcription]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestGetIpAccessList(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("GetIpAccessList retrieval successful", t, func() {
		response, err := client.GetIpAccessList("AL7c889084ddbef1bd6a6a4bceb421d4e3")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "AL7c889084ddbef1bd6a6a4bceb421d4e3")
	})
}

func TestCreateIpAccessList(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("CreateIpAccessList creation successful", t, func() {
		response, err := client.CreateIpAccessList(map[string]string{
			"FriendlyName": "Testing",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "Testing")
	})
}

func TestUpdateIpAccessList(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("UpdateIpAccessList retrieval successful", t, func() {
		response, err := client.UpdateIpAccessList("AL7c889084ddbef1bd6a6a4bceb421d4e3", map[string]string{
			"FriendlyName": "testing-update",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "testing-update")
	})
}

func TestDeleteIpAccessList(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Transcription retrieval successful", t, func() {
		response, err := client.DeleteIpAccessList("AL7c889084ddbef1bd6a6a4bceb421d4e2")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "AL7c889084ddbef1bd6a6a4bceb421d4e3")
	})
}

func TestListSipIpAddresses(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("ListSipIpAddresses retrieval successful", t, func() {
		response, err := client.ListSipIpAddresses("AL7c889084ddbef1bd6a6a4bceb421d4e3", map[string]string{
			"Page":     "1",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("credentials")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for credential := range subResource {
			So(subResource[credential]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestGetSipIpAddress(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("GetSipIpAddress retrieval successful", t, func() {
		response, err := client.GetSipIpAddress("AL7c889084ddbef1bd6a6a4bceb421d4e3", "IP65889084cd3aec4ad39742e98491f3bd")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "SCe18890849a99187d6d454d348783d15c")
	})
}

func TestCreateSipIpAddress(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("CreateSipIpAddress creation successful", t, func() {
		response, err := client.CreateSipIpAddress("AL7c889084ddbef1bd6a6a4bceb421d4e3", map[string]string{
			"FriendlyName": "Testing",
			"IpAddress":    "70.23.122.11",
		})

		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "Testing")
	})
}

func TestUpdateSipIpAddress(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("UpdateSipIpAddress retrieval successful", t, func() {
		response, err := client.UpdateSipIpAddress("AL7c889084ddbef1bd6a6a4bceb421d4e3", "IP65889084cd3aec4ad39742e98491f3bd", map[string]string{
			"FriendlyName": "testing-update",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "testing-update")
	})
}

func TestDeleteSipIpAccess(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("DeleteSipIpAddress retrieval successful", t, func() {
		response, err := client.DeleteSipIpAddress("AL7c889084ddbef1bd6a6a4bceb421d4e3", "IP65889084cd3aec4ad39742e98491f3bc")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "IP65889084cd3aec4ad39742e98491f3bd")
	})
}
