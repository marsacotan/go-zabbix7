// Copyright 2024 The tantianran Authors. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.

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
