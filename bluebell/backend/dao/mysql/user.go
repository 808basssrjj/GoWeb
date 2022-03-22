package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const (
	SECRET = "zhaoning"
)

var (
	ErrorUserExist    = errors.New("用户已存在！")
	ErrorUserNotExist = errors.New("用户不存在！")
	ErrorInvalidPwd   = errors.New("密码错误！")
)

// CheckUserExist 用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := "select id from user where username = ?"
	var count int
	err = db.Get(&count, sqlStr, username)
	if count > 0 {
		err = ErrorUserExist
		return
	}
	return nil
}

// InsertUser 保存用户
func InsertUser(user *models.User) (err error) {
	// 密码加密
	user.Password = encryptPassword(user.Password)
	sqlStr := "insert into user(user_id, username, password) values (?, ?, ?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

// Login 用户登录
func Login(user *models.User) (err error) {
	oPwd := user.Password

	sqlStr := "select user_id, password from user where username = ? "
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	if user.Password != encryptPassword(oPwd) {
		return ErrorInvalidPwd
	}
	return nil
}

// encryptPassword 密码加密
func encryptPassword(pwd string) string {
	h := md5.New()
	h.Write([]byte(SECRET))
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}

// UserDetail 用户信息
func UserDetail(id int64) (detail *models.User, err error) {
	detail = new(models.User)
	sqlStr := "select user_id, password, username from user where user_id = ? "
	err = db.Get(detail, sqlStr, id)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
	}
	return
}
