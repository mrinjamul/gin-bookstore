package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/gin-bookstore/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})

	})

	r.Run()

}
