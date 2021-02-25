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

type GetBookOpt struct {
	ID string `json:"id" binding:"required"`
}

func (bc *BookController) GetBook(c *gin.Context) {
	opt := GetBookOpt{}
	err := c.ShouldBind(&opt)
	if err != nil {
		c.JSON(400, gin.H{
			"data": nil,
			"msg": "Bad bookID",
		})
		return
	}
	bookID, err := strconv.Atoi(opt.ID)
	book, err := bc.service.FindBookById(bookID)
	if err != nil {
		c.JSON(404, gin.H{
			"data": nil,
			"msg": "Book Not Found",
		})
		return
	}
	c.JSON(200, book)
}

func (bc *BookController) GetBooks(c *gin.Context) {
	books := bc.service.GetBooks()
	c.JSON(200, books)
}

func (bc *BookController) Bind(r *gin.RouterGroup) {
	r.POST("/getBooks", middleware.AuthRequire(1), bc.GetBooks)
	r.POST("/getBook", middleware.AuthRequire(1), bc.GetBook)
}

func NewBookController(service service.BookService) *BookController {
	return &BookController{
		service: service,
	}
}
