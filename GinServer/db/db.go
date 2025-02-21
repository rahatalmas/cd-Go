package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("mysql", "root:@/pet_vet_bd")
	if err != nil {
		fmt.Println("failed to connect database")
	}
}
