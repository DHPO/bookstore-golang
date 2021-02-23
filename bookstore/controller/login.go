package controller

import (
	"bookstore/bookstore/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

type LoginController struct {
	service service.UserService
}

type LoginForm struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (lc *LoginController) Login(c *gin.Context) {
	var form LoginForm
	err := c.ShouldBindJSON(&form)
	if err != nil{
		c.JSON(400, gin.H{
			"status": -100,
			"msg": "Bad Request",
		})
		return
	}
	userAuth, err := lc.service.CheckUser(form.Username, form.Password)
	if err != nil {
		c.JSON(404, gin.H{
			"status": -100,
			"msg": "用户名或密码错误",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("auth", 1)
	err = session.Save()
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"status": 0,
		"msg": "登录成功",
		"data": userAuth,
	})
}

func (lc *LoginController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{
		"status": 0,
		"msg": "logout",
	})
}

func (lc *LoginController) CheckSession(c *gin.Context) {
	session := sessions.Default(c)
	auth := session.Get("auth")
	if auth == nil {
		c.JSON(200, gin.H{
			"status": -101,
			"msg": "登录失效",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 0,
		"msg": "登录成功",
		"data": auth,
	})
}

func (lc *LoginController) Bind(r *gin.RouterGroup) {
	r.POST("/login", lc.Login)
	r.POST("/logout", lc.Logout)
	r.POST("/checkSession", lc.CheckSession)
}

func NewLoginController(service service.UserService) *LoginController {
	return &LoginController{
		service: service,
	}
}