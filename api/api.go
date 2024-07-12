package api

import (
	"net/http"
)

type API struct {
	Config *Config
	User   UserAPI
	Token  TokenAPI
}

type UserAPI struct {
	Config *Config
	Client *http.Client
}

type TokenAPI struct {
	Config *Config
	Client *http.Client
}
