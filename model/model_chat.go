package model

// GroupData -
type GroupData struct {
	Avatar      string `json:"avatar"`
	ChatID      string `json:"chat_id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	OwnerOpenID string `json:"owner_open_id"`
	OwnerUserID string `json:"owner_user_id"`
}

// GroupListData -
type GroupListData struct {
	HasMore   bool        `json:"has_more"`
	PageToken string      `json:"page_token"`
	Groups    []GroupData `json:"groups"`
}

// GroupListResp -
type GroupListResp struct {
	CommonResp
	Data *GroupListData `json:"data"`
}
