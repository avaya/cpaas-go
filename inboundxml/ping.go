// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// The <Ping> element is provided as a flexible way to send the default voice request parameters to a URL during a call.
// The URL to ping is nested within the opening and closing tags of this element.
type Ping struct {
	Value  string `xml:",chardata"`
	Method string `xml:"method,attr,omitempty"`
}
