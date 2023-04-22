package models

import (
	"time"

	"github.com/jxeldotdev/url-backend/internal/db"
	"gorm.io/gorm"
)

/*type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	Password string `gorm:"size:255;not null" json:"-"`
	Urls     []Url  `gorm:"foreignKey:UserId"`
	IsAdmin  bool
}
*/

type Url struct {
	gorm.Model
	Id        uint
	LongUrl   string    `gorm:"size:255;not null" json:"long_url"`
	ShortUrl  string    `gorm:"size:10;not null;unique" json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uint
}

func (url *Url) Save() (*Url, error) {
	err := db.Database.Create(&url).Error
	if err != nil {
		return &Url{}, err
	}
	return url, nil
}
