package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mikietechie/gopeople/config"
	"github.com/mikietechie/gopeople/tp"
)

func Index(c *fiber.Ctx) error {
	var data tp.IndexResData
	data.Start = config.START_TIME
	data.Current = time.Now()
	return c.JSON(data)
}
