package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/firebaseClient"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Router) {
	(*r).Get("/", GetHome)
	(*r).Post("/users/new", addUser)
	(*r).Post("/contracts/new/{initiatorId}/{recipientId}", addContract)
}

func GetHome(resWriter http.ResponseWriter, req *http.Request) {
	print("hello\n")
	resWriter.Write([]byte("<h1>Hello World!</h1>"))
}

func addUser(resWriter http.ResponseWriter, req *http.Request) {
	user := &firebaseClient.User{}
	err := json.NewDecoder(req.Body).Decode(user)

	if err != nil {
		print(err)
	}

	userId := ""

	var result *string = firebaseClient.AddUser(user, &userId)
	if *result == "" {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(500)
		resWriter.Write([]byte(`{"error": "Something went wrong while trying to add user."}`))
	} else {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(200)
		resWriter.Write([]byte(fmt.Sprintf(`{"userId": "%s"}`, *result)))
	}
}

func addContract(resWriter http.ResponseWriter, req *http.Request) {
	contract := &firebaseClient.Contract{}
	initiatorId := chi.URLParam(req, "initiatorId")
	recipientId := chi.URLParam(req, "recipientId")

	contract.InitiatorId = initiatorId
	contract.RecipientId = recipientId

	contractId := ""

	result := firebaseClient.AddContract(contract, &contractId)

	if *result == "" {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(500)
		resWriter.Write([]byte(`{"error": "Something went wrong while trying to add contract."}`))
	} else {
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(200)
		resWriter.Write([]byte(fmt.Sprintf(`{"contractId": "%s"}`, *result)))
	}

}
