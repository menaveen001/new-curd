package controller

import (
	"log"
	"strconv"

	"example.com/m/database"
	"example.com/m/entity"
	"github.com/gin-gonic/gin"
)

func InsertUser(c *gin.Context) {
	var student entity.Student
	if err := c.BindJSON(&student); err != nil {
		log.Printf("Invalid Request JSON %v", err)
	}

	result := database.InsertUser(student)
	c.JSON(200, result)

}

func GetAllUser(c *gin.Context) {

	c.JSON(200, database.SelectData())
}

func GetDataByID(c *gin.Context) {
	var student entity.Student
	id := c.Param("id")
	nid, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err, nid, student)
	}

	c.JSON(200, database.SelectDataByID(nid))

}

// func UpdateById(c *gin.Context) {
// 	var student entity.Student
// 	c.BindJSON(&student)
// 	c.Param("id")

// 	nid, err := strconv.Atoi(id)
// 	if err != nil {
// 		log.Println(err, nid)

// 	}
// 	c.JSON(200, database.UpdateUserById(student))

// }

func DeleteByID(c *gin.Context) {
	var student entity.Student
	id := c.Param("id")
	c.BindJSON(&student)
	nid, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err, nid)

	}
	c.JSON(200, database.DeleteDataByID(nid))

}
