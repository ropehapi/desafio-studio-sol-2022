package models

/*
	Struct utilizada como modelo para "encodarmos" nossas responses, de modo a atender o modelo de response pedida.
*/
type Response struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}
