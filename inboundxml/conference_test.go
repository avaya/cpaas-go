// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConferenceElement(t *testing.T) {
	Convey("Conference should pass", t, func() {
		ixml, err := New(Response{Dial: &Dial{
			Conference: &Conference{
				StartConferenceOnEnter: true,
				CallbackUrl:            "http://webhookr.com/conf-callback",
				HangupOnStar:           true,
				MaxParticipants:        5,
			},
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})

}
