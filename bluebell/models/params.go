package models

// ParamSignUp 定义请求的参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"require"`
	Password   string `json:"password" binding:"require"`
	RePassword string `json:"re_password" binding:"require,eqfield=Password"`
}
