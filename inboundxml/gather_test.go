// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGatherElement(t *testing.T) {
	Convey("Gather should pass", t, func() {
		ixml, err := New(Response{Gather: &Gather{
			Input:		 "speech",
			Language:    "en-US",
			Hints:		 "sales",	
			Action:      "http://example.com/example-callback-url/say?example=simple.xml",
			Method:      "GET",
			NumDigits:   4,
			FinishOnKey: "#",
			Say: &Say{
				Value: "Please enter your 4 digit pin.",
			},
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})
}
