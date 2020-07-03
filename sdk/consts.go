package sdk

const (
	// FeishuAPI Feishu API Entrypoint
	FeishuAPI = "https://open.feishu.cn/open-apis"
)

const (
	// AccessRoleReader 只读
	AccessRoleReader = "reader"
	// AccessRoleFreeBusyReader -
	AccessRoleFreeBusyReader = "free_busy_reader"
	// AccessRoleWriter -
	AccessRoleWriter = "writer"
	// AccessRoleOwner -
	AccessRoleOwner = "owner"
)

const (
	// ActionInvite -
	ActionInvite = "invite"
	// ActionRemove -
	ActionRemove = "remove"
)

/**
 * Auth
 */

const (
	// APIAppAccessTokenInternal 获取 app_access_token（企业自建应用）
	APIAppAccessTokenInternal = "/auth/v3/app_access_token/internal/"
	// APIAppAccessToken 获取 app_access_token（应用商店应用）
	APIAppAccessToken = "/auth/v3/app_access_token/"
	// APIRefreshAccessToken 刷新access_token
	APIRefreshAccessToken = "/authen/v1/refresh_access_token"
)

/**
 * Message (Robot) 信息 (机器人)
 */

const (
	// APIRobotSendMessage 机器人发送消息
	APIRobotSendMessage = "/message/v4/send/"
	// APIRobotSendBatchMessage 机器人批量发送消息
	APIRobotSendBatchMessage = "/message/v4/batch_send/"

	// APIUploadImage 上传图片
	APIUploadImage = "/image/v4/put/"
	// APIGetImage 获取图片
	APIGetImage = "/image/v4/get"
)

/**
 * Chat 群
 */

const (
	// APIChatInfo 获取群信息
	APIChatInfo = "/chat/v4/info"
	// APIChatList 获取群信息
	APIChatList = "/chat/v4/list"
)

// From: https://github.com/galaxy-book/feishu-sdk-golang/blob/master/core/consts/consts.go
//API Host const， v3
// const (
// 	//获取 tenant_access_token（企业自建应用）
// 	ApiTenantAccessTokenInternal = "/auth/v3/tenant_access_token/internal/"
// 	//重新推送 app_ticket
// 	ApiAppTicketResend = "/auth/v3/app_ticket/resend/"
// 	//获取登录用户身份
// 	ApiOAuth2AccessToken = "https://open.feishu.cn/connect/qrconnect/oauth2/access_token/"
// 	//code2session
// 	ApiTokenLoginValidate = "/mina/v2/tokenLoginValidate"

// 	//////////////////部门和用户
// 	//获取通讯录授权范围
// 	ApiScope = "/contact/v1/scope/get"
// 	//获取通讯录授权范围v2
// 	ApiScopeV2 = "/contact/v2/scope/get"

// 	//获取部门列表
// 	ApiDepartmentSimpleList = "/contact/v1/department/simple/list"
// 	//获取部门列表 v2
// 	ApiDepartmentSimpleListV2 = "/contact/v2/department/simple/list"

// 	//获取部门详情
// 	ApiDepartmentInfoGet = "/contact/v1/department/info/get"
// 	//批量获取部门详情
// 	ApiDepartmentInfoBatchGet = "/contact/v2/department/detail/batch_get"

// 	//获取部门用户列表
// 	ApiDepartmentUserList = "/contact/v1/department/user/list"
// 	//获取部门用户列表v2
// 	ApiDepartmentUserListV2 = "/contact/v2/department/user/list"
// 	//获取用户详情
// 	ApiDepartmentUserDetailList = "/contact/v1/department/user/detail/list"
// 	//获取用户详情v2
// 	ApiDepartmentUserDetailListV2 = "/contact/v2/department/user/detail/list"

// 	//批量获取用户信息
// 	ApiUserBatchGet = "/contact/v1/user/batch_get"
// 	//批量获取用户信息v2
// 	ApiUserBatchGetV2 = "/contact/v2/user/batch_get"

// 	//////////////////角色
// 	//获取角色列表
// 	ApiRoleList = "/contact/v2/role/list"
// 	//获取角色成员列表
// 	ApiRoleMemberList = "/contact/v2/role/members"

// 	/////////////////////////////////////////////////////////
// 	//创建日历
// 	ApiCalendarCreate = "/calendar/v3/calendars"
// 	//获取日历
// 	ApiCalendarGet = "/calendar/v3/calendar_list/%s"
// 	//获取日历列表
// 	ApiCalendarListGet = "/calendar/v3/calendar_list"
// 	//更新日历
// 	ApiCalendarUpdate = "/calendar/v3/calendars/%s"
// 	//创建日程
// 	ApiCalendarEventCreate = "/calendar/v3/calendars/%s/events"
// 	//删除日程
// 	ApiCalendarEventDelete = "/calendar/v3/calendars/%s/events/%s"
// 	//邀请/移除日程参与者
// 	ApiCalendarEventAttendeesUpdate = "/calendar/v3/calendars/%s/events/%s/attendees"
// 	//获取访问控制列表
// 	ApiCalendarAttendeesGet = "/calendar/v3/calendars/%s/acl"
// 	//删除访问空值
// 	ApiCalendarAttendeesDelete = "/calendar/v3/calendars/%s/acl/%s"

// 	//搜索用户F
// 	ApiSearchUser = "/search/v1/user"

// 	//检验应用管理员
// 	ApiIsUserAdmin = "/application/v3/is_user_admin"

// 	////////用户群组
// 	//获取用户所在的群列表
// 	ApiUserGroupLIst = "/user/v4/group_list"
// 	//获取群成员列表
// 	ApiChatMembers = "/chat/v4/members"
// 	//搜索用户所在的群列表
// 	ApiChatSearch = "/chat/v4/search"

// 	////////群信息和群管理
// 	//创建群
// 	ApiCreateChat = "/chat/v4/create/"
// 	//获取群列表
// 	ApiChatList = "/chat/v4/list"
// 	//更新群信息
// 	ApiUpdateChat = "/chat/v4/update/"
// 	//拉用户进群
// 	ApiAddChatUser = "/chat/v4/chatter/add/"
// 	//移除用户出群
// 	ApiRemoveChatUser = "/chat/v4/chatter/delete/"
// 	//解散群
// 	ApiDisbandChat = "/chat/v4/disband"
// 	//拉机器人进群
// 	ApiAddBot = "/bot/v4/add"
// )
