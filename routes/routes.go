package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ropehapi/teste-studio-sol/controllers"
	"github.com/ropehapi/teste-studio-sol/middleware"
)

/*
	Aqui, criamos o nosso router baseado no GorillaMux, um terminal de requisições mais poderoso.
	Definimos nossa rota /verify, que dispara o método VerifyPassword, que é a função responsável por validar nossa senha.
	Logo após, subimos uma instância do nosso servidor na porta 8080.
*/
func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/verify", controllers.VerifyPassword).Methods("Post")

	log.Fatal(http.ListenAndServe(":8080", r))
}
