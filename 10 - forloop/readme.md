# For Loop 

O loop for em Go é uma estrutura de controle que permite repetir um bloco de código várias vezes. Ele é especialmente útil quando você sabe antecipadamente quantas vezes deseja executar um determinado bloco de código.

A sintaxe básica do loop for em Go é a seguinte:

```
for inicialização; condição; pós-instrução {
    // bloco de código a ser repetido
}
```
Aqui está o que cada parte da sintaxe significa:

- inicialização: é onde você pode inicializar uma variável de controle do loop. Isso geralmente é usado para definir o valor inicial da variável.
- condição: é uma expressão booleana que é verificada antes de cada iteração do loop. Se a condição for verdadeira, o bloco de código dentro do loop será executado. Se a condição for falsa, o loop será encerrado.
- pós-instrução: é uma instrução que é executada após cada iteração do loop. Geralmente é usado para atualizar a variável de controle do loop.

Aqui está um exemplo simples de um loop for em Go que imprime os números de 1 a 5:

```
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```
Saída esperada: 
```
1
2
3
4
5
```

Você pode personalizar o comportamento do loop for de acordo com suas necessidades, ajustando a inicialização, a condição e a pós-instrução. O loop for em Go é uma ferramenta poderosa para executar tarefas repetitivas de forma eficiente.