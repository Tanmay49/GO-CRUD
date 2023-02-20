package Routes

import (
	"main/Init"
	"main/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Trial(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Worldsss",
	})
}

func Insert(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	post := Models.Post{Title: body.Title, Body: body.Body}
	result := Init.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": post})

}

func GetALlposts(c *gin.Context) {
	var posts []Models.Post
	Init.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"message": posts})
}

func GetPostID(c *gin.Context) {
	id := c.Param("id")

	var post Models.Post
	Init.DB.First(&post, id)
	c.JSON(http.StatusOK, gin.H{
		"message": post})

}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var post Models.Post

	Init.DB.First(&post, id)

	Init.DB.Model(&post).Updates(Models.Post{Title: body.Title, Body: body.Body})
	c.JSON(http.StatusOK, gin.H{
		"message": post})

}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	post := Models.Post{}
	Init.DB.Delete(&post, id)

	var newposts []Models.Post

	Init.DB.Find(&newposts)
	c.JSON(http.StatusOK, gin.H{
		"message": newposts})
}
