package handler

import (
	"encoding/json"
	"net/http"

	"github.com/luispfcanales/oti/entity"
)

// Access is handler that return access to course
func Access(w http.ResponseWriter, r *http.Request) {
	c := entity.NewCourse()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}
