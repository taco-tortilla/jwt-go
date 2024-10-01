package main

import (
	"github.com/taco-tortilla/jwt-go/initializers"
	"github.com/taco-tortilla/jwt-go/server"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	server.Init()
}
