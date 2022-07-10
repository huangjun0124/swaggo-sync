package model

import (
	"time"
)

type QueryAppResponse struct {
	AppName   string    `json:"app_name"`   // 应用名
	AppNo     int32     `json:"app_no"`     // 应用序号
	AppSecret string    `json:"app_secret"` // secret
	CreatedAt time.Time `json:"created_at"` // 创建时间
	IsDel     bool      `json:"is_del"`     // 是否删除
}

type QueryAppParam struct {
	AppName string `validate:"required,min=2"` // 应用名
}

type UserLogInReq struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type UserBaseReq struct {
	UserId int64 `json:"user_id"`
}

type UserLogInResp struct {
	UserBaseReq
	UserName string `json:"user_name"`
	Token    string `json:"token"`
	ExpireAt string `json:"expire"` // yyyy-MM-dd HH:mm:ss
}

type UserPermissionResp struct {
	Permission []*Permission `json:"permission"`
	Role       Role          `json:"role"`
}

type Permission struct {
	Menu     string        `json:"menu"`               // 菜单
	Children []*Permission `json:"children,omitempty"` // 子菜单，空则默认具有所有子菜单权限
}

type Role struct {
	Type int32  `json:"type"`
	Name string `json:"name"`
}
