// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/yahoo/yfuzz/pkg/schema"
)

// GetJobStatus retrieves the status of a yFuzz job.
// See https://godoc.org/github.com/yahoo/yfuzz/services/yfuzz-server/api#GetJobStatus
func GetJobStatus(job string) error {
	color.Green("Getting job %s...", job)

	endpoint := fmt.Sprintf("/jobs/%s", job)

	var status schema.GetJobStatusResponse

	err := get(endpoint, &status)
	if err != nil {
		return err
	}

	noneFound := true

	if status.StatusCrashFound > 0 {
		color.Green("%d found crash", status.StatusCrashFound)
		noneFound = false
	}

	if status.StatusNoCrash > 0 {
		color.Green("%d finished without finding a crash", status.StatusNoCrash)
		noneFound = false
	}

	if status.StatusRunning > 0 {
		color.Green("%d running", status.StatusRunning)
		noneFound = false
	}

	if status.StatusPending > 0 {
		color.Green("%d queued", status.StatusPending)
		noneFound = false
	}

	if status.StatusFailed > 0 {
		color.Red("%d failed", status.StatusFailed)
		noneFound = false
	}

	if status.StatusUnknown > 0 {
		color.Red("%d unknown", status.StatusUnknown)
		noneFound = false
	}

	if noneFound {
		color.Red("No events found for job %s", job)
	}

	return nil
}
