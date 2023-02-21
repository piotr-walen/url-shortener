package controllers

import (
	"net/http"
	"strings"
	"url-shortener/models"
)

func HashRedirect(w http.ResponseWriter, r *http.Request) {
	hash := strings.TrimPrefix(r.URL.Path, "/")
	url, err := models.GetUrl(hash)
	if err != nil {
		http.Error(w, "Error while resolving url", http.StatusInternalServerError)
		return
	}
	if url.Exists {
		http.Redirect(w, r, url.Value, http.StatusSeeOther)
		return
	}
	http.NotFound(w, r)
}
