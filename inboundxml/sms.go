// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// Avaya CPaaS allows SMS messages to be sent during a call using the <Sms> element.
// The SMS receiver (to attribute) must be a valid phone number.
// The sender (from attribute) must be one of your registered Avaya CPaaS numbers.
// The text of the message should be placed inside the element and can not be longer than 160 characters.

// The action attribute can be used to direct the SMS to a new InboundXML document for processing.
// If directed to a new InboundXML using the action attribute, all InboundXML beneath the <Sms> element in the originating InboundXML document is disregarded.
// Similarly, the statusCallback attribute is provided to report the outcome of the SMS transmission.
type Sms struct {
	Value          string `xml:",chardata"`
	To             string `xml:"to,attr,omitempty"`
	From           string `xml:"from,attr,omitempty"`
	Action         string `xml:"action,attr,omitempty"`
	Method         string `xml:"method,attr,omitempty"`
	StatusCallback string `xml:"statusCallback,attr,omitempty"`
}
