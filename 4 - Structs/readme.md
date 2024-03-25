# Structs 

## Structs em Go

Em Go, as structs são tipos de dados compostos que permitem agrupar diferentes tipos de dados sob um único nome. Elas são úteis para representar registros de dados ou objetos complexos. Podemos fazer uma analogia com Structs e Classes em outras linguagens, como Java ou Python.

## Declaração de Structs

Uma struct em Go é declarada da seguinte forma:

```
type NomeDaStruct struct {
    campo1 tipo1
    campo2 tipo2
    // ...
}
```
- type: Palavra-chave utilizada para definir um novo tipo.
- NomeDaStruct: Nome da struct.
- campo1, campo2, etc.: Campos da struct, que representam diferentes partes ou características do objeto.
- tipo1, tipo2, etc.: Tipos de dados dos campos.

```
type Pessoa struct {
    nome   string
    idade  int
    email  string
    altura float64
}
```