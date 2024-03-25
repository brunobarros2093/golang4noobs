# Padrões de Concorrencia 

**1. Worker Pool**

O padrão Worker Pool, também conhecido como Thread Pool, é usado para restringir o número de goroutines que são executadas simultaneamente. Isso é útil para controlar o uso de recursos e evitar a sobrecarga do sistema. 

Aqui está um exemplo básico de um Worker Pool:

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= 9; a++ {
        <-results
    }
}
```

**2. Generators**

O padrão Generator é um tipo de função que retorna um canal. Os geradores são usados para encapsular goroutines que produzem dados, permitindo que os consumidores leiam os dados como se fossem uma sequência.

Aqui está um exemplo de um Generator:

```go
func count(start, end int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := start; i <= end; i++ {
            ch <- i
        }
        close(ch)
    }()
    return ch
}

func main() {
    for num := range count(1, 5) {
        fmt.Println(num)
    }
}
```

**3. Multiplexador**

O padrão Multiplexador, também conhecido como Fan-in, é usado para combinar múltiplos canais em um. Isso é útil quando você tem várias goroutines produzindo dados, mas você quer consumir os dados como se viessem de uma única fonte.

Aqui está um exemplo de um Multiplexador:

```go
func merge(chans ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }

    wg.Add(len(chans))
    for _, c := range chans {
        go output(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

func main() {
    in1 := count(1, 5)
    in2 := count(6, 10)

    for num := range merge(in1, in2) {
        fmt.Println(num)
    }
}
```

Neste exemplo, `merge` é uma função que combina múltiplos canais em um. Ela inicia uma nova goroutine para cada canal de entrada que lê os dados do canal e os envia para o canal de saída. Quando todos os canais de entrada foram fechados, ela fecha o canal de saída.