package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string    `gorm:"uniqueIndex;not null"`
	Email      string    `gorm:"uniqueIndex;not null"`
	Password   string    `gorm:"not null"`
	Role       string    `gorm:"default:user"`
	IsBanned   bool      `gorm:"default:false"`
	Posts      []Post    `gorm:"foreignKey:AuthorID"`
	Comments   []Comment `gorm:"foreignKey:AuthorID"`
	LikedPosts []Post    `gorm:"many2many:user_liked_posts;"`
}

type Post struct {
	gorm.Model
	Title      string    `gorm:"not null"`
	Content    string    `gorm:"not null"`
	AuthorID   uint      `gorm:"not null"`
	Author     User      `gorm:"foreignKey:AuthorID"`
	LikesCount int       `gorm:"default:0"`
	IsFlagged  bool      `gorm:"default:false"`
	Status     string    `gorm:"default:active"`
	Comments   []Comment `gorm:"foreignKey:PostID"`
	LikedBy    []User    `gorm:"many2many:user_liked_posts;"`
}

type Comment struct {
	gorm.Model
	Content  string `gorm:"not null"`
	PostID   uint   `gorm:"not null"`
	Post     Post   `gorm:"foreignKey:PostID"`
	AuthorID uint   `gorm:"not null"`
	Author   User   `gorm:"foreignKey:AuthorID"`
}

type UserLikedPost struct {
	UserID    uint `gorm:"primaryKey"`
	PostID    uint `gorm:"primaryKey"`
	CreatedAt time.Time
}
