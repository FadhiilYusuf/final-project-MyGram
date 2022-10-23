package main

import (
	"fmt"
	"mygram/config"
	"mygram/database"
	"mygram/router"
)

func main() {
	r := router.StartApp()
	err := database.StartDB()
	if err != nil {
		fmt.Println("Error starting database: ", err)
		return
	}
	r.Run(config.SERVER_PORT)
}
