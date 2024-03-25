# Unit Tests 

Os testes unitários são uma parte fundamental do desenvolvimento de software. Eles ajudam a garantir que o seu código funcione como esperado e permitem que você faça alterações com confiança, sabendo que se algo quebrar, os testes irão falhar.

Em Go, os testes unitários são escritos em arquivos que terminam com `_test.go`. Dentro desses arquivos, as funções de teste são definidas com nomes que começam com `Test`, seguido pelo nome da função que está sendo testada.

Aqui está um exemplo de um teste unitário para uma função `Add`:

```go
// add.go
package main

func Add(x, y int) int {
    return x + y
}
```

```go
// add_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

Neste exemplo, `TestAdd` é uma função de teste para a função `Add`. Ela chama `Add` com os argumentos `2` e `3` e verifica se o resultado é `5`. Se o resultado não for `5`, ela chama `t.Errorf` para registrar um erro.

Para executar os testes, você pode usar o comando `go test` no terminal. Se todos os testes passarem, ele imprimirá `ok`. Se algum teste falhar, ele imprimirá uma descrição do teste que falhou e o que era esperado.

Além disso, o pacote `testing` fornece várias funções e tipos que ajudam a escrever testes unitários, incluindo funções para verificar se os valores são iguais (`Equal`), para verificar se um valor é verdadeiro (`True`), para verificar se um valor é falso (`False`), e assim por diante.