package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"scrappythoughts.com/scrappythoughts-repo/internal/models"
)

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetPosts(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 10
		}

		var posts []models.Post
		var total int64

		db.Model(&models.Post{}).Count(&total)
		offset := (page - 1) * limit

		if err := db.Preload("Author").Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
			http.Error(w, "Error fetching posts", http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"posts":       posts,
			"total":       total,
			"page":        page,
			"total_pages": (total + int64(limit) - 1) / int64(limit),
		}

		json.NewEncoder(w).Encode(response)
	}
}

func CreatePost(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		userID := r.Context().Value("user_id").(uint)
		isBanned := r.Context().Value("is_banned").(bool)

		// Check if user is banned
		if isBanned {
			http.Error(w, "Banned users cannot create posts", http.StatusForbidden)
			return
		}

		post := models.Post{
			Title:    req.Title,
			Content:  req.Content,
			AuthorID: userID,
		}

		if err := db.Create(&post).Error; err != nil {
			http.Error(w, "Error creating post", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	}
}

func GetPost(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		var post models.Post

		if err := db.Preload("Author").Preload("Comments.Author").First(&post, postID).Error; err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(post)
	}
}

func UpdatePost(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		userRole := r.Context().Value("role").(string)
		isBanned := r.Context().Value("is_banned").(bool)

		// Check if user is banned
		if isBanned {
			http.Error(w, "Banned users cannot update posts", http.StatusForbidden)
			return
		}

		var post models.Post
		if err := db.First(&post, postID).Error; err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		if userRole != "user" && userRole != "admin" && userRole != "moderator" {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}

		var req PostRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		post.Title = req.Title
		post.Content = req.Content

		if err := db.Save(&post).Error; err != nil {
			http.Error(w, "Error updating post", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(post)
	}
}

func DeletePost(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		userID := r.Context().Value("user_id").(uint)
		userRole := r.Context().Value("role").(string)
		isBanned := r.Context().Value("is_banned").(bool)

		// Check if user is banned
		if isBanned {
			http.Error(w, "Banned users cannot delete posts", http.StatusForbidden)
			return
		}

		var post models.Post
		if err := db.First(&post, postID).Error; err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		if post.AuthorID != userID && userRole != "admin" && userRole != "moderator" {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}

		if err := db.Delete(&post).Error; err != nil {
			http.Error(w, "Error deleting post", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func LikePost(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		userID := r.Context().Value("user_id").(uint)

		var post models.Post
		if err := db.First(&post, postID).Error; err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if err := db.Model(&post).Association("LikedBy").Append(&user); err != nil {
			http.Error(w, "Error liking post", http.StatusInternalServerError)
			return
		}

		post.LikesCount++
		db.Save(&post)

		w.WriteHeader(http.StatusOK)
	}
}

func UnlikePost(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		userID := r.Context().Value("user_id").(uint)

		var post models.Post
		if err := db.First(&post, postID).Error; err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if err := db.Model(&post).Association("LikedBy").Delete(&user); err != nil {
			http.Error(w, "Error unliking post", http.StatusInternalServerError)
			return
		}

		post.LikesCount--
		db.Save(&post)

		w.WriteHeader(http.StatusOK)
	}
}
