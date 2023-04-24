package helper

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jxeldotdev/url-backend/models"
)

var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345678910"
const length int = 6

func genUrlStr() string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}
	fmt.Printf("Generated URL: %s", string(b))
	return string(b)
}

func GenUniqueShortUrl() string {
	var str string
	for {
		str = genUrlStr()
		urlExists, _ := models.FindUrlByShortUrl(str)
		if urlExists.ShortUrl == "" {
			break
		}
	}
	return str
}
