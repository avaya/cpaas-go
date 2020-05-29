// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// Like the <Number> element, the <Conference> element is only nested within the <Dial> element.
// Instead of dialing a number, the <Conference> element allows the ongoing call to connect to a conference room.

// The conference room name is nested within the <Conference> element, and all calls connected to the same conference room name will be in the room together.

// By default, all callers will hear hold music until two callers are in the room.
// To change this behavior, startConferenceOnEnter may be set to ‘true’ or ‘false’.
// The waitUrl attribute may be used to set a custom MP3.
// For greater customization, the URL of an InboundXML document can be used while callers are waiting as well.
// If an InboundXML document is used, the <Gather>, <Record> and <Dial> elements are not allowed.

// When callers enter or exit the room, a beep is heard if the beep attribute is set to the default value of ‘true’.
// A participant can be initially muted by setting the muted attribute to ‘true’.

// The conference room can be limited to a certain number of participants by setting the maxParticipants attribute.
// The endConferenceOnExit attribute is used to end a conference when an specific user (or any one of many users) exits.
type Conference struct {
	Value                  string `xml:",chardata"`
	Muted                  bool   `xml:"muted,attr,omitempty"`
	Beep                   bool   `xml:"beep,attr,omitempty"`
	StartConferenceOnEnter bool   `xml:"startConferenceOnEnter,attr,omitempty"`
	EndConferenceOnExit    bool   `xml:"endConferenceOnExit,attr,omitempty"`
	MaxParticipants        int    `xml:"maxParticipants,attr,omitempty"`
	WaitSound              string `xml:"waitSound,attr,omitempty"`
	HangupOnStar           bool   `xml:"hangupOnStar,attr,omitempty"`
	CallbackUrl            string `xml:"callbackUrl,attr,omitempty"`
	CallbackMethod         string `xml:"callbackMethod,attr,omitempty"`
	DigitsMatch            string `xml:"digitsMatch,attr,omitempty"`
	StayAlone              bool   `xml:"stayAlone,attr,omitempty"`
	Record                 bool   `xml:"record,attr,omitempty"`
	RecordCallbackUrl      string `xml:"recordCallbackUrl,attr,omitempty"`
	RecordFileFormat       string `xml:"recordFileFormat,attr,omitempty"`
}
