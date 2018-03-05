package routing

import (
	"belajar/bot/controllers"

	"github.com/gin-gonic/gin"
)

func BootRouting() {
	router := gin.Default()

	router.POST("/callback", controllers.Callback)

	router.Run()
}
