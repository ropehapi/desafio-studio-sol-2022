/*
	Antes de mais nada, que fique claro que estou ciente de que o código deve ser limpo a ponto de dispensar comentários.
	Porém, por se tratar de uma avaliação, eu os usarei para explicar o fluxo do script.
*/
package main

import "github.com/ropehapi/teste-studio-sol/routes"

/*
	A função main tem uma única funcionalidade, chamar o método HandleRequest,
	que contém uma instância do nosso router e inicializa o nosso servidor.
*/
func main() {
	routes.HandleRequest()
}