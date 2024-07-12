// Copyright 2024 The tantianran Authors. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.
//
// Package api provides a structured interface for interacting with the Zabbix API,
// focusing primarily on user-related operations.
//
// The API struct manages the configuration and client instances for making HTTP requests
// to the Zabbix API server.
//
// Example:
// The UserAPI struct encapsulates user-related methods such as authentication,
// management, and profile retrieval, following a similar structure to other components
// within this package.

package api

import (
	"net/http"
)

type API struct {
	Config *Config
	User   UserAPI
}

type UserAPI struct {
	Config *Config
	Client *http.Client
}
