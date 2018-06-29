// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package schema

// GetVersionResponse ...
// Schema for GET /version response
type GetVersionResponse struct {
	Version string `json:"version"`
}
