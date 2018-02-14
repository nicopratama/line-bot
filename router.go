package main

import (
	"belajar/bot/controllers"

	"github.com/gin-gonic/gin"
)

func StartRouting() {
	router := gin.Default()

	router.POST("/callback", controllers.Callback)

	router.Run()
}
