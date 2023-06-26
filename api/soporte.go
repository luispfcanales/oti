package handler

import "net/http"

const url = "https://docs.google.com/forms/d/e/1FAIpQLSer6WCD7JdmeMTP-C9iuS5tsL2BcMTcq4OIz_ReN0YY0KTqBg/viewform?usp=sharing"

// Soporte is handler to redirect
func Soporte(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, url, http.StatusFound)
}
