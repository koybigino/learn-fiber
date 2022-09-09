package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
	gorm.Model
}

type User struct {
	Id         int       `json:"id" gorm:"primaryKey;not null"`
	Email      string    `json:"email" gorm:"unique;not null" validate:"email"`
	Password   string    `json:"password" gorm:"unique;not null"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
}
