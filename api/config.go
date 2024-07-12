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
