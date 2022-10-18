package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	// gin.Context封装了request和response,请求的上下文
	c.String(http.StatusOK, "hello gin!")
}

// gin框架的示例
func main() {
	// fmt.Println(gin.Version)
	// 实例化gin.Engine结构体对象
	r := gin.Default()
	// 注册路由规则，以及绑定处理函数
	r.GET("/hello", helloHandler)
	// 监听端口，启动
	fmt.Println("运行地址：http://127.0.0.1:8080")
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("gin启动失败,err: ", err)
	}
}
