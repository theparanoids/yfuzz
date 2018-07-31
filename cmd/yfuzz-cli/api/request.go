// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
	"github.com/yahoo/yfuzz/pkg/schema"
)

func get(endpoint string, result interface{}) error {
	path := viper.GetString("api") + endpoint

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	return request(req, result)
}

func post(endpoint string, body interface{}, result interface{}) error {
	marshaledBody, err := json.Marshal(body)
	if err != nil {
		return nil
	}
	bodyReader := bytes.NewReader(marshaledBody)

	path := viper.GetString("api") + endpoint

	req, err := http.NewRequest("POST", path, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json")

	return request(req, result)
}

func delete(endpoint string) error {
	path := viper.GetString("api") + endpoint

	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return request(req, nil)
}

// Send a request to the API with the appropriate headers
func request(req *http.Request, res interface{}) error {
	client, err := newClient()
	if err != nil {
		return err
	}

	req.Close = true

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Handle errors
	if isError(resp.StatusCode) {
		var mh schema.MessageHolder
		var message string
		if err = json.Unmarshal(contents, &mh); err != nil {
			message = string(contents)
		} else {
			message = mh.Message
		}
		return fmt.Errorf("%d: %s", resp.StatusCode, message)
	}

	// Don't bother trying to unmarshal if we don't expect a response
	if res == nil {
		return nil
	}

	return json.Unmarshal(contents, res)
}

func isError(statusCode int) bool {
	return statusCode < http.StatusOK || statusCode >= http.StatusBadRequest
}
