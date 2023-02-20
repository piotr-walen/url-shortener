package main

import (
	"fmt"
	"net/http"
	"url-shortener/controllers"
	"url-shortener/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.HashRedirect)
	mux.HandleFunc("/url-shorten", controllers.UrlShorten)

	var handler http.Handler = mux
	handler = utils.AttachLogger(handler)

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
