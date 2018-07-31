// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"net/http"

	"github.com/gorilla/mux"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/yahoo/yfuzz/pkg/schema"
	"github.com/yahoo/yfuzz/pkg/types"
)

// GetJobStatus retrieves the status of all pods for a given job.
// Handler for GET /jobs/:job
//
// URL Parameters:
// - job: the name of the job (string)
//
// Response Fields:
// - crash_found: number of pods that found a crash (int)
// - no_crash: number of pods that timed out or hit the input limit with no crash (int)
// - pending: number of pods waiting to be scheduled (int)
// - running: number of pods currently running (int)
// - failed: number of pods that experienced an internal error (int)
// - unknown: number of pods that cannot be contacted (int)
func GetJobStatus(r *http.Request, dependencies EndpointDependencies) (int, interface{}) {
	job := mux.Vars(r)["job"]
	if job == "" {
		return http.StatusBadRequest, ResponseMessage("\"job\" is required.")
	}

	status, err := dependencies.Kubernetes.GetJobStatus(job)
	if err != nil {
		jww.WARN.Println(err.Error())
		return http.StatusInternalServerError, ResponseFromError(err)
	}

	return http.StatusOK, schema.GetJobStatusResponse{
		StatusCrashFound: status[types.StatusCrashFound],
		StatusNoCrash:    status[types.StatusNoCrash],
		StatusPending:    status[types.StatusPending],
		StatusRunning:    status[types.StatusRunning],
		StatusFailed:     status[types.StatusFailed],
		StatusUnknown:    status[types.StatusUnknown],
	}
}
