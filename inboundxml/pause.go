// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// The <Pause> element will pause the call for the number of seconds set by the length attribute.
type Pause struct {
	Length int `xml:"length,attr,omitempty"`
}
