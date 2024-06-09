package handlers

import (
	"database/sql"
	"net/http"
	"post/internals/tools"
	md "post/models"
)

func deletePost(w http.ResponseWriter, post md.Post, db *sql.DB) {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", post.Id)
	if err != nil {
		http.Error(w, "Error while deleting post : "+err.Error(), http.StatusBadRequest)
		return
	}
	tools.WriteResponse(w, "Post well deleted", http.StatusOK)
}
