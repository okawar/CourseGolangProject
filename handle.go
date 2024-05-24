package main

import (
	"golang_pr/service"
	"log"
	"net/http"
	"strings"
)

func HandleApi(w http.ResponseWriter, r *http.Request) {
	var ctx service.Context
	ctx.Request = r
	ctx.Response = w
	for action, srvc := range cfg.Api {
		if r.Method == srvc.Method && r.URL.Path == srvc.Url {
			log.Println("MAIN", "API REQ:", r.Method, r.URL.String())
			if srvc.NeedAuth {
				auth, _ := service.Authentificate(&ctx)
				if auth {
					service.CRUDS[action](&ctx)
				} else {
					w.WriteHeader(401)
				}
			} else {
				service.CRUDS[action](&ctx)
			}
		}
	}
}

func HandleResource(w http.ResponseWriter, r *http.Request) {
	log.Println("MAIN", "RES REQ", r.URL.Path)
	http.ServeFile(w, r, cfg.Server.Workdir+r.URL.Path)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.URL.Path, ".") || !strings.HasSuffix(r.URL.Path, ".html") {
		log.Println("MAIN", "PAGE REQ", r.Method, r.URL.Path)
		http.ServeFile(w, r, cfg.Server.Workdir+"/resources/pages"+r.URL.Path+".html")
	} else {
		log.Println("MAIN", "Incorrect url for page '", r.URL.Path, "' URl mustn't have '.'")
	}
}
