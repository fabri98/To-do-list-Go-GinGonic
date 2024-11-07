package routes

import (
	"gin-mvc/controllers"
	"gin-mvc/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	// Rutas públicas
	router.GET("/", controllers.ShowLogin)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)

	// Rutas protegidas
	routes := router.Group("/api",
		middlewares.AuthRequired(),
		middlewares.SessionTimeoutMiddleware(300),
		middlewares.CsrfToken(),
	)
	{
		routes.GET("/users", controllers.GetUsers)
		routes.POST("/users", controllers.CreateUser)
		routes.PUT("/users/:id", controllers.UpdateUser)
		routes.POST("/users/:id/update", controllers.UpdateUser)
		routes.DELETE("/users/:id", controllers.DeleteUser)
		routes.POST("/users/:id/delete", controllers.DeleteUser)
	}

}
