// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListCarrierLookups(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List carrier lookups retrieval successful", t, func() {
		response, err := client.ListCarrierLookups(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("carrier_lookups")

		So(len(subResource), ShouldBeGreaterThan, 0)

		for application := range subResource {
			So(subResource[application]["date_updated"], ShouldNotBeBlank)
		}
	})

}

func TestCarrierLookup(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Will initiate carrier lookup and expect successful response", t, func() {
		response, err := client.CarrierLookup(map[string]string{
			"PhoneNumber": "+XXXXXX",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("carrier_lookups")
		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)
	})
}

func TestListCnamLookups(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List carrier lookups retrieval successful", t, func() {
		response, err := client.ListCnamLookups(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("cnam_dips")

		So(len(subResource), ShouldBeGreaterThan, 0)

		for application := range subResource {
			So(subResource[application]["date_updated"], ShouldNotBeBlank)
		}
	})

}

func TestBnaLookup(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Will initiate bna lookup and expect successful response", t, func() {
		response, err := client.CarrierLookup(map[string]string{
			"PhoneNumber": "+1XXX", // ONLY +1 NUMBERS
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("bna_lookups")
		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)
	})
}

func TestListBnaLookups(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List bna lookups retrieval successful", t, func() {
		response, err := client.ListBnaLookups(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("bna_lookups")

		So(len(subResource), ShouldBeGreaterThan, 0)

		for application := range subResource {
			So(subResource[application]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestCnamLookup(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Will initiate bna lookup and expect successful response", t, func() {
		response, err := client.CnamLookup(map[string]string{
			"PhoneNumber": "+1XXX", // ONLY +1 NUMBERS
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("cnam_dips")
		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)
	})
}
