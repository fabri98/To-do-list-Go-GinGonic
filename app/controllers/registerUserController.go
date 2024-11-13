package controllers

import (
	"gin-mvc/config"
	"gin-mvc/middlewares"
	"gin-mvc/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func ShowRegister(c *gin.Context) {
	session := sessions.Default(c)

	if userID := session.Get("userID"); userID != nil {
		c.Redirect(http.StatusSeeOther, "/api/tasks")
		return
	}
	c.HTML(http.StatusOK, "register.html", gin.H{
		"csrfField": csrf.TemplateField(c.Request),
	})
}
func RegisterUser(c *gin.Context) {
	var user models.User

	// Binding de los datos del formulario
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Encriptar la contraseña antes de guardar
	passwordInput := c.PostForm("Password") // Capturamos la contraseña desde el formulario
	err := user.SetPassword(passwordInput)  // encriptamos

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encriptando la contraseña"})
		return
	}

	// Guardar el usuario en la base de datos
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando el usuario"})
		return
	}

	session := middlewares.SetSession(c, user.ID, user.Name)
	session.Save()

	// Redireccionar o responder con éxito
	c.Redirect(http.StatusSeeOther, "/api/tasks")
}
