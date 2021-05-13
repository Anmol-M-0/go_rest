package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post Home Page",
	})
}
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Home page",
	})
}
func PutHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Put Home Page",
	})
}
func DeleteHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete Home Page",
	})
}
func PatchHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Patch Home Page",
	})
}
func HeadHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Head Home Page",
	})
	// fmt.Println("in head")
}
func OptionsHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Options Home Page",
	})
}
func main() {
	fmt.Println("hello world")
	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.PUT("/", PutHomePage)
	r.DELETE("/", DeleteHomePage)
	r.PATCH("/", PatchHomePage)
	r.HEAD("/", HeadHomePage)
	r.OPTIONS("/", OptionsHomePage)
	r.Run()
}
