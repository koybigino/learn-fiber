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
	UserId    int    `json:"user_id" gorm:"not null"`
	Users     []User `json:"users" gorm:"many2many:votes"`
	gorm.Model
}
type User struct {
	Id         int       `json:"id" gorm:"primaryKey;not null"`
	Email      string    `json:"email" gorm:"unique;not null" validate:"email"`
	Password   string    `json:"password" gorm:"unique;not null"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
	Posts      []Post    `json:"posts" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
}

type Vote struct {
	PostId int `json:"post_id"`
	UserId int `json:"user_id"`
}

type BodyVote struct {
	PostId int `json:"post_id"`
}

// type Token struct {
// 	AccessToken string `json:"access_token" gorm:"unique"`
// 	TokenType   string `json:"token_type"`
// }

// type TokenData struct {
// 	ID int `json:"id" gorm:"primaryKey"`
// }
