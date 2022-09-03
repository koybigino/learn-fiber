package main

import "gorm.io/gorm"

type Post struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
	gorm.Model
}
