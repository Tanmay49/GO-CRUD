package main

import (
	"main/Init"
	"main/Models"
	"main/Routes"

	"github.com/gin-gonic/gin"
)

func init() {
	Init.LoadEnv()
	Init.ConnectDB()
}

func main() {
	Init.DB.AutoMigrate(&Models.Post{})

	router := gin.Default()
	router.GET("/", Routes.Trial)
	router.GET("/getallposts", Routes.GetALlposts)
	router.GET("/getpost/:id", Routes.GetPostID)
	router.POST("/insert", Routes.Insert)
	router.PUT("/update/:id", Routes.UpdatePost)
	router.DELETE("/delete/:id", Routes.DeletePost)

	router.Run()
}
