package models

type Request struct {
	Password string
	Rules    []Rule `json:"rules"`
}