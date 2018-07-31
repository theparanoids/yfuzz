// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Init reads in configuration values from configuration files.
func Init() {
	// Override from file if present
	viper.SetConfigName("cli-config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.yfuzz")
	viper.AddConfigPath("/etc/yfuzz")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Could not read config file, using defaults.")
	}
}
