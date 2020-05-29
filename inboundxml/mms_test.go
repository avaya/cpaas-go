// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMmsElement(t *testing.T) {
	Convey("Mms should pass", t, func() {
		ixml, err := New(Response{Mms: &Mms{
			From:  		 "(phone_number)",
			To:    		 "(phone_number)",
			Method:      "GET",
			Value:		 "This is MMS sent from Zang.",
			MediaUrl:	 "https://media.giphy.com/media/zZJzLrxmx5ZFS/giphy.gif",
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}
