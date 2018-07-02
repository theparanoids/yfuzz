// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/yahoo/yfuzz/pkg/schema"
	"github.com/yahoo/yfuzz/services/yfuzz-server/kubernetes"
)

func TestCreateJob(t *testing.T) {
	params, err := json.Marshal(schema.CreateJobRequest{
		Image: "test-image",
	})
	if err != nil {
		t.Errorf("Could not marshal test payload for create job, %s", err.Error())
	}

	req, _ := http.NewRequest("POST", "/jobs", bytes.NewReader(params))

	fakeKube := kubernetes.NewFake()

	w := endpointTest("/jobs", CreateJob, EndpointDependencies{
		Kubernetes: fakeKube,
	}, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Create job returned %v instead of %v", w.Code, http.StatusCreated)
	}

	var response schema.CreateJobResponse

	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Could not unmarshal create job response, got error: %s", err.Error())
	}

	if response.Job != "test-image" {
		t.Errorf("Create job endpoint didn't return the created job. Got %s", response.Job)
	}
}
