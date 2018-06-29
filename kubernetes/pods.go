// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package kubernetes

import (
	"github.com/spf13/viper"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetPods retrieves all pods belonging to a job
func (k API) GetPods(job string) (*corev1.PodList, error) {
	label := "job-name = " + job

	return k.client.CoreV1().Pods(viper.GetString("kubernetes.namespace")).List(metav1.ListOptions{
		LabelSelector: label,
	})
}
