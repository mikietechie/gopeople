package config

import (
	"strconv"

	"github.com/gofiber/fiber/v2/log"
)

func SetupLogging() {
	loggingLevel, err := strconv.Atoi(LOGGING_LEVEL)
	if err != nil {
		panic(err)
	}
	log.SetLevel(log.Level(loggingLevel))
	log.SetOutput(LOGGING_WRITER)
}
