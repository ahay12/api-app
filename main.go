package main

import (
	"api-apps/database"
	"log"
)

func main() {
	database.InitDatabase()

	log.Fatal(app.Listen(":4000"))
}
