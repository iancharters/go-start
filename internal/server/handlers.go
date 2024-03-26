package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	user, err := s.app.GetUser(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User '%s' does not exist."+err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to fetch user from database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal user: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resp)
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := &CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to unmarshal request body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	user, err := s.app.CreateUser(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal user: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resp)
}
