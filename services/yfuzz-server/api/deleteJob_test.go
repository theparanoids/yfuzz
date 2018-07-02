// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.
package api

import (
	"net/http"
	"testing"

	"github.com/yahoo/yfuzz/services/yfuzz-server/kubernetes"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDeleteJob(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/jobs/test-job", nil)

	fakeKube := kubernetes.NewFake(&batchv1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind: "job",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-job",
		},
	})

	w := endpointTest("/jobs/{job}", DeleteJob, EndpointDependencies{
		Kubernetes: fakeKube,
	}, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("List jobs endpoint returned %v instead of %v", w.Code, http.StatusNoContent)
	}
}
