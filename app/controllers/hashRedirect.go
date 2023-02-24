package controllers

import (
	"net/http"
	"strings"
	"url-shortener/models"
)

func HashRedirect(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")
	namespace := params[0]
	segment := params[1]

	url, err := models.GetUrl(namespace, segment)
	if err != nil {
		http.Error(w, "Error while resolving url", http.StatusInternalServerError)
		return
	}
	if !url.Exists {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url.Value, http.StatusSeeOther)
}
