// Copyright 2024 The tantianran Authors. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.
//
// Package api provides methods for interacting with user-related functionalities
// in the Zabbix API.
//
// The main functionalities encapsulated in this file include:
// - user.checkAuthentication
// - user.create
// - user.delete
// - user.get
// - user.login
// - user.logout
// - user.provision
// - user.resettotp
// - user.unblock
// - user.update
//
// These methods facilitate authentication, management, and manipulation of user
// entities within a Zabbix system using the JSON-RPC protocol.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marsacotan/go-zabbix7/types"
)

// This is the implementation of the user.login method in the Zabbix API.
func (u *UserAPI) GetToken() (types.UserLoginTokenResponse, error) {
	resp, err := u.makeUserLoginRequest(false)
	if err != nil {
		return types.UserLoginTokenResponse{}, err
	}
	return resp.(types.UserLoginTokenResponse), nil
}

func (u *UserAPI) GetUserData() (types.UserLoginUserDataResponse, error) {
	resp, err := u.makeUserLoginRequest(true)
	if err != nil {
		return types.UserLoginUserDataResponse{}, err
	}
	return resp.(types.UserLoginUserDataResponse), nil
}

func (u *UserAPI) makeUserLoginRequest(userData bool) (interface{}, error) {
	reqBody := types.UserLoginRequest{
		JSONRPC: DefaultJSONRPC,
		Method:  UserLogin,
		Params: types.UserLoginParams{
			Username: u.Config.User,
			Password: u.Config.Passwd,
			UserData: userData,
		},
		ID: 1,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(DefaultPost, u.Config.URL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", DefaultContentType)

	resp, err := u.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if userData {
		var resBody types.UserLoginUserDataResponse
		if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
			return nil, err
		}
		return resBody, nil
	} else {
		var resBody types.UserLoginTokenResponse
		if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
			return nil, err
		}
		return resBody, nil
	}
}

// This is the implementation of the user.create method in the Zabbix API.
