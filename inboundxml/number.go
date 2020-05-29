// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// The <Number> element must be nested within the <Dial> element.
// <Number> can be used to dial multiple phones simultaneously by using additional <Number> elements.
// If multiple <Number> elements are used to specify additional calls, the first call to answer is connected, and the rest of the outgoing calls are canceled.
// <Number> can also be used to send DTMF tones to called parties with the sendDigits attribute.
type Number struct {
	Value      string `xml:",chardata"`
	SendDigits string `xml:"sendDigits,attr,omitempty"`
}
