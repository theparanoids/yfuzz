// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"net/http"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/yahoo/yfuzz/api/schema"
)

// ListJobs returns a list of all existing jobs on the cluster.
// Handler for GET /jobs
//
// Parameters:
//	- none
//
// Response Fields:
//	- jobs: all YFuzz jobs (array of strings)
func ListJobs(r *http.Request, dependencies EndpointDependencies) (int, interface{}) {
	jobs, err := dependencies.Kubernetes.ListJobs()
	if err != nil {
		jww.WARN.Println(err.Error())
		return http.StatusInternalServerError, responseFromError(err)
	}

	return http.StatusOK, schema.ListJobsResponse{Jobs: jobs}
}
