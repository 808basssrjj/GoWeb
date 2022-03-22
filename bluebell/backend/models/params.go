package models

// 定义请求参数的结构体

// SignUpParam 注册参数
// gin中使用binding tag来做参数校验
type SignUpParam struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	//Email      string `json:"email" binding:"required,email"`
}

// LoginParam 登录参数
type LoginParam struct {
	// json tag  后加string可以序列化是转为string 反序列化是转为int
	// 用来解决前端int类型小 导致失真的问题
	UserID   int    `json:"user_id,string" binding:""`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// VoteParam 投票参数
type VoteParam struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"required,oneof=1 0 -1"` //只能是1 0 -1
}

type PostListParam struct {
	Page  int64  `form:"page"`
	Size  int64  `form:"size"`
	Order string `form:"order"`
}
