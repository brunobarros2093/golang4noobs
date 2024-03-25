# Defer 

O defer é uma palavra-chave em Go que permite que você adie a execução de uma função até que a tudo tenha sido executado, seja normalmente ou através de um panic. É comumente usado para limpeza, como fechar um arquivo ou desbloquear um mutex.

Aqui está um exemplo de como você pode usar defer:

```
func main() {
    f, _ := os.Open("file.txt")
    defer f.Close()

    // Operações com o arquivo
}
```
Neste exemplo, <b>defer f.Close()</b> garante que o arquivo será fechado quando a função main retornar, não importa em que ponto isso aconteça. Isso é útil para garantir que os recursos sejam limpos, mesmo em caso de erro.