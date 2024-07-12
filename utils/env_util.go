// Copyright 2024 The tantianran Authors. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.
//
// Package utils provides utility functions for environment variables management in the go-zabbix7 library.
// The GetEnv function retrieves the value of a specified environment variable.
// If the variable is not set, it logs a fatal error and terminates the program.
//
// Functions:
// - GetEnv(envVar string) string: Retrieves the value of the specified environment variable.
//
// This utility is crucial for accessing environment configuration values required by the go-zabbix7 library.

package utils

import (
	"log"
	"os"
)

func GetEnv(envVar string) string {
	value, exists := os.LookupEnv(envVar)
	if !exists {
		log.Fatalf("Environment variable %s not set", envVar)
	}
	return value
}
