package app

import (
	"bookstore/bookstore/controller"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/gin-contrib/cors"

	// "bookstore/bookstore/middleware"
	"github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
)

type Application struct {
	viper *viper.Viper
	engine *gin.Engine
}

func NewApplication(
	viper *viper.Viper,
	bc *controller.BookController,
	uc *controller.UserController,
	lc *controller.LoginController,
) *Application {
	app := &Application{
		engine: gin.Default(),
		viper: viper,
	}
	// set middleware
	store := cookie.NewStore([]byte("secretKeyssa;kldjhasodhfa;"))
	app.engine.Use(sessions.Sessions("gin-session", store))
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("*")
	corsConfig.AddAllowMethods("*")
	corsConfig.AllowOrigins = append(corsConfig.AllowOrigins, "http://10.0.0.75:3000")
	corsConfig.AllowCredentials = true
	corsConfig.AllowWildcard = true
	corsMiddleware := cors.New(corsConfig)
	app.engine.Use(corsMiddleware)

	// set controller
	bc.Bind(app.engine.Group(""))
	uc.Bind(app.engine.Group(""))
	lc.Bind(app.engine.Group(""))
	return app
}

func (app *Application) Serve() {
	addr := fmt.Sprintf("%s:%s", app.viper.GetString("server.host"), app.viper.GetString("server.port"))
	gin.SetMode(app.viper.GetString("server.mode"))
	app.engine.Run(addr)
}
