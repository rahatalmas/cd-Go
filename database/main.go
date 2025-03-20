package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var username, host, password, dbname string
	fmt.Print("Enter User name: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)
	fmt.Print("Enter database name: ")
	fmt.Scanln(&dbname)
	fmt.Println(host, password, dbname)
	var dbinfo = fmt.Sprintf("%s:%s@/%s", username, password, dbname)
	fmt.Println("mysql", dbinfo)
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		fmt.Println("Failed To Open Database: ", err.Error())
	}
	cerr := db.Ping()
	if cerr != nil {
		fmt.Println("Failed To Open Database: ", cerr.Error())
	} else {
		fmt.Println("db:", db)
		fmt.Println("Welcome Back\nHow Can i Help you?")
		for {
			var option int
			fmt.Println("1.Manage DB \t 2.Execute Query\n3.DB Info \t 4.help")
			fmt.Scanln(&option)
			if option == 3 {
				fmt.Println("Query Result: ")
				rows, err := db.Query("SHOW TABLES")
				if err != nil {
					fmt.Println("Fail to execute Query: ", err.Error())
				}
				for rows.Next() {
					var tableName string
					if err := rows.Scan(&tableName); err != nil {
						fmt.Println("failed to scan database: ", err.Error())
						break
					}
					fmt.Println(tableName)
				}
				fmt.Println("\nWhats Next? ")
			} else {
				break
			}
		}
	}
}
