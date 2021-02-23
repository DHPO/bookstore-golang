package controller

import (
	"bookstore/bookstore/middleware"
	"bookstore/bookstore/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service service.BookService
}

func (bc *BookController) GetBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"data": nil,
			"msg": "Bad bookID",
		})
		return
	}
	book, err := bc.service.FindBookById(bookID)
	if err != nil {
		c.JSON(404, gin.H{
			"data": nil,
			"msg": "Book Not Found",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": book,
	})
}

func (bc *BookController) GetBooks(c *gin.Context) {
	books := bc.service.GetBooks()
	c.JSON(200, gin.H{
		"data": books,
	})
}

func (bc *BookController) Bind(r *gin.RouterGroup) {
	r.GET("/getBooks", middleware.AuthRequire(1), bc.GetBooks)
	r.GET("/getBook/:id", middleware.AuthRequire(1), bc.GetBook)
}

func NewBookController(service service.BookService) *BookController {
	return &BookController{
		service: service,
	}
}
