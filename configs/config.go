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
	config.Database.JsonFile = os.Getenv("DB_JSON_FILE")
	config.WebPresentation.PageRefresh, _ = strconv.Atoi(os.Getenv("PAGE_REFRESH"))
}

func GetEnv() Config {
	return config
}
