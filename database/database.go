package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"crudapp/entity"
)

//make connection to db

var db *sql.DB

func CreatConnectionTodb() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	var err error
	db, err = sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+databaseName)
	if err != nil {

		fmt.Println("error validating sql.Open argument")

	}

	log.Println("successful connection to Database")

}

//insert data into db

func InsertUser(student entity.Student) int64 {
	CreatConnectionTodb()
	insForm, err := db.Prepare("INSERT INTO student(firstname,lastname,id) VALUES(?,?,?)")
	if err != nil {
		fmt.Println(err)
		return 0
	}

	//execute

	res, err := insForm.Exec(student.Firstname, student.Lastname, student.Id)
	if err != nil {
		fmt.Println("query not executed", err)
		return 0
	}

	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows not inserted", err)
		return 0
	}
	log.Printf("rows inserted,%v\n", rows)
	return rows

}

// updateUser by id

func UpdateUserById(student entity.Student) int64 {
	CreatConnectionTodb()
	updateRes, err := db.Prepare("update student set firstname = ?, lastname =? where id = ?")
	if err != nil {
		fmt.Println(err)
	}

	// execute
	res, err := updateRes.Exec(student.Firstname, student.Lastname, student.Id)
	if err != nil {
		fmt.Println(err, res)
	}
	a, err := res.RowsAffected()
	if err != nil {
		fmt.Println("data is not updated", err, a)

	}

	fmt.Printf("data is updated ,%v\n", a)
	log.Println(a)
	return 0

}

// shows all data
func SelectData() []entity.Student {

	CreatConnectionTodb()
	var student1 []entity.Student
	queryRes, err := db.Query("SELECT * FROM student")
	if err != nil {
		fmt.Println(err, queryRes)
	}
	for queryRes.Next() {

		var student entity.Student
		err := queryRes.Scan(&student.Id, &student.Firstname, &student.Lastname)

		if err != nil {
			fmt.Println(err)
		}
		student1 = append(student1, student)
	}
	log.Println(student1)
	return student1
}

// get data by id
func SelectDataByID(id int) []entity.Student {

	CreatConnectionTodb()
	var student1 []entity.Student
	queryRes, err := db.Query("SELECT *FROM student WHERE id =?", id)
	if err != nil {
		fmt.Println(err, queryRes)
	}

	if queryRes.Next() {
		var student entity.Student
		err := queryRes.Scan(&student.Id, &student.Firstname, &student.Lastname)
		if err != nil {
			log.Println(err)

		}
		student1 = append(student1, student)

	}

	log.Println(student1)
	return student1

}

// delete by id
func DeleteDataByID(id int) []entity.Student {

	CreatConnectionTodb()
	var student1 []entity.Student
	queryRes, err := db.Query("delete from student where id = ?", id)
	if err != nil {
		fmt.Println(err, queryRes)
	}

	if queryRes.Next() {
		var student entity.Student
		err := queryRes.Scan(&student.Id, &student.Firstname, &student.Lastname)
		if err != nil {
			log.Println(err)

		}
		student1 = append(student1, student)

	}

	log.Println(student1)
	return student1

}
