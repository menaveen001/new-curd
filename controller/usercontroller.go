package controller

import (
	"log"

	"example.com/m/database"
	"example.com/m/entity"
	"github.com/gin-gonic/gin"
)

func InsertUser(c *gin.Context) {
	var student entity.Student
	if err := c.BindJSON(&student); err != nil {
		log.Printf("Invalid Request JSON %v", err)
	}
	// student.Id = 11
	// student.Firstname = "amren"
	// student.Lastname = "y"
	result := database.InsertUser(student)
	c.JSON(200, result)

}
