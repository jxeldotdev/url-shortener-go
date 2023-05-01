package models

import (
	"github.com/jxeldotdev/url-backend/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"-"`
	Urls     []Url  `gorm:"foreignKey:UserId" json:"urls"`
	IsAdmin  bool   `json:"is_admin"`
	RoleId   Role   `gorm:"foreignKey:Id"`
}

func (user *User) Save() (*User, error) {
	pw, err := HashPassword(user.Password)
	user.Password = pw
	err = db.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
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

func UserLoginCheck(username string, password string) (bool, error) {
	user := User{}

	err := db.Database.Model(User{}).Where("username = ?", username).Take(&user).Error

	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return false, err
	}
	return true, nil
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
