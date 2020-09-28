// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

import (
	"encoding/json"
	"fmt"
	"time"
)

type DatastoreOperationWaiter struct {
	Config    *Config
	UserAgent string
	Project   string
	CommonOperationWaiter
}

func (w *DatastoreOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("https://datastore.googleapis.com/v1/%s", w.CommonOperationWaiter.Op.Name)

	return sendRequest(w.Config, "GET", w.Project, url, w.UserAgent, nil, datastoreIndex409Contention)
}

func createDatastoreWaiter(config *Config, op map[string]interface{}, project, activity, userAgent string) (*DatastoreOperationWaiter, error) {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil, nil
	}
	w := &DatastoreOperationWaiter{
		Config:    config,
		UserAgent: userAgent,
		Project:   project,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return nil, err
	}
	return w, nil
}

// nolint: deadcode,unused
func datastoreOperationWaitTimeWithResponse(config *Config, op map[string]interface{}, response *map[string]interface{}, project, activity, userAgent string, timeout time.Duration) error {
	w, err := createDatastoreWaiter(config, op, project, activity, userAgent)
	if err != nil || w == nil {
		// If w is nil, the op was synchronous.
		return err
	}
	if err := OperationWait(w, activity, timeout, config.PollInterval); err != nil {
		return err
	}
	return json.Unmarshal([]byte(w.CommonOperationWaiter.Op.Response), response)
}

func datastoreOperationWaitTime(config *Config, op map[string]interface{}, project, activity, userAgent string, timeout time.Duration) error {
	w, err := createDatastoreWaiter(config, op, project, activity, userAgent)
	if err != nil || w == nil {
		// If w is nil, the op was synchronous.
		return err
	}
	return OperationWait(w, activity, timeout, config.PollInterval)
}
