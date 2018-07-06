// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package kubernetes

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/spf13/viper"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

var jobTemplateHead = `
apiVersion: batch/v1
kind: Job
metadata:
  name: {{ .Name }}
spec:
  parallelism: {{ .JobConfig.GetString "parallelism" }}
  template:
    metadata:
      name: {{ .Name }}
    spec:
      containers:
        - name: yfuzz
          image: {{ .Image }}
          resources:
            requests:
              memory: {{ .JobConfig.GetString "memory.request" }}
              cpu: {{ .JobConfig.GetString "cpu.request" }}
            limits:
              memory: {{ .JobConfig.GetString "memory.limit" }}
              cpu: {{ .JobConfig.GetString "cpu.limit" }}
          command: ["/bin/sh"]
          args: ["-c", "/run_fuzzer.sh"]
          volumeMounts:
            - name: yfuzz-volume
              mountPath: /shared_data
            - name: varlog
              mountPath: /var/log/yfuzz`

var logContainerTemplate = `
        - name: logging
          image: {{ .JobConfig.GetString("log.image") }}
          volumeMounts:
            - name: "varlog"
              mountPath: "/var/log/yfuzz"
          envFrom:
            - configMapRef:
                name: {{ .JobConfig.GetString("log.config-map") }}
            - secretRef:
                name: {{ .JobConfig.GetString("log.secret") }}`

var jobTemplateTail = `
      restartPolicy: Never
      volumes:
        - name: yfuzz-volume
          persistentVolumeClaim:
            claimName: {{ .JobConfig.GetString "persistent-volume-claim" }}
        - name: varlog
          emptyDir: {}`

type options struct {
	JobConfig *viper.Viper
	Name      string
	Image     string
}

// CreateJob creates a new job in the Kubernetes cluster.
func (k API) CreateJob(name, registryLink string) (*batchv1.Job, error) {
	var jobTemplateString string

	if viper.IsSet("kubernetes.job-config.log.image") {
		jobTemplateString = jobTemplateHead + logContainerTemplate + jobTemplateTail
	} else {
		jobTemplateString = jobTemplateHead + jobTemplateTail
	}

	jobTemplate, err := template.New("jobTemplate").Parse(jobTemplateString)
	if err != nil {
		return nil, err
	}

	opts := options{
		JobConfig: viper.Sub("kubernetes.job-config"),
		Name:      name,
		Image:     registryLink,
	}

	var rawJob bytes.Buffer
	err = jobTemplate.Execute(&rawJob, opts)
	if err != nil {
		return nil, err
	}

	decode := scheme.Codecs.UniversalDeserializer().Decode

	obj, _, err := decode(rawJob.Bytes(), nil, nil)
	if err != nil {
		return nil, err
	}

	jobSpec, ok := obj.(*batchv1.Job)
	if !ok {
		return nil, errors.New("could not parse job template")
	}

	return k.client.BatchV1().Jobs(viper.GetString("kubernetes.namespace")).Create(jobSpec)
}
