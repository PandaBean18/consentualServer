package main

import (
	"net/http"

	"server/firebaseClient"
	"server/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var r chi.Router = chi.NewRouter()

	r.Use(middleware.Logger)
	//r.Get("/", routes.GetHome)
	routes.RegisterRoutes(&r)
	client := firebaseClient.CreateClient()
	client.Close()

	http.ListenAndServe(":3000", r)

}
