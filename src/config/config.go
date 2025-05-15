package config

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var _ = godotenv.Overload("../.insecure.env")
var _ = godotenv.Overload("../.env")
var Ctx = context.Background()

/* System */
var (
	DEVMODE          = GetEnvOrDef("ENV", "PROD") == "DEV"
	HOST             = GetEnvOrDef("HOST", "localhost")
	PORT             = GetEnvOrDef("PORT", "8000")
	DB_LOGGING_LEVEL = GetEnvOrDef("DB_LOGGING_LEVEL", "3")
	LOGGING_LEVEL    = GetEnvOrDef("LOGGING_LEVEL", "0")
	LOGGING_FILE     = GetEnvOrDef("LOGGING_FILE", "logs.log")
)
var SERVER_ADDRESS = fmt.Sprintf("%s:%s", HOST, PORT)
var file, _ = os.OpenFile(LOGGING_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
var LOGGING_WRITER = io.MultiWriter(os.Stdout, file)

/* POSTGRES */
var (
	pg_user     = GetEnvOrDef("POSTGRES_USER", "Db")
	pg_password = GetEnvOrDef("POSTGRES_PASSWORD", "password")
	pg_host     = GetEnvOrDef("POSTGRES_HOST", "0.0.0.0")
	pg_port     = GetEnvOrDef("POSTGRES_PORT", "5432")
	pg_db_name  = GetEnvOrDef("POSTGRES_DB", "gopeople")
)
var DATABASE_CONNECTION = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pg_user, pg_password, pg_host, pg_port, pg_db_name)

var (
	GENDER_API_URL = GetEnvOrDef("GENDER_API_URL", `https://api.genderize.io/?name=%s`)
	NATION_API_URL = GetEnvOrDef("NATION_API_URL", `https://api.nationalize.io/?name=%s`)
	AGE_API_URL    = GetEnvOrDef("AGE_API_URL", `https://api.agify.io/?name=%s`)
)

var START_TIME = time.Now()

func init() {
	SetupLogging()
}
