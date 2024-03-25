# Chamadas HTTP 


Primeiro, importamos os pacotes necessários. O pacote `net/http` fornece funções para construir servidores HTTP.

```go
var templates *template.Template
```

Aqui, declaramos uma variável global [`templates`](command:_github.copilot.openSymbolInFile?%5B%2224%20-%20HTTP%2Fmain.go%22%2C%22templates%22%5D "24 - HTTP/main.go") que irá armazenar os templates HTML que serão usados para renderizar as respostas.

```go
type usuario struct {
	Nome  string
	Email string
}
```

Definimos uma struct [`usuario`](command:_github.copilot.openSymbolInFile?%5B%2224%20-%20HTTP%2Fmain.go%22%2C%22usuario%22%5D "24 - HTTP/main.go") que representa um usuário com um nome e um email.

```go
func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{"Pedro", "pedrim@gmail.com"}
	err := templates.ExecuteTemplate(w, "home.html", u)
	if err != nil {
		return
	}
}
```

Aqui, definimos uma função [`home`](command:_github.copilot.openSymbolInFile?%5B%2224%20-%20HTTP%2Fmain.go%22%2C%22home%22%5D "24 - HTTP/main.go") que será chamada quando uma solicitação HTTP for feita para a rota [`/`](command:_github.copilot.openRelativePath?%5B%22%2F%22%5D "/"). Esta função recebe dois argumentos: um `http.ResponseWriter` e um `*http.Request`.

- `http.ResponseWriter` é uma interface que permite que você escreva a resposta que será enviada de volta ao cliente.
- `*http.Request` é um objeto que contém todas as informações sobre a solicitação HTTP que foi feita, incluindo os cabeçalhos, os parâmetros da query, o corpo, etc.

Dentro da função [`home`](command:_github.copilot.openSymbolInFile?%5B%2224%20-%20HTTP%2Fmain.go%22%2C%22home%22%5D "24 - HTTP/main.go"), criamos um novo [`usuario`](command:_github.copilot.openSymbolInFile?%5B%2224%20-%20HTTP%2Fmain.go%22%2C%22usuario%22%5D "24 - HTTP/main.go"), e então usamos o método `ExecuteTemplate` do pacote `html/template` para renderizar o template [`home.html`](command:_github.copilot.openSymbolInFile?%5B%2224%20-%20HTTP%2Fmain.go%22%2C%22home.html%22%5D "24 - HTTP/main.go") com os dados do usuário. O primeiro argumento para `ExecuteTemplate` é o `http.ResponseWriter` onde o HTML renderizado será escrito. O segundo argumento é o nome do template a ser renderizado. O terceiro argumento são os dados que serão passados para o template.

Se houver um erro ao renderizar o template, a função retorna imediatamente.

Note que este código não inclui a inicialização do servidor HTTP ou a definição das rotas. Normalmente, você veria algo como isto em uma função [`main`](command:_github.copilot.openSymbolInFile?%5B%2224%20-%20HTTP%2Fmain.go%22%2C%22main%22%5D "24 - HTTP/main.go"):

```go
func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Aqui, `http.HandleFunc` é usado para associar a função [`home`](command:_github.copilot.openSymbolInFile?%5B%2224%20-%20HTTP%2Fmain.go%22%2C%22home%22%5D "24 - HTTP/main.go") à rota [`/`](command:_github.copilot.openRelativePath?%5B%22%2F%22%5D "/"). `http.ListenAndServe` é então chamado para iniciar o servidor HTTP na porta 8080.