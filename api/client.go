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
		Token:  TokenAPI{Config: config, Client: httpClient},
	}
}
