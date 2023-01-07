package main

import (
	"fmt"
	"net/http"
	"todo/bin/app/handler"
	c "todo/bin/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var s = &handler.Server{Chi: chi.NewRouter()}
	var r = s.Chi
	r.Use(middleware.Logger)

	s.Routes()

	var colorGreen, colorReset = "\033[32m", "\033[0m"
	var url string = c.Env.ApiUrl
	if url == "" {
		url = fmt.Sprintf("%s:%s", c.Env.Host, c.Env.Port)
	}
	if url == ":" {
		url += "3030"
	}
	fmt.Println(url)
	fmt.Printf("⇨ http server started on %s%s%s\n", colorGreen, url, colorReset)
	http.ListenAndServe(url, r)
}
