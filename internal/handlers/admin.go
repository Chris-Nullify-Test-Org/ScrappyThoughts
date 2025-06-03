package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"scrappythoughts.com/scrappythoughts-repo/internal/models"
)

func AdminDashboard(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var totalUsers, totalPosts, totalComments int64

		db.Model(&models.User{}).Count(&totalUsers)
		db.Model(&models.Post{}).Count(&totalPosts)
		db.Model(&models.Comment{}).Count(&totalComments)

		response := map[string]interface{}{
			"total_users":    totalUsers,
			"total_posts":    totalPosts,
			"total_comments": totalComments,
		}

		json.NewEncoder(w).Encode(response)
	}
}

func GetModeratorPosts(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")
		flagged := r.URL.Query().Get("flagged")

		query := db.Model(&models.Post{})

		if status != "" {
			query = query.Where("status = ?", status)
		}

		if flagged == "true" {
			query = query.Where("is_flagged = ?", true)
		}

		var posts []models.Post
		if err := query.Preload("Author").Find(&posts).Error; err != nil {
			http.Error(w, "Error fetching posts", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(posts)
	}
}

func FlagPost(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		var req struct {
			Reason string `json:"reason"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		var post models.Post
		if err := db.First(&post, postID).Error; err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		post.IsFlagged = true
		if err := db.Save(&post).Error; err != nil {
			http.Error(w, "Error flagging post", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(post)
	}
}

func UpdatePostStatus(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")
		var req struct {
			Status string `json:"status"`
			Reason string `json:"reason"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		var post models.Post
		if err := db.First(&post, postID).Error; err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		post.Status = req.Status
		if err := db.Save(&post).Error; err != nil {
			http.Error(w, "Error updating post status", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(post)
	}
}

func BanUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "id")
		var req struct {
			Reason       string `json:"reason"`
			DurationDays int    `json:"duration_days"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		user.IsBanned = true
		if err := db.Save(&user).Error; err != nil {
			http.Error(w, "Error banning user", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value("role").(string)
		if userRole != "admin" {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ModeratorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value("role").(string)
		if userRole != "moderator" && userRole != "admin" {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
