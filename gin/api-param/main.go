package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func bookDetailHandler(c *gin.Context) {
	bookId := c.Param("id")
	// gin.Context封装了request和response,请求的上下文
	c.String(http.StatusOK, fmt.Sprintf("成功获取书籍详情：%s", bookId))
}

// gin框架的示例
func main() {
	// 实例化gin.Engine结构体对象
	r := gin.Default()
	// 注册路由规则，以及绑定处理函数
	// 基本路由 /book/24/
	r.GET("/book/:id", bookDetailHandler)
	// 监听端口，启动
	fmt.Println("运行地址：http://127.0.0.1:8080/book/24")
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("gin启动失败,err: ", err)
	}
}
