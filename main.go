package main

import (
	"goServer/db"
	"goServer/routers"
)

func main() {
	db.InitDb()
	router := routers.SetupRouter()
	router.Run("localhost:8080")
}
