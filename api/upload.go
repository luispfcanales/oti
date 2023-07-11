package handler

import (
	"encoding/json"
	"net/http"

	"github.com/luispfcanales/oti/entity"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	var pc *entity.Computer
	var msg *entity.ReponseApi

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	switch r.Method {
	case http.MethodPost:
		defer r.Body.Close()
		json.NewDecoder(r.Body).Decode(pc)

		msg.Status = http.StatusCreated
		msg.Message = "success"
		msg.Data = pc

		json.NewEncoder(w).Encode(msg)
		break
	default:
		http.NotFound(w, r)
		return
	}
}
