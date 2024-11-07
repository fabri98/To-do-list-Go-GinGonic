package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Conexion a DB MySQL
func ConnectToDB() {
	var err error

	// Para hacer la conexion a una DB mySql, el dsn debe tener este formato
	dsn := "root:@tcp(localhost:3306)/gin_crud?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}
}
