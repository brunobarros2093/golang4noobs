# Genéricos 

Os tipos genéricos foram introduzidos no Go 1.18 e permitem que você escreva funções e tipos que são agnósticos em relação ao tipo de dados com os quais estão trabalhando. Antes do Go 1.18, você tinha que escrever funções separadas para diferentes tipos de dados ou usar interfaces vazias (`interface{}`), que perdem a segurança de tipo.

Os tipos genéricos são definidos usando parâmetros de tipo. Aqui está um exemplo de uma função genérica que pode trabalhar com qualquer tipo de slice:

```go
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
```

Neste exemplo, `T` é um parâmetro de tipo que pode representar qualquer tipo. `any` é uma restrição de tipo que indica que `T` pode ser qualquer tipo. A função `PrintSlice` pode ser chamada com um slice de qualquer tipo:

```go
PrintSlice([]int{1, 2, 3})       // Imprime: 1 2 3
PrintSlice([]string{"a", "b"})   // Imprime: a b
```

Os tipos genéricos também podem ser usados para definir tipos de dados genéricos. Aqui está um exemplo de uma estrutura genérica que pode armazenar qualquer tipo de dado:

```go
type Box[T any] struct {
    content T
}

func (b Box[T]) Content() T {
    return b.content
}
```

Neste exemplo, `Box` é um tipo genérico que pode armazenar qualquer tipo de dado. A função `Content` retorna o conteúdo da caixa.

Os tipos genéricos permitem que você escreva código mais reutilizável e seguro. No entanto, eles também podem tornar o código mais complexo e mais difícil de entender, então devem ser usados com cuidado.