// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package schema

// GetVersionResponse is the schema for the response from the GET /version endpoint.
// See https://godoc.org/github.com/yahoo/yfuzz/services/yfuzz-server/api#GetVersion
type GetVersionResponse struct {
	Version string `json:"version"`
}

// MessageHolder is a way to return a string that will serialize nicely.
type MessageHolder struct {
	Message string `json:"message"`
}
