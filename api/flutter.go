package handler

import (
	"net/http"
)

// Flutter is handler to redirect
func Flutter(w http.ResponseWriter, r *http.Request) {
	c := newCourse()
	http.Redirect(w, r, c.Link, http.StatusFound)
}
