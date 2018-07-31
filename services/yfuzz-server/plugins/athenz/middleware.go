// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package athenz

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"github.com/yahoo/athenz/clients/go/zms"
)

// Middleware initializes a connection to Athenz for use in verifying users have permission to access yFuzz.
func Middleware(h http.Handler) http.Handler {
	// Set up client for athenz
	keyPEM, err := ioutil.ReadFile(viper.GetString("middleware.athenz.key-file"))
	if err != nil {
		panic(err)
	}
	certPEM, err := ioutil.ReadFile(viper.GetString("middleware.athenz.cert-file"))
	if err != nil {
		panic(err)
	}

	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	authURL := viper.GetString("middleware.athenz.url")

	client := zms.NewClient(authURL, tr)

	// Handler checks that the user making the request has access to the yFuzz service in Athenz.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorized, err := getAuthorization(client, r.TLS.VerifiedChains)
		if err != nil {
			jww.WARN.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if authorized {
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// Extract the principal from the user x.509 certificate and query athenz to see if that user is authorized to access yFuzz
func getAuthorization(client zms.ZMSClient, verifiedChains [][]*x509.Certificate) (bool, error) {
	// Find the cert for athenz
	for _, chain := range verifiedChains {
		for _, cert := range chain {
			if cert.Issuer.CommonName == viper.GetString("middleware.athenz.ca-issuer-name") {
				principal := zms.EntityName(cert.Subject.CommonName)
				action := zms.ActionName(viper.GetString("middleware.athenz.action"))
				resource := zms.ResourceName(viper.GetString("middleware.athenz.resource"))

				access, err := client.GetAccess(action, resource, "", principal)
				if err != nil {
					return false, err
				}

				if access.Granted {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
