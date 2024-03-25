# Panic e recover 

<b>panic</b> e <b>recover</b> são mecanismos usados para lidar com erros e panics em tempo de execução.

O panic é usado para indicar uma situação excepcional ou um erro grave que ocorreu durante a execução do programa. Quando um panic é chamado, a execução normal do programa é interrompida e o programa começa a desenrolar a pilha de chamadas de função (stack unwinding) até que seja encontrada uma chamada de função que tenha um bloco <b>recover</b> associado a ela.

O recover é uma função que é usada para recuperar o controle do programa após um panic. Ele só pode ser chamado dentro de uma função [deferida](https://github.com/brunobarros2093/golang4noobs/tree/main/12%20-%20Defer), que é uma função que é executada quando a função atual retorna. O recover captura o valor passado para o panic e permite que o programa continue a execução normalmente.

```
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Erro recuperado:", r)
        }
    }()

    fmt.Println("Antes do panic")
    panic("Ocorreu um erro!")
    fmt.Println("Depois do panic") // Esta linha não será executada
}
```
Neste exemplo, o programa imprime "Antes do panic" e, em seguida, chama panic com a mensagem de erro "Ocorreu um erro!". O programa entra em pânico e a função deferida é executada. O recover captura o valor passado para o panic e imprime a mensagem de erro. A execução do programa continua normalmente após o recover.

É importante notar que o recover só funciona dentro de uma função deferida. Se você chamar o recover fora de uma função deferida, ele não terá efeito e o programa continuará em pânico.

Para entender o Defer, [veja o tópico aqui](https://github.com/brunobarros2093/golang4noobs/tree/main/12%20-%20Defer).

Em Go, `panic` e `recover` são mecanismos usados para lidar com erros e panics em tempo de execução.

O `panic` é usado para indicar uma situação excepcional ou um erro grave que ocorreu durante a execução do programa. Quando um `panic` é chamado, a execução normal do programa é interrompida e o programa começa a desenrolar a pilha de chamadas de função (stack unwinding) até que seja encontrada uma chamada de função que tenha um bloco `recover` associado a ela.

O `recover` é uma função que é usada para recuperar o controle do programa após um `panic`. Ele só pode ser chamado dentro de uma função deferida, que é uma função que é executada quando a função atual retorna. O `recover` captura o valor passado para o `panic` e permite que o programa continue a execução normalmente.

Aqui está um exemplo de como usar `panic` e `recover` em Go:

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Erro recuperado:", r)
        }
    }()

    fmt.Println("Antes do panic")
    panic("Ocorreu um erro!")
    fmt.Println("Depois do panic") // Esta linha não será executada
}
```

Neste exemplo, o programa imprime "Antes do panic" e, em seguida, chama `panic` com a mensagem de erro "Ocorreu um erro!". O programa entra em pânico e a função deferida é executada. O `recover` captura o valor passado para o `panic` e imprime a mensagem de erro. A execução do programa continua normalmente após o `recover`.

É importante notar que o `recover` só funciona dentro de uma função deferida. Se você chamar o `recover` fora de uma função deferida, ele não terá efeito e o programa continuará em pânico.

O uso adequado de `panic` e `recover` é importante para lidar com erros de forma controlada e garantir que o programa não pare abruptamente. No entanto, é recomendado evitar o uso excessivo de `panic` e `recover` e, em vez disso, tratar os erros de forma apropriada sempre que possível.
