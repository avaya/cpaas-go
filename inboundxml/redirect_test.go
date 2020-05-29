// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRedirectElement(t *testing.T) {
	Convey("Redirect should pass", t, func() {
		ixml, err := New(Response{Redirect: &Redirect{
			Value:  "http://example.com/rest.xml",
			Method: "POST",
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})
}
