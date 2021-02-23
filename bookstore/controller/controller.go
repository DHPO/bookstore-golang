package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type Controller interface {
	Bind(r *gin.RouterGroup)
}

var ProviderSet = wire.NewSet(
	NewBookController,
	NewUserController,
	NewLoginController,
)