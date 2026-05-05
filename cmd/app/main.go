package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/nseve/first-go-restapi/internal/db"
	"github.com/nseve/first-go-restapi/internal/handler"
	"github.com/nseve/first-go-restapi/internal/middleware"
	"github.com/nseve/first-go-restapi/internal/repository"
	"github.com/nseve/first-go-restapi/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using system env")
	}

	database, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.RunMigrations(database)
	if err != nil {
		log.Fatal(err)
	}

	projectRepo := repository.NewProjectRepository(database)
	taskRepo := repository.NewTaskRepository(database)
	userRepo := repository.NewUserRepository(database)

	projectService := service.NewProjectService(projectRepo)
	taskService := service.NewTaskService(taskRepo, projectRepo)
	authService := service.NewAuthService(userRepo, os.Getenv("JWT_SECRET"))

	projectHandler := handler.NewProjectHandler(projectService)
	taskHandler := handler.NewTaskHandler(taskService)
	authHandler := handler.NewAuthHandler(authService)

	r := chi.NewRouter()

	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
	})

	jwtSecret := os.Getenv("JWT_SECRET")

	r.Route("/projects", func(r chi.Router) {
		r.Use(middleware.JWTAuth(jwtSecret))

		r.Get("/", projectHandler.GetAll)
		r.Post("/", projectHandler.Create)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", projectHandler.GetByID)
			r.Put("/", projectHandler.Update)
			r.Delete("/", projectHandler.Delete)
		})

		r.Route("/{projectId}/tasks", func(r chi.Router) {
			r.Get("/", taskHandler.GetByProjectID)
			r.Post("/", taskHandler.Create)
		})
	})

	r.Route("/tasks", func(r chi.Router) {
		r.Use(middleware.JWTAuth(jwtSecret))
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", taskHandler.GetByID)
			r.Put("/", taskHandler.Update)
			r.Delete("/", taskHandler.Delete)
		})
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
