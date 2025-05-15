package main

import (
	"github.com/mikietechie/gopeople/models"
)

func main() {
	models.PgConnect()
	defer models.PgDisconnect()
}
