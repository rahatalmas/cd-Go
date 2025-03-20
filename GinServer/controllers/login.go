package controllers

import (
	"fmt"
	"ginserver/models"
	"ginserver/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var authData models.Auth
	fmt.Println(c.Request)
	if err := c.BindJSON(&authData); err != nil {
		c.JSON(200, gin.H{"message": "Invalid Data"})
		return
	}
	if authData.Name == "" || authData.Password == "" {
		c.JSON(200, gin.H{"message": "Invalid Data"})
		return
	}
	accessToken, err := utils.AccessToken(authData.Name)
	if err != nil {
		fmt.Println("token err: ", err)
	}
	c.JSON(200, gin.H{"message": "success", "accesstoken": accessToken})
}
