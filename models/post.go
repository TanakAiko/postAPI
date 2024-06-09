package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type Post struct {
	Id       int
	UserId   int    `json:"userID"`
	Content  string `json:"content"`
	CreateAt time.Time
}

func (post *Post) CreatePost(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	content, err := os.ReadFile("./databases/sqlRequests/insertNewPost.sql")
	if err != nil {
		return err
	}

	// This code snippet is preparing a SQL statement for execution within a transaction.
	stmt, err := tx.Prepare(string(content))
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		post.UserId,
		post.Content,
		time.Now().Format(time.RFC3339),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		return err
	}

	return err
}

func (post *Post) GetOnePost(db *sql.DB) error {
	err := db.QueryRow("SELECT userId, content, createdAt FROM posts WHERE id = ?", post.Id).Scan(
		&post.UserId,
		&post.Content,
		&post.CreateAt,
	)
	return err
}
