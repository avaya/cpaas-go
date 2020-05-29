// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumberElement(t *testing.T) {
	Convey("Number should pass", t, func() {
		ixml, err := New(Response{Dial: &Dial{
			Number: &Number{
				Value:      "(555)555-5555",
				SendDigits: "ww12w3221",
			},
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})
}
