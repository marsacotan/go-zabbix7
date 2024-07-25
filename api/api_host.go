// Copyright 2024 The tantianran Authors. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.
//
// The api package provides methods to interact with Zabbix API functionalities.
// These methods use the JSON-RPC protocol to facilitate authentication, management,
// and operations of user entities within the Zabbix system.

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marsacotan/go-zabbix7/types"
	"github.com/marsacotan/go-zabbix7/utils"
)

func (u *HostAPI) HostCreate(hostname string, options ...func(*types.HostCreateRequest)) (*types.HostCreateResult, error) {
	reqBody := types.HostCreateRequest{
		JSONRPC: DefaultJSONRPC,
		Method:  HostCreate,
		Params: types.HostCreateParams{
			Host: hostname,
		},
		ID: 1,
	}

	for _, opt := range options {
		opt(&reqBody)
	}

	req := utils.SendReqBody(reqBody, u.Config.URL, DefaultPost, DefaultContentType, u.Config.AuthToken)
	resp, err := u.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var resBody types.HostCreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		return nil, err
	}

	if resBody.Error != nil {
		return nil, fmt.Errorf("create failed: %d %v %v", resBody.Error.Code, resBody.Error.Message, resBody.Error.Data)
	}

	return resBody.Result, nil
}

func (u *HostAPI) HostDelete(hostid []string) (*types.HostDeleteResult, error) {
	reqBody := types.HostDeleteRequest{
		JSONRPC: DefaultJSONRPC,
		Method:  HostDelete,
		Params:  hostid,
		ID:      1,
	}

	req := utils.SendReqBody(reqBody, u.Config.URL, DefaultPost, DefaultContentType, u.Config.AuthToken)
	resp, err := u.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var resBody types.HostDeleteResponse
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		return nil, err
	}

	if resBody.Error != nil {
		return nil, fmt.Errorf("delete failed: %d %v %v", resBody.Error.Code, resBody.Error.Message, resBody.Error.Data)
	}

	return resBody.Result, nil
}
