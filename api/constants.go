// Copyright 2024 The tantianran Authors. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.
//
// Package api defines constants and method names used in interactions with the Zabbix API.
// These constants include default HTTP request methods, JSON-RPC versioning, and specific method names.
// They are used throughout the project to maintain consistency and clarity in API interactions.
//
// Constants:
// - DefaultPost: Default HTTP request method ("POST").
// - DefaultJSONRPC: Default JSON-RPC version ("2.0").
// - DefaultContentType: Default content type for JSON-RPC requests ("application/json-rpc").
//
// Method Names:
// - UserLogin: Zabbix API method for user authentication (example).
//
// These constants and method names are central to the functionality of the API client.

package api

const (
	DefaultPost        = "POST"
	DefaultJSONRPC     = "2.0"
	DefaultContentType = "application/json-rpc"

	// All methods are defined here.
	UserLogin   = "user.login"
	UserCreate  = "user.create"
	UserDelete  = "user.delete"
	UserUnblock = "user.unblock"

	HostCreate = "host.create"
	HostDelete = "host.delete"
)
