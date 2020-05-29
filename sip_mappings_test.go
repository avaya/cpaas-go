// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListSipDomainIpAclMappings(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("ListSipDomainIpAclMappings retrieval successful", t, func() {
		response, err := client.ListSipDomainIpAclMappings("SD8f889084962719d55a674fd7b84f4402", map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("domains")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for transcription := range subResource {
			So(subResource[transcription]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestMapSipDomainIpAcl(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("MapSipDomainIpAcl creation successful", t, func() {
		response, err := client.MapSipDomainIpAcl("SD8f889084962719d55a674fd7b84f4402", map[string]string{
			"IpAccessControlListSid": "AL7c889084ddbef1bd6a6a4bceb421d4e3",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

	})
}

func TestDeleteSipDomainIpAclMapping(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("DeleteSipDomainIpAclMapping retrieval successful", t, func() {
		response, err := client.DeleteSipDomainIpAclMapping("SD8f889084962719d55a674fd7b84f4402", "AL7c889084ddbef1bd6a6a4bceb421d4e2")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "AL7c889084ddbef1bd6a6a4bceb421d4e3")
	})
}

func TestListSipDomainCredentialMappings(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("ListSipDomainCredentialMappings retrieval successful", t, func() {
		response, err := client.ListSipDomainCredentialMappings("SD8f889084962719d55a674fd7b84f4403", map[string]string{
			"Page":     "2",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("domains")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for transcription := range subResource {
			So(subResource[transcription]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func MapSipDomainCredentialList(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("MapSipDomainIpAcl creation successful", t, func() {
		response, err := client.MapSipDomainIpAcl("SD8f889084962719d55a674fd7b84f4402", map[string]string{
			"CredentialListSid": "SCe18890849a99187d6d454d348783d15c",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

	})
}

func TestDeleteSipDomainCredentialListMapping(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("DeleteSipDomainCredentialListMapping retrieval successful", t, func() {
		response, err := client.DeleteSipDomainCredentialListMapping("SD8f889084962719d55a674fd7b84f4402", "AL7c889084ddbef1bd6a6a4bceb421d4e2")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "AL7c889084ddbef1bd6a6a4bceb421d4e3")
	})
}
