package api

import (
	"net/http"

	"github.com/marsacotan/go-zabbix7/utils"
)

type TLSVerifyOptions struct {
	InsecureSkipVerify bool
	CertPath           string
}

type Client struct {
	HTTPClient *http.Client
	APIURL     string
	APITOKEN   string
}

func NewClient(apiURL string, opt TLSVerifyOptions, token ...string) *Client {

	var apiToken string

	if len(token) > 0 {
		apiToken = token[0]
	}

	if !opt.InsecureSkipVerify {

		return &Client{
			HTTPClient: utils.CreateTLSVerifyHTTPClient(opt.CertPath),
			APIURL:     apiURL,
			APITOKEN:   apiToken,
		}
	}

	return &Client{
		HTTPClient: utils.CreateHTTPClient(),
		APIURL:     apiURL,
		APITOKEN:   apiToken,
	}
}
