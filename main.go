package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	id        int
	firstname string
	lastname  string
}

func main() {
	db, err := sql.Open("mysql", "root:ishu123@tcp(localhost:3306)/userdb")
	if err != nil {

		fmt.Println("error validating sql.Open argument")
		panic(err.Error())

	}
	defer db.Close()

	//	err = db.Ping()
	//if err != nil {

	//	fmt.Println("error verifying connection with db.Ping")
	//	panic(err.Error())

	// insForm, err := db.Prepare("INSERT INTO student(id,firstname,lastname) VALUES(?,?,?)")
	// if err != nil {
	// 	fmt.Println("data is added")

	// }
	// res, err := insForm.Exec(2, "amren", "y")
	// if err != nil {
	// 	fmt.Println("query not executed", err)

	// }
	// //if err != nil {

	// //	}
	// rows, err := res.LastInsertId()
	// if err != nil {
	// 	fmt.Println("data is added")

	// }
	// //	if err != nil {

	// //	}
	// //}
	//fmt.Println("rows inserted", rows)

	queryRes, err := db.Query("SELECT * FROM student")
	if err != nil {
		fmt.Println(err, queryRes)
	}
	for queryRes.Next() {

		var student Student
		err := queryRes.Scan(&student.id, &student.firstname, &student.lastname)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%v\n", student)
	}
	updateRes, err := db.Prepare("update student set firstname=? where id=?")
	if err != nil {
		fmt.Println(err)
	}

	// execute
	res, err := updateRes.Exec("ankush", "1")
	if err != nil {
		fmt.Println(err, res)
	}
	a, err := res.RowsAffected()
	if err != nil {
		fmt.Println("data is not updated", err)

	}

	fmt.Printf("data is updated ,%v\n", a)

	insForm, err := db.Prepare("INSERT INTO student(id,firstname,lastname) VALUES(?,?,?)")
	if err != nil {
		fmt.Println(err)

	}
	res, err = insForm.Exec(4, "ankush", "pal")
	if err != nil {
		fmt.Println("query not executed", err)

	}

	rows, err := res.LastInsertId()
	if err != nil {
		fmt.Println("data is added")

	}

	fmt.Printf("rows inserted,%v\n", rows)

	fmt.Println("successful connection to Database")
}
