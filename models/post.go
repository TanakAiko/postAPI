package models

import (
	"database/sql"
	"time"
)

type Post struct {
	Id       int
	UserId   int
	Content  string
	CreateAt time.Time
}

func (post *Post) CreatePost(db *sql.DB) error {

	return nil
}

func (post *Post) GetPost(db *sql.DB) error {

	return nil
}
