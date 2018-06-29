// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package kubernetes

import (
	"fmt"
	"net/http"
	"strings"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	corev1 "k8s.io/api/core/v1"
)

// GetPodLogs retrieves the last tailLines lines of logs from a pod.
// If tailLines <= 0, returns the entire log.
func (k API) GetPodLogs(podName string, tailLines int) (string, error) {
	tailLines64 := int64(tailLines)

	logOptions := &corev1.PodLogOptions{}
	if tailLines > 0 {
		logOptions.TailLines = &tailLines64
	}

	res := k.client.CoreV1().Pods(viper.GetString("kubernetes.namespace")).GetLogs(podName, logOptions).Do()
	body, err := res.Raw()
	if err != nil {
		return "", err
	}

	var statusCode int
	res.StatusCode(&statusCode)
	if statusCode != http.StatusOK {
		return "", fmt.Errorf("Got %d from pod %s/logs", statusCode, podName)
	}

	return string(body), nil
}

// GetJobLogs retrieves the logs from all pods belonging to a given job.
// If crashesOnly is true, only returns logs from pods that found a crash.
// If tailLines <= 0, returns the entire log.
func (k API) GetJobLogs(job string, crashesOnly bool, tailLines int) ([]string, error) {
	pods, err := k.GetPods(job)
	if err != nil {
		return nil, err
	}

	var logs []string

	for _, pod := range pods.Items {
		// If the pod isn't finished, it hasn't found a crash
		if crashesOnly && pod.Status.Phase != corev1.PodSucceeded {
			continue
		}

		log, err := k.GetPodLogs(pod.Name, tailLines)
		if err != nil {
			jww.WARN.Println(err.Error())
			logs = append(logs, fmt.Sprintf("could not retrieve logs for pod %s", pod.Name))
			continue
		}

		// Check if it found a crash
		if crashesOnly && !strings.Contains(log, viper.GetString("kubernetes.job-config.log.crash-found-message")) {
			continue
		}

		logs = append(logs, log)
	}

	return logs, nil
}
