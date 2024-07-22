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

// This is the implementation of the user.user.login method in the Zabbix API.
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
	req := utils.SendReqBody(reqBody, u.Config.URL, DefaultPost, DefaultContentType, u.Config.AuthToken)
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

	req := utils.SendReqBody(reqBody, u.Config.URL, DefaultPost, DefaultContentType, u.Config.AuthToken)
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

// This is the implementation of the user.delete method in the Zabbix API.
func (u *UserAPI) UserDelete(arrayUsersId []string) (*types.UserDeleteResult, error) {
	reqBody := types.UserDeleteRequest{
		JSONRPC: DefaultJSONRPC,
		Method:  UserDelete,
		Params:  arrayUsersId,
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

	var resBody types.UserDeleteResponse
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		return nil, err
	}

	if resBody.Error != nil {
		return nil, fmt.Errorf("delete failed: %d %v %v", resBody.Error.Code, resBody.Error.Message, resBody.Error.Data)
	}

	return resBody.Result, nil
}

// This is the implementation of the user.unblock method in the Zabbix API.
func (u *UserAPI) UserUnblock(arrayUsersId []string) (*types.UserUnblockResult, error) {
	reqBody := types.UserUnblockRequest{
		JSONRPC: DefaultJSONRPC,
		Method:  UserUnblock,
		Params:  arrayUsersId,
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

	var resBody types.UserUnblockResponse
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		return nil, err
	}

	if resBody.Error != nil {
		return nil, fmt.Errorf("unblock failed: %d %v %v", resBody.Error.Code, resBody.Error.Message, resBody.Error.Data)
	}

	return resBody.Result, nil
}
