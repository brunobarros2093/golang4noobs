package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

//-----------------\/quem envia a response para o cliquente
// ------------------------------------\/-quem fez a requisição
func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{"Pedro", "pedrim@gmail.com"}
	//----------------------------------------------------\/-valores a serem passados para o HTML/front
	err := templates.ExecuteTemplate(w, "home.html", u)
	if err != nil {
		return
	}
}

func main() {

	http.HandleFunc("/home", home)
	// permitindo todos os arquivos .html - registrando no template
	templates = template.Must(template.ParseGlob("*.html"))
	fmt.Println("Rodando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
