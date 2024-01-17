package logic

import (
	"bluebell/dao/mysql"
	"bluebell/pkg/snowflake"
)

//存放业务逻辑代码

func SignUp() {
	//判断用户是否存在
	mysql.QueryUserByUsername()
	//生成UID
	snowflake.GenID()
	//保存到数据库
	mysql.InsertUser()
}
