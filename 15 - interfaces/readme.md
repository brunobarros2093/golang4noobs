# Interfaces 

Em Go, uma interface é um tipo que define um conjunto de métodos. Um tipo implementa uma interface implementando todos os métodos especificados na interface. Diferentemente de muitas outras linguagens, em Go não é necessário declarar explicitamente que um tipo implementa uma interface. Se um tipo possui todos os métodos que uma interface requer, ele implementa essa interface.

Aqui está um exemplo de uma interface e um tipo que a implementa:

```go
type Vehiculo interface {
    Mover() string
}

type Carro struct {
    Marca string
}

// Carro implementa a interface Vehiculo
func (c Carro) Mover() string {
    return "O carro está se movendo"
}

func main() {
    var v Vehiculo = Carro{"Toyota"}
    fmt.Println(v.Mover()) // Imprime: O carro está se movendo
}
```

Neste exemplo, `Vehiculo` é uma interface que requer um método `Mover` que retorna uma string. `Carro` é um tipo que tem um método `Mover`, então ele implementa a interface `Vehiculo`.

