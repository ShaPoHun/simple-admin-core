// Code generated by goctl. DO NOT EDIT.
package types

// The basic response with data | 基础带数据信息
// swagger:model BaseDataInfo
type BaseDataInfo struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response with data | 基础带数据信息
// swagger:model BaseListInfo
type BaseListInfo struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response without data | 基础不带数据信息
// swagger:model BaseMsgResp
type BaseMsgResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
}

// The simplest message | 最简单的信息
// swagger:response SimpleMsg
type SimpleMsg struct {
	// Message | 信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageInfo
type PageInfo struct {
	// Page number | 第几页
	// Required: true
	Page uint64 `json:"page" validate:"number"`
	// Page size | 单页数据行数
	// Required: true
	// Maximum: 100000
	PageSize uint64 `json:"pageSize" validate:"number,max=100000"`
}

// Basic ID request | 基础ID参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
}

// Basic IDs request | 基础ID数组参数请求
// swagger:model IDsReq
type IDsReq struct {
	// IDs
	// Required: true
	Ids []uint64 `json:"ids"`
}

// Basic ID request | 基础ID地址参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// ID
	// Required: true
	// Max length: 36
	Id string `json:"id" validate:"len=36"`
}

// Basic UUID array request | 基础UUID数组参数请求
// swagger:model UUIDsReq
type UUIDsReq struct {
	// Ids
	// Required: true
	Ids []string `json:"ids"`
}

// The base response data | 基础信息
// swagger:model BaseInfo
type BaseInfo struct {
	// ID
	Id uint64 `json:"id"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
}

// The base UUID response data | 基础信息
// swagger:model BaseUUIDInfo
type BaseUUIDInfo struct {
	// ID
	Id string `json:"id"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
}

// The response data of role information | 角色信息
// swagger:model RoleInfo
type RoleInfo struct {
	BaseInfo
	// Translated Name | 展示名称
	Trans string `json:"trans,optional"`
	// Status | 状态
	Status uint32 `json:"status,optional"`
	// Name | 角色名称
	Name string `json:"name,optional"`
	// Role code | 角色码
	Code string `json:"code,optional"`
	// DefaultRouter | 默认首页
	DefaultRouter string `json:"defaultRouter,optional"`
	// Remark | 备注
	Remark string `json:"remark,optional"`
	// Sort | 排序
	Sort uint32 `json:"sort,optional"`
}

// The response data of role list | 角色列表数据
// swagger:model RoleListResp
type RoleListResp struct {
	BaseDataInfo
	// Role list data | 角色列表数据
	Data RoleListInfo `json:"data"`
}

// Role list data | 角色列表数据
// swagger:model RoleListInfo
type RoleListInfo struct {
	BaseListInfo
	// The API list data | 角色列表数据
	Data []RoleInfo `json:"data"`
}

// Get role list request params | 角色列表请求参数
// swagger:model RoleListReq
type RoleListReq struct {
	PageInfo
	// Name
	Name string `json:"name,optional"`
}

// Role information response | 角色信息返回体
// swagger:model RoleInfoResp
type RoleInfoResp struct {
	BaseDataInfo
	// Role information | 角色数据
	Data RoleInfo `json:"data"`
}

// The response data of user information | 用户信息
// swagger:model UserInfo
type UserInfo struct {
	BaseUUIDInfo
	// Status | 状态
	Status uint32 `json:"status,optional"`
	// Username | 用户名
	Username string `json:"username,optional"`
	// Nickname | 昵称
	Nickname string `json:"nickname,optional"`
	// Password | 密码
	Password string `json:"password,optional"`
	// Description | 描述
	Description string `json:"description,optional"`
	// HomePath | 首页
	HomePath string `json:"homePath,optional"`
	// RoleId | 角色ID
	RoleIds []uint64 `json:"roleIds,optional"`
	// Mobile | 手机号
	Mobile string `json:"mobile,optional"`
	// Email | 邮箱
	Email string `json:"email,optional"`
	// Avatar | 头像地址
	Avatar string `json:"avatar,optional"`
	// Department ID | 部门ID
	DepartmentId uint64 `json:"departmentId,optional"`
	// Position ID | 职位ID
	PositionIds []uint64 `json:"positionId,optional"`
}

// The response data of user list | 用户列表数据
// swagger:model UserListResp
type UserListResp struct {
	BaseDataInfo
	// User list data | User列表数据
	Data UserListInfo `json:"data"`
}

// User list data | 用户列表数据
// swagger:model UserListInfo
type UserListInfo struct {
	BaseListInfo
	// The API list data | User列表数据
	Data []UserInfo `json:"data"`
}

// Get user list request params | 用户列表请求参数
// swagger:model UserListReq
type UserListReq struct {
	PageInfo
	// User Name | 用户名
	// Max length: 20
	Username string `json:"username,optional" validate:"omitempty,alphanum,max=20"`
	// User's nickname | 用户的昵称
	// Max length: 10
	Nickname string `json:"nickname,optional" validate:"omitempty,alphanumunicode,max=10"`
	// User's mobile phone number | 用户的手机号码
	// Max length: 18
	Mobile string `json:"mobile,optional" validate:"omitempty,numeric,max=18"`
	// The user's email address | 用户的邮箱
	// Max length: 100
	Email string `json:"email,optional" validate:"omitempty,email,max=100"`
	// User's role ID | 用户的角色ID
	// Maximum: 1000
	RoleIds []uint64 `json:"roleIds,optional"`
	// The user's department ID | 用户所属部门ID
	DepartmentId uint64 `json:"departmentId,optional"`
	// User's position id | 用户的职位ID
	PositionId uint64 `json:"positionId,optional"`
}

// User information response | 用户信息返回体
// swagger:model UserInfoResp
type UserInfoResp struct {
	BaseDataInfo
	// User information | User数据
	Data UserInfo `json:"data"`
}

// register request | 注册参数
// swagger:model RegisterReq
type RegisterReq struct {
	// User Name | 用户名
	// Required: true
	// Max length: 20
	Username string `json:"username" validate:"alphanum,max=20"`
	// Password | 密码
	// Required: true
	// Min length: 6
	// Max length: 30
	Password string `json:"password" validate:"max=30,min=6"`
	// Captcha ID which store in redis | 验证码编号, 存在redis中
	// Required: true
	// Max length: 20
	CaptchaId string `json:"captchaId" validate:"len=20"`
	// The Captcha which users input | 用户输入的验证码
	// Required: true
	// Max length: 5
	Captcha string `json:"captcha" validate:"len=5"`
	// The user's email address | 用户的邮箱
	// Required: true
	// Max length: 100
	Email string `json:"email" validate:"email,max=100"`
}

// change user's password request | 修改密码请求参数
// swagger:model ChangePasswordReq
type ChangePasswordReq struct {
	// User's old password | 用户旧密码
	// Required: true
	// Max length: 30
	OldPassword string `json:"oldPassword" validate:"max=30"`
	// User's new password | 用户新密码
	// Required: true
	// Max length: 30
	NewPassword string `json:"newPassword" validate:"max=30"`
}

// The log in information | 登陆返回的数据信息
// swagger:model LoginInfo
type LoginInfo struct {
	// User's UUID | 用户的UUID
	UserId string `json:"userId"`
	// Token for authorization | 验证身份的token
	Token string `json:"token"`
	// Expire timestamp | 过期时间戳
	Expire uint64 `json:"expire"`
}

// The simple role data | 简单的角色数据
// swagger:model RoleInfoSimple
type RoleInfoSimple struct {
	// Role name | 角色名
	RoleName string `json:"roleName"`
	// Role value | 角色值
	Value string `json:"value"`
}

// The response data of user's basic information | 用户基本信息返回数据
// swagger:model UserBaseInfoResp
type UserBaseInfoResp struct {
	BaseDataInfo
	// The  data of user's basic information | 用户基本信息
	Data UserBaseInfo `json:"data"`
}

// The  data of user's basic information | 用户基本信息
// swagger:model UserBaseInfo
type UserBaseInfo struct {
	// User's UUID | 用户的UUID
	UUID string `json:"userId"`
	// User's name | 用户名
	Username string `json:"username"`
	// User's nickname | 用户的昵称
	Nickname string `json:"nickname"`
	// The user's avatar path | 用户的头像路径
	Avatar string `json:"avatar"`
	// The home page that the user enters after logging in | 用户登陆后进入的首页
	HomePath string `json:"homePath"`
	// The description of user | 用户的描述信息
	Description string `json:"desc"`
}

// The permission code for front end permission control | 权限码： 用于前端权限控制
// swagger:model PermCodeResp
type PermCodeResp struct {
	BaseDataInfo
	// Permission code data | 权限码数据
	Data []string `json:"data"`
}

// Login request | 登录参数
// swagger:model LoginReq
type LoginReq struct {
	// User Name | 用户名
	// Required: true
	// Max length: 20
	Username string `json:"username" validate:"alphanum,max=20"`
	// Password | 密码
	// Required: true
	// Min length: 6
	// Max length: 30
	Password string `json:"password" validate:"max=30,min=6"`
	// Captcha ID which store in redis | 验证码编号, 存在redis中
	// Required: true
	// Max length: 20
	CaptchaId string `json:"captchaId"  validate:"len=20"`
	// The Captcha which users input | 用户输入的验证码
	// Required: true
	// Max length: 5
	Captcha string `json:"captcha" validate:"len=5"`
}

// The log in response data | 登录返回数据
// swagger:model LoginResp
type LoginResp struct {
	BaseDataInfo
	// The log in information | 登陆返回的数据信息
	Data LoginInfo `json:"data"`
}

// The profile information | 个人信息
// swagger:model ProfileInfo
type ProfileInfo struct {
	// user's nickname | 用户的昵称
	Nickname string `json:"nickname"`
	// The user's avatar path | 用户的头像路径
	Avatar string `json:"avatar"`
	// User's mobile phone number | 用户的手机号码
	Mobile string `json:"mobile"`
	// The user's email address | 用户的邮箱
	Email string `json:"email"`
}

// The profile response data | 个人信息返回数据
// swagger:model ProfileResp
type ProfileResp struct {
	BaseDataInfo
	// The profile information | 个人信息
	Data ProfileInfo `json:"data"`
}

// The profile request data | 个人信息请求参数
// swagger:model ProfileReq
type ProfileReq struct {
	// user's nickname | 用户的昵称
	// Max length: 10
	Nickname string `json:"nickname" validate:"omitempty,alphanumunicode,max=10"`
	// The user's avatar path | 用户的头像路径
	Avatar string `json:"avatar"`
	// User's mobile phone number | 用户的手机号码
	// Max length: 18
	Mobile string `json:"mobile" validate:"omitempty,numeric,max=18"`
	// The user's email address | 用户的邮箱
	// Max length: 100
	Email string `json:"email" validate:"omitempty,email,max=100"`
}

// The response data of menu information | 菜单信息
// swagger:model MenuInfo
type MenuInfo struct {
	BaseInfo
	// Translated Name | 国际化展示名称
	Trans string `json:"trans,optional"`
	// Level | 菜单层级
	Level uint32 `json:"level,optional"`
	// ParentId | 父级菜单ID
	ParentId uint64 `json:"parentId,optional"`
	// Path | 菜单访问路径
	Path string `json:"path,optional"`
	// Menu name | 菜单名称
	Name string `json:"name,optional"`
	// Redirect | 跳转地址
	Redirect string `json:"redirect,optional"`
	// Component | 组件地址
	Component string `json:"component,optional"`
	// Sort | 排序
	Sort uint32 `json:"sort,optional"`
	// Disabled | 是否启用
	Disabled bool `json:"disabled,optional"`
	// Meta
	Meta Meta `json:"meta"`
	// MenuType | 菜单类型
	MenuType uint32 `json:"menuType,optional"`
}

// The meta data of menu | 菜单的meta数据
// swagger:model Meta
type Meta struct {
	// Menu title show in page | 菜单显示名
	// Max length: 50
	Title string `json:"title" validate:"max=50"`
	// Menu Icon | 菜单图标
	// Max length: 50
	Icon string `json:"icon" validate:"max=50"`
	// Hide menu | 隐藏菜单
	HideMenu bool `json:"hideMenu" validate:"boolean"`
	// If hide the breadcrumb | 隐藏面包屑
	HideBreadcrumb bool `json:"hideBreadcrumb,optional" validate:"boolean"`
	// Do not keep alive the tab | 不缓存Tab
	IgnoreKeepAlive bool `json:"ignoreKeepAlive,optional" validate:"boolean"`
	// Hide the tab header | 当前路由不在标签页显示
	HideTab bool `json:"hideTab,optional" validate:"boolean"`
	// Iframe path | 内嵌iframe的地址
	FrameSrc string `json:"frameSrc,optional"`
	// The route carries parameters or not | 如果该路由会携带参数，且需要在tab页上面显示。则需要设置为true
	CarryParam bool `json:"carryParam,optional" validate:"boolean"`
	// Hide children menu or not | 隐藏所有子菜单
	HideChildrenInMenu bool `json:"hideChildrenInMenu,optional" validate:"boolean"`
	// Affix tab | 是否固定标签
	Affix bool `json:"affix,optional" validate:"boolean"`
	// The maximum number of pages the router can open | 动态路由可打开Tab页数
	DynamicLevel uint32 `json:"dynamicLevel" validate:"number,lt=30"`
	// The real path of the route without dynamic part | 动态路由的实际Path, 即去除路由的动态部分
	RealPath string `json:"realPath,optional"`
}

// The response data of menu list | 菜单列表返回数据
// swagger:model MenuListResp
type MenuListResp struct {
	BaseDataInfo
	// Menu list data | Menu列表数据
	Data MenuListInfo `json:"data"`
}

// Menu list data | Menu列表数据
// swagger:model MenuListInfo
type MenuListInfo struct {
	BaseListInfo
	// The menu list data | 菜单列表数据
	Data []MenuInfo `json:"data"`
}

// Menu information response | 菜单信息返回体
// swagger:model MenuInfoResp
type MenuInfoResp struct {
	BaseDataInfo
	// Menu information | Menu数据
	Data MenuInfo `json:"data"`
}

// Menu information plain | 菜单信息无嵌套
// swagger:model MenuPlainInfo
type MenuPlainInfo struct {
	Id uint64 `json:"id"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
	// Translated Name | 国际化展示名称
	Trans string `json:"trans,optional"`
	// Level | 菜单层级
	Level uint32 `json:"level,optional"`
	// ParentId | 父级菜单ID
	ParentId uint64 `json:"parentId,optional"`
	// Path | 菜单访问路径
	Path string `json:"path,optional"`
	// Menu name | 菜单名称
	Name string `json:"name,optional"`
	// Redirect | 跳转地址
	Redirect string `json:"redirect,optional"`
	// Component | 组件地址
	Component string `json:"component,optional"`
	// Sort | 排序
	Sort uint32 `json:"sort,optional"`
	// Disabled | 是否启用
	Disabled bool `json:"disabled,optional"`
	// MenuType | 菜单类型
	MenuType uint32 `json:"menuType,optional"`
	// Menu title show in page | 菜单显示名
	// Max length: 50
	Title string `json:"title" validate:"max=50"`
	// Menu Icon | 菜单图标
	// Max length: 50
	Icon string `json:"icon" validate:"max=50"`
	// Hide menu | 隐藏菜单
	HideMenu bool `json:"hideMenu" validate:"boolean"`
	// If hide the breadcrumb | 隐藏面包屑
	HideBreadcrumb bool `json:"hideBreadcrumb,optional" validate:"boolean"`
	// Do not keep alive the tab | 不缓存Tab
	IgnoreKeepAlive bool `json:"ignoreKeepAlive,optional" validate:"boolean"`
	// Hide the tab header | 当前路由不在标签页显示
	HideTab bool `json:"hideTab,optional" validate:"boolean"`
	// Iframe path | 内嵌iframe的地址
	FrameSrc string `json:"frameSrc,optional"`
	// The route carries parameters or not | 如果该路由会携带参数，且需要在tab页上面显示。则需要设置为true
	CarryParam bool `json:"carryParam,optional" validate:"boolean"`
	// Hide children menu or not | 隐藏所有子菜单
	HideChildrenInMenu bool `json:"hideChildrenInMenu,optional" validate:"boolean"`
	// Affix tab | 是否固定标签
	Affix bool `json:"affix,optional" validate:"boolean"`
	// The maximum number of pages the router can open | 动态路由可打开Tab页数
	DynamicLevel uint32 `json:"dynamicLevel" validate:"number,lt=30"`
	// The real path of the route without dynamic part | 动态路由的实际Path, 即去除路由的动态部分
	RealPath string `json:"realPath,optional"`
}

// Menu list data | 菜单列表数据
type MenuPlainInfoList struct {
	BaseListInfo
	// The menu list data | 菜单列表数据
	Data []MenuPlainInfo `json:"data"`
}

// Menu list data response | 菜单列表数据返回体
// swagger:model MenuPlainInfoListResp
type MenuPlainInfoListResp struct {
	BaseDataInfo
	// Menu list data | Menu列表数据
	Data MenuPlainInfoList `json:"data"`
}

// The information of captcha | 验证码数据
// swagger:model CaptchaInfo
type CaptchaInfo struct {
	CaptchaId string `json:"captchaId"`
	ImgPath   string `json:"imgPath"`
}

// The response data of captcha | 验证码返回数据
// swagger:model CaptchaResp
type CaptchaResp struct {
	BaseDataInfo
	// The menu authorization data | 菜单授权信息数据
	Data CaptchaInfo `json:"data"`
}

// The API information | API信息
// swagger:model ApiInfo
type ApiInfo struct {
	BaseInfo
	// Translated Name | 多语言名称
	Trans string `json:"trans,optional"`
	// API path | API路径
	// Required: true
	// Min length: 1
	// Max length: 50
	Path string `json:"path,optional" validate:"omitempty,min=1,max=50"`
	// API Description | API 描述
	// Required: true
	// Max length: 50
	Description string `json:"description,optional" validate:"omitempty,max=50"`
	// API group | API分组
	// Require: true
	// Min length: 1
	// Max length: 10
	Group string `json:"group,optional" validate:"omitempty,alphanum,min=1,max=10"`
	// API request method e.g. POST | API请求类型 如POST
	// Required: true
	// Min length: 3
	// Max length: 4
	Method string `json:"method,optional" validate:"omitempty,uppercase,min=3,max=4"`
}

// The response data of API list | API列表数据
// swagger:model ApiListResp
type ApiListResp struct {
	BaseDataInfo
	// API list data | API 列表数据
	Data ApiListInfo `json:"data"`
}

// API list data | API 列表数据
// swagger:model ApiListInfo
type ApiListInfo struct {
	BaseListInfo
	// The API list data | API列表数据
	Data []ApiInfo `json:"data"`
}

// Get API list request params | API列表请求参数
// swagger:model ApiListReq
type ApiListReq struct {
	PageInfo
	// API path | API路径
	// Max length: 100
	Path string `json:"path,optional" validate:"omitempty,max=100"`
	// API Description | API 描述
	// Max length: 50
	Description string `json:"description,optional" validate:"omitempty,max=50"`
	// API group | API分组
	// Max length: 20
	Group string `json:"group,optional" validate:"omitempty,max=20"`
	// API request method e.g. POST | API请求类型 如POST
	// Max length: 4
	Method string `json:"method,optional" validate:"omitempty,uppercase,max=4"`
}

// API information response | API信息返回体
// swagger:model ApiInfoResp
type ApiInfoResp struct {
	BaseDataInfo
	// API information | API数据
	Data ApiInfo `json:"data"`
}

// The response data of api authorization | API授权数据
// swagger:model ApiAuthorityInfo
type ApiAuthorityInfo struct {
	// API path | API 路径
	Path string `json:"path"`
	// API method | API请求方法
	Method string `json:"method"`
}

// Create or update api authorization information request | 创建或更新API授权信息
// swagger:model CreateOrUpdateApiAuthorityReq
type CreateOrUpdateApiAuthorityReq struct {
	// Role ID | 角色ID
	// Required: true
	// Maximum: 1000
	RoleId uint64 `json:"roleId" validate:"number,max=1000"`
	// API authorization list | API授权列表数据
	// Required: true
	Data []ApiAuthorityInfo `json:"data"`
}

// The response data of api authorization list | API授权列表返回数据
// swagger:model ApiAuthorityListResp
type ApiAuthorityListResp struct {
	BaseDataInfo
	// The api authorization list data | API授权列表数据
	Data ApiAuthorityListInfo `json:"data"`
}

// The  data of api authorization list | API授权列表数据
// swagger:model ApiAuthorityListInfo
type ApiAuthorityListInfo struct {
	BaseListInfo
	// The api authorization list data | API授权列表数据
	Data []ApiAuthorityInfo `json:"data"`
}

// Create or update menu authorization information request params | 创建或更新菜单授权信息参数
// swagger:model MenuAuthorityInfoReq
type MenuAuthorityInfoReq struct {
	// role ID | 角色ID
	// Required: true
	// Maximum: 1000
	RoleId uint64 `json:"roleId" validate:"number,max=1000"`
	// menu ID array | 菜单ID数组
	// Required: true
	MenuIds []uint64 `json:"menuIds"`
}

// Menu authorization response data | 菜单授权信息数据
// swagger:model MenuAuthorityInfoResp
type MenuAuthorityInfoResp struct {
	BaseDataInfo
	// The menu authorization data | 菜单授权信息数据
	Data MenuAuthorityInfoReq `json:"data"`
}

// The response data of dictionary information | 字典信息
// swagger:model DictionaryInfo
type DictionaryInfo struct {
	BaseInfo
	// Translated Name | 字典多语言名称
	Trans string `json:"trans,optional"`
	// Title | 字典多语言名称
	Title string `json:"title,optional"`
	// Name | 字典名称
	Name string `json:"name,optional"`
	// Status | 状态
	Status uint32 `json:"status,optional"`
	// Description of dictionary | 字典描述
	Desc string `json:"desc,optional"`
}

// The response data of dictionary list | 字典列表数据
// swagger:model DictionaryListResp
type DictionaryListResp struct {
	BaseDataInfo
	// Dictionary list data | 字典列表数据
	Data DictionaryListInfo `json:"data"`
}

// Dictionary list data | 字典列表数据
// swagger:model DictionaryListInfo
type DictionaryListInfo struct {
	BaseListInfo
	// The API list data | 字典列表数据
	Data []DictionaryInfo `json:"data"`
}

// Get dictionary list request params | 字典列表请求参数
// swagger:model DictionaryListReq
type DictionaryListReq struct {
	PageInfo
	// Name
	Name string `json:"name,optional"`
}

// Dictionary information response | 字典信息返回体
// swagger:model DictionaryInfoResp
type DictionaryInfoResp struct {
	BaseDataInfo
	// Dictionary information | 字典数据
	Data DictionaryInfo `json:"data"`
}

// The response data of oauth provider information | 第三方信息
// swagger:model OauthProviderInfo
type OauthProviderInfo struct {
	BaseInfo
	// Provider name | 第三方提供商名称
	Name string `json:"name,optional"`
	// ClientId | 客户端ID
	ClientId string `json:"clientId,optional"`
	// ClientSecret | 客户端密钥
	ClientSecret string `json:"clientSecret,optional"`
	// Redirect URL| 跳转地址
	RedirectUrl string `json:"redirectUrl,optional"`
	// Scopes | 授权范围
	Scopes string `json:"scopes,optional"`
	// Authority URL | 授权地址
	AuthUrl string `json:"authUrl,optional"`
	// The URL to get token | 获取Token的地址
	TokenUrl string `json:"tokenUrl,optional"`
	// The type of auth | 鉴权方式
	AuthStyle uint64 `json:"authStyle,optional"`
	// The URL to get user information | 获取信息地址
	InfoUrl string `json:"infoUrl,optional"`
}

// The response data of oauth provider list | 第三方列表数据
// swagger:model OauthProviderListResp
type OauthProviderListResp struct {
	BaseDataInfo
	// OauthProvider list data | 第三方列表数据
	Data OauthProviderListInfo `json:"data"`
}

// OauthProvider list data | 第三方列表数据
// swagger:model OauthProviderListInfo
type OauthProviderListInfo struct {
	BaseListInfo
	// The API list data | 第三方列表数据
	Data []OauthProviderInfo `json:"data"`
}

// Get oauth provider list request params | 第三方列表请求参数
// swagger:model OauthProviderListReq
type OauthProviderListReq struct {
	PageInfo
	// Name
	Name string `json:"name,optional"`
	// ClientId
	ClientId string `json:"clientId,optional"`
	// ClientSecret
	ClientSecret string `json:"clientSecret,optional"`
}

// Oauth provider information response | 第三方信息返回体
// swagger:model OauthProviderInfoResp
type OauthProviderInfoResp struct {
	BaseDataInfo
	// OauthProvider information | 第三方数据
	Data OauthProviderInfo `json:"data"`
}

// Oauth log in request | Oauth 登录请求
// swagger:model OauthLoginReq
type OauthLoginReq struct {
	// State code to avoid hack | 状态码，请求前后相同避免安全问题
	// Required: true
	// Max Length: 30
	State string `json:"state" validate:"max=30"`
	// Provider name | 提供商名字
	// Required: true
	// Max Length: 40
	// Example: google
	Provider string `json:"provider" validate:"max=40"`
}

// Redirect response | 跳转网址返回信息
// swagger:model RedirectResp
type RedirectResp struct {
	BaseDataInfo
	// Redirect information | 跳转网址
	Data RedirectInfo `json:"data"`
}

// Redirect information | 跳转网址
// swagger:model RedirectInfo
type RedirectInfo struct {
	// Redirect URL | 跳转网址
	URL string `json:"URL"`
}

// The oauth callback response data | Oauth回调数据
// swagger:model CallbackResp
type CallbackResp struct {
	// User's UUID | 用户的UUID
	UserId string `json:"userId"`
	// User's role information| 用户的角色信息
	// in: body
	Role RoleInfoSimple `json:"role"`
	// Token for authorization | 验证身份的token
	Token string `json:"token"`
	// Expire timestamp | 过期时间戳
	Expire uint64 `json:"expire"`
}

// The response data of token information | Token信息
// swagger:model TokenInfo
type TokenInfo struct {
	BaseUUIDInfo
	// Status
	Status uint32 `json:"status,optional"`
	// User's UUID | 用户的UUID
	Uuid string `json:"uuid,optional"`
	// Token | 用户的Token
	Token string `json:"token,optional"`
	// Source | Token 来源
	Source string `json:"source,optional"`
	// ExpiredAt | 过期时间
	ExpiredAt int64 `json:"expiredAt,optional"`
}

// The response data of token list | Token列表数据
// swagger:model TokenListResp
type TokenListResp struct {
	BaseDataInfo
	// Token list data | Token列表数据
	Data TokenListInfo `json:"data"`
}

// Token list data | Token列表数据
// swagger:model TokenListInfo
type TokenListInfo struct {
	BaseListInfo
	// The API list data | Token列表数据
	Data []TokenInfo `json:"data"`
}

// Get token list request params | Token列表请求参数
// swagger:model TokenListReq
type TokenListReq struct {
	PageInfo
	// Username
	Username string `json:"username,optional"`
	// Nickname
	Nickname string `json:"nickname,optional"`
	// Email
	Email string `json:"email,optional"`
	// Uuid
	Uuid string `json:"uuid,optional"`
}

// Token information response | Token信息返回体
// swagger:model TokenInfoResp
type TokenInfoResp struct {
	BaseDataInfo
	// Token information | Token数据
	Data TokenInfo `json:"data"`
}

// The response data of department information | 部门信息
// swagger:model DepartmentInfo
type DepartmentInfo struct {
	BaseInfo
	// Translated Name
	Trans string `json:"trans,optional"`
	// Status | 状态
	Status uint32 `json:"status,optional"`
	// Sort | 排序
	Sort uint32 `json:"sort,optional"`
	// Name | 部门名称
	Name string `json:"name,optional"`
	// Ancestors | 父级部门列表
	Ancestors string `json:"ancestors,optional"`
	// Leader | 部门负责人
	Leader string `json:"leader,optional"`
	// Phone | 电话号码
	Phone string `json:"phone,optional"`
	// Email | 邮箱
	Email string `json:"email,optional"`
	// Remark | 备注
	Remark string `json:"remark,optional"`
	// ParentId | 父级 ID
	ParentId uint64 `json:"parentId,optional"`
}

// The response data of department list | 部门列表数据
// swagger:model DepartmentListResp
type DepartmentListResp struct {
	BaseDataInfo
	// Department list data | 部门列表数据
	Data DepartmentListInfo `json:"data"`
}

// Department list data | 部门列表数据
// swagger:model DepartmentListInfo
type DepartmentListInfo struct {
	BaseListInfo
	// The API list data | 部门列表数据
	Data []DepartmentInfo `json:"data"`
}

// Get department list request params | 部门列表请求参数
// swagger:model DepartmentListReq
type DepartmentListReq struct {
	PageInfo
	// Name
	Name string `json:"name,optional"`
	// Leader
	Leader string `json:"leader,optional"`
}

// Department information response | 部门信息返回体
// swagger:model DepartmentInfoResp
type DepartmentInfoResp struct {
	BaseDataInfo
	// Department information | 部门数据
	Data DepartmentInfo `json:"data"`
}

// The response data of position information | 职位信息
// swagger:model PositionInfo
type PositionInfo struct {
	BaseInfo
	// Translated Name
	Trans string `json:"trans,optional"`
	// Status
	Status uint32 `json:"status,optional"`
	// Sort
	Sort uint32 `json:"sort,optional"`
	// Name
	Name string `json:"name,optional"`
	// Code
	Code string `json:"code,optional"`
	// Remark
	Remark string `json:"remark,optional"`
}

// The response data of position list | 职位列表数据
// swagger:model PositionListResp
type PositionListResp struct {
	BaseDataInfo
	// Position list data | 职位列表数据
	Data PositionListInfo `json:"data"`
}

// Position list data | 职位列表数据
// swagger:model PositionListInfo
type PositionListInfo struct {
	BaseListInfo
	// The API list data | 职位列表数据
	Data []PositionInfo `json:"data"`
}

// Get position list request params | 职位列表请求参数
// swagger:model PositionListReq
type PositionListReq struct {
	PageInfo
	// Name
	Name string `json:"name,optional"`
	// Code
	Code string `json:"code,optional"`
	// Remark
	Remark string `json:"remark,optional"`
}

// Position information response | 职位信息返回体
// swagger:model PositionInfoResp
type PositionInfoResp struct {
	BaseDataInfo
	// Position information | 职位数据
	Data PositionInfo `json:"data"`
}

// The response data of dictionary detail information | 字典键值信息
// swagger:model DictionaryDetailInfo
type DictionaryDetailInfo struct {
	BaseInfo
	// Status | 状态
	Status uint32 `json:"status,optional"`
	// Title | 显示名称
	Title string `json:"title,optional"`
	// Key | 键
	Key string `json:"key,optional"`
	// Value | 值
	Value string `json:"value,optional"`
	// Dictionary ID | 所属字典ID
	DictionaryId uint64 `json:"dictionaryId,optional"`
	// Sort | 排序
	Sort uint32 `json:"sort,optional"`
}

// The response data of dictionary detail list | 字典键值列表数据
// swagger:model DictionaryDetailListResp
type DictionaryDetailListResp struct {
	BaseDataInfo
	// DictionaryDetail list data | 字典键值列表数据
	Data DictionaryDetailListInfo `json:"data"`
}

// DictionaryDetail list data | 字典键值列表数据
// swagger:model DictionaryDetailListInfo
type DictionaryDetailListInfo struct {
	BaseListInfo
	// The API list data | 字典键值列表数据
	Data []DictionaryDetailInfo `json:"data"`
}

// Get dictionary detail list request params | 字典键值列表请求参数
// swagger:model DictionaryDetailListReq
type DictionaryDetailListReq struct {
	PageInfo
	// Key | 键
	Key string `json:"key,optional"`
	// Dictionary ID
	DictionaryId uint64 `json:"dictionaryId,optional"`
}

// DictionaryDetail information response | 字典键值信息返回体
// swagger:model DictionaryDetailInfoResp
type DictionaryDetailInfoResp struct {
	BaseDataInfo
	// DictionaryDetail information | 字典键值数据
	Data DictionaryDetailInfo `json:"data"`
}

// The response data of menu param information | 菜单参数信息
// swagger:model MenuParamInfo
type MenuParamInfo struct {
	BaseInfo
	// Type | 参数类型
	Type string `json:"type,optional"`
	// Key | 参数键
	Key string `json:"key,optional"`
	// Value | 参数值
	Value string `json:"value,optional"`
	// Parent menu ID | 父级菜单ID
	MenuId uint64 `json:"menuId,optional"`
}

// The response data of menu param list | 菜单参数列表数据
// swagger:model MenuParamListResp
type MenuParamListResp struct {
	BaseDataInfo
	// MenuParam list data | 菜单参数列表数据
	Data MenuParamListInfo `json:"data"`
}

// MenuParam list data | 菜单参数列表数据
// swagger:model MenuParamListInfo
type MenuParamListInfo struct {
	BaseListInfo
	// The API list data | 菜单参数列表数据
	Data []MenuParamInfo `json:"data"`
}

// Get menu param list request params | 菜单参数列表请求参数
// swagger:model MenuParamListReq
type MenuParamListReq struct {
	PageInfo
	// Menu ID | 菜单ID
	MenuId uint64 `json:"menuId"`
}

// MenuParam information response | 菜单参数信息返回体
// swagger:model MenuParamInfoResp
type MenuParamInfoResp struct {
	BaseDataInfo
	// MenuParam information | MenuParam数据
	Data MenuParamInfo `json:"data"`
}

// The response data of task information | 定时任务信息
// swagger:model TaskInfo
type TaskInfo struct {
	BaseInfo
	// Status
	Status uint32 `json:"status,optional"`
	// Name
	Name string `json:"name,optional"`
	// TaskGroup
	TaskGroup string `json:"taskGroup,optional"`
	// CronExpression
	CronExpression string `json:"cronExpression,optional"`
	// Pattern
	Pattern string `json:"pattern,optional"`
	// Payload
	Payload string `json:"payload,optional"`
}

// The response data of task list | 定时任务列表数据
// swagger:model TaskListResp
type TaskListResp struct {
	BaseDataInfo
	// Task list data | 定时任务列表数据
	Data TaskListInfo `json:"data"`
}

// Task list data | 定时任务列表数据
// swagger:model TaskListInfo
type TaskListInfo struct {
	BaseListInfo
	// The API list data | 定时任务列表数据
	Data []TaskInfo `json:"data"`
}

// Get task list request params | 定时任务列表请求参数
// swagger:model TaskListReq
type TaskListReq struct {
	PageInfo
	// Name
	Name string `json:"name,optional"`
	// TaskGroup
	TaskGroup string `json:"taskGroup,optional"`
}

// Task information response | 定时任务信息返回体
// swagger:model TaskInfoResp
type TaskInfoResp struct {
	BaseDataInfo
	// Task information | 定时任务数据
	Data TaskInfo `json:"data"`
}

// The response data of task log information | 任务日志信息
// swagger:model TaskLogInfo
type TaskLogInfo struct {
	// ID
	Id uint64 `json:"id"`
	// StartedAt
	StartedAt int64 `json:"startedAt,optional"`
	// FinishedAt
	FinishedAt int64 `json:"finishedAt,optional"`
	// Result
	Result uint32 `json:"result,optional"`
}

// The response data of task log list | 任务日志列表数据
// swagger:model TaskLogListResp
type TaskLogListResp struct {
	BaseDataInfo
	// TaskLog list data | 任务日志列表数据
	Data TaskLogListInfo `json:"data"`
}

// TaskLog list data | 任务日志列表数据
// swagger:model TaskLogListInfo
type TaskLogListInfo struct {
	BaseListInfo
	// The API list data | 任务日志列表数据
	Data []TaskLogInfo `json:"data"`
}

// Get task log list request params | 任务日志列表请求参数
// swagger:model TaskLogListReq
type TaskLogListReq struct {
	PageInfo
	TaskId uint64 `json:"taskId"`
	Result uint32 `json:"result"`
}

// TaskLog information response | 任务日志信息返回体
// swagger:model TaskLogInfoResp
type TaskLogInfoResp struct {
	BaseDataInfo
	// TaskLog information | 任务日志数据
	Data TaskLogInfo `json:"data"`
}
