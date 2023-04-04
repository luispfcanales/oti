package handler

import (
	"encoding/json"
	"net/http"
)

// Credentials is handler that return access to course
func Credentials(w http.ResponseWriter, r *http.Request) {
	c := newCourse()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&c)
}
