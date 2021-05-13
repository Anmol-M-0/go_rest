package main

import (
	"errors"
	"fmt"
	"io/ioutil"

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

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
func ReadPost(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": errors.New("an error occured, that's all we know"),
		})
	} else {
		c.JSON(200, gin.H{
			"message": string(value),
		})
	}

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
	//query strings, path params body of post request

	//query strings
	r.GET("/query", QueryStrings)             // /query?name=anmol&age=24
	r.GET("/path/:name/:age", PathParameters) // /path/anmol/24

	//reading the body of a post request
	r.POST("/read", ReadPost) //read from te body of post req
	r.Run()
}
