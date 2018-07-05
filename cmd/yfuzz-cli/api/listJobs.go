// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"github.com/fatih/color"
	"github.com/yahoo/yfuzz/pkg/schema"
)

// ListJobs lists all of the existing YFuzz jobs
func ListJobs() error {
	color.Green("Listing all jobs...")

	var jobList schema.ListJobsResponse

	err := get("/jobs", &jobList)
	if err != nil {
		return err
	}

	for _, job := range jobList.Jobs {
		color.Green(job)
	}

	return nil
}
