package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/ropehapi/teste-studio-sol/models"
)

/*
	Essa função é o núcleo da nossa API, onde processaremos a requisição vinda do cliente e devolveremos uma resposta em JSON.
	Depois de "parsearmos" o corpo da requisição, o passamos para um switch responsável por delegar cada validação para seu respectivo método.
	Todos esses métodos de validação seguem um padrão, primeiro é feita a regra de negócio da nossa validação, e em seguida uma condicional para ver se 
	a senha é válida, para caso não seja, chamar um método responsável por dar nossa senha como inválida e armazenar o tipo de validação que falhou em um array.
*/
func VerifyPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var request models.Request
	json.Unmarshal(body, &request)

	password := request.Password
	var response models.Response

	response.Verify = true
	for _, value := range request.Rules {

		switch value.Rule {
			case "minSize":
				validateMinSize(password, &value, &response)
			case "minUppercase":
				validateMinUppercase(password, &value, &response)
			case "minLowercase":
				validateMinLowercase(password, &value, &response)
			case "minDigit":
				validateMinDigit(password, &value, &response)
			case "minSpecialChars":
				validateMinSpecialChars(password, &value, &response)
			case "noRepeted":
				validateNoRepeted(password, &value, &response)
		}
	}

	json.NewEncoder(w).Encode(response)
}

/*
	Esse é o método responsável por dar uma validação como falsa, e armazenar em um array as regras de validação que falharam.
*/
func invalidatePassword(response *models.Response, rule string) {
	response.Verify = false
	response.NoMatch = append(response.NoMatch, rule)
}

func validateMinSize(password string, value *models.Rule, response *models.Response){
	if len(password) < value.Value {
		invalidatePassword(response, "minSize")
	}
}

func validateMinUppercase(password string, value *models.Rule, response *models.Response){
	counter := 0
	for _, r := range password {
		if unicode.IsUpper(r) {
			counter += 1
		}
	}
	if counter < value.Value {
		invalidatePassword(response, "minUppercase")
	}
}

func validateMinLowercase(password string, value *models.Rule, response *models.Response){
	counter := 0
	for _, r := range password {
		if unicode.IsLower(r) {
			counter += 1
		}
	}
	if counter < value.Value {
		invalidatePassword(response, "minLowercase")
	}
}

func validateMinDigit(password string, value *models.Rule, response *models.Response){
	counter := 0
	for _, r := range password {
		if unicode.IsNumber(r) {
			counter += 1
		}
	}
	if counter < value.Value {
		invalidatePassword(response, "minDigit")
	}
}

func validateMinSpecialChars(password string, value *models.Rule, response *models.Response){
	counter := 0
	for _, r := range password {
		if strings.ContainsAny(string(r), "!@#$%^&*()-+\\/{}[]") {
			counter += 1
		}
	}
	if counter < value.Value {
		invalidatePassword(response, "minSpecialChars")
	}
}

func validateNoRepeted(password string, value *models.Rule, response *models.Response){
	repeatCount := 1
	maximumRepeat := 2
	lastChar := ""
	for _, r := range password {
		c := string(r)
		if c == lastChar {
			repeatCount++
			if repeatCount == maximumRepeat {
				invalidatePassword(response, "noRepeted")
			}
		} else {
			repeatCount = 1
		}
		lastChar = c
	}
}
