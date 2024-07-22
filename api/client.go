// Copyright 2024 The tantianran Authors. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.
//
// Package api provides utilities for creating a Zabbix API client instance.
//
// The NewClient function initializes a new API client with the provided configuration.
// It configures the HTTP client based on security settings and returns an API instance
// configured to interact with the Zabbix API server.

package api

import "github.com/marsacotan/go-zabbix7/utils"

func NewClient(config *Config) *API {
	httpClient := utils.CreateHTTPClient()
	if !config.InsecureSkipVerify {
		httpClient = utils.CreateTLSVerifyHTTPClient(config.CertPath)
	}

	return &API{
		Config: config,
		User:   UserAPI{Config: config, Client: httpClient},
		Host:   HostAPI{Config: config, Client: httpClient},
	}
}
