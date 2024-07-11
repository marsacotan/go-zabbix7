package api

import "github.com/marsacotan/go-zabbix7/utils"

type ClientConfig struct {
	URL                string
	User               string
	Passwd             string
	Token              string
	InsecureSkipVerify bool
	CertPath           string
}

func NewClientConfig(url string) *ClientConfig {
	return &ClientConfig{
		URL: url,
	}
}

func (c *ClientConfig) WithLoginCred(user, passwd string) *ClientConfig {
	c.User = user
	c.Passwd = passwd
	return c
}

func (c *ClientConfig) WithToken(token string) *ClientConfig {
	c.Token = token
	return c
}

func (c *ClientConfig) WithSkipTlsVerify(insecureSkipVerify bool) *ClientConfig {
	c.InsecureSkipVerify = insecureSkipVerify
	return c
}

func (c *ClientConfig) WithCertPath(certPath string) *ClientConfig {
	c.CertPath = certPath
	return c
}

func UseLocalEnvToken(envVar string) string {
	return utils.GetEnv(envVar)
}
