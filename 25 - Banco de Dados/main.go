package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	stringConn := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConn)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão aberta com sucesso!")

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		log.Fatal(err)
	}
	//agora o defer retorna um erro e o objeto
	defer func(linhas *sql.Rows) {
		_ = linhas.Close()
	}(linhas)
	// retorna em bytes
	fmt.Println(linhas)
}
