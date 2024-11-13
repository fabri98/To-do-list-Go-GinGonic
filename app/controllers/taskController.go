package controllers

import (
	"gin-mvc/config"
	"gin-mvc/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func ListTasks(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	userName := session.Get(("userName"))
	var tasks []models.Task
	// Filtra las tareas donde el campo `UserID` coincide con el `userID` de la sesión
	if err := config.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "tasks.html", gin.H{
			"error": "Error al recuperar las tareas.",
		})
		return
	}

	c.HTML(http.StatusOK, "tasks.html", gin.H{
		"tasks":     tasks,
		"userName":  userName,
		"csrfField": csrf.TemplateField(c.Request)})
}

func UpdateTaskForm(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
	}

	c.HTML(http.StatusOK, "updateTask.html", gin.H{
		"ID":          task.ID,
		"Title":       task.Title,
		"Description": task.Description,
		"csrfField":   csrf.TemplateField(c.Request),
	})
}
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
	}

	task.Title = c.PostForm("Title")
	task.Description = c.PostForm("Description")

	config.DB.Save(&task)

	c.Redirect(http.StatusSeeOther, "/api/tasks")
}
