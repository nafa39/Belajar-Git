// Go & mysql connection
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Nafa3999!@tcp(127.0.0.1:3306)/Hacktiv8")
	if err != nil {
		log.Print("Error occured while connecting to the database")
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Print("Error occured while pinging to the database")
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}
