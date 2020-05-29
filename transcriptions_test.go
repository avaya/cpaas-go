// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListTranscriptions(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("List transcriptions retrieval successful", t, func() {
		response, err := client.ListTranscriptions(map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("transcriptions")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for transcription := range subResource {
			So(subResource[transcription]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestGetTranscription(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Transcription retrieval successful", t, func() {
		response, err := client.GetTranscription("TR24b6d24c57c94c58b3d6f35e0ca9bd71")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "TRO07889084c63ecd8c3c3f49c7b552ba6f")
	})
}

func TestTranscribeRecording(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Transcription retrieval successful", t, func() {
		response, err := client.TranscribeRecording("RE24b6d24c57c94c58b3d6f35e0ca9bd71", map[string]string{
			"TranscribeCallback": "http://webhookr.com/testing-12345",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "TRO07889084c63ecd8c3c3f49c7b552ba6f")
	})
}

func TestTranscribeAudioUrl(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Transcription retrieval successful", t, func() {
		response, err := client.TranscribeAudioUrl(map[string]string{
			"AudioUrl":           "http://",
			"TranscribeCallback": "http://webhookr.com/testing-12345",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "TRO07889084c63ecd8c3c3f49c7b552ba6f")
	})
}
