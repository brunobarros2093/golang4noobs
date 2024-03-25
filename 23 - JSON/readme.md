# JSON - Marshal e Unmarshal

Em Go, `encoding/json` é o pacote que fornece funções para trabalharmos com JSON. As duas funções principais que você usará são `json.Marshal` e `json.Unmarshal`.

**json.Marshal**

`json.Marshal` é usado para converter um objeto Go (como uma struct ou um mapa) em um JSON. Aqui está um exemplo:

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "John", Age: 30}
    jsonData, err := json.Marshal(p)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(string(jsonData))
}
```

Neste exemplo, `json.Marshal` pega uma instância de `Person` e retorna uma representação JSON dessa instância como uma fatia (slice) de bytes. Se houver um erro durante a marshalling, ele será retornado como o segundo valor.

**json.Unmarshal**

`json.Unmarshal` é o oposto de `json.Marshal`. Ele é usado para converter um JSON em um objeto Go. Aqui está um exemplo:

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    jsonData := []byte(`{"Name":"John","Age":30}`)
    var p Person
    err := json.Unmarshal(jsonData, &p)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(p)
}
```

Neste exemplo, `json.Unmarshal` pega uma fatia de bytes contendo um JSON e um ponteiro para uma instância de `Person`. Ele preenche a instância de `Person` com os dados do JSON. Se houver um erro durante a unmarshalling, ele será retornado.

Note que os nomes dos campos na struct `Person` começam com letras maiúsculas. Isso é necessário porque o pacote `encoding/json` só pode acessar campos exportados. Se os nomes dos campos começarem com letras minúsculas, eles serão ignorados pelo `json.Marshal` e `json.Unmarshal`.