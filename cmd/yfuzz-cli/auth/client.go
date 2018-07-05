// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package auth

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// GetClient creates a client with the proper certs to access the YFuzz api
func GetClient() (*http.Client, error) {
	keyPEMBlock, err := ioutil.ReadFile(viper.GetString("athenz.user-key"))
	if err != nil {
		return nil, err
	}

	certPEMBLock, err := ioutil.ReadFile(viper.GetString("athenz.user-cert"))
	if err != nil {
		return nil, err
	}

	CAPEMBlock, err := ioutil.ReadFile(viper.GetString("athenz.ca-cert"))
	if err != nil {
		return nil, err
	}

	cert, err := tls.X509KeyPair(certPEMBLock, keyPEMBlock)
	if err != nil {
		return nil, err
	}

	caCert, err := x509.ParseCertificate(CAPEMBlock)
	if err != nil {
		return nil, err
	}

	rootCA := x509.NewCertPool()
	rootCA.AddCert(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      rootCA,
	}
	tlsConfig.BuildNameToCertificate()

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	return client, nil
}
