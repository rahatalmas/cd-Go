package controllers

import (
	"encoding/json"
	"fmt"
	"ginserver/db"
	"ginserver/models"
	"ginserver/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(200, gin.H{"message": "Invalid Data"})
		return
	}
	accessToken, err := utils.AccessToken(user.Name)
	if err != nil {
		fmt.Println("token err: ", err)
	}
	refreshToken, err2 := utils.AccessToken(user.Name)
	if err != nil {
		fmt.Println("token err: ", err2)
	}
	user.RefreshToken = refreshToken
	_, err = db.DB.Exec(`
	    INSERT INTO user 
		  (user_name, user_password, user_contact, role_key, refresh_token)
		VALUES 
		  (?, ?, ?, ?, ?)`,
		user.Name, user.Password, user.Contact, user.Role, user.RefreshToken,
	)

	if err != nil {
		fmt.Println("Error executing query:", err)
	}
	//jsonString := fmt.Sprintf(`{"user_name":"%s","password":"%s"}`, user.Name, user.Password)
	//jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{"message": "success", "accesstoken": accessToken, "user": user})
}

func SignUpDummy(c *gin.Context) {
	var user models.User
	fmt.Println(c.Request)
	if err := c.BindJSON(&user); err != nil {
		c.JSON(200, gin.H{"message": "Invalid Data"})
		return
	}
	accessToken, err := utils.AccessToken(user.Name)
	if err != nil {
		fmt.Println("token err: ", err)
	}
	refreshToken, err2 := utils.AccessToken(user.Name)
	if err != nil {
		fmt.Println("token err: ", err2)
	}
	user.RefreshToken = refreshToken
	file, err := os.OpenFile("dummydata/dummydatafile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file opened", file.Name())
	//jsonString := fmt.Sprintf(`{"user_name":"%s","password":"%s"}`, user.Name, user.Password)
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(string(jsonData))
	c.JSON(200, gin.H{"message": "success", "accesstoken": accessToken, "user": user})
}
