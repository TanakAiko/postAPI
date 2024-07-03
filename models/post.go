package models

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"time"
)

type Post struct {
	Id         int       `json:"postID"`
	UserId     int       `json:"userID"`
	Categorie  []string  `json:"categorie"`
	Content    string    `json:"content"`
	Img        string    `json:"img"`
	NbrLike    int       `json:"nbrLike"`
	NbrDislike int       `json:"nbrDislike"`
	CreateAt   time.Time `json:"createAt"`
}

func (post *Post) CreatePost(db *sql.DB) error {
	categorieJSON, err := json.Marshal(post.Categorie)
	if err != nil {
		return err
	}

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
		string(categorieJSON),
		post.Content,
		post.Img,
		post.NbrLike,
		post.NbrDislike,
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
	var categorieJSON string
	err := db.QueryRow("SELECT * FROM posts WHERE id = ?", post.Id).Scan(
		&post.UserId,
		&categorieJSON,
		&post.Content,
		&post.Img,
		&post.NbrLike,
		&post.NbrDislike,
		&post.CreateAt,
	)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(categorieJSON), &post.Categorie)
	return err
}
