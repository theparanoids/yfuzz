// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package mtls

import (
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"reflect"

	"github.com/spf13/viper"

	jww "github.com/spf13/jwalterweatherman"
)

// Middleware initializes a connection to Athenz for use in verifying users have permission to access yFuzz.
func Middleware(h http.Handler) http.Handler {
	// Handler checks that the public key from user's certificate is in the list of authorized keys.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorized, err := getAuthorization(r.TLS.PeerCertificates)
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

// Compare the public keys in the peer certificates to a whitelist.
func getAuthorization(peerCertificates []*x509.Certificate) (bool, error) {
	// Use an interface, because the keys could be RSA, DSA, ECDSA. We don't care, as long as they match.
	var authorizedKeys []interface{}

	for i, rawKey := range viper.GetStringSlice("plugins.mtls.authorized-keys") {
		block, _ := pem.Decode([]byte(rawKey))
		if block == nil {
			jww.INFO.Printf("Cannot decode public key %d in authorized keys.\n", i)
			continue
		}
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			jww.INFO.Printf("Cannot parse public key %d in authorized keys: %s\n", i, err.Error())
			continue
		}
		authorizedKeys = append(authorizedKeys, key)
	}

	// Quadratic, but there shouldn't be many peer certificates on the request.
	for _, cert := range peerCertificates {
		for _, key := range authorizedKeys {
			if reflect.DeepEqual(cert.PublicKey, key) {
				return true, nil
			}
		}
	}

	// If there was no match, the user doesn't have access.
	return false, nil
}
