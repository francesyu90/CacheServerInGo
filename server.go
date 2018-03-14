package main

import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"

	"./app/src/utils"
)

func testing(c *gin.Context) {
	log.Println("Hello World")
}

func getMainEngine(u *utils.Utilities) *gin.Engine {
	router := gin.Default()
	api := router.Group(u.GetStringConfigValue("api.api")) 
	{
		api.GET(u.GetStringConfigValue("qa.api.test"), testing)
	}

	return router
}



func main() {
	
	v := utils.ReadInConfig()

	u := utils.GetUtilities(v)

	utils.InitLoggers(u)

	utils.RedisConnect(u)

	router := getMainEngine(u)

	portStr := fmt.Sprintf(":%d", u.GetIntConfigValue("general.port"))
	router.Run(portStr)

}
