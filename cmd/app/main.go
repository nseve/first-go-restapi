package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nseve/first-go-restapi/internal/db"
	"github.com/nseve/first-go-restapi/internal/handler"
	"github.com/nseve/first-go-restapi/internal/repository"
	"github.com/nseve/first-go-restapi/internal/service"
)

func main() {
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

	projectService := service.NewProjectService(projectRepo)
	taskService := service.NewTaskService(taskRepo, projectRepo)

	projectHandler := handler.NewProjectHandler(projectService)
	taskHandler := handler.NewTaskHandler(taskService)

	r := chi.NewRouter()

	r.Route("/projects", func(r chi.Router) {
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
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", taskHandler.GetByID)
			r.Put("/", taskHandler.Update)
			r.Delete("/", taskHandler.Delete)
		})
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
