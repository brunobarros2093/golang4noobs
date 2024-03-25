# Select 

O `select` em Go é uma instrução que permite que você 'espere' durante a comunicação entre canais. Ele escolhe o 'caso' pronto para a comunicação e executa o respectivo bloco de código. Se mais de um caso estiver pronto, ele escolhe um aleatoriamente.

Aqui está um exemplo básico de como usar `select` com canais:

```go
func main() {
    chan1 := make(chan string)
    chan2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        chan1 <- "from chan1"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        chan2 <- "from chan2"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-chan1:
            fmt.Println(msg1)
        case msg2 := <-chan2:
            fmt.Println(msg2)
        }
    }
}
```

Neste exemplo, temos duas goroutines enviando dados para dois canais diferentes. O `select` está esperando dados de ambos os canais. Quando um canal está pronto para receber dados, o `select` executa o caso correspondente.

Note que se nenhum dos casos estiver pronto, o `select` irá bloquear até que um caso esteja pronto. Se você quiser que o `select` não bloqueie, você pode adicionar um caso `default`:

```go
select {
case msg1 := <-chan1:
    fmt.Println(msg1)
case msg2 := <-chan2:
    fmt.Println(msg2)
default:
    fmt.Println("no message received")
}
```

Neste exemplo, se nenhum dos canais estiver pronto para receber dados, o `select` irá executar o caso `default` e continuar.