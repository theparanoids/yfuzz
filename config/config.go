// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package config

import (
	"fmt"
	"os"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

// Version holds the current version of YFuzz.
// Injected by the linker, see build.sh
var Version string

// Init reads in configuration from a configuration file and environment variables.
func Init() {
	// Set some sensible defaults
	viper.SetDefault("port", 443)
	viper.SetDefault("log-file", "yfuzz.log")
	viper.SetDefault("access-log-file", "yfuzz-access.log")

	// Tell viper where to look
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.yfuzz")
	viper.AddConfigPath("/etc/yfuzz")

	// Read values from any environment variable of the form YFUZZ_*
	viper.SetEnvPrefix("yfuzz")
	viper.AutomaticEnv()

	// Get config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %s", err))
	}

	// Set up logging
	jww.SetStdoutThreshold(jww.LevelInfo)
	jww.SetLogThreshold(jww.LevelInfo)

	logFile, err := os.OpenFile(viper.GetString("log-file"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("fatal error opening log file: %s", err))
	}
	jww.SetLogOutput(logFile)
}

// InitFake loads a fake configuration for running tests.
func InitFake() {
	viper.SetConfigName("test-config")
	viper.AddConfigPath("../test-data")
	err := viper.ReadInConfig()

	jww.SetStdoutThreshold(jww.LevelInfo)

	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %s", err))
	}
}
