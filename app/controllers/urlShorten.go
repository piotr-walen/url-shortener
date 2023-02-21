package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"url-shortener/config"
	"url-shortener/models"
	"url-shortener/utils"
)

type UrlShortenBody struct {
	Url string `json:"url"`
}

type UrlShortenResponse struct {
	Hash string `json:"hash"`
}

func UrlShorten(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read request body", http.StatusBadRequest)
		return
	}

	payload := UrlShortenBody{}
	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		http.Error(w, "Malformed request body", http.StatusBadRequest)
		return
	}

	v, err := url.ParseRequestURI(payload.Url)
	if err != nil {
		http.Error(w, "Malformed url in request body", http.StatusBadRequest)
		return
	}

	hash := utils.GetShortHash(v.RawPath, config.GetConfig().MAX_URL_LENGTH)
	models.AddUrl(hash, payload.Url)

	response := UrlShortenResponse{Hash: hash}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Cannot send request response", http.StatusInternalServerError)
		return
	}

}
