// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package kubernetes

import (
	"strings"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"github.com/yahoo/yfuzz/types"
	corev1 "k8s.io/api/core/v1"
)

// GetJobStatus retrieves the status of all pods belonging to a given job.
func (k API) GetJobStatus(job string) (map[types.PodStatus]int, error) {
	pods, err := k.GetPods(job)
	if err != nil {
		return nil, err
	}

	count := make(map[types.PodStatus]int)

	for _, pod := range pods.Items {
		if pod.Status.Phase == corev1.PodSucceeded {
			// Job finished normally, check if we found a crash
			status := k.getFinishedPodStatus(pod.GetName())
			count[status]++
		} else {
			status := podPhaseToPodStatus(pod.Status.Phase)
			count[status]++
		}
	}

	return count, nil
}

func (k API) getFinishedPodStatus(podName string) types.PodStatus {
	log, err := k.GetPodLogs(podName, viper.GetInt("kubernetes.job-config.log.tail-lines"))
	if err != nil {
		jww.WARN.Println(err.Error())
		return types.StatusUnknown
	}

	if strings.Contains(log, viper.GetString("kubernetes.job-config.log.crash-found-message")) {
		return types.StatusCrashFound
	}

	return types.StatusNoCrash
}

// Convert a kubernetes pod phase into a YFuzz pod status
func podPhaseToPodStatus(podPhase corev1.PodPhase) types.PodStatus {
	switch podPhase {
	case corev1.PodPending:
		return types.StatusPending
	case corev1.PodRunning:
		return types.StatusRunning
	case corev1.PodFailed:
		return types.StatusFailed
	case corev1.PodUnknown:
		return types.StatusUnknown
	default:
		jww.WARN.Printf("Faield to convert %s to YFuzz Job Status.\n", podPhase)
		return types.StatusUnknown
	}
}
