// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package plugins

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/yahoo/yfuzz/services/yfuzz-server/api"
	"github.com/yahoo/yfuzz/services/yfuzz-server/plugins/athenz"
	"github.com/yahoo/yfuzz/services/yfuzz-server/plugins/mtls"
)

// Plugin is the interface used for any addons to yFuzz.
type Plugin interface {
	Register(r *mux.Router, e api.EndpointDependencies)
}

// Plugins lists all available plugins.
var Plugins = map[string]Plugin{
	"athenz": athenz.Plugin,
	"mtls":   mtls.Plugin,
}

// Register adds handlers and middleware from plugins.
func Register(r *mux.Router, e api.EndpointDependencies) {
	for name := range viper.GetStringMap("plugins") {
		if plugin, ok := Plugins[name]; ok {
			plugin.Register(r, e)
		} else {
			panic(fmt.Sprintf("Unknown plugin: %s", name))
		}
	}
}
