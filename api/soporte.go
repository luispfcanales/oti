package handler

import "net/http"

const url = "https://sites.google.com/unamad.edu.pe/hoja-de-servicio-oti/pagina-principal"

// Soporte is handler to redirect
func Soporte(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, url, http.StatusFound)
}
