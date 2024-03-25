# Banco de Dados 

Como estabelecer uma conexão com um banco de dados MySQL em Go usando o pacote `database/sql` e o driver `go-sql-driver/mysql`. Aqui está uma explicação detalhada do código na main.go:

```go
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
```

Primeiro, importamos os pacotes necessários. O pacote `database/sql` fornece uma interface genérica para interagir com qualquer tipo de banco de dados SQL. O driver `go-sql-driver/mysql` é específico para MySQL e implementa a interface definida por `database/sql`.

```go
stringConn := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
db, err := sql.Open("mysql", stringConn)
```

Aqui, usamos a função `sql.Open` para abrir uma conexão com o banco de dados. O primeiro argumento é o nome do driver do banco de dados (neste caso, "mysql"). O segundo argumento é a string de conexão do banco de dados, que contém informações como o nome de usuário, a senha, o nome do banco de dados e outros parâmetros de conexão.

```go
if err != nil {
	log.Fatal(err)
}
```

Se houver um erro ao abrir a conexão, o programa irá registrar o erro e parar a execução.

```go
defer func(db *sql.DB) {
	_ = db.Close()
}(db)
```

Aqui, usamos a palavra-chave `defer` para garantir que a conexão com o banco de dados será fechada quando a função [`main`](command:_github.copilot.openSymbolInFile?%5B%2225%20-%20Banco%20de%20Dados%2Fmain.go%22%2C%22main%22%5D "25 - Banco de Dados/main.go") terminar. Isso é importante para liberar os recursos que a conexão está usando.

```go
if err = db.Ping(); err != nil {
	log.Fatal(err)
}
```

Aqui, usamos o método `Ping` para verificar se a conexão com o banco de dados está funcionando corretamente. Se houver um erro, o programa irá registrar o erro e parar a execução.

```go
fmt.Println("Conexão aberta com sucesso!")
```

Se não houver erros, o programa irá imprimir uma mensagem dizendo que a conexão foi aberta com sucesso.

```go
linhas, err := db.Query("select * from usuarios")
if err != nil {
```

Finalmente, usamos o método `Query` para executar uma consulta SQL no banco de dados. O resultado da consulta é retornado como um `*sql.Rows`, que pode ser usado para iterar sobre as linhas retornadas pela consulta. Se houver um erro ao executar a consulta, o programa irá registrar o erro e parar a execução.