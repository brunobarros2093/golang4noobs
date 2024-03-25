# Go routines 

As goroutines são uma das características mais poderosas do Go. Elas são funções que podem ser executadas concorrentemente com outras funções. Você pode iniciar uma goroutine simplesmente usando a palavra-chave `go` antes de uma chamada de função. Aqui está um exemplo:

```go
func hello() {
    fmt.Println("Hello from goroutine")
}

func main() {
    go hello()
    fmt.Println("Hello from main function")

    // Aguarda para que a goroutine termine
    time.Sleep(time.Second)
}
```

Neste exemplo, `hello` é uma função que imprime uma mensagem. Na função `main`, `go hello()` inicia uma nova goroutine que executa a função `hello`. A função `main` continua a execução imediatamente, sem esperar que a função `hello` termine.

As goroutines são mais leves do que threads e são gerenciadas pelo runtime do Go, não pelo sistema operacional. Isso significa que você pode iniciar milhares ou mesmo dezenas de milhares de goroutines em um único programa.

No entanto, as goroutines são independentes umas das outras, o que significa que quando a função `main` termina, todas as goroutines são abruptamente paradas, mesmo que não tenham terminado a execução. É por isso que o exemplo acima usa `time.Sleep` para dar à goroutine a chance de terminar antes que a função `main` termine.

Para um controle mais preciso sobre as goroutines, você pode usar canais e o pacote `sync`, que fornecem a sincronização e comunicação entre goroutines.