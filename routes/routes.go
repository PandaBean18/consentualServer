package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Router) {
	(*r).Get("/", GetHome)
}

func GetHome(resWriter http.ResponseWriter, req *http.Request) {
	print("hello\n")
	resWriter.Write([]byte("<h1>Hello World!</h1>"))
}
