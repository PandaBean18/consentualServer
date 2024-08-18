package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"server/routes"
)

func main() {
	var r chi.Router = chi.NewRouter()

	r.Use(middleware.Logger)
	routes.RegisterRoutes(&r)
	http.ListenAndServe(":3000", r)
}
