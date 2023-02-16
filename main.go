package main

import (
	"main/Init"
	"main/Routes"

	"github.com/gin-gonic/gin"
)

func init() {
	Init.LoadEnv()
	Init.ConnectDB()
}

func main() {
	router := gin.Default()
	router.GET("/", Routes.Trial)
	router.Run("localhost:8080")
}
