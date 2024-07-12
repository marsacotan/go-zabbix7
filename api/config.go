// Copyright 2024 The tantianran Authors. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.
//
// Package api provides configuration utilities for setting up connections to the Zabbix API.
//
// The Config struct defines parameters such as URL, authentication credentials, and security settings
// for connecting to the Zabbix API server.
//
// The ConfigConn function initializes a new Config instance with the specified URL.
//
// The WithLoginCred method sets the username and password for authentication in the Config instance.
//
// The WithToken method sets the authentication token in the Config instance.
//
// The WithSkipTlsVerify method configures whether to skip TLS certificate verification in the Config instance.
//
// The WithCertPath method sets the path to the TLS certificate in the Config instance.
//
// The WithLocalEnvToken function retrieves a token from the local environment using the provided environment variable.

package api

import "github.com/marsacotan/go-zabbix7/utils"

type Config struct {
	URL                string
	User               string
	Passwd             string
	AuthToken          string
	InsecureSkipVerify bool
	CertPath           string
}

func ConfigConn(url string) *Config {
	return &Config{
		URL: url,
	}
}

func (c *Config) WithLoginCred(user, passwd string) *Config {
	c.User = user
	c.Passwd = passwd
	return c
}

func (c *Config) WithToken(token string) *Config {
	c.AuthToken = token
	return c
}

func (c *Config) WithSkipTlsVerify(insecureSkipVerify bool) *Config {
	c.InsecureSkipVerify = insecureSkipVerify
	return c
}

func (c *Config) WithCertPath(certPath string) *Config {
	c.CertPath = certPath
	return c
}

func WithLocalEnvToken(envVar string) string {
	return utils.GetEnv(envVar)
}
