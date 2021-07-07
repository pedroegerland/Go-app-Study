package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "0660bibo"
	dbname   = "alura_loja"
)

// ConectaComBancoDeDados -
func ConectaComBancoDeDados() *sql.DB {
	conexao := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
