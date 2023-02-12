package main

import (
	"github.com/ddiox/task-5-vix-btpns-Glenn_Claudio/database"
	"github.com/ddiox/task-5-vix-btpns-Glenn_Claudio/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
