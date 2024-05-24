package main

import (
	"golang_pr/config"
	"golang_pr/sql"
	"log"
	"net/http"
)

var cfg *config.Config

func main() {
	sql.Migrate()
	cfg = config.Get()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlePage)
	mux.HandleFunc("/api/", HandleApi)
	mux.HandleFunc("/resources/", HandleResource)
	server := http.Server{
		Addr:    config.Location(),
		Handler: mux,
	}
	log.Println("MAIN", "Server opened on address", cfg.Server.Host, "with port", cfg.Server.Port)
	log.Println("MAIN", "Server work in dir:", cfg.Server.Workdir)
	server.ListenAndServe()
}
