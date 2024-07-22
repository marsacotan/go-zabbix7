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

// UserLogin performs the user.login method in the Zabbix API.
// It returns either UserLoginTokenResponse or UserLoginUserDataResponse based on the userData parameter.
func (u *UserAPI) userLogin(userData *bool) (interface{}, error) {
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

	if *userData {
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

// GetToken retrieves the user login token.
func (u *UserAPI) GetToken() (types.UserLoginTokenResponse, error) {
	var flag bool = false
	resp, err := u.userLogin(&flag)
	if err != nil {
		return types.UserLoginTokenResponse{}, err
	}
	tokenResp, ok := resp.(types.UserLoginTokenResponse)
	if !ok {
		return types.UserLoginTokenResponse{}, fmt.Errorf("unexpected response type")
	}
	return tokenResp, nil
}

// GetUserData retrieves the user data.
func (u *UserAPI) GetUserData() (types.UserLoginUserDataResponse, error) {
	var flag bool = true
	resp, err := u.userLogin(&flag)
	if err != nil {
		return types.UserLoginUserDataResponse{}, err
	}
	userDataResp, ok := resp.(types.UserLoginUserDataResponse)
	if !ok {
		return types.UserLoginUserDataResponse{}, fmt.Errorf("unexpected response type")
	}
	return userDataResp, nil
}

// This is the implementation of the user.create method in the Zabbix API.
func (u *UserAPI) UserCreate(username string, options ...func(*types.UserCreateRequest)) (*types.UserCreateResult, error) {
	reqBody := types.UserCreateRequest{
		JSONRPC: DefaultJSONRPC,
		Method:  UserCreate,
		Params: types.UserCreateParams{
			Username: username,
		},
		ID: 1,
	}

	for _, opt := range options {
		opt(&reqBody)
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
	req.Header.Set("Authorization", "Bearer "+u.Config.AuthToken)

	resp, err := u.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var resBody types.UserCreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		return nil, err
	}

	if resBody.Error != nil {
		return nil, fmt.Errorf("create failed: %d %v %v", resBody.Error.Code, resBody.Error.Message, resBody.Error.Data)
	}

	return resBody.Result, nil
}
