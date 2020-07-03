package sdk

import (
	"fmt"
	"testing"
)

func TestGetChatList(t *testing.T) {
	gl, err := GetChatList(100, "")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(gl)
}
