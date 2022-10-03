package main

import (
	"golang-practice/controllers"
	"golang-practice/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.ConnectDB()
	route.GET("/addresses/validate", controllers.GetAddressById)
	route.Run()
}
