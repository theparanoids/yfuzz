// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/yahoo/yfuzz/config"
)

// Helper to take a bit of the repetition out of testing endpoints
func endpointTest(path string, endpoint endpointFunc, dependencies EndpointDependencies, req *http.Request) *httptest.ResponseRecorder {
	config.InitFake()
	router := mux.NewRouter()
	router.Path(path).Handler(Endpoint(endpoint, dependencies))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
