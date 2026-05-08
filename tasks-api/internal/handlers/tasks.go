package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"tasks-api/internal/models"
	"tasks-api/internal/storage"
)

type Handler struct {
	Store storage.Storage
}

type errorResponse struct {
	Error string `json:"error"`
}

func New(s storage.Storage) *Handler {
	return &Handler{Store: s}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}

func (h *Handler) TasksCollection(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s %s", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		tasks := h.Store.List()
		writeJSON(w, http.StatusOK, tasks)
	case http.MethodPost:
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			writeError(w, http.StatusBadRequest, "Invalid JSON")
			return
		}

		if task.Title == "" {
			writeError(w, http.StatusBadRequest, "Title is required")
			return
		}

		createdTask, err := h.Store.Create(task)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to create task")
			return
		}
		writeJSON(w, http.StatusCreated, createdTask)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s %s", r.Method, r.URL.Path)

	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	switch r.Method {
	case http.MethodGet:
		task, exists := h.Store.Get(id)
		if !exists {
			writeError(w, http.StatusNotFound, "Task not found")
			return
		}
		writeJSON(w, http.StatusOK, task)
	case http.MethodPut:
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			writeError(w, http.StatusBadRequest, "Invalid JSON")
			return
		}

		if task.Title == "" {
			writeError(w, http.StatusBadRequest, "Title is required")
			return
		}

		updatedTask, err := h.Store.Update(id, task)
		if err != nil {
			writeError(w, http.StatusNotFound, "Task not found")
			return
		}
		writeJSON(w, http.StatusOK, updatedTask)
	case http.MethodDelete:
		if err := h.Store.Delete(id); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to delete task")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
