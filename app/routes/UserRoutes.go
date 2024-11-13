package routes

import (
	"gin-mvc/controllers"
	"gin-mvc/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	// Rutas públicas
	// Login routes
	router.GET("/", controllers.ShowLogin)
	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)
	// Register routes
	router.GET("/register", controllers.ShowRegister)
	router.POST("/register", controllers.RegisterUser)

	// Rutas protegidas
	routes := router.Group("/api",
		middlewares.AuthRequired(),
		middlewares.CsrfToken(),
	)
	{
		routes.GET("/users", controllers.GetUsers)
		routes.PUT("/users/:id", controllers.UpdateUser)
		routes.POST("/users/:id/update", controllers.UpdateUser)
		routes.DELETE("/users/:id", controllers.DeleteUser)
		routes.POST("/users/:id/delete", controllers.DeleteUser)

		routes.GET("/tasks", controllers.ListTasks)
		routes.GET("/tasks/:id/update", controllers.UpdateTaskForm)
		routes.POST("/tasks/:id/update", controllers.UpdateTask)
	}

}
