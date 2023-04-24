package models

import (
	"fmt"
	"time"

	"github.com/jxeldotdev/url-backend/internal/db"
)

type Url struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	LongUrl   string    `gorm:"not null" binding:"required" json:"long_url"`
	ShortUrl  string    `gorm:"unique" json:"short_url"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UserId    uint      `json:"user_id"`
}

func (url *Url) Save() (*Url, error) {
	err := db.Database.Create(&url).Error
	if err != nil {
		return &Url{}, err
	}
	return url, nil
}

func FindUrlByShortUrl(shortUrl string) (Url, error) {
	var urlInDb Url
	err := db.Database.Where("short_url = ?", shortUrl).First(&urlInDb).Error
	if err != nil {
		return Url{}, err
	}
	fmt.Printf("FindUrlByShortUrl: +%v", urlInDb)
	return urlInDb, nil
}

func (url *Url) GetAll() ([]Url, error) {
	var urls []Url
	err := db.Database.Limit(100).Find(&urls).Error
	if err != nil {
		return urls, err
	}
	return urls, nil
}
