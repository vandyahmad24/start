package main

import (
	"test-start/config"
	"test-start/router"
)

func main() {
	config.InitDB()
	router.Router()
}
