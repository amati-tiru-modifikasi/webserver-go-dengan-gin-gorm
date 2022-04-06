package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", GetHome)

	router.Run()
}

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":   "Berhasil",
		"messages": "Berhasil Tampil Home",
	})
}
