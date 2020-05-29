// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListCredentialLists(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("ListCredentialLists retrieval successful", t, func() {
		response, err := client.ListCredentialLists(map[string]string{
			"Page":     "1",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("credential_lists")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for transcription := range subResource {
			So(subResource[transcription]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestGetCredentialList(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("GetCredentialList retrieval successful", t, func() {
		response, err := client.GetCredentialList("CL22889084b07c733fd6ec4a30842bce3a")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "CL22889084b07c733fd6ec4a30842bce3a")
	})
}

func TestCreateCredentialList(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("CreateCredentialList creation successful", t, func() {
		response, err := client.CreateCredentialList(map[string]string{
			"FriendlyName": "Testing",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "Testing")
	})
}

func TestUpdateCredentialList(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("UpdateCredentialList retrieval successful", t, func() {
		response, err := client.UpdateCredentialList("CL22889084b07c733fd6ec4a30842bce3a", map[string]string{
			"FriendlyName": "testing-update",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "testing-update")
	})
}

func TestDeleteCredentialList(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Transcription retrieval successful", t, func() {
		response, err := client.DeleteCredentialList("CL22889084b07c733fd6ec4a30842bce3b")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "CL22889084b07c733fd6ec4a30842bce3a")
	})
}

func TestListSipCredentials(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("ListSipCredentials retrieval successful", t, func() {
		response, err := client.ListSipCredentials("CL22889084b07c733fd6ec4a30842bce3a", map[string]string{
			"Page":     "1",
			"PageSize": "1",
		})
		So(response, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)

		subResource := response.GetSubResource("credentials")

		So(len(subResource), ShouldBeGreaterThanOrEqualTo, 0)

		for credential := range subResource {
			So(subResource[credential]["date_updated"], ShouldNotBeBlank)
		}
	})
}

func TestGetSipCredential(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("GetSipCredential retrieval successful", t, func() {
		response, err := client.GetSipCredential("CL22889084b07c733fd6ec4a30842bce3a", "SCe18890849a99187d6d454d348783d15c")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "SCe18890849a99187d6d454d348783d15c")
	})
}

func TestCreateSipCredential(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("CreateCredentialList creation successful", t, func() {
		response, err := client.CreateSipCredential("CL22889084b07c733fd6ec4a30842bce3a", map[string]string{
			"FriendlyName": "Testing",
			"Username":     "testing-gohelper",
			"Password":     "somepassword123",
		})

		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "Testing")
	})
}

func TestUpdateSipCredential(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("UpdateSipCredential retrieval successful", t, func() {
		response, err := client.UpdateSipCredential("CL22889084b07c733fd6ec4a30842bce3a", "SCe18890849a99187d6d454d348783d15c", map[string]string{
			"FriendlyName": "testing-update",
		})
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["friendly_name"], ShouldEqual, "testing-update")
	})
}

func TestDeleteSipCredential(t *testing.T) {
	client, err := NewClient()

	Convey("New Client With Defaults", t, func() {
		So(client, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Transcription retrieval successful", t, func() {
		response, err := client.DeleteSipCredential("CL22889084b07c733fd6ec4a30842bce3a", "SCe18890849a99187d6d454d348783d15b")
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)

		So(response.Response, ShouldHaveSameTypeAs, &http.Response{})
		So(response.Resource, ShouldNotBeNil)
		So(response.Resource["sid"], ShouldEqual, "SCe18890849a99187d6d454d348783d15c")
	})
}
