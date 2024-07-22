// Copyright 2024 The tantianran Authors. All rights reserved.
// Use of this source code is governed by an Apache-2.0 license
// that can be found in the LICENSE file.
//
// Package utils provides utility functions for HTTP client configuration in the go-zabbix7 library.
// These functions facilitate the creation of HTTP clients with different TLS configurations,
// supporting both insecure skip verification and verified TLS connections using custom CA certificates.
//
// Functions:
// - CreateHTTPClient(): Creates an HTTP client with insecure TLS verification.
// - CreateTLSVerifyHTTPClient(caCertPath string): Creates an HTTP client with verified TLS using a specified CA certificate file.
//
// These utilities are essential for establishing secure and configurable HTTP connections
// required by the Zabbix API client.

package utils

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func CreateHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func CreateTLSVerifyHTTPClient(caCertPath string) *http.Client {

	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("TLS authentication is required by default, and cannot read the CA certificate file: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            caCertPool,
		},
	}

	return &http.Client{Transport: tr}
}

func SendReqBody(reqBody interface{}, url string, reqMethod string, contentType string, token ...string) *http.Request {
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(reqMethod, url, bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Fatalln(err)
	}

	if len(token) != 0 {
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Authorization", "Bearer "+token[0])
	} else {
		req.Header.Set("Content-Type", contentType)
	}

	return req
}
