// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package kubernetes

import (
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// API holds an interface for talking to Kubernetes.
type API struct {
	client kubernetes.Interface
}

// New creates a client for interacting with kubernetes.
func New() (*API, error) {
	var kubeConfig *rest.Config
	var err error

	if viper.IsSet("kubernetes.config-path") {
		kubeConfig, err = clientcmd.BuildConfigFromFlags("", viper.GetString("kubernetes.config-path"))
	} else {
		kubeConfig, err = rest.InClusterConfig()
	}
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}

	return &API{
		client: clientset,
	}, nil
}

// NewFake creates a fake Kubernetes client for running tests.
func NewFake(objects ...runtime.Object) *API {
	return &API{
		client: fake.NewSimpleClientset(objects...),
	}
}
