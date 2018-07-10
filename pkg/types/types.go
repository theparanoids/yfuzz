// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package types

// PodStatus is an enum of possible yFuzz pod statuses.
type PodStatus string

// Available yFuzz pod statuses:
// CrashFound: the pod found a crash.
// NoCrash: the pod exited (e.g. due to timeout) without finding a crash.
// Pending: the pod is waiting for an available node.
// Running: the pod is in the process of fuzzing.
// Failed: a problem occurred with the pod.
// Unknown: the pod cannot be reached to check its status.
const (
	StatusCrashFound PodStatus = "CrashFound"
	StatusNoCrash    PodStatus = "NoCrash"
	StatusPending    PodStatus = "Pending"
	StatusRunning    PodStatus = "Running"
	StatusFailed     PodStatus = "Failed"
	StatusUnknown    PodStatus = "Unknown"
)
