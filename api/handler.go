// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/yahoo/yfuzz/kubernetes"
)

// EndpointDependencies holds objects endpoints rely on, such as a Kubernetes or database connection.
type EndpointDependencies struct {
	Kubernetes *kubernetes.API
}

// Signature for an endpoint function that can be turned into a handler.
type endpointFunc func(*http.Request, EndpointDependencies) (int, interface{})

// Endpoint wraps an endpoint handler in some logic to catch errors and marshal a response to JSON
func Endpoint(handler endpointFunc, dependencies EndpointDependencies) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status, data := handler(r, dependencies)

		marshaledResponse, err := json.Marshal(data)
		if err != nil {
			jww.WARN.Printf("Failed to serialize: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(status)

		if data != nil {
			w.Header().Set("content-type", "application/json")
			w.Write(marshaledResponse)
		}
	})
}

// Wrappers for generating responses that will serialize nicely
type messageHolder struct {
	Message string `json:"message"`
}

func responseMessage(message string, i ...interface{}) messageHolder {
	return messageHolder{fmt.Sprintf(message, i...)}
}

func responseFromError(err error) messageHolder {
	return messageHolder{err.Error()}
}
