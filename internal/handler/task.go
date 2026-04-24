package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nseve/first-go-restapi/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	projectIDStr := chi.URLParam(r, "projectId")

	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		// http.Error(w, "invalid id", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var req struct {
		Title    string `json:"title"`
		Duration int    `json:"duration"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// http.Error(w, "invalid body", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	task, err := h.service.Create(uint(projectID), req.Title, req.Duration)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(task)
	writeJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) GetByProjectID(w http.ResponseWriter, r *http.Request) {
	projectIDStr := chi.URLParam(r, "projectId")

	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		// http.Error(w, "invalid id", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	tasks, err := h.service.GetByProjectID(uint(projectID))
	if err != nil {
		// http.Error(w, "failed to get projects", http.StatusInternalServerError)
		writeError(w, http.StatusInternalServerError, "Failed to get projects")
		return
	}

	// json.NewEncoder(w).Encode(tasks)
	writeJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "invalid id", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	task, err := h.service.GetByID(uint(id))
	if err != nil {
		// http.Error(w, "task not found", http.StatusNotFound)
		writeError(w, http.StatusNotFound, "Task not found")
		return
	}

	// json.NewEncoder(w).Encode(task)
	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "invalid id", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var req struct {
		Title    string `json:"title"`
		Duration int    `json:"duration"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// http.Error(w, "invalid body", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	task, err := h.service.Update(uint(id), req.Title, req.Duration)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	// json.NewEncoder(w).Encode(task)
	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "invalid id", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		// http.Error(w, "failed to delete task", http.StatusInternalServerError)
		writeError(w, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
