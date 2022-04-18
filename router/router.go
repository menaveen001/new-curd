package router

import (
	"net/http"

	"example.com/m/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.POST("/user", controller.InsertUser)
	http.ListenAndServe(":3777", router)

}
