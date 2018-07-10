// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/yahoo/yfuzz/pkg/schema"
)

// GetJobLogs retrieves the logs of a given yFuzz job, truncated to a given number of lines.
// See https://godoc.org/github.com/yahoo/yfuzz/services/yfuzz-server/api#GetJobLogs
func GetJobLogs(job string, tail int) error {
	color.Green("Getting logs for job %s...", job)

	endpoint := fmt.Sprintf("/jobs/%s/logs?tail=%d", job, tail)

	var logs schema.GetJobLogsResponse

	err := get(endpoint, &logs)
	if err != nil {
		return err
	}

	if len(logs.Logs) == 0 {
		color.Red("No logs found for job %s", job)
		return nil
	}

	if len(logs.Logs) == 1 {
		fmt.Println(logs.Logs[0])
	} else {
		for podNum, log := range logs.Logs {
			color.Green("Log %d:", podNum)
			fmt.Println(log)
		}
	}

	return nil
}
