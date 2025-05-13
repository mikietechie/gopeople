package config

import (
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

var _ = godotenv.Overload("../.insecure.env")
var _ = godotenv.Overload("../.env")
var Ctx = context.Background()

/* System */
var (
	DEVMODE = GetEnvOrDef("ENV", "PROD") == "DEV"
	HOST    = GetEnvOrDef("HOST", "localhost")
	PORT    = GetEnvOrDef("PORT", "8000")
)
var SERVER_ADDRESS = fmt.Sprintf("%s:%s", HOST, PORT)

/* POSTGRES */
var (
	pg_user     = GetEnvOrDef("POSTGRES_USER", "Db")
	pg_password = GetEnvOrDef("POSTGRES_PASSWORD", "password")
	pg_host     = GetEnvOrDef("POSTGRES_HOST", "0.0.0.0")
	pg_port     = GetEnvOrDef("POSTGRES_PORT", "5432")
	pg_db_name  = GetEnvOrDef("POSTGRES_DB", "gopeople")
)
var DATABASE_CONNECTION = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pg_user, pg_password, pg_host, pg_port, pg_db_name)
var START_TIME = time.Now()

func init() {}
