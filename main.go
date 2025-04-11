package main

import (
	"fmt"

	"github.com/FragsterDev/Go-Todo-API/config"
	"github.com/FragsterDev/Go-Todo-API/database"
)

func main() {

	cnfg := config.LoadConfig()

	db := database.ConnectDB(cnfg.MONGO_URI, cnfg.DBNAME)

	fmt.Printf("Connected to Database: %v", db.Name())
}
