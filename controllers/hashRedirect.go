package controllers

import (
	"net/http"
	"strings"
	"url-shortener/models"
)

func HashRedirect(w http.ResponseWriter, r *http.Request) {
	hash := strings.TrimPrefix(r.URL.Path, "/")
	url, ok := models.GetUrl(hash)
	if ok {
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}
	http.NotFound(w, r)
}
