package main

import (
	"github.com/yummycupcakes/task-5-vix-btpns-gilly/database"
	"github.com/yummycupcakes/task-5-vix-btpns-gilly/router"
)

func main() {
	database.DatabaseConnection()
	r := router.StartApp()
	r.Run(":8080")
}