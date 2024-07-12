package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marsacotan/go-zabbix7/types"
)

// Get a token.
func (u *UserAPI) GetToken() (types.UserLoginTokenResponse, error) {
	resp, err := u.makeRequest(false)
	if err != nil {
		return types.UserLoginTokenResponse{}, err
	}
	return resp.(types.UserLoginTokenResponse), nil
}

// Get other information about the authenticated user.
func (u *UserAPI) GetUserData() (types.UserLoginUserDataResponse, error) {
	resp, err := u.makeRequest(true)
	if err != nil {
		return types.UserLoginUserDataResponse{}, err
	}
	return resp.(types.UserLoginUserDataResponse), nil
}

// makeRequest makes a request based on userData parameter.
func (u *UserAPI) makeRequest(userData bool) (interface{}, error) {
	reqBody := types.UserLoginRequest{
		JSONRPC: "2.0",
		Method:  "user.login",
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

	req, err := http.NewRequest("POST", u.Config.URL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json-rpc")

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
