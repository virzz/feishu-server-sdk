package sdk

import (
	"errors"
	"fmt"

	"github.com/kirinlabs/HttpRequest"
	m "github.com/virzz/feishu-server-sdk/model"
)

// GetChatList Get Chat List
func GetChatList(pageSize int, pageToken string) (g *m.GroupListData, err error) {
	res, err := HttpRequest.Debug(Debug).
		SetHeaders(map[string]string{"Authorization": GetAuthorization()}).
		Get(
			fmt.Sprintf("%s%s", FeishuAPI, APIChatList),
			fmt.Sprintf("page_size=%d&page_token=%s", pageSize, pageToken),
		)
	gl := &m.GroupListResp{}
	if err = res.Json(gl); err != nil {
		return
	}
	if gl.Code == 0 {
		return gl.Data, nil
	}
	return nil, errors.New(gl.Msg)
}
