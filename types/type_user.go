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
	UserData *bool  `json:"userData,omitempty"`
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

// Here are the struct definitions for the request and response of the user.create method.
type CreateUserParams struct {
	UserID          *int            `json:"userid,omitempty"`
	Username        *string         `json:"username,omitempty"`
	Password        *string         `json:"passwd,omitempty"`
	RoleID          *int            `json:"roleid,omitempty"`
	AttemptClock    *int64          `json:"attempt_clock,omitempty"`
	AttemptFailed   *int            `json:"attempt_failed,omitempty"`
	AttemptIP       *string         `json:"attempt_ip,omitempty"`
	Autologin       *int            `json:"autologin,omitempty"`
	Autologout      *string         `json:"autologout,omitempty"`
	Lang            *string         `json:"lang,omitempty"`
	Name            *string         `json:"name,omitempty"`
	Refresh         *string         `json:"refresh,omitempty"`
	RowsPerPage     *int            `json:"rows_per_page,omitempty"`
	Surname         *string         `json:"surname,omitempty"`
	Theme           *string         `json:"theme,omitempty"`
	TSProvisioned   *int64          `json:"ts_provisioned,omitempty"`
	URL             *string         `json:"url,omitempty"`
	UserDirectoryID *int            `json:"userdirectoryid,omitempty"`
	UsrGrps         []*UsrGrpParams `json:"usrgrps,omitempty"`
	Medias          []*MediaParams  `json:"medias,omitempty"`
}

type UserCreateRequest struct {
	JSONRPC string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  CreateUserParams `json:"params"`
	ID      int              `json:"id"`
}

type UsrGrpParams struct {
	UsrGrpID        *int    `json:"usrgrpid,omitempty"`
	Name            *string `json:"name,omitempty"`
	DebugMode       *int    `json:"debug_mode,omitempty"`
	GUIAccess       *int    `json:"gui_access,omitempty"`
	MFAStatus       *int    `json:"mfa_status,omitempty"`
	MFAID           *int    `json:"mfaid,omitempty"`
	UsersStatus     *int    `json:"users_status,omitempty"`
	UserDirectoryID *int    `json:"userdirectoryid,omitempty"`
}

type MediaParams struct {
	MediaTypeID          *int      `json:"mediatypeid,omitempty"`
	SendTo               *[]string `json:"sendto,omitempty"`
	Active               *int      `json:"active,omitempty"`
	Severity             *int      `json:"severity,omitempty"`
	Period               *string   `json:"period,omitempty"`
	UserDirectoryMediaID *int      `json:"userdirectory_mediaid,omitempty"`
}
