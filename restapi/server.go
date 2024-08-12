package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx       context.Context = context.Background()
	db, dberr                 = sql.Open("mysql", "root:1234@/users")
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func (U *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userPage(w, r)
}

func userPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == http.MethodGet {
		rows, err := db.QueryContext(ctx, "SELECT id,name,email from users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var users []User
		for rows.Next() {
			var id int
			var name, email string
			if err := rows.Scan(&id, &name, &email); err != nil {
				log.Fatal(err)
			}
			users = append(users, User{Id: id, Name: name, Email: email})
			fmt.Println(id, " ", name, " ", email)
		}
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(users); err != nil {
			log.Fatal(err)
		}
	}
	if r.Method == http.MethodPost {
		//e := r.ParseForm()
		//fmt.Println(e)
		t := r.Header.Get("content-type")
		if t == "application/json" {
			var new_user User
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&new_user); err != nil {
				log.Fatal(err)
			}
			fmt.Println(new_user)
		} else {
			fmt.Println("type: ", t)
			name := r.FormValue("name")
			email := r.FormValue("email")
			password := r.FormValue("password")
			fmt.Println(name, " ", email)
			result, err := db.ExecContext(
				ctx,
				"INSERT INTO users (name,email,password) VALUES(?,?,?)",
				name, email, password,
			)
			if err != nil {
				log.Fatal(err)
			}
			lq, err := result.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintln(w, lq)
		}
	}
}

func main() {
	if dberr != nil {
		log.Fatal(dberr)
	}

	fmt.Println("Server started on port 5000")
	http.Handle("/api/user", new(User))

	log.Fatal(http.ListenAndServe(":5000", nil))
}
