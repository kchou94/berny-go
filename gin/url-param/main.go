package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultQuery()若参数不存在，返回默认值，Query()若不存在，返回空串
func userDetailHandler(c *gin.Context) {
	// username := c.DefaultQuery("name", "xxx")
	username := c.Query("name")
	c.String(http.StatusOK, fmt.Sprintf("姓名：%s", username))
}

func main() {
	r := gin.Default()
	// 基本路由 /user?name=root
	r.GET("/user/", userDetailHandler)
	fmt.Println("运行地址：http://127.0.0.1:8080/user?name=root")
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("gin启动失败,err: ", err)
	}
}
