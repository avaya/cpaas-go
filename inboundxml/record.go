// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// The <Record> element is used to record audio during a call.
// It can occur anywhere within an InboundXML document but will only begin recording once it has been reached.
// This means, it would have to be the first element after <Response> for the entire call to be recorded.
// When the recording is complete, a URL of the recorded audio is created and submitted as a GET or POST to the action URL.

//Similar to the <Gather> element, a timeout value sets how much silence to allow before the recording ends, maxLength sets how long the recording may be, and the
// finishOnKey is used to set which keys will end the recording.
// By default, the action and method specify that <Record> should make a POST to the URL of the current InboundXML document.

// Note: no more than than 5 recordings may be attached to a call
type Record struct {
	Action             string `xml:"action,attr,omitempty"`
	Method             string `xml:"method,attr,omitempty"`
	Timeout            int    `xml:"timeout,attr,omitempty"`
	FinishOnKey        string `xml:"finishOnKey,attr,omitempty"`
	MaxLength          int    `xml:"maxLength,attr,omitempty"`
	Transcribe         bool   `xml:"transcribe,attr,omitempty"`
	TranscribeQuality  string `xml:"transcribeQuality,attr,omitempty"`
	TranscribeCallback string `xml:"transcribeCallback,attr,omitempty"`
	PlayBeep           bool   `xml:"playBeep,attr,omitempty"`
	Direction          string `xml:"direction,attr,omitempty"`
	FileFormat         string `xml:"fileFormat,attr,omitempty"`
	Background         bool   `xml:"background,attr,omitempty"`
	TrimSilence        bool   `xml:"trimSilence,attr,omitempty"`
}
