package controller

import (
	"bookstore/bookstore/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func (uc *UserController) CheckUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := uc.service.CheckUser(username, password)
	if err != nil {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
		return
	}
	c.JSON(200, user)
}

func (uc *UserController) Bind(r *gin.RouterGroup) {
	r.GET("/checkUser", uc.CheckUser)
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}
