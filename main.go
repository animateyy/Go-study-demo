package main

import (
	"ginweb/apis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/login", apis.GetUser)
	r.POST("/register", apis.RegisterUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
