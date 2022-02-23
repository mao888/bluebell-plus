package models

import (
	"encoding/json"
	"errors"
)
/**
 * @Author huchao
 * @Description //TODO 定义请求参数结构体
 * @Date 22:09 2022/2/10
 **/
type User struct {
	UserID   uint64 `json:"user_id,string" db:"user_id"` // 指定json序列化/反序列化时使用小写user_id
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	AccessToken    string
	RefreshToken   string
}

// UnmarshalJSON 为User类型实现自定义的UnmarshalJSON方法
func (u *User) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName string `json:"username" db:"username"`
		Password string `json:"password" db:"password"`
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
	}
	return
}

/**
 * @Author huchao
 * @Description //TODO 注册请求参数
 * @Date 22:09 2022/2/10
 **/
type RegisterForm struct {
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

/**
 * @Author huchao
 * @Description //TODO 登录请求参数
 * @Date 22:09 2022/2/10
 **/
type LoginForm struct {
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

// UnmarshalJSON 为RegisterForm类型实现自定义的UnmarshalJSON方法
func (r *RegisterForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else if required.Password != required.ConfirmPassword {
		err = errors.New("两次密码不一致")
	} else {
		r.UserName = required.UserName
		r.Password = required.Password
		r.ConfirmPassword = required.ConfirmPassword
	}
	return
}

/**
 * @Author huchao
 * @Description //TODO 投票数据
 * @Date 11:01 2022/2/14
 **/
type VoteDataForm struct {
	//UserID int 从请求中获取当前的用户
	PostID    string  `json:"post_id" binding:"required"`	 // 帖子id
	Direction int8     `json:"direction,string" binding:"oneof=1 0 -1"`  // 赞成票(1)还是反对票(-1)取消投票(0)
}
// UnmarshalJSON 为VoteDataForm类型实现自定义的UnmarshalJSON方法
func (v *VoteDataForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID    string  `json:"post_id"`
		Direction int8 `json:"direction"`
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