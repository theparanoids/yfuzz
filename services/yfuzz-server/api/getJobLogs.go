// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/yahoo/yfuzz/pkg/schema"
)

// GetJobLogs retrieves the logs for a specific job.
// Handler for GET /jobs/:job/logs
//
// URL Parameters:
//	- job: the name of the job (string)
//
// Query Parameters:
//  - crashes_only: only show crashes (boolean)
//  - tail: number of tail lines to truncate the log to
//
// Response Fields:
// 	- logs: logs from jobs
func GetJobLogs(r *http.Request, dependencies EndpointDependencies) (int, interface{}) {
	job := mux.Vars(r)["job"]
	if job == "" {
		return http.StatusBadRequest, responseMessage("\"job\" is required.")
	}

	crashesOnly := r.URL.Query().Get("crashes_only") == "true"
	tailLinesString := r.URL.Query().Get("tail")
	tailLines, err := strconv.Atoi(tailLinesString)

	if tailLinesString != "" && err != nil {
		return http.StatusBadRequest, responseMessage("\"tail\" must be an integer")
	}

	logs, err := dependencies.Kubernetes.GetJobLogs(job, crashesOnly, tailLines)
	if err != nil {
		jww.WARN.Println(err.Error())
		return http.StatusInternalServerError, responseFromError(err)
	}

	return http.StatusOK, schema.GetJobLogsResponse{Logs: logs}
}
