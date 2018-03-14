package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	// "./app/src/utils"
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
	

	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./app/config")

	err := viper.ReadInConfig()
	if err != nil {
		// log.Println("Config file not found...")
		log.Fatalln(err)
	}

	log.Println(viper.GetInt("general.active"))

	// utils.InitLoggers()

	// utils.RedisConnect()

	// router := getMainEngine()

	// router.Run(":8080")

	

}
