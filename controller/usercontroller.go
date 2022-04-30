package controller

import (
	"log"
	"strconv"

	"example.com/m/database"
	"example.com/m/entity"
	"github.com/gin-gonic/gin"
)

// insertdata into db
func InsertUser(c *gin.Context) {
	var student entity.Student
	if err := c.BindJSON(&student); err != nil {
		log.Printf("Invalid Request JSON %v", err)
	}

	result := database.InsertUser(student)
	c.JSON(200, result)

}

// get all data from db
func GetAllUser(c *gin.Context) {

	c.JSON(200, database.SelectData())
}

// get data by id from db

func GetDataByID(c *gin.Context) {
	var student entity.Student
	id := c.Param("id")
	nid, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err, nid, student)
	}

	c.JSON(200, database.SelectDataByID(nid))

}

// delete data by id from db

func DeleteByID(c *gin.Context) {
	var student entity.Student
	id := c.Param("id")
	c.BindJSON(&student)
	nid, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err, nid)

	}
	c.JSON(200, database.DeleteDataByID(nid))
	log.Println("data is deleted")

}

// update data by id into db
func UpdateById(c *gin.Context) {

	var student entity.Student

	id := c.Param("id")
	// c.BindJSON(&student)
	nid, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err, nid)

	}
	if err := c.BindJSON(&student); err != nil {
		log.Printf("Invalid Request JSON %v", err)
	}
	c.JSON(200, database.UpdateUserById(student))
	log.Println(student)
}
