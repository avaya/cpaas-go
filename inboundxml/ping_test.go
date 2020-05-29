// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPingElement(t *testing.T) {
	Convey("Ping should pass", t, func() {
		ixml, err := New(Response{Ping: &Ping{
			Value: "http://webhookr.com/ping-test",
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})
}
