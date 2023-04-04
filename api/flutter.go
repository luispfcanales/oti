package handler

import (
	"encoding/json"
	"net/http"
)

type Course struct {
	Link     string `json:"link,omitempty"`
	Password string `json:"password,omitempty"`
}

func newCourse() Course {
	return Course{
		Password: "apps#12curs-flu2204",
		Link:     "https://dignal.com/academy-flutter/",
	}
}

// Flutter is handler to redirect
func Flutter(w http.ResponseWriter, r *http.Request) {
	c := newCourse()
	http.Redirect(w, r, c.Link, http.StatusFound)
}

// Credentials is handler that return access to course
func Credentials(w http.ResponseWriter, r *http.Request) {
	c := newCourse()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&c)
}
