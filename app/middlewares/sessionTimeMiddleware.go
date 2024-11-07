package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SessionTimeoutMiddleware actualiza el tiempo de expiración en cada solicitud
func SessionTimeoutMiddleware(timeout int) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// Verifica si el usuario tiene una sesión activa
		if session.Get("userID") != nil {
			// Renueva el tiempo de expiración de la sesión
			session.Options(sessions.Options{
				MaxAge: timeout, // Tiempo en segundos
			})
			session.Save()
		}

		c.Next()
	}
}
