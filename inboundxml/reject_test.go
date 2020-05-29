// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRejectVerb(t *testing.T) {
	Convey("Reject should pass if no reason is provided", t, func() {
		ixml, err := New(Response{Reject: &Reject{}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response><Reject></Reject></Response>`)
	})

	Convey("Reject should pass if valid reason is provided", t, func() {
		ixml, err := New(Response{Reject: &Reject{Reason: "busy"}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response><Reject reason="busy"></Reject></Response>`)
	})

	Convey("Reject should fail if invalid reason is passed", t, func() {
		ixml, err := New(Response{Reject: &Reject{Reason: "busy2"}})
		So(ixml, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}
