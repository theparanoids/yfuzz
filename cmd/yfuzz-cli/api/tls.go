// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// newClient creates a http client with a user x509 certificate.
func newClient() (*http.Client, error) {
	keyBlock, err := ioutil.ReadFile(viper.GetString("tls.user-key"))
	if err != nil {
		return nil, err
	}

	certBlock, err := ioutil.ReadFile(viper.GetString("tls.user-cert"))
	if err != nil {
		return nil, err
	}

	cert, err := tls.X509KeyPair(certBlock, keyBlock)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tlsConfig.BuildNameToCertificate()

	if caCertFile := viper.GetString("tls.ca-cert"); caCertFile != "" {
		caCertBlock, err := ioutil.ReadFile(caCertFile)
		if err != nil {
			return nil, err
		}
		tlsConfig.RootCAs.AppendCertsFromPEM(caCertBlock)
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	return client, nil
}
