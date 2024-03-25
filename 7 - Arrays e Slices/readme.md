# Arrays e Slices

## Arrays

Declaração: Um array em Go é uma coleção ordenada de elementos de um tipo específico, com tamanho <b>fixo</b> e pré-definido. A sintaxe para declarar um array é:

```
var numeros [3]int            // Array de inteiros com tamanho 3
numeros := [3]int{1, 2, 3}    // Inicialização com valores
```
## Slices

Declaração: Um slice em Go é uma estrutura de dados dinâmica construída sobre um array subjacente, que fornece uma interface mais poderosa e flexível para trabalhar com sequências de dados. Slices podem ser criados usando a função make ou através da notação de slice em arrays existentes.A sintaxe para declarar um slice é:

```
s := make([]int, 3, 5)     // Cria um slice de inteiros com comprimento 3 e capacidade 5
s := []int{1, 2, 3, 4, 5}  // Cria um slice literal
```

Slices oferecem funcionalidades adicionais, como alterar o tamanho do slice dinamicamente, fatiar (criar sub-slices) e trabalhar com a capacidade subjacente.

```
s := make([]int, 3)         // Cria um slice de inteiros com comprimento 3
s = append(s, 4)             // Adiciona um elemento ao slice
s = append(s, 5, 6, 7)       // Adiciona múltiplos elementos ao slice
s = s[1:4]                   // Fatiamento de slice
```

## Comparação

- Tamanho: Arrays têm tamanho fixo, enquanto slices têm tamanho dinâmico.
- Flexibilidade: Slices são mais flexíveis e versáteis do que arrays.
- Eficiência: Arrays tendem a ser mais eficientes em termos de memória e processamento do que slices, devido à alocação contígua de memória.

Em geral, slices são muito mais comuns em Go devido à sua flexibilidade e praticidade em comparação com arrays fixos. Eles são amplamente utilizados para manipular coleções de dados dinâmicas e são uma parte fundamental da linguagem Go.