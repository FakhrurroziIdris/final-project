package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var config = Config{}

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.WebServer.Port = os.Getenv("PORT")

	config.CRON.SensorInterval = os.Getenv("SENSOR_INTERVAL")

	config.WebPresentation.TemplateFolder = os.Getenv("TEMPLATE_WEB_FOLDER")
	config.WebPresentation.PageRefresh, _ = strconv.Atoi(os.Getenv("PAGE_REFRESH"))

	config.Database.JsonFile = os.Getenv("DB_JSON_FILE")
	config.Database.Host = os.Getenv("PGSQL_HOST")
	config.Database.Name = os.Getenv("PGSQL_NAME")
	config.Database.Password = os.Getenv("PGSQL_PASSWORD")
	config.Database.Port = os.Getenv("PGSQL_PORT")
	config.Database.User = os.Getenv("PGSQL_USER")

	config.Authentication.ExpiresInMinute, _ = strconv.Atoi(os.Getenv("JWT_LIFE"))
}

func GetEnv() Config {
	return config
}
