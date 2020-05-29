// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

// StringInSlice - Will check if string in list. This is equivalent to python if x in []
func StringInSlice(str string, list []string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}
	return false
}

// ToBool - Will return string value back as bool OR respond with defaults
func ToBool(value string, def bool) bool {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return def
	}

	return b
}

// Random -
func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// MapToUrlValues -
func MapToUrlValues(data map[string]string) *url.Values {
	data_values := &url.Values{}

	for key, value := range data {
		data_values.Add(key, value)
	}

	return data_values
}
