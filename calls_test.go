// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListCalls(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List calls retrieval successful", t, func() {
		response, err := client.ListCalls(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("calls")

		So(len(subResource), ShouldBeGreaterThan, 0)

		for application := range subResource {
			So(subResource[application]["date_updated"], ShouldNotBeBlank)
		}
	})

}

func TestGetCall(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Call retrieval successful", t, func() {
		response, err := client.GetCall("CA42889084446a6643ee2c4950a4d38b2d")
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		So(response.Resource["sid"], ShouldEqual, "CA42889084446a6643ee2c4950a4d38b2d")
	})
}

func TestCreateCalls(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Create call", t, func() {
		response, err := client.CreateCall(map[string]string{
			"From": "+XXX",
			"To":   "+XXX",
			"Url":  "http://static.zang.io/ivr/welcome/call.xml",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldNotBeEmpty)
	})
}

func TestInterruptCall(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Should interrupt the call", t, func() {
		response, err := client.InterruptCall("CA42889084446a6643ee2c4950a4d38b2d", map[string]string{
			"Url":    "http://static.zang.io/ivr/welcome/call.xml",
			"Method": "GET",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldNotBeEmpty)
	})
}

func TestCallSendDigits(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Should send digits the call", t, func() {
		response, err := client.SendDigitsCall("CA42889084446a6643ee2c4950a4d38b2d", map[string]string{
			"Digits": "12345",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldNotBeEmpty)
	})
}

func TestRecordCall(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Should start call recording", t, func() {
		response, err := client.RecordCall("CA42889084446a6643ee2c4950a4d38b2d", map[string]string{
			"Record": "true",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldNotBeEmpty)
	})
}

func TestPlayAudioCall(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Should start playing in the call", t, func() {
		response, err := client.PlayAudioCall("CA42889084446a6643ee2c4950a4d38b2d", map[string]string{
			"AudioUrl":  "true",
			"Direction": "both",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldNotBeEmpty)
	})
}

func TestApplyEffectsCall(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Should apply effects to the call", t, func() {
		response, err := client.ApplyEffectsCall("CA42889084446a6643ee2c4950a4d38b2d", map[string]string{
			"Rate":  "10",
			"Tempo": "10",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldNotBeEmpty)
	})
}
