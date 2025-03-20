package main

import (
	"fmt"
	"ginserver/controllers"
	"ginserver/controllers/doctors"
	"ginserver/controllers/users"
	"ginserver/db"

	"github.com/gin-gonic/gin"
)

// type OptionFunc func(*gin.Engine)
// type ginC *gin.Context

func HandleNotFound() gin.OptionFunc {
	return func(engine *gin.Engine) {
		engine.NoRoute(func(c *gin.Context) {
			c.HTML(404, "404.html", "GO-GIN")
		})
	}
}

func createEngine(options ...gin.OptionFunc) *gin.Engine {
	engine := gin.Default()
	for _, option := range options {
		option(engine)
	}
	return engine
}

func main() {
	db.InitDB()
	if err := db.DB.Ping(); err != nil {
		fmt.Println("failed to connect db")
	}
	// s := gin.Default()
	// s.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"name": "almas", "wife": "pretty"})
	// })
	engine := createEngine(HandleNotFound())
	engine.LoadHTMLGlob("templates/*") // -> load all html files from the templates directory
	//engine.LoadHTMLFiles("templates/404.html")  // -> loads the files by name given as parameters

	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"name": "Rahat Almas", "bio": "Software Engineer"})
	})

	engine.GET("api/v1/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"type": "Route Testing", "message": "success"})
	})

	engine.POST("api/v1/login", controllers.Login)
	engine.POST("api/v1/signup", controllers.SignUp)

	engine.GET("api/v1/users", users.List)
	engine.GET("api/v1/doctors", doctors.DoctorsList)
	engine.GET("api/v1/doctors/:id", doctors.DoctorById)

	engine.Run("localhost:5000")
}
