package main

import (
	"fmt"
	"net/http"
	"url-shortener/config"
	"url-shortener/controllers"
	"url-shortener/utils"
)

func main() {
	config.ParseConfig()
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.HashRedirect)
	mux.HandleFunc("/url-shorten", controllers.UrlShorten)

	var handler http.Handler = mux
	handler = utils.AttachLogger(handler)

	srv := &http.Server{
		Handler: handler,
		Addr:    config.GetConfig().ADDR,
	}

	fmt.Println("Listening on port " + config.GetConfig().ADDR)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
