// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package mtls

import (
	"github.com/gorilla/mux"
	"github.com/yahoo/yfuzz/services/yfuzz-server/api"
)

// type to implement yFuzz plugin interface
type mtls struct{}

// Plugin is the exported yFuzz plugin.
var Plugin mtls

// Register initializes the plugin in yFuzz.
func (mtls) Register(r *mux.Router, _ api.EndpointDependencies) {
	r.Use(Middleware)
}
