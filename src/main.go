package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

//a simple post end point
//returns a json {message: {message}}
func PostHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post Home Page",
	})
}

//a simple get end point
//returns a json {message: {message}}
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Home page",
	})
}

//a simple put end point
//returns a json {message: {message}}
func PutHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Put Home Page",
	})
}

//a simple delete end point
//returns a json {message: {message}}
func DeleteHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete Home Page",
	})
}

//a simple patch end point
//returns a json {message: {message}}
func PatchHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Patch Home Page",
	})
}

//a simple head end point
func HeadHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Head Home Page",
	})
	// fmt.Println("in head")
}

//a simple options end point
//returns a json {message: {message}}
func OptionsHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Options Home Page",
	})
}

//a get end point, /query?name={name}&age={age}
//to test out query params
//returns a json {name: {name}, age:{age}}
func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

//a get end point, /path/:{name}/:{age}
//to test out path params
//returns a json {name: {name}, age:{age}}
func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

//a post end point, /read, that reads the content of the body and returns
//that in the form of json {"message":{body}}
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
