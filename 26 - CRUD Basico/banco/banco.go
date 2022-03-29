package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Driver de conexão obrigatório para o SQL - tem que ser adicionado manualmente
)

func Conectar() (*sql.DB, error) {
	stringConn := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConn)
	// checa se houve um erro ao abrir a conexão
	if err != nil {
		return nil, err
	}
	// faz literalmente um PING no banco
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
