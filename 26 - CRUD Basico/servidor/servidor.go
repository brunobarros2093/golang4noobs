package servidor

import (
	"crud-basico/banco"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	id    int32  `json: "id"`
	nome  string `json: "nome"`
	email string `json: "email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := ioutil.ReadAll(r.Body)

	if err != nil {
		//Não farei uso do retornos
		_, _ = w.Write([]byte("Falha ao ler corpo da requisição!"))
		return
	}
	// a requisição vem com um corpo
	var usuario usuario
	if err = json.Unmarshal(bodyReq, &usuario); err != nil {
		_, _ = w.Write([]byte("Erro ao converter usuário!"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		_, _ = w.Write([]byte("Erro ao conectar no banco!"))
		return
	}
	defer db.Close()
	// prepare statement
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?,?)")
	if erro != nil {
		_, _ = w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer func(statement *sql.Stmt) {
		_ = statement.Close()

	}(statement)
	insert, erro := statement.Exec(usuario.nome, usuario.email)
	if erro != nil {
		_, _ = w.Write([]byte("Erro ao executar statement"))
		return
	}
	id, err := insert.LastInsertId()
	if err != nil {
		_, _ = w.Write([]byte("Erro ao retornar Id inserido"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", id)))

}
