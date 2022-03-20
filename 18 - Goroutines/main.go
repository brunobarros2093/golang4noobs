package main

import (
	"fmt"
	"time"
)

// CONCORRENCIA É DIFERENTE DE PARALELISMO
// Paralelismo: duas coisas ocorrendo exatamente ao mesmo tempo (processadores com mais de 1 nucleo)
// Concorrencia: Não necessariamente executam ao mesmo tempo (mas podem estar) - revesamento entre tarefas
func main() {

	// Grupo para sincronizar as GoRoutines e garantir que o programa só acabe após a execução de
	// tudo que está no grupo
	//var waitGroup sync.WaitGroup
	//	////Passando a quantidade de goroutines na fila
	//	//waitGroup.Add(2)
	//	//
	//	//go func() {
	//	//	// com Go Routines, basicamente o método abaixo é executado antes mesmo da conclusão do primeiro método
	//	//	//go escrever("Olá Mundo!")
	//	//	escrever("Olá Mundo!")
	//	//	waitGroup.Done() // -1
	//	//}()
	//	//
	//	//go func() {
	//	//	// sem goroutines a chamada do segundo método só é chamada após o término do primeiro
	//	//	escrever("Programando em Go Coroutines")
	//	//	waitGroup.Done() // -1
	//	//}()
	//	//
	//	//waitGroup.Wait() // quando zerar o programa termina execução

	fmt.Println("==== Channels ====")
	// cria um canal (chan) de string
	canal := make(chan string)

	go escrever("Rodando em primeiro", canal)
	//escrever("Rodando em segundo", canal)
	//----------aqui o canal esta "enviando" algo para fora dele
	// COM LOOP INFINITO
	//for {
	//	// checagem se o canal esta aberto ou fechado
	//	mensagem, aberto := <-canal
	//	if !aberto {
	//		break
	//	}
	//	fmt.Println(mensagem)
	//}
	// sem o loop infinito
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("---Fim do programa")
}

//-------------------------a função recebe um canal de string
func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// o canal esta RECEBENDO o texto
		// mandando um valor para dentro do canal
		canal <- texto
		time.Sleep(time.Second)
	}
	// fecha o canal após execução
	close(canal)
}
