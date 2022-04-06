package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", GetHome)
	router.GET("/artikel/:judul", GetArtikel)
	router.POST("/artikel", PostArtikel)

	router.Run()
}

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":   "Berhasil",
		"messages": "Berhasil Tampil Home",
	})
}

func GetArtikel(c *gin.Context) {
	judul := c.Param("judul")
	c.JSON(200, gin.H{
		"status":   "Berhasil",
		"messages": judul,
	})
}

func PostArtikel(c *gin.Context) {
	judul := c.PostForm("judul")
	desc := c.PostForm("desc")

	c.JSON(200, gin.H{
		"status":   "Berhasil",
		"messages": desc,
		"title":    judul,
	})
}
