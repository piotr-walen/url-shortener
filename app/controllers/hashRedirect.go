package controllers

import (
	"log"
	"net/http"
	"strings"
	"url-shortener/models"
)

func HashRedirect(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")
	namespace := params[0]
	segment := params[1]

	log.Println(namespace, segment)

	url, err := models.GetUrl(namespace, segment)
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
