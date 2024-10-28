package main

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("hello world");
	dsn := "root:@/academia_db_main"
	db,err := sql.Open("mysql",dsn);
	defer db.Close()
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(db)
	if err := db.Ping(); err!=nil{
		fmt.Println(err)
	}
	fmt.Println("database connected")
	fmt.Println("1. Query \t 2.CreateTable")
	var s int
	fmt.Scanln(&s)
	if(s == 1){
		var tableNames []string
		query := "SELECT table_name FROM information_schema.tables WHERE table_schema = 'academia_db_main'"
		rows,err := db.Query(query)
		if err != nil{
			fmt.Println(err)
		}
		defer rows.Close();
		for rows.Next() {
			var tableName string 
			if err:= rows.Scan(&tableName); err!=nil{
				fmt.Println(err)
			}
			tableNames = append(tableNames,tableName)
		}
		fmt.Println("Table name: ")
		for k,v := range tableNames{
			fmt.Println(k," - ",v)
		}
                   
		var choice int 
		fmt.Scanln(&choice)
		colTemplate := fmt.Sprintf("SELECT COUNT(*) FROM information_schema.columns WHERE table_schema = academia_db_main AND table_name = %s", tableNames[choice])
		qtemplate := fmt.Sprintf("SELECT * FROM %s",tableNames[choice])
		fmt.Println(colTemplate)
		var count int
		e := db.QueryRow(colTemplate).Scan(&count)
		if e!=nil{
			fmt.Println(err)
		}
		fmt.Println("count ",count)

		res,err := db.Query(qtemplate)
		if err!=nil{
			fmt.Println(err)
		}
		for res.Next(){
			var one,two,three,four string 
			if err := res.Scan(&one,&two,&three,&four);err!=nil{
				fmt.Println(err)
			}
			fmt.Println(one, two, three, four)
		}
	}else{
		m2mquery := `CREATE TABLE studentcourses (
		  student_id VARCHAR(200),
		  course_code VARCHAR(200),
		  PRIMARY KEY (student_id,course_code),
		  FOREIGN KEY (student_id) REFERENCES student(student_id),
		  FOREIGN KEY (course_code) REFERENCES course(course_code)
		);`
		res,err := db.Exec(m2mquery)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}