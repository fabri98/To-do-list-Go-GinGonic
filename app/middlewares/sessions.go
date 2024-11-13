package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSession(c *gin.Context, id uint, name string) sessions.Session {
	session := sessions.Default(c)
	session.Set("userID", id)
	session.Set("userName", name)
	return session
}
