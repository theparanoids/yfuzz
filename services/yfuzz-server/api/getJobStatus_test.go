// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/yahoo/yfuzz/pkg/schema"
	"github.com/yahoo/yfuzz/services/yfuzz-server/kubernetes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetJobStatus(t *testing.T) {
	req, _ := http.NewRequest("GET", "/jobs/test-job", nil)

	fakeKube := kubernetes.NewFake(&corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind: "pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   "test-job-pod-1",
			Labels: map[string]string{"job-name": "test-job"},
		},
		Status: corev1.PodStatus{
			Phase: corev1.PodPending,
		},
	})

	w := endpointTest("/jobs/{job}", GetJobStatus, EndpointDependencies{
		Kubernetes: fakeKube,
	}, req)

	if w.Code != http.StatusOK {
		t.Errorf("Job status returned %v instead of %v", w.Code, http.StatusOK)
	}

	var response schema.GetJobStatusResponse

	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Could not unmarshal job status response, got error: %s", err.Error())
	}

	if response.StatusPending != 1 {
		t.Errorf("Did not report pending pod.")
	}
}
