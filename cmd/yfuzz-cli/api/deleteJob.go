// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"fmt"

	"github.com/fatih/color"
)

// DeleteJob deletes the given YFuzz job.
func DeleteJob(job string) error {
	color.Green("Deleting job %s...", job)

	endpoint := fmt.Sprintf("/jobs/%s", job)

	err := delete(endpoint)
	if err != nil {
		return err
	}

	color.Green("Job %s deleted", job)

	return nil
}
