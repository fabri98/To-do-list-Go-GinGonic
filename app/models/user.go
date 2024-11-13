package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Task     []Task `gorm:"foreingKey:UserID"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	gorm.Model
}

func (User) TableName() string {
	return "users"
}

// SetPassword encripta la contraseña y la asigna al campo Password
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compara la contraseña en texto plano con la encriptada
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
