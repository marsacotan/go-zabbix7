package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/marsacotan/go-zabbix7/types"
)

// Get a token.
func (c *Client) GetToken() (types.UserLoginTokenResponse, error) {

	reqBody := types.UserLoginRequest{
		JSONRPC: "2.0",
		Method:  "user.login",
		Params: types.UserLoginParams{
			Username: c.Config.User,
			Password: c.Config.Passwd,
		},
		ID: 1,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return types.UserLoginTokenResponse{}, err
	}

	req, err := http.NewRequest("POST", c.Config.URL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return types.UserLoginTokenResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json-rpc")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return types.UserLoginTokenResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return types.UserLoginTokenResponse{}, err
	}

	defer resp.Body.Close()

	var resBody types.UserLoginTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		return types.UserLoginTokenResponse{}, err
	}
	return resBody, nil
}

// Get other information about the authenticated user.
func (c *Client) GetUserData() (types.UserLoginUserDataResponse, error) {

	reqBody := types.UserLoginRequest{
		JSONRPC: "2.0",
		Method:  "user.login",
		Params: types.UserLoginParams{
			Username: c.Config.User,
			Password: c.Config.Passwd,
			UserData: true,
		},
		ID: 1,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return types.UserLoginUserDataResponse{}, err
	}

	req, err := http.NewRequest("POST", c.Config.URL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return types.UserLoginUserDataResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json-rpc")

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return types.UserLoginUserDataResponse{}, err
	}
	defer resp.Body.Close()

	var resBody types.UserLoginUserDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		return types.UserLoginUserDataResponse{}, err
	}
	return resBody, nil
}
