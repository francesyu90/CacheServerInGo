package main

import (
	"log"
	"github.com/gin-gonic/gin"

	"./app/src/utils"
)

func testing(c *gin.Context) {
	log.Println("Hello World")
}

func getMainEngine() *gin.Engine {

	router := gin.Default()
	api := router.Group("/api") 
	{
		api.GET("/test", testing)
	}

	return router
}



func main() {
	

	v := utils.ReadInConfig()

	log.Println(v.GetInt("general.active"))

	// utils.InitLoggers()

	// utils.RedisConnect()

	// router := getMainEngine()

	// router.Run(":8080")

	

}
