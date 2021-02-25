package middleware

import (
	"bookstore/bookstore/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequire(levelRequire int) gin.HandlerFunc {
	return func (c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("auth")
		if v == nil{
			c.JSON(403, gin.H{
				"msg": "Permission Denied",
			})
			return
		}
		auth, ok := v.(model.UserAuth)
		if !ok || auth.UserType > levelRequire {
			c.JSON(403, gin.H{
				"msg": "Permission Denied",
			})
			return
		}
		c.Next()
	}
}
