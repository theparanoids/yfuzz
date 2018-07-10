// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"github.com/fatih/color"
	"github.com/yahoo/yfuzz/pkg/schema"
)

// CreateJob creates a yFuzz job from a link to a docker image.
// See https://godoc.org/github.com/yahoo/yfuzz/services/yfuzz-server/api#CreateJob
func CreateJob(image string) error {
	color.Green("Creating job from image %s...", image)

	var created schema.CreateJobResponse

	params := schema.CreateJobRequest{
		Image: image,
	}

	err := post("/jobs", params, &created)
	if err != nil {
		return err
	}

	color.Green("Job %s created", created.Job)

	return nil
}
