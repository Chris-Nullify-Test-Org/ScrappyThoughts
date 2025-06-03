package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"scrappythoughts.com/scrappythoughts-repo/internal/models"
)

type CommentRequest struct {
	Content string `json:"content"`
}

func GetComments(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		var comments []models.Comment

		if err := db.Where("post_id = ?", postID).Preload("Author").Find(&comments).Error; err != nil {
			http.Error(w, "Error fetching comments", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(comments)
	}
}

func CreateComment(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		userID := r.Context().Value("user_id").(uint)

		var req CommentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Convert postID from string to uint
		postIDUint, err := strconv.ParseUint(postID, 10, 32)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		comment := models.Comment{
			Content:  req.Content,
			PostID:   uint(postIDUint),
			AuthorID: userID,
		}

		if err := db.Create(&comment).Error; err != nil {
			http.Error(w, "Error creating comment", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(comment)
	}
}
