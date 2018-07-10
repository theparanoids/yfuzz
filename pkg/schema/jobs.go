// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package schema

// ListJobsResponse is the schema for the response from the GET /jobs endpoint.
// See https://godoc.org/github.com/yahoo/yfuzz/services/yfuzz-server/api#ListJobs
type ListJobsResponse struct {
	Jobs []string `json:"jobs"`
}

// CreateJobRequest is the schema for the request to the POST /jobs endpoint.
// See https://godoc.org/github.com/yahoo/yfuzz/services/yfuzz-server/api#CreateJob
type CreateJobRequest struct {
	Image string `json:"image"`
}

// CreateJobResponse is the schema for the response from the POST /jobs endpoint.
// See https://godoc.org/github.com/yahoo/yfuzz/services/yfuzz-server/api#CreateJob
type CreateJobResponse struct {
	Job string `json:"job"`
}

// GetJobStatusResponse is the schema for the respone from the GET /jobs/:job endpoint.
// See https://godoc.org/github.com/yahoo/yfuzz/services/yfuzz-server/api#GetJobStatus
type GetJobStatusResponse struct {
	StatusCrashFound int `json:"crash_found"`
	StatusNoCrash    int `json:"no_crash"`
	StatusPending    int `json:"pending"`
	StatusRunning    int `json:"running"`
	StatusFailed     int `json:"failed"`
	StatusUnknown    int `json:"unknown"`
}

// GetJobLogsResponse is the schema for the response from the GET /jobs/:job/logs endpoint.
// See https://godoc.org/github.com/yahoo/yfuzz/services/yfuzz-server/api#GetJobLogs
type GetJobLogsResponse struct {
	Logs []string `json:"logs"`
}
