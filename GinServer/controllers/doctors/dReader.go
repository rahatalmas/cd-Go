package doctors

import (
	"fmt"
	"ginserver/dummydata"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DoctorsList(c *gin.Context) {
	c.JSON(200, dummydata.Doctors)
}

func DoctorById(c *gin.Context) {
	id, err := c.Params.Get("id")
	if err {
		fmt.Println(err)
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	Id, err2 := strconv.Atoi(id)
	if err2 != nil {
		fmt.Println("Failed to convert string to integer")
		c.JSON(500, gin.H{"message": "Enternal Server Error"})
	}
	fmt.Println("Doctor Id: ", id)
	for _, v := range dummydata.Doctors {
		if v.Id == Id {
			c.JSON(200, v)
			return
		}
	}
	c.JSON(404, gin.H{"message": "Data Not Found"})
}
