package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nseve/first-go-restapi/internal/service"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler(s *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: s}
}

func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Ttile string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// http.Error(w, "invalid request body", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	project, err := h.service.Create(req.Ttile)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(project)
	writeJSON(w, http.StatusCreated, project)
}

func (h *ProjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	projects, err := h.service.GetAll()
	if err != nil {
		// http.Error(w, "failed to get projects", http.StatusInternalServerError)
		writeError(w, http.StatusInternalServerError, "Failed to get projects")
		return
	}

	// json.NewEncoder(w).Encode(projects)
	writeJSON(w, http.StatusOK, projects)
}

func (h *ProjectHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "invalid id", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	project, err := h.service.GetByID(uint(id))
	if err != nil {
		// http.Error(w, "project not found", http.StatusNotFound)
		writeError(w, http.StatusNotFound, "Project not found")
		return
	}

	// json.NewEncoder(w).Encode(project)
	writeJSON(w, http.StatusOK, project)
}

func (h *ProjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "invalid id", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var req struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// http.Error(w, "invalid body", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	project, err := h.service.Update(uint(id), req.Title)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	// json.NewEncoder(w).Encode(project)
	writeJSON(w, http.StatusOK, project)
}

func (h *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "invalid id", http.StatusBadRequest)
		writeError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		// http.Error(w, "failed to delete project", http.StatusInternalServerError)
		writeError(w, http.StatusInternalServerError, "Failed to delete project")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
