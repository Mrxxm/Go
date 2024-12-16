package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		fmt.Println("日志中间件 before request")

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		for k, v := range c.Request.Header {
			if k == "Authorization" {
				token = v[0]
			} else {
				fmt.Println(k, v)
			}
		}

		if token != "xxm" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "认证失败",
			})
			//return // 退出 不生效
			c.Abort() // 退出 生效
		}

		// before request
		fmt.Println("认证中间件 before request")

		c.Next()
	}
}

func main() {
	// 创建一个不包含中间件的路由器
	r := gin.New()

	// 全局中间件
	// 使用 Logger 中间件
	//r.Use(Logger())

	// 使用 Recovery 中间件
	r.Use(gin.Recovery())

	// 路由组中添加中间件
	// authorized := r.Group("/", AuthRequired())
	authorized := r.Group("/")

	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", Login)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8082")
}

func Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

	return
}
