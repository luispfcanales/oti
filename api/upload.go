package handler

import (
	"encoding/json"
	"net/http"

	"github.com/luispfcanales/oti/entity"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	var pc entity.Computer

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	json.NewEncoder(w).Encode(pc)
}
