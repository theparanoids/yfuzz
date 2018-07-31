// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"github.com/yahoo/yfuzz/pkg/version"
	"github.com/yahoo/yfuzz/services/yfuzz-server/api"
	"github.com/yahoo/yfuzz/services/yfuzz-server/config"
	"github.com/yahoo/yfuzz/services/yfuzz-server/kubernetes"
	"github.com/yahoo/yfuzz/services/yfuzz-server/plugins"
)

// Helper for converting the string loaded from the settings file to the correct type
func stringToAuthType(s string) tls.ClientAuthType {
	switch s {
	case "NoClientCert":
		return tls.NoClientCert
	case "RequestClientCert":
		return tls.RequestClientCert
	case "RequireAnyClientCert":
		return tls.RequireAnyClientCert
	case "VerifyClientCertIfGiven":
		return tls.VerifyClientCertIfGiven
	case "RequireAndVerifyClientCert":
		return tls.RequireAndVerifyClientCert
	default:
		jww.WARN.Printf("Warning: %s is not a valid client auth type.\n", s)
		return tls.VerifyClientCertIfGiven
	}
}

func main() {
	config.Init()

	jww.INFO.Printf("yFuzz %s, built on %s\n", version.Version, version.Timestamp)

	router := mux.NewRouter()

	// Add some basic wrappers to log requests and catch panics
	accessFile, err := os.OpenFile(viper.GetString("access-log-file"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("fatal error opening access log file: %s", err))
	}
	router.Use(handlers.RecoveryHandler())
	router.Use(func(h http.Handler) http.Handler {
		return handlers.CombinedLoggingHandler(accessFile, h)
	})

	// Set up dependencies for endpoints
	kubernetesAPI, err := kubernetes.New()
	if err != nil {
		panic(fmt.Sprintf("Could not connect to kubernetes API: %s", err.Error()))
	}
	dependencies := api.EndpointDependencies{
		Kubernetes: kubernetesAPI,
	}

	// Register our endpoints
	router.Methods("GET").Path("/version").Handler(api.Endpoint(api.GetVersion, dependencies))
	router.Methods("GET").Path("/jobs").Handler(api.Endpoint(api.ListJobs, dependencies))
	router.Methods("GET").Path("/jobs/{job}").Handler(api.Endpoint(api.GetJobStatus, dependencies))
	router.Methods("GET").Path("/jobs/{job}/logs").Handler(api.Endpoint(api.GetJobLogs, dependencies))
	router.Methods("POST").Path("/jobs").Handler(api.Endpoint(api.CreateJob, dependencies))
	router.Methods("DELETE").Path("/jobs/{job}").Handler(api.Endpoint(api.DeleteJob, dependencies))

	// Register plugins
	plugins.Register(router, dependencies)

	// Set up TLS
	// Different authentication plugsins will require different client auth types.
	tlsConfig := &tls.Config{
		ClientAuth: stringToAuthType(viper.GetString("tls.client-auth-type")),
	}

	// Optional: require certificates to be signed by a custom CA.
	if caCertFile := viper.GetString("tls.ca-cert-file"); caCertFile != "" {
		caCertBlock, err := ioutil.ReadFile(caCertFile)
		if err != nil {
			panic("Can't read TLS cert file! " + err.Error())
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCertBlock)
		tlsConfig.ClientCAs = caCertPool
	}

	port := viper.GetString("port")
	server := &http.Server{
		Addr:      ":" + port,
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	cert := viper.GetString("tls.cert-file")
	key := viper.GetString("tls.key-file")

	// Launch the server
	jww.INFO.Printf("About to listen on %s\n", port)
	jww.FATAL.Println(server.ListenAndServeTLS(cert, key))
}
