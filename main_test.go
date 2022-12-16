package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ropehapi/teste-studio-sol/controllers"
	"github.com/ropehapi/teste-studio-sol/middleware"
	"github.com/stretchr/testify/assert"
)

/*
	Aqui eu faço um método que devolve uma instância do router para que eu possa usar nos testes.
*/
func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	return r
}

/*
	Esse teste unitário joga uma senha com todas as suas regras exigidas, e deve retornar verdadeiro.
	Caso contrário, alguma coisa quebrou o teste lá na controller.
*/
func TestVerifyPasswordEndpoint(t *testing.T) {
	r := setupRouter()
	r.HandleFunc("/verify", controllers.VerifyPassword).Methods("Post")

	var buf []byte
	jsonBody := "{\"password\": \"ABCdef123!@#\",\"rules\": [{\"rule\": \"minSize\",\"value\": 12},{\"rule\": \"minUppercase\",\"value\": 3},{\"rule\": \"minLowercase\",\"value\": 3},{\"rule\": \"minDigit\",\"value\": 3},{\"rule\": \"minSpecialChars\",\"value\": 3},{\"rule\": \"noRepeted\",\"value\": 0}]}"
	buf, _ = json.Marshal(jsonBody)

	req, _ := http.NewRequest("POST", "/verify", bytes.NewBuffer(buf))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "{\"verify\":true,\"noMatch\":null}\n", response.Body.String())
}
