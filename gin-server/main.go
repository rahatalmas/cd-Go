package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct{
	Id string `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Department string `json:"department" form:"department"`
}

func main(){
   fmt.Println("hello world")
   s := Student{Id:"1",Name:"Almas",Department:"Cse"}
   r:=gin.Default()
   r.LoadHTMLGlob("templates/*")
   //r.LoadHTMLFiles("templates/index.html","templates/student.html") // i can also pass multiple html files here using comma separation
   r.GET("/ping",func(c *gin.Context){
		c.JSON(200,s)
   })
   r.GET("/student",func(c *gin.Context){
	    c.JSON(200,gin.H{
			"message":"success",
			"student":s,
		})
   })
   r.GET("/html",func(c *gin.Context){
	   c.HTML(200,"index.html",gin.H{
		  "message":"success",
		  "student":s,
	   })
   })
   r.GET("/html2",func(c *gin.Context){
	    c.HTML(200,"student.html",s)
   })
   r.GET("/form",func(c *gin.Context){
        c.HTML(200,"form.html",gin.H{"message":"form"})
   })
   r.POST("/submit",func(c *gin.Context){
	    var s Student
		if err:= c.ShouldBind(&s); err!=nil{
			c.JSON(400,gin.H{"error":err.Error()})
			return
		}
		fmt.Println(s)
		c.HTML(200,"index.html",gin.H{"message":"successful form data","student":s})
   })
   r.Run()
}