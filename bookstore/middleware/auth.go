package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequire(levelRequire int) gin.HandlerFunc {
	return func (c *gin.Context) {
		session := sessions.Default(c)
		level := session.Get("level")
		if level == nil || level.(int) > levelRequire {
			c.AbortWithStatusJSON(403, gin.H{
				"msg": "Permission Denied",
			})
		}
		c.Next()
	}
}
