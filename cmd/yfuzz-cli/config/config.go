// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Init reads in configuration values.
func Init() {
	home := os.Getenv("HOME")

	// Set the defaults
	viper.SetDefault("athenz.user-cert", home+"/.yfuzz/certs/userx509.pem")
	viper.SetDefault("athenz.user-key", home+"/.yfuzz/certs/userkey.pem")
	viper.SetDefault("athenz.ca-cert", home+"/.yfuzz/certs/athenz-ca.pem")

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
