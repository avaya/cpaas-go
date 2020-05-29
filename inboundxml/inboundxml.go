// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"unsafe"

	"github.com/jbussdieker/golibxml"
	"github.com/krolaw/xsd"
)

// InboundXML -
type InboundXML struct {
	Response    Response `xml:"Response"`
	RenderCache []byte
}

// GetSchema -
func (i *InboundXML) GetSchema() ([]byte, error) {
	b, err := ioutil.ReadFile("./schema/inboundxml.xsd")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Validate -
func (i *InboundXML) Validate() error {
	rdoc, err := i.Marshal()
	if err != nil {
		return err
	}

	schema, scherr := i.GetSchema()
	if scherr != nil {
		return scherr
	}

	xsdSchema, err := xsd.ParseSchema([]byte(schema))
	if err != nil {
		return err
	}

	doc := golibxml.ParseDoc(rdoc)
	if doc == nil {
		return fmt.Errorf("Could not parse xml document: %v", rdoc)
	}
	defer doc.Free()

	if err := xsdSchema.Validate(xsd.DocPtr(unsafe.Pointer(doc.Ptr))); err != nil {
		return err
	}

	return nil
}

// Marshal -
func (i *InboundXML) Marshal() (string, error) {
	ixml, err := xml.Marshal(i.Response)
	if err != nil {
		return "", err
	}

	// Save cache here so later on we can use the same for .String()
	i.RenderCache = ixml

	return string(ixml), nil
}

// String -
func (i *InboundXML) String() (response string) {
	if i.RenderCache != nil {
		response = string(i.RenderCache)
	}
	return
}

// New -
func New(r Response) (*InboundXML, error) {
	document := &InboundXML{Response: r}

	if err := document.Validate(); err != nil {
		return nil, err
	}

	return document, nil
}
