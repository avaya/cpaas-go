// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// The <Gather> element allows callers to input digits to the call using their keypads which are then sent via POST or GET to a URL for further processing.
// There are many ways to get creative with <Gather> but its most common use case is in creating IVR menus.
// This is accomplished by nesting prompts for input from the caller using the <Say> or <Play> elements.

// By default, an unlimited number of digits can be gathered.
// <Gather> will timeout after 5 seconds or once the ‘#’ key is pressed.
// The gathered digits will then be submitted to the current InboundXML document.
// This default behavior of <Gather> can be altered using the provided element attributes.
type Gather struct {
	Input       string `xml:"input,attr,omitempty"`
	Language    string `xml:"language,attr,omitempty"`
	Hints       string `xml:"hints,attr,omitempty"`
	Action      string `xml:"action,attr,omitempty"`
	Method      string `xml:"method,attr,omitempty"`
	Timeout     int    `xml:"timeout,attr,omitempty"`
	FinishOnKey string `xml:"finishOnKey,attr,omitempty"`
	NumDigits   int    `xml:"numDigits,attr,omitempty"`
	Say         *Say   `xml:"Say,omitempty"`
	Play        *Play  `xml:"Play,omitempty"`
	Pause       *Pause `xml:"Pause,omitempty"`
}
