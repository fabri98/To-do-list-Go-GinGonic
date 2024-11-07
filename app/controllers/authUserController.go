package controllers

import (
	"gin-mvc/config"
	"gin-mvc/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func ShowLogin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("userID") != nil {
		c.Redirect(http.StatusSeeOther, "/api/users")
		return
	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"csrfField": csrf.TemplateField(c.Request), // Asegúrate de que esto esté bien
	})
}

func Login(c *gin.Context) {
	inputEmail := c.PostForm("Email")
	inputPassword := c.PostForm("Password")

	var user models.User
	if err := config.DB.Where("email = ?", inputEmail).First(&user).Error; err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Email o contraseña incorrecto!"})
		return
	}

	if !user.CheckPassword(inputPassword) {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Email o contraseña incorrecto!"})
		return
	}

	// Crear sesión
	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Set("userName", user.Name)

	session.Save()

	c.Redirect(http.StatusSeeOther, "/api/users")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusSeeOther, "/")
}
