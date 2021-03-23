package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 利用组来收敛鉴权
	v1 := r.Group("", func(c *gin.Context) {
		fmt.Println("auth login success")
	})
	v1.GET("/users", func(c *gin.Context) {
		c.String(200, "/v1/users")
	})

	v1.GET("/products", func(c *gin.Context) {
		c.String(200, "/v1/products")
	})
	ginAddr := "127.0.0.1:3452"
	r.Run(ginAddr)
}
