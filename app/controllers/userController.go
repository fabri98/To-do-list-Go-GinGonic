package controllers

import (
	"gin-mvc/config"
	"gin-mvc/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func GetUsers(c *gin.Context) {
	session := sessions.Default(c)
	userName := session.Get("userName")

	// Verificar si la sesión ha expirado
	userID := session.Get("userID")
	if userID == nil {
		// Si `userID` es nulo, la sesión ha expirado o no existe
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "La sesión ha expirado. Por favor, inicia sesión nuevamente.",
		})
		return
	}

	var users []models.User
	config.DB.Find(&users)

	c.HTML(http.StatusOK, "users.html", gin.H{
		"users":     users,
		"userName":  userName,
		"csrfField": csrf.TemplateField(c.Request),
	})
}

func UpdateUser(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	id := c.Param("id")

	if userID == nil {
		// Si no hay usuario en la sesión, devuelve Unauthorized
		c.Status(http.StatusUnauthorized)
		return
	}
	var user models.User
	// Buscar usuario existente
	if err := config.DB.First(&user, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "users.html", gin.H{
			"error": "Usuario no encontrado",
		})
		return
	}

	user.Name = c.PostForm("Name")
	user.Email = c.PostForm("Email")

	// esto es para que se actualice el nombre en la bienvenida, en resumen solo es algo visual
	if userID == user.ID {
		session.Set("userName", user.Name)
		session.Save()
	}

	// Actualizar usuario
	config.DB.Save(&user)

	c.Redirect(http.StatusSeeOther, "/api/users")
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	session := sessions.Default(c)
	userID := session.Get("userID")

	// Verificar si la sesión ha expirado
	if userID == nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "La sesión ha expirado. Por favor, inicia sesión nuevamente.",
		})
		return
	}
	// Intentar eliminar el usuario
	result := config.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "users.html", gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/api/users")
}
