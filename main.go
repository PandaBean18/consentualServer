package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func getHome(resWriter http.ResponseWriter, req *http.Request) {
	print("hello\n")
	resWriter.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
	var r chi.Router = chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", getHome)
	http.ListenAndServe(":3000", r)
}
