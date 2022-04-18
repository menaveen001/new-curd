package main

import (
	"fmt"

	"example.com/m/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// type Student struct {
// 	id        int
// 	firstname string
// 	lastname  string
// }

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	router.InitRouter()
	// ername := os.Getenv("USERNAME")
	// password := os.Getenv("PASSWORD")
	// databaseName := os.Getenv("DATABASE_NAME")
	// db, err := sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+databaseName)
	// if err !us= nil {

	// 	fmt.Println("error validating sql.Open argument")
	// 	panic(err.Error())

	// }
	// defer db.Close()
	// fmt.Println("successful connection to Database")

	// queryRes, err := db.Query("SELECT * FROM student")
	// if err != nil {
	// 	fmt.Println(err, queryRes)
	// }
	// for queryRes.Next() {

	// 	var student Student
	// 	err := queryRes.Scan(&student.id, &student.firstname, &student.lastname)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Printf("%v\n", student)
	// }

	// deleteRes, err := db.Prepare("delete  student set firstname=? where id=?")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// res, err := deleteRes.Exec("praveen", 5)
	// if err != nil {
	// 	fmt.Println(err, res)
	// }
	// b, err := res.RowsAffected()
	// if err != nil {
	// 	fmt.Println("not delete", err)
	// }
	// fmt.Printf("deleted ,%v\n", b)
}
