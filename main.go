package main

import (
	"fmt"

	"example.com/m/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")

	}
	router.InitRouter()

}
