package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

// SignUp 用户注册
func SignUp(param *models.SignUpParam) (err error) {
	// 1.用户是否存在
	if err = mysql.CheckUserExist(param.Username); err != nil {
		return err
	}

	// 2,生成userid
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		Username: param.Username,
		Password: param.Password,
	}
	return mysql.InsertUser(user)
}

// Login 用户登录
func Login(param *models.LoginParam) (aToken, rToken string, err error) {
	user := &models.User{
		Username: param.Username,
		Password: param.Password,
	}
	if err = mysql.Login(user); err != nil {
		return "", "", err
	}

	param.UserID = int(user.UserID)
	// 生成token
	return jwt.GenToken(user.UserID)
}
