# **Teste Studio Sol**
*Antes de mais nada, gostaria de agradecer pela oportunidade de participar desse teste.*

*Por mais que eu não passe, esse desafio foi muito bom para o meu aprendizado, pois era o que eu estava precisando para cair de cabeça nos estudos da tecnologia. Eu vinha estudando Golang a pouco mais de duas semanas, e estava precisando de motivação, um norte, e um caso de uso real para aplicar meus conhecimentos. Muito obrigado.*

## **Sobre o código**
- ### **Explicação do script**
    Todo o código dessa API foi explicado através de blocos de comentários em cima dos métodos. Tenho em mente que comentários não são elegantes para o código, mas por se tratar de um teste, que deve ter sua escrita mais didática, optei por fazer dessa forma.

- ### **Testes**
    Também é importante notar que o projeto conta com um teste unitário testando seu único endpoint. Esse teste, que pode ser rodado através do comando `go test` valida a integridade da resposta dado uma requisição mockada.

- ### **Conteinerização**
    O projeto foi dispõe de um Dockerfile, afim de que possamos subir a aplicação em diferentes ambientes sem mais problemas.
- ### **Versionamento**
    Esse projeto foi versionado do início ao fim. Portanto, basta rodar seus comandos `git` para conferir qualquer coisa que desejar sobre o versionamento do projeto, ou até mesmo acessar o seu [repositório no GitHub](https://github.com/ropehapi/teste-studio-sol).

---

## **Rodando o projeto**
Afim de facilitar a execução do programa e evitar problemas de configuração de ambiente, eu disponibilizei um Dockerfile na raíz do projeto, que ao ter sua imagem buildada e container de pé, expõe a aplicação através de `localhost:8080`. Basta rodar os seguintes comandos:

> docker build . -t teste-studio-sol

> docker run -p 8080:8080 -t teste-studio-sol

## **Consumindo o endpoint /verify**
Nossa API dispõe de apenas uma rota, que é a responsável por processar toda a regra de negócio do teste. No caso, validar se uma senha atende aos requisitos.

Para consumir a rota, devemos fazer uma requisição de método `POST`, para o endereço `localhost:8080/verify`, informando o corpo da requisição em `Json`, que deve conter um campo `password`, e um um campo `rules`, contendo um array de regras que devem possuir o nome da regra (`rule`), e seu respectivo valor (`value`).

    {
        "password": "ABCdef123!@#",
        "rules": [
            {
                "rule": "minSize",
                "value": 12
            },
            {
                "rule": "minUppercase",
                "value": 3
            },
            {
                "rule": "minLowercase",
                "value": 3
            },
            {
                "rule": "minDigit",
                "value": 3
            },
            {
                "rule": "minSpecialChars",
                "value": 3
            },
            {
                "rule": "noRepeted",
                "value": 0
            }
            
        ]
    }

Essa requisição, de acordo com as regras informadas, nos devolverá um Json contendo os campos `verify` e `noMatch`, tendo o verify um valor booleano, responsável por dizer se a senha é válida ou não, e o noMatch sendo um array de strings contendo os nomes das regras que falharam, como por exemplo:

    {
    	"verify": true,
	    "noMatch": null
    }

    ou

    {
        "verify": false,
        "noMatch": [
            "minSize",
            "minUppercase",
            "minLowercase",
            "minSpecialChars"
        ]
    }

## **Considerações finais**
Acredito que seja isso, gostaria de ter feito também uma alternativa em GraphQL, mas eu já tive que correr contra o tempo para aprender quase que tudo isso durante a semana, então não consegui, mas ficou o incentivo para eu aprender assim que possível.

Mais uma vez, muito obrigado pela oportunidade.


