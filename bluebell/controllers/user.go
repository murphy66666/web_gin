package controllers

import (
	"bluebell/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验

	//2.业务处理
	logic.SignUp()
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{})
}
