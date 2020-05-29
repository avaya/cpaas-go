// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlayElement(t *testing.T) {
	Convey("Reject should NOT pass if no audio is provided", t, func() {
		ixml, err := New(Response{Play: &Play{}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response><Play></Play></Response>`)
	})

	Convey("Play should render with valid loop element", t, func() {
		ixml, err := New(Response{Play: &Play{Loop: 10, Value: "tone_stream://%(10000,0,350,440)"}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response><Play loop="10">tone_stream://%(10000,0,350,440)</Play></Response>`)
	})

	Convey("Play should render with valid method element", t, func() {
		ixml, err := New(Response{Play: &Play{Loop: 10, Method: "POST", Value: "tone_stream://%(10000,0,350,440)"}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response><Play loop="10" method="POST">tone_stream://%(10000,0,350,440)</Play></Response>`)
	})
}
