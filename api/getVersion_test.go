// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/yahoo/yfuzz/api/schema"
	"github.com/yahoo/yfuzz/config"
)

func TestGetVersion(t *testing.T) {
	req, _ := http.NewRequest("GET", "/version", nil)
	w := endpointTest("/version", GetVersion, EndpointDependencies{}, req)

	if w.Code != http.StatusOK {
		t.Errorf("Version returned %v instead of %v", w.Code, http.StatusOK)
	}

	var response schema.GetVersionResponse

	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Could not unmarshal version response, got error: %s", err.Error())
	}

	if response.Version != config.Version {
		t.Errorf("Version endpoint returned %s instead of %s", response.Version, config.Version)
	}
}
