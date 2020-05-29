// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListNumbers(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List numbers retrieval successful", t, func() {
		response, err := client.ListNumbers(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("phone_numbers")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for notification := range subResource {
			So(subResource[notification]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestGetNumber(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Number retrieval successful", t, func() {
		response, err := client.GetNumber("DI65889084ddbf05d8f49a44d3ac173731")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "DI65889084ddbf05d8f49a44d3ac173731")
	})
}

func TestUpdateNumber(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Number update successful", t, func() {
		response, err := client.UpdateNumber("DI65889084ddbf05d8f49a44d3ac173731", map[string]string{
			"FriendlyName": "Testing",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "Testing")
	})
}

func TestDeleteNumber(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Number deletion successful", t, func() {
		response, err := client.DeleteNumber("DIe1889084349f3883ba094fea8f252572")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "DIe1889084349f3883ba094fea8f252572")
	})
}

func TestPurchaseNumber(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Number purchase successful", t, func() {
		response, err := client.PurchaseNumber(map[string]string{
			"AreaCode":     "908",
			"FriendlyName": "Test purchase",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "Test purchase")
	})
}

func TestListAvailableNumbers(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Number list successful", t, func() {
		response, err := client.ListAvailableNumbers("US", "Local", map[string]string{
			"Page":     "1",
			"PageSize": "1",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("phone_numbers")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for notification := range subResource {
			So(subResource[notification]["date_updated"], ShouldNotBeBlank)
		}
	})
}
