package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ropehapi/teste-studio-sol/controllers"
	"github.com/ropehapi/teste-studio-sol/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/verify", controllers.VerifyPassword).Methods("Post")

	log.Fatal(http.ListenAndServe(":8080", r))
}
