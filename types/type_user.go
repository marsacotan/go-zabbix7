package types

type UserLoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserData bool   `json:"userData"`
}

type UserLoginRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  UserLoginParams `json:"params"`
	ID      int             `json:"id"`
}

type UserLoginTokenResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      int    `json:"id"`
}

type UserLoginUserDataResult struct {
	UserID          string `json:"userid"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	URL             string `json:"url"`
	AutoLogin       string `json:"autologin"`
	AutoLogout      string `json:"autologout"`
	Lang            string `json:"lang"`
	Refresh         string `json:"refresh"`
	Theme           string `json:"theme"`
	AttemptFailed   string `json:"attempt_failed"`
	AttemptIP       string `json:"attempt_ip"`
	AttemptClock    string `json:"attempt_clock"`
	RowsPerPage     string `json:"rows_per_page"`
	Timezone        string `json:"timezone"`
	RoleID          string `json:"roleid"`
	UserDirectoryID string `json:"userdirectoryid"`
	TSProvisioned   string `json:"ts_provisioned"`
	DebugMode       int    `json:"debug_mode"`
	Deprovisioned   bool   `json:"deprovisioned"`
	GUIAccess       string `json:"gui_access"`
	MFAID           int    `json:"mfaid"`
	AuthType        int    `json:"auth_type"`
	Type            int    `json:"type"`
	UserIP          string `json:"userip"`
	SessionID       string `json:"sessionid"`
	Secret          string `json:"secret"`
}

type UserLoginUserDataResponse struct {
	JSONRPC string                  `json:"jsonrpc"`
	Result  UserLoginUserDataResult `json:"result"`
	ID      int                     `json:"id"`
}
