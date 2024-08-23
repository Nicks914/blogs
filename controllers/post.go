package controllers

import (
	"blogs/models"
	"blogs/utils"
	"encoding/json"
	"net/http"
)

// Create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	post.UserID = claims.UserID

	utils.DB.Create(&post)
	json.NewEncoder(w).Encode(post)
}

// Get all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	utils.DB.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

// Get a single post by ID
func GetPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var post models.Post
	utils.DB.First(&post, id)

	if post.ID == 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}

// Update a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id := r.URL.Query().Get("id")
	var post models.Post
	utils.DB.First(&post, id)

	if post.UserID != claims.UserID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	json.NewDecoder(r.Body).Decode(&post)
	utils.DB.Save(&post)
	json.NewEncoder(w).Encode(post)
}

// Delete a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id := r.URL.Query().Get("id")
	var post models.Post
	utils.DB.First(&post, id)

	if post.UserID != claims.UserID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	utils.DB.Delete(&post)
	json.NewEncoder(w).Encode("Post deleted")
}
