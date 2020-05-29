// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDialElement(t *testing.T) {
	Convey("Dial should pass", t, func() {
		ixml, err := New(Response{Dial: &Dial{
			Action: "http://webhookr.com/cpaas-inbound-dial-example",
			Method: "GET",
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})
}
