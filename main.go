package main

import (
	"im/init"
	router2 "im/router"
)

func main() {

	init.InitLogger()

	init.InitConfig()
	
	init.InitDB()
	router := router2.Router()
	router.Run(":8080")
}
