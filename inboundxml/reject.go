// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// Reject - InboundXML Verb designed reject the call
// More information at: http://docs.zang.io/docs/reject
type Reject struct {
	Reason string `xml:"reason,attr,omitempty"`
}
