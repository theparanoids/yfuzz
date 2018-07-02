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

func newClient() (*zms.ZMSClient, error) {
	// Set up transport for athenz
	keyPEM, err := ioutil.ReadFile(viper.GetString("auth.athenz.key-file"))
	if err != nil {
		return nil, err
	}
	certPEM, err := ioutil.ReadFile(viper.GetString("auth.athenz.cert-file"))
	if err != nil {
		return nil, err
	}

	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	authURL := viper.GetString("auth.athenz.url")

	client := zms.NewClient(authURL, tr)

	return &client, nil
}

func getAuthorization(verifiedChains [][]*x509.Certificate) (bool, error) {
	client, err := newClient()
	if err != nil {
		return false, err
	}

	// Find the cert for athenz
	for _, chain := range verifiedChains {
		for _, cert := range chain {
			if cert.Issuer.CommonName == viper.GetString("auth.athenz.ca-issuer-name") {
				principal := zms.EntityName(cert.Subject.CommonName)
				action := zms.ActionName(viper.GetString("auth.athenz.action"))
				resource := zms.ResourceName(viper.GetString("auth.athenz.resource"))

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

// Verify checks that the user making the request has access to the YFuzz service in Athenz.
func Verify(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorized, err := getAuthorization(r.TLS.VerifiedChains)
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
