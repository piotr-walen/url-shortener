package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"url-shortener/models"
	"url-shortener/utils"
)

type UrlShortenBody struct {
	Url string `json:"url"`
}

type UrlShortenResponse struct {
	Hash string `json:"hash"`
}

const MAX_URL_LENGTH = 6

func UrlShorten(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	payload := UrlShortenBody{}
	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		log.Fatal(err)
	}

	hash := utils.GetShortHash(payload.Url, MAX_URL_LENGTH)
	models.AddUrl(hash, payload.Url)

	response := UrlShortenResponse{Hash: hash}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}

}
