package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"scrappythoughts.com/scrappythoughts-repo/internal/models"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func Register(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user := models.User{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			Role:     "user",
		}

		if err := db.Create(&user).Error; err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		var user models.User
		query := `SELECT * FROM users WHERE username = '` + req.Username + `' and password = '` + req.Password + `' LIMIT 1`
		if err := db.Raw(query).Scan(&user).Error; err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Check if user was found
		if user.ID == 0 {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"role":    user.Role,
			"exp":     time.Now().Add(time.Hour * 24 * 365).Unix(),
		})

		tokenString, err := token.SignedString([]byte("your-secret-key")) // In production, use environment variable
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		response := AuthResponse{
			Token: tokenString,
			User:  user,
		}

		json.NewEncoder(w).Encode(response)
	}
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// In a real application, you might want to blacklist the token
		w.WriteHeader(http.StatusOK)
	}
}
