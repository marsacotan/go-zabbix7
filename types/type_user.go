// Copyright 2024 The tantianran Authors. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.
//
// Package api defines structs for request and response bodies used in interactions with the Zabbix API.
// This file specifically defines structs related to user-related API methods.
// Each struct corresponds to a specific API method's request or response format, ensuring type safety
// and clarity in data exchange with the Zabbix server.
//
// The structs defined here are integral to encoding and decoding JSON-RPC messages sent and received
// by the API client.

package types

// Here are the struct definitions for the request and response of the user.login method.
type UserLoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserData bool   `json:"userData,omitempty"`
}

type UserLoginRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  UserLoginParams `json:"params"`
	ID      int             `json:"id"`
}

type UserLoginTokenResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      int    `json:"id"`
}

type UserLoginUserDataResult struct {
	UserID          string `json:"userid"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	URL             string `json:"url"`
	AutoLogin       string `json:"autologin"`
	AutoLogout      string `json:"autologout"`
	Lang            string `json:"lang"`
	Refresh         string `json:"refresh"`
	Theme           string `json:"theme"`
	AttemptFailed   string `json:"attempt_failed"`
	AttemptIP       string `json:"attempt_ip"`
	AttemptClock    string `json:"attempt_clock"`
	RowsPerPage     string `json:"rows_per_page"`
	Timezone        string `json:"timezone"`
	RoleID          string `json:"roleid"`
	UserDirectoryID string `json:"userdirectoryid"`
	TSProvisioned   string `json:"ts_provisioned"`
	DebugMode       int    `json:"debug_mode"`
	Deprovisioned   bool   `json:"deprovisioned"`
	GUIAccess       string `json:"gui_access"`
	MFAID           int    `json:"mfaid"`
	AuthType        int    `json:"auth_type"`
	Type            int    `json:"type"`
	UserIP          string `json:"userip"`
	SessionID       string `json:"sessionid"`
	Secret          string `json:"secret"`
}

type UserLoginUserDataResponse struct {
	JSONRPC string                  `json:"jsonrpc"`
	Result  UserLoginUserDataResult `json:"result"`
	ID      int                     `json:"id"`
}
