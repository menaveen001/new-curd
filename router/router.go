package router

import (
	"net/http"

	"example.com/m/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.POST("/user", controller.InsertUser)
	router.GET("/user", controller.GetAllUser)
	router.GET("/user/:id", controller.GetDataByID)
	router.POST("/user/:id", controller.UpdateById)
	router.DELETE("/:id", controller.DeleteByID)
	http.ListenAndServe(":3777", router)

}
