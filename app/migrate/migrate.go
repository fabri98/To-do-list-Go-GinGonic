package main

import (
	"gin-mvc/config"
	"gin-mvc/models"
)

func init() {
	config.ConnectToDB()
}

func main() {

	config.DB.AutoMigrate(&models.User{}, &models.Task{})
}
