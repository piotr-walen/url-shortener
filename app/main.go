package main

import (
	"log"
	"net/http"
	"url-shortener/config"
	"url-shortener/controllers"
	"url-shortener/storage"
	"url-shortener/utils"
)

func main() {
	err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = storage.Connect()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.HashRedirect)
	mux.HandleFunc("/url-shorten", controllers.UrlShorten)

	var handler http.Handler = mux
	handler = utils.NewLogger(utils.NewCorsHeaders(handler))

	srv := &http.Server{
		Handler: handler,
		Addr:    config.GetConfig().Addr,
	}

	log.Println("Listening on port " + config.GetConfig().Addr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
