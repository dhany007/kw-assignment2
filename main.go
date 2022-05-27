package main

import (
	"assignment2/databases"
	"assignment2/routers"
)

func main() {
	db := databases.ConnectDB()

	port := ":3000"
	router := routers.NewRouter(db)

	router.Run(port)
}
