package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func loginHandler(c *gin.Context) {
	var login Login
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.String(http.StatusOK, fmt.Sprintf("姓名： %s -- 密码：%s", login.Username, login.Password))
}

func main() {
	r := gin.Default()
	r.POST("/login", loginHandler)
	fmt.Println("运行地址：http://127.0.0.1:8080/login")
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("gin启动失败,err: ", err)
	}
}
