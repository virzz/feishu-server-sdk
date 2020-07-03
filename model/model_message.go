package model

// MsgReq 文本消息
type MsgReq struct {
	OpenID      string      `json:"open_id,omitempty"`
	RootID      string      `json:"root_id,omitempty"`
	ChatID      string      `json:"chat_id,omitempty"`
	UserID      string      `json:"user_id,omitempty"`
	Email       string      `json:"email,omitempty"`
	MsgType     string      `json:"msg_type"`
	Content     *MsgContent `json:"content,omitempty"`
	UpdateMulti bool        `json:"update_multi"`
	// Card        *Card       `json:"card,omitempty"`
}

// MsgContent Msg Content
type MsgContent struct {
	Text     string   `json:"text"`
	ImageKey string   `json:"image_key"`
	Post     *MsgPost `json:"post,omitempty"`
}

// MsgPost MsgPost
type MsgPost struct {
	ZhCn *MsgPostValue `json:"zh_cn,omitempty"`
	EnUs *MsgPostValue `json:"en_us,omitempty"`
	JaJp *MsgPostValue `json:"ja_jp,omitempty"`
}

// MsgPostValue MsgPostValue
type MsgPostValue struct {
	Title   string      `json:"title"`
	Content interface{} `json:"content"`
}

// MsgResp MsgResp
type MsgResp struct {
	Code int64 `json:"code"`
	Data struct {
		MessageID string `json:"message_id,omitempty"`
		ImageKey  string `json:"image_key,omitempty"`
	} `json:"data"`
	Msg string `json:"msg"`
}
