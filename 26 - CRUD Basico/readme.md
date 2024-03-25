# CRUD em Golang 

Como implementar operações CRUD (Create, Read, Update, Delete) em um servidor HTTP em Go. Ele usa o pacote `github.com/gorilla/mux` para roteamento e manipulação de solicitações HTTP. Aqui está uma explicação detalhada do código:

```go
router := mux.NewRouter()
```

Aqui, criamos um novo roteador usando a função `mux.NewRouter`. Este roteador será usado para registrar as rotas e as funções de manipulação correspondentes.

```go
router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods("POST")
```

Aqui, registramos a rota `/usuarios` com o método HTTP `POST`. Quando uma solicitação POST é feita para `/usuarios`, a função `servidor.CriarUsuario` será chamada. Esta função é responsável por criar (Create) um novo usuário.

```go
router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods("GET")
```

Aqui, registramos a rota `/usuarios` com o método HTTP `GET`. Quando uma solicitação GET é feita para `/usuarios`, a função `servidor.BuscarUsuarios` será chamada. Esta função é responsável por buscar (Read) todos os usuários.

```go
router.HandleFunc("/usuario/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
```

Aqui, registramos a rota `/usuario/{id}` com o método HTTP `GET`. Quando uma solicitação GET é feita para `/usuario/{id}`, a função `servidor.BuscarUsuario` será chamada. Esta função é responsável por buscar (Read) um usuário específico com base no `id` fornecido na URL.

```go
router.HandleFunc("/usuario/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
```

Aqui, registramos a rota `/usuario/{id}` com o método HTTP `PUT`. Quando uma solicitação PUT é feita para `/usuario/{id}`, a função `servidor.AtualizarUsuario` será chamada. Esta função é responsável por atualizar (Update) um usuário específico com base no `id` fornecido na URL.

```go
router.HandleFunc("/usuario/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)
```

Aqui, registramos a rota `/usuario/{id}` com o método HTTP `DELETE`. Quando uma solicitação DELETE é feita para `/usuario/{id}`, a função `servidor.DeletarUsuario` será chamada. Esta função é responsável por deletar (Delete) um usuário específico com base no `id` fornecido na URL.

```go
log.Fatal(http.ListenAndServe(":5000", router))
```

Finalmente, iniciamos o servidor HTTP na porta 5000 usando a função `http.ListenAndServe`. O segundo argumento para esta função é o roteador que criamos, que contém todas as nossas rotas registradas.

Note que as funções `servidor.CriarUsuario`, `servidor.BuscarUsuarios`, `servidor.BuscarUsuario`, `servidor.AtualizarUsuario` e `servidor.DeletarUsuario` não são definidas neste arquivo main, mas você encontrará os arquivos na pasta 26. Acredito que se você chegou até aqui deverá entender perfeitamente os métodos listados acima, mas sinta-se a vontade para abrir um Issue ou PR para tirar dúvidas! 

<hr>

# servidor.go

O arquivo [`servidor.go`](command:_github.copilot.openRelativePath?%5B%2226%20-%20CRUD%20Basico%2Fservidor%2Fservidor.go%22%5D "26 - CRUD Basico\servidor\servidor.go") é nosso servidor web em Go que interage com um banco de dados para realizar operações CRUD (Create, Read, Update, Delete) em usuários. A função [`BuscarUsuarios`](command:_github.copilot.openSymbolInFile?%5B%2226%20-%20CRUD%20Basico%2Fservidor%2Fservidor.go%22%2C%22BuscarUsuarios%22%5D "26 - CRUD Basico/servidor/servidor.go") é uma dessas operações, especificamente a operação "Read".

Aqui está uma explicação detalhada da função [`BuscarUsuarios`](command:_github.copilot.openSymbolInFile?%5B%2226%20-%20CRUD%20Basico%2Fservidor%2Fservidor.go%22%2C%22BuscarUsuarios%22%5D "26 - CRUD Basico/servidor/servidor.go"):

```go
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
```

A função [`BuscarUsuarios`](command:_github.copilot.openSymbolInFile?%5B%2226%20-%20CRUD%20Basico%2Fservidor%2Fservidor.go%22%2C%22BuscarUsuarios%22%5D "26 - CRUD Basico/servidor/servidor.go") é um manipulador de solicitações HTTP. Ela recebe dois argumentos: um `http.ResponseWriter` para enviar respostas HTTP, e um `*http.Request`, que representa a solicitação HTTP recebida.

```go
db, erro := banco.Conectar()
if erro != nil {
	w.Write([]byte("Erro ao conectar no banco!"))
	return
}
```

Aqui, a função tenta conectar-se ao banco de dados usando a função `banco.Conectar()`. Se houver um erro ao conectar-se ao banco de dados, a função envia uma resposta HTTP com a mensagem "Erro ao conectar no banco!" e retorna imediatamente.

```go
defer db.Close()
```

Aqui, a função garante que a conexão com o banco de dados será fechada quando a função [`BuscarUsuarios`](command:_github.copilot.openSymbolInFile?%5B%2226%20-%20CRUD%20Basico%2Fservidor%2Fservidor.go%22%2C%22BuscarUsuarios%22%5D "26 - CRUD Basico/servidor/servidor.go") terminar, usando a palavra-chave `defer`.

```go
linhasDoDb, erro := db.Query("select * from usuarios")
if erro != nil {
	w.Write([]byte("Erro buscar usuários"))
	return
}
```

Aqui, a função executa uma consulta SQL para buscar todos os usuários do banco de dados. Se houver um erro ao executar a consulta, a função envia uma resposta HTTP com a mensagem "Erro buscar usuários" e retorna imediatamente.

```go
defer linhasDoDb.Close()
```

Aqui, a função garante que o conjunto de resultados da consulta SQL será fechado quando a função [`BuscarUsuarios`](command:_github.copilot.openSymbolInFile?%5B%2226%20-%20CRUD%20Basico%2Fservidor%2Fservidor.go%22%2C%22BuscarUsuarios%22%5D "26 - CRUD Basico/servidor/servidor.go") terminar, usando a palavra-chave `defer`.

```go
var usuarios []usuario
for linhasDoDb.Next() {
	var usuario usuario
	if erro := linhasDoDb.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
		w.Write([]byte("Erro ao scannear usuário"))
		return
	}
```

Aqui, a função itera sobre o conjunto de resultados da consulta SQL. Para cada linha, ela cria um novo [`usuario`](command:_github.copilot.openSymbolInFile?%5B%2226%20-%20CRUD%20Basico%2Fservidor%2Fservidor.go%22%2C%22usuario%22%5D "26 - CRUD Basico/servidor/servidor.go") e usa o método `Scan` para preencher os campos do [`usuario`](command:_github.copilot.openSymbolInFile?%5B%2226%20-%20CRUD%20Basico%2Fservidor%2Fservidor.go%22%2C%22usuario%22%5D "26 - CRUD Basico/servidor/servidor.go") com os valores da linha. Se houver um erro ao fazer isso, a função envia uma resposta HTTP com a mensagem "Erro ao scannear usuário" e retorna imediatamente.

<hr>

# banco.go 

O arquivo [`banco.go`](command:_github.copilot.openRelativePath?%5B%2226%20-%20CRUD%20Basico%2Fbanco%2Fbanco.go%22%5D "26 - CRUD Basico\banco\banco.go") é responsável por estabelecer uma conexão com um banco de dados MySQL. Ele usa o pacote `database/sql` para interagir com o banco de dados e o driver `github.com/go-sql-driver/mysql` para conectar-se ao MySQL.

Aqui está uma explicação detalhada da função [`Conectar`](command:_github.copilot.openSymbolInFile?%5B%2226%20-%20CRUD%20Basico%2Fbanco%2Fbanco.go%22%2C%22Conectar%22%5D "26 - CRUD Basico/banco/banco.go"):

```go
func Conectar() (*sql.DB, error) {
```

A função [`Conectar`](command:_github.copilot.openSymbolInFile?%5B%2226%20-%20CRUD%20Basico%2Fbanco%2Fbanco.go%22%2C%22Conectar%22%5D "26 - CRUD Basico/banco/banco.go") retorna um ponteiro para uma conexão de banco de dados `*sql.DB` e um `error`. Se a conexão for bem-sucedida, o `error` será `nil`.

```go
stringConn := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
```

Aqui, definimos a string de conexão para o banco de dados MySQL. A string de conexão contém as informações necessárias para conectar-se ao banco de dados, incluindo o nome de usuário (`golang`), a senha (`golang`), o nome do banco de dados (`devbook`), e outros parâmetros de conexão.

```go
db, err := sql.Open("mysql", stringConn)
```

Aqui, tentamos abrir uma conexão com o banco de dados usando a função `sql.Open`. O primeiro argumento é o nome do driver do banco de dados (`mysql`), e o segundo argumento é a string de conexão que definimos anteriormente.

```go
if err != nil {
	return nil, err
}
```

Se houver um erro ao abrir a conexão, a função retorna `nil` e o erro.

```go
if err = db.Ping(); err != nil {
	return nil, err
}
```

Aqui, tentamos fazer um "ping" no banco de dados para garantir que a conexão está funcionando corretamente. Se houver um erro ao fazer o ping, a função retorna `nil` e o erro.

```go
return db, nil
```

Se a conexão for bem-sucedida e o ping não retornar nenhum erro, a função retorna a conexão de banco de dados e `nil` para o erro.

Note que o driver MySQL (`github.com/go-sql-driver/mysql`) é importado com um underscore (`_`) antes do nome do pacote. Isso significa que o pacote é importado apenas por seus efeitos colaterais (neste caso, registrar o driver MySQL para que possa ser usado com `database/sql`), e não é usado diretamente no código.


