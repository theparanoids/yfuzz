// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/yahoo/yfuzz/api/schema"
)

// CreateJob creates new YFuzz jobs.
// Handler for POST /jobs
//
// Body Parameters:
//	- image: link to a docker image for a YFuzz job (string)
//
// Response Fields:
//	- job: name of the created job (string)
//
func CreateJob(r *http.Request, dependencies EndpointDependencies) (int, interface{}) {
	params := &schema.CreateJobRequest{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(params)
	if err != nil {
		return http.StatusBadRequest, responseFromError(err)
	}

	jobName, err := extractName(params.Image)
	if err != nil {
		return http.StatusBadRequest, responseFromError(err)
	}

	_, err = dependencies.Kubernetes.CreateJob(jobName, params.Image)
	if err != nil {
		// TODO: return 409 if conflict
		jww.WARN.Println(err.Error())
		return http.StatusInternalServerError, responseFromError(err)
	}

	jww.INFO.Printf("Created job %s\n", jobName)

	return http.StatusCreated, schema.CreateJobResponse{Job: jobName}
}

// extractName: takes a link to a docker registry and converts it to a name kubernetes can handle
func extractName(registryLink string) (string, error) {
	// take everything after the first forward slash following the last dot
	components := strings.Split(registryLink, ".")
	fragment := strings.SplitN(components[len(components)-1], "/", 2)
	path := fragment[len(fragment)-1]

	re := regexp.MustCompile("[^a-zA-Z0-9]")
	sanitizedPath := re.ReplaceAllString(path, "-")

	if sanitizedPath == "" {
		return "", fmt.Errorf("Could not turn %s into a job name", registryLink)
	}

	return strings.ToLower(sanitizedPath), nil
}
