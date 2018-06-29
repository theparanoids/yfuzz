// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package kubernetes

import (
	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeleteJob deletes the Kubernetes job with the given name
func (k API) DeleteJob(name string) error {
	propagationPolicy := metav1.DeletePropagationForeground

	return k.client.BatchV1().Jobs(viper.GetString("kubernetes.namespace")).Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &propagationPolicy,
	})
}
