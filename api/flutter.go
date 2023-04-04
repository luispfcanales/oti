package handler

import (
	"net/http"

	"github.com/luispfcanales/oti/entity"
)

// Flutter is handler to redirect
func Flutter(w http.ResponseWriter, r *http.Request) {
	c := entity.NewCourse()
	http.Redirect(w, r, c.Link, http.StatusFound)
}
