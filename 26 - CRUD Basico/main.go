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
	// registro de rotas e funções
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods("POST")
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods("GET")
	router.HandleFunc("/usuario/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	fmt.Println("Escutando na Porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
