package go_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "INSERT INTO customer(id, name) VALUES('dono', 'Eko')"
	_, err := db.ExecContext(ctx, querySql)

	if err != nil {
		panic(err)
	}

	fmt.Println("Berhasil Insert di DB")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	querySql := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, querySql)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id ", id)
		fmt.Println("name ", name)
	}

}
