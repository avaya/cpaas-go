// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// The <Sip> element is nested within the <Dial> element, and is used to call SIP addresses.
// The desired SIP address to call is nested within <Sip> </Sip>, just as with a number when using the <Number> element.
// However, the opening and closing sip tags (<Sip> </Sip>) may be omitted completely by simply prefixing the desired SIP address with “sip:” when using it within the <Dial> element.
// For example: <Dial>sip:username@domain.com<Dial>

// If multiple <Sip> elements are used, the first call to answer is connected and the rest of the outgoing calls are canceled.
type Sip struct {
	Value    string `xml:",chardata"`
	Username string `xml:"username,attr,omitempty"`
	Password string `xml:"password,attr,omitempty"`
}
