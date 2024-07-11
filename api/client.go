package api

import (
	"net/http"

	"github.com/marsacotan/go-zabbix7/utils"
)

type Client struct {
	HTTPClient *http.Client
	Config     *Config
}

// Creating an http client is secure by default, which means TLS verification is required.
// The default value of config.InsecureSkipVerify is False, which returns an http client with TLS verification enabled.
// If you want to skip TLS verification, you need to set it to True
func NewZbxClient(config *Config) *Client {
	if !config.InsecureSkipVerify {
		return &Client{
			HTTPClient: utils.CreateTLSVerifyHTTPClient(config.CertPath),
			Config:     config,
		}
	}
	return &Client{
		HTTPClient: utils.CreateHTTPClient(),
		Config:     config,
	}
}
