// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// Response - Main InboundXML verbs container element. REQUIRED
// More information at: http://docs.zang.io/docs/response
type Response struct {
	Say               *Say               `xml:"Say,omitempty"`
	Play              *Play              `xml:"Play,omitempty"`
	Reject            *Reject            `xml:"Reject,omitempty"`
	Invalid           *Invalid           `xml:"Invalid,omitempty"`
	PlayLastRecording *PlayLastRecording `xml:"PlayLastRecording,omitempty"`
	Dial              *Dial              `xml:"Dial,omitempty"`
	Gather            *Gather            `xml:"Gather,omitempty"`
	Hangup            *Hangup            `xml:"Hangup,omitempty"`
	Pause             *Pause             `xml:"Pause,omitempty"`
	Ping              *Ping              `xml:"Ping,omitempty"`
	Record            *Record            `xml:"Record,omitempty"`
	Redirect          *Redirect          `xml:"Redirect,omitempty"`
	Sms               *Sms               `xml:"Sms,omitempty"`
	Mms               *Mms               `xml:"Mms,omitempty"`
}
