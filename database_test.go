package go_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "gamal:adgjmptw12@tcp(localhost:3307)/go_db")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
