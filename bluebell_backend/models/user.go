package models

import (
	"encoding/json"
	"errors"
)

// User 定义请求参数结构体
type User struct {
	UserID       uint64 `json:"user_id,string" db:"user_id"` // 指定json序列化/反序列化时使用小写user_id
	UserName     string `json:"username" db:"username"`
	Password     string `json:"password" db:"password"`
	Email        string `json:"email" db:"gender"`  // 邮箱
	Gender       int    `json:"gender" db:"gender"` // 性别
	AccessToken  string
	RefreshToken string
}

// UnmarshalJSON 为User类型实现自定义的UnmarshalJSON方法
func (u *User) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName string `json:"username" db:"username"`
		Password string `json:"password" db:"password"`
		Email    string `json:"email" db:"gender"`  // 邮箱
		Gender   int    `json:"gender" db:"gender"` // 性别
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else {
		u.UserName = required.UserName
		u.Password = required.Password
		u.Email = required.Email
		u.Gender = required.Gender
	}
	return
}

// RegisterForm 注册请求参数
type RegisterForm struct {
	UserName        string `json:"username" binding:"required"`  // 用户名
	Email           string `json:"email" binding:"required"`     // 邮箱
	Gender          int    `json:"gender" binding:"oneof=0 1 2"` // 性别 0:未知 1:男 2:女
	Password        string `json:"password" binding:"required"`  // 密码
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

// LoginForm 登录请求参数
type LoginForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UnmarshalJSON 为RegisterForm类型实现自定义的UnmarshalJSON方法
func (r *RegisterForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName        string `json:"username"`
		Email           string `json:"email"`    // 邮箱
		Gender          int    `json:"gender"`   // 性别 0:未知 1:男 2:女
		Password        string `json:"password"` // 密码
		ConfirmPassword string `json:"confirm_password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else if len(required.Email) == 0 {
		err = errors.New("缺少必填字段email")
	} else if required.Password != required.ConfirmPassword {
		err = errors.New("两次密码不一致")
	} else {
		r.UserName = required.UserName
		r.Email = required.Email
		r.Gender = required.Gender
		r.Password = required.Password
		r.ConfirmPassword = required.ConfirmPassword
	}
	return
}

// VoteDataForm 投票数据
type VoteDataForm struct {
	//UserID int 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1)还是反对票(-1)取消投票(0)
}

// UnmarshalJSON 为VoteDataForm类型实现自定义的UnmarshalJSON方法
func (v *VoteDataForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID    string `json:"post_id"`
		Direction int8   `json:"direction"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.PostID) == 0 {
		err = errors.New("缺少必填字段post_id")
	} else if required.Direction == 0 {
		err = errors.New("缺少必填字段direction")
	} else {
		v.PostID = required.PostID
		v.Direction = required.Direction
	}
	return
}
