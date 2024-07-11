package api

import (
	"bytes"
	"encoding/json"
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

func Post(c *Client, reqBody interface{}) (*http.Response, error) {
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.Config.URL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json-rpc")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return resp, nil
}
