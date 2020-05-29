// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// The <Dial> element starts an outgoing dial from the current call.
// Once the dial is complete, the next element in the InboundXML document will be processed unless the action attribute is set.
// If the action attribute is set, the result of the dial is submitted as a GET or POST (depending on the method attribute) to the action URL, and the call will continue using the InboundXML in said URL.

// By default, the outgoing call will timeout after 30 seconds if it is not answered. However, the timeout attribute can be used to set a custom time.
// The length of the call is limited by the timeLimit attribute, which is 4 hours (14400 seconds) by default.

// Setting the hangupOnStar attribute to ‘true’ will allow the original call to terminate the outgoing call without having to hangup by dialing ‘*’.
// The original call then continues with the current InboundXML document, or if a URL was passed, the response from the action attribute.

// The callerId attribute can be set to any number, and will default to the caller id of the original caller.
// The number to be dialed should be nested within the <Dial> element.
// For more options, use the <Number>, <Sip> or <Conference> elements instead of a simple phone number.
type Dial struct {
	Value             string      `xml:",chardata"`
	Action            string      `xml:"action,attr,omitempty"`
	Method            string      `xml:"method,attr,omitempty"`
	TimeLimit         int         `xml:"timeLimit,attr,omitempty"`
	CallerId          string      `xml:"callerId,attr,omitempty"`
	HideCallerId      bool        `xml:"hideCallerId,attr,omitempty"`
	DialMusic         string      `xml:"dialMusic,attr,omitempty"`
	CallbackUrl       string      `xml:"callbackUrl,attr,omitempty"`
	CallbackMethod    string      `xml:"callbackMethod,attr,omitempty"`
	ConfirmSound      string      `xml:"confirmSound,attr,omitempty"`
	DigitsMatch       string      `xml:"digitsMatch,attr,omitempty"`
	StraightToVm      bool        `xml:"straightToVm,attr,omitempty"`
	HeartbeatUrl      string      `xml:"heartbeatUrl,attr,omitempty"`
	HeartbeatMethod   string      `xml:"heartbeatMethod,attr,omitempty"`
	ForwardedFrom     int         `xml:"forwardedFrom,attr,omitempty"`
	IfMachine         string      `xml:"ifMachine,attr,omitempty"`
	IfMachineUrl      string      `xml:"ifMachineUrl,attr,omitempty"`
	IfMachineMethod   string      `xml:"ifMachineMethod,attr,omitempty"`
	Record            bool        `xml:"record,attr,omitempty"`
	RecordDirection   string      `xml:"recordDirection,attr,omitempty"`
	RecordCallbackUrl string      `xml:"recordCallbackUrl,attr,omitempty"`
	Conference        *Conference `xml:"Conference,omitempty"`
	Number            *Number     `xml:"Number,omitempty"`
	Sip               *Sip        `xml:"Sip,omitempty"`
}
