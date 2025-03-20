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
	Id       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
}

func (U *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userPage(w, r)
}

func userPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == http.MethodGet {
		d := "users"
		rows, err := db.QueryContext(ctx, "SELECT * from "+d)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ") // Optional: For pretty-printed JSON
		if err := encoder.Encode(users); err != nil {
			log.Fatal(err)
		}
	}
	if r.Method == http.MethodPost {
		userAgent := r.Header.Get("User-Agent")
		fmt.Println(userAgent)
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
			result, err := db.ExecContext(
				ctx,
				"INSERT INTO users (name,email,password) VALUES(?,?,?)",
				new_user.Name, new_user.Email, new_user.Password,
			)
			if err != nil {
				log.Fatal(err)
			}
			lq, err := result.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintln(w, lq)
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
