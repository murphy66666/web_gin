package gin_demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func func1(c *gin.Context) {
	fmt.Println("func1")
}
func func2(c *gin.Context) {
	fmt.Println("func2 before")
	c.Next()
	fmt.Println("func2 after")
}
func func3(c *gin.Context) {
	fmt.Println("func3")
	c.Abort()
}
func func4(c *gin.Context) {
	fmt.Println("func4")
	c.Set("name", "战鹰")
}
func func5(c *gin.Context) {
	fmt.Println("func5")
	v, ok := c.Get("name")
	if ok {
		vstr := v.(string)
		fmt.Println(vstr)
	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": struct{}{},
		})
	})

	shopGroup := r.Group("/shop", func1, func2)
	shopGroup.Use(func3)
	{
		shopGroup.GET("/index", func4, func5)
	}

	r.Run(":8093")
}
