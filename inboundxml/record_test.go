// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRecordElement(t *testing.T) {
	Convey("Record should pass", t, func() {
		ixml, err := New(Response{Record: &Record{
			Action:      "http://webhookr.com/cpaas-inboundxml-recording-action-example",
			Background:  false,
			Method:      "POST",
			FinishOnKey: "#",
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})
}
