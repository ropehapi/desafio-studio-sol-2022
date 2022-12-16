package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"
)

type Request struct {
	Password string
	Rules    []Rule `json:"rules"`
}

type Rule struct {
	Rule  string
	Value int
}

type Response struct {
	Verify  bool     `json:"verify"`
	noMatch []string `json:"noMatch"`
}

func VerifyPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var request Request
	json.Unmarshal(body, &request)

	password := request.Password
	var response Response

	response.Verify = true
	for _, value := range request.Rules {
		counter := 0

		if value.Rule == "minSize" {
			if len(password) < value.Value {
				invalidatePassword(&response, "minSize")
			}
		}

		if value.Rule == "minUppercase" {
			for _, r := range password {
				if unicode.IsUpper(r) {
					counter += 1
				}
			}
			if counter < value.Value {
				invalidatePassword(&response, "minUppercase")
			}
		}

		if value.Rule == "minLowercase" {
			for _, r := range password {
				if unicode.IsLower(r) {
					counter += 1
				}
			}
			if counter < value.Value {
				invalidatePassword(&response, "minLowercase")
			}
		}

		if value.Rule == "minDigit" {
			for _, r := range password {
				if unicode.IsNumber(r) {
					counter += 1
				}
			}
			if counter < value.Value {
				invalidatePassword(&response, "minDigit")
			}
		}

		if value.Rule == "minSpecialChars" {
			for _, r := range password {
				if strings.ContainsAny(string(r), "!@#$%^&*()-+\\/{}[]") {
					counter += 1
				}
			}
			if counter < value.Value {
				invalidatePassword(&response, "minSpecialChars")
			}
		}

		if value.Rule == "noRepeted" {
			repeatCount := 1
			maximumRepeat := 2
			lastChar := ""
			for _, r := range password {
				c := string(r)
				if c == lastChar {
					repeatCount++
					if repeatCount == maximumRepeat {
						invalidatePassword(&response, "noRepeted")
					}
				} else {
					repeatCount = 1
				}
				lastChar = c
			}
		}
	}

	/**
	* Parei aqui, resolvendo a questão de como devolver um JSON
	 */
	// w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	// resp:= make(map[string]string)
	// resp["verify"] = strconv.FormatBool(response.Verify)
	// resp["noMatch"] = response.noMatch
	// jsonResp, err := json.Marshal(resp)
	// if err != nil {
	// 	log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	// }
	// w.Write(jsonResp)
	// return
}

func invalidatePassword(response *Response, rule string) {
	response.Verify = false
	response.noMatch = append(response.noMatch, rule)
}
