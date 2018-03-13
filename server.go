package main

import (
	"log"
	// "time"

	"./app/src/utils"
)

func main() {

	utils.InitLoggers()
	
	// time.Sleep(5000 * time.Millisecond)
	c := utils.RedisConnect()
	log.Println(c)
}
