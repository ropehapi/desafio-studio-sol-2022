package models

type Response struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}
