// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"net/http"

	"github.com/yahoo/yfuzz/api/schema"
	"github.com/yahoo/yfuzz/config"
)

// GetVersion retrieves the server version.
// Handler for GET /version
//
// Parameters:
//	- none
//
// Response Fields:
//	- version: version of the YFuzz server (string)
func GetVersion(r *http.Request, _ EndpointDependencies) (int, interface{}) {
	return http.StatusOK, schema.GetVersionResponse{Version: config.Version}
}
