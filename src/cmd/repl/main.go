package main

import (
	"fmt"

	"github.com/mikietechie/gopeople/models"
	"github.com/mikietechie/gopeople/utils"
)

func main() {
	models.PgConnect()
	defer models.PgDisconnect()
	fmt.Println(utils.GetNationalityByName("Mike"))
	fmt.Println(utils.GetAgeByName("Mike"))
	fmt.Println(utils.GetGenderByName("Mike"))
}
