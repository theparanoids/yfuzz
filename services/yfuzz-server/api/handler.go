// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/yahoo/yfuzz/pkg/schema"
	"github.com/yahoo/yfuzz/services/yfuzz-server/kubernetes"
)

// Generic error message
const internalServerError = "An internal server error occurred."

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

		// Handle redirects
		if http.StatusMultipleChoices <= status && status < http.StatusBadRequest {
			if url, ok := data.(string); ok {
				http.Redirect(w, r, url, status)
			} else {
				jww.WARN.Printf("Cannot redirect to location %v of type %T\n", data, data)
				http.Error(w, internalServerError, http.StatusInternalServerError)
			}
			return
		}

		marshaledResponse, err := json.Marshal(data)
		if err != nil {
			jww.WARN.Printf("Failed to serialize: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(status)

		if data != nil {
			w.Header().Set("content-type", "application/json")
			_, err := w.Write(marshaledResponse)
			if err != nil {
				jww.WARN.Printf("Failed to serialize error: %s\n", err.Error())
			}
		}
	})
}

// ResponseMessage is a helper to return an object with a simple message.
func ResponseMessage(message string, i ...interface{}) schema.MessageHolder {
	return schema.MessageHolder{Message: fmt.Sprintf(message, i...)}
}

// ResponseFromError is a helper to return an object with an error message.
func ResponseFromError(err error) schema.MessageHolder {
	return schema.MessageHolder{Message: err.Error()}
}
