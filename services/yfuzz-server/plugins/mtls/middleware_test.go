// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package mtls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
)

// Ensure authenticated requests go through and plain requests do not
func TestMiddleware(t *testing.T) {
	// Test user's x509 keypair
	cert, err := tls.LoadX509KeyPair("./test_data/testCert.pem", "./test_data/testPrivateKey.pem")
	if err != nil {
		t.Fatalf("Failed setup: %s", err.Error())
	}
	// Test user's public key
	publicKey, err := ioutil.ReadFile("./test_data/testPublicKey.pem")
	if err != nil {
		t.Fatalf("Failed setup: %s", err.Error())
	}
	// Public key of a different test user
	otherPublicKey, err := ioutil.ReadFile("./test_data/otherPublicKey.pem")
	if err != nil {
		t.Fatalf("Failed setup: %s", err.Error())
	}

	// Set up the server
	server := httptest.NewTLSServer(Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})))
	defer server.Close()
	server.TLS.ClientAuth = tls.RequireAnyClientCert

	// Set up a client to use the keys above
	client := server.Client()
	rootCAs := x509.NewCertPool()
	rootCAs.AddCert(server.Certificate())
	clientConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      rootCAs,
	}
	client.Transport = &http.Transport{TLSClientConfig: clientConfig}

	// Request where the key is not in the whitelist
	viper.Set("plugins.mtls.authorized-keys", []string{string(otherPublicKey)})
	res, err := client.Get(server.URL)
	if err != nil {
		t.Errorf("Failed request: %s", err.Error())
	} else if res.StatusCode != http.StatusForbidden {
		t.Errorf("Request with non-whitelisted key should be forbidden.")
	}

	// Request where they key is in the whitelist
	viper.Set("plugins.mtls.authorized-keys", []string{string(otherPublicKey), string(publicKey)})
	res, err = client.Get(server.URL)
	if err != nil {
		t.Errorf("Failed request: %s", err.Error())
	} else if res.StatusCode != http.StatusOK {
		t.Errorf("Request with whitelisted key should be authenticated.")
	}
}
