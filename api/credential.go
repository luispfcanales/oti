package handler

import (
	"encoding/json"
	"net/http"
)

// Credential is handler that return access to course
func Credential(w http.ResponseWriter, r *http.Request) {
	c := newCourse()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&c)
}
