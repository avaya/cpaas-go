// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// The <Hangup> element will hangup the current call.
type Hangup struct {
	Schedule int    `xml:"schedule,attr,omitempty"`
	Reason   string `xml:"reason,attr,omitempty"`
}
