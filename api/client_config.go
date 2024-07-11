package api

type ClientConfig struct {
	URL                string
	User               string
	Passwd             string
	Token              string
	InsecureSkipVerify bool
	CertPath           string
}

func NewClientConfig(url, user, passwd string) *ClientConfig {
	return &ClientConfig{
		URL:    url,
		User:   user,
		Passwd: passwd,
	}
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
