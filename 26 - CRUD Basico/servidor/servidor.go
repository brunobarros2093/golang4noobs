package servidor

import (
	"crud-basico/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Letra maiuscula permite acesso as propriedades da Struct, se deixar letra minuscula não é alterado!
type usuario struct {
	id    int32  `json: "id"`
	Nome  string `json: "nome"`
	Email string `json: "email"`
}

// CriarUsuario - Cria usuários e insere no banco
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
	defer statement.Close()

	insert, erro := statement.Exec(usuario.Nome, usuario.Email)
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

// BuscarUsuarios -  Traz uma lista de usuários do banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco!"))
		return
	}
	defer db.Close()

	linhasDoDb, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro buscar usuários"))
		return
	}
	defer linhasDoDb.Close()

	var usuarios []usuario
	for linhasDoDb.Next() {
		var usuario usuario
		if erro := linhasDoDb.Scan(&usuario.id, &usuario.Nome, &usuario.Email); erro != nil {
			// Estou ignorando e não tratando o erro que o Write pode retornar!
			w.Write([]byte("Erro ao scannear usuário"))
			return
		}
		// preenche o slice acima
		usuarios = append(usuarios, usuario)
	}
	// retorna o status 200
	w.WriteHeader(http.StatusOK)
	// transformando o slice de usuários em json
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro converter usuário para JSON"))
		return
	}

}

// BuscarUsuario -  Traz um usuário especifico do banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}
