package middlewares

import (
	"github.com/gin-gonic/gin"
)

func MethodOverride() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.PostFormValue("_method")
		if c.Request.Method == "POST" {
			if method == "PUT" || method == "DELETE" {
				c.Request.Method = method
			}
		}
		c.Next()
	}
}
