package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"scrappythoughts.com/scrappythoughts-repo/internal/database"
	"scrappythoughts.com/scrappythoughts-repo/internal/models"
)

func AuthMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
				http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
				return
			}

			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				return []byte("your-secret-key"), nil // In production, use environment variable
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			userID := uint(claims["user_id"].(float64))

			// Use the shared database connection
			sharedDB := database.GetDB()

			// Check if user is banned
			var user models.User
			if err := sharedDB.First(&user, userID).Error; err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			isBanned := user.IsBanned

			ctx := context.WithValue(r.Context(), "user_id", userID)
			ctx = context.WithValue(ctx, "role", claims["role"].(string))
			ctx = context.WithValue(ctx, "is_banned", isBanned)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
