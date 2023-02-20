package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/controllers"
)

const MAX_URL_LENGTH = 6

func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		uri := r.URL.String()
		method := r.Method
		log.Println(uri, method)
	}

	return http.HandlerFunc(fn)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.HashRedirect)
	mux.HandleFunc("/url-shorten", controllers.UrlShorten)

	var handler http.Handler = mux
	handler = logRequestHandler(handler)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":8000",
	}

	fmt.Println("Listening on port :8000")
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
