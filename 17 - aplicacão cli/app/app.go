package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Com G maiusculo, pois é exportável
// Retorna aplicação cli.App pronta para ser executada
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação linha de comando estudos"
	app.Usage = "Busca IPs e nomes de servidores na internet"
	// flags --host
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores online",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	// pacote Net busca os IPs
	// retorna um slice ou erro
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
func buscarServidores(c *cli.Context) {
	host := c.String("host")

	if servidores, err := net.LookupNS(host); err != nil {
		log.Fatal(err)
	} else {
		for _, name := range servidores {
			fmt.Println(name)
		}
	}
}
