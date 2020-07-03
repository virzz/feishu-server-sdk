package sdk

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/kirinlabs/HttpRequest"
	m "github.com/virzz/feishu-server-sdk/model"
)

var (
	// AppID App ID
	AppID string
	// AppSecret App Secret
	AppSecret string
	// AccessToken Access Token Struct
	AccessToken = &m.AccessTokenResp{}
	// Session Session Infomation
	Session = &m.SessionResp{}
	// Authorization -
	Authorization = "Bearer "
)

// SetSecret Set App Secret
func SetSecret(appID, appSecret string) {
	AppID = appID
	AppSecret = appSecret
}

// GetAccessToken Get Access Token
func GetAccessToken() (err error) {
	if Debug {
		fmt.Println("[D] GetAccessToken")
	}
	data := &m.AppKeyReq{
		AppID:     AppID,
		AppSecret: AppSecret,
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}
	res, err := HttpRequest.Debug(Debug).
		SetHeaders(map[string]string{}).JSON().
		Post(
			fmt.Sprintf("%s%s", FeishuAPI, APIAppAccessTokenInternal),
			string(dataJSON),
		)
	if err = res.Json(AccessToken); err != nil {
		return err
	}
	AccessToken.Expire = time.Now().Unix() + AccessToken.Expire
	return nil
}

// GetAppAccessToken Get App Access Token
func GetAppAccessToken() (data string, err error) {
	if Debug {
		fmt.Println("[D] GetAppAccessToken")
		fmt.Println(AccessToken.Expire)
		fmt.Println(AccessToken.Code)
		fmt.Println(AccessToken.Msg)
		fmt.Println(AccessToken.AppAccessToken)
	}
	if AccessToken.Expire == 0 || time.Now().Unix() >= AccessToken.Expire {
		if err := GetAccessToken(); err != nil {
			return "", err
		}
	}
	return AccessToken.AppAccessToken, nil
}

// GetAuthorization -
func GetAuthorization() string {
	if Debug {
		fmt.Println("[D] GetAuthorization")
	}
	token, err := GetAppAccessToken()
	if err != nil {
		log.Println(err)
		return ""
	}
	// Authorization = fmt.Sprintf("Bearer %s", token)
	return fmt.Sprintf("Bearer %s", token)
}

// Code2Session Code2Session
func Code2Session(code string) (err error) {
	appAccessToken, err := GetAppAccessToken()
	if err != nil {
		return
	}
	res, err := HttpRequest.Debug(Debug).
		SetHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", appAccessToken),
			"Content-Type":  "application/json",
		}).
		Post(
			fmt.Sprintf("%s/mina/v2/tokenLoginValidate", FeishuAPI),
			fmt.Sprintf(`{"code": "%s"}`, code),
		)
	if err = res.Json(Session); err != nil {
		return err
	}
	// open_id, session_key, access_token, expires_in, refresh_token
	return nil
}

// DecryptUserInfo Decrypt UserInfo From Client
func DecryptUserInfo(encryptedData, key, iv, signature string) (data *m.EncryptData, err error) {
	res, err := aesDecrypt(encryptedData, key, iv)
	if err != nil {
		return data, err
	}
	if err = json.Unmarshal(res, data); err != nil {
		return data, err
	}
	return data, nil
}
