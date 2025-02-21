package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func JsonResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("err response write - ", err)
	}
}

func (U *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user := []User{
		{
			Id:       1,
			Name:     "B.M. Rahat Almas",
			Password: "123456",
		},
		{
			Id:       2,
			Name:     "Pretty Dey",
			Password: "123456",
		},
		{
			Id:       3,
			Name:     "Itachi Uchiha",
			Password: "123456",
		},
	}
	path := r.URL.Path
	switch path {
	case "/user/alluser":
		JsonResponse(w, r, user)
	case "/user/post":
		fmt.Fprintf(w, "post req")
	default:
		fmt.Fprintf(w, "Not Found")
	}
}

func (P *Post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	post := []Post{
		{Id: 1, Title: "apple", Content: "apple is red"},
	}
	path := r.URL.Path
	switch path {
	case "/post/":
		JsonResponse(w, r, post)
	default:
		fmt.Fprintf(w, "not found")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/user/", new(User))
	mux.Handle("/post/", new(Post))
	fmt.Println("server is running on port 5000")
	http.ListenAndServe(":5000", mux)
}
