package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Security(c *fiber.Ctx) error {
	log.Info(c.GetReqHeaders()["User-Agent"])

	// In another universe where security matters, here we get auth headers and verify them e.t.c
	return c.Next()
}
