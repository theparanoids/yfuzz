// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"net/http"

	"github.com/gorilla/mux"
	jww "github.com/spf13/jwalterweatherman"
)

// DeleteJob removes yFuzz jobs.
// Handler for DELETE /jobs/:job
//
// URL Parameters:
// - job: the name of the job (string)
//
// Response Fields:
// - none
func DeleteJob(r *http.Request, dependencies EndpointDependencies) (int, interface{}) {
	job := mux.Vars(r)["job"]
	if job == "" {
		return http.StatusBadRequest, ResponseMessage("\"job\" is required.")
	}

	err := dependencies.Kubernetes.DeleteJob(job)
	if err != nil {
		jww.WARN.Println(err.Error())
		return http.StatusInternalServerError, ResponseFromError(err)
	}

	jww.INFO.Printf("Job %s deleted.\n", job)

	return http.StatusNoContent, nil
}
