package models

/*
	Struct utilizada para "parsear" nossa requisição, de modo que tenhamos um objeto com a nossa senha, e um array de 
	regras, que é outra struct que também está detalhada.
*/
type Request struct {
	Password string
	Rules    []Rule `json:"rules"`
}