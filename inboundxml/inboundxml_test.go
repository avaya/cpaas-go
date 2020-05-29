// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestInboundXMLResponse -
func TestInboundXMLResponse(t *testing.T) {
	Convey("New InboundXML should return valid empty response", t, func() {
		ixml, err := New(Response{})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("New InboundXML render should return empty response", t, func() {
		ixml, err := New(Response{})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

		r, rerr := ixml.Marshal()
		So(r, ShouldEqual, `<Response></Response>`)
		So(rerr, ShouldBeNil)

		So(ixml.String(), ShouldEqual, `<Response></Response>`)
	})

	Convey("New InboundXML render should be invoked with new", t, func() {
		ixml, err := New(Response{})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response></Response>`)
	})

	Convey("New InboundXML document should not pass the validation", t, func() {
		ixml, err := New(Response{Invalid: &Invalid{}})
		So(ixml, ShouldBeNil)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "Document validation error")
	})
}
