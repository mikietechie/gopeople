package main

import (
	"github.com/mikietechie/gopeople/api/app"
	"github.com/mikietechie/gopeople/config"
	"github.com/mikietechie/gopeople/models"
)

func main() {
	models.PgConnect()
	defer models.PgDisconnect()
	app := app.New()
	app.Listen(config.SERVER_ADDRESS)
}
