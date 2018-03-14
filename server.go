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

	utils.InitLoggers()

	router := getMainEngine()

	router.Run(":8080")

	
	// c := utils.RedisConnect()
	// log.Println(c)
}
