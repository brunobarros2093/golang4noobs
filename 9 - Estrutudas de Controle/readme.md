# Switch 

O switch em Go é uma estrutura de controle de fluxo que permite fazer comparações entre uma expressão e uma lista de possíveis valores.

## Sintaxe Básica

A sintaxe básica do switch em Go é a seguinte:

```
switch expressao {
case valor1:
    // Código a ser executado se a expressao for igual a valor1
case valor2:
    // Código a ser executado se a expressao for igual a valor2
default:
    // Código a ser executado se a expressao não corresponder a nenhum dos casos anteriores
}
```
- A expressao é avaliada uma vez.
- Cada case especifica um valor a ser comparado com a expressão.
- O default é opcional e é executado se nenhum dos case corresponder à expressão.

Exemplo

```
func main() {
    dia := "quarta"

    switch dia {
    case "segunda", "terça", "quarta", "quinta", "sexta":
        fmt.Println("Dia útil")
    case "sábado", "domingo":
        fmt.Println("Final de semana")
    default:
        fmt.Println("Dia inválido")
    }
}
```
## Uso Avançado

O switch em Go pode ser usado sem uma expressão, o que permite uma lógica mais complexa em cada case. Também permite o uso de expressões booleanas e a verificação de tipos de variáveis.

## Considerações Finais

- O switch em Go é uma maneira eficiente e limpa de lidar com várias condições.
- É comumente usado em vez de múltiplas declarações if-else.
- Se múltiplos case correspondem à expressão, apenas o primeiro case correspondente é executado.