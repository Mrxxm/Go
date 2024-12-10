package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type SignUpParam struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func main() {
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var json LoginForm
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "xxm" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "账号或密码错误"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "登录成功"})
	})

	router.POST("/loginForm", func(c *gin.Context) {
		var form LoginForm

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "xxm" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "账号或密码错误"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "登录成功"})
	})

	router.POST("/signup", func(c *gin.Context) {
		var u SignUpParam
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// 保存入库等业务逻辑代码...

		c.JSON(http.StatusOK, "success")
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8081")
}
