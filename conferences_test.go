// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListConferences(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List conferences retrieval successful", t, func() {
		response, err := client.ListConferences(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("conferences")

		So(len(subResource), ShouldBeGreaterThan, 0)

		for application := range subResource {
			So(subResource[application]["date_updated"], ShouldNotBeBlank)
		}
	})

}

func TestGetConference(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Conference retrieval successful", t, func() {
		response, err := client.GetConference("CF42889084446a6643ee2c4950a4d38b2d")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		So(response.Resource["sid"], ShouldEqual, "CF42889084446a6643ee2c4950a4d38b2d")
	})
}

func TestListParticipants(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List participants retrieval successful", t, func() {
		response, err := client.ListParticipants("CF42889084446a6643ee2c4950a4d38b2d", map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("participants")

		So(len(subResource), ShouldBeGreaterThan, 0)

		for application := range subResource {
			So(subResource[application]["date_updated"], ShouldNotBeBlank)
		}
	})

}

func TestGetParticipant(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Conference participant retrieval successful", t, func() {
		response, err := client.GetParticipant("CF42889084446a6643ee2c4950a4d38b2d", "XXXX")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		So(response.Resource["sid"], ShouldEqual, "XXX")
	})
}

func TestMuteDeafParticipant(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Conference participant mute and no deaf successful", t, func() {
		response, err := client.MuteOrDeafParticipant("CF42889084446a6643ee2c4950a4d38b2d", "XXXX", map[string]string{
			"Muted": "true",
			"Deaf":  "false",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		So(response.Resource["sid"], ShouldEqual, "XXX")
	})
}

func TestHangupParticipant(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Conference participant should hangup successfully", t, func() {
		response, err := client.HangupParticipant("CF42889084446a6643ee2c4950a4d38b2d", "XXXX")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		So(response.Resource["sid"], ShouldEqual, "XXX")
	})
}

func TestPlayAudioParticipant(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Conference participant mute and no deaf successful", t, func() {
		response, err := client.PlayParticipant("CF42889084446a6643ee2c4950a4d38b2d", "XXXX", map[string]string{
			"AudioUrl": "http://full-url-to-audio-file.mp3",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		So(response.Resource["sid"], ShouldEqual, "XXX")
	})
}
