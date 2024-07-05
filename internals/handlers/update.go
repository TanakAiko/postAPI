package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"post/internals/tools"
	md "post/models"
	"strconv"
)

func updateLike(w http.ResponseWriter, post md.Post, db *sql.DB) {
	query := `
        UPDATE posts
        SET nbrLike = ?, nbrDislike = ?
        WHERE id = ?;
    `

	fmt.Printf("\n\npost.Id: %v\n\n", post.Id)
	fmt.Printf("post.NbrLike: %v\n", post.NbrLike)
	fmt.Printf("post.NbrDislike: %v\n", post.NbrDislike)

	result, err := db.Exec(query, post.NbrLike, post.NbrDislike, post.Id)
	if err != nil {
		fmt.Println("ERROR 1")
		http.Error(w, "Error while deleting post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("ERROR 2")
		http.Error(w, "Error while checking rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("ERROR 3")
		http.Error(w, "No post found with ID: "+strconv.Itoa(post.Id), http.StatusBadRequest)
		return
	}

	tools.WriteResponse(w, "Post well updated", http.StatusOK)
}
