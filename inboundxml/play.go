// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// Say - InboundXML element plays an audio file for the caller.
// More information at: http://docs.zang.io/docs/play
//
// Example:
// ixml, err := New(Response{Play: &Play{
//			Loop: 1,
//			Value: "tone_stream://%(10000,0,350,440)",
// }})
// fmt.Println(inxml.String())
type Play struct {
	Value  string `xml:",chardata"`
	Loop   int    `xml:"loop,attr,omitempty"`
	Method string `xml:"method,attr,omitempty"`
}
