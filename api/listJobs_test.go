// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/yahoo/yfuzz/api/schema"
	"github.com/yahoo/yfuzz/kubernetes"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestListJobs(t *testing.T) {
	req, _ := http.NewRequest("GET", "/jobs", nil)

	fakeKube := kubernetes.NewFake(&batchv1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind: "job",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-job",
		},
	})

	w := endpointTest("/jobs", ListJobs, EndpointDependencies{
		Kubernetes: fakeKube,
	}, req)

	if w.Code != http.StatusOK {
		t.Errorf("List jobs endpoint returned %v instead of %v", w.Code, http.StatusOK)
	}

	var response schema.ListJobsResponse

	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Could not unmarshal jobs response, got error: %s", err.Error())
	}

	if len(response.Jobs) != 1 || response.Jobs[0] != "test-job" {
		t.Errorf("List jobs endpoint didn't return the test job.")
	}
}
