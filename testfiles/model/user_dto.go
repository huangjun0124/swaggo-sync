package model

import "time"

type CreateUserReq struct {
	Username string `json:"userName"  binding:"required,max=255"`       // 用户名
	NickName string `json:"nickName"  binding:"required,max=100"`       // 用户昵称
	Password string `json:"password"  binding:"required,min=3,max=255"` // 密码
	Gender   int32  `json:"gender"  binding:"max=2"`                    // 性别(1 男 2 女)
	Phone    string `json:"phone"  binding:"max=30"`                    // 手机号
}

type CreateUserResp struct {
	Id       int64  `json:"id"`                                         // 用户id
	Username string `json:"username"  binding:"required,min=3,max=255"` // 用户名
}

type ModifyUserReq struct {
	UserId   int64  `json:"userId" binding:"required"`            // 用户id
	NickName string `json:"nickName"  binding:"required,max=100"` // 用户昵称
	Phone    string `json:"phone"  binding:"max=30"`              // 手机号
}

type ChangePwdReq struct {
	UserId   int64  `json:"userId" binding:"required"`                  // 用户id
	Password string `json:"password"  binding:"required,min=3,max=255"` // 密码
}

type QueryUsersReq struct {
	UserName  string  `json:"userName"`                            // 登录账户名
	NickName  string  `json:"nickName"`                            // 用户昵称
	Roles     []int64 `json:"roles"`                               // 角色id
	PageIndex int32   `json:"pageIndex" binding:"required"`        // 当前页码(从1开始)
	PageSize  int32   `json:"pageSize" binding:"required,max=500"` // 每页行数
}

type UserInfo struct {
	UserId     int64      `json:"userId"`   // 用户id
	UserName   string     `json:"userName"` // 登录账户名
	NickName   string     `json:"nickName"` // 用户昵称
	Phone      string     `json:"phone"`    // 手机号
	Roles      []UserRole `json:"roles"`    // 角色名
	Status     int32      `json:"status"`   // 状态(1:正常,2:停用)
	CreateTime time.Time  `json:"createTime"`
	UpdateTime time.Time  `json:"updateTime"`
}

type QueryUsersResp struct {
	Users []*UserInfo `json:"users"`
	Total int32       `json:"total"` // 总记录数
}

type RoleInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type UserRole struct {
	UserId int64  `json:"userId"`
	Id     int64  `json:"roleId"`
	Name   string `json:"roleName"`
}

type ChangeUserStateReq struct {
	UserId int64 `json:"userId"`
	Oper   int32 `json:"oper"` // 操作类型(1:停用,2:启用,3:删除)
}

type UserSelector struct {
	UserId   int64  `json:"userId"`   // 用户id
	UserName string `json:"userName"` // 登录账户名
	NickName string `json:"nickName"` // 用户昵称
}

type BindUserRoleReq struct {
	UserId  int64   `json:"userId"`  // 用户id
	RoleIds []int64 `json:"roleIds"` // 角色列表
}
