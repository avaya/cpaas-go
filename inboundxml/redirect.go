// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// The <Redirect> element directs the call to another InboundXML document.
// The URL of the InboundXML document is nested within the <Redirect> element, and the method attribute sets if the request will be a GET or POST.
// In addition to the default voice request parameters, the parameter UrlBase will also be forwarded when the redirect request is made.
// UrlBase points to the base InboundXML document where the <Redirect> occurred.
type Redirect struct {
	Value  string `xml:",chardata"`
	Method string `xml:"method,attr,omitempty"`
}
