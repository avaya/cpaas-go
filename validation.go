// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"
	"strings"
)

// ValidateAccountSid -
func ValidateSid(sid string) error {
	if len(sid) != 34 {
		return fmt.Errorf("Make sure to provide valid SID. You've provided: %s", sid)
	}
	return nil
}

// ValidateAccountSid -
func ValidateAccountSid(accountSid string) error {
	if len(accountSid) != 34 || !strings.HasPrefix(accountSid, "AC") {
		return fmt.Errorf("Make sure to provide valid AccountSid. You've provided: %s", accountSid)
	}
	return nil
}

// ValidateAuthToken -
func ValidateAuthToken(authToken string) error {
	if len(authToken) != 32 {
		return fmt.Errorf("Make sure to provide valid AuthToken. You've provided: %s", authToken)
	}
	return nil
}
