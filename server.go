package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"./app/src/utils"
)

func testing(c *gin.Context) {
	log.Println("Hello World")
}

func getMainEngine(v *viper.Viper) *gin.Engine {

	router := gin.Default()
	api := router.Group(v.GetString("api.api")) 
	{
		api.GET(v.GetString("qa.api.test"), testing)
	}

	return router
}



func main() {
	
	v := utils.ReadInConfig()

	u := utils.GetUtilities(v)

	// utils.InitLoggers()

	utils.RedisConnect(u)

	// router := getMainEngine(v)
	// router.Run(":8080")

	

}
