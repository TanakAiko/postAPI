package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"post/internals/tools"
	md "post/models"
)

func getOnePost(w http.ResponseWriter, post md.Post, db *sql.DB) {
	if err := post.GetOnePost(db); err != nil {
		http.Error(w, "Error while getting post : "+err.Error(), http.StatusBadRequest)
		return
	}
	tools.WriteResponse(w, post, http.StatusOK)
}

func getAllPost(w http.ResponseWriter, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM posts ORDER BY createdAt DESC")
	if err != nil {
		http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []md.Post{}
	for rows.Next() {
		var post md.Post
		var categorieJSON string
		if err := rows.Scan(&post.Id, &post.UserId, &categorieJSON, &post.Content, &post.Img, &post.NbrLike, &post.NbrDislike, &post.CreateAt); err != nil {
			http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
			return
		}
		if err = json.Unmarshal([]byte(categorieJSON), &post.Categorie); err != nil {
			http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	tools.WriteResponse(w, posts, http.StatusOK)
}
