package sdk

import (
	"fmt"
	"testing"
)

func init() {
	SetDebug(true)
	SetSecret("cli_9e737c0e236b500e", "px5uq1u7nkowvRXT3YjH3e7IBtTaR2tA")
}

func TestGetAuthorization(t *testing.T) {
	token := GetAuthorization()
	fmt.Println(token)
}

func TestGetAccessToken(t *testing.T) {
	GetAccessToken()
	fmt.Println(AccessToken)
}

func TestGetAppAccessToken(t *testing.T) {
	token, err := GetAppAccessToken()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token)
}

func TestCode2Session(t *testing.T) {
	err := Code2Session("code")
	if err != nil {
		t.Error(err)
	}
}
