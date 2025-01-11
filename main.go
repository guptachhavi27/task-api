package main

import (
    
	"task-api/database"
	"task-api/routes"
    "task-api/configs"
)


func main() {
    configs.LoadEnv()
	database.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run(":8080") // Start the server
}