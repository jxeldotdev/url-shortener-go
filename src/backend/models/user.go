package models

import (
	"html"
	"strings"

	"github.com/jxeldotdev/url-backend/internal/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"-"`
	Urls     []Url  `gorm:"foreignKey:UserId" json:"urls"`
	IsAdmin  bool   `json:"is_admin"`
}

func (user *User) Save() (*User, error) {
	user.BeforeSave(db.Database)
	err := db.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
func (user *User) GetAllUsers() ([]User, error) {
	var users []User
	err := db.Database.Limit(100).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (user *User) Update() error {
	if _, err := FindUserById(user.Id); err != nil {
		return err
	}
	db.Database.Save(&user)
	return nil
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := db.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := db.Database.Preload("Urls").Where("Id=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
