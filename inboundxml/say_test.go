// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestInboundXMLResponse -
func TestSayVerb(t *testing.T) {
	Convey("New InboundXML should return valid say response", t, func() {
		ixml, err := New(Response{Say: &Say{}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response><Say></Say></Response>`)
	})

	Convey("New InboundXML should return valid say response (with content)", t, func() {
		ixml, err := New(Response{Say: &Say{Value: "Lorem Ipsum dolor sit amet"}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response><Say>Lorem Ipsum dolor sit amet</Say></Response>`)
	})

	Convey("Say, if voice attribute passed, requires valid voice", t, func() {
		ixml, err := New(Response{Say: &Say{
			Voice: "be",
			Value: "Lorem Ipsum dolor sit amet",
		}})
		So(ixml, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("Say voice attribute renders correctly", t, func() {
		ixml, err := New(Response{Say: &Say{
			Voice: "female",
			Value: "Lorem Ipsum dolor sit amet",
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response><Say voice="female">Lorem Ipsum dolor sit amet</Say></Response>`)
	})

	Convey("Say language attribute renders correctly", t, func() {
		ixml, err := New(Response{Say: &Say{
			Voice:    "female",
			Language: "en-gb",
			Value:    "Lorem Ipsum dolor sit amet",
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response><Say voice="female" language="en-gb">Lorem Ipsum dolor sit amet</Say></Response>`)
	})

	Convey("Say loop attribute renders correctly", t, func() {
		ixml, err := New(Response{Say: &Say{
			Voice:    "female",
			Language: "en",
			Loop:     10,
			Value:    "Lorem Ipsum dolor sit amet",
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(ixml.String(), ShouldEqual, `<Response><Say voice="female" language="en" loop="10">Lorem Ipsum dolor sit amet</Say></Response>`)
	})
}
