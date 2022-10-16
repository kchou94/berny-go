package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin框架的示例
func main() {
	// fmt.Println(gin.Version)
	// 实例化gin.Engine结构体对象
	r := gin.Default()
	// 注册路由规则，以及绑定处理函数
	// gin.Context封装了request和response,请求的上下文
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin!")
	})
	// 监听端口，启动
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("gin启动失败,err: ", err)
	}
}
