package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/gin-bookstore/controllers"
	"github.com/mrinjamul/gin-bookstore/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})

	})

	r.Run()

}
