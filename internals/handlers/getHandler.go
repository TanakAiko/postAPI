package handlers

import (
	"database/sql"
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
	rows, err := db.Query("SELECT id, userID, content, createdAt FROM posts ORDER BY createdAt DESC")
	if err != nil {
		http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []md.Post{}
	for rows.Next() {
		var post md.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.Content, &post.CreateAt); err != nil {
			http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	tools.WriteResponse(w, posts, http.StatusOK)
}
