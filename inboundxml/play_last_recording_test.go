// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlayLastRecordingElement(t *testing.T) {
	Convey("PlayLastRecording should pass", t, func() {
		ixml, err := New(Response{PlayLastRecording: &PlayLastRecording{}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response><PlayLastRecording></PlayLastRecording></Response>`)
	})
}
