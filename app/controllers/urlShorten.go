package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"url-shortener/models"
)

type UrlShortenBody struct {
	TargetUrl string `json:"targetUrl"`
	Namespace string `json:"namespace"`
	Segment   string `json:"segment"`
}

var alfanumericRegex = regexp.MustCompile(`^\w+$`)

func (p *UrlShortenBody) validate() error {
	errs := []error{}

	_, err := url.ParseRequestURI(p.TargetUrl)
	if err != nil {
		errs = append(errs, errors.New("invalid targetUrl field"))
	}
	if ok := alfanumericRegex.MatchString(p.Namespace); !ok {
		errs = append(errs, errors.New("invalid namespace field"))
	}
	if ok := alfanumericRegex.MatchString(p.Segment); !ok {
		errs = append(errs, errors.New("invalid segment field"))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
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

	err = payload.validate()
	if err != nil {
		http.Error(w, "Malformed request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	ok, err := models.AddUrl(payload.Namespace, payload.Segment, payload.TargetUrl)
	if err != nil {
		http.Error(w, "Error when adding url", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Url already exists", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
