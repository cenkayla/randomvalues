package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Value struct {
	UUID      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Number    int       `gorm:"-" json:"number,omitempty"`
	Name      string
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func generateWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateURL(length int) string {
	return generateWithCharset(length, charset)
}
