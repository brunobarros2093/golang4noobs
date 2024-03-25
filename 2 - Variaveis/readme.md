# Variáveis 
Em Go, as variáveis são espaços de armazenamento que contêm valores. Aqui está uma breve explicação das variáveis em Go:

Declaração: As variáveis em Go são declaradas usando a sintaxe var nome tipo. Por exemplo: var x int declara uma variável chamada x do tipo int.

Inicialização: As variáveis podem ser inicializadas durante a declaração, fornecendo um valor após o tipo. Por exemplo: var x int = 10 declara e inicializa a variável x com o valor 10.

Inferência de Tipo: Go permite inferência de tipo, onde o tipo pode ser omitido da declaração se o compilador puder inferi-lo do valor inicial. Por exemplo: x := 10 declara e inicializa x como int.

Escopo: As variáveis em Go têm escopo léxico, o que significa que estão disponíveis apenas dentro do bloco em que são declaradas, incluindo blocos de funções.

Tipos de Variáveis: Go suporta uma variedade de tipos de variáveis, incluindo tipos básicos como int, float, bool, etc., bem como tipos compostos como structs, arrays, slices, maps, channels, e pointers.

Mutabilidade: A mutabilidade das variáveis depende do tipo. Por exemplo, as variáveis de tipo int são mutáveis, enquanto as de tipo string são imutáveis.

Zero Values: Se uma variável é declarada sem ser inicializada explicitamente, ela recebe um "zero value" padrão. Por exemplo, 0 para números inteiros, 0.0 para floats, false para booleanos e "" para strings.