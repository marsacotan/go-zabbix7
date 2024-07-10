package types

type CreateUserTokenWithExpirationParams struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	ExpiresAt int64  `json:"expires_at"`
}

type CreateUserTokenWithExpirationRequest struct {
	JSONRPC string                              `json:"jsonrpc"`
	Method  string                              `json:"method"`
	Params  CreateUserTokenWithExpirationParams `json:"params"`
	ID      int                                 `json:"id"`
}

type CreatePermanentTokenForUserParams struct {
	Name   string `json:"name"`
	UserID string `json:"userid"`
}

type CreatePermanentTokenForUserRequest struct {
	JSONRPC string                            `json:"jsonrpc"`
	Method  string                            `json:"method"`
	Params  CreatePermanentTokenForUserParams `json:"params"`
	ID      int                               `json:"id"`
}

type TokenCreateResult struct {
	TokenIDs []string `json:"tokenids"`
}

type TokenCreateResponse struct {
	JSONRPC string            `json:"jsonrpc"`
	Result  TokenCreateResult `json:"result"`
	ID      int               `json:"id"`
}
