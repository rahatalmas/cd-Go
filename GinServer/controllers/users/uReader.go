package users

import (
	"encoding/json"
	"fmt"
	"ginserver/db"
	"ginserver/models"
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListDummy(c *gin.Context) {
	data, err := os.ReadFile("dummydata/dummydatafile.txt")
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	fmt.Println(string(data))
	var users []models.Auth
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	for {
		var user models.Auth
		err := decoder.Decode(&user)
		if err == io.EOF {
			break
		} else if err != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		users = append(users, user)
	}
	fmt.Println(users)
	c.JSON(200, users)
}

func List(c *gin.Context) {
	rows, err1 := db.DB.Query("SELECT * FROM user")
	if err1 != nil {
		fmt.Println("failed to fetch users")
		c.JSON(500, gin.H{"message": "internal server error"})
		return
	}
	var users []models.User
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Picture, &user.Contact, &user.Role, &user.RefreshToken)
		if err != nil {
			fmt.Println(user)
			fmt.Println("failed to scan user rows", err.Error())
			//continue
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("failed to scan fetch rows", err.Error())
		c.JSON(500, gin.H{"message": "internal server error"})
		return
	}
	c.JSON(200, users)
}
