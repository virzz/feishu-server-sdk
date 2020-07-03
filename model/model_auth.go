package model

// AppKeyReq -
type AppKeyReq struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

// AccessTokenResp -
type AccessTokenResp struct {
	AppAccessToken    string `json:"app_access_token"`
	Code              int64  `json:"code"`
	Expire            int64  `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

// SessionResp -
type SessionResp struct {
	CommonResp
	Data struct {
		AccessToken  string `json:"access_token"`
		EmployeeID   string `json:"employee_id"`
		ExpiresIn    int64  `json:"expires_in"`
		OpenID       string `json:"open_id"`
		RefreshToken string `json:"refresh_token"`
		SessionKey   string `json:"session_key"`
		TenantKey    string `json:"tenant_key"`
		TokenType    string `json:"token_type"`
		UID          string `json:"uid"`
		UnionID      string `json:"union_id"`
	} `json:"data"`
}

// EncryptData -
type EncryptData struct {
	AvatarURL string `json:"avatarUrl"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Gender    int64  `json:"gender"`
	NickName  string `json:"nickName"`
	OpenID    string `json:"openId"`
	Province  string `json:"province"`
	Watermark struct {
		Appid     string `json:"appid"`
		Timestamp int64  `json:"timestamp"`
	} `json:"watermark"`
}
