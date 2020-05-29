// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSipElement(t *testing.T) {
	Convey("Sip should pass", t, func() {
		ixml, err := New(Response{Dial: &Dial{
			Sip: &Sip{
				Username: "(someusername)",
				Password: "(somepassword)",
				Value:    "username@domain.com",
			},
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})
}
