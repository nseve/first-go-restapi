package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nseve/first-go-restapi/internal/middleware"
	"github.com/nseve/first-go-restapi/internal/response"
	"github.com/nseve/first-go-restapi/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		response.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	projectIDStr := chi.URLParam(r, "projectId")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var req struct {
		Title    string `json:"title"`
		Duration int    `json:"duration"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	task, err := h.service.Create(uint(projectID), userID, req.Title, req.Duration)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) GetByProjectID(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		response.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	projectIDStr := chi.URLParam(r, "projectId")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	tasks, err := h.service.GetByProjectID(uint(projectID), userID)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to get projects")
		return
	}

	response.WriteJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		response.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	task, err := h.service.GetByID(uint(id), userID)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, "Task not found")
		return
	}

	response.WriteJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		response.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var req struct {
		Title    string `json:"title"`
		Duration int    `json:"duration"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	task, err := h.service.Update(uint(id), userID, req.Title, req.Duration)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)
	if !ok {
		response.WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	if err := h.service.Delete(uint(id), userID); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
