package api

import "github.com/marsacotan/go-zabbix7/utils"

// Creating an http client is secure by default, which means TLS verification is required.
// The default value of config.InsecureSkipVerify is False, which returns an http client with TLS verification enabled.
// If you want to skip TLS verification, you need to set it to True
func NewClient(config *Config) *API {
	if !config.InsecureSkipVerify {
		httpClient := utils.CreateTLSVerifyHTTPClient(config.CertPath)
		return &API{
			Config: config,
			User:   UserAPI{Config: config, Client: httpClient},
			Token:  TokenAPI{Config: config, Client: httpClient},
		}
	}
	httpClient := utils.CreateHTTPClient()
	return &API{
		Config: config,
		User:   UserAPI{Config: config, Client: httpClient},
		Token:  TokenAPI{Config: config, Client: httpClient},
	}
}
