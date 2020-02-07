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
)

type RedisOperationWaiter struct {
	Config  *Config
	Project string
	CommonOperationWaiter
}

func (w *RedisOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("https://redis.googleapis.com/v1/%s", w.CommonOperationWaiter.Op.Name)
	return sendRequest(w.Config, "GET", w.Project, url, nil)
}

func createRedisWaiter(config *Config, op map[string]interface{}, project, activity string) (*RedisOperationWaiter, error) {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil, nil
	}
	w := &RedisOperationWaiter{
		Config:  config,
		Project: project,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return nil, err
	}
	return w, nil
}

func redisOperationWaitTimeWithResponse(config *Config, op map[string]interface{}, response *map[string]interface{}, project, activity string, timeoutMinutes int) error {
	w, err := createRedisWaiter(config, op, project, activity)
	if err != nil || w == nil {
		// If w is nil, the op was synchronous.
		return err
	}
	if err := OperationWait(w, activity, timeoutMinutes); err != nil {
		return err
	}
	return json.Unmarshal([]byte(w.CommonOperationWaiter.Op.Response), response)
}

func redisOperationWaitTime(config *Config, op map[string]interface{}, project, activity string, timeoutMinutes int) error {
	w, err := createRedisWaiter(config, op, project, activity)
	if err != nil || w == nil {
		// If w is nil, the op was synchronous.
		return err
	}
	return OperationWait(w, activity, timeoutMinutes)
}
