package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	DB    *sql.DB
	DBERR error
	Posts Post
}

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

func (p *Post) Blogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//ctx := context.Background()
	//db.QueryContext(ctx,"get query")
	//db.ExecContext(ctx,"post query")
	switch r.Method {
	case http.MethodGet:
		var posts []Post
		if rows, err := db.Query("SELECT * FROM post"); err == nil {
			var post Post
			for rows.Next() {
				err := rows.Scan(&post.Id, &post.Title, &post.Details)
				if err != nil {
					log.Fatal("error fetching")
				}
				posts = append(posts, post)
			}
		}
		data, err := json.Marshal(posts)
		if err != nil {
			log.Fatal("error json encode")
		}
		w.Write(data)
	case http.MethodPost:
		title := r.FormValue("title")
		details := r.FormValue("details")
		var query string = "INSERT INTO post (title,details) VALUES(?,?)"
		row, err := db.Query(query, title, details)
		if err != nil {
			log.Fatal("error post query")
		}
		var result []Post
		for row.Next() {
			var post Post
			err := row.Scan(&post.Title, &post.Details)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, post)
		}
		data, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(data)
	case http.MethodPut:
		w.Write([]byte("update request"))
	case http.MethodDelete:
		w.Write([]byte("delete request"))
	}
}

func (S *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server")
	S.DB, S.DBERR = sql.Open("mysql", "root:1234@/post")
	if S.DBERR == nil {
		fmt.Println("db connected")
		switch r.URL.Path {
		case "/":
			w.Write([]byte("home page"))
		case "/api/blogs":
			S.Posts.Blogs(w, r, S.DB)
		default:
			fmt.Fprintf(w, "Wrong Routes request")
		}
	} else {
		fmt.Println(S.DBERR)
	}
}

func main() {
	http.ListenAndServe(":8000", new(Server))
}

func C(w http.ResponseWriter, r *http.Request) {
	client := new(http.Client)
	res, err := client.Get("http://software.diu.edu.bd:8006/result?grecaptcha=&semesterId=241&studentId=203-15-3914")
	if err != nil {
		fmt.Println("error req")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("io read error")
	}
	w.Write([]byte("not found"))
	fmt.Println(string(data))
}
