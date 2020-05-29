// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListApplications(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List application retrieval successful", t, func() {
		response, err := client.ListApplications(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("applications")

		So(len(subResource), ShouldBeGreaterThan, 0)

		for application := range subResource {
			So(subResource[application]["date_updated"], ShouldNotBeBlank)
		}
	})

}

func TestGetApplication(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Account retrieval successful", t, func() {
		response, err := client.GetApplication("AP42889084abfad04ab67c45df8b142f12")
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		So(response.Resource["sid"], ShouldEqual, "AP42889084abfad04ab67c45df8b142f12")
	})
}

func TestCreateApplication(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Account retrieval successful", t, func() {
		response, err := client.CreateApplication(map[string]string{
			"FriendlyName": "Application From Golang Client",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		So(response.Resource["friendly_name"], ShouldEqual, "Application From Golang Client")
	})
}

func TestUpdateApplication(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Account retrieval successful", t, func() {
		response, err := client.UpdateApplication("AP07889084badaa62fca894604b652cbca", map[string]string{
			"FriendlyName": "Application From Golang Client - Updated",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		So(response.Resource["friendly_name"], ShouldEqual, "Application From Golang Client - Updated")
		So(response.Resource["sid"], ShouldEqual, "AP07889084badaa62fca894604b652cbca")
	})
}

func TestDeleteApplication(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Account retrieval successful", t, func() {
		response, err := client.DeleteApplication("AP42889084c052fce31b104af4a955e1ce")
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		So(response.Resource["sid"], ShouldEqual, "AP42889084c052fce31b104af4a955e1ce")
	})
}

func TestListApplicationClients(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Account retrieval successful", t, func() {
		response, err := client.ListApplicationClients("AP42889084c052fce31b104af4a955e1ce", map[string]string{})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("clients")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for apclient := range subResource {
			So(subResource[apclient]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestGetApplicationClient(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Account retrieval successful", t, func() {
		response, err := client.GetApplicationClient("AP42889084c052fce31b104af4a955e1ce", "XXXXXX")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		So(response.Resource["sid"], ShouldEqual, "XXXXXX")
	})
}
