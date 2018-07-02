// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package schema

// ListJobsResponse ...
// Schema for GET /jobs response
type ListJobsResponse struct {
	Jobs []string `json:"jobs"`
}

// CreateJobRequest ...
// Schema for POST /jobs request
type CreateJobRequest struct {
	Image string `json:"image"`
}

// CreateJobResponse ...
// Schema for POST /jobs response
type CreateJobResponse struct {
	Job string `json:"job"`
}

// GetJobStatusResponse ...
// Schema for GET /jobs/:job response
type GetJobStatusResponse struct {
	StatusCrashFound int `json:"crash_found"`
	StatusNoCrash    int `json:"no_crash"`
	StatusPending    int `json:"pending"`
	StatusRunning    int `json:"running"`
	StatusFailed     int `json:"failed"`
	StatusUnknown    int `json:"unknown"`
}

// GetJobLogsResponse ...
// Schema for GET /jobs/:job/logs response
type GetJobLogsResponse struct {
	Logs []string `json:"logs"`
}
