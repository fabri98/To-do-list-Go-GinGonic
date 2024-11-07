package main

import (
	"gin-mvc/config"
	"gin-mvc/middlewares"
	"gin-mvc/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.Use(
		middlewares.MethodOverride(),
		sessions.Sessions("mysession", middlewares.SetCookie()),
		middlewares.CsrfToken(),
	)

	routes.UserRouter(router)

	router.Run(":8080")
}
