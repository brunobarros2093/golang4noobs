# Testes parte 2 


```go
package formas_test

import (
	"math"
	"testes-avancado/formas"
	"testing"
)
```

Este código define o pacote de teste e importa os pacotes necessários. O pacote `testing` é o pacote padrão de Go para escrever testes unitários, como já vimos na [parte 1](https://github.com/brunobarros2093/golang4noobs/tree/main/21%20-%20Testes). O pacote `math` é usado para cálculos matemáticos. O pacote `testes-avancado/formas` é o pacote que contém o código que está sendo testado.

```go
func TestArea(t *testing.T) {
```

Esta é a definição da função de teste. Todas as funções de teste em Go devem ter um único argumento do tipo `*testing.T`.

```go
t.Run("Retangulo", func(t *testing.T) {
		ret := formas.Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()
		if areaEsperada != areaRecebida {
			t.Errorf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
```

Aqui, `t.Run` é usado para definir um subteste. Este subteste está testando a função `Area` para um `Retangulo`. Ele cria um `Retangulo`, calcula a área, e verifica se a área é o que se esperava. Se a área não for o que se esperava, ele chama `t.Errorf` para registrar um erro.

```go
t.Run("Circulo", func(t *testing.T) {
		circ := formas.Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()
```

Este é outro subteste, similar ao primeiro, mas testando a função `Area` para um `Circulo`.

Para executar os testes, você pode usar o comando `go test` no terminal. Se todos os testes passarem, ele imprimirá `ok`. Se algum teste falhar, ele imprimirá uma descrição do teste que falhou e o que era esperado.