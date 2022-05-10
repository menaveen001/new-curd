package router

import (
	"net/http"

	"crudapp/controller"
	"github.com/gin-gonic/gin"
)

// InitRouter() : is used to initialize the routes for the application
func InitRouter() {
	router := gin.Default()
	router.POST("/login",controller.Login)
	//add new routes here
	router.POST("/user", controller.InsertUser)
	router.GET("/user", controller.GetAllUser)
	router.GET("/user/:id", controller.GetDataByID)
	router.POST("/user/update", controller.UpdateById)
	router.DELETE("/:id", controller.DeleteByID)
	http.ListenAndServe(":3777", router)

}
