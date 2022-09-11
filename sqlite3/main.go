package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./sample.db")

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT)")
	statement.Exec()

	insert, _ := database.Prepare("INSERT INTO people VALUES (?, ?)")
	insert.Exec(1, "fulano")

	resultSet, _ := database.Query("SELECT * FROM people")
	for resultSet.Next() {
		var id int
		var firstname string

		resultSet.Scan(&id, &firstname)

		fmt.Println("user id:", id)
		fmt.Println("user name:", firstname)
	}
}
