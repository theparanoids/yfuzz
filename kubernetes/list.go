// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package kubernetes

import (
	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListJobs lists all jobs on the cluster
func (k API) ListJobs() ([]string, error) {
	jobs, err := k.client.BatchV1().Jobs(viper.GetString("kubernetes.namespace")).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var status []string

	for _, job := range jobs.Items {
		status = append(status, job.GetName())
	}

	return status, nil
}
