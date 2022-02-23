package logic

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"bluebell_backend/pkg/jwt"
	"bluebell_backend/pkg/snowflake"
	//"go.uber.org/zap"
)

//func UserPasswordValid() {
//
//}

/**
 * @Author huchao
 * @Description //TODO 存放注册业务逻辑的代码
 * @Date 21:52 2022/2/10
 **/
func SignUp(p *models.RegisterForm) (error error) {
	// 1、判断用户存不存在
	err := mysql.CheckUserExist(p.UserName)
	if err !=nil {
		// 数据库查询出错
		return err
	}

	// 2、生成UID
	userId, err := snowflake.GetID()
	if err != nil {
		return mysql.ErrorGenIDFailed
	}
	// 构造一个User实例
	u := models.User{
		UserID:   userId,
		UserName: p.UserName,
		Password: p.Password,
	}
	// 3、保存进数据库
	return mysql.InsertUser(u)
}

/**
 * @Author huchao
 * @Description //TODO 存放登录业务逻辑代码
 * @Date 21:52 2022/2/10
 **/
func Login(p *models.LoginForm) (user *models.User, error error)  {
	user = &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil,err
	}
	// 生成JWT
	//return jwt.GenToken(user.UserID,user.UserName)
	atoken,rtoken, err := jwt.GenToken(user.UserID,user.UserName)
	if err != nil {
		return
	}
	user.AccessToken = atoken
	user.RefreshToken = rtoken
	return
}
