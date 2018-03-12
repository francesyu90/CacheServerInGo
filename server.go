package main

import (
	"log"

	"./app/src/utils"
)

func main() {

	utils.InitLoggers()
	c := utils.RedisConnect()
	log.Println(c)
}
