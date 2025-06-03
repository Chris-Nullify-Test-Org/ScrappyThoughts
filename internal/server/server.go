package server

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
	"scrappythoughts.com/scrappythoughts-repo/internal/database"
	"scrappythoughts.com/scrappythoughts-repo/internal/handlers"
	auth "scrappythoughts.com/scrappythoughts-repo/internal/middleware"
	"scrappythoughts.com/scrappythoughts-repo/internal/models"
	"scrappythoughts.com/scrappythoughts-repo/internal/seed"
)

type Server struct {
	Router *chi.Mux
	db     *gorm.DB
}

var (
	// Global database connection that will be shared across all requests
	globalDB *gorm.DB
)

func New() (*Server, error) {
	// Initialize the database
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	// Auto migrate the schema
	err = db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
		&models.UserLikedPost{},
	)
	if err != nil {
		return nil, err
	}

	// Seed the database with dummy data
	if err := seed.SeedData(db); err != nil {
		return nil, err
	}

	s := &Server{
		Router: chi.NewRouter(),
		db:     db,
	}

	s.setupMiddleware()
	s.setupRoutes()

	return s, nil
}

// GetDB returns the global database connection
func GetDB() *gorm.DB {
	return globalDB
}

func (s *Server) setupMiddleware() {
	s.Router.Use(chimiddleware.Logger)
	s.Router.Use(chimiddleware.Recoverer)
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

func (s *Server) setupRoutes() {
	// API routes
	s.Router.Route("/api/v1", func(r chi.Router) {
		// Public routes
		r.Post("/auth/register", handlers.Register(s.db))
		r.Post("/auth/login", handlers.Login(s.db))
		r.Get("/posts", handlers.GetPosts(s.db))
		r.Get("/posts/{id}", handlers.GetPost(s.db))
		r.Get("/posts/{id}/comments", handlers.GetComments(s.db))

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(auth.AuthMiddleware(s.db))
			r.Post("/auth/logout", handlers.Logout())
			r.Post("/posts", handlers.CreatePost(s.db))
			r.Put("/posts/{id}", handlers.UpdatePost(s.db))
			r.Delete("/posts/{id}", handlers.DeletePost(s.db))
			r.Post("/posts/{id}/like", handlers.LikePost(s.db))
			r.Delete("/posts/{id}/like", handlers.UnlikePost(s.db))
			r.Post("/posts/{id}/comments", handlers.CreateComment(s.db))
		})

		// Admin routes
		r.Group(func(r chi.Router) {
			r.Use(auth.AuthMiddleware(s.db))
			r.Use(handlers.AdminMiddleware)
			r.Get("/admin", handlers.AdminDashboard(s.db))
		})

		// Moderator routes
		r.Group(func(r chi.Router) {
			r.Use(auth.AuthMiddleware(s.db))
			r.Use(handlers.ModeratorMiddleware)
			r.Get("/moderator/posts", handlers.GetModeratorPosts(s.db))
			r.Post("/moderator/posts/{id}/flag", handlers.FlagPost(s.db))
			r.Put("/moderator/posts/{id}/status", handlers.UpdatePostStatus(s.db))
			r.Post("/moderator/users/{id}/ban", handlers.BanUser(s.db))
		})
	})
}
