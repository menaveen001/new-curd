package database

import (
	"database/sql"
	"fmt"
	"os"

	"example.com/m/entity"
)

var db *sql.DB

func CreatConnectionTodb() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	var err error
	db, err = sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+databaseName)
	if err != nil {

		fmt.Println("error validating sql.Open argument")
		//panic(err.Error())
		//	return nil
	}

	fmt.Println("successful connection to Database")
	//return db
}

func InsertUser(student entity.Student) int64 {
	CreatConnectionTodb()
	insForm, err := db.Prepare("INSERT INTO student(firstname,lastname,id) VALUES(?,?,?)")
	if err != nil {
		fmt.Println(err)
		return 0
	}
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
	defer db.Close()
	fmt.Printf("rows inserted,%v\n", rows)
	return rows

}
func UpdateUser(student entity.Student) {
	CreatConnectionTodb()
	updateRes, err := db.Prepare("update student set firstname=? where id=?")
	if err != nil {
		fmt.Println(err)
	}

	// execute
	res, err := updateRes.Exec(student.Firstname, student.Id)
	if err != nil {
		fmt.Println(err, res)
	}
	a, err := res.RowsAffected()
	if err != nil {
		fmt.Println("data is not updated", err)

	}

	fmt.Printf("data is updated ,%v\n", a)

}
