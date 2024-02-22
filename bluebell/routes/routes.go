package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册路由
	r.POST("/signup", controllers.SignUpHandler)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": struct{}{}})
	})

	r.GET("version", func(c *gin.Context) {
		version := viper.GetString("app.version")
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": version})
	})

	return r
}
