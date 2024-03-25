# Métodos 

Um método é uma função associada a um tipo específico. É semelhante a uma função, mas tem um "receiver" especial que define em qual tipo este método deve ser anexado.

No nosso código, definimos um método chamado `salvar` para o tipo `usuario`. O receiver do método é `(u usuario)`. Isso significa que o método `salvar` pode ser chamado em qualquer valor do tipo `usuario`.

```go
func (u usuario) salvar() {
    fmt.Println(u.idade)
    fmt.Println("Metodo salvar do usuario")
}
```

No método `salvar`, `u` é uma cópia do `usuario` original sobre o qual o método foi chamado. Isso significa que se você modificar `u` dentro do método `salvar`, isso não afetará o `usuario` original.

No entanto, se você quiser modificar o `usuario` original dentro do método, você pode usar um ponteiro para o `usuario` como o receiver. Isso seria escrito como `(u *usuario)`.

Na função `main`, você cria um `usuario` e chama o método `salvar` nele:

```go
func main() {
    usuario1 := usuario{"usuario1", 20}
    fmt.Println(usuario1)
    usuario1.salvar()
}
```

Quando você executa `usuario1.salvar()`, o método `salvar` é chamado no `usuario1`, e a idade do `usuario1` é impressa, seguida pela string "Metodo salvar do usuario".