package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/marsacotan/go-zabbix7/types"
)

// Creates an enabled token that never expires and authenticates the specified user.
func (t *TokenAPI) CreatePermanentTokenForUser(name string, userid string) (types.TokenCreateResponse, error) {

	reqBody := types.CreatePermanentTokenForUserRequest{
		JSONRPC: "2.0",
		Method:  "token.create",
		Params: types.CreatePermanentTokenForUserParams{
			Name:   name,
			UserID: userid,
		},
		ID: 1,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return types.TokenCreateResponse{}, err
	}

	req, err := http.NewRequest("POST", t.Config.URL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return types.TokenCreateResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json-rpc")
	req.Header.Set("Authorization", "Bearer "+t.Config.AuthToken)

	resp, err := t.Client.Do(req)
	if err != nil {
		return types.TokenCreateResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return types.TokenCreateResponse{}, err
	}

	defer resp.Body.Close()

	var resBody types.TokenCreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		return types.TokenCreateResponse{}, err
	}
	return resBody, nil
}

// Creates a token with a specified expiration time that authenticates the current user.
func (c *TokenAPI) CreateUserTokenWithExpiration(name string, status string, expires_at int64) (types.TokenCreateResponse, error) {
	reqBody := types.CreateUserTokenWithExpirationRequest{
		JSONRPC: "2.0",
		Method:  "token.create",
		Params: types.CreateUserTokenWithExpirationParams{
			Name:      name,
			Status:    status,
			ExpiresAt: expires_at,
		},
		ID: 1,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return types.TokenCreateResponse{}, err
	}

	req, err := http.NewRequest("POST", c.Config.URL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return types.TokenCreateResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json-rpc")
	req.Header.Set("Authorization", "Bearer "+c.Config.AuthToken)

	resp, err := c.Client.Do(req)
	if err != nil {
		return types.TokenCreateResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return types.TokenCreateResponse{}, err
	}

	defer resp.Body.Close()

	var resBody types.TokenCreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		return types.TokenCreateResponse{}, err
	}
	return resBody, nil
}
