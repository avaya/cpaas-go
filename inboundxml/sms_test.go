// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSmsElement(t *testing.T) {
	Convey("Sms should pass", t, func() {
		ixml, err := New(Response{Sms: &Sms{
			From:  "(phone_number)",
			To:    "(phone_number)",
			Value: "Test message sent from Avaya CPaaS!",
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})
}
