# Aplicação CLI 

No arquivo [`app.go`](command:_github.copilot.openRelativePath?%5B%2217%20-%20aplicac%C3%A3o%20cli%2Fapp%2Fapp.go%22%5D "17 - aplicacão cli\app\app.go"), definimos uma função [`Gerar`](command:_github.copilot.openSymbolInFile?%5B%2217%20-%20aplicac%C3%A3o%20cli%2Fapp%2Fapp.go%22%2C%22Gerar%22%5D "17 - aplicacão cli/app/app.go") que cria e configura uma nova aplicação CLI.

```go
func Gerar() *cli.App {
    app := cli.NewApp()
    app.Name = "Aplicação linha de comando estudos"
    app.Usage = "Busca IPs e nomes de servidores na internet"
    // ...
}
```

A função `cli.NewApp` cria uma nova aplicação CLI passando nome e descrição do app. Em seguida, você definirá uma flag `host` que pode ser usada para especificar o host para o qual a aplicação deve buscar informações.

```go
flags := []cli.Flag{
    cli.StringFlag{
        Name:  "host",
        Value: "devbook.com.br",
    },
}
```

A flag `host` é uma flag de string, o que significa que ela espera um valor de string. O valor padrão é "devbook.com.br".

Finalmente, você começa a definir os comandos que a aplicação deve suportar. Cada comando é definido como um `cli.Command` com um nome e outros campos que definem como o comando deve funcionar.

```go
app.Commands = []cli.Command{
    {
        Name:   "ip",
        // ...
    },
}
```

O arquivo `main.go` seria responsável por chamar a função [`Gerar`](command:_github.copilot.openSymbolInFile?%5B%2217%20-%20aplicac%C3%A3o%20cli%2Fapp%2Fapp.go%22%2C%22Gerar%22%5D "17 - aplicacão cli/app/app.go") para criar a aplicação e então chamar `app.Run` para iniciar a aplicação e processar os argumentos de linha de comando.

```go
func main() {
    app := app.Gerar()
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
```

Neste exemplo, `os.Args` é passado para `app.Run`. `os.Args` é um slice que contém os argumentos de linha de comando que foram passados para o programa.