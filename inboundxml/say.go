// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// Say - InboundXML Verb designed to do a TTS speech towards callee
// More information at: http://docs.zang.io/docs/say
//
// Example:
// ixml, err := New(Response{Say: &Say{
//			Voice: "female",
//			Value: "Lorem Ipsum dolor sit amet",
// }})
// fmt.Println(inxml.String())
type Say struct {
	Value    string `xml:",chardata"`
	Voice    string `xml:"voice,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
	Loop     int    `xml:"loop,attr,omitempty"`
}
