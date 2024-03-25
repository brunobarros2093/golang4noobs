# Maps 

Um map em Go é uma estrutura de dados que associa chaves únicas a valores correspondentes como Map<chave,valor>. 

```
// Usando make
meuMap := make(map[string]int)

// Map literal
meuMap := map[string]int{
    "chave1": valor1,
    "chave2": valor2,
}
```
Os valores de um map são acessados e modificados usando a sintaxe de colchetes [].

```
meuMap["novaChave"] = novoValor
valor := meuMap["chaveExistente"]
```
## Verificação de Existência de Chave 
Você pode verificar se uma chave existe em um map usando uma atribuição múltipla.

```
valor, existe := meuMap["chave"]
if existe {
    // A chave existe
} else {
    // A chave não existe
}
```

Você pode remover uma chave de um map usando a função delete.

```
delete(meuMap, "chave")
```

Você pode iterar sobre os pares chave-valor de um map usando um loop for range.

```
for chave, valor := range meuMap {
    fmt.Println(chave, valor)
}
```
O comprimento de um map pode ser obtido usando a função len: 

```
comprimento := len(meuMap)
```
Maps em Go são úteis para armazenar dados associativos, onde cada valor é identificado por uma chave exclusiva.