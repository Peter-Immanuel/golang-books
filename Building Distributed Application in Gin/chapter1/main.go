package main

import "github.com/gin-gonic/gin"

func mainHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}

func main() {
	router := gin.Default()
	router.GET("/", mainHandler)
	router.Run()
}
