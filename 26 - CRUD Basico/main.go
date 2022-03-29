package main

import (
	"crud-basico/servidor"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// contem as rotas para interagir com o BD
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods("POST")

	fmt.Println("Escutando na Porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
